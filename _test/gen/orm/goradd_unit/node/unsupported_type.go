// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// UnsupportedTypeI is the builder interface to the UnsupportedType nodes.
type UnsupportedTypeNodeI interface {
	query.NodeI
	PrimaryKeyNode() *query.ColumnNode

	TypeSerial() *query.ColumnNode
	TypeSet() *query.ColumnNode
	TypeEnum() *query.ColumnNode
	TypeDecimal() *query.ColumnNode
	TypeDouble() *query.ColumnNode
	TypeGeo() *query.ColumnNode
	TypeTinyBlob() *query.ColumnNode
	TypeMediumBlob() *query.ColumnNode
	TypeVarbinary() *query.ColumnNode
	TypeLongtext() *query.ColumnNode
	TypeBinary() *query.ColumnNode
	TypeSmall() *query.ColumnNode
	TypeMedium() *query.ColumnNode
	TypeBig() *query.ColumnNode
	TypePolygon() *query.ColumnNode
	TypeUnsigned() *query.ColumnNode
	TypeMultfk1() *query.ColumnNode
	TypeMultifk2() *query.ColumnNode
}

// UnsupportedTypeNode represents the unsupported_type table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the UnsupportedTypeNode, call [UnsupportedType] to start a reference chain when querying the unsupported_type table.
type UnsupportedTypeNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the unsupportedTypeNode functions here.
	query.ReferenceNodeI
}

// UnsupportedType returns a table node that starts a node chain that begins with the unsupported_type table.
func UnsupportedType() UnsupportedTypeNodeI {
	n := UnsupportedTypeNode{
		query.NewTableNode("goradd_unit", "unsupported_type", "UnsupportedType"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *UnsupportedTypeNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.TypeSerial())
	nodes = append(nodes, n.TypeSet())
	nodes = append(nodes, n.TypeEnum())
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

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *UnsupportedTypeNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *UnsupportedTypeNode) Copy_() query.NodeI {
	return &UnsupportedTypeNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *UnsupportedTypeNode) PrimaryKeyNode() *query.ColumnNode {
	return n.TypeSerial()
}

// TypeSerial represents the type_serial column in the database.
func (n *UnsupportedTypeNode) TypeSerial() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_serial",
		"TypeSerial",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeSet represents the type_set column in the database.
func (n *UnsupportedTypeNode) TypeSet() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_set",
		"TypeSet",
		query.ColTypeUnknown,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeEnum represents the type_enum column in the database.
func (n *UnsupportedTypeNode) TypeEnum() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_enum",
		"TypeEnum",
		query.ColTypeUnknown,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeDecimal represents the type_decimal column in the database.
func (n *UnsupportedTypeNode) TypeDecimal() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_decimal",
		"TypeDecimal",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeDouble represents the type_double column in the database.
func (n *UnsupportedTypeNode) TypeDouble() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_double",
		"TypeDouble",
		query.ColTypeFloat64,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeGeo represents the type_geo column in the database.
func (n *UnsupportedTypeNode) TypeGeo() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_geo",
		"TypeGeo",
		query.ColTypeUnknown,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeTinyBlob represents the type_tiny_blob column in the database.
func (n *UnsupportedTypeNode) TypeTinyBlob() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_tiny_blob",
		"TypeTinyBlob",
		query.ColTypeBytes,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeMediumBlob represents the type_medium_blob column in the database.
func (n *UnsupportedTypeNode) TypeMediumBlob() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_medium_blob",
		"TypeMediumBlob",
		query.ColTypeBytes,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeVarbinary represents the type_varbinary column in the database.
func (n *UnsupportedTypeNode) TypeVarbinary() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_varbinary",
		"TypeVarbinary",
		query.ColTypeUnknown,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeLongtext represents the type_longtext column in the database.
func (n *UnsupportedTypeNode) TypeLongtext() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_longtext",
		"TypeLongtext",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeBinary represents the type_binary column in the database.
func (n *UnsupportedTypeNode) TypeBinary() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_binary",
		"TypeBinary",
		query.ColTypeUnknown,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeSmall represents the type_small column in the database.
func (n *UnsupportedTypeNode) TypeSmall() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_small",
		"TypeSmall",
		query.ColTypeInteger,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeMedium represents the type_medium column in the database.
func (n *UnsupportedTypeNode) TypeMedium() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_medium",
		"TypeMedium",
		query.ColTypeInteger,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeBig represents the type_big column in the database.
func (n *UnsupportedTypeNode) TypeBig() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_big",
		"TypeBig",
		query.ColTypeInteger64,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypePolygon represents the type_polygon column in the database.
func (n *UnsupportedTypeNode) TypePolygon() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_polygon",
		"TypePolygon",
		query.ColTypeUnknown,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeUnsigned represents the type_unsigned column in the database.
func (n *UnsupportedTypeNode) TypeUnsigned() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_unsigned",
		"TypeUnsigned",
		query.ColTypeUnsigned,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeMultfk1 represents the type_multFk1 column in the database.
func (n *UnsupportedTypeNode) TypeMultfk1() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_multFk1",
		"TypeMultfk1",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// TypeMultifk2 represents the type_multiFk2 column in the database.
func (n *UnsupportedTypeNode) TypeMultifk2() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"unsupported_type",
		"type_multiFk2",
		"TypeMultifk2",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

type unsupportedTypeNodeEncoded struct {
	RefNode query.ReferenceNodeI
}

// GobEncode makes the node serializable.
// doc:hide
func (n *UnsupportedTypeNode) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	s := unsupportedTypeNodeEncoded{
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
func (n *UnsupportedTypeNode) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	var s unsupportedTypeNodeEncoded
	if err = dec.Decode(&s); err != nil {
		panic(err)
	}
	n.ReferenceNodeI = s.RefNode
	query.SetParentNode(n, query.ParentNode(n)) // Reinforce types
	return
}
func init() {
	gob.Register(&UnsupportedTypeNode{})
}
