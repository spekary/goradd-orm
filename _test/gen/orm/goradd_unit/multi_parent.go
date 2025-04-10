package goradd_unit

// This is the implementation file for the MultiParent ORM object.
// This is where you build the api to your data model for your web application and potentially mobile apps.
// Your edits to this file will be preserved.

import (
	"context"
	"encoding/gob"
	"fmt"
)

// MultiParent represents an item in the multi_parent table in the database.
type MultiParent struct {
	multiParentBase
}

// NewMultiParent creates a new MultiParent object and initializes it to default values.
func NewMultiParent() *MultiParent {
	o := new(MultiParent)
	o.Initialize()
	return o
}

// Initialize will initialize or re-initialize a MultiParent database object to default values.
func (o *MultiParent) Initialize() {
	o.multiParentBase.Initialize()
	// Add your own initializations here
}

// String implements the Stringer interface and returns a description of the record, primarily for debugging.
func (o *MultiParent) String() string {
	if o == nil {
		return ""
	}
	return fmt.Sprintf("MultiParent %v", o.PrimaryKey())
}

// Key returns a unique key for the object, among a list of similar objects.
func (o *MultiParent) Key() string {
	if o == nil {
		return ""
	}
	return fmt.Sprintf("%v", o.PrimaryKey())
}

// Label returns a human-readable label of the object.
// This would be what a user would see as a description of the object if choosing from a list.
func (o *MultiParent) Label() string {
	if o == nil {
		return ""
	}
	return o.Name()
}

// Save will update or insert the object, depending on the state of the object.
// If it has an auto-generated primary key, it will be changed after an insert.
// Database errors generally will be handled by a panic and not returned here,
// since those indicate a problem with a database driver or configuration.
//
// Save will return a db.OptimisticLockError if it detects a collision when two users
// are attempting to change the same database record.
//
// It will return a db.UniqueValueError if it detects a collision when an attempt
// is made to add a record with a unique column that is given a value that is already in the database.
//
// Updating a record that has not changed will have no effect on the database.
func (o *MultiParent) Save(ctx context.Context) error {
	return o.save(ctx)
}

// QueryMultiParents returns a new query builder.
func QueryMultiParents(ctx context.Context) MultiParentBuilder {
	return queryMultiParents(ctx)
}

// queryMultiParents creates a new builder and is the central spot where all queries are directed.
// You can modify this function to enforce restrictions on queries, for example to make sure the user is authorized to
// access the data.
func queryMultiParents(ctx context.Context) MultiParentBuilder {
	return newMultiParentBuilder(ctx)
}

// DeleteMultiParent deletes the multi_parent record with primary key pk from the database.
// Note that you can also delete loaded MultiParent objects by calling Delete on them.
// Returns an error only if there was a problem with the database during the delete.
// If the record was not found, no error will be returned.
// doc: type=MultiParent
func DeleteMultiParent(ctx context.Context, pk string) error {
	return deleteMultiParent(ctx, pk)
}

func init() {
	gob.RegisterName("goradd_unitMultiParent", new(MultiParent))
}
