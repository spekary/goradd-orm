// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"github.com/goradd/orm/pkg/query"
)

// ForwardRestrictNode represents the forward_restrict table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the ForwardRestrictNode, call [ForwardRestrict] to start a reference chain when querying the forward_restrict table.
type ForwardRestrictNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the forwardRestrictNode functions here.
	query.ReferenceNodeI
}

// ForwardRestrict returns a table node that starts a node chain that begins with the forward_restrict table.
func ForwardRestrict() *ForwardRestrictNode {
	n := ForwardRestrictNode{
		query.NewTableNode("goradd_unit", "forward_restrict", "ForwardRestrict"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *ForwardRestrictNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.ReverseID())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *ForwardRestrictNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *ForwardRestrictNode) Copy_() query.NodeI {
	return &ForwardRestrictNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *ForwardRestrictNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *ForwardRestrictNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_restrict",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Name represents the name column in the database.
func (n *ForwardRestrictNode) Name() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_restrict",
		"name",
		"Name",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ReverseID represents the reverse_id column in the database.
func (n *ForwardRestrictNode) ReverseID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_restrict",
		"reverse_id",
		"ReverseID",
		query.ColTypeInteger,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Reverse represents the link to a Reverse object.
func (n *ForwardRestrictNode) Reverse() *ReverseNode {
	cn := &ReverseNode{
		query.NewReferenceNode(
			"goradd_unit",
			"forward_restrict",
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
