// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// RootNode is the builder interface to the Root nodes.
type RootNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
	// OptionalLeafID represents the optional_leaf_id column in the database.
	OptionalLeafID() *query.ColumnNode
	// OptionalLeaf represents the OptionalLeaf reference to a Leaf object.
	OptionalLeaf() LeafNode
	// RequiredLeafID represents the required_leaf_id column in the database.
	RequiredLeafID() *query.ColumnNode
	// RequiredLeaf represents the RequiredLeaf reference to a Leaf object.
	RequiredLeaf() LeafNode
	// OptionalLeafUniqueID represents the optional_leaf_unique_id column in the database.
	OptionalLeafUniqueID() *query.ColumnNode
	// OptionalLeafUnique represents the OptionalLeafUnique reference to a Leaf object.
	OptionalLeafUnique() LeafNode
	// RequiredLeafUniqueID represents the required_leaf_unique_id column in the database.
	RequiredLeafUniqueID() *query.ColumnNode
	// RequiredLeafUnique represents the RequiredLeafUnique reference to a Leaf object.
	RequiredLeafUnique() LeafNode
	// ParentID represents the parent_id column in the database.
	ParentID() *query.ColumnNode
	// Parent represents the Parent reference to a Root object.
	Parent() RootNode
	// ParentRoots represents the ParentRoots reference to Root objects.
	ParentRoots() RootNode
}

// rootTable represents the root table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the rootTable, call [Root()] to start a reference chain when querying the root table.
type rootTable struct {
}

type rootReference struct {
	rootTable
	query.ReferenceNode
}

type rootReverse struct {
	rootTable
	query.ReverseNode
}

// Root returns a table node that starts a node chain that begins with the root table.
func Root() RootNode {
	return rootTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n rootTable) TableName_() string {
	return "root"
}

// NodeType_ returns the query.NodeType of the node.
func (n rootTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n rootTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n rootTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.OptionalLeafID())
	nodes = append(nodes, n.RequiredLeafID())
	nodes = append(nodes, n.OptionalLeafUniqueID())
	nodes = append(nodes, n.RequiredLeafUniqueID())
	nodes = append(nodes, n.ParentID())
	return nodes
}

func (n *rootReference) ColumnNodes_() (nodes []query.Node) {
	nodes = n.rootTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

func (n *rootReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.rootTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n rootTable) IsEnum_() bool {
	return false
}

func (n *rootReference) NodeType_() query.NodeType {
	return query.ReferenceNodeType
}

func (n *rootReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKey returns a node that points to the primary key column.
func (n rootTable) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *rootReference) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *rootReverse) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

func (n rootTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) ID() *query.ColumnNode {
	cn := n.rootTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) ID() *query.ColumnNode {
	cn := n.rootTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n rootTable) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) Name() *query.ColumnNode {
	cn := n.rootTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) Name() *query.ColumnNode {
	cn := n.rootTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n rootTable) OptionalLeafID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "optional_leaf_id",
		Identifier:   "OptionalLeafID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) OptionalLeafID() *query.ColumnNode {
	cn := n.rootTable.OptionalLeafID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) OptionalLeafID() *query.ColumnNode {
	cn := n.rootTable.OptionalLeafID()
	query.NodeSetParent(cn, n)
	return cn
}

// OptionalLeaf represents the link to a Leaf object.
func (n rootTable) OptionalLeaf() LeafNode {
	cn := &leafReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "optional_leaf_id",
			Identifier:      "OptionalLeaf",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) OptionalLeaf() LeafNode {
	cn := n.rootTable.OptionalLeaf().(*leafReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) OptionalLeaf() LeafNode {
	cn := n.rootTable.OptionalLeaf().(*leafReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n rootTable) RequiredLeafID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "required_leaf_id",
		Identifier:   "RequiredLeafID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) RequiredLeafID() *query.ColumnNode {
	cn := n.rootTable.RequiredLeafID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) RequiredLeafID() *query.ColumnNode {
	cn := n.rootTable.RequiredLeafID()
	query.NodeSetParent(cn, n)
	return cn
}

// RequiredLeaf represents the link to a Leaf object.
func (n rootTable) RequiredLeaf() LeafNode {
	cn := &leafReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "required_leaf_id",
			Identifier:      "RequiredLeaf",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) RequiredLeaf() LeafNode {
	cn := n.rootTable.RequiredLeaf().(*leafReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) RequiredLeaf() LeafNode {
	cn := n.rootTable.RequiredLeaf().(*leafReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n rootTable) OptionalLeafUniqueID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "optional_leaf_unique_id",
		Identifier:   "OptionalLeafUniqueID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) OptionalLeafUniqueID() *query.ColumnNode {
	cn := n.rootTable.OptionalLeafUniqueID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) OptionalLeafUniqueID() *query.ColumnNode {
	cn := n.rootTable.OptionalLeafUniqueID()
	query.NodeSetParent(cn, n)
	return cn
}

// OptionalLeafUnique represents the link to a Leaf object.
func (n rootTable) OptionalLeafUnique() LeafNode {
	cn := &leafReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "optional_leaf_unique_id",
			Identifier:      "OptionalLeafUnique",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) OptionalLeafUnique() LeafNode {
	cn := n.rootTable.OptionalLeafUnique().(*leafReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) OptionalLeafUnique() LeafNode {
	cn := n.rootTable.OptionalLeafUnique().(*leafReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n rootTable) RequiredLeafUniqueID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "required_leaf_unique_id",
		Identifier:   "RequiredLeafUniqueID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) RequiredLeafUniqueID() *query.ColumnNode {
	cn := n.rootTable.RequiredLeafUniqueID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) RequiredLeafUniqueID() *query.ColumnNode {
	cn := n.rootTable.RequiredLeafUniqueID()
	query.NodeSetParent(cn, n)
	return cn
}

// RequiredLeafUnique represents the link to a Leaf object.
func (n rootTable) RequiredLeafUnique() LeafNode {
	cn := &leafReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "required_leaf_unique_id",
			Identifier:      "RequiredLeafUnique",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) RequiredLeafUnique() LeafNode {
	cn := n.rootTable.RequiredLeafUnique().(*leafReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) RequiredLeafUnique() LeafNode {
	cn := n.rootTable.RequiredLeafUnique().(*leafReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n rootTable) ParentID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "parent_id",
		Identifier:   "ParentID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) ParentID() *query.ColumnNode {
	cn := n.rootTable.ParentID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) ParentID() *query.ColumnNode {
	cn := n.rootTable.ParentID()
	query.NodeSetParent(cn, n)
	return cn
}

// Parent represents the link to a Root object.
func (n rootTable) Parent() RootNode {
	cn := &rootReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "parent_id",
			Identifier:      "Parent",
			ReceiverType:    query.ColTypeString,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) Parent() RootNode {
	cn := n.rootTable.Parent().(*rootReference)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) Parent() RootNode {
	cn := n.rootTable.Parent().(*rootReference)
	query.NodeSetParent(cn, n)
	return cn
}

// ParentRoots represents the many-to-one relationship formed by the reverse reference from the
// parent_id column in the root table.
func (n rootTable) ParentRoots() RootNode {
	cn := &rootReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "parent_id",
			Identifier:      "ParentRoots",
			ReceiverType:    query.ColTypeString,
			IsUnique:        false,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReference) ParentRoots() RootNode {
	cn := n.rootTable.ParentRoots().(*rootReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *rootReverse) ParentRoots() RootNode {
	cn := n.rootTable.ParentRoots().(*rootReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n rootTable) GobEncode() (data []byte, err error) {
	return
}

func (n *rootTable) GobDecode(data []byte) (err error) {
	return
}

func (n *rootReference) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *rootReference) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	return
}

func (n *rootReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *rootReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(rootTable))
	gob.Register(new(rootReference))
	gob.Register(new(rootReverse))
}
