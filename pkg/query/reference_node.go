package query

import (
	"bytes"
	"encoding/gob"
)

type ReferenceNodeI interface {
	NodeI
	conditioner
	NodeLinker
}

// A ReferenceNode is a forward-pointing foreign key relationship. If the other side of the relationship is
// not a enum table, then the other table will have a matching ReverseNode.
type ReferenceNode struct {
	// The query name of the column that is the foreign key
	ColumnQueryName string
	// The identifier that will be used to identify this object in source code.
	// Equals the key for the Get() function on an object.
	Identifier string
	// The type of item acting as a pointer. This should be the same on both sides of the reference.
	ReceiverType ReceiverType
	nodeCondition
	nodeLink
}

func (n *ReferenceNode) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	if err = e.Encode(n.ColumnQueryName); err != nil {
		panic(err)
	}
	if err = e.Encode(n.Identifier); err != nil {
		panic(err)
	}
	if err = e.Encode(n.ReceiverType); err != nil {
		panic(err)
	}
	if err = e.Encode(&n.nodeCondition.condition); err != nil {
		panic(err)
	}
	if err = e.Encode(&n.nodeLink.parentNode); err != nil {
		panic(err)
	}
	data = buf.Bytes()
	return
}

func (n *ReferenceNode) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	if err = dec.Decode(&n.ColumnQueryName); err != nil {
		panic(err)
	}
	if err = dec.Decode(&n.Identifier); err != nil {
		panic(err)
	}
	if err = dec.Decode(&n.ReceiverType); err != nil {
		panic(err)
	}
	if err = dec.Decode(&n.nodeCondition.condition); err != nil {
		panic(err)
	}
	if err = dec.Decode(&n.nodeLink.parentNode); err != nil {
		panic(err)
	}
	return
}

func init() {
	gob.Register(&ReferenceNode{})
}