// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// LoginNode is the builder interface to the Login nodes.
type LoginNode interface {
	query.Node
	PrimaryKeyNode() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// PersonID represents the person_id column in the database.
	PersonID() *query.ColumnNode
	// Person represents the Person reference to a Person object.
	Person() PersonNode
	// Username represents the username column in the database.
	Username() *query.ColumnNode
	// Password represents the password column in the database.
	Password() *query.ColumnNode
	// IsEnabled represents the is_enabled column in the database.
	IsEnabled() *query.ColumnNode
}

// LoginExpander is the builder interface for Logins that are expandable.
type LoginExpander interface {
	LoginNode
	query.Expander
}

// loginTable represents the login table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the loginTable, call [Login()] to start a reference chain when querying the login table.
type loginTable struct {
}

type loginReverse struct {
	loginTable
	query.ReverseNode
}

// Login returns a table node that starts a node chain that begins with the login table.
func Login() LoginNode {
	return loginTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n loginTable) TableName_() string {
	return "login"
}

// NodeType_ returns the query.NodeType of the node.
func (n loginTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n loginTable) DatabaseKey_() string {
	return "goradd"
}

// ColumnNodes_ is used internally by the framework to return the list of all the column nodes.
func (n loginTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.PersonID())
	nodes = append(nodes, n.Username())
	nodes = append(nodes, n.Password())
	nodes = append(nodes, n.IsEnabled())
	return nodes
}

func (n *loginReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.loginTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n loginTable) IsEnum_() bool {
	return false
}

func (n *loginReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKeyNode returns a node that points to the primary key column.
func (n loginTable) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n *loginReverse) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n loginTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *loginReverse) ID() *query.ColumnNode {
	cn := n.loginTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n loginTable) PersonID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "person_id",
		Identifier:   "PersonID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *loginReverse) PersonID() *query.ColumnNode {
	cn := n.loginTable.PersonID()
	query.NodeSetParent(cn, n)
	return cn
}

// Person represents the link to a Person object.
func (n loginTable) Person() PersonNode {
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

func (n *loginReverse) Person() PersonNode {
	cn := n.loginTable.Person().(*personReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n loginTable) Username() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "username",
		Identifier:   "Username",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *loginReverse) Username() *query.ColumnNode {
	cn := n.loginTable.Username()
	query.NodeSetParent(cn, n)
	return cn
}

func (n loginTable) Password() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "password",
		Identifier:   "Password",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *loginReverse) Password() *query.ColumnNode {
	cn := n.loginTable.Password()
	query.NodeSetParent(cn, n)
	return cn
}

func (n loginTable) IsEnabled() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "is_enabled",
		Identifier:   "IsEnabled",
		ReceiverType: query.ColTypeBool,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *loginReverse) IsEnabled() *query.ColumnNode {
	cn := n.loginTable.IsEnabled()
	query.NodeSetParent(cn, n)
	return cn
}

func (n loginTable) GobEncode() (data []byte, err error) {
	return
}

func (n *loginTable) GobDecode(data []byte) (err error) {
	return
}

func (n *loginReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *loginReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(loginTable))
	gob.Register(new(loginReverse))
}
