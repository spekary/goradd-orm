// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// ProjectNode is the builder interface to the Project nodes.
type ProjectNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// Num represents the num column in the database.
	Num() *query.ColumnNode
	// Status represents the status_enum column in the database.
	Status() *query.ColumnNode
	// ManagerID represents the manager_id column in the database.
	ManagerID() *query.ColumnNode
	// Manager represents the Manager reference to a Person object.
	Manager() PersonNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
	// Description represents the description column in the database.
	Description() *query.ColumnNode
	// StartDate represents the start_date column in the database.
	StartDate() *query.ColumnNode
	// EndDate represents the end_date column in the database.
	EndDate() *query.ColumnNode
	// Budget represents the budget column in the database.
	Budget() *query.ColumnNode
	// Spent represents the spent column in the database.
	Spent() *query.ColumnNode
	// Children represents the Children reference to Project objects.
	Children() ProjectNode
	// Parents represents the Parents reference to Project objects.
	Parents() ProjectNode
	// TeamMembers represents the TeamMembers reference to Person objects.
	TeamMembers() PersonNode
	// Milestones represents the Milestones reference to Milestone objects.
	Milestones() MilestoneNode
}

// projectTable represents the project table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the projectTable, call [Project()] to start a reference chain when querying the project table.
type projectTable struct {
}

type projectReference struct {
	projectTable
	query.ReferenceNode
}

type projectReverse struct {
	projectTable
	query.ReverseNode
}

type projectAssociation struct {
	projectTable
	query.ManyManyNode
}

// Project returns a table node that starts a node chain that begins with the project table.
func Project() ProjectNode {
	return projectTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n projectTable) TableName_() string {
	return "project"
}

// NodeType_ returns the query.NodeType of the node.
func (n projectTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n projectTable) DatabaseKey_() string {
	return "goradd"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n projectTable) ColumnNodes_() (nodes []query.Node) {
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

func (n *projectReference) ColumnNodes_() (nodes []query.Node) {
	nodes = n.projectTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

func (n *projectReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.projectTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

func (n *projectAssociation) ColumnNodes_() (nodes []query.Node) {
	nodes = n.projectTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n projectTable) IsEnum_() bool {
	return false
}

func (n *projectReference) NodeType_() query.NodeType {
	return query.ReferenceNodeType
}

func (n *projectReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

func (n *projectAssociation) NodeType_() query.NodeType {
	return query.ManyManyNodeType
}

// PrimaryKey returns a node that points to the primary key column.
func (n projectTable) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *projectReference) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *projectReverse) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *projectAssociation) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

func (n projectTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) ID() *query.ColumnNode {
	cn := n.projectTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) ID() *query.ColumnNode {
	cn := n.projectTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) ID() *query.ColumnNode {
	cn := n.projectTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n projectTable) Num() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "num",
		Identifier:   "Num",
		ReceiverType: query.ColTypeInteger,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) Num() *query.ColumnNode {
	cn := n.projectTable.Num()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) Num() *query.ColumnNode {
	cn := n.projectTable.Num()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) Num() *query.ColumnNode {
	cn := n.projectTable.Num()
	query.NodeSetParent(cn, n)
	return cn
}

func (n projectTable) Status() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "status_enum",
		Identifier:   "Status",
		ReceiverType: query.ColTypeInteger,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) Status() *query.ColumnNode {
	cn := n.projectTable.Status()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) Status() *query.ColumnNode {
	cn := n.projectTable.Status()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) Status() *query.ColumnNode {
	cn := n.projectTable.Status()
	query.NodeSetParent(cn, n)
	return cn
}

func (n projectTable) ManagerID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "manager_id",
		Identifier:   "ManagerID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) ManagerID() *query.ColumnNode {
	cn := n.projectTable.ManagerID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) ManagerID() *query.ColumnNode {
	cn := n.projectTable.ManagerID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) ManagerID() *query.ColumnNode {
	cn := n.projectTable.ManagerID()
	query.NodeSetParent(cn, n)
	return cn
}

// Manager represents the link to a Person object.
func (n projectTable) Manager() PersonNode {
	cn := &personReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "manager_id",
			Identifier:      "Manager",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) Manager() PersonNode {
	cn := n.projectTable.Manager().(*personReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) Manager() PersonNode {
	cn := n.projectTable.Manager().(*personReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) Manager() PersonNode {
	cn := n.projectTable.Manager().(*personReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n projectTable) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) Name() *query.ColumnNode {
	cn := n.projectTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) Name() *query.ColumnNode {
	cn := n.projectTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) Name() *query.ColumnNode {
	cn := n.projectTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n projectTable) Description() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "description",
		Identifier:   "Description",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) Description() *query.ColumnNode {
	cn := n.projectTable.Description()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) Description() *query.ColumnNode {
	cn := n.projectTable.Description()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) Description() *query.ColumnNode {
	cn := n.projectTable.Description()
	query.NodeSetParent(cn, n)
	return cn
}

func (n projectTable) StartDate() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "start_date",
		Identifier:   "StartDate",
		ReceiverType: query.ColTypeTime,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) StartDate() *query.ColumnNode {
	cn := n.projectTable.StartDate()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) StartDate() *query.ColumnNode {
	cn := n.projectTable.StartDate()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) StartDate() *query.ColumnNode {
	cn := n.projectTable.StartDate()
	query.NodeSetParent(cn, n)
	return cn
}

func (n projectTable) EndDate() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "end_date",
		Identifier:   "EndDate",
		ReceiverType: query.ColTypeTime,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) EndDate() *query.ColumnNode {
	cn := n.projectTable.EndDate()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) EndDate() *query.ColumnNode {
	cn := n.projectTable.EndDate()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) EndDate() *query.ColumnNode {
	cn := n.projectTable.EndDate()
	query.NodeSetParent(cn, n)
	return cn
}

func (n projectTable) Budget() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "budget",
		Identifier:   "Budget",
		ReceiverType: query.ColTypeUnknown,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) Budget() *query.ColumnNode {
	cn := n.projectTable.Budget()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) Budget() *query.ColumnNode {
	cn := n.projectTable.Budget()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) Budget() *query.ColumnNode {
	cn := n.projectTable.Budget()
	query.NodeSetParent(cn, n)
	return cn
}

func (n projectTable) Spent() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "spent",
		Identifier:   "Spent",
		ReceiverType: query.ColTypeUnknown,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) Spent() *query.ColumnNode {
	cn := n.projectTable.Spent()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) Spent() *query.ColumnNode {
	cn := n.projectTable.Spent()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) Spent() *query.ColumnNode {
	cn := n.projectTable.Spent()
	query.NodeSetParent(cn, n)
	return cn
}

// Children represents the many-to-many relationship formed by the related_project_assn table.
func (n projectTable) Children() ProjectNode {
	cn := &projectAssociation{
		ManyManyNode: query.ManyManyNode{
			AssnTableQueryName:       "related_project_assn",
			ParentColumnQueryName:    "parent_id",
			ParentColumnReceiverType: query.ColTypeString,
			Identifier:               "Children",
			RefColumnQueryName:       "child_id",
			RefColumnReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) Children() ProjectNode {
	cn := n.projectTable.Children().(*projectAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) Children() ProjectNode {
	cn := n.projectTable.Children().(*projectAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) Children() ProjectNode {
	cn := n.projectTable.Children().(*projectAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

// Parents represents the many-to-many relationship formed by the related_project_assn table.
func (n projectTable) Parents() ProjectNode {
	cn := &projectAssociation{
		ManyManyNode: query.ManyManyNode{
			AssnTableQueryName:       "related_project_assn",
			ParentColumnQueryName:    "child_id",
			ParentColumnReceiverType: query.ColTypeString,
			Identifier:               "Parents",
			RefColumnQueryName:       "parent_id",
			RefColumnReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) Parents() ProjectNode {
	cn := n.projectTable.Parents().(*projectAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) Parents() ProjectNode {
	cn := n.projectTable.Parents().(*projectAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) Parents() ProjectNode {
	cn := n.projectTable.Parents().(*projectAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

// TeamMembers represents the many-to-many relationship formed by the team_member_project_assn table.
func (n projectTable) TeamMembers() PersonNode {
	cn := &personAssociation{
		ManyManyNode: query.ManyManyNode{
			AssnTableQueryName:       "team_member_project_assn",
			ParentColumnQueryName:    "project_id",
			ParentColumnReceiverType: query.ColTypeString,
			Identifier:               "TeamMembers",
			RefColumnQueryName:       "team_member_id",
			RefColumnReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) TeamMembers() PersonNode {
	cn := n.projectTable.TeamMembers().(*personAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) TeamMembers() PersonNode {
	cn := n.projectTable.TeamMembers().(*personAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) TeamMembers() PersonNode {
	cn := n.projectTable.TeamMembers().(*personAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

// Milestones represents the many-to-one relationship formed by the reverse reference from the
// project_id column in the milestone table.
func (n projectTable) Milestones() MilestoneNode {
	cn := &milestoneReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "project_id",
			Identifier:      "Milestones",
			ReceiverType:    query.ColTypeString,
			IsUnique:        false,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReference) Milestones() MilestoneNode {
	cn := n.projectTable.Milestones().(*milestoneReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectReverse) Milestones() MilestoneNode {
	cn := n.projectTable.Milestones().(*milestoneReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *projectAssociation) Milestones() MilestoneNode {
	cn := n.projectTable.Milestones().(*milestoneReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n projectTable) GobEncode() (data []byte, err error) {
	return
}

func (n *projectTable) GobDecode(data []byte) (err error) {
	return
}

func (n *projectReference) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *projectReference) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	return
}

func (n *projectReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *projectReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func (n *projectAssociation) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ManyManyNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *projectAssociation) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ManyManyNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(projectTable))
	gob.Register(new(projectReference))
	gob.Register(new(projectReverse))
	gob.Register(new(projectAssociation))
}
