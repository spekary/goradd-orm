// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// MilestoneNode is the builder interface to the Milestone nodes.
type MilestoneNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// ProjectID represents the project_id column in the database.
	ProjectID() *query.ColumnNode
	// Project represents the Project reference to a Project object.
	Project() ProjectNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
}

// milestoneTable represents the milestone table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the milestoneTable, call [Milestone()] to start a reference chain when querying the milestone table.
type milestoneTable struct {
}

type milestoneReverse struct {
	milestoneTable
	query.ReverseNode
}

// Milestone returns a table node that starts a node chain that begins with the milestone table.
func Milestone() MilestoneNode {
	return milestoneTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n milestoneTable) TableName_() string {
	return "milestone"
}

// NodeType_ returns the query.NodeType of the node.
func (n milestoneTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n milestoneTable) DatabaseKey_() string {
	return "goradd"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n milestoneTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.ProjectID())
	nodes = append(nodes, n.Name())
	return nodes
}

func (n *milestoneReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.milestoneTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

func (n *milestoneReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKey returns a node that points to the primary key column.
func (n milestoneTable) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *milestoneReverse) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

func (n milestoneTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *milestoneReverse) ID() *query.ColumnNode {
	cn := n.milestoneTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n milestoneTable) ProjectID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "project_id",
		Identifier:   "ProjectID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *milestoneReverse) ProjectID() *query.ColumnNode {
	cn := n.milestoneTable.ProjectID()
	query.NodeSetParent(cn, n)
	return cn
}

// Project represents the link to a Project object.
func (n milestoneTable) Project() ProjectNode {
	cn := &projectReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "project_id",
			Identifier:      "Project",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *milestoneReverse) Project() ProjectNode {
	cn := n.milestoneTable.Project().(*projectReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n milestoneTable) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *milestoneReverse) Name() *query.ColumnNode {
	cn := n.milestoneTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n milestoneTable) GobEncode() (data []byte, err error) {
	return
}

func (n *milestoneTable) GobDecode(data []byte) (err error) {
	return
}

func (n *milestoneReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *milestoneReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(milestoneTable))
	gob.Register(new(milestoneReverse))
}
