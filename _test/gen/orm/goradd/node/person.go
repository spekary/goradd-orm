// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// PersonNode is the builder interface to the Person nodes.
type PersonNode interface {
	query.Node
	PrimaryKeyNode() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// FirstName represents the first_name column in the database.
	FirstName() *query.ColumnNode
	// LastName represents the last_name column in the database.
	LastName() *query.ColumnNode
	// Types represents the type_enum column in the database.
	Types() *query.ColumnNode
	// Projects represents the Projects reference to Project objects.
	Projects() ProjectExpander
	// Addresses represents the Addresses reference to Address objects.
	Addresses() AddressExpander
	// EmployeeInfo represents the EmployeeInfo reference to a EmployeeInfo object.
	EmployeeInfo() EmployeeInfoNode
	// Login represents the Login reference to a Login object.
	Login() LoginNode
	// ManagerProjects represents the ManagerProjects reference to Project objects.
	ManagerProjects() ProjectExpander
}

// PersonExpander is the builder interface for People that are expandable.
type PersonExpander interface {
	PersonNode
	query.Expander
}

// personTable represents the person table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the personTable, call [Person()] to start a reference chain when querying the person table.
type personTable struct {
}

type personReference struct {
	personTable
	query.ReferenceNode
}

type personAssociation struct {
	personTable
	query.ManyManyNode
}

// Person returns a table node that starts a node chain that begins with the person table.
func Person() PersonNode {
	return personTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n personTable) TableName_() string {
	return "person"
}

// NodeType_ returns the query.NodeType of the node.
func (n personTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n personTable) DatabaseKey_() string {
	return "goradd"
}

// ColumnNodes_ is used internally by the framework to return the list of all the column nodes.
func (n personTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.FirstName())
	nodes = append(nodes, n.LastName())
	nodes = append(nodes, n.Types())
	return nodes
}

func (n *personReference) ColumnNodes_() (nodes []query.Node) {
	nodes = n.personTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

func (n *personAssociation) ColumnNodes_() (nodes []query.Node) {
	nodes = n.personTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n personTable) IsEnum_() bool {
	return false
}

func (n *personReference) NodeType_() query.NodeType {
	return query.ReferenceNodeType
}

func (n *personAssociation) NodeType_() query.NodeType {
	return query.ManyManyNodeType
}

// PrimaryKeyNode returns a node that points to the primary key column.
func (n personTable) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n *personReference) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n *personAssociation) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n personTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personReference) ID() *query.ColumnNode {
	cn := n.personTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personAssociation) ID() *query.ColumnNode {
	cn := n.personTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n personTable) FirstName() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "first_name",
		Identifier:   "FirstName",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personReference) FirstName() *query.ColumnNode {
	cn := n.personTable.FirstName()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personAssociation) FirstName() *query.ColumnNode {
	cn := n.personTable.FirstName()
	query.NodeSetParent(cn, n)
	return cn
}

func (n personTable) LastName() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "last_name",
		Identifier:   "LastName",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personReference) LastName() *query.ColumnNode {
	cn := n.personTable.LastName()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personAssociation) LastName() *query.ColumnNode {
	cn := n.personTable.LastName()
	query.NodeSetParent(cn, n)
	return cn
}

func (n personTable) Types() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_enum",
		Identifier:   "Types",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personReference) Types() *query.ColumnNode {
	cn := n.personTable.Types()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personAssociation) Types() *query.ColumnNode {
	cn := n.personTable.Types()
	query.NodeSetParent(cn, n)
	return cn
}

// Projects represents the many-to-many relationship formed by the team_member_project_assn table.
func (n personTable) Projects() ProjectExpander {
	cn := &projectAssociation{
		ManyManyNode: query.ManyManyNode{
			AssnTableQueryName:       "team_member_project_assn",
			ParentColumnQueryName:    "team_member_id",
			ParentColumnReceiverType: query.ColTypeString,
			Identifier:               "Projects",
			RefColumnQueryName:       "project_id",
			RefColumnReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personReference) Projects() ProjectExpander {
	cn := n.personTable.Projects().(*projectAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personAssociation) Projects() ProjectExpander {
	cn := n.personTable.Projects().(*projectAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

// Addresses represents the many-to-one relationship formed by the reverse reference from the
// person_id column in the address table.
func (n personTable) Addresses() AddressExpander {
	cn := &addressReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "person_id",
			Identifier:      "Addresses",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personReference) Addresses() AddressExpander {
	cn := n.personTable.Addresses().(*addressReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personAssociation) Addresses() AddressExpander {
	cn := n.personTable.Addresses().(*addressReverse)
	query.NodeSetParent(cn, n)
	return cn
}

// EmployeeInfo represents the one-to-one relationship formed by the reverse reference from the
// person_id column in the employee_info table.
func (n personTable) EmployeeInfo() EmployeeInfoNode {
	cn := &employeeInfoReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "person_id",
			Identifier:      "EmployeeInfo",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personReference) EmployeeInfo() EmployeeInfoNode {
	cn := n.personTable.EmployeeInfo().(*employeeInfoReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personAssociation) EmployeeInfo() EmployeeInfoNode {
	cn := n.personTable.EmployeeInfo().(*employeeInfoReverse)
	query.NodeSetParent(cn, n)
	return cn
}

// Login represents the one-to-one relationship formed by the reverse reference from the
// person_id column in the login table.
func (n personTable) Login() LoginNode {
	cn := &loginReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "person_id",
			Identifier:      "Login",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personReference) Login() LoginNode {
	cn := n.personTable.Login().(*loginReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personAssociation) Login() LoginNode {
	cn := n.personTable.Login().(*loginReverse)
	query.NodeSetParent(cn, n)
	return cn
}

// ManagerProjects represents the many-to-one relationship formed by the reverse reference from the
// manager_id column in the project table.
func (n personTable) ManagerProjects() ProjectExpander {
	cn := &projectReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "manager_id",
			Identifier:      "ManagerProjects",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personReference) ManagerProjects() ProjectExpander {
	cn := n.personTable.ManagerProjects().(*projectReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *personAssociation) ManagerProjects() ProjectExpander {
	cn := n.personTable.ManagerProjects().(*projectReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n personTable) GobEncode() (data []byte, err error) {
	return
}

func (n *personTable) GobDecode(data []byte) (err error) {
	return
}

func (n *personReference) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *personReference) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	return
}

func (n *personAssociation) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ManyManyNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *personAssociation) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ManyManyNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(personTable))
	gob.Register(new(personReference))
	gob.Register(new(personAssociation))
}
