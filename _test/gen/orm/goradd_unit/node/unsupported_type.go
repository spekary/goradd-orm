// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
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
	// TypeDecimal represents the type_decimal column in the database.
	TypeDecimal() *query.ColumnNode
	// TypeDouble represents the type_double column in the database.
	TypeDouble() *query.ColumnNode
	// TypeGeo represents the type_geo column in the database.
	TypeGeo() *query.ColumnNode
	// TypeTinyBlob represents the type_tiny_blob column in the database.
	TypeTinyBlob() *query.ColumnNode
	// TypeMediumBlob represents the type_medium_blob column in the database.
	TypeMediumBlob() *query.ColumnNode
	// TypeVarbinary represents the type_varbinary column in the database.
	TypeVarbinary() *query.ColumnNode
	// TypeLongtext represents the type_longtext column in the database.
	TypeLongtext() *query.ColumnNode
	// TypeBinary represents the type_binary column in the database.
	TypeBinary() *query.ColumnNode
	// TypeSmall represents the type_small column in the database.
	TypeSmall() *query.ColumnNode
	// TypeMedium represents the type_medium column in the database.
	TypeMedium() *query.ColumnNode
	// TypeBig represents the type_big column in the database.
	TypeBig() *query.ColumnNode
	// TypePolygon represents the type_polygon column in the database.
	TypePolygon() *query.ColumnNode
	// TypeUnsigned represents the type_unsigned column in the database.
	TypeUnsigned() *query.ColumnNode
	// TypeMultfk1 represents the type_multFk1 column in the database.
	TypeMultfk1() *query.ColumnNode
	// TypeMultifk2 represents the type_multiFk2 column in the database.
	TypeMultifk2() *query.ColumnNode
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
	nodes = append(nodes, n.TypeDecimal())
	nodes = append(nodes, n.TypeDouble())
	nodes = append(nodes, n.TypeGeo())
	nodes = append(nodes, n.TypeTinyBlob())
	nodes = append(nodes, n.TypeMediumBlob())
	nodes = append(nodes, n.TypeVarbinary())
	nodes = append(nodes, n.TypeLongtext())
	nodes = append(nodes, n.TypeBinary())
	nodes = append(nodes, n.TypeSmall())
	nodes = append(nodes, n.TypeMedium())
	nodes = append(nodes, n.TypeBig())
	nodes = append(nodes, n.TypePolygon())
	nodes = append(nodes, n.TypeUnsigned())
	nodes = append(nodes, n.TypeMultfk1())
	nodes = append(nodes, n.TypeMultifk2())
	return nodes
}

// PrimaryKey returns a node that points to the primary key column.
func (n unsupportedTypeTable) PrimaryKey() *query.ColumnNode {
	return n.TypeSerial()
}

func (n unsupportedTypeTable) TypeSerial() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_serial",
		Identifier:   "TypeSerial",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeSet() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_set",
		Identifier:   "TypeSet",
		ReceiverType: query.ColTypeUnknown,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeEnumerated() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_enumerated",
		Identifier:   "TypeEnumerated",
		ReceiverType: query.ColTypeUnknown,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeDecimal() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_decimal",
		Identifier:   "TypeDecimal",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeDouble() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_double",
		Identifier:   "TypeDouble",
		ReceiverType: query.ColTypeFloat64,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeGeo() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_geo",
		Identifier:   "TypeGeo",
		ReceiverType: query.ColTypeUnknown,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeTinyBlob() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_tiny_blob",
		Identifier:   "TypeTinyBlob",
		ReceiverType: query.ColTypeBytes,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeMediumBlob() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_medium_blob",
		Identifier:   "TypeMediumBlob",
		ReceiverType: query.ColTypeBytes,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeVarbinary() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_varbinary",
		Identifier:   "TypeVarbinary",
		ReceiverType: query.ColTypeUnknown,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeLongtext() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_longtext",
		Identifier:   "TypeLongtext",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeBinary() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_binary",
		Identifier:   "TypeBinary",
		ReceiverType: query.ColTypeUnknown,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeSmall() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_small",
		Identifier:   "TypeSmall",
		ReceiverType: query.ColTypeInteger,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeMedium() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_medium",
		Identifier:   "TypeMedium",
		ReceiverType: query.ColTypeInteger,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeBig() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_big",
		Identifier:   "TypeBig",
		ReceiverType: query.ColTypeInteger64,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypePolygon() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_polygon",
		Identifier:   "TypePolygon",
		ReceiverType: query.ColTypeUnknown,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeUnsigned() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_unsigned",
		Identifier:   "TypeUnsigned",
		ReceiverType: query.ColTypeUnsigned,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeMultfk1() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_multFk1",
		Identifier:   "TypeMultfk1",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n unsupportedTypeTable) TypeMultifk2() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "type_multiFk2",
		Identifier:   "TypeMultifk2",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
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
