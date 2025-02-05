// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// ForwardCascadeNode is the builder interface to the ForwardCascade nodes.
type ForwardCascadeNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
	// ReverseID represents the reverse_id column in the database.
	ReverseID() *query.ColumnNode
	// Reverse represents the Reverse reference to a Reverse object.
	Reverse() ReverseNode
}

// forwardCascadeTable represents the forward_cascade table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the forwardCascadeTable, call [ForwardCascade()] to start a reference chain when querying the forward_cascade table.
type forwardCascadeTable struct {
}

type forwardCascadeReverse struct {
	forwardCascadeTable
	query.ReverseNode
}

// ForwardCascade returns a table node that starts a node chain that begins with the forward_cascade table.
func ForwardCascade() ForwardCascadeNode {
	return forwardCascadeTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n forwardCascadeTable) TableName_() string {
	return "forward_cascade"
}

// NodeType_ returns the query.NodeType of the node.
func (n forwardCascadeTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n forwardCascadeTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n forwardCascadeTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.ReverseID())
	return nodes
}

func (n *forwardCascadeReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.forwardCascadeTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n forwardCascadeTable) IsEnum_() bool {
	return false
}

func (n *forwardCascadeReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKey returns a node that points to the primary key column.
func (n forwardCascadeTable) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *forwardCascadeReverse) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

func (n forwardCascadeTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *forwardCascadeReverse) ID() *query.ColumnNode {
	cn := n.forwardCascadeTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n forwardCascadeTable) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *forwardCascadeReverse) Name() *query.ColumnNode {
	cn := n.forwardCascadeTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n forwardCascadeTable) ReverseID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "reverse_id",
		Identifier:   "ReverseID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *forwardCascadeReverse) ReverseID() *query.ColumnNode {
	cn := n.forwardCascadeTable.ReverseID()
	query.NodeSetParent(cn, n)
	return cn
}

// Reverse represents the link to a Reverse object.
func (n forwardCascadeTable) Reverse() ReverseNode {
	cn := &reverseReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "reverse_id",
			Identifier:      "Reverse",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *forwardCascadeReverse) Reverse() ReverseNode {
	cn := n.forwardCascadeTable.Reverse().(*reverseReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n forwardCascadeTable) GobEncode() (data []byte, err error) {
	return
}

func (n *forwardCascadeTable) GobDecode(data []byte) (err error) {
	return
}

func (n *forwardCascadeReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *forwardCascadeReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(forwardCascadeTable))
	gob.Register(new(forwardCascadeReverse))
}
