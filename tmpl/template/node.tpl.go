//** This file was code generated by GoT. DO NOT EDIT. ***

package template

import (
	"fmt"
	"io"
	"path/filepath"
	"strconv"

	"github.com/goradd/orm/pkg/codegen"
	"github.com/goradd/orm/pkg/model"
)

func init() {
	codegen.RegisterTemplate(new(NodeTemplate))
}

// NodeTemplate generates the code for table nodes.
type NodeTemplate struct {
}

func (n *NodeTemplate) FileName(table *model.Table) string {
	return filepath.Join("orm", table.DbKey, "node", table.FileName()+".go")
}

func (n *NodeTemplate) GenerateTable(table *model.Table, _w io.Writer, importPath string) (err error) {
	return n.gen(table, _w)
}

func (n *NodeTemplate) Overwrite() bool {
	return true
}

//*** node.tmpl

func (n *NodeTemplate) gen(table *model.Table, _w io.Writer) (err error) {
	if err = n.genHeader(table, _w); err != nil {
		return
	}
	if err = n.genStruct(table, _w); err != nil {
		return
	}
	if err = n.genTableNode(table, _w); err != nil {
		return
	}
	if err = n.genPrivate(table, _w); err != nil {
		return
	}
	if err = n.genPrimaryKey(table, _w); err != nil {
		return
	}
	if err = n.genColumns(table, _w); err != nil {
		return
	}
	if err = n.genAssn(table, _w); err != nil {
		return
	}
	if err = n.genReverse(table, _w); err != nil {
		return
	}
	if err = n.genGob(table, _w); err != nil {
		return
	}
	if err = n.genInit(table, _w); err != nil {
		return
	}
	return
}

func (n *NodeTemplate) genHeader(table *model.Table, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// Code generated by goradd-orm. DO NOT EDIT.

package node


import (
    "bytes"
    "encoding/gob"
	"github.com/goradd/orm/pkg/query"
)
`); err != nil {
		return
	}

	return
}

func (n *NodeTemplate) genStruct(table *model.Table, _w io.Writer) (err error) {

	//*** struct.tmpl

	if _, err = io.WriteString(_w, `// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node represents the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` table in a query. It uses a builder pattern to chain
// together other tables and columns to form a node in a query.
//
// To use the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node, call [`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `] to start a reference chain when querying the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` table.
type `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node struct {
    // ReferenceNodeI is an internal object that represents the capabilities of the node. Since it is embedded, all
    // of its functions are exported and are callable along with the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DecapIdentifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node functions here.
	query.ReferenceNodeI
}
`); err != nil {
		return
	}

	return
}

func (n *NodeTemplate) genTableNode(table *model.Table, _w io.Writer) (err error) {

	//*** table.tmpl

	if _, err = io.WriteString(_w, `// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` returns a table node that starts a node chain that begins with the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` table.
func `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `() *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node {
	n := `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node {
		query.NewTableNode("`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DbKey); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `", "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",  "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `"),
	}
	query.SetParentNode(&n, nil)
	return &n
}
`); err != nil {
		return
	}

	return
}

func (n *NodeTemplate) genPrivate(table *model.Table, _w io.Writer) (err error) {

	//*** private.tmpl

	if _, err = io.WriteString(_w, `// SelectNodes_ is used internally by the framework to return the list of all the column nodes.
// doc: hide
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) SelectNodes_() (nodes []*query.ColumnNode) {
`); err != nil {
		return
	}

	for _, col := range table.Columns {

		if _, err = io.WriteString(_w, `	nodes = append(nodes, n.`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, fmt.Sprint(col.Identifier)); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `())
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `	return nodes
}

// EmbeddedNode is used internally by the framework to return the embedded Reference node.
// doc: hide
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) EmbeddedNode_() query.NodeI {
	return n.ReferenceNodeI
}

// Copy_ is used internally by the framework to deep copy the node.
// doc: hide
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) Copy_() query.NodeI {
	return &`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node{query.CopyNode(n.ReferenceNodeI)}
}
`); err != nil {
		return
	}

	return
}

func (n *NodeTemplate) genPrimaryKey(table *model.Table, _w io.Writer) (err error) {
	if col := table.PrimaryKeyColumn(); col != nil {

		//*** pk.tmpl

		if _, err = io.WriteString(_w, `// PrimaryKeyNode returns a node that points to the primary key column, if
// a single primary key exists in the table.
func (n *`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `Node) PrimaryKeyNode() (*query.ColumnNode) {
    return n.`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, fmt.Sprint(col.Identifier)); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `()
}
`); err != nil {
			return
		}

	}
	return
}

func (n *NodeTemplate) genColumns(table *model.Table, _w io.Writer) (err error) {
	for _, col := range table.Columns {
		if err = n.genColumn(table, col, _w); err != nil {
			return
		}
	}
	return
}

func (n *NodeTemplate) genGob(table *model.Table, _w io.Writer) (err error) {

	//*** gob.tmpl

	if _, err = io.WriteString(_w, `
type `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DecapIdentifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `NodeEncoded struct {
	RefNode query.ReferenceNodeI
}

// GobEncode makes the node serializable.
// doc:hide
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) GobEncode() (data []byte, err error) {
	var buf bytes.Buffer
	e := gob.NewEncoder(&buf)

	s := `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DecapIdentifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `NodeEncoded {
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
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) GobDecode(data []byte) (err error) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	var s `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DecapIdentifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `NodeEncoded
	if err = dec.Decode(&s); err != nil {
		panic(err)
	}
	n.ReferenceNodeI = s.RefNode
	query.SetParentNode(n, query.ParentNode(n)) // Reinforce types
	return
}
`); err != nil {
		return
	}

	return
}

func (n *NodeTemplate) genInit(table *model.Table, _w io.Writer) (err error) {

	//*** init.tmpl

	if _, err = io.WriteString(_w, `func init() {
	gob.Register(&`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node{})
}
`); err != nil {
		return
	}

	return
}

//*** column.tmpl

func (n *NodeTemplate) genColumn(table *model.Table, col *model.Column, _w io.Writer) (err error) {

	if col.IsReference() {
		if err = n.genColumnNode(table, col, _w); err != nil {
			return
		}
		if err = n.genTableRefNode(table, col, _w); err != nil {
			return
		}
	} else {
		if err = n.genColumnNode(table, col, _w); err != nil {
			return
		}
	}
	return
}

func (n *NodeTemplate) genColumnNode(table *model.Table, col *model.Column, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` represents the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` column in the database.
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `() *query.ColumnNode {
	cn := query.NewColumnNode (
		"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DbKey); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
		"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
		"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
		"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
		query.`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Type.String()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `,
		`); err != nil {
		return
	}

	if col.IsPk {

		if _, err = io.WriteString(_w, `true`); err != nil {
			return
		}

	} else {

		if _, err = io.WriteString(_w, `false`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `,
	)
	query.SetParentNode(cn, n)
	return cn
}
`); err != nil {
		return
	}

	return
}

func (n *NodeTemplate) genTableRefNode(table *model.Table, col *model.Column, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.ReferenceIdentifier()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` represents the link to a `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.ReferenceType()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` object.
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.ReferenceIdentifier()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `() *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.ReferenceType()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node {
	cn := &`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.ReferenceType()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node {
		query.NewReferenceNode (
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DbKey); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.ReferenceIdentifier()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Reference.Table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Reference.Table.PrimaryKeyColumn().QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			false,
			query.`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Type.String()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}
`); err != nil {
		return
	}

	return
}

func (n *NodeTemplate) genEnumTableRefNode(table *model.Table, col *model.Column, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.ReferenceIdentifier()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` represents the link to a `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Reference.EnumTable.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` object.
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.ReferenceIdentifier()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `() *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Reference.EnumTable.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node {
    cn := &`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Reference.EnumTable.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node {
        query.NewReferenceNode (
            "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DbKey); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
            "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
            "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
            "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
            "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Reference.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
            "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Reference.EnumTable.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
            "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Reference.EnumTable.PkQueryName()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
            true,
            query.`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, col.Type.String()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `,
        ),
    }
    query.SetParentNode(cn, n)
    return cn
}
`); err != nil {
		return
	}

	return
}

//*** assn.tmpl

func (n *NodeTemplate) genAssn(table *model.Table, _w io.Writer) (err error) {
	for _, mm := range table.ManyManyReferences {
		if err = n.genAssnTable(table, mm, _w); err != nil {
			return
		}
	}
	return
}

func (n *NodeTemplate) genAssnTable(table *model.Table, mm *model.ManyManyReference, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` represents the many-to-many relationship formed by the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.AssnTableName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` table.
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `() *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.ObjectType()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node  {
	cn := &`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.ObjectType()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node {
		query.NewManyManyNode (
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DbKey); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.AssnTableName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.AssnSourceColumnName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.QueryName()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.AssnDestColumnName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, mm.PrimaryKey()); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, strconv.FormatBool(mm.DestinationEnumTable != nil)); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}
`); err != nil {
		return
	}

	return
}

//*** reverse.tmpl

func (n *NodeTemplate) genReverse(table *model.Table, _w io.Writer) (err error) {
	for _, rev := range table.ReverseReferences {
		if rev.IsUnique {
			if err = n.genReverseOne(table, rev, _w); err != nil {
				return
			}
		} else {
			if err = n.genReverseMany(table, rev, _w); err != nil {
				return
			}
		}
	}
	return
}

func (n *NodeTemplate) genReverseOne(table *model.Table, rev *model.Column, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Reference.ReverseIdentifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` represents the one-to-one relationship formed by the reverse reference from the
// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` column in the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` table.
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Reference.ReverseIdentifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `() *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node  {

	cn := &`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node {
		query.NewReverseReferenceNode (
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DbKey); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.PrimaryKeyColumn().QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Reference.ReverseIdentifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			false,
		),
	}
	query.SetParentNode(cn, n)
	return cn

}
`); err != nil {
		return
	}

	return
}

func (n *NodeTemplate) genReverseMany(table *model.Table, rev *model.Column, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `
// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Reference.ReverseIdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` represents the many-to-one relationship formed by the reverse reference from the
// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` column in the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` table.
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Reference.ReverseIdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `() *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node  {
	cn := &`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node {
		query.NewReverseReferenceNode (
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DbKey); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.PrimaryKeyColumn().QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Reference.ReverseIdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.Table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			"`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, rev.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `",
			true,
		),
	}
	query.SetParentNode(cn, n)
	return cn
}
`); err != nil {
		return
	}

	return
}
