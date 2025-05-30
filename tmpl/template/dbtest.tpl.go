//** This file was code generated by GoT. DO NOT EDIT. ***

package template

// This template generates a got template for the db.go file in the orm directory

import (
	"io"

	"github.com/goradd/orm/pkg/codegen"
	"github.com/goradd/orm/pkg/model"
)

func init() {
	t := DbTestTemplate{}
	codegen.RegisterTemplate(&t)
}

// DbTestTemplate generates the db_test.go file
type DbTestTemplate struct {
	Package string
}

func (tmpl *DbTestTemplate) FileName(dbKey string) string {
	return "db_test.go"
}

func (tmpl *DbTestTemplate) GenerateDatabase(database *model.Database, _w io.Writer, importPath string) (err error) {
	tmpl.Package = database.Key

	//*** dbtest.tmpl

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
    "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, importPath); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `/node"
	"testing"
)
`); err != nil {
		return
	}

	//*** test_main.tmpl

	if _, err = io.WriteString(_w, `
func TestMain(m *testing.M) {
	os.Exit(runTests(m))
}

func runTests(m *testing.M) int {
	setup(m)
	defer teardown()
	return m.Run()
}

func setup(m *testing.M) {
	fmt.Println("Setting up tests...")

	//fmt.Println(test.NewSeed())

	// uncomment the next line to re-use a previously created seed to recreate a test
	test.UseSeed(1000)

	InitDB()
}

func teardown() {
	// Cleanup logic here
	fmt.Println("Cleaning up after tests...")
}


`); err != nil {
		return
	}

	//*** test_json.tmpl

	orderedTables := database.MarshalOrder()
	//assnTables := database.UniqueManyManyReferences()

	if _, err = io.WriteString(_w, `// TestDbJson will export the entire database as JSON into a memory buffer, clear the database, then
// import the entire database from the buffer. It will then do some sanity checks.
func TestDbJson(t *testing.T) {
return
    ctx := db.NewContext(nil)

    // get single comparison objects and data sizes
    // database must be pre-populated for test

`); err != nil {
		return
	}

	for _, table := range orderedTables {

		if _, err = io.WriteString(_w, `    v_`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `, _ := Query`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx).OrderBy(node.`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `().PrimaryKey()).Get() // gets first record
`); err != nil {
			return
		}

	}

	for _, table := range orderedTables {

		if _, err = io.WriteString(_w, `    v_`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `Count, _ := Count`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx)
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `
    var b bytes.Buffer
    w := bufio.NewWriter(&b)
    assert.NoError(t, JsonEncodeAll(ctx, w))

    ClearAll(ctx)
`); err != nil {
		return
	}

	for _, table := range orderedTables {

		if _, err = io.WriteString(_w, `    assert.Equal(t, 0, func() int {i,_ := Count`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx); return i}() )
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `
    r := bufio.NewReader(&b)
    assert.NoError(t, JsonDecodeAll(ctx, r))

`); err != nil {
		return
	}

	for _, table := range orderedTables {

		if _, err = io.WriteString(_w, `    if v_`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, ` != nil {
`); err != nil {
			return
		}

		if !table.NoTest {

			if _, err = io.WriteString(_w, `        obj, _ := Query`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(ctx).OrderBy(node.`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `().PrimaryKey()).Get()
        assertEqualFields`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `(t, v_`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `, obj)
`); err != nil {
				return
			}

		}

		if _, err = io.WriteString(_w, `    }
`); err != nil {
			return
		}

	}

	for _, table := range orderedTables {

		if _, err = io.WriteString(_w, `    assert.Equal(t, v_`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.Identifier); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `Count, func() int {i,_ := Count`); err != nil {
			return
		}

		if _, err = io.WriteString(_w, table.IdentifierPlural); err != nil {
			return
		}

		if _, err = io.WriteString(_w, `(ctx); return i}() )
`); err != nil {
			return
		}

	}

	if _, err = io.WriteString(_w, `}

`); err != nil {
		return
	}

	return
}

func (tmpl *DbTestTemplate) Overwrite() bool {
	return true
}
