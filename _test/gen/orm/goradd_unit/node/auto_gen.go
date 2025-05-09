// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
	"github.com/goradd/orm/pkg/schema"
)

// AutoGenNode is the builder interface to the AutoGen nodes.
type AutoGenNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// GroLock represents the gro_lock column in the database.
	GroLock() *query.ColumnNode
	// GroTimestamp represents the gro_timestamp column in the database.
	GroTimestamp() *query.ColumnNode
	// Created represents the created column in the database.
	Created() *query.ColumnNode
	// Modified represents the modified column in the database.
	Modified() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
}

// autoGenTable represents the auto_gen table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the autoGenTable, call [AutoGen()] to start a reference chain when querying the auto_gen table.
type autoGenTable struct {
}

// AutoGen returns a table node that starts a node chain that begins with the auto_gen table.
func AutoGen() AutoGenNode {
	return autoGenTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n autoGenTable) TableName_() string {
	return "auto_gen"
}

// NodeType_ returns the query.NodeType of the node.
func (n autoGenTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n autoGenTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n autoGenTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.GroLock())
	nodes = append(nodes, n.GroTimestamp())
	nodes = append(nodes, n.Created())
	nodes = append(nodes, n.Modified())
	nodes = append(nodes, n.Name())
	return nodes
}

// PrimaryKey returns a node that points to the primary key column.
func (n autoGenTable) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

func (n autoGenTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "id",
		Identifier:    "ID",
		ReceiverType:  query.ColTypeString,
		SchemaType:    schema.ColTypeAutoPrimaryKey,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n autoGenTable) GroLock() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "gro_lock",
		Identifier:    "GroLock",
		ReceiverType:  query.ColTypeInteger64,
		SchemaType:    schema.ColTypeInt,
		SchemaSubType: schema.ColSubTypeLock,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n autoGenTable) GroTimestamp() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "gro_timestamp",
		Identifier:    "GroTimestamp",
		ReceiverType:  query.ColTypeInteger64,
		SchemaType:    schema.ColTypeInt,
		SchemaSubType: schema.ColSubTypeTimestamp,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n autoGenTable) Created() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "created",
		Identifier:    "Created",
		ReceiverType:  query.ColTypeTime,
		SchemaType:    schema.ColTypeTime,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n autoGenTable) Modified() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "modified",
		Identifier:    "Modified",
		ReceiverType:  query.ColTypeTime,
		SchemaType:    schema.ColTypeTime,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n autoGenTable) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "name",
		Identifier:    "Name",
		ReceiverType:  query.ColTypeString,
		SchemaType:    schema.ColTypeString,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n autoGenTable) GobEncode() (data []byte, err error) {
	return
}

func (n *autoGenTable) GobDecode(data []byte) (err error) {
	return
}

func init() {
	gob.Register(new(autoGenTable))
}
