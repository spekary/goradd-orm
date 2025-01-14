//** This file was code generated by GoT. DO NOT EDIT. ***

package template

import (
	"io"
	"path/filepath"

	"github.com/goradd/orm/pkg/codegen"
	"github.com/goradd/orm/pkg/model"
)

func init() {
	t := EnumNodeTemplate{}
	codegen.RegisterTemplate(&t)
}

// EnumNodeTemplate generates the node code for enumerated types
type EnumNodeTemplate struct {
	Package string
}

func (tmpl *EnumNodeTemplate) FileName(table *model.Enum) string {
	return filepath.Join("orm", table.DbKey, "node", table.FileName()+".go")
}

func (tmpl *EnumNodeTemplate) GenerateEnum(table *model.Enum, _w io.Writer, importPath string) (err error) {
	return tmpl.gen(table, _w)
}

func (tmpl *EnumNodeTemplate) Overwrite() bool {
	return true
}

//*** enum_node.tmpl

func (tmpl *EnumNodeTemplate) gen(table *model.Enum, _w io.Writer) (err error) {
	if err = tmpl.genHeader(table, _w); err != nil {
		return
	}
	if err = tmpl.genPrivate(table, _w); err != nil {
		return
	}
	if err = tmpl.genLeafNodes(table, _w); err != nil {
		return
	}
	return
}

func (tmpl *EnumNodeTemplate) genHeader(table *model.Enum, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// Code generated by goradd-orm. DO NOT EDIT.

package node

import (
    "encoding/gob"
	"github.com/goradd/orm/pkg/query"
)

type `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node struct {
	query.ReferenceNodeI
}

// PrimaryKeyNode returns a node representing the primary key column.
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

	if _, err = io.WriteString(_w, table.FieldIdentifier(0)); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `()
}

func init() {
   gob.Register(new(`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node))
}

`); err != nil {
		return
	}

	return
}

func (tmpl *EnumNodeTemplate) genPrivate(table *model.Enum, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// SelectNodes_ is used internally by the framework to return the list of column nodes.
// doc: hide
func (n *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Node) SelectNodes_() ([]*query.ColumnNode) {
    return []*query.ColumnNode {
`); err != nil {
		return
	}

	for i := 0; i < len(table.Fields); i++ {

		if _, err = io.WriteString(_w, `	    n.`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.FieldIdentifier(i)); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(),
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `	}
}

// EmbeddedNode_ is used internally by the framework to return the embedded ReferenceNodeI.
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

func (tmpl *EnumNodeTemplate) genLeafNodes(table *model.Enum, _w io.Writer) (err error) {

	for i := 0; i < len(table.Fields); i++ {
		fn := table.FieldIdentifier(i)

		if _, err = io.WriteString(_w, `
func (n *`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `Node) `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, fn); err != nil {
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

		if _, err = io.WriteString(_w, table.FieldQueryName(i)); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `",
		"`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, fn); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `",
		query.`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.FieldReceiverType(i).String()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `,
		`); err != nil {
			return
		}

		if i == 0 {

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

	}
	return
}
