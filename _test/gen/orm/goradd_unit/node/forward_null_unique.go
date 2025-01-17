// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// ForwardNullUniqueNodeI is the builder interface to the ForwardNullUnique nodes.
type ForwardNullUniqueNodeI interface {
	query.Node
	PrimaryKeyNode() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
	// ReverseID represents the reverse_id column in the database.
	ReverseID() *query.ColumnNode
	// Reverse represents the Reverse reference to a Reverse object.
	Reverse() ReverseNodeI
}

// ForwardNullUniqueExpander is the builder interface for ForwardNullUniques that are expandable.
type ForwardNullUniqueExpander interface {
	ForwardNullUniqueNodeI
	// Expand causes the node to produce separate rows with individual items, rather than a single row with an array of items.
	Expand()
}

// forwardNullUniqueTable represents the forward_null_unique table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the forwardNullUniqueTable, call [ForwardNullUnique()] to start a reference chain when querying the forward_null_unique table.
type forwardNullUniqueTable struct {
}

type forwardNullUniqueReverse struct {
	forwardNullUniqueTable
	query.ReverseNode
}

// ForwardNullUnique returns a table node that starts a node chain that begins with the forward_null_unique table.
func ForwardNullUnique() ForwardNullUniqueNodeI {
	return forwardNullUniqueTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n forwardNullUniqueTable) TableName_() string {
	return "forward_null_unique"
}

// NodeType_ returns the query.NodeType of the node.
func (n forwardNullUniqueTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n forwardNullUniqueTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ is used internally by the framework to return the list of all the column nodes.
// This may include reference nodes to enum types.
func (n forwardNullUniqueTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.ReverseID())
	return nodes
}

func (n *forwardNullUniqueReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.forwardNullUniqueTable.ColumnNodes_()
	for _, cn := range nodes {
		cn.(query.Linker).SetParent(n)
	}
	return
}

// Columns_ is used internally by the framework to return the list of all the columns in the table.
func (n forwardNullUniqueTable) Columns_() []string {
	return []string{
		"id",
		"name",
		"reverse_id",
	}
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n forwardNullUniqueTable) IsEnum_() bool {
	return false
}

func (n *forwardNullUniqueReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKeyNode returns a node that points to the primary key column.
func (n forwardNullUniqueTable) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n *forwardNullUniqueReverse) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n forwardNullUniqueTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardNullUniqueReverse) ID() *query.ColumnNode {
	cn := n.forwardNullUniqueTable.ID()
	cn.SetParent(n)
	return cn
}

func (n forwardNullUniqueTable) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardNullUniqueReverse) Name() *query.ColumnNode {
	cn := n.forwardNullUniqueTable.Name()
	cn.SetParent(n)
	return cn
}

func (n forwardNullUniqueTable) ReverseID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "reverse_id",
		Identifier:   "ReverseID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardNullUniqueReverse) ReverseID() *query.ColumnNode {
	cn := n.forwardNullUniqueTable.ReverseID()
	cn.SetParent(n)
	return cn
}

// Reverse represents the link to a Reverse object.
func (n forwardNullUniqueTable) Reverse() ReverseNodeI {
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

func (n *forwardNullUniqueReverse) Reverse() ReverseNodeI {
	cn := n.forwardNullUniqueTable.Reverse().(*reverseReference)
	cn.SetParent(n)
	return cn
}

func (n forwardNullUniqueTable) GobEncode() (data []byte, err error) {
	return
}

func (n *forwardNullUniqueTable) GobDecode(data []byte) (err error) {
	return
}

func (n *forwardNullUniqueReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *forwardNullUniqueReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(forwardNullUniqueTable))
	gob.Register(new(forwardNullUniqueReverse))
}
