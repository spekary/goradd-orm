//** This file was code generated by GoT. DO NOT EDIT. ***

package template

import (
	"io"
	"path/filepath"

	"github.com/goradd/orm/pkg/codegen"
	"github.com/goradd/orm/pkg/model"
)

func init() {
	t := TableTemplate{}
	codegen.RegisterTemplate(&t)
}

// TableTemplate generates the code for tables.
// This is one-time generated and allows the programmer to edit it to override the base table code.
type TableTemplate struct {
	Package string
}

func (tmpl *TableTemplate) FileName(table *model.Table) string {
	return filepath.Join("orm", table.DbKey, table.FileName()+".go")
}

func (tmpl *TableTemplate) GenerateTable(table *model.Table, _w io.Writer, importPath string) (err error) {
	tmpl.Package = table.DbKey
	//*** table.tmpl

	// The master template for the table classes

	if _, err = io.WriteString(_w, `package `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, tmpl.Package); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `

// This is the implementation file for the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` ORM object.
// This is where you build the api to your data model for your web application and potentially mobile apps.
// Your edits to this file will be preserved.

import (
    "fmt"
    "context"
)

// `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` represents an item in the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` table in the database.
type `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` struct {
	`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DecapIdentifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Base
}

// New`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` creates a new `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` object and initializes it to default values.
func New`); err != nil {
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
	o := new(`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `)
	o.Initialize()
	return o
}

// Initialize will initialize or re-initialize a `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` database object to default values.
func (o *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `) Initialize() {
	o.`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DecapIdentifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Base.Initialize()
    // Add your own initializations here
}

// String implements the Stringer interface and returns the default label for the object as it appears in html lists.
// Typically you would change this to whatever was pertinent to your application.
func (o *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `) String() string  {
    if o == nil {
        return ""   // Possibly - Select One -?
    }
`); err != nil {
		return
	}

	if col := table.ColumnByName("name"); col != nil {

		if _, err = io.WriteString(_w, `    return o.`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, col.VariableIdentifier()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

	} else {

		if _, err = io.WriteString(_w, `    return fmt.Sprintf("`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` %v", o.PrimaryKey())
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `}

// SqlQuery`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` returns a new query builder.
func SqlQuery`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(ctx context.Context) *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Builder {
	return query`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(ctx)
}

// query`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` creates a new builder and is the central spot where all queries are directed.
// You can modify this function to enforce restrictions on queries, for example to make sure the user is authorized to
// access the data.
func query`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `(ctx context.Context) *`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Builder {
	return new`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Builder(ctx)
}


`); err != nil {
		return
	}

	if table.PrimaryKeyColumn() != nil {

		if _, err = io.WriteString(_w, `// Delete`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` deletes a `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.QueryName); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` record from the database given its primary key.
// Note that you can also delete loaded `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` objects by calling Delete on them.
// doc: type=`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `
func Delete`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx context.Context, pk `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.PrimaryKeyColumn().GoType()); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `) {
	delete`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx, pk)
}
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `

func init() {
    gob.RegisterName("`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DbKey); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `", new(`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `))
}

`); err != nil {
		return
	}

	return
}

func (tmpl *TableTemplate) Overwrite() bool {
	return false
}
