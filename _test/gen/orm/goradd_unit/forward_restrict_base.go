// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"unicode/utf8"

	"github.com/goradd/all"
	"github.com/goradd/orm/_test/gen/orm/goradd_unit/node"
	"github.com/goradd/orm/pkg/broadcast"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/query"
)

// ForwardRestrictBase is embedded in a ForwardRestrict object and provides the ORM access to the database.
// The member variables of the structure are private and should not normally be accessed by the ForwardRestrict embedder.
// Instead, use the accessor functions.
type forwardRestrictBase struct {
	id        string
	idIsValid bool
	idIsDirty bool

	name        string
	nameIsValid bool
	nameIsDirty bool

	reverseID        string
	reverseIDIsValid bool
	reverseIDIsDirty bool
	objReverse       *Reverse

	// Custom aliases, if specified
	_aliases map[string]any

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update.
	_restored bool

	_originalPK string
}

// IDs used to access the ForwardRestrict object fields by name using the Get function.
// doc: type=ForwardRestrict
const (
	ForwardRestrict_ID        = `ID`
	ForwardRestrict_Name      = `Name`
	ForwardRestrict_ReverseID = `ReverseID`
	ForwardRestrict_Reverse   = `Reverse`
)

const ForwardRestrictNameMaxLength = 100 // The number of runes the column can hold

// Initialize or re-initialize a ForwardRestrict database object to default values.
// The primary key will get a temporary negative number which will be replaced when the object is saved.
// Multiple calls to Initialize are not guaranteed to create sequential values for the primary key.
func (o *forwardRestrictBase) Initialize() {

	newObjectPkCounter = newObjectPkCounter - 1
	o.id = fmt.Sprintf("%d", newObjectPkCounter)

	o.idIsValid = false
	o.idIsDirty = false

	o.name = ""

	o.nameIsValid = false
	o.nameIsDirty = false

	o.reverseID = ""

	o.reverseIDIsValid = false
	o.reverseIDIsDirty = false

	o._restored = false
}

// PrimaryKey returns the current value of the primary key field.
func (o *forwardRestrictBase) PrimaryKey() string {
	return o.id
}

// OriginalPrimaryKey returns the value of the primary key that was originally loaded into the object when it was
// read from the database.
func (o *forwardRestrictBase) OriginalPrimaryKey() string {
	return o._originalPK
}

// Copy copies all valid fields to a new ForwardRestrict object.
// Forward reference ids will be copied, but reverse and many-many references will not.
// Attached objects will not be included in the copy.
// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *forwardRestrictBase) Copy() (newObject *ForwardRestrict) {
	newObject = NewForwardRestrict()
	if o.nameIsValid {
		newObject.SetName(o.name)
	}
	if o.reverseIDIsValid {
		newObject.SetReverseID(o.reverseID)
	}
	return
}

// ID returns the loaded value of ID or
// the zero value if not loaded. Call IDIsValid() to determine
// if it is loaded.
func (o *forwardRestrictBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *forwardRestrictBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

// Name returns the loaded value of Name.
func (o *forwardRestrictBase) Name() string {
	if o._restored && !o.nameIsValid {
		panic("Name was not selected in the last query and has not been set, and so is not valid")
	}
	return o.name
}

// NameIsValid returns true if the value was loaded from the database or has been set.
func (o *forwardRestrictBase) NameIsValid() bool {
	return o.nameIsValid
}

// SetName sets the value of Name in the object, to be saved later using the Save() function.
func (o *forwardRestrictBase) SetName(name string) {
	o.nameIsValid = true
	if utf8.RuneCountInString(name) > ForwardRestrictNameMaxLength {
		panic("attempted to set ForwardRestrict.Name to a value larger than its maximum length in runes")
	}
	if o.name != name || !o._restored {
		o.name = name
		o.nameIsDirty = true
	}

}

// ReverseID returns the loaded value of ReverseID.
func (o *forwardRestrictBase) ReverseID() string {
	if o._restored && !o.reverseIDIsValid {
		panic("ReverseID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.reverseID
}

// ReverseIDIsValid returns true if the value was loaded from the database or has been set.
func (o *forwardRestrictBase) ReverseIDIsValid() bool {
	return o.reverseIDIsValid
}

// SetReverseID sets the value of ReverseID in the object, to be saved later using the Save() function.
func (o *forwardRestrictBase) SetReverseID(reverseID string) {
	o.reverseIDIsValid = true
	if o.reverseID != reverseID || !o._restored {
		o.reverseID = reverseID
		o.reverseIDIsDirty = true
		o.objReverse = nil
	}

}

// Reverse returns the current value of the loaded Reverse, and nil if its not loaded.
func (o *forwardRestrictBase) Reverse() *Reverse {
	return o.objReverse
}

// LoadReverse returns the related Reverse. If it is not already loaded,
// it will attempt to load it, provided the ReverseID column has been loaded first.
func (o *forwardRestrictBase) LoadReverse(ctx context.Context) *Reverse {
	if !o.reverseIDIsValid {
		return nil
	}

	if o.objReverse == nil {
		// Load and cache
		o.objReverse = LoadReverse(ctx, o.reverseID)
	}
	return o.objReverse
}

// SetReverse sets the value of Reverse in the object, to be saved later using the Save() function.
func (o *forwardRestrictBase) SetReverse(objReverse *Reverse) {
	if objReverse == nil {
		panic("Cannot set Reverse to a null value.")
	} else {
		o.objReverse = objReverse
		o.reverseIDIsValid = true
		if o.reverseID != objReverse.PrimaryKey() {
			o.reverseID = objReverse.PrimaryKey()
			o.reverseIDIsDirty = true
		}
	}
}

// GetAlias returns the alias for the given key.
func (o *forwardRestrictBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *forwardRestrictBase) IsNew() bool {
	return !o._restored
}

// LoadForwardRestrict returns a ForwardRestrict from the database.
// selectNodes lets you provide nodes for selecting specific fields or additional fields from related tables.
// See [ForwardRestrictsBuilder.Select] for more info.
func LoadForwardRestrict(ctx context.Context, id string, selectNodes ...query.Node) *ForwardRestrict {
	return queryForwardRestricts(ctx).
		Where(op.Equal(node.ForwardRestrict().ID(), id)).
		Select(selectNodes...).
		Get()
}

// HasForwardRestrict returns true if a ForwardRestrict with the given primaryKey exists in the database.
// doc: type=ForwardRestrict
func HasForwardRestrict(ctx context.Context, id string) bool {
	return queryForwardRestricts(ctx).
		Where(op.Equal(node.ForwardRestrict().ID(), id)).
		Count(false) == 1
}

// The ForwardRestrictBuilder uses the query.BuilderI interface to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, LoadCursor, Get, Count, or Delete
type ForwardRestrictBuilder interface {
	// Join(alias string, joinedTable query.Node, condition query.Node) ForwardRestrictBuilder

	// Expand turns a Reverse or ManyMany node into individual rows.
	Expand(n query.Expander) ForwardRestrictBuilder

	// Where adds a condition to filter what gets selected.
	// Calling Where multiple times will AND the conditions together.
	Where(c query.Node) ForwardRestrictBuilder

	// OrderBy specifies how the resulting data should be sorted.
	// By default, the given nodes are sorted in ascending order.
	// Add Descending() to the node to specify that it should be sorted in descending order.
	OrderBy(nodes ...query.Sorter) ForwardRestrictBuilder

	// Limit will return a subset of the data, limited to the offset and number of rows specified.
	// For large data sets and specific types of queries, this can be slow, because it will perform
	// the entire query before computing the limit.
	// You cannot limit a query that has embedded arrays.
	Limit(maxRowCount int, offset int) ForwardRestrictBuilder

	// Select optimizes the query to only return the specified fields.
	// Once you put a Select in your query, you must specify all the fields that you will eventually read out.
	// Some fields, like primary keys, are always selected.
	// If you are using a GroupBy, most database drivers will only allow selecting on fields in the GroupBy, and
	// doing otherwise will result in an error.
	Select(nodes ...query.Node) ForwardRestrictBuilder

	// Calculation adds a calculation node with an aliased name.
	// After the query, you can read the data using GetAlias() on a returned object.
	Calculation(name string, n query.Aliaser) ForwardRestrictBuilder

	// Distinct removes duplicates from the results of the query.
	// Adding a Select() is usually required.
	Distinct() ForwardRestrictBuilder

	// GroupBy controls how results are grouped when using aggregate functions with Calculation.
	GroupBy(nodes ...query.Node) ForwardRestrictBuilder

	// Having does additional filtering on the results of the query after the query is performed.
	Having(node query.Node) ForwardRestrictBuilder

	// Load terminates the query builder, performs the query, and returns a slice of ForwardRestrict objects.
	// If there are any errors, nil is returned and the specific error is stored in the context.
	// If no results come back from the query, it will return a non-nil empty slice.
	Load() []*ForwardRestrict
	// Load terminates the query builder, performs the query, and returns a slice of interfaces.
	// This can then satisfy a general interface that loads arrays of objects.
	// If there are any errors, nil is returned and the specific error is stored in the context.
	// If no results come back from the query, it will return a non-nil empty slice.
	LoadI() []any

	// LoadCursor terminates the query builder, performs the query, and returns a cursor to the query.
	//
	// A query cursor is useful for dealing with large amounts of query results. However, there are some
	// limitations to its use. When working with SQL databases, you cannot use a cursor while querying
	// many-to-many or reverse relationships that will create an array of values.
	//
	// Call Next() on the returned cursor object to step through the results. Make sure you call Close
	// on the cursor object when you are done. You should use
	//   defer cursor.Close()
	// to make sure the cursor gets closed.
	LoadCursor() forwardRestrictsCursor

	// Get is a convenience method to return only the first item found in a query.
	// The entire query is performed, so you should generally use this only if you know
	// you are selecting on one or very few items.
	//
	// If an error occurs, or no results are found, a nil is returned.
	// In the case of an error, the error is returned in the context.
	Get() *ForwardRestrict

	// Count terminates a query and returns just the number of items selected.
	// distinct wll count the number of distinct items, ignoring duplicates.
	// nodes will select individual fields, and should be accompanied by a GroupBy.
	Count(distinct bool, nodes ...query.Node) int

	// Delete uses the query builder to delete a group of records that match the criteria
	Delete()

	// Subquery terminates the query builder and tags it as a subquery within a larger query.
	// You MUST include what you are selecting by adding Calculation or Select functions on the subquery builder.
	// Generally you would use this as a node to a Calculation function on the surrounding query builder.
	// Subquery() *query.SubqueryNode

}

type forwardRestrictQueryBuilder struct {
	builder *query.Builder
}

func newForwardRestrictBuilder(ctx context.Context) ForwardRestrictBuilder {
	b := forwardRestrictQueryBuilder{
		builder: query.NewBuilder(ctx, node.ForwardRestrict()),
	}
	return &b
}

// Load terminates the query builder, performs the query, and returns a slice of ForwardRestrict objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *forwardRestrictQueryBuilder) Load() (forwardRestricts []*ForwardRestrict) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd_unit")
	results := database.BuilderQuery(b.builder.Ctx, b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(ForwardRestrict)
		o.load(item, o)
		forwardRestricts = append(forwardRestricts, o)
	}
	return
}

// Load terminates the query builder, performs the query, and returns a slice of interfaces.
// This can then satisfy a general interface that loads arrays of objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *forwardRestrictQueryBuilder) LoadI() (forwardRestricts []any) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd_unit")
	results := database.BuilderQuery(b.builder.Ctx, b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(ForwardRestrict)
		o.load(item, o)
		forwardRestricts = append(forwardRestricts, o)
	}
	return
}

// LoadCursor terminates the query builder, performs the query, and returns a cursor to the query.
//
// A query cursor is useful for dealing with large amounts of query results. However, there are some
// limitations to its use. When working with SQL databases, you cannot use a cursor while querying
// many-to-many or reverse relationships that will create an array of values.
//
// Call Next() on the returned cursor object to step through the results. Make sure you call Close
// on the cursor object when you are done. You should use
//
//	defer cursor.Close()
//
// to make sure the cursor gets closed.
func (b *forwardRestrictQueryBuilder) LoadCursor() forwardRestrictsCursor {
	b.builder.Command = query.BuilderCommandLoadCursor
	database := db.GetDatabase("goradd_unit")
	result := database.BuilderQuery(b.builder.Ctx, b.builder)
	if result == nil {
		return forwardRestrictsCursor{}
	}
	cursor := result.(query.CursorI)

	return forwardRestrictsCursor{cursor}
}

type forwardRestrictsCursor struct {
	query.CursorI
}

// Next returns the current ForwardRestrict object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c forwardRestrictsCursor) Next() *ForwardRestrict {
	if c.CursorI == nil {
		return nil
	}

	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(ForwardRestrict)
	o.load(row, o)
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
//
// If an error occurs, or no results are found, a nil is returned.
// In the case of an error, the error is returned in the context.
func (b *forwardRestrictQueryBuilder) Get() *ForwardRestrict {
	results := b.Load()
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *forwardRestrictQueryBuilder) Expand(n query.Expander) ForwardRestrictBuilder {
	b.builder.Expand(n)
	return b
}

/*
// Join attaches the table referred to by joinedTable, filtering the join process using the operation node specified
// by condition.
// The joinedTable node will be modified by this process so that you can use it in subsequent builder operations.
// Call GetAlias to return the resulting object from the query result.
func (b *forwardRestrictQueryBuilder) Join(alias string, joinedTable query.Node, condition query.Node) ForwardRestrictBuilder {
    if query.RootNode(n).TableName_() != "forward_restrict" {
        panic("you can only join a node that is rooted at node.ForwardRestrict()")
    }
    // TODO: make sure joinedTable is a table node
	b.builder.Join(alias, joinedTable, condition)
	return b
}
*/

// Where adds a condition to filter what gets selected.
// Calling Where multiple times will AND the conditions together.
func (b *forwardRestrictQueryBuilder) Where(c query.Node) ForwardRestrictBuilder {
	b.builder.Where(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
// By default, the given nodes are sorted in ascending order.
// Add Descending() to the node to specify that it should be sorted in descending order.
func (b *forwardRestrictQueryBuilder) OrderBy(nodes ...query.Sorter) ForwardRestrictBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified.
// For large data sets and specific types of queries, this can be slow, because it will perform
// the entire query before computing the limit.
// You cannot limit a query that has embedded arrays.
func (b *forwardRestrictQueryBuilder) Limit(maxRowCount int, offset int) ForwardRestrictBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select specifies what specific columns will be loaded with data.
// By default, all the columns of the forward_restrict table will be queried and loaded.
// If nodes contains columns from the forward_restrict table, that will limit the columns queried and loaded to only those columns.
// If related tables are specified, then all the columns from those tables are queried, selected and joined to the result.
// If columns in related tables are specified, then only those columns will be queried and loaded.
// Depending on the query, additional columns may automatically be added to the query. In particular, primary key columns
// will be added in most situations. The exception to this would be in distinct queries, group by queries, or subqueries.
func (b *forwardRestrictQueryBuilder) Select(nodes ...query.Node) ForwardRestrictBuilder {
	b.builder.Select(nodes...)
	return b
}

// Calculation adds a calculation node with an aliased name.
// After the query, you can read the data using GetAlias() on the returned object.
func (b *forwardRestrictQueryBuilder) Calculation(name string, n query.Aliaser) ForwardRestrictBuilder {
	b.builder.Calculation(name, n)
	return b
}

// Distinct removes duplicates from the results of the query.
// Adding a Select() is usually required.
func (b *forwardRestrictQueryBuilder) Distinct() ForwardRestrictBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions with Calculation.
func (b *forwardRestrictQueryBuilder) GroupBy(nodes ...query.Node) ForwardRestrictBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query after the query is performed.
func (b *forwardRestrictQueryBuilder) Having(node query.Node) ForwardRestrictBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
// distinct wll count the number of distinct items, ignoring duplicates.
// nodes will select individual fields, and should be accompanied by a GroupBy.
func (b *forwardRestrictQueryBuilder) Count(distinct bool, nodes ...query.Node) int {
	b.builder.Command = query.BuilderCommandCount
	if distinct {
		b.builder.Distinct()
	}
	database := db.GetDatabase("goradd_unit")
	results := database.BuilderQuery(b.builder.Ctx, b.builder)
	if results == nil {
		return 0
	}
	return results.(int)
}

// Delete uses the query builder to delete a group of records that match the criteria.
func (b *forwardRestrictQueryBuilder) Delete() {
	b.builder.Command = query.BuilderCommandDelete
	database := db.GetDatabase("goradd_unit")
	database.BuilderQuery(b.builder.Ctx, b.builder)
	broadcast.BulkChange(b.builder.Context(), "goradd_unit", "forward_restrict")
}

/*
// Subquery terminates the query builder and tags it as a subquery within a larger query.
// You MUST include what you are selecting by adding Calculation or Select functions on the subquery builder.
// Generally you would use this as a node to a Calculation function on the surrounding query builder.
func (b *forwardRestrictQueryBuilder)  Subquery() *query.SubqueryNode {
	 return b.builder.Subquery()
}
*/

// CountForwardRestrictByID queries the database and returns the number of ForwardRestrict objects that
// have id.
// doc: type=ForwardRestrict
func CountForwardRestrictByID(ctx context.Context, id string) int {
	return int(queryForwardRestricts(ctx).Where(op.Equal(node.ForwardRestrict().ID(), id)).Count(false))
}

// CountForwardRestrictByName queries the database and returns the number of ForwardRestrict objects that
// have name.
// doc: type=ForwardRestrict
func CountForwardRestrictByName(ctx context.Context, name string) int {
	return int(queryForwardRestricts(ctx).Where(op.Equal(node.ForwardRestrict().Name(), name)).Count(false))
}

// CountForwardRestrictByReverseID queries the database and returns the number of ForwardRestrict objects that
// have reverseID.
// doc: type=ForwardRestrict
func CountForwardRestrictByReverseID(ctx context.Context, reverseID string) int {
	if reverseID == "" {
		return 0
	}
	return int(queryForwardRestricts(ctx).Where(op.Equal(node.ForwardRestrict().ReverseID(), reverseID)).Count(false))
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
func (o *forwardRestrictBase) load(m map[string]interface{}, objThis *ForwardRestrict) {

	if v, ok := m["id"]; ok && v != nil {
		if o.id, ok = v.(string); ok {
			o.idIsValid = true
			o.idIsDirty = false

			o._originalPK = o.id

		} else {
			panic("Wrong type found for id.")
		}
	} else {
		o.idIsValid = false
		o.id = ""
	}

	if v, ok := m["name"]; ok && v != nil {
		if o.name, ok = v.(string); ok {
			o.nameIsValid = true
			o.nameIsDirty = false

		} else {
			panic("Wrong type found for name.")
		}
	} else {
		o.nameIsValid = false
		o.name = ""
	}

	if v, ok := m["reverse_id"]; ok && v != nil {
		if o.reverseID, ok = v.(string); ok {
			o.reverseIDIsValid = true
			o.reverseIDIsDirty = false

		} else {
			panic("Wrong type found for reverse_id.")
		}
	} else {
		o.reverseIDIsValid = false
		o.reverseID = ""
	}

	if v, ok := m["Reverse"]; ok {
		if objReverse, ok2 := v.(map[string]interface{}); ok2 {
			o.objReverse = new(Reverse)
			o.objReverse.load(objReverse, o.objReverse)
			o.reverseIDIsValid = true
			o.reverseIDIsDirty = false
		} else {
			panic("Wrong type found for Reverse object.")
		}
	} else {
		o.objReverse = nil
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}

	o._restored = true

}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *forwardRestrictBase) Save(ctx context.Context) {
	if o._restored {
		o.update(ctx)
	} else {
		o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *forwardRestrictBase) update(ctx context.Context) {
	if !o._restored {
		panic("cannot update a record that was not originally read from the database.")
	}

	var modifiedFields map[string]interface{}
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		// TODO: Perform all reads and consistency checks before saves

		// Save loaded Reverse object to get its new pk and update it here.
		if o.objReverse != nil {
			o.objReverse.Save(ctx)
			id := o.objReverse.PrimaryKey()
			o.SetReverseID(id)
		}

		// Save all modified fields to the database
		modifiedFields = o.getModifiedFields()
		if len(modifiedFields) != 0 {
			d.Update(ctx, "forward_restrict", modifiedFields, map[string]any{"id": o._originalPK})
		}

	}) // transaction

	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd_unit", "forward_restrict", o._originalPK, all.SortedKeys(modifiedFields)...)
	}
}

// insert will insert the object into the database. Related items will be saved.
func (o *forwardRestrictBase) insert(ctx context.Context) {
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		if o.objReverse != nil {
			o.objReverse.Save(ctx)
			id := o.objReverse.PrimaryKey()
			o.SetReverseID(id)
		}

		if !o.nameIsValid {
			panic("a value for Name is required, and there is no default value. Call SetName() before inserting the record.")
		}

		if !o.reverseIDIsValid {
			panic("a value for ReverseID is required, and there is no default value. Call SetReverseID() before inserting the record.")
		}

		m := o.getValidFields()

		id := d.Insert(ctx, "forward_restrict", m)
		o.id = id
		o._originalPK = id

	}) // transaction

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd_unit", "forward_restrict", o.PrimaryKey())
}

// getModifiedFields returns the database columns that have been modified. This
// will determine which specific fields are sent to the database to be changed.
func (o *forwardRestrictBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.nameIsDirty {
		fields["name"] = o.name
	}
	if o.reverseIDIsDirty {
		fields["reverse_id"] = o.reverseID
	}
	return
}

// getValidFields returns the fields that have valid data in them in a form ready to send to the database.
func (o *forwardRestrictBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.nameIsValid {
		fields["name"] = o.name
	}
	if o.reverseIDIsValid {
		fields["reverse_id"] = o.reverseID
	}
	return
}

// Delete deletes the record from the database.
func (o *forwardRestrictBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "forward_restrict", map[string]any{"ID": o.id})
	broadcast.Delete(ctx, "goradd_unit", "forward_restrict", fmt.Sprint(o.id))
}

// deleteForwardRestrict deletes the associated record from the database.
func deleteForwardRestrict(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd_unit")
	d.Delete(ctx, "forward_restrict", map[string]any{"ID": pk})
	broadcast.Delete(ctx, "goradd_unit", "forward_restrict", fmt.Sprint(pk))
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *forwardRestrictBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.nameIsDirty = false
	o.reverseIDIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database.
func (o *forwardRestrictBase) IsDirty() (dirty bool) {
	dirty = o.idIsDirty ||
		o.nameIsDirty ||
		o.reverseIDIsDirty ||
		(o.objReverse != nil && o.objReverse.IsDirty())

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *forwardRestrictBase) Get(key string) interface{} {

	switch key {

	case "ID":
		if !o.idIsValid {
			return nil
		}
		return o.id

	case "Name":
		if !o.nameIsValid {
			return nil
		}
		return o.name

	case "ReverseID":
		if !o.reverseIDIsValid {
			return nil
		}
		return o.reverseID

	case "Reverse":
		return o.Reverse()

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *forwardRestrictBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.id); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict.id: %w", err)
	}
	if err := encoder.Encode(o.idIsValid); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict.idIsValid: %w", err)
	}
	if err := encoder.Encode(o.idIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict.idIsDirty: %w", err)
	}

	if err := encoder.Encode(o.name); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict.name: %w", err)
	}
	if err := encoder.Encode(o.nameIsValid); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict.nameIsValid: %w", err)
	}
	if err := encoder.Encode(o.nameIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict.nameIsDirty: %w", err)
	}

	if err := encoder.Encode(o.reverseID); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict.reverseID: %w", err)
	}
	if err := encoder.Encode(o.reverseIDIsValid); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict.reverseIDIsValid: %w", err)
	}
	if err := encoder.Encode(o.reverseIDIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict.reverseIDIsDirty: %w", err)
	}

	if o.objReverse == nil {
		if err := encoder.Encode(false); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(true); err != nil {
			return nil, err
		}
		if err := encoder.Encode(o.objReverse); err != nil {
			return nil, fmt.Errorf("error encoding ForwardRestrict.objReverse: %w", err)
		}
	}

	if o._aliases == nil {
		if err := encoder.Encode(false); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(true); err != nil {
			return nil, err
		}
		if err := encoder.Encode(o._aliases); err != nil {
			return nil, fmt.Errorf("error encoding ForwardRestrict._aliases: %w", err)
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict._restored: %w", err)
	}
	if err := encoder.Encode(o._originalPK); err != nil {
		return nil, fmt.Errorf("error encoding ForwardRestrict._originalPK: %w", err)
	}

	return buf.Bytes(), nil

	return buf.Bytes(), nil
}

// UnmarshalBinary converts a structure that was created with MarshalBinary into a ForwardRestrict object.
func (o *forwardRestrictBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var isPtr bool

	_ = isPtr
	if err = dec.Decode(&o.id); err != nil {
		return fmt.Errorf("error decoding ForwardRestrict.id: %w", err)
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return fmt.Errorf("error decoding ForwardRestrict.idIsValid: %w", err)
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return fmt.Errorf("error decoding ForwardRestrict.idIsDirty: %w", err)
	}

	if err = dec.Decode(&o.name); err != nil {
		return fmt.Errorf("error decoding ForwardRestrict.name: %w", err)
	}
	if err = dec.Decode(&o.nameIsValid); err != nil {
		return fmt.Errorf("error decoding ForwardRestrict.nameIsValid: %w", err)
	}
	if err = dec.Decode(&o.nameIsDirty); err != nil {
		return fmt.Errorf("error decoding ForwardRestrict.nameIsDirty: %w", err)
	}

	if err = dec.Decode(&o.reverseID); err != nil {
		return fmt.Errorf("error decoding ForwardRestrict.reverseID: %w", err)
	}
	if err = dec.Decode(&o.reverseIDIsValid); err != nil {
		return fmt.Errorf("error decoding ForwardRestrict.reverseIDIsValid: %w", err)
	}
	if err = dec.Decode(&o.reverseIDIsDirty); err != nil {
		return fmt.Errorf("error decoding ForwardRestrict.reverseIDIsDirty: %w", err)
	}

	if err = dec.Decode(&isPtr); err != nil {
		return fmt.Errorf("error decoding ForwardRestrict.objReverse isPtr: %w", err)
	}
	if isPtr {
		if err = dec.Decode(&o.objReverse); err != nil {
			return fmt.Errorf("error decoding ForwardRestrict.objReverse: %w", err)
		}
	}
	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *forwardRestrictBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *forwardRestrictBase) MarshalStringMap() map[string]interface{} {
	v := make(map[string]interface{})

	if o.idIsValid {
		v["id"] = o.id
	}

	if o.nameIsValid {
		v["name"] = o.name
	}

	if o.reverseIDIsValid {
		v["reverseID"] = o.reverseID
	}

	if val := o.Reverse(); val != nil {
		v["reverse"] = val.MarshalStringMap()
	}
	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the ForwardRestrict. The ForwardRestrict can be a
// newly created object, or one loaded from the database.
//
// After unmarshalling, the object is not  saved. You must call Save to insert it into the database
// or update it.
//
// Unmarshalling of sub-objects, as in objects linked via foreign keys, is not currently supported.
//
// The fields it expects are:
//
//	"id" - string
//	"name" - string
//	"reverseID" - string
func (o *forwardRestrictBase) UnmarshalJSON(data []byte) (err error) {
	var v map[string]interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	return o.UnmarshalStringMap(v)
}

// UnmarshalStringMap will load the values from the stringmap into the object.
//
// Override this in ForwardRestrict to modify the json before sending it here.
func (o *forwardRestrictBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
	for k, v := range m {
		switch k {

		case "name":
			{
				if v == nil {
					return fmt.Errorf("json field %s cannot be null", k)
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetName(s)
				}
			}

		case "reverseID":
			{
				if v == nil {
					return fmt.Errorf("json field %s cannot be null", k)
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetReverseID(s)
				}

			}

		}
	}
	return
}
