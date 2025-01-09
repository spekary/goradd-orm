// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// ForwardRestrictNodeI is the builder interface to the ForwardRestrict nodes.
type ForwardRestrictNodeI interface {
	query.NodeI
	PrimaryKeyNode() *query.ColumnNode
	// ID represents the id column in the database.
	ID() *query.ColumnNode
	// Name represents the name column in the database.
	Name() *query.ColumnNode
	// ReverseID represents the reverse_id column in the database.
	ReverseID() *query.ColumnNode
	// Reverse represents the Reverse reference to a Reverse object.
	Reverse() ReverseNodeI
}

// ForwardRestrictExpander is the builder interface for ForwardRestricts that are expandable.
type ForwardRestrictExpander interface {
	ForwardRestrictNodeI
	// Expand causes the node to produce separate rows with individual items, rather than a single row with an array of items.
	Expand()
}

// forwardRestrictTable represents the forward_restrict table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node chain in a query.
//
// To use the forwardRestrictTable, call [ForwardRestrict()] to start a reference chain when querying the forward_restrict table.
type forwardRestrictTable struct {
}

type forwardRestrictReverse struct {
	forwardRestrictTable
	query.ReverseNode
}

// ForwardRestrict returns a table node that starts a node chain that begins with the forward_restrict table.
func ForwardRestrict() ForwardRestrictNodeI {
	return forwardRestrictTable{}
}

// TableName_ returns the query name of the table the node is associated with.
func (n forwardRestrictTable) TableName_() string {
	return "forward_restrict"
}

// NodeType_ returns the query.NodeType of the node.
func (n forwardRestrictTable) NodeType_() query.NodeType {
	return query.TableNodeType
}

// DatabaseKey_ returns the database key of the database the node is associated with.
func (n forwardRestrictTable) DatabaseKey_() string {
	return "goradd_unit"
}

// ColumnNodes_ is used internally by the framework to return the list of all the column nodes.
// This may include reference nodes to enum types.
func (n forwardRestrictTable) ColumnNodes_() (nodes []query.NodeI) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.ReverseID())
	return nodes
}

// Columns_ is used internally by the framework to return the list of all the columns in the table.
func (n forwardRestrictTable) Columns_() []string {
	return []string{
		"id",
		"name",
		"reverse_id",
	}
}

// IsEnum_ is used internally by the framework to determine if the current table is an enumerated type.
func (n forwardRestrictTable) IsEnum_() bool {
	return false
}

func (n *forwardRestrictReverse) NodeType_() query.NodeType {
	return query.ReverseNodeType
}

// PrimaryKeyNode returns a node that points to the primary key column.
func (n forwardRestrictTable) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n *forwardRestrictReverse) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

func (n forwardRestrictTable) ID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "id",
		Identifier:   "ID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: true,
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardRestrictReverse) ID() *query.ColumnNode {
	cn := n.forwardRestrictTable.ID()
	cn.SetParent(n)
	return cn
}

func (n forwardRestrictTable) Name() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "name",
		Identifier:   "Name",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardRestrictReverse) Name() *query.ColumnNode {
	cn := n.forwardRestrictTable.Name()
	cn.SetParent(n)
	return cn
}

func (n forwardRestrictTable) ReverseID() *query.ColumnNode {
	cn := &query.ColumnNode{
		QueryName:    "reverse_id",
		Identifier:   "ReverseID",
		ReceiverType: query.ColTypeString,
		IsPrimaryKey: false,
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardRestrictReverse) ReverseID() *query.ColumnNode {
	cn := n.forwardRestrictTable.ReverseID()
	cn.SetParent(n)
	return cn
}

// Reverse represents the link to a Reverse object.
func (n forwardRestrictTable) Reverse() ReverseNodeI {
	cn := &reverseReference{
		ReferenceNode: query.ReferenceNode{
			ColumnQueryName: "reverse_id",
			Identifier:      "ReverseID",
			ReceiverType:    query.ColTypeString,
		},
	}
	cn.SetParent(n)
	return cn
}

func (n *forwardRestrictReverse) Reverse() ReverseNodeI {
	cn := n.forwardRestrictTable.Reverse().(*reverseReference)
	cn.SetParent(n)
	return cn
}

func (n forwardRestrictTable) GobEncode() (data []byte, err error) {
	return
}

func (n *forwardRestrictTable) GobDecode(data []byte) (err error) {
	return
}

func (n *forwardRestrictReverse) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(n.ReverseNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *forwardRestrictReverse) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	if err = dec.Decode(&n.ReverseNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(new(forwardRestrictTable))
	gob.Register(new(forwardRestrictReverse))
}
