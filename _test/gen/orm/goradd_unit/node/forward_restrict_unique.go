// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// ForwardRestrictUniqueNode represents the forward_restrict_unique table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the ForwardRestrictUniqueNode, call [ForwardRestrictUnique] to start a reference chain when querying the forward_restrict_unique table.
type ForwardRestrictUniqueNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the forwardRestrictUniqueNode functions here.
	query.ReferenceNodeI
}

// ForwardRestrictUnique returns a table node that starts a node chain that begins with the forward_restrict_unique table.
func ForwardRestrictUnique() *ForwardRestrictUniqueNode {
	n := ForwardRestrictUniqueNode{
		query.NewTableNode("goradd_unit", "forward_restrict_unique", "ForwardRestrictUnique"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *ForwardRestrictUniqueNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.Name())
	nodes = append(nodes, n.ReverseID())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *ForwardRestrictUniqueNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *ForwardRestrictUniqueNode) Copy_() query.NodeI {
	return &ForwardRestrictUniqueNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *ForwardRestrictUniqueNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *ForwardRestrictUniqueNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_restrict_unique",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Name represents the name column in the database.
func (n *ForwardRestrictUniqueNode) Name() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_restrict_unique",
		"name",
		"Name",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// ReverseID represents the reverse_id column in the database.
func (n *ForwardRestrictUniqueNode) ReverseID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd_unit",
		"forward_restrict_unique",
		"reverse_id",
		"ReverseID",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Reverse represents the link to a Reverse object.
func (n *ForwardRestrictUniqueNode) Reverse() *ReverseNode {
	cn := &ReverseNode{
		query.NewReferenceNode(
			"goradd_unit",
			"forward_restrict_unique",
			"reverse_id",
			"ReverseID",
			"Reverse",
			"reverse",
			"id",
			false,
			query.ColTypeString,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

type forwardRestrictUniqueNodeEncoded struct {
	RefNode query.ReferenceNodeI
}

// GobEncode makes the node serializable.
// doc:hide
func (n *ForwardRestrictUniqueNode) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	s := forwardRestrictUniqueNodeEncoded{
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
func (n *ForwardRestrictUniqueNode) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	var s forwardRestrictUniqueNodeEncoded
	if err = dec.Decode(&s); err != nil {
		panic(err)
	}
	n.ReferenceNodeI = s.RefNode
	query.SetParentNode(n, query.ParentNode(n)) // Reinforce types
	return
}
func init() {
	gob.Register(&ForwardRestrictUniqueNode{})
}
