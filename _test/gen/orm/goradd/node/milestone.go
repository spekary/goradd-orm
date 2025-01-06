// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// MilestoneNode represents the milestone table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the MilestoneNode, call [Milestone] to start a reference chain when querying the milestone table.
type MilestoneNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the milestoneNode functions here.
	query.ReferenceNodeI
}

// Milestone returns a table node that starts a node chain that begins with the milestone table.
func Milestone() *MilestoneNode {
	n := MilestoneNode{
		query.NewTableNode("goradd", "milestone", "Milestone"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *MilestoneNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.ProjectID())
	nodes = append(nodes, n.Name())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *MilestoneNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *MilestoneNode) Copy_() query.NodeI {
	return &MilestoneNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *MilestoneNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *MilestoneNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"milestone",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ProjectID represents the project_id column in the database.
func (n *MilestoneNode) ProjectID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"milestone",
		"project_id",
		"ProjectID",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Project represents the link to a Project object.
func (n *MilestoneNode) Project() *ProjectNode {
	cn := &ProjectNode{
		query.NewReferenceNode(
			"goradd",
			"milestone",
			"project_id",
			"ProjectID",
			"Project",
			"project",
			"id",
			false,
			query.ColTypeString,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

// Name represents the name column in the database.
func (n *MilestoneNode) Name() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"milestone",
		"name",
		"Name",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

type milestoneNodeEncoded struct {
	RefNode query.ReferenceNodeI
}

// GobEncode makes the node serializable.
// doc:hide
func (n *MilestoneNode) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	s := milestoneNodeEncoded{
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
func (n *MilestoneNode) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	var s milestoneNodeEncoded
	if err = dec.Decode(&s); err != nil {
		panic(err)
	}
	n.ReferenceNodeI = s.RefNode
	query.SetParentNode(n, query.ParentNode(n)) // Reinforce types
	return
}
func init() {
	gob.Register(&MilestoneNode{})
}
