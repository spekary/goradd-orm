// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"github.com/goradd/orm/pkg/query"
)

// ForwardCascadeNode represents the forward_cascade table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the ForwardCascadeNode, call [ForwardCascade] to start a reference chain when querying the forward_cascade table.
type ForwardCascadeNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the forwardCascadeNode functions here.
	query.ReferenceNodeI
}

// ForwardCascade returns a table node that starts a node chain that begins with the forward_cascade table.
func ForwardCascade() *ForwardCascadeNode {
	n := ForwardCascadeNode{
		query.NewTableNode("goradd_unit", "forward_cascade", "ForwardCascade"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *ForwardCascadeNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.ReverseID())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *ForwardCascadeNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *ForwardCascadeNode) Copy_() query.NodeI {
	return &ForwardCascadeNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *ForwardCascadeNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *ForwardCascadeNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_cascade",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Name represents the name column in the database.
func (n *ForwardCascadeNode) Name() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_cascade",
		"name",
		"Name",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ReverseID represents the reverse_id column in the database.
func (n *ForwardCascadeNode) ReverseID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_cascade",
		"reverse_id",
		"ReverseID",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Reverse represents the link to a Reverse object.
func (n *ForwardCascadeNode) Reverse() *ReverseNode {
	cn := &ReverseNode{
		query.NewReferenceNode(
			"goradd_unit",
			"forward_cascade",
			"reverse_id",
			"ReverseID",
			"Reverse",
			"reverse",
			"id",
			false,
			query.ColTypeString,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}
