// Code generated by goradd-orm. DO NOT EDIT.

// Package goradd_unit contains the object relational model for the goradd_unit database.
//
// Queries use a builder pattern, started with a SqlQuery* function. Add functions to the builder to further constrain the query,
// using nodes from the [node] package to refer to tables and columns in the database. End the query with either a Load call to get a
// list of items, a Get call to get one item, or a Count call to count the number of items in the query.
//
// Some Examples
//
//	projects := model.QueryProjects().Load()
//
// Returns all the projects in the database.
//
//	projects := model.QueryProjects().
//	    Join(node.Project().Manager()).
//	    Where(op.GreaterOrEqual(node.Project().StartDate(), time.NewDate(2006, 1, 1)).
//	    OrderBy(node.Project().Num()).
//	    Load()
//
// Returns the projects that started in 2006 or later, with the manager objects attached, and ordered by project number.
// To get the manager of the first project returned, you can do this:
//
//	firstManager := projects[0].Manager()
//
// See the goradd-orm documentation for more information.
package goradd_unit

import (
	"github.com/goradd/orm/pkg/db"
)

// The newObjectPkCounter is used to assign a temporary primary key to an auto generated primary key that
// has not been saved. Auto generated primary keys are generated by the database driver, and until the record
// is saved, it does not have the final value. But having unique values for the primary key of new values helps
// with association tables and in user interfaces with associating rows of a table for example, with new objects
// that are yet to be saved. This does not need to be a thread safe object since we just need to guarantee uniqueness
// within a thread. Duplicates or skipped values between threads are OK.
var newObjectPkCounter int64

// Database returns the database object corresponding to goradd_unit in the global database cluster.
func Database() db.DatabaseI {
	return db.GetDatabase("goradd_unit")
}
