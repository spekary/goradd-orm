// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
	"github.com/goradd/orm/pkg/schema"
)

// GiftNode is the builder interface to the Gift nodes.
type GiftNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// Number represents the number column in the database.
	Number() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
}

// giftTable represents the gift table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the giftTable, call [Gift()] to start a reference chain when querying the gift table.
type giftTable struct {
}

// Gift returns a table node that starts a node chain that begins with the gift table.
func Gift() GiftNode {
	return giftTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n giftTable) TableName_() string {
	return "gift"
}

// NodeType_ returns the query.NodeType of the node.
func (n giftTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n giftTable) DatabaseKey_() string {
	return "goradd"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n giftTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.Number())
	nodes = append(nodes, n.Name())
	return nodes
}

// PrimaryKey returns a node that points to the primary key column.
func (n giftTable) PrimaryKey() *query.ColumnNode {
	return n.Number()
}

func (n giftTable) Number() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "number",
		Identifier:    "Number",
		ReceiverType:  query.ColTypeInteger,
		SchemaType:    schema.ColTypeInt,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n giftTable) Name() *query.ColumnNode {
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

func (n giftTable) GobEncode() (data []byte, err error) {
	return
}

func (n *giftTable) GobDecode(data []byte) (err error) {
	return
}

func init() {
	gob.Register(new(giftTable))
}
