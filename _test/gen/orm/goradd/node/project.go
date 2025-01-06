// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// ProjectNode represents the project table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the ProjectNode, call [Project] to start a reference chain when querying the project table.
type ProjectNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the projectNode functions here.
	query.ReferenceNodeI
}

// Project returns a table node that starts a node chain that begins with the project table.
func Project() *ProjectNode {
	n := ProjectNode{
		query.NewTableNode("goradd", "project", "Project"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *ProjectNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Num())
	nodes = append(nodes, n.Status())
	nodes = append(nodes, n.ManagerID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.Description())
	nodes = append(nodes, n.StartDate())
	nodes = append(nodes, n.EndDate())
	nodes = append(nodes, n.Budget())
	nodes = append(nodes, n.Spent())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *ProjectNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *ProjectNode) Copy_() query.NodeI {
	return &ProjectNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *ProjectNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *ProjectNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Num represents the num column in the database.
func (n *ProjectNode) Num() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"num",
		"Num",
		query.ColTypeInteger,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Status represents the status_id column in the database.
func (n *ProjectNode) Status() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"status_id",
		"Status",
		query.ColTypeInteger,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ManagerID represents the manager_id column in the database.
func (n *ProjectNode) ManagerID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"manager_id",
		"ManagerID",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Manager represents the link to a Person object.
func (n *ProjectNode) Manager() *PersonNode {
	cn := &PersonNode{
		query.NewReferenceNode(
			"goradd",
			"project",
			"manager_id",
			"ManagerID",
			"Manager",
			"person",
			"id",
			false,
			query.ColTypeString,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

// Name represents the name column in the database.
func (n *ProjectNode) Name() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"name",
		"Name",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Description represents the description column in the database.
func (n *ProjectNode) Description() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"description",
		"Description",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// StartDate represents the start_date column in the database.
func (n *ProjectNode) StartDate() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"start_date",
		"StartDate",
		query.ColTypeTime,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// EndDate represents the end_date column in the database.
func (n *ProjectNode) EndDate() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"end_date",
		"EndDate",
		query.ColTypeTime,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Budget represents the budget column in the database.
func (n *ProjectNode) Budget() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"budget",
		"Budget",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Spent represents the spent column in the database.
func (n *ProjectNode) Spent() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"project",
		"spent",
		"Spent",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Children represents the many-to-many relationship formed by the related_project_assn table.
func (n *ProjectNode) Children() *ProjectNode {
	cn := &ProjectNode{
		query.NewManyManyNode(
			"goradd",
			"related_project_assn",
			"parent_id",
			"Children",
			"project",
			"child_id",
			"id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

// Parents represents the many-to-many relationship formed by the related_project_assn table.
func (n *ProjectNode) Parents() *ProjectNode {
	cn := &ProjectNode{
		query.NewManyManyNode(
			"goradd",
			"related_project_assn",
			"child_id",
			"Parents",
			"project",
			"parent_id",
			"id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

// TeamMembers represents the many-to-many relationship formed by the team_member_project_assn table.
func (n *ProjectNode) TeamMembers() *PersonNode {
	cn := &PersonNode{
		query.NewManyManyNode(
			"goradd",
			"team_member_project_assn",
			"project_id",
			"TeamMembers",
			"person",
			"team_member_id",
			"id",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

// Milestones represents the many-to-one relationship formed by the reverse reference from the
// project_id column in the milestone table.
func (n *ProjectNode) Milestones() *MilestoneNode {
	cn := &MilestoneNode{
		query.NewReverseReferenceNode(
			"goradd",
			"project",
			"id",
			"Milestones",
			"milestone",
			"project_id",
			true,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

type projectNodeEncoded struct {
	RefNode query.ReferenceNodeI
}

// GobEncode makes the node serializable.
// doc:hide
func (n *ProjectNode) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	s := projectNodeEncoded{
		RefNode: n.ReferenceNodeI,
	}

	if err = e.Encode(s); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

// GobDecode makes the node deserializable.
// doc: hide
func (n *ProjectNode) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	var s projectNodeEncoded
	if err = dec.Decode(&s); err != nil {
		panic(err)
	}
	n.ReferenceNodeI = s.RefNode
	query.SetParentNode(n, query.ParentNode(n)) // Reinforce types
	return
}
func init() {
	gob.Register(&ProjectNode{})
}
