// Code generated by GoRADD. DO NOT EDIT.

package node

import (
	"github.com/goradd/orm/pkg/query"
)

// PersonWithLockNode represents the person_with_lock table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the PersonWithLockNode, call [PersonWithLock] to start a reference chain when querying the person_with_lock table.
type PersonWithLockNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the personWithLockNode functions here.
	query.ReferenceNodeI
}

// PersonWithLock returns a table node that starts a node chain that begins with the person_with_lock table.
func PersonWithLock() *PersonWithLockNode {
	n := PersonWithLockNode{
		query.NewTableNode("goradd", "person_with_lock", "PersonWithLock"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *PersonWithLockNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.FirstName())
	nodes = append(nodes, n.LastName())
	nodes = append(nodes, n.SysTimestamp())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *PersonWithLockNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *PersonWithLockNode) Copy_() query.NodeI {
	return &PersonWithLockNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *PersonWithLockNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *PersonWithLockNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"person_with_lock",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// FirstName represents the first_name column in the database.
func (n *PersonWithLockNode) FirstName() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"person_with_lock",
		"first_name",
		"FirstName",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// LastName represents the last_name column in the database.
func (n *PersonWithLockNode) LastName() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"person_with_lock",
		"last_name",
		"LastName",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// SysTimestamp represents the sys_timestamp column in the database.
func (n *PersonWithLockNode) SysTimestamp() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"person_with_lock",
		"sys_timestamp",
		"SysTimestamp",
		query.ColTypeTime,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}