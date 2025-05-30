// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
	"github.com/goradd/orm/pkg/schema"
)

// LeafNlNode is the builder interface to the LeafNl nodes.
type LeafNlNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
	// RootNlID represents the root_nl_id column in the database.
	RootNlID() *query.ColumnNode
	// RootNl represents the RootNl reference to a RootNl object.
	RootNl() RootNlNode
	// GroLock represents the gro_lock column in the database.
	GroLock() *query.ColumnNode
	// Leaf2s represents the Leaf2s reference to LeafNl objects.
	Leaf2s() LeafNlNode
	// Leaf1s represents the Leaf1s reference to LeafNl objects.
	Leaf1s() LeafNlNode
}

// leafNlTable represents the leaf_nl table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the leafNlTable, call [LeafNl()] to start a reference chain when querying the leaf_nl table.
type leafNlTable struct {
}

type leafNlReverse struct {
	leafNlTable
	query.ReverseNode
}

type leafNlAssociation struct {
	leafNlTable
	query.ManyManyNode
}

// LeafNl returns a table node that starts a node chain that begins with the leaf_nl table.
func LeafNl() LeafNlNode {
	return leafNlTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n leafNlTable) TableName_() string {
	return "leaf_nl"
}

// NodeType_ returns the query.NodeType of the node.
func (n leafNlTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n leafNlTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n leafNlTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.RootNlID())
	nodes = append(nodes, n.GroLock())
	return nodes
}

func (n *leafNlReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.leafNlTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

func (n *leafNlAssociation) ColumnNodes_() (nodes []query.Node) {
	nodes = n.leafNlTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

func (n *leafNlReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

func (n *leafNlAssociation) NodeType_() query.NodeType {
	return query.ManyManyNodeType
}

// PrimaryKey returns a node that points to the primary key column.
func (n leafNlTable) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *leafNlReverse) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *leafNlAssociation) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

func (n leafNlTable) ID() *query.ColumnNode {
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

func (n *leafNlReverse) ID() *query.ColumnNode {
	cn := n.leafNlTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlAssociation) ID() *query.ColumnNode {
	cn := n.leafNlTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n leafNlTable) Name() *query.ColumnNode {
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

func (n *leafNlReverse) Name() *query.ColumnNode {
	cn := n.leafNlTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlAssociation) Name() *query.ColumnNode {
	cn := n.leafNlTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n leafNlTable) RootNlID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:     "root_nl_id",
		Identifier:    "RootNlID",
		ReceiverType:  query.ColTypeString,
		SchemaType:    schema.ColTypeReference,
		SchemaSubType: schema.ColSubTypeNone,
		IsPrimaryKey:  false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlReverse) RootNlID() *query.ColumnNode {
	cn := n.leafNlTable.RootNlID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlAssociation) RootNlID() *query.ColumnNode {
	cn := n.leafNlTable.RootNlID()
	query.NodeSetParent(cn, n)
	return cn
}

// RootNl represents the link to a RootNl object.
func (n leafNlTable) RootNl() RootNlNode {
	cn := &rootNlReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "root_nl_id",
			Identifier:      "RootNl",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlReverse) RootNl() RootNlNode {
	cn := n.leafNlTable.RootNl().(*rootNlReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlAssociation) RootNl() RootNlNode {
	cn := n.leafNlTable.RootNl().(*rootNlReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n leafNlTable) GroLock() *query.ColumnNode {
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

func (n *leafNlReverse) GroLock() *query.ColumnNode {
	cn := n.leafNlTable.GroLock()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlAssociation) GroLock() *query.ColumnNode {
	cn := n.leafNlTable.GroLock()
	query.NodeSetParent(cn, n)
	return cn
}

// Leaf2s represents the many-to-many relationship formed by the leaf_nl_assn table.
func (n leafNlTable) Leaf2s() LeafNlNode {
	cn := &leafNlAssociation{
		ManyManyNode: query.ManyManyNode{
			AssnTableQueryName:       "leaf_nl_assn",
			ParentColumnQueryName:    "leaf_1_id",
			ParentColumnReceiverType: query.ColTypeString,
			Identifier:               "Leaf2s",
			RefColumnQueryName:       "leaf_2_id",
			RefColumnReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlReverse) Leaf2s() LeafNlNode {
	cn := n.leafNlTable.Leaf2s().(*leafNlAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlAssociation) Leaf2s() LeafNlNode {
	cn := n.leafNlTable.Leaf2s().(*leafNlAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

// Leaf1s represents the many-to-many relationship formed by the leaf_nl_assn table.
func (n leafNlTable) Leaf1s() LeafNlNode {
	cn := &leafNlAssociation{
		ManyManyNode: query.ManyManyNode{
			AssnTableQueryName:       "leaf_nl_assn",
			ParentColumnQueryName:    "leaf_2_id",
			ParentColumnReceiverType: query.ColTypeString,
			Identifier:               "Leaf1s",
			RefColumnQueryName:       "leaf_1_id",
			RefColumnReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlReverse) Leaf1s() LeafNlNode {
	cn := n.leafNlTable.Leaf1s().(*leafNlAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *leafNlAssociation) Leaf1s() LeafNlNode {
	cn := n.leafNlTable.Leaf1s().(*leafNlAssociation)
	query.NodeSetParent(cn, n)
	return cn
}

func (n leafNlTable) GobEncode() (data []byte, err error) {
	return
}

func (n *leafNlTable) GobDecode(data []byte) (err error) {
	return
}

func (n *leafNlReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *leafNlReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func (n *leafNlAssociation) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ManyManyNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *leafNlAssociation) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ManyManyNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(leafNlTable))
	gob.Register(new(leafNlReverse))
	gob.Register(new(leafNlAssociation))
}
