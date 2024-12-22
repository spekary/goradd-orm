// Code generated by GoRADD. DO NOT EDIT.

package node

import (
	"github.com/goradd/orm/pkg/query"
)

// ForwardNullUniqueNode represents the forward_null_unique table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the ForwardNullUniqueNode, call [ForwardNullUnique] to start a reference chain when querying the forward_null_unique table.
type ForwardNullUniqueNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the forwardNullUniqueNode functions here.
	query.ReferenceNodeI
}

// ForwardNullUnique returns a table node that starts a node chain that begins with the forward_null_unique table.
func ForwardNullUnique() *ForwardNullUniqueNode {
	n := ForwardNullUniqueNode{
		query.NewTableNode("goradd_unit", "forward_null_unique", "ForwardNullUnique"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *ForwardNullUniqueNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.ReverseID())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *ForwardNullUniqueNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *ForwardNullUniqueNode) Copy_() query.NodeI {
	return &ForwardNullUniqueNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *ForwardNullUniqueNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *ForwardNullUniqueNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_null_unique",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Name represents the name column in the database.
func (n *ForwardNullUniqueNode) Name() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_null_unique",
		"name",
		"Name",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ReverseID represents the reverse_id column in the database.
func (n *ForwardNullUniqueNode) ReverseID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_null_unique",
		"reverse_id",
		"ReverseID",
		query.ColTypeInteger,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Reverse represents the link to a Reverse object.
func (n *ForwardNullUniqueNode) Reverse() *ReverseNode {
	cn := &ReverseNode{
		query.NewReferenceNode(
			"goradd_unit",
			"forward_null_unique",
			"reverse_id",
			"ReverseID",
			"Reverse",
			"reverse",
			"id",
			false,
			query.ColTypeInteger,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}