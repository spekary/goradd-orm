// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
	"bytes"
	"encoding/gob"

	"github.com/goradd/orm/pkg/query"
)

// AddressNode represents the address table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the AddressNode, call [Address] to start a reference chain when querying the address table.
type AddressNode struct {
	// ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
	// of its functions are exported and are callable along with the addressNode functions here.
	query.ReferenceNodeI
}

// Address returns a table node that starts a node chain that begins with the address table.
func Address() *AddressNode {
	n := AddressNode{
		query.NewTableNode("goradd", "address", "Address"),
	}
	query.SetParentNode(&n, nil)
	return &n
}

// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *AddressNode) SelectNodes_() (nodes []*query.ColumnNode) {
	nodes = append(nodes, n.ID())
	nodes = append(nodes, n.PersonID())
	nodes = append(nodes, n.Street())
	nodes = append(nodes, n.City())
	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *AddressNode) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *AddressNode) Copy_() query.NodeI {
	return &AddressNode{query.CopyNode(n.ReferenceNodeI)}
}

// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *AddressNode) PrimaryKeyNode() *query.ColumnNode {
	return n.ID()
}

// ID represents the id column in the database.
func (n *AddressNode) ID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"address",
		"id",
		"ID",
		query.ColTypeString,
		true,
	)
	query.SetParentNode(cn, n)
	return cn
}

// PersonID represents the person_id column in the database.
func (n *AddressNode) PersonID() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"address",
		"person_id",
		"PersonID",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// Person represents the link to a Person object.
func (n *AddressNode) Person() *PersonNode {
	cn := &PersonNode{
		query.NewReferenceNode(
			"goradd",
			"address",
			"person_id",
			"PersonID",
			"Person",
			"person",
			"id",
			false,
			query.ColTypeString,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}

// Street represents the street column in the database.
func (n *AddressNode) Street() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"address",
		"street",
		"Street",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

// City represents the city column in the database.
func (n *AddressNode) City() *query.ColumnNode {
	cn := query.NewColumnNode(
		"goradd",
		"address",
		"city",
		"City",
		query.ColTypeString,
		false,
	)
	query.SetParentNode(cn, n)
	return cn
}

type addressNodeEncoded struct {
	RefNode query.ReferenceNodeI
}

// GobEncode makes the node serializable.
// doc:hide
func (n *AddressNode) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	s := addressNodeEncoded{
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
func (n *AddressNode) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	var s addressNodeEncoded
	if err = dec.Decode(&s); err != nil {
		panic(err)
	}
	n.ReferenceNodeI = s.RefNode
	query.SetParentNode(n, query.ParentNode(n)) // Reinforce types
	return
}
func init() {
	gob.Register(&AddressNode{})
}
