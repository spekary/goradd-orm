// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// ProjectStatusNodeI is the builder interface to the ProjectStatus nodes.
type ProjectStatusNodeI interface {
	query.NodeI
	PrimaryKeyNode() *query.ColumnNode
	ID() *query.ColumnNode
	Name() *query.ColumnNode
	Description() *query.ColumnNode
	Guidelines() *query.ColumnNode
	IsActive() *query.ColumnNode
}

type projectStatusEnum struct {
}

type projectStatusReference struct {
	projectStatusEnum
	query.ReferenceNode
}

type projectStatusAssociation struct {
	projectStatusEnum
	query.ManyManyNode
}

// PrimaryKeyNode returns a node representing the primary key column.
func (n projectStatusEnum) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n projectStatusReference) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n projectStatusAssociation) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func init() {
	gob.Register(new(projectStatusEnum))
	gob.Register(new(projectStatusAssociation))
	gob.Register(new(projectStatusReference))
}

// SelectNodes_ is used internally by the framework to return the list of column nodes.
func (n projectStatusEnum) SelectNodes_() []*query.ColumnNode {
	return []*query.ColumnNode{
		n.ID(),
		n.Name(),
		n.Description(),
		n.Guidelines(),
		n.IsActive(),
	}
}

func (n projectStatusEnum) NodeType_() query.NodeType {
	return query.EnumNodeType
}

func (n *projectStatusAssociation) NodeType_() query.NodeType {
	return query.ManyEnumNodeType
}

// TableName_ returns the query name of the table the node is associated with.
func (n projectStatusEnum) TableName_() string {
	return "project_status_enum"
}

// TableName_ returns the query name of the table the node is associated with.
func (n projectStatusAssociation) TableName_() string {
	return n.Parent().TableName_()
}

// TableName_ returns the query name of the table the node is associated with.
func (n projectStatusReference) TableName_() string {
	return n.Parent().TableName_()
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n projectStatusEnum) DatabaseKey_() string {
	return "goradd"
}

func (n projectStatusEnum) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeInteger,
		IsPrimaryKey: true,
	}
	cn.SetParent(n)
	return cn
}

func (n *projectStatusReference) ID() *query.ColumnNode {
	cn := n.projectStatusEnum.ID()
	cn.SetParent(n)
	return cn
}

func (n *projectStatusAssociation) ID() *query.ColumnNode {
	cn := n.projectStatusEnum.ID()
	cn.SetParent(n)
	return cn
}

func (n projectStatusEnum) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *projectStatusReference) Name() *query.ColumnNode {
	cn := n.projectStatusEnum.Name()
	cn.SetParent(n)
	return cn
}

func (n *projectStatusAssociation) Name() *query.ColumnNode {
	cn := n.projectStatusEnum.Name()
	cn.SetParent(n)
	return cn
}

func (n projectStatusEnum) Description() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "description",
		Identifier:   "Description",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *projectStatusReference) Description() *query.ColumnNode {
	cn := n.projectStatusEnum.Description()
	cn.SetParent(n)
	return cn
}

func (n *projectStatusAssociation) Description() *query.ColumnNode {
	cn := n.projectStatusEnum.Description()
	cn.SetParent(n)
	return cn
}

func (n projectStatusEnum) Guidelines() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "guidelines",
		Identifier:   "Guidelines",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *projectStatusReference) Guidelines() *query.ColumnNode {
	cn := n.projectStatusEnum.Guidelines()
	cn.SetParent(n)
	return cn
}

func (n *projectStatusAssociation) Guidelines() *query.ColumnNode {
	cn := n.projectStatusEnum.Guidelines()
	cn.SetParent(n)
	return cn
}

func (n projectStatusEnum) IsActive() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "is_active",
		Identifier:   "IsActive",
		ReceiverType: query.ColTypeBool,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *projectStatusReference) IsActive() *query.ColumnNode {
	cn := n.projectStatusEnum.IsActive()
	cn.SetParent(n)
	return cn
}

func (n *projectStatusAssociation) IsActive() *query.ColumnNode {
	cn := n.projectStatusEnum.IsActive()
	cn.SetParent(n)
	return cn
}

func (n projectStatusEnum) GobEncode() (data []byte, err error) {
	return
}

func (n *projectStatusEnum) GobDecode(data []byte) (err error) {
	return
}

func (n *projectStatusReference) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *projectStatusReference) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	return
}

func (n *projectStatusAssociation) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ManyManyNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *projectStatusAssociation) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ManyManyNode); err != nil {
		panic(err)
	}
	return
}
