// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// PersonTypeNodeI is the builder interface to the PersonType nodes.
type PersonTypeNodeI interface {
	query.Node
	PrimaryKeyNode() *query.ColumnNode
	ID() *query.ColumnNode
	Name() *query.ColumnNode
}

type personTypeEnum struct {
}

type personTypeReference struct {
	personTypeEnum
	query.ReferenceNode
}

type personTypeAssociation struct {
	personTypeEnum
	query.ManyManyNode
}

// PrimaryKeyNode returns a node representing the primary key column.
func (n personTypeEnum) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n personTypeReference) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n personTypeAssociation) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func init() {
	gob.Register(new(personTypeEnum))
	gob.Register(new(personTypeAssociation))
	gob.Register(new(personTypeReference))
}

// ColumnNodes_ is used internally by the framework to return the list of column nodes.
func (n personTypeEnum) ColumnNodes_() []query.Node {
	return []query.Node{
		n.ID(),
		n.Name(),
	}
}

func (n personTypeAssociation) ColumnNodes_() []query.Node {
	return []query.Node{
		n.ID(),
		n.Name(),
	}
}

// Columns_ is used internally by the framework to return the list of column names.
func (n personTypeEnum) Columns_() []string {
	return []string{
		"id",
		"name",
	}
}

func (n personTypeEnum) NodeType_() query.NodeType {
	return query.EnumNodeType
}

func (n *personTypeAssociation) NodeType_() query.NodeType {
	return query.ManyEnumNodeType
}

// TableName_ returns the query name of the table the node is associated with.
func (n personTypeEnum) TableName_() string {
	return "person_type_enum"
}

// TableName_ returns the query name of the table the node is associated with.
func (n personTypeAssociation) TableName_() string {
	return n.Parent().TableName_()
}

// TableName_ returns the query name of the table the node is associated with.
func (n personTypeReference) TableName_() string {
	return n.Parent().TableName_()
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n personTypeEnum) DatabaseKey_() string {
	return "goradd"
}

func (n personTypeEnum) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeInteger,
		IsPrimaryKey: true,
	}
	cn.SetParent(n)
	return cn
}

func (n *personTypeReference) ID() *query.ColumnNode {
	cn := n.personTypeEnum.ID()
	cn.SetParent(n)
	return cn
}

func (n *personTypeAssociation) ID() *query.ColumnNode {
	cn := n.personTypeEnum.ID()
	cn.SetParent(n)
	return cn
}

func (n personTypeEnum) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *personTypeReference) Name() *query.ColumnNode {
	cn := n.personTypeEnum.Name()
	cn.SetParent(n)
	return cn
}

func (n *personTypeAssociation) Name() *query.ColumnNode {
	cn := n.personTypeEnum.Name()
	cn.SetParent(n)
	return cn
}

func (n personTypeEnum) GobEncode() (data []byte, err error) {
	return
}

func (n *personTypeEnum) GobDecode(data []byte) (err error) {
	return
}

func (n *personTypeReference) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *personTypeReference) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	return
}

func (n *personTypeAssociation) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ManyManyNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *personTypeAssociation) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ManyManyNode); err != nil {
		panic(err)
	}
	return
}
