//** This file was code generated by GoT. DO NOT EDIT. ***

package template

// This template generates a got template for the db.go file in the orm directory

import (
	"fmt"
	"io"
	"path/filepath"
	"slices"

	"github.com/goradd/maps"
	"github.com/goradd/orm/pkg/codegen"
	"github.com/goradd/orm/pkg/model"
)

func init() {
	t := DbTemplate{}
	codegen.RegisterTemplate(&t)
}

// DbTemplate generates the db.go file
type DbTemplate struct {
	Package string
}

func (tmpl *DbTemplate) FileName(dbKey string) string {
	return filepath.Join("orm", dbKey, "db.go")
}

func (tmpl *DbTemplate) GenerateDatabase(database *model.Database, _w io.Writer) (err error) {
	tmpl.Package = database.Key

	//*** db.tmpl

	if _, err = io.WriteString(_w, `// Code generated by goradd-orm. DO NOT EDIT.

// Package `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, tmpl.Package); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` contains the object relational model for the `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, database.Key); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` database.
//
// Queries use a builder pattern, started with a Query* function. Add functions to the builder to further constrain the query,
// using nodes from the [node] package to refer to tables and columns in the database. End the query with either a Load call to get a
// list of items, a Get call to get one item, or a Count call to count the number of items in the query.
//
// Some Examples
//
//   projects := model.QueryProjects().Load()
//
// Returns all the projects in the database.
//
//   projects := model.QueryProjects().
//       Join(node.Project().Manager()).
//       Where(op.GreaterOrEqual(node.Project().StartDate(), time.NewDate(2006, 1, 1)).
//       OrderBy(node.Project().Num()).
//       Load()
//
// Returns the projects that started in 2006 or later, with the manager objects attached, and ordered by project number.
// To get the manager of the first project returned, you can do this:
//
//   firstManager := projects[0].Manager()
//
// See the goradd-orm documentation for more information.
//
package `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, tmpl.Package); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `

import (
	"github.com/goradd/orm/pkg/db"
	"context"
	"io"
	"encoding/json"
)

// The newObjectPkCounter is used to assign a temporary primary key to an auto generated primary key that
// has not been saved. Auto generated primary keys are generated by the database driver, and until the record
// is saved, it does not have the final value. But having unique values for the primary key of new values helps
// with association tables and in user interfaces with associating rows of a table for example, with new objects
// that are yet to be saved. This does not need to be a thread safe object since we just need to guarantee uniqueness
// within a thread. Duplicates or skipped values between threads are OK.
var newObjectPkCounter int64


// Database returns the database object corresponding to `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, database.Key); err != nil {
		return
	}

	if _, err = io.WriteString(_w, ` in the global database cluster.
func Database() db.DatabaseI {
    return db.GetDatabase("`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, database.Key); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `")
}

`); err != nil {
		return
	}

	//*** clear_all.tmpl

	if _, err = io.WriteString(_w, `
`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `
// ClearAll deletes all the data in the database, except for data in Enum tables.
func ClearAll(ctx context.Context) {
    db := Database()

`); err != nil {
		return
	}

	for _, mm := range database.UniqueManyManyReferences() {

		if _, err = io.WriteString(_w, `    db.Delete(ctx, `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, fmt.Sprintf("%#v", mm.AssnTableName)); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `, nil)
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `
`); err != nil {
		return
	}

	for _, table := range slices.Backward(database.MarshalOrder()) {

		if _, err = io.WriteString(_w, `    db.Delete(ctx, `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, fmt.Sprintf("%#v", table.QueryName)); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `, nil)
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `
}

`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `
`); err != nil {
		return
	}

	//*** json_encode.tmpl

	{
		orderedTables := database.MarshalOrder()
		var processedTables maps.Set[string]
		assnTables := database.UniqueManyManyReferences()

		if _, err = io.WriteString(_w, `

`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `
func JsonEncodeAll(ctx context.Context, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	if _,err := io.WriteString(writer, "[\n"); err != nil {
		return err
	}

`); err != nil {
			return
		}

		for tableNum, table := range orderedTables {

			processedTables.Add(table.QueryName)

			if _, err = io.WriteString(_w, `	{	// Write `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `
		if _,err := io.WriteString(writer, "["); err != nil {
			return err
		}

		if _,err := io.WriteString(writer, `+"`"+`"`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.QueryName); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `"`+"`"+`); err != nil {
			return err
		}
        if _,err := io.WriteString(writer, ",\n["); err != nil {
            return err
        }

		cursor := Query`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(ctx).LoadCursor()
		defer cursor.Close()
		if obj := cursor.Next(); obj != nil {
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		for obj := cursor.Next(); obj != nil; obj = cursor.Next() {
			if _, err := io.WriteString(writer, ",\n"); err != nil {
				return err
			}
			if err := encoder.Encode(obj); err != nil {
				return err
			}
		}

		if _,err := io.WriteString(writer, "]\n]"); err != nil {
			return err
		}

`); err != nil {
				return
			}

			if tableNum < len(orderedTables)-1 {

				if _, err = io.WriteString(_w, `        if _,err := io.WriteString(writer, ","); err != nil {
            return err
        }
`); err != nil {
					return
				}

			}

			if _, err = io.WriteString(_w, `		if _,err := io.WriteString(writer, "\n"); err != nil {
			return err
		}
	}
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, ` `); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `

`); err != nil {
			return
		}

		if len(assnTables) > 0 {

			if _, err = io.WriteString(_w, `    db := Database()
`); err != nil {
				return
			}

			for _, mm := range assnTables {

				if _, err = io.WriteString(_w, `    {
        if _,err := io.WriteString(writer, ",\n["); err != nil {
            return err
        }
        if _,err := io.WriteString(writer, `+"`"+``); err != nil {
					return
				}

				if _, err = io.WriteString(_w, fmt.Sprintf("%#v", mm.AssnTableName)); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `,[`+"`"+`); err != nil {
            return err
        }


        cursor := db.Query(ctx, `); err != nil {
					return
				}

				if _, err = io.WriteString(_w, fmt.Sprintf("%#v", mm.AssnTableName)); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `,
            map[string]query.ReceiverType{
                `); err != nil {
					return
				}

				if _, err = io.WriteString(_w, fmt.Sprintf("%#v", mm.AssnSourceColumnName)); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `: query.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, mm.AssnSourceColumnType.String()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `,
                `); err != nil {
					return
				}

				if _, err = io.WriteString(_w, fmt.Sprintf("%#v", mm.AssnDestColumnName)); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `: query.`); err != nil {
					return
				}

				if _, err = io.WriteString(_w, mm.AssnDestColumnType.String()); err != nil {
					return
				}

				if _, err = io.WriteString(_w, `,
            },
            nil,
            nil)
        if rec := cursor.Next(); rec != nil {
            if err := encoder.Encode(rec); err != nil {
                return err
            }
        }
        for rec := cursor.Next(); rec != nil; rec = cursor.Next() {
            if _,err := io.WriteString(writer, ",\n"); err != nil {
                return err
            }
            if err := encoder.Encode(rec); err != nil {
                return err
            }
        }
        if _,err := io.WriteString(writer, `+"`"+`]]`+"`"+`); err != nil {
            return err
        }
    }
`); err != nil {
					return
				}

			}

		}

		if _, err = io.WriteString(_w, `
	_, err := io.WriteString(writer, "]")
	return err
}

`); err != nil {
			return
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

	//*** json_decode.tmpl

	{
		orderedTables := database.MarshalOrder()
		assnTables := database.UniqueManyManyReferences()

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `
func JsonDecodeAll(ctx context.Context,  reader io.Reader) error {
	decoder := json.NewDecoder(reader)

	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the JSON to start with an array")
		return err
	}

	for decoder.More() {
		if err := jsonDecodeTable(ctx, decoder); err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}

func jsonDecodeTable(ctx context.Context,  decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the duple to start with an array")
		return err
	}

	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading name of table:", err)
		return err
	}

	if tableName, ok := token.(string); !ok {
		fmt.Println("Error: Expected a name of a table.")
		return err
	} else {
		switch tableName {
`); err != nil {
			return
		}

		for _, table := range orderedTables {

			if _, err = io.WriteString(_w, `		case `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, fmt.Sprintf("%#v", table.QueryName)); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `:
			err = jsonDecode`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(ctx, decoder)
`); err != nil {
				return
			}

		}

		for _, mm := range assnTables {

			if _, err = io.WriteString(_w, `		case `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, fmt.Sprintf("%#v", mm.AssnTableName)); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `:
			err = jsonDecode`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, mm.TableIdentifier()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(ctx, decoder)
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `
		}
		if err != nil {
			return err
		}
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}

`); err != nil {
			return
		}

		for _, table := range orderedTables {

			if _, err = io.WriteString(_w, `func jsonDecode`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(ctx context.Context,  decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, ` list to start with an array")
		return err
	}

	for decoder.More() {
		obj := New`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `()
		if err = decoder.Decode(&obj); err != nil {
			return err
		}
		obj.Save(ctx)
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `
`); err != nil {
			return
		}

		for _, mm := range assnTables {

			if _, err = io.WriteString(_w, `func jsonDecode`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, mm.TableIdentifier()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(ctx context.Context,  decoder *json.Decoder) error {
	token, err := decoder.Token()
	if err != nil {
		fmt.Println("Error reading opening token:", err)
		return err
	}
	// Ensure the first token is a start of an array
	if delim, ok := token.(json.Delim); !ok || delim != '[' {
		fmt.Println("Error: Expected the `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, mm.TableIdentifier()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, ` list to start with an array")
		return err
	}

    database := Database()
	for decoder.More() {
	    var imp struct {
            Src `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, mm.AssnSourceColumnType.GoType()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, ` `+"`"+`json:"`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, mm.AssnSourceColumnName); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `"`+"`"+`
            Dest `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, mm.AssnDestColumnType.GoType()); err != nil {
				return
			}

			if _, err = io.WriteString(_w, ` `+"`"+`json:"`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, mm.AssnDestColumnName); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `"`+"`"+`
        }

		if err = decoder.Decode(&imp); err != nil {
			return err
		}
		db.Associate(ctx, database, `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, fmt.Sprintf("%#v", mm.AssnTableName)); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `, `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, fmt.Sprintf("%#v", mm.AssnSourceColumnName)); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `, imp.Src, `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, fmt.Sprintf("%#v", mm.AssnDestColumnName)); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `, imp.Dest)
	}

	// Check if the last token is the end of the array
	token, err = decoder.Token()
	if err != nil {
		fmt.Println("Error reading the last token:", err)
		return err
	}

	if delim, ok := token.(json.Delim); !ok || delim != ']' {
		fmt.Println("Error: Expected the JSON to end with a closing array token")
		return err
	}

	return nil
}
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `

`); err != nil {
			return
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

	return
}

func (tmpl *DbTemplate) Overwrite() bool {
	return true
}
