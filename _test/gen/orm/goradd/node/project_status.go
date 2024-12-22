// Code generated by goradd. DO NOT EDIT.

package node

import (
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

type ProjectStatusNode struct {
	query.ReferenceNodeI
}

// PrimaryKeyNode returns a node representing the primary key column.
func (n *ProjectStatusNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func init() {
	gob.Register(new(ProjectStatusNode))
}

// SelectNodes_ is used internally by the framework to return the list of column nodes.
// doc: hide
func (n *ProjectStatusNode) SelectNodes_() []*query.ColumnNode {
	return []*query.ColumnNode{
		n.ID(),
		n.Name(),
		n.Description(),
		n.Guidelines(),
		n.IsActive(),
	}
}

// EmbeddedNode_ is used internally by the framework to return the embedded ReferenceNodeI.
// doc: hide
func (n *ProjectStatusNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *ProjectStatusNode) Copy_() query.NodeI {
	return &ProjectStatusNode{query.CopyNode(n.ReferenceNodeI)}
}

func (n *ProjectStatusNode) ID() *query.ColumnNode {

	cn := query.NewColumnNode(
		"goradd",
		"project_status_enum",
		"id",
		"ID",
		query.ColTypeInteger,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

func (n *ProjectStatusNode) Name() *query.ColumnNode {

	cn := query.NewColumnNode(
		"goradd",
		"project_status_enum",
		"name",
		"Name",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

func (n *ProjectStatusNode) Description() *query.ColumnNode {

	cn := query.NewColumnNode(
		"goradd",
		"project_status_enum",
		"description",
		"Description",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

func (n *ProjectStatusNode) Guidelines() *query.ColumnNode {

	cn := query.NewColumnNode(
		"goradd",
		"project_status_enum",
		"guidelines",
		"Guidelines",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

func (n *ProjectStatusNode) IsActive() *query.ColumnNode {

	cn := query.NewColumnNode(
		"goradd",
		"project_status_enum",
		"is_active",
		"IsActive",
		query.ColTypeBool,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}
