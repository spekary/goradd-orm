// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// EmployeeInfoI is the builder interface to the EmployeeInfo nodes.
type EmployeeInfoNodeI interface {
	query.NodeI
	PrimaryKeyNode() *query.ColumnNode

	ID() *query.ColumnNode
	PersonID() *query.ColumnNode
	Person() PersonNodeI
	EmployeeNumber() *query.ColumnNode
}

// EmployeeInfoNode represents the employee_info table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the EmployeeInfoNode, call [EmployeeInfo] to start a reference chain when querying the employee_info table.
type EmployeeInfoNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the employeeInfoNode functions here.
	query.ReferenceNodeI
}

// EmployeeInfo returns a table node that starts a node chain that begins with the employee_info table.
func EmployeeInfo() EmployeeInfoNodeI {
	n := EmployeeInfoNode{
		query.NewTableNode("goradd", "employee_info", "EmployeeInfo"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *EmployeeInfoNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.PersonID())
	nodes = append(nodes, n.EmployeeNumber())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *EmployeeInfoNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *EmployeeInfoNode) Copy_() query.NodeI {
	return &EmployeeInfoNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *EmployeeInfoNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *EmployeeInfoNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"employee_info",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// PersonID represents the person_id column in the database.
func (n *EmployeeInfoNode) PersonID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"employee_info",
		"person_id",
		"PersonID",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Person represents the link to a Person object.
func (n *EmployeeInfoNode) Person() PersonNodeI {
	cn := &PersonNode{
		query.NewReferenceNode(
			"goradd",
			"employee_info",
			"person_id",
			"PersonID",
			"Person",
			"person",
			"id",
			false,
			query.ColTypeString,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

// EmployeeNumber represents the employee_number column in the database.
func (n *EmployeeInfoNode) EmployeeNumber() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"employee_info",
		"employee_number",
		"EmployeeNumber",
		query.ColTypeInteger,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

type employeeInfoNodeEncoded struct {
	RefNode query.ReferenceNodeI
}

// GobEncode makes the node serializable.
// doc:hide
func (n *EmployeeInfoNode) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	s := employeeInfoNodeEncoded{
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
func (n *EmployeeInfoNode) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	var s employeeInfoNodeEncoded
	if err = dec.Decode(&s); err != nil {
		panic(err)
	}
	n.ReferenceNodeI = s.RefNode
	query.SetParentNode(n, query.ParentNode(n)) // Reinforce types
	return
}
func init() {
	gob.Register(&EmployeeInfoNode{})
}
