// Code generated by GoRADD. DO NOT EDIT.

package node

import (
	"github.com/goradd/orm/pkg/query"
)

// ForwardNullNode represents the forward_null table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the ForwardNullNode, call [ForwardNull] to start a reference chain when querying the forward_null table.
type ForwardNullNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the forwardNullNode functions here.
	query.ReferenceNodeI
}

// ForwardNull returns a table node that starts a node chain that begins with the forward_null table.
func ForwardNull() *ForwardNullNode {
	n := ForwardNullNode{
		query.NewTableNode("goradd_unit", "forward_null", "ForwardNull"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *ForwardNullNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.ReverseID())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *ForwardNullNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *ForwardNullNode) Copy_() query.NodeI {
	return &ForwardNullNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *ForwardNullNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *ForwardNullNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_null",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Name represents the name column in the database.
func (n *ForwardNullNode) Name() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_null",
		"name",
		"Name",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ReverseID represents the reverse_id column in the database.
func (n *ForwardNullNode) ReverseID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_null",
		"reverse_id",
		"ReverseID",
		query.ColTypeInteger,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Reverse represents the link to a Reverse object.
func (n *ForwardNullNode) Reverse() *ReverseNode {
	cn := &ReverseNode{
		query.NewReferenceNode(
			"goradd_unit",
			"forward_null",
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