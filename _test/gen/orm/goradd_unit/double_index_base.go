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

// DoubleIndexBase is embedded in a DoubleIndex object and provides the ORM access to the database.
// The member variables of the structure are private and should not normally be accessed by the DoubleIndex embedder.
// Instead, use the accessor functions.
type doubleIndexBase struct {
	id        int
	idIsValid bool
	idIsDirty bool

	fieldInt        int
	fieldIntIsValid bool
	fieldIntIsDirty bool

	fieldString        string
	fieldStringIsValid bool
	fieldStringIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]any

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update.
	_restored bool

	_originalPK int
}

// IDs used to access the DoubleIndex object fields by name using the Get function.
// doc: type=DoubleIndex
const (
	DoubleIndex_ID          = `ID`
	DoubleIndex_FieldInt    = `FieldInt`
	DoubleIndex_FieldString = `FieldString`
)

const DoubleIndexIDMax = 2147483647
const DoubleIndexIDMin = -2147483648
const DoubleIndexFieldIntMax = 2147483647
const DoubleIndexFieldIntMin = -2147483648
const DoubleIndexFieldStringMaxLength = 50 // The number of runes the column can hold

// Initialize or re-initialize a DoubleIndex database object to default values.
func (o *doubleIndexBase) Initialize() {

	o.id = 0

	o.idIsValid = false
	o.idIsDirty = false

	o.fieldInt = 0

	o.fieldIntIsValid = false
	o.fieldIntIsDirty = false

	o.fieldString = ""

	o.fieldStringIsValid = false
	o.fieldStringIsDirty = false

	o._aliases = nil

	o._restored = false
}

// PrimaryKey returns the current value of the primary key field.
func (o *doubleIndexBase) PrimaryKey() int {
	return o.id
}

// OriginalPrimaryKey returns the value of the primary key that was originally loaded into the object when it was
// read from the database.
func (o *doubleIndexBase) OriginalPrimaryKey() int {
	return o._originalPK
}

// Copy copies all valid fields to a new DoubleIndex object.
// Forward reference ids will be copied, but reverse and many-many references will not.
// Attached objects will not be included in the copy.
// You will need to manually set the primary key field before saving.
// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *doubleIndexBase) Copy() (newObject *DoubleIndex) {
	newObject = NewDoubleIndex()
	if o.idIsValid {
		newObject.SetID(o.id)
	}
	if o.fieldIntIsValid {
		newObject.SetFieldInt(o.fieldInt)
	}
	if o.fieldStringIsValid {
		newObject.SetFieldString(o.fieldString)
	}
	return
}

// ID returns the loaded value of ID.
func (o *doubleIndexBase) ID() int {
	if o._restored && !o.idIsValid {
		panic("ID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.id
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *doubleIndexBase) IDIsValid() bool {
	return o.idIsValid
}

// SetID sets the value of ID in the object, to be saved later in the database using the Save() function.
func (o *doubleIndexBase) SetID(v int) {
	if o._restored &&
		o.idIsValid && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.id == v {
		// no change
		return
	}

	o.idIsValid = true
	o.id = v
	o.idIsDirty = true
}

// FieldInt returns the loaded value of FieldInt.
func (o *doubleIndexBase) FieldInt() int {
	if o._restored && !o.fieldIntIsValid {
		panic("FieldInt was not selected in the last query and has not been set, and so is not valid")
	}
	return o.fieldInt
}

// FieldIntIsValid returns true if the value was loaded from the database or has been set.
func (o *doubleIndexBase) FieldIntIsValid() bool {
	return o.fieldIntIsValid
}

// SetFieldInt sets the value of FieldInt in the object, to be saved later in the database using the Save() function.
func (o *doubleIndexBase) SetFieldInt(v int) {
	if o._restored &&
		o.fieldIntIsValid && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.fieldInt == v {
		// no change
		return
	}

	o.fieldIntIsValid = true
	o.fieldInt = v
	o.fieldIntIsDirty = true
}

// FieldString returns the loaded value of FieldString.
func (o *doubleIndexBase) FieldString() string {
	if o._restored && !o.fieldStringIsValid {
		panic("FieldString was not selected in the last query and has not been set, and so is not valid")
	}
	return o.fieldString
}

// FieldStringIsValid returns true if the value was loaded from the database or has been set.
func (o *doubleIndexBase) FieldStringIsValid() bool {
	return o.fieldStringIsValid
}

// SetFieldString sets the value of FieldString in the object, to be saved later in the database using the Save() function.
func (o *doubleIndexBase) SetFieldString(v string) {
	if utf8.RuneCountInString(v) > DoubleIndexFieldStringMaxLength {
		panic("attempted to set DoubleIndex.FieldString to a value larger than its maximum length in runes")
	}
	if o._restored &&
		o.fieldStringIsValid && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.fieldString == v {
		// no change
		return
	}

	o.fieldStringIsValid = true
	o.fieldString = v
	o.fieldStringIsDirty = true
}

// GetAlias returns the value for the Alias node aliasKey that was returned in the most
// recent query.
func (o *doubleIndexBase) GetAlias(aliasKey string) query.AliasValue {
	if a, ok := o._aliases[aliasKey]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + aliasKey + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *doubleIndexBase) IsNew() bool {
	return !o._restored
}

// LoadDoubleIndex returns a DoubleIndex from the database.
// selectNodes lets you provide nodes for selecting specific fields or additional fields from related tables.
// See [DoubleIndicesBuilder.Select] for more info.
func LoadDoubleIndex(ctx context.Context, id int, selectNodes ...query.Node) *DoubleIndex {
	return queryDoubleIndices(ctx).
		Where(op.Equal(node.DoubleIndex().ID(), id)).
		Select(selectNodes...).
		Get()
}

// HasDoubleIndex returns true if a DoubleIndex with the given primaryKey exists in the database.
// doc: type=DoubleIndex
func HasDoubleIndex(ctx context.Context, id int) bool {
	return queryDoubleIndices(ctx).
		Where(op.Equal(node.DoubleIndex().ID(), id)).
		Count() == 1
}

// LoadDoubleIndexByID queries for a single DoubleIndex object by the given unique index values.
// selectNodes optionally let you provide nodes for joining to other tables or selecting specific fields.
// See [DoubleIndicesBuilder.Select].
// If you need a more elaborate query, use QueryDoubleIndices() to start a query builder.
func LoadDoubleIndexByID(ctx context.Context, id int, selectNodes ...query.Node) *DoubleIndex {
	q := queryDoubleIndices(ctx)
	q = q.Where(op.Equal(node.DoubleIndex().ID(), id))
	return q.Select(selectNodes...).Get()
}

// HasDoubleIndexByID returns true if the
// given unique index values exist in the database.
// doc: type=DoubleIndex
func HasDoubleIndexByID(ctx context.Context, id int) bool {
	q := queryDoubleIndices(ctx)
	q = q.Where(op.Equal(node.DoubleIndex().ID(), id))
	return q.Count() == 1
}

// LoadDoubleIndexByFieldIntFieldString queries for a single DoubleIndex object by the given unique index values.
// selectNodes optionally let you provide nodes for joining to other tables or selecting specific fields.
// See [DoubleIndicesBuilder.Select].
// If you need a more elaborate query, use QueryDoubleIndices() to start a query builder.
func LoadDoubleIndexByFieldIntFieldString(ctx context.Context, fieldInt int, fieldString string, selectNodes ...query.Node) *DoubleIndex {
	q := queryDoubleIndices(ctx)
	q = q.Where(op.Equal(node.DoubleIndex().FieldInt(), fieldInt))
	q = q.Where(op.Equal(node.DoubleIndex().FieldString(), fieldString))
	return q.Select(selectNodes...).Get()
}

// HasDoubleIndexByFieldIntFieldString returns true if the
// given unique index values exist in the database.
// doc: type=DoubleIndex
func HasDoubleIndexByFieldIntFieldString(ctx context.Context, fieldInt int, fieldString string) bool {
	q := queryDoubleIndices(ctx)
	q = q.Where(op.Equal(node.DoubleIndex().FieldInt(), fieldInt))
	q = q.Where(op.Equal(node.DoubleIndex().FieldString(), fieldString))
	return q.Count() == 1
}

// The DoubleIndexBuilder uses the query.BuilderI interface to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, LoadCursor, Get, Count, or Delete
type DoubleIndexBuilder interface {
	// Join(alias string, joinedTable query.Node, condition query.Node) DoubleIndexBuilder

	// Where adds a condition to filter what gets selected.
	// Calling Where multiple times will AND the conditions together.
	Where(c query.Node) DoubleIndexBuilder

	// OrderBy specifies how the resulting data should be sorted.
	// By default, the given nodes are sorted in ascending order.
	// Add Descending() to the node to specify that it should be sorted in descending order.
	OrderBy(nodes ...query.Sorter) DoubleIndexBuilder

	// Limit will return a subset of the data, limited to the offset and number of rows specified.
	// For large data sets and specific types of queries, this can be slow, because it will perform
	// the entire query before computing the limit.
	// You cannot limit a query that has selected a "many" relationship".
	Limit(maxRowCount int, offset int) DoubleIndexBuilder

	// Select performs two functions:
	//  - Passing a table type node will join the object or objects from that table to this object.
	//  - Passing a column node will optimize the query to only return the specified fields.
	// Once you select at least one column, you must select all the columns that you want in the result.
	// Some fields, like primary keys, are always selected.
	// If you are using a GroupBy, you must select the fields in the GroupBy.
	Select(nodes ...query.Node) DoubleIndexBuilder

	// Calculation adds a calculation described by operation with the name alias.
	// After the query, you can read the data using GetAlias() on the object identified by base.
	Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) DoubleIndexBuilder

	// Distinct removes duplicates from the results of the query.
	// Adding a Select() is required.
	Distinct() DoubleIndexBuilder

	// GroupBy controls how results are grouped when using aggregate functions with Calculation.
	GroupBy(nodes ...query.Node) DoubleIndexBuilder

	// Having does additional filtering on the results of the query after the query is performed.
	Having(node query.Node) DoubleIndexBuilder

	// Load terminates the query builder, performs the query, and returns a slice of DoubleIndex objects.
	// If there are any errors, nil is returned and the specific error is stored in the context.
	// If no results come back from the query, it will return a non-nil empty slice.
	Load() []*DoubleIndex
	// Load terminates the query builder, performs the query, and returns a slice of interfaces.
	// This can then satisfy a general interface that loads arrays of objects.
	// If there are any errors, nil is returned and the specific error is stored in the context.
	// If no results come back from the query, it will return a non-nil empty slice.
	LoadI() []query.OrmObj

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
	LoadCursor() doubleIndicesCursor

	// Get is a convenience method to return only the first item found in a query.
	// The entire query is performed, so you should generally use this only if you know
	// you are selecting on one or very few items.
	//
	// If an error occurs, or no results are found, a nil is returned.
	// In the case of an error, the error is returned in the context.
	Get() *DoubleIndex

	// Count terminates a query and returns just the number of items in the result.
	// If you have Select or Calculation columns in the query, it will count NULL results as well.
	// To not count NULL values, use Where in the builder with a NotNull operation.
	// To count distinct combinations of items, call Distinct() on the builder.
	Count() int

	// Subquery terminates the query builder and tags it as a subquery within a larger query.
	// You MUST include what you are selecting by adding Calculation or Select functions on the subquery builder.
	// Generally you would use this as a node to a Calculation function on the surrounding query builder.
	// Subquery() *query.SubqueryNode

}

type doubleIndexQueryBuilder struct {
	builder *query.Builder
}

func newDoubleIndexBuilder(ctx context.Context) DoubleIndexBuilder {
	b := doubleIndexQueryBuilder{
		builder: query.NewBuilder(ctx, node.DoubleIndex()),
	}
	return &b
}

// Load terminates the query builder, performs the query, and returns a slice of DoubleIndex objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *doubleIndexQueryBuilder) Load() (doubleIndices []*DoubleIndex) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd_unit")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(DoubleIndex)
		o.load(item, o)
		doubleIndices = append(doubleIndices, o)
	}
	return
}

// Load terminates the query builder, performs the query, and returns a slice of interfaces.
// This can then satisfy a variety of interfaces that load arrays of objects, including KeyLabeler.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *doubleIndexQueryBuilder) LoadI() (doubleIndices []query.OrmObj) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd_unit")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(DoubleIndex)
		o.load(item, o)
		doubleIndices = append(doubleIndices, o)
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
func (b *doubleIndexQueryBuilder) LoadCursor() doubleIndicesCursor {
	b.builder.Command = query.BuilderCommandLoadCursor
	database := db.GetDatabase("goradd_unit")
	result := database.BuilderQuery(b.builder)
	cursor := result.(query.CursorI)

	return doubleIndicesCursor{cursor}
}

type doubleIndicesCursor struct {
	query.CursorI
}

// Next returns the current DoubleIndex object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c doubleIndicesCursor) Next() *DoubleIndex {
	if c.CursorI == nil {
		return nil
	}

	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(DoubleIndex)
	o.load(row, o)
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
//
// If an error occurs, or no results are found, a nil is returned.
// In the case of an error, the error is returned in the context.
func (b *doubleIndexQueryBuilder) Get() *DoubleIndex {
	results := b.Load()
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

/*
// Join attaches the table referred to by joinedTable, filtering the join process using the operation node specified
// by condition.
// The joinedTable node will be modified by this process so that you can use it in subsequent builder operations.
// Call GetAlias to return the resulting object from the query result.
func (b *doubleIndexQueryBuilder) Join(alias string, joinedTable query.Node, condition query.Node) DoubleIndexBuilder {
    if query.RootNode(n).TableName_() != "double_index" {
        panic("you can only join a node that is rooted at node.DoubleIndex()")
    }
    // TODO: make sure joinedTable is a table node
	b.builder.Join(alias, joinedTable, condition)
	return b
}
*/

// Where adds a condition to filter what gets selected.
// Calling Where multiple times will AND the conditions together.
func (b *doubleIndexQueryBuilder) Where(c query.Node) DoubleIndexBuilder {
	b.builder.Where(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
// By default, the given nodes are sorted in ascending order.
// Add Descending() to the node to specify that it should be sorted in descending order.
func (b *doubleIndexQueryBuilder) OrderBy(nodes ...query.Sorter) DoubleIndexBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified.
// For large data sets and specific types of queries, this can be slow, because it will perform
// the entire query before computing the limit.
// You cannot limit a query that has embedded arrays.
func (b *doubleIndexQueryBuilder) Limit(maxRowCount int, offset int) DoubleIndexBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select specifies what specific columns will be loaded with data.
// By default, all the columns of the double_index table will be queried and loaded.
// If nodes contains columns from the double_index table, that will limit the columns queried and loaded to only those columns.
// If related tables are specified, then all the columns from those tables are queried, selected and joined to the result.
// If columns in related tables are specified, then only those columns will be queried and loaded.
// Depending on the query, additional columns may automatically be added to the query. In particular, primary key columns
// will be added in most situations. The exception to this would be in distinct queries, group by queries, or subqueries.
func (b *doubleIndexQueryBuilder) Select(nodes ...query.Node) DoubleIndexBuilder {
	b.builder.Select(nodes...)
	return b
}

// Calculation adds operation as an aliased value onto base.
// After the query, you can read the data by passing alias to GetAlias on the returned object.
func (b *doubleIndexQueryBuilder) Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) DoubleIndexBuilder {
	b.builder.Calculation(base, alias, operation)
	return b
}

// Distinct removes duplicates from the results of the query.
// Adding a Select() is usually required.
func (b *doubleIndexQueryBuilder) Distinct() DoubleIndexBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions with Calculation.
func (b *doubleIndexQueryBuilder) GroupBy(nodes ...query.Node) DoubleIndexBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query after the query is performed.
func (b *doubleIndexQueryBuilder) Having(node query.Node) DoubleIndexBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items in the result.
// If you have Select or Calculation columns in the query, it will count NULL results as well.
// To not count NULL values, use Where in the builder with a NotNull operation.
// To count distinct combinations of items, call Distinct() on the builder.
func (b *doubleIndexQueryBuilder) Count() int {
	b.builder.Command = query.BuilderCommandCount
	database := db.GetDatabase("goradd_unit")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return 0
	}
	return results.(int)
}

/*
// Subquery terminates the query builder and tags it as a subquery within a larger query.
// You MUST include what you are selecting by adding Calculation or Select functions on the subquery builder.
// Generally you would use this as a node to a Calculation function on the surrounding query builder.
func (b *doubleIndexQueryBuilder)  Subquery() *query.SubqueryNode {
	 return b.builder.Subquery()
}
*/

// CountDoubleIndicesByID queries the database and returns the number of DoubleIndex objects that
// have id.
// doc: type=DoubleIndex
func CountDoubleIndicesByID(ctx context.Context, id int) int {
	return queryDoubleIndices(ctx).Where(op.Equal(node.DoubleIndex().ID(), id)).Count()
}

// CountDoubleIndicesByFieldInt queries the database and returns the number of DoubleIndex objects that
// have fieldInt.
// doc: type=DoubleIndex
func CountDoubleIndicesByFieldInt(ctx context.Context, fieldInt int) int {
	return queryDoubleIndices(ctx).Where(op.Equal(node.DoubleIndex().FieldInt(), fieldInt)).Count()
}

// CountDoubleIndicesByFieldString queries the database and returns the number of DoubleIndex objects that
// have fieldString.
// doc: type=DoubleIndex
func CountDoubleIndicesByFieldString(ctx context.Context, fieldString string) int {
	return queryDoubleIndices(ctx).Where(op.Equal(node.DoubleIndex().FieldString(), fieldString)).Count()
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
func (o *doubleIndexBase) load(m map[string]interface{}, objThis *DoubleIndex) {

	if v, ok := m["id"]; ok && v != nil {
		if o.id, ok = v.(int); ok {
			o.idIsValid = true
			o.idIsDirty = false

			o._originalPK = o.id

		} else {
			panic("Wrong type found for id.")
		}
	} else {
		o.idIsValid = false
		o.id = 0
		o.idIsDirty = false
	}

	if v, ok := m["field_int"]; ok && v != nil {
		if o.fieldInt, ok = v.(int); ok {
			o.fieldIntIsValid = true
			o.fieldIntIsDirty = false

		} else {
			panic("Wrong type found for field_int.")
		}
	} else {
		o.fieldIntIsValid = false
		o.fieldInt = 0
		o.fieldIntIsDirty = false
	}

	if v, ok := m["field_string"]; ok && v != nil {
		if o.fieldString, ok = v.(string); ok {
			o.fieldStringIsValid = true
			o.fieldStringIsDirty = false

		} else {
			panic("Wrong type found for field_string.")
		}
	} else {
		o.fieldStringIsValid = false
		o.fieldString = ""
		o.fieldStringIsDirty = false
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = v.(map[string]any)
	}

	o._restored = true

}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
// Database errors generally will be handled by the logger and not returned here,
// since those indicate a problem with database driver or configuration.
// Save will return a db.OptimisticLockError if it detects a collision when two users
// are attempting to change the same database record.
func (o *doubleIndexBase) Save(ctx context.Context) error {
	if o._restored {
		return o.update(ctx)
	} else {
		return o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *doubleIndexBase) update(ctx context.Context) error {
	if !o._restored {
		panic("cannot update a record that was not originally read from the database.")
	}
	if !o.IsDirty() {
		return nil // nothing to save
	}

	var modifiedFields map[string]interface{}

	d := Database()
	err := db.ExecuteTransaction(ctx, d, func() error {

		modifiedFields = o.getModifiedFields()
		if len(modifiedFields) != 0 {
			var err2 error

			_, err2 = d.Update(ctx, "double_index", "id", o._originalPK, modifiedFields, "", 0)
			if err2 != nil {
				return err2
			}
		}

		return nil
	}) // transaction
	if err != nil {
		return err
	}

	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd_unit", "double_index", o._originalPK, all.SortedKeys(modifiedFields)...)
	}

	return nil
}

// insert will insert the object into the database. Related items will be saved.
func (o *doubleIndexBase) insert(ctx context.Context) (err error) {
	d := Database()
	err = db.ExecuteTransaction(ctx, d, func() error {

		if !o.idIsValid {
			panic("a value for ID is required, and there is no default value. Call SetID() before inserting the record.")
		}
		if !o.fieldIntIsValid {
			panic("a value for FieldInt is required, and there is no default value. Call SetFieldInt() before inserting the record.")
		}
		if !o.fieldStringIsValid {
			panic("a value for FieldString is required, and there is no default value. Call SetFieldString() before inserting the record.")
		}

		m := o.getValidFields()

		d.Insert(ctx, "double_index", m)
		id := o.PrimaryKey()
		o._originalPK = id

		return nil

	}) // transaction

	if err != nil {
		return
	}

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd_unit", "double_index", o.PrimaryKey())
	return
}

// getModifiedFields returns the database columns that have been modified. This
// will determine which specific fields are sent to the database to be changed.
func (o *doubleIndexBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.fieldIntIsDirty {
		fields["field_int"] = o.fieldInt
	}
	if o.fieldStringIsDirty {
		fields["field_string"] = o.fieldString
	}
	return
}

// getValidFields returns the fields that have valid data in them in a form ready to send to the database.
func (o *doubleIndexBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsValid {
		fields["id"] = o.id
	}
	if o.fieldIntIsValid {
		fields["field_int"] = o.fieldInt
	}
	if o.fieldStringIsValid {
		fields["field_string"] = o.fieldString
	}
	return
}

// Delete deletes the record from the database.
func (o *doubleIndexBase) Delete(ctx context.Context) (err error) {
	if o == nil {
		return // allow deleting of a nil object to be a noop
	}
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "double_index", map[string]any{"ID": o.id})
	return nil
	broadcast.Delete(ctx, "goradd_unit", "double_index", fmt.Sprint(o.id))
	return
}

// deleteDoubleIndex deletes the DoubleIndex with primary key pk from the database
// and handles associated records.
func deleteDoubleIndex(ctx context.Context, pk int) error {
	d := db.GetDatabase("goradd_unit")
	d.Delete(ctx, "double_index", map[string]any{"ID": pk})
	broadcast.Delete(ctx, "goradd_unit", "double_index", fmt.Sprint(pk))
	return nil
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *doubleIndexBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.fieldIntIsDirty = false
	o.fieldStringIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database or created.
// However, a new object that has a column with a default value will be automatically marked as dirty upon creation.
func (o *doubleIndexBase) IsDirty() (dirty bool) {
	dirty = o.idIsDirty ||
		o.fieldIntIsDirty ||
		o.fieldStringIsDirty

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil.
// Get can be used to retrieve a value by using the Identifier of a node.
func (o *doubleIndexBase) Get(key string) interface{} {

	switch key {

	case "ID":
		if !o.idIsValid {
			return nil
		}
		return o.id

	case "FieldInt":
		if !o.fieldIntIsValid {
			return nil
		}
		return o.fieldInt

	case "FieldString":
		if !o.fieldStringIsValid {
			return nil
		}
		return o.fieldString

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *doubleIndexBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.id); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex.id: %w", err)
	}
	if err := encoder.Encode(o.idIsValid); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex.idIsValid: %w", err)
	}
	if err := encoder.Encode(o.idIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex.idIsDirty: %w", err)
	}

	if err := encoder.Encode(o.fieldInt); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex.fieldInt: %w", err)
	}
	if err := encoder.Encode(o.fieldIntIsValid); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex.fieldIntIsValid: %w", err)
	}
	if err := encoder.Encode(o.fieldIntIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex.fieldIntIsDirty: %w", err)
	}

	if err := encoder.Encode(o.fieldString); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex.fieldString: %w", err)
	}
	if err := encoder.Encode(o.fieldStringIsValid); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex.fieldStringIsValid: %w", err)
	}
	if err := encoder.Encode(o.fieldStringIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex.fieldStringIsDirty: %w", err)
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
			return nil, fmt.Errorf("error encoding DoubleIndex._aliases: %w", err)
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex._restored: %w", err)
	}
	if err := encoder.Encode(o._originalPK); err != nil {
		return nil, fmt.Errorf("error encoding DoubleIndex._originalPK: %w", err)
	}

	return buf.Bytes(), nil

	return buf.Bytes(), nil
}

// UnmarshalBinary converts a structure that was created with MarshalBinary into a DoubleIndex object.
func (o *doubleIndexBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var isPtr bool

	_ = isPtr
	if err = dec.Decode(&o.id); err != nil {
		return fmt.Errorf("error decoding DoubleIndex.id: %w", err)
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return fmt.Errorf("error decoding DoubleIndex.idIsValid: %w", err)
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return fmt.Errorf("error decoding DoubleIndex.idIsDirty: %w", err)
	}

	if err = dec.Decode(&o.fieldInt); err != nil {
		return fmt.Errorf("error decoding DoubleIndex.fieldInt: %w", err)
	}
	if err = dec.Decode(&o.fieldIntIsValid); err != nil {
		return fmt.Errorf("error decoding DoubleIndex.fieldIntIsValid: %w", err)
	}
	if err = dec.Decode(&o.fieldIntIsDirty); err != nil {
		return fmt.Errorf("error decoding DoubleIndex.fieldIntIsDirty: %w", err)
	}

	if err = dec.Decode(&o.fieldString); err != nil {
		return fmt.Errorf("error decoding DoubleIndex.fieldString: %w", err)
	}
	if err = dec.Decode(&o.fieldStringIsValid); err != nil {
		return fmt.Errorf("error decoding DoubleIndex.fieldStringIsValid: %w", err)
	}
	if err = dec.Decode(&o.fieldStringIsDirty); err != nil {
		return fmt.Errorf("error decoding DoubleIndex.fieldStringIsDirty: %w", err)
	}

	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *doubleIndexBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *doubleIndexBase) MarshalStringMap() map[string]interface{} {
	v := make(map[string]interface{})

	if o.idIsValid {
		v["id"] = o.id
	}

	if o.fieldIntIsValid {
		v["fieldInt"] = o.fieldInt
	}

	if o.fieldStringIsValid {
		v["fieldString"] = o.fieldString
	}

	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the DoubleIndex. The DoubleIndex can be a
// newly created object, or one loaded from the database.
//
// After unmarshalling, the object is not  saved. You must call Save to insert it into the database
// or update it.
//
// Unmarshalling of sub-objects, as in objects linked via foreign keys, is not currently supported.
//
// The fields it expects are:
//
//	"id" - int
//	"fieldInt" - int
//	"fieldString" - string
func (o *doubleIndexBase) UnmarshalJSON(data []byte) (err error) {
	var v map[string]interface{}
	if len(data) == 0 {
		return
	}
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber() // use a number to avoid precision errors
	if err = d.Decode(&v); err != nil {
		return err
	}
	return o.UnmarshalStringMap(v)
}

// UnmarshalStringMap will load the values from the stringmap into the object.
//
// Override this in DoubleIndex to modify the json before sending it here.
func (o *doubleIndexBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
	for k, v := range m {
		switch k {

		case "id":
			{
				if v == nil {
					return fmt.Errorf("field %s cannot be null", k)
				}

				switch n := v.(type) {
				case json.Number:
					n2, err := n.Int64()
					if err != nil {
						return err
					}
					o.SetID(int(n2))
				case int:
					o.SetID(n)
				case float64:
					o.SetID(int(n))
				default:
					return fmt.Errorf("field %s must be a number", k)
				}
			}

		case "fieldInt":
			{
				if v == nil {
					return fmt.Errorf("field %s cannot be null", k)
				}

				switch n := v.(type) {
				case json.Number:
					n2, err := n.Int64()
					if err != nil {
						return err
					}
					o.SetFieldInt(int(n2))
				case int:
					o.SetFieldInt(n)
				case float64:
					o.SetFieldInt(int(n))
				default:
					return fmt.Errorf("field %s must be a number", k)
				}
			}

		case "fieldString":
			{
				if v == nil {
					return fmt.Errorf("field %s cannot be null", k)
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetFieldString(s)
				}
			}

		}
	}
	return
}
