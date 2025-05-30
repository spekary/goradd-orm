// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
	"github.com/goradd/orm/pkg/schema"
)

// UnsupportedTypeNode is the builder interface to the UnsupportedType nodes.
type UnsupportedTypeNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// TypeSerial represents the type_serial column in the database.
	TypeSerial() *query.ColumnNode
	// TypeSet represents the type_set column in the database.
	TypeSet() *query.ColumnNode
	// TypeEnumerated represents the type_enumerated column in the database.
	TypeEnumerated() *query.ColumnNode
	// TypeGeo represents the type_geo column in the database.
	TypeGeo() *query.ColumnNode
	// TypeTinyblob represents the type_tinyblob column in the database.
	TypeTinyblob() *query.ColumnNode
	// TypeBinary represents the type_binary column in the database.
	TypeBinary() *query.ColumnNode
	// TypeSmall represents the type_small column in the database.
	TypeSmall() *query.ColumnNode
	// TypeMedium represents the type_medium column in the database.
	TypeMedium() *query.ColumnNode
	// TypePolygon represents the type_polygon column in the database.
	TypePolygon() *query.ColumnNode
	// TypeMultFk1 represents the type_mult_fk1 column in the database.
	TypeMultFk1() *query.ColumnNode
	// TypeMultiFk2 represents the type_multi_fk2 column in the database.
	TypeMultiFk2() *query.ColumnNode
}

// unsupportedTypeTable represents the unsupported_type table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the unsupportedTypeTable, call [UnsupportedType()] to start a reference chain when querying the unsupported_type table.
type unsupportedTypeTable struct {
}

// UnsupportedType returns a table node that starts a node chain that begins with the unsupported_type table.
func UnsupportedType() UnsupportedTypeNode {
	return unsupportedTypeTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n unsupportedTypeTable) TableName_() string {
	return "unsupported_type"
}

// NodeType_ returns the query.NodeType of the node.
func (n unsupportedTypeTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n unsupportedTypeTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n unsupportedTypeTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.TypeSerial())
	nodes = append(nodes, n.TypeSet())
	nodes = append(nodes, n.TypeEnumerated())
	nodes = append(nodes, n.TypeGeo())
	nodes = append(nodes, n.TypeTinyblob())
	nodes = append(nodes, n.TypeBinary())
	nodes = append(nodes, n.TypeSmall())
	nodes = append(nodes, n.TypeMedium())
	nodes = append(nodes, n.TypePolygon())
	nodes = append(nodes, n.TypeMultFk1())
	nodes = append(nodes, n.TypeMultiFk2())
	return nodes
}

// PrimaryKey returns a node that points to the primary key column.
func (n unsupportedTypeTable) PrimaryKey() *query.ColumnNode {
	return n.TypeSerial()
}

func (n unsupportedTypeTable) TypeSerial() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_serial",
		Identifier:    "TypeSerial",
		ReceiverType:  query.ColTypeUnsigned64,
		SchemaType:    schema.ColTypeUint,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeSet() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_set",
		Identifier:    "TypeSet",
		ReceiverType:  query.ColTypeUnknown,
		SchemaType:    schema.ColTypeUnknown,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeEnumerated() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_enumerated",
		Identifier:    "TypeEnumerated",
		ReceiverType:  query.ColTypeUnknown,
		SchemaType:    schema.ColTypeUnknown,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeGeo() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_geo",
		Identifier:    "TypeGeo",
		ReceiverType:  query.ColTypeUnknown,
		SchemaType:    schema.ColTypeUnknown,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeTinyblob() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_tinyblob",
		Identifier:    "TypeTinyblob",
		ReceiverType:  query.ColTypeBytes,
		SchemaType:    schema.ColTypeBytes,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeBinary() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_binary",
		Identifier:    "TypeBinary",
		ReceiverType:  query.ColTypeUnknown,
		SchemaType:    schema.ColTypeUnknown,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeSmall() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_small",
		Identifier:    "TypeSmall",
		ReceiverType:  query.ColTypeInteger,
		SchemaType:    schema.ColTypeInt,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeMedium() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_medium",
		Identifier:    "TypeMedium",
		ReceiverType:  query.ColTypeInteger,
		SchemaType:    schema.ColTypeInt,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypePolygon() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_polygon",
		Identifier:    "TypePolygon",
		ReceiverType:  query.ColTypeUnknown,
		SchemaType:    schema.ColTypeUnknown,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeMultFk1() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_mult_fk1",
		Identifier:    "TypeMultFk1",
		ReceiverType:  query.ColTypeString,
		SchemaType:    schema.ColTypeString,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeMultiFk2() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "type_multi_fk2",
		Identifier:    "TypeMultiFk2",
		ReceiverType:  query.ColTypeString,
		SchemaType:    schema.ColTypeString,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) GobEncode() (data []byte, err error) {
	return
}

func (n *unsupportedTypeTable) GobDecode(data []byte) (err error) {
	return
}

func init() {
	gob.Register(new(unsupportedTypeTable))
}
