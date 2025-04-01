package goradd_unit

// This is the implementation file for the LeafLock ORM object.
// This is where you build the api to your data model for your web application and potentially mobile apps.
// Your edits to this file will be preserved.

import (
	"context"
	"encoding/gob"
	"fmt"
)

// LeafLock represents an item in the leaf_lock table in the database.
type LeafLock struct {
	leafLockBase
}

// NewLeafLock creates a new LeafLock object and initializes it to default values.
func NewLeafLock() *LeafLock {
	o := new(LeafLock)
	o.Initialize()
	return o
}

// Initialize will initialize or re-initialize a LeafLock database object to default values.
func (o *LeafLock) Initialize() {
	o.leafLockBase.Initialize()
	// Add your own initializations here
}

// String implements the Stringer interface and returns a description of the record, primarily for debugging.
func (o *LeafLock) String() string {
	if o == nil {
		return ""
	}
	return fmt.Sprintf("LeafLock %v", o.PrimaryKey())
}

// Key returns a unique key for the object, among a list of similar objects.
func (o *LeafLock) Key() string {
	if o == nil {
		return ""
	}
	return fmt.Sprintf("%v", o.PrimaryKey())
}

// Label returns a human-readable label of the object.
// This would be what a user would see as a description of the object if choosing from a list.
func (o *LeafLock) Label() string {
	if o == nil {
		return ""
	}
	return o.Name()
}

// QueryLeafLocks returns a new query builder.
func QueryLeafLocks(ctx context.Context) LeafLockBuilder {
	return queryLeafLocks(ctx)
}

// queryLeafLocks creates a new builder and is the central spot where all queries are directed.
// You can modify this function to enforce restrictions on queries, for example to make sure the user is authorized to
// access the data.
func queryLeafLocks(ctx context.Context) LeafLockBuilder {
	return newLeafLockBuilder(ctx)
}

// DeleteLeafLock deletes the leaf_lock record wtih primary key pk from the database.
// Note that you can also delete loaded LeafLock objects by calling Delete on them.
// doc: type=LeafLock
func DeleteLeafLock(ctx context.Context, pk string) {
	deleteLeafLock(ctx, pk)
}

func init() {
	gob.RegisterName("goradd_unitLeafLock", new(LeafLock))
}
