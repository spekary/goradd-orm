//** This file was code generated by GoT. DO NOT EDIT. ***

package template

import (
	"io"
	"path/filepath"
	"strconv"

	"github.com/goradd/orm/pkg/codegen"
	"github.com/goradd/orm/pkg/model"
	"github.com/goradd/orm/pkg/query"
)

func init() {
	t := TableBaseTestTemplate{}
	codegen.RegisterTemplate(&t)
}

// TableBaseTestTemplate generates code for the base table tests.
type TableBaseTestTemplate struct {
	Package string
}

func (tmpl *TableBaseTestTemplate) FileName(table *model.Table) string {
	return filepath.Join("orm", table.DbKey, table.FileName()+"_base_test.go")
}

func (tmpl *TableBaseTestTemplate) GenerateTable(table *model.Table, _w io.Writer, importPath string) (err error) {
	tmpl.Package = table.DbKey

	//*** table_base_test.tmpl

	if _, err = io.WriteString(_w, `// Code generated by goradd-orm. DO NOT EDIT.

package `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, tmpl.Package); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `

import (
	"github.com/stretchr/testify/assert"
	strings2 "github.com/goradd/strings"
	"github.com/goradd/orm/pkg/test"
	"testing"
)
`); err != nil {
		return
	}

	//*** set.tmpl

	for _, col := range table.Columns {
		if col.IsAutoId {
			continue
		}

		if _, err = io.WriteString(_w, `func Test`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `_Set`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(t *testing.T) {

    obj := New`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `()

    `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.VariableIdentifier()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` := 	test.RandomValue[`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.GoType()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `](`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, strconv.FormatUint(uint64(col.Size), 10)); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `)
    obj.Set`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.VariableIdentifier()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `)
    assert.Equal(t, `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.VariableIdentifier()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `, obj.`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `())
`); err != nil {
			return
		}

		if col.IsNullable {

			if _, err = io.WriteString(_w, `    assert.False(t, obj.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `IsNull())
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

		if col.IsNullable {

			if _, err = io.WriteString(_w, `    // Test nil
    obj.Set`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(nil)
    assert.Equal(t, `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.DefaultValueAsValue()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `, obj.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(), "set nil")
    assert.True(t, obj.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `IsNull())
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `
    // test zero
    obj.Set`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.DefaultValueAsValue()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `)
    assert.Equal(t, `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.DefaultValueAsValue()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `, obj.`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(), "set empty")
`); err != nil {
			return
		}

		if col.IsNullable {

			if _, err = io.WriteString(_w, `    assert.False(t, obj.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `IsNull())
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

		if col.Size > 0 && (col.Type == query.ColTypeBytes || col.Type == query.ColTypeString) {

			if _, err = io.WriteString(_w, `    // test panic on setting value larger than maximum size allowed
    `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.VariableIdentifier()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, ` = test.RandomValue[`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.GoType()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `](`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, strconv.FormatUint(uint64(col.Size+1), 10)); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `)
    assert.Panics(t, func() {
        obj.Set`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.VariableIdentifier()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `)
    })
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `}
`); err != nil {
			return
		}

	}
	return
}

func (tmpl *TableBaseTestTemplate) Overwrite() bool {
	return true
}
