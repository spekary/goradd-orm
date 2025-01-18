// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// ForwardNullNode is the builder interface to the ForwardNull nodes.
type ForwardNullNode interface {
	query.Node
	PrimaryKeyNode() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
	// ReverseID represents the reverse_id column in the database.
	ReverseID() *query.ColumnNode
	// Reverse represents the Reverse reference to a Reverse object.
	Reverse() ReverseNode
}

// ForwardNullExpander is the builder interface for ForwardNulls that are expandable.
type ForwardNullExpander interface {
	ForwardNullNode
	// Expand causes the node to produce separate rows with individual items, rather than a single row with an array of items.
	Expand()
}

// forwardNullTable represents the forward_null table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the forwardNullTable, call [ForwardNull()] to start a reference chain when querying the forward_null table.
type forwardNullTable struct {
}

type forwardNullReverse struct {
	forwardNullTable
	query.ReverseNode
}

// ForwardNull returns a table node that starts a node chain that begins with the forward_null table.
func ForwardNull() ForwardNullNode {
	return forwardNullTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n forwardNullTable) TableName_() string {
	return "forward_null"
}

// NodeType_ returns the query.NodeType of the node.
func (n forwardNullTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n forwardNullTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ is used internally by the framework to return the list of all the column nodes.
// This may include reference nodes to enum types.
func (n forwardNullTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.ReverseID())
	return nodes
}

func (n *forwardNullReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.forwardNullTable.ColumnNodes_()
	for _, cn := range nodes {
		cn.(query.Linker).SetParent(n)
	}
	return
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n forwardNullTable) IsEnum_() bool {
	return false
}

func (n *forwardNullReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKeyNode returns a node that points to the primary key column.
func (n forwardNullTable) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n *forwardNullReverse) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n forwardNullTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardNullReverse) ID() *query.ColumnNode {
	cn := n.forwardNullTable.ID()
	cn.SetParent(n)
	return cn
}

func (n forwardNullTable) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardNullReverse) Name() *query.ColumnNode {
	cn := n.forwardNullTable.Name()
	cn.SetParent(n)
	return cn
}

func (n forwardNullTable) ReverseID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "reverse_id",
		Identifier:   "ReverseID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardNullReverse) ReverseID() *query.ColumnNode {
	cn := n.forwardNullTable.ReverseID()
	cn.SetParent(n)
	return cn
}

// Reverse represents the link to a Reverse object.
func (n forwardNullTable) Reverse() ReverseNode {
	cn := &reverseReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "reverse_id",
			Identifier:      "ReverseID",
			ReceiverType:    query.ColTypeString,
		},
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardNullReverse) Reverse() ReverseNode {
	cn := n.forwardNullTable.Reverse().(*reverseReference)
	cn.SetParent(n)
	return cn
}

func (n forwardNullTable) GobEncode() (data []byte, err error) {
	return
}

func (n *forwardNullTable) GobDecode(data []byte) (err error) {
	return
}

func (n *forwardNullReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *forwardNullReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(forwardNullTable))
	gob.Register(new(forwardNullReverse))
}
