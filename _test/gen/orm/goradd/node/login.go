// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"github.com/goradd/orm/pkg/query"
)

// LoginNode represents the login table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the LoginNode, call [Login] to start a reference chain when querying the login table.
type LoginNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the loginNode functions here.
	query.ReferenceNodeI
}

// Login returns a table node that starts a node chain that begins with the login table.
func Login() *LoginNode {
	n := LoginNode{
		query.NewTableNode("goradd", "login", "Login"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *LoginNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.PersonID())
	nodes = append(nodes, n.Username())
	nodes = append(nodes, n.Password())
	nodes = append(nodes, n.IsEnabled())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *LoginNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *LoginNode) Copy_() query.NodeI {
	return &LoginNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *LoginNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *LoginNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"login",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// PersonID represents the person_id column in the database.
func (n *LoginNode) PersonID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"login",
		"person_id",
		"PersonID",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Person represents the link to a Person object.
func (n *LoginNode) Person() *PersonNode {
	cn := &PersonNode{
		query.NewReferenceNode(
			"goradd",
			"login",
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

// Username represents the username column in the database.
func (n *LoginNode) Username() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"login",
		"username",
		"Username",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Password represents the password column in the database.
func (n *LoginNode) Password() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"login",
		"password",
		"Password",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// IsEnabled represents the is_enabled column in the database.
func (n *LoginNode) IsEnabled() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"login",
		"is_enabled",
		"IsEnabled",
		query.ColTypeBool,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}
