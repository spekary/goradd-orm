// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// EmployeeInfoNode is the builder interface to the EmployeeInfo nodes.
type EmployeeInfoNode interface {
	query.Node
	PrimaryKey() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// PersonID represents the person_id column in the database.
	PersonID() *query.ColumnNode
	// Person represents the Person reference to a Person object.
	Person() PersonNode
	// EmployeeNumber represents the employee_number column in the database.
	EmployeeNumber() *query.ColumnNode
}

// employeeInfoTable represents the employee_info table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the employeeInfoTable, call [EmployeeInfo()] to start a reference chain when querying the employee_info table.
type employeeInfoTable struct {
}

type employeeInfoReverse struct {
	employeeInfoTable
	query.ReverseNode
}

// EmployeeInfo returns a table node that starts a node chain that begins with the employee_info table.
func EmployeeInfo() EmployeeInfoNode {
	return employeeInfoTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n employeeInfoTable) TableName_() string {
	return "employee_info"
}

// NodeType_ returns the query.NodeType of the node.
func (n employeeInfoTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n employeeInfoTable) DatabaseKey_() string {
	return "goradd"
}

// ColumnNodes_ is used internally by the framework to return the list of all the column nodes.
func (n employeeInfoTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.PersonID())
	nodes = append(nodes, n.EmployeeNumber())
	return nodes
}

func (n *employeeInfoReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.employeeInfoTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n employeeInfoTable) IsEnum_() bool {
	return false
}

func (n *employeeInfoReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKey returns a node that points to the primary key column.
func (n employeeInfoTable) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *employeeInfoReverse) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

func (n employeeInfoTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *employeeInfoReverse) ID() *query.ColumnNode {
	cn := n.employeeInfoTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n employeeInfoTable) PersonID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "person_id",
		Identifier:   "PersonID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *employeeInfoReverse) PersonID() *query.ColumnNode {
	cn := n.employeeInfoTable.PersonID()
	query.NodeSetParent(cn, n)
	return cn
}

// Person represents the link to a Person object.
func (n employeeInfoTable) Person() PersonNode {
	cn := &personReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "person_id",
			Identifier:      "PersonID",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *employeeInfoReverse) Person() PersonNode {
	cn := n.employeeInfoTable.Person().(*personReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n employeeInfoTable) EmployeeNumber() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "employee_number",
		Identifier:   "EmployeeNumber",
		ReceiverType: query.ColTypeInteger,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *employeeInfoReverse) EmployeeNumber() *query.ColumnNode {
	cn := n.employeeInfoTable.EmployeeNumber()
	query.NodeSetParent(cn, n)
	return cn
}

func (n employeeInfoTable) GobEncode() (data []byte, err error) {
	return
}

func (n *employeeInfoTable) GobDecode(data []byte) (err error) {
	return
}

func (n *employeeInfoReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *employeeInfoReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(employeeInfoTable))
	gob.Register(new(employeeInfoReverse))
}
