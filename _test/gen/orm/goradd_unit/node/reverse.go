// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// ReverseNode is the builder interface to the Reverse nodes.
type ReverseNode interface {
	query.TableNodeI
	PrimaryKey() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
	// ForwardCascades represents the ForwardCascades reference to ForwardCascade objects.
	ForwardCascades() ForwardCascadeNode
	// ForwardCascadeUnique represents the ForwardCascadeUnique reference to a ForwardCascadeUnique object.
	ForwardCascadeUnique() ForwardCascadeUniqueNode
	// ForwardNulls represents the ForwardNulls reference to ForwardNull objects.
	ForwardNulls() ForwardNullNode
	// ForwardNullUnique represents the ForwardNullUnique reference to a ForwardNullUnique object.
	ForwardNullUnique() ForwardNullUniqueNode
	// ForwardRestricts represents the ForwardRestricts reference to ForwardRestrict objects.
	ForwardRestricts() ForwardRestrictNode
	// ForwardRestrictUnique represents the ForwardRestrictUnique reference to a ForwardRestrictUnique object.
	ForwardRestrictUnique() ForwardRestrictUniqueNode
}

// reverseTable represents the reverse table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the reverseTable, call [Reverse()] to start a reference chain when querying the reverse table.
type reverseTable struct {
}

type reverseReference struct {
	reverseTable
	query.ReferenceNode
}

type reverseReverse struct {
	reverseTable
	query.ReverseNode
}

// Reverse returns a table node that starts a node chain that begins with the reverse table.
func Reverse() ReverseNode {
	return reverseTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n reverseTable) TableName_() string {
	return "reverse"
}

// NodeType_ returns the query.NodeType of the node.
func (n reverseTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n reverseTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ returns a list of all the column nodes in this node.
func (n reverseTable) ColumnNodes_() (nodes []query.Node) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	return nodes
}

func (n *reverseReference) ColumnNodes_() (nodes []query.Node) {
	nodes = n.reverseTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

func (n *reverseReverse) ColumnNodes_() (nodes []query.Node) {
	nodes = n.reverseTable.ColumnNodes_()
	for _, cn := range nodes {
		query.NodeSetParent(cn, n)
	}
	return
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n reverseTable) IsEnum_() bool {
	return false
}

func (n *reverseReference) NodeType_() query.NodeType {
	return query.ReferenceNodeType
}

func (n *reverseReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKey returns a node that points to the primary key column.
func (n reverseTable) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *reverseReference) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

// PrimaryKey returns a node that points to the primary key column.
func (n *reverseReverse) PrimaryKey() *query.ColumnNode {
	return n.ID()
}

func (n reverseTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReference) ID() *query.ColumnNode {
	cn := n.reverseTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReverse) ID() *query.ColumnNode {
	cn := n.reverseTable.ID()
	query.NodeSetParent(cn, n)
	return cn
}

func (n reverseTable) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReference) Name() *query.ColumnNode {
	cn := n.reverseTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReverse) Name() *query.ColumnNode {
	cn := n.reverseTable.Name()
	query.NodeSetParent(cn, n)
	return cn
}

// ForwardCascades represents the many-to-one relationship formed by the reverse reference from the
// reverse_id column in the forward_cascade table.
func (n reverseTable) ForwardCascades() ForwardCascadeNode {
	cn := &forwardCascadeReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "reverse_id",
			Identifier:      "ForwardCascades",
			ReceiverType:    query.ColTypeString,
			IsUnique:        false,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReference) ForwardCascades() ForwardCascadeNode {
	cn := n.reverseTable.ForwardCascades().(*forwardCascadeReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReverse) ForwardCascades() ForwardCascadeNode {
	cn := n.reverseTable.ForwardCascades().(*forwardCascadeReverse)
	query.NodeSetParent(cn, n)
	return cn
}

// ForwardCascadeUnique represents the one-to-one relationship formed by the reverse reference from the
// reverse_id column in the forward_cascade_unique table.
func (n reverseTable) ForwardCascadeUnique() ForwardCascadeUniqueNode {
	cn := &forwardCascadeUniqueReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "reverse_id",
			Identifier:      "ForwardCascadeUnique",
			ReceiverType:    query.ColTypeString,
			IsUnique:        true,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReference) ForwardCascadeUnique() ForwardCascadeUniqueNode {
	cn := n.reverseTable.ForwardCascadeUnique().(*forwardCascadeUniqueReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReverse) ForwardCascadeUnique() ForwardCascadeUniqueNode {
	cn := n.reverseTable.ForwardCascadeUnique().(*forwardCascadeUniqueReverse)
	query.NodeSetParent(cn, n)
	return cn
}

// ForwardNulls represents the many-to-one relationship formed by the reverse reference from the
// reverse_id column in the forward_null table.
func (n reverseTable) ForwardNulls() ForwardNullNode {
	cn := &forwardNullReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "reverse_id",
			Identifier:      "ForwardNulls",
			ReceiverType:    query.ColTypeString,
			IsUnique:        false,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReference) ForwardNulls() ForwardNullNode {
	cn := n.reverseTable.ForwardNulls().(*forwardNullReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReverse) ForwardNulls() ForwardNullNode {
	cn := n.reverseTable.ForwardNulls().(*forwardNullReverse)
	query.NodeSetParent(cn, n)
	return cn
}

// ForwardNullUnique represents the one-to-one relationship formed by the reverse reference from the
// reverse_id column in the forward_null_unique table.
func (n reverseTable) ForwardNullUnique() ForwardNullUniqueNode {
	cn := &forwardNullUniqueReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "reverse_id",
			Identifier:      "ForwardNullUnique",
			ReceiverType:    query.ColTypeString,
			IsUnique:        true,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReference) ForwardNullUnique() ForwardNullUniqueNode {
	cn := n.reverseTable.ForwardNullUnique().(*forwardNullUniqueReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReverse) ForwardNullUnique() ForwardNullUniqueNode {
	cn := n.reverseTable.ForwardNullUnique().(*forwardNullUniqueReverse)
	query.NodeSetParent(cn, n)
	return cn
}

// ForwardRestricts represents the many-to-one relationship formed by the reverse reference from the
// reverse_id column in the forward_restrict table.
func (n reverseTable) ForwardRestricts() ForwardRestrictNode {
	cn := &forwardRestrictReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "reverse_id",
			Identifier:      "ForwardRestricts",
			ReceiverType:    query.ColTypeString,
			IsUnique:        false,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReference) ForwardRestricts() ForwardRestrictNode {
	cn := n.reverseTable.ForwardRestricts().(*forwardRestrictReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReverse) ForwardRestricts() ForwardRestrictNode {
	cn := n.reverseTable.ForwardRestricts().(*forwardRestrictReverse)
	query.NodeSetParent(cn, n)
	return cn
}

// ForwardRestrictUnique represents the one-to-one relationship formed by the reverse reference from the
// reverse_id column in the forward_restrict_unique table.
func (n reverseTable) ForwardRestrictUnique() ForwardRestrictUniqueNode {
	cn := &forwardRestrictUniqueReverse{
		ReverseNode: query.ReverseNode{
			ColumnQueryName: "reverse_id",
			Identifier:      "ForwardRestrictUnique",
			ReceiverType:    query.ColTypeString,
			IsUnique:        true,
		},
	}
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReference) ForwardRestrictUnique() ForwardRestrictUniqueNode {
	cn := n.reverseTable.ForwardRestrictUnique().(*forwardRestrictUniqueReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n *reverseReverse) ForwardRestrictUnique() ForwardRestrictUniqueNode {
	cn := n.reverseTable.ForwardRestrictUnique().(*forwardRestrictUniqueReverse)
	query.NodeSetParent(cn, n)
	return cn
}

func (n reverseTable) GobEncode() (data []byte, err error) {
	return
}

func (n *reverseTable) GobDecode(data []byte) (err error) {
	return
}

func (n *reverseReference) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *reverseReference) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReferenceNode); err != nil {
		panic(err)
	}
	return
}

func (n *reverseReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(&n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *reverseReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(reverseTable))
	gob.Register(new(reverseReference))
	gob.Register(new(reverseReverse))
}
