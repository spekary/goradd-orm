// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
	"github.com/goradd/orm/pkg/schema"
)

// AltLeafUnNode is the builder interface to the AltLeafUn nodes.
type AltLeafUnNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
	// AltRootUnID represents the alt_root_un_id column in the database.
	AltRootUnID() *query.ColumnNode
	// AltRootUn represents the AltRootUn reference to a AltRootUn object.
	AltRootUn() AltRootUnNode
}

// altLeafUnTable represents the alt_leaf_un table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the altLeafUnTable, call [AltLeafUn()] to start a reference chain when querying the alt_leaf_un table.
type altLeafUnTable struct {
}

type altLeafUnReverse struct {
	altLeafUnTable
	query.ReverseNode
}

// AltLeafUn returns a table node that starts a node chain that begins with the alt_leaf_un table.
func AltLeafUn() AltLeafUnNode {
	return altLeafUnTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n altLeafUnTable) TableName_() string {
	return "alt_leaf_un"
}

// NodeType_ returns the query.NodeType of the node.
func (n altLeafUnTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n altLeafUnTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n altLeafUnTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.AltRootUnID())
	return nodes
}

func (n *altLeafUnReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.altLeafUnTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

func (n *altLeafUnReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKey returns a node that points to the primary key column.
func (n altLeafUnTable) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *altLeafUnReverse) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

func (n altLeafUnTable) ID() *query.ColumnNode {
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

func (n *altLeafUnReverse) ID() *query.ColumnNode {
	cn := n.altLeafUnTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n altLeafUnTable) Name() *query.ColumnNode {
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

func (n *altLeafUnReverse) Name() *query.ColumnNode {
	cn := n.altLeafUnTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n altLeafUnTable) AltRootUnID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "alt_root_un_id",
		Identifier:    "AltRootUnID",
		ReceiverType:  query.ColTypeFloat32,
		SchemaType:    schema.ColTypeReference,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *altLeafUnReverse) AltRootUnID() *query.ColumnNode {
	cn := n.altLeafUnTable.AltRootUnID()
	query.NodeSetParent(cn, n)
	return cn
}

// AltRootUn represents the link to a AltRootUn object.
func (n altLeafUnTable) AltRootUn() AltRootUnNode {
	cn := &altRootUnReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "alt_root_un_id",
			Identifier:      "AltRootUn",
			ReceiverType:    query.ColTypeFloat32,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *altLeafUnReverse) AltRootUn() AltRootUnNode {
	cn := n.altLeafUnTable.AltRootUn().(*altRootUnReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n altLeafUnTable) GobEncode() (data []byte, err error) {
	return
}

func (n *altLeafUnTable) GobDecode(data []byte) (err error) {
	return
}

func (n *altLeafUnReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *altLeafUnReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(altLeafUnTable))
	gob.Register(new(altLeafUnReverse))
}
