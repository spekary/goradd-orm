// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// LeafLockNode is the builder interface to the LeafLock nodes.
type LeafLockNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
	// GroLock represents the gro_lock column in the database.
	GroLock() *query.ColumnNode
}

// leafLockTable represents the leaf_lock table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the leafLockTable, call [LeafLock()] to start a reference chain when querying the leaf_lock table.
type leafLockTable struct {
}

// LeafLock returns a table node that starts a node chain that begins with the leaf_lock table.
func LeafLock() LeafLockNode {
	return leafLockTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n leafLockTable) TableName_() string {
	return "leaf_lock"
}

// NodeType_ returns the query.NodeType of the node.
func (n leafLockTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n leafLockTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n leafLockTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.GroLock())
	return nodes
}

// PrimaryKey returns a node that points to the primary key column.
func (n leafLockTable) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

func (n leafLockTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n leafLockTable) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n leafLockTable) GroLock() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "gro_lock",
		Identifier:   "GroLock",
		ReceiverType: query.ColTypeInteger64,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n leafLockTable) GobEncode() (data []byte, err error) {
	return
}

func (n *leafLockTable) GobDecode(data []byte) (err error) {
	return
}

func init() {
	gob.Register(new(leafLockTable))
}
