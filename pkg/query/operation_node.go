package query

import (
	"bytes"
	"encoding/gob"
	"log/slog"
	"strings"
)

type OperationNodeI interface {
	Node
	container
}

// Operator is used internally by the framework to specify an operation to be performed by the database.
// Not all databases can perform all the operations. It will be up to the database driver to sort this out.
type Operator string

const (
	// Standard logical operators

	OpEqual        Operator = "="
	OpNotEqual     Operator = "<>"
	OpAnd          Operator = "AND"
	OpOr           Operator = "OR"
	OpXor          Operator = "XOR"
	OpGreater      Operator = ">"
	OpGreaterEqual Operator = ">="
	OpLess         Operator = "<"
	OpLessEqual    Operator = "<="

	// Unary logical
	OpNot  Operator = "NOT"
	OpAll  Operator = "1=1"
	OpNone Operator = "1=0"

	// Math operators
	OpAdd      Operator = "+"
	OpSubtract Operator = "-"
	OpMultiply Operator = "*"
	OpDivide   Operator = "/"
	OpModulo   Operator = "%"

	// Unary math
	OpNegate Operator = " -"

	// Bit operators
	OpBitAnd     Operator = "&"
	OpBitOr      Operator = "|"
	OpBitXor     Operator = "^"
	OpShiftLeft  Operator = "<<"
	OpShiftRight Operator = ">>"

	// Unary bit
	OpBitInvert Operator = "~"

	// Function operator
	// The function name is followed by the operators in parenthesis
	OpFunc Operator = "func"

	// SQL functions that act like operators in that the operator is put in between the operands
	OpLike  Operator = "LIKE" // This is very SQL specific and may not be supported in NoSql
	OpIn    Operator = "IN"
	OpNotIn Operator = "NOT IN"

	// Special NULL tests
	OpNull    Operator = "NULL"
	OpNotNull Operator = "NOT NULL"

	// Our own custom operators for universal support
	OpStartsWith     Operator = "StartsWith"
	OpEndsWith       Operator = "EndsWith"
	OpContains       Operator = "Contains"
	OpDateAddSeconds Operator = "AddSeconds" // Adds the given number of seconds to a datetime
)

// String returns a string representation of the Operator type. For convenience, this also corresponds to the SQL
// representation of an operator
func (o Operator) String() string {
	return string(o)
}

// An OperationNode is a general purpose structure that specifies an operation on a node or group of nodes.
// The operation could be arithmetic, boolean, or a function.
type OperationNode struct {
	op           Operator
	operands     []Node
	functionName string // for function operations specific to the db driver
	distinct     bool   // some aggregate queries, particularly count, allow this inside the function
	isAggregate  bool
}

// NewOperationNode returns a new operation node.
func NewOperationNode(op Operator, operands ...interface{}) *OperationNode {
	n := &OperationNode{
		op: op,
	}
	n.assignOperands(operands...)
	return n
}

// NewFunctionNode returns an operation node that executes a database function.
func NewFunctionNode(functionName string, operands ...interface{}) *OperationNode {
	n := &OperationNode{
		op:           OpFunc,
		functionName: functionName,
	}
	n.assignOperands(operands...)
	return n
}

// NewAggregateFunctionNode returns an operation node that executes an aggregate function.
func NewAggregateFunctionNode(functionName string, operands ...interface{}) *OperationNode {
	n := &OperationNode{
		op:           OpFunc,
		functionName: functionName,
		isAggregate:  true,
	}
	n.assignOperands(operands...)
	return n
}

// NewCountNode creates a Count function node. If no operands are given, it will use * as the parameter to the function
// which means it will count nulls. To NOT count nulls, a node needs to be specified. Only up to one node can be specified.
func NewCountNode(operands ...Node) *OperationNode {
	n := &OperationNode{
		op:           OpFunc,
		functionName: "COUNT",
		isAggregate:  true,
	}
	// Note: Some SQLs like MySQL and Postgres allow multiple operands combined with DISTINCT.
	// However, others do not. We support only the universally accepted way.
	// There are workarounds.
	if len(operands) > 1 {
		panic("can only specify one operand in a count operation")
	}
	for _, op := range operands { // 0 or 1
		n.operands = append(n.operands, op)
	}

	return n
}

func (n *OperationNode) NodeType_() NodeType {
	return OperationNodeType
}

// assignOperands processes the list of operands at run time, making sure all static values are escaped.
func (n *OperationNode) assignOperands(operands ...interface{}) {
	var op interface{}

	if operands != nil {
		for _, op = range operands {
			if ni, ok := op.(Node); ok {
				n.operands = append(n.operands, ni)
			} else {
				n.operands = append(n.operands, NewValueNode(op))
			}
		}
	}
}

// Distinct sets the operation to return distinct results
func (n *OperationNode) Distinct() *OperationNode {
	n.distinct = true
	return n
}

func (n *OperationNode) containedNodes() (nodes []Node) {
	for _, op := range n.operands {
		if nc, ok := op.(container); ok {
			nodes = append(nodes, nc.containedNodes()...)
		} else {
			nodes = append(nodes, op)
		}
	}
	return
}

func (n *OperationNode) TableName_() string {
	return ""
}

func (n *OperationNode) DatabaseKey_() string {
	return ""
}

func (n *OperationNode) log(level int) {
	tabs := strings.Repeat("\t", level)
	slog.Debug(tabs + "Op: " + n.op.String())
}

func (n *OperationNode) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(n.op); err != nil {
		panic(err)
	}
	if err = e.Encode(n.operands); err != nil {
		panic(err)
	}
	if err = e.Encode(n.functionName); err != nil {
		panic(err)
	}
	if err = e.Encode(n.distinct); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *OperationNode) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	if err = dec.Decode(&n.op); err != nil {
		panic(err)
	}
	if err = dec.Decode(&n.operands); err != nil {
		panic(err)
	}
	if err = dec.Decode(&n.functionName); err != nil {
		panic(err)
	}
	if err = dec.Decode(&n.distinct); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(&OperationNode{})
}

// OperationNodeOperator is used internally by the framework to get the operator.
func OperationNodeOperator(n *OperationNode) Operator {
	return n.op
}

// OperationNodeOperands is used internally by the framework to get the operands.
func OperationNodeOperands(n *OperationNode) []Node {
	return n.operands
}

// OperationNodeFunction is used internally by the framework to get the function.
func OperationNodeFunction(n *OperationNode) string {
	return n.functionName
}

// OperationNodeDistinct is used internally by the framework to get the distinct value.
func OperationNodeDistinct(n *OperationNode) bool {
	return n.distinct
}

// NodeHasAggregate is used by the orm to detect if this node or its subnodes has an aggregate function.
func NodeHasAggregate(n Node) bool {
	if on, ok := n.(*OperationNode); ok {
		if on.isAggregate {
			return true
		}
		for _, op := range on.operands {
			if NodeHasAggregate(op) {
				return true
			}
		}
	} else if cn, ok2 := n.(container); ok2 {
		for _, n := range cn.containedNodes() {
			if NodeHasAggregate(n) {
				return true
			}
		}
	}
	return false
}
