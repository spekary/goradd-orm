//** This file was code generated by GoT. DO NOT EDIT. ***

package template

// This template generates a got template for the db.go file in the orm directory

import (
	"fmt"
	"io"
	"path/filepath"
	"spekary/goradd/orm/pkg/codegen"
	"spekary/goradd/orm/pkg/model"
)

func init() {
	codegen.RegisterTemplate(new(DbTemplate))
}

type DbTemplate struct{}

func (n DbTemplate) FileName(dbKey string) string {
	return filepath.Join(dbKey, "orm", "db.go")
}

func (n DbTemplate) GenerateDatabase(database *model.Database, _w io.Writer) (err error) {
	packageName := "orm"

	if _, err = io.WriteString(_w, `// Code generated by goradd-orm. DO NOT EDIT.

// Package `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, packageName); err != nil {
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

	if _, err = io.WriteString(_w, packageName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `

import (
	"spekary/goradd/orm/pkg/db"
)

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

	if _, err = io.WriteString(_w, fmt.Sprint(database.Key)); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `")
}

`); err != nil {
		return
	}

	return
}

func (n DbTemplate) Overwrite() bool {
	return true
}
