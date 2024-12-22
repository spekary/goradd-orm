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
	t := EnumTableTemplate{}
	codegen.RegisterTemplate(&t)
}

type EnumTableTemplate struct {
	Package string
}

func (tmpl *EnumTableTemplate) FileName(table *model.EnumTable) string {
	return filepath.Join("orm", table.DbKey, table.FileName()+".go")
}

func (tmpl *EnumTableTemplate) GenerateEnum(table *model.EnumTable, _w io.Writer, importPath string) (err error) {
	tmpl.Package = table.DbKey
	return tmpl.gen(table, _w)
}

func (tmpl *EnumTableTemplate) Overwrite() bool {
	return true
}

//*** enumTable.tmpl

func (tmpl *EnumTableTemplate) gen(table *model.EnumTable, _w io.Writer) (err error) {
	if err = tmpl.genHeader(table, _w); err != nil {
		return
	}
	if err = tmpl.genConstants(table, _w); err != nil {
		return
	}
	if err = tmpl.genString(table, _w); err != nil {
		return
	}
	if err = tmpl.genId(table, _w); err != nil {
		return
	}
	if err = tmpl.genPlurals(table, _w); err != nil {
		return
	}
	if err = tmpl.genInterfaces(table, _w); err != nil {
		return
	}
	if err = tmpl.genFields(table, _w); err != nil {
		return
	}
	return
}

func (tmpl *EnumTableTemplate) genHeader(table *model.EnumTable, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// Code generated by goradd-orm. DO NOT EDIT.

package `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, tmpl.Package); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `

import (
	"strconv"
)
`); err != nil {
		return
	}

	return
}

func (tmpl *EnumTableTemplate) genConstants(table *model.EnumTable, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `type `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` int

const (
`); err != nil {
		return
	}

	maxKey := 0
	for _, con := range table.Constants {
		if con.Value > maxKey {
			maxKey = con.Value
		}

		if _, err = io.WriteString(_w, `	`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, con.Const); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` = `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, strconv.Itoa(con.Value)); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `)

// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `MaxValue is the maximum enumerated value of `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `
// doc: type=`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `
const `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `MaxValue = `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, strconv.Itoa(maxKey)); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `

`); err != nil {
		return
	}

	return
}

func (tmpl *EnumTableTemplate) genString(table *model.EnumTable, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// String returns the name value of the type and satisfies the fmt.Stringer interface
func (e `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `) String() string {
	switch e {
	case 0: return ""
`); err != nil {
		return
	}

	for _, con := range table.Constants {

		if _, err = io.WriteString(_w, `	case `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, con.Const); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `: return `+"`"+``); err != nil {
			return
		}

		if _, err = io.WriteString(_w, fmt.Sprint(table.FieldValue(con.Value, 1))); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ``+"`"+`
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `	default: panic("index out of range")
	}
	return "" // prevent warning
}
`); err != nil {
		return
	}

	return
}

func (tmpl *EnumTableTemplate) genId(table *model.EnumTable, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// ID returns a string representation of the id and satisfies the IDer interface
func (e `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `) ID() string {
	return strconv.Itoa(int(e))
}

// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `FromID converts a `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` ID to a `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `
func `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `FromID (id string) `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` {
	switch id {
`); err != nil {
		return
	}

	for _, con := range table.Constants {

		if _, err = io.WriteString(_w, `    case `+"`"+``); err != nil {
			return
		}

		if _, err = io.WriteString(_w, strconv.Itoa(con.Value)); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ``+"`"+`: return `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, con.Const); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `	}
	return `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(0)
}

// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `FromIDs converts a slice of `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` IDs to a slice of `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `
func `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `FromIDs (ids []string) (values []`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `) {
    values = make([]`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `, 0, len(ids))
    for _,id := range ids {
        values = append(values, `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `FromID(id))
    }
    return
}

// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `FromName converts a `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` name to a `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `
func `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `FromName (name string) `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` {
	switch name {
`); err != nil {
		return
	}

	for _, con := range table.Constants {

		if _, err = io.WriteString(_w, `	case `+"`"+``); err != nil {
			return
		}

		if _, err = io.WriteString(_w, fmt.Sprint(table.FieldValue(con.Value, 1))); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ``+"`"+`: return `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, con.Const); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `	}
	return `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(0)
}

`); err != nil {
		return
	}

	return
}

func (tmpl *EnumTableTemplate) genPlurals(table *model.EnumTable, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` returns a slice of all the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` values
// in key order.
func `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `() ([]`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `) {
    return []`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `{ 
`); err != nil {
		return
	}

	for _, con := range table.Constants {

		if _, err = io.WriteString(_w, `        `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, con.Const); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `,
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `    }
}

// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `I returns a slice of all the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` values as generic interfaces.
// doc: type=`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `
func `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `I() (values []any) {
    return []any{ 
`); err != nil {
		return
	}

	for _, con := range table.Constants {

		if _, err = io.WriteString(_w, `        `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, con.Const); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `,
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `    }
}

`); err != nil {
		return
	}

	return
}

func (tmpl *EnumTableTemplate) genInterfaces(table *model.EnumTable, _w io.Writer) (err error) {

	if _, err = io.WriteString(_w, `// Label returns the string that will be displayed to a user for this item.
func (e `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `) Label() string {
	return e.String()
}

// Value returns the value as an interface. It satisfies goradd's Valuer interface.
func (e `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `) Value() any {
	return e.ID()
}

`); err != nil {
		return
	}

	return
}

func (tmpl *EnumTableTemplate) genFields(table *model.EnumTable, _w io.Writer) (err error) {

	for i, field := range table.Fields[2:] {
		fieldNum := i + 2

		if _, err = io.WriteString(_w, `func (e `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `) `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, field.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `() `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, field.GoType()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` {
	switch e {
	case 0: return `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, field.Type.DefaultValueString()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

		for _, con := range table.Constants {

			if _, err = io.WriteString(_w, `	case `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, con.Const); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `: return `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, fmt.Sprintf("%#v", table.FieldValue(con.Value, fieldNum))); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `	default: panic("Index out of range")
	}
	return `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, field.Type.DefaultValueString()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` // prevent warning
}

// `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, field.IdentifierPlural); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` returns a slice of all the `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, field.TitlePlural); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` associated with `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` values.
// doc: type=`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `
func `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, field.IdentifierPlural); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `() []`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, field.GoType()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` {
	return []`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, field.GoType()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` {
	    // 0 item will be a zero value
	    `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, field.Type.DefaultValueString()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `,
`); err != nil {
			return
		}

		for _, con := range table.Constants {

			if _, err = io.WriteString(_w, `	    `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, fmt.Sprintf("%#v", table.FieldValue(con.Value, fieldNum))); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `,
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `    }
}

`); err != nil {
			return
		}

	}
	return
}
