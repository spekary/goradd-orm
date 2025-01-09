// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// PersonWithLockNodeI is the builder interface to the PersonWithLock nodes.
type PersonWithLockNodeI interface {
	query.NodeI
	PrimaryKeyNode() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// FirstName represents the first_name column in the database.
	FirstName() *query.ColumnNode
	// LastName represents the last_name column in the database.
	LastName() *query.ColumnNode
	// SysTimestamp represents the sys_timestamp column in the database.
	SysTimestamp() *query.ColumnNode
}

// PersonWithLockExpander is the builder interface for PersonWithLocks that are expandable.
type PersonWithLockExpander interface {
	PersonWithLockNodeI
	// Expand causes the node to produce separate rows with individual items, rather than a single row with an array of items.
	Expand()
}

// personWithLockTable represents the person_with_lock table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the personWithLockTable, call [PersonWithLock()] to start a reference chain when querying the person_with_lock table.
type personWithLockTable struct {
	_self query.NodeI
}

type personWithLockReverse struct {
	personWithLockTable
	query.ReverseNode
}

// PersonWithLock returns a table node that starts a node chain that begins with the person_with_lock table.
func PersonWithLock() PersonWithLockNodeI {
	var n personWithLockTable
	n._self = n
	return n
}

// TableName_ returns the query name of the table the node is associated with.
func (n personWithLockTable) TableName_() string {
	return "person_with_lock"
}

// NodeType_ returns the query.NodeType of the node.
func (n personWithLockTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n personWithLockTable) DatabaseKey_() string {
	return "goradd"
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
func (n personWithLockTable) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.FirstName())
	nodes = append(nodes, n.LastName())
	nodes = append(nodes, n.SysTimestamp())
	return nodes
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n personWithLockTable) IsEnum_() bool {
	return false
}

func (n *personWithLockReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n personWithLockTable) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n personWithLockTable) ID() *query.ColumnNode {
	cn := query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	cn.SetParent(n._self)
	return &cn
}

// FirstName represents the first_name column in the database.
func (n personWithLockTable) FirstName() *query.ColumnNode {
	cn := query.ColumnNode{
		QueryName:    "first_name",
		Identifier:   "FirstName",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n._self)
	return &cn
}

// LastName represents the last_name column in the database.
func (n personWithLockTable) LastName() *query.ColumnNode {
	cn := query.ColumnNode{
		QueryName:    "last_name",
		Identifier:   "LastName",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n._self)
	return &cn
}

// SysTimestamp represents the sys_timestamp column in the database.
func (n personWithLockTable) SysTimestamp() *query.ColumnNode {
	cn := query.ColumnNode{
		QueryName:    "sys_timestamp",
		Identifier:   "SysTimestamp",
		ReceiverType: query.ColTypeTime,
		IsPrimaryKey: false,
	}
	cn.SetParent(n._self)
	return &cn
}

func (n *personWithLockTable) GobEncode() (data []byte, err error) {
	return
}

func (n *personWithLockTable) GobDecode(data []byte) (err error) {
	n._self = n
	return
}

func (n *personWithLockReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *personWithLockReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	n._self = n
	return
}

func init() {
	gob.Register(new(personWithLockTable))
	gob.Register(new(personWithLockReverse))
}
