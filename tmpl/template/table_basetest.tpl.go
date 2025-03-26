//** This file was code generated by GoT. DO NOT EDIT. ***

package template

import (
	"io"
	"path/filepath"
	"strconv"

	"github.com/goradd/orm/pkg/codegen"
	"github.com/goradd/orm/pkg/model"
	"github.com/goradd/orm/pkg/query"
	"github.com/goradd/orm/pkg/schema"
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
	"github.com/stretchr/testify/require"
	strings2 "github.com/goradd/strings"
	"github.com/goradd/orm/pkg/test"
	"testing"
)
`); err != nil {
		return
	}

	var hasRequiredUnknown bool
	for _, col := range table.Columns {
		if col.ReceiverType == query.ColTypeUnknown && !col.IsNullable {
			hasRequiredUnknown = true
		} // cannot know what the set of valid input characters are.
	}

	//*** sample.tmpl

	if _, err = io.WriteString(_w, `// createMinimalSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` creates an unsaved minimal version of a `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` object
// for testing.
func createMinimalSample`); err != nil {
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

	if _, err = io.WriteString(_w, ` {
    obj := New`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `()
    updateMinimalSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(obj)
    return obj
}

// updateMinimalSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` sets the values of a minimal sample to new, random values.
func updateMinimalSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(obj *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `)  {
`); err != nil {
		return
	}

	for _, col := range table.Columns {

		if col.ReceiverType == query.ColTypeUnknown {
			continue
		} // cannot know what the set of valid input characters are.
		testSize := min(col.Size, 100000)

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

		if col.HasSetter() {

			if col.IsReference() {

				if !col.IsNullable {

					if _, err = io.WriteString(_w, `    // A required forward reference will need to be fulfilled just to save the minimal version of this object
    // If the database is configured so that the referenced object points back here, either directly or through multiple
    // forward references, it possible this could create an endless loop. Such a structure could not be saved anyways.
    obj.Set`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.ReferenceIdentifier()); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `(createMinimalSample`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.ReferenceType()); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `())
`); err != nil {
						return
					}

				}

				if _, err = io.WriteString(_w, `
`); err != nil {
					return
				}

			} else if col.IsEnum() {

				if _, err = io.WriteString(_w, `
`); err != nil {
					return
				}

				if col.IsEnumArray() {

					if _, err = io.WriteString(_w, `     obj.Set`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Identifier); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `(test.RandomEnumArray(`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Reference.EnumTable.IdentifierPlural); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `()))
`); err != nil {
						return
					}

				} else {

					if _, err = io.WriteString(_w, `     obj.Set`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Identifier); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `(test.RandomEnum(`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Reference.EnumTable.IdentifierPlural); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `()))
`); err != nil {
						return
					}

				}

				if _, err = io.WriteString(_w, `
`); err != nil {
					return
				}

			} else {

				if _, err = io.WriteString(_w, `
`); err != nil {
					return
				}

				if col.SchemaSubType == schema.ColSubTypeNumeric {

					if _, err = io.WriteString(_w, `    obj.Set`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Identifier); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `(test.RandomDecimal(`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, strconv.FormatUint(uint64(col.DecimalPrecision()), 10)); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `, `); err != nil {
						return
					}

					if _, err = io.WriteString(_w, strconv.FormatUint(uint64(col.DecimalScale()), 10)); err != nil {
						return
					}

					if _, err = io.WriteString(_w, ` ))
`); err != nil {
						return
					}

				} else {

					if _, err = io.WriteString(_w, `    obj.Set`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Identifier); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `(test.RandomValue[`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.GoType()); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `](`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, strconv.FormatUint(uint64(testSize), 10)); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `))
`); err != nil {
						return
					}

				}

				if _, err = io.WriteString(_w, `
`); err != nil {
					return
				}

			}

			if _, err = io.WriteString(_w, `
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `}

// createMaximalSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` creates an unsaved version of a `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` object
// for testing that includes references to minimal objects.
func createMaximalSample`); err != nil {
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

	if _, err = io.WriteString(_w, ` {
    obj := New`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `()
    updateMaximalSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(obj)
    return obj
}

// updateMaximalSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` sets all the maximal sample values to new values.
func updateMaximalSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(obj *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `) {
    updateMinimalSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(obj)
`); err != nil {
		return
	}

	for _, col := range table.Columns {

		if col.HasSetter() {

			if _, err = io.WriteString(_w, `    `); err != nil {
				return
			}

			if col.IsReference() {

				if _, err = io.WriteString(_w, `
        `); err != nil {
					return
				}

				if col.IsNullable {

					if _, err = io.WriteString(_w, `
    obj.Set`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.ReferenceIdentifier()); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `(createMinimalSample`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.ReferenceType()); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `())
        `); err != nil {
						return
					}

				}

				if _, err = io.WriteString(_w, `
    `); err != nil {
					return
				}

			}

			if _, err = io.WriteString(_w, `
`); err != nil {
				return
			}

		}

	}

	if _, err = io.WriteString(_w, `
`); err != nil {
		return
	}

	for _, col := range table.ReverseReferences {

		if _, err = io.WriteString(_w, `    obj.Set`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.ReverseIdentifier()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(createMinimalSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.Table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `())
`); err != nil {
			return
		}

	}

	for _, mm := range table.ManyManyReferences {

		if _, err = io.WriteString(_w, `    obj.Set`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, mm.IdentifierPlural); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(createMinimalSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, mm.DestinationTable.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `())
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `}

// deleteSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` deletes an object created and saved by one of the sample creator functions.
func deleteSample`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(ctx context.Context, obj *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `) {
    if obj == nil {
        return
    }

`); err != nil {
		return
	}

	for _, col := range table.ReverseReferences {

		if col.IsUnique {

			if _, err = io.WriteString(_w, `    deleteSample`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.Table.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(ctx, obj.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.ReverseIdentifier()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `())
`); err != nil {
				return
			}

		} else {

			if _, err = io.WriteString(_w, `    for _,item := range obj.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.ReverseIdentifier()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `() {
        deleteSample`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.Table.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(ctx, item)
    }
`); err != nil {
				return
			}

		}

	}

	if _, err = io.WriteString(_w, `
`); err != nil {
		return
	}

	for _, mm := range table.ManyManyReferences {

		if _, err = io.WriteString(_w, `    for _,item := range obj.`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, mm.IdentifierPlural); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `() {
        deleteSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, mm.ObjectType()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx, item)
    }
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `
    obj.Delete(ctx)
`); err != nil {
		return
	}

	for _, col := range table.Columns {

		if col.HasSetter() {

			if _, err = io.WriteString(_w, `    `); err != nil {
				return
			}

			if col.IsReference() {

				if _, err = io.WriteString(_w, `
    deleteSample`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.ReferenceType()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `(ctx, obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.ReferenceIdentifier()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `())
    `); err != nil {
					return
				}

			}

			if _, err = io.WriteString(_w, `
`); err != nil {
				return
			}

		}

	}

	if _, err = io.WriteString(_w, `}


`); err != nil {
		return
	}

	if !hasRequiredUnknown {
		//*** set.tmpl

		for _, col := range table.Columns {
			if !col.HasSetter() {
				continue
			}
			testSize := min(col.Size, 100000)

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

			if col.IsEnum() {

				if _, err = io.WriteString(_w, `    `); err != nil {
					return
				}

				if col.IsEnumArray() {

					if _, err = io.WriteString(_w, `
    val := 	test.RandomEnumArray(`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Reference.EnumTable.IdentifierPlural); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `())
    obj.Set`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Identifier); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `(val)
    `); err != nil {
						return
					}

				} else {

					if _, err = io.WriteString(_w, `
    val := test.RandomEnum(`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Reference.EnumTable.IdentifierPlural); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `())
    obj.Set`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Identifier); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `(val)
   `); err != nil {
						return
					}

				}

				if _, err = io.WriteString(_w, `
`); err != nil {
					return
				}

			} else {

				if col.SchemaSubType == schema.ColSubTypeNumeric {

					if _, err = io.WriteString(_w, `    val := 	test.RandomDecimal(`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, strconv.FormatUint(uint64(col.DecimalPrecision()), 10)); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `, `); err != nil {
						return
					}

					if _, err = io.WriteString(_w, strconv.FormatUint(uint64(col.DecimalScale()), 10)); err != nil {
						return
					}

					if _, err = io.WriteString(_w, ` )
`); err != nil {
						return
					}

				} else {

					if _, err = io.WriteString(_w, `    val := 	test.RandomValue[`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.GoType()); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `](`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, strconv.FormatUint(uint64(testSize), 10)); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `)
`); err != nil {
						return
					}

				}

				if _, err = io.WriteString(_w, `    obj.Set`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `(val)
`); err != nil {
					return
				}

			}

			if col.SchemaSubType == schema.ColSubTypeDateOnly {

				if _, err = io.WriteString(_w, `    val = obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `()
    assert.Zero(t, val.Minute()) // make sure minute part is zero'd
    assert.Zero(t, val.Hour()) // make sure hour part is zero'd
    assert.Zero(t, val.Second()) // make sure second part is zero'd
`); err != nil {
					return
				}

			} else if col.SchemaSubType == schema.ColSubTypeTimeOnly {

				if _, err = io.WriteString(_w, `    val = obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `()
    assert.EqualValues(t, 1, val.Year()) // make sure year part is zero'd. The zero value of time.Time has a year of 1.
    assert.EqualValues(t, 1, val.Month()) // make sure month part is zero'd
    assert.EqualValues(t, 1, val.Day()) // make sure day part is zero'd
`); err != nil {
					return
				}

			} else {

				if _, err = io.WriteString(_w, `    assert.Equal(t, val, obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `())
`); err != nil {
					return
				}

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

				if _, err = io.WriteString(_w, `    // Test NULL
    obj.Set`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `ToNull()
`); err != nil {
					return
				}

				if col.DefaultValueAsValue() == "nil" {

					if _, err = io.WriteString(_w, `    assert.Nil(t, obj.`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Identifier); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `())
`); err != nil {
						return
					}

				} else {

					if _, err = io.WriteString(_w, `    assert.EqualValues(t, `); err != nil {
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

					if _, err = io.WriteString(_w, `())
`); err != nil {
						return
					}

				}

				if _, err = io.WriteString(_w, `    assert.True(t, obj.`); err != nil {
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
    // test default
`); err != nil {
				return
			}

			if col.IsEnumArray() {

				if _, err = io.WriteString(_w, `    obj.Set`); err != nil {
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
    assert.True(t, obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `().Equal(`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.GoType()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `(`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.DefaultValueAsValue()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `)), "set default")
`); err != nil {
					return
				}

			} else {

				if _, err = io.WriteString(_w, `    obj.Set`); err != nil {
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
    assert.EqualValues(t, `); err != nil {
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

				if _, err = io.WriteString(_w, `(), "set default")
`); err != nil {
					return
				}

			}

			if _, err = io.WriteString(_w, `
`); err != nil {
				return
			}

			if col.Size > 0 &&
				col.Size == testSize &&
				!col.IsEnum() &&
				(col.ReceiverType == query.ColTypeBytes || col.ReceiverType == query.ColTypeString || col.ReceiverType == query.ColTypeUnknown) {

				if _, err = io.WriteString(_w, `    // test panic on setting value larger than maximum size allowed
    val = test.RandomValue[`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.GoType()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `](`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, strconv.FormatUint(uint64(col.MaxLength()+1), 10)); err != nil {
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

				if _, err = io.WriteString(_w, `(val)
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

		if _, err = io.WriteString(_w, `
func Test`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `_Copy(t *testing.T) {
    obj := createMinimalSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `()

    obj2 := obj.Copy()

`); err != nil {
			return
		}

		for _, col := range table.Columns {

			if col.HasSetter() {

				if _, err = io.WriteString(_w, `    assert.Equal(t, obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `(), obj2.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `())
`); err != nil {
					return
				}

			}

		}

		if _, err = io.WriteString(_w, `
}

`); err != nil {
			return
		}

		//*** save.tmpl

		if _, err = io.WriteString(_w, `func Test`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `_BasicInsert(t *testing.T) {
    obj := createMinimalSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `()
    ctx := db.NewContext(nil)
    err := obj.Save(ctx)
	assert.NoError(t, err)
    defer deleteSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx, obj)

    // Test retrieval
    obj2 := Load`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx, obj.PrimaryKey())
    require.NotNil(t, obj2)

    assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

`); err != nil {
			return
		}

		for _, col := range table.Columns {

			if col.ReceiverType == query.ColTypeUnknown {
				continue
			} // cannot know what the set of valid input characters are.
			if col.IsReference() {
				continue
			} // forward references will be tested in the References test.

			if _, err = io.WriteString(_w, `
    assert.True(t, obj2.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, col.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `IsValid())
`); err != nil {
				return
			}

			if col.IsNullable {

				if _, err = io.WriteString(_w, `    assert.False(t, obj2.`); err != nil {
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

			if col.HasSetter() {

				if col.IsEnumArray() {

					if _, err = io.WriteString(_w, `    assert.True(t, obj.`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Identifier); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `().Equal(obj2.`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.Identifier); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `()))
`); err != nil {
						return
					}

				} else {

					if _, err = io.WriteString(_w, `
`); err != nil {
						return
					}

					if col.SchemaSubType != schema.ColSubTypeNumeric {

						if _, err = io.WriteString(_w, `    assert.EqualValues(t, obj.`); err != nil {
							return
						}

						if _, err = io.WriteString(_w, col.Identifier); err != nil {
							return
						}

						if _, err = io.WriteString(_w, `(), obj2.`); err != nil {
							return
						}

						if _, err = io.WriteString(_w, col.Identifier); err != nil {
							return
						}

						if _, err = io.WriteString(_w, `())
`); err != nil {
							return
						}

					}

					if _, err = io.WriteString(_w, `
`); err != nil {
						return
					}

				}

				if _, err = io.WriteString(_w, `    // test that setting it to the same value will not change the dirty bit
    assert.False(t, obj2.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.VariableIdentifier()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `IsDirty)
    obj2.Set`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `(obj2.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `())
    assert.False(t, obj2.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.VariableIdentifier()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `IsDirty)
`); err != nil {
					return
				}

			}

		}

		if _, err = io.WriteString(_w, `
}

func Test`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `_InsertPanics(t *testing.T) {
    obj := createMinimalSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `()
    ctx := db.NewContext(nil)

 `); err != nil {
			return
		}

		for _, col := range table.Columns {

			if _, err = io.WriteString(_w, `
 `); err != nil {
				return
			}

			if !col.IsNullable && col.HasSetter() && col.DefaultValue == nil {

				if _, err = io.WriteString(_w, `
    obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.VariableIdentifier()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `IsValid = false
    assert.Panics(t, func() {obj.Save(ctx)})
    obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.VariableIdentifier()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `IsValid = true

 `); err != nil {
					return
				}

			}

			if _, err = io.WriteString(_w, `
 `); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `
}

func Test`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `_BasicUpdate(t *testing.T) {
    obj := createMinimalSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `()
    ctx := db.NewContext(nil)
    assert.NoError(t, obj.Save(ctx))
    defer deleteSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx, obj)
    updateMinimalSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(obj)
    assert.NoError(t, obj.Save(ctx))
    obj2 := Load`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx, obj.PrimaryKey())

`); err != nil {
			return
		}

		for _, col := range table.Columns {

			if col.ReceiverType == query.ColTypeTime {

				if _, err = io.WriteString(_w, ` `); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `
    assert.WithinDuration(t, obj2.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `(), obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `(), time.Second, "`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, ` not within one second")
`); err != nil {
					return
				}

			} else if col.SchemaSubType != schema.ColSubTypeNumeric {

				if _, err = io.WriteString(_w, `    assert.Equal(t, obj2.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `(), obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `(), "`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, ` did not update")
`); err != nil {
					return
				}

			}

		}

		if _, err = io.WriteString(_w, `}



func Test`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `_References(t *testing.T) {
    obj := createMaximalSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `()
    ctx := db.NewContext(nil)
    obj.Save(ctx)
    defer deleteSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx, obj)

    // Test that referenced objects were saved and assigned ids
`); err != nil {
			return
		}

		for _, col := range table.Columns {

			if col.IsReference() {

				if _, err = io.WriteString(_w, `    assert.NotNil(t, obj.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.ReferenceIdentifier()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `())
`); err != nil {
					return
				}

				if col.Reference.Table.PrimaryKeyColumn().IsAutoId {

					if _, err = io.WriteString(_w, `    assert.NotEqual(t, '-', obj.`); err != nil {
						return
					}

					if _, err = io.WriteString(_w, col.ReferenceIdentifier()); err != nil {
						return
					}

					if _, err = io.WriteString(_w, `().PrimaryKey()[0])
`); err != nil {
						return
					}

				}

				if _, err = io.WriteString(_w, `
`); err != nil {
					return
				}

			}

		}

		if _, err = io.WriteString(_w, `

}
`); err != nil {
			return
		}

		//*** get.tmpl

		if _, err = io.WriteString(_w, `func Test`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `_EmptyPrimaryKeyGetter(t *testing.T) {
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

		if table.PrimaryKeyColumn().IsAutoId {

			if _, err = io.WriteString(_w, `    i,err := strconv.Atoi(obj.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.PrimaryKeyColumn().Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `())
    assert.NoError(t, err)
    assert.True(t, i < 0)
`); err != nil {
				return
			}

		} else {

			if _, err = io.WriteString(_w, `    assert.Zero(t, obj.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.PrimaryKeyColumn().Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `())
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `}

func Test`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `_Getters(t *testing.T) {
    obj := createMinimalSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `()

`); err != nil {
			return
		}

		if table.PrimaryKeyColumn().IsAutoId {

			if _, err = io.WriteString(_w, `    i,err := strconv.Atoi(obj.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.PrimaryKeyColumn().Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `())
    assert.NoError(t, err)
    assert.True(t, i < 0)
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `
    ctx := db.NewContext(nil)
    require.NoError(t, obj.Save(ctx))
    defer deleteSample`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx, obj)

    obj2 := Load`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx, obj.PrimaryKey(), node.`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `().PrimaryKey())
`); err != nil {
			return
		}

		if !table.PrimaryKeyColumn().IsAutoId {

			if _, err = io.WriteString(_w, `    assert.Equal(t, obj.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.PrimaryKeyColumn().Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(), obj2.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.PrimaryKeyColumn().Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `())
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

		for _, col := range table.Columns {

			if !col.IsPrimaryKey {

				if _, err = io.WriteString(_w, `    assert.Panics(t, func() { obj2.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, col.Identifier); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `() })
`); err != nil {
					return
				}

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
