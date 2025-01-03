// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"github.com/goradd/orm/pkg/query"
)

// PersonNode represents the person table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the PersonNode, call [Person] to start a reference chain when querying the person table.
type PersonNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the personNode functions here.
	query.ReferenceNodeI
}

// Person returns a table node that starts a node chain that begins with the person table.
func Person() *PersonNode {
	n := PersonNode{
		query.NewTableNode("goradd", "person", "Person"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *PersonNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.FirstName())
	nodes = append(nodes, n.LastName())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *PersonNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *PersonNode) Copy_() query.NodeI {
	return &PersonNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *PersonNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *PersonNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"person",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// FirstName represents the first_name column in the database.
func (n *PersonNode) FirstName() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"person",
		"first_name",
		"FirstName",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// LastName represents the last_name column in the database.
func (n *PersonNode) LastName() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"person",
		"last_name",
		"LastName",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// PersonTypes represents the many-to-many relationship formed by the person_persontype_assn table.
func (n *PersonNode) PersonTypes() *PersonTypeNode {
	cn := &PersonTypeNode{
		query.NewManyManyNode(
			"goradd",
			"person_persontype_assn",
			"person_id",
			"PersonTypes",
			"",
			"person_type_id",
			"id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

// Projects represents the many-to-many relationship formed by the team_member_project_assn table.
func (n *PersonNode) Projects() *ProjectNode {
	cn := &ProjectNode{
		query.NewManyManyNode(
			"goradd",
			"team_member_project_assn",
			"team_member_id",
			"Projects",
			"project",
			"project_id",
			"id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}
