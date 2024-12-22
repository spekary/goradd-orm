// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"github.com/goradd/orm/pkg/query"
)

// DoubleIndexNode represents the double_index table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the DoubleIndexNode, call [DoubleIndex] to start a reference chain when querying the double_index table.
type DoubleIndexNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the doubleIndexNode functions here.
	query.ReferenceNodeI
}

// DoubleIndex returns a table node that starts a node chain that begins with the double_index table.
func DoubleIndex() *DoubleIndexNode {
	n := DoubleIndexNode{
		query.NewTableNode("goradd_unit", "double_index", "DoubleIndex"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *DoubleIndexNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.FieldInt())
	nodes = append(nodes, n.FieldString())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *DoubleIndexNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *DoubleIndexNode) Copy_() query.NodeI {
	return &DoubleIndexNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *DoubleIndexNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *DoubleIndexNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"double_index",
		"id",
		"ID",
		query.ColTypeInteger,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// FieldInt represents the field_int column in the database.
func (n *DoubleIndexNode) FieldInt() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"double_index",
		"field_int",
		"FieldInt",
		query.ColTypeInteger,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// FieldString represents the field_string column in the database.
func (n *DoubleIndexNode) FieldString() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"double_index",
		"field_string",
		"FieldString",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}
