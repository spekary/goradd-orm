// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/goradd/all"
	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/broadcast"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/query"
)

// PersonWithLockBase is embedded in a PersonWithLock object and provides the ORM access to the database.
// The member variables of the structure are private and should not normally be accessed by the PersonWithLock embedder.
// Instead, use the accessor functions.
type personWithLockBase struct {
	id        string
	idIsValid bool
	idIsDirty bool

	firstName        string
	firstNameIsValid bool
	firstNameIsDirty bool

	lastName        string
	lastNameIsValid bool
	lastNameIsDirty bool

	sysTimestamp        time.Time
	sysTimestampIsNull  bool
	sysTimestampIsValid bool
	sysTimestampIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]any

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update.
	_restored bool

	_originalPK string
}

// IDs used to access the PersonWithLock object fields by name using the Get function.
// doc: type=PersonWithLock
const (
	PersonWithLock_ID           = `ID`
	PersonWithLock_FirstName    = `FirstName`
	PersonWithLock_LastName     = `LastName`
	PersonWithLock_SysTimestamp = `SysTimestamp`
)

const PersonWithLockFirstNameMaxLength = 50 // The number of runes the column can hold
const PersonWithLockLastNameMaxLength = 50  // The number of runes the column can hold

// Initialize or re-initialize a PersonWithLock database object to default values.
// The primary key will get a temporary negative number which will be replaced when the object is saved.
// Multiple calls to Initialize are not guaranteed to create sequential values for the primary key.
func (o *personWithLockBase) Initialize() {

	newObjectPkCounter = newObjectPkCounter - 1
	o.id = fmt.Sprintf("%d", newObjectPkCounter)

	o.idIsValid = false
	o.idIsDirty = false

	o.firstName = ""

	o.firstNameIsValid = false
	o.firstNameIsDirty = false

	o.lastName = ""

	o.lastNameIsValid = false
	o.lastNameIsDirty = false

	o.sysTimestamp = time.Time{}

	o.sysTimestampIsNull = false
	o.sysTimestampIsValid = true
	o.sysTimestampIsDirty = true

	o._restored = false
}

// PrimaryKey returns the current value of the primary key field.
func (o *personWithLockBase) PrimaryKey() string {
	return o.id
}

// OriginalPrimaryKey returns the value of the primary key that was originally loaded into the object when it was
// read from the database.
func (o *personWithLockBase) OriginalPrimaryKey() string {
	return o._originalPK
}

// Copy copies all valid fields to a new PersonWithLock object.
// Forward reference ids will be copied, but reverse and many-many references will not.
// Attached objects will not be included in the copy.
// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *personWithLockBase) Copy() (newObject *PersonWithLock) {
	newObject = NewPersonWithLock()
	if o.firstNameIsValid {
		newObject.SetFirstName(o.firstName)
	}
	if o.lastNameIsValid {
		newObject.SetLastName(o.lastName)
	}
	return
}

// ID returns the loaded value of ID or
// the zero value if not loaded. Call IDIsValid() to determine
// if it is loaded.
func (o *personWithLockBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *personWithLockBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

// FirstName returns the loaded value of FirstName.
func (o *personWithLockBase) FirstName() string {
	if o._restored && !o.firstNameIsValid {
		panic("FirstName was not selected in the last query and has not been set, and so is not valid")
	}
	return o.firstName
}

// FirstNameIsValid returns true if the value was loaded from the database or has been set.
func (o *personWithLockBase) FirstNameIsValid() bool {
	return o.firstNameIsValid
}

// SetFirstName sets the value of FirstName in the object, to be saved later in the database using the Save() function.
func (o *personWithLockBase) SetFirstName(v string) {
	if utf8.RuneCountInString(v) > PersonWithLockFirstNameMaxLength {
		panic("attempted to set PersonWithLock.FirstName to a value larger than its maximum length in runes")
	}
	if o._restored &&
		o.firstNameIsValid && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.firstName == v {
		// no change
		return
	}

	o.firstNameIsValid = true
	o.firstName = v
	o.firstNameIsDirty = true
}

// LastName returns the loaded value of LastName.
func (o *personWithLockBase) LastName() string {
	if o._restored && !o.lastNameIsValid {
		panic("LastName was not selected in the last query and has not been set, and so is not valid")
	}
	return o.lastName
}

// LastNameIsValid returns true if the value was loaded from the database or has been set.
func (o *personWithLockBase) LastNameIsValid() bool {
	return o.lastNameIsValid
}

// SetLastName sets the value of LastName in the object, to be saved later in the database using the Save() function.
func (o *personWithLockBase) SetLastName(v string) {
	if utf8.RuneCountInString(v) > PersonWithLockLastNameMaxLength {
		panic("attempted to set PersonWithLock.LastName to a value larger than its maximum length in runes")
	}
	if o._restored &&
		o.lastNameIsValid && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.lastName == v {
		// no change
		return
	}

	o.lastNameIsValid = true
	o.lastName = v
	o.lastNameIsDirty = true
}

// SysTimestamp returns the loaded value of SysTimestamp.
func (o *personWithLockBase) SysTimestamp() time.Time {
	if o._restored && !o.sysTimestampIsValid {
		panic("SysTimestamp was not selected in the last query and has not been set, and so is not valid")
	}
	return o.sysTimestamp
}

// SysTimestampIsValid returns true if the value was loaded from the database or has been set.
func (o *personWithLockBase) SysTimestampIsValid() bool {
	return o.sysTimestampIsValid
}

// SysTimestampIsNull returns true if the related database value is null.
func (o *personWithLockBase) SysTimestampIsNull() bool {
	return o.sysTimestampIsNull
}

// SysTimestamp_I returns the loaded value of SysTimestamp as an interface.
// If the value in the database is NULL, a nil interface is returned.
func (o *personWithLockBase) SysTimestamp_I() interface{} {
	if o._restored && !o.sysTimestampIsValid {
		panic("sysTimestamp was not selected in the last query and has not been set, and so is not valid")
	} else if o.sysTimestampIsNull {
		return nil
	}
	return o.sysTimestamp
}

// SetSysTimestampToNull() will set the sys_timestamp value in the database to NULL.
// SysTimestamp() will return the column's default value after this.
func (o *personWithLockBase) SetSysTimestampToNull() {
	if !o.sysTimestampIsValid || !o.sysTimestampIsNull {
		// If we know it is null in the database, don't save it
		o.sysTimestampIsDirty = true
	}
	o.sysTimestampIsValid = true
	o.sysTimestampIsNull = true
	o.sysTimestamp = time.Time{}
}

// GetAlias returns the alias for the given key.
func (o *personWithLockBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *personWithLockBase) IsNew() bool {
	return !o._restored
}

// LoadPersonWithLock returns a PersonWithLock from the database.
// selectNodes lets you provide nodes for selecting specific fields or additional fields from related tables.
// See [PersonWithLocksBuilder.Select] for more info.
func LoadPersonWithLock(ctx context.Context, id string, selectNodes ...query.Node) *PersonWithLock {
	return queryPersonWithLocks(ctx).
		Where(op.Equal(node.PersonWithLock().ID(), id)).
		Select(selectNodes...).
		Get()
}

// HasPersonWithLock returns true if a PersonWithLock with the given primaryKey exists in the database.
// doc: type=PersonWithLock
func HasPersonWithLock(ctx context.Context, id string) bool {
	return queryPersonWithLocks(ctx).
		Where(op.Equal(node.PersonWithLock().ID(), id)).
		Count() == 1
}

// The PersonWithLockBuilder uses the query.BuilderI interface to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, LoadCursor, Get, Count, or Delete
type PersonWithLockBuilder interface {
	// Join(alias string, joinedTable query.Node, condition query.Node) PersonWithLockBuilder

	// Where adds a condition to filter what gets selected.
	// Calling Where multiple times will AND the conditions together.
	Where(c query.Node) PersonWithLockBuilder

	// OrderBy specifies how the resulting data should be sorted.
	// By default, the given nodes are sorted in ascending order.
	// Add Descending() to the node to specify that it should be sorted in descending order.
	OrderBy(nodes ...query.Sorter) PersonWithLockBuilder

	// Limit will return a subset of the data, limited to the offset and number of rows specified.
	// For large data sets and specific types of queries, this can be slow, because it will perform
	// the entire query before computing the limit.
	// You cannot limit a query that has embedded arrays.
	Limit(maxRowCount int, offset int) PersonWithLockBuilder

	// Select optimizes the query to only return the specified fields.
	// Once you put a Select in your query, you must specify all the fields that you will eventually read out.
	// Some fields, like primary keys, are always selected.
	// If you are using a GroupBy, most database drivers will only allow selecting on fields in the GroupBy, and
	// doing otherwise will result in an error.
	Select(nodes ...query.Node) PersonWithLockBuilder

	// Calculation adds a calculation node with an aliased name.
	// After the query, you can read the data using GetAlias() on a returned object.
	Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) PersonWithLockBuilder

	// Distinct removes duplicates from the results of the query.
	// Adding a Select() is required.
	Distinct() PersonWithLockBuilder

	// GroupBy controls how results are grouped when using aggregate functions with Calculation.
	GroupBy(nodes ...query.Node) PersonWithLockBuilder

	// Having does additional filtering on the results of the query after the query is performed.
	Having(node query.Node) PersonWithLockBuilder

	// Load terminates the query builder, performs the query, and returns a slice of PersonWithLock objects.
	// If there are any errors, nil is returned and the specific error is stored in the context.
	// If no results come back from the query, it will return a non-nil empty slice.
	Load() []*PersonWithLock
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
	LoadCursor() personWithLocksCursor

	// Get is a convenience method to return only the first item found in a query.
	// The entire query is performed, so you should generally use this only if you know
	// you are selecting on one or very few items.
	//
	// If an error occurs, or no results are found, a nil is returned.
	// In the case of an error, the error is returned in the context.
	Get() *PersonWithLock

	// Count terminates a query and returns just the number of items in the result.
	// If you have Select or Calculation columns in the query, it will count NULL results as well.
	// To not count NULL values, use Where in the builder with a NotNull operation.
	// To count distinct combinations of items, call Distinct() on the builder.
	Count() int

	// Delete uses the query builder to delete a group of records that match the criteria
	Delete()

	// Subquery terminates the query builder and tags it as a subquery within a larger query.
	// You MUST include what you are selecting by adding Calculation or Select functions on the subquery builder.
	// Generally you would use this as a node to a Calculation function on the surrounding query builder.
	// Subquery() *query.SubqueryNode

}

type personWithLockQueryBuilder struct {
	builder *query.Builder
}

func newPersonWithLockBuilder(ctx context.Context) PersonWithLockBuilder {
	b := personWithLockQueryBuilder{
		builder: query.NewBuilder(ctx, node.PersonWithLock()),
	}
	return &b
}

// Load terminates the query builder, performs the query, and returns a slice of PersonWithLock objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *personWithLockQueryBuilder) Load() (personWithLocks []*PersonWithLock) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(PersonWithLock)
		o.load(item, o)
		personWithLocks = append(personWithLocks, o)
	}
	return
}

// Load terminates the query builder, performs the query, and returns a slice of interfaces.
// This can then satisfy a general interface that loads arrays of objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *personWithLockQueryBuilder) LoadI() (personWithLocks []any) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(PersonWithLock)
		o.load(item, o)
		personWithLocks = append(personWithLocks, o)
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
func (b *personWithLockQueryBuilder) LoadCursor() personWithLocksCursor {
	b.builder.Command = query.BuilderCommandLoadCursor
	database := db.GetDatabase("goradd")
	result := database.BuilderQuery(b.builder)
	if result == nil {
		return personWithLocksCursor{}
	}
	cursor := result.(query.CursorI)

	return personWithLocksCursor{cursor}
}

type personWithLocksCursor struct {
	query.CursorI
}

// Next returns the current PersonWithLock object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c personWithLocksCursor) Next() *PersonWithLock {
	if c.CursorI == nil {
		return nil
	}

	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(PersonWithLock)
	o.load(row, o)
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
//
// If an error occurs, or no results are found, a nil is returned.
// In the case of an error, the error is returned in the context.
func (b *personWithLockQueryBuilder) Get() *PersonWithLock {
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
func (b *personWithLockQueryBuilder) Join(alias string, joinedTable query.Node, condition query.Node) PersonWithLockBuilder {
    if query.RootNode(n).TableName_() != "person_with_lock" {
        panic("you can only join a node that is rooted at node.PersonWithLock()")
    }
    // TODO: make sure joinedTable is a table node
	b.builder.Join(alias, joinedTable, condition)
	return b
}
*/

// Where adds a condition to filter what gets selected.
// Calling Where multiple times will AND the conditions together.
func (b *personWithLockQueryBuilder) Where(c query.Node) PersonWithLockBuilder {
	b.builder.Where(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
// By default, the given nodes are sorted in ascending order.
// Add Descending() to the node to specify that it should be sorted in descending order.
func (b *personWithLockQueryBuilder) OrderBy(nodes ...query.Sorter) PersonWithLockBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified.
// For large data sets and specific types of queries, this can be slow, because it will perform
// the entire query before computing the limit.
// You cannot limit a query that has embedded arrays.
func (b *personWithLockQueryBuilder) Limit(maxRowCount int, offset int) PersonWithLockBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select specifies what specific columns will be loaded with data.
// By default, all the columns of the person_with_lock table will be queried and loaded.
// If nodes contains columns from the person_with_lock table, that will limit the columns queried and loaded to only those columns.
// If related tables are specified, then all the columns from those tables are queried, selected and joined to the result.
// If columns in related tables are specified, then only those columns will be queried and loaded.
// Depending on the query, additional columns may automatically be added to the query. In particular, primary key columns
// will be added in most situations. The exception to this would be in distinct queries, group by queries, or subqueries.
func (b *personWithLockQueryBuilder) Select(nodes ...query.Node) PersonWithLockBuilder {
	b.builder.Select(nodes...)
	return b
}

// Calculation adds operation as an aliased value onto base.
// After the query, you can read the data by passing alias to GetAlias on the returned object.
func (b *personWithLockQueryBuilder) Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) PersonWithLockBuilder {
	b.builder.Calculation(base, alias, operation)
	return b
}

// Distinct removes duplicates from the results of the query.
// Adding a Select() is usually required.
func (b *personWithLockQueryBuilder) Distinct() PersonWithLockBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions with Calculation.
func (b *personWithLockQueryBuilder) GroupBy(nodes ...query.Node) PersonWithLockBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query after the query is performed.
func (b *personWithLockQueryBuilder) Having(node query.Node) PersonWithLockBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items in the result.
// If you have Select or Calculation columns in the query, it will count NULL results as well.
// To not count NULL values, use Where in the builder with a NotNull operation.
// To count distinct combinations of items, call Distinct() on the builder.
func (b *personWithLockQueryBuilder) Count() int {
	b.builder.Command = query.BuilderCommandCount
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return 0
	}
	return results.(int)
}

// Delete uses the query builder to delete a group of records that match the criteria.
func (b *personWithLockQueryBuilder) Delete() {
	b.builder.Command = query.BuilderCommandDelete
	database := db.GetDatabase("goradd")
	database.BuilderQuery(b.builder)
	broadcast.BulkChange(b.builder.Context(), "goradd", "person_with_lock")
}

/*
// Subquery terminates the query builder and tags it as a subquery within a larger query.
// You MUST include what you are selecting by adding Calculation or Select functions on the subquery builder.
// Generally you would use this as a node to a Calculation function on the surrounding query builder.
func (b *personWithLockQueryBuilder)  Subquery() *query.SubqueryNode {
	 return b.builder.Subquery()
}
*/

// CountPersonWithLockByID queries the database and returns the number of PersonWithLock objects that
// have id.
// doc: type=PersonWithLock
func CountPersonWithLockByID(ctx context.Context, id string) int {
	return queryPersonWithLocks(ctx).Where(op.Equal(node.PersonWithLock().ID(), id)).Count()
}

// CountPersonWithLockByFirstName queries the database and returns the number of PersonWithLock objects that
// have firstName.
// doc: type=PersonWithLock
func CountPersonWithLockByFirstName(ctx context.Context, firstName string) int {
	return queryPersonWithLocks(ctx).Where(op.Equal(node.PersonWithLock().FirstName(), firstName)).Count()
}

// CountPersonWithLockByLastName queries the database and returns the number of PersonWithLock objects that
// have lastName.
// doc: type=PersonWithLock
func CountPersonWithLockByLastName(ctx context.Context, lastName string) int {
	return queryPersonWithLocks(ctx).Where(op.Equal(node.PersonWithLock().LastName(), lastName)).Count()
}

// CountPersonWithLockBySysTimestamp queries the database and returns the number of PersonWithLock objects that
// have sysTimestamp.
// doc: type=PersonWithLock
func CountPersonWithLockBySysTimestamp(ctx context.Context, sysTimestamp time.Time) int {
	return queryPersonWithLocks(ctx).Where(op.Equal(node.PersonWithLock().SysTimestamp(), sysTimestamp)).Count()
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
func (o *personWithLockBase) load(m map[string]interface{}, objThis *PersonWithLock) {

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

	if v, ok := m["first_name"]; ok && v != nil {
		if o.firstName, ok = v.(string); ok {
			o.firstNameIsValid = true
			o.firstNameIsDirty = false

		} else {
			panic("Wrong type found for first_name.")
		}
	} else {
		o.firstNameIsValid = false
		o.firstName = ""
	}

	if v, ok := m["last_name"]; ok && v != nil {
		if o.lastName, ok = v.(string); ok {
			o.lastNameIsValid = true
			o.lastNameIsDirty = false

		} else {
			panic("Wrong type found for last_name.")
		}
	} else {
		o.lastNameIsValid = false
		o.lastName = ""
	}

	if v, ok := m["sys_timestamp"]; ok {
		if v == nil {
			o.sysTimestamp = time.Time{}
			o.sysTimestampIsNull = true
			o.sysTimestampIsValid = true
			o.sysTimestampIsDirty = false
		} else if o.sysTimestamp, ok = v.(time.Time); ok {
			o.sysTimestampIsNull = false
			o.sysTimestampIsValid = true
			o.sysTimestampIsDirty = false
		} else {
			panic("Wrong type found for sys_timestamp.")
		}
	} else {
		o.sysTimestampIsValid = false
		o.sysTimestampIsNull = true
		o.sysTimestamp = time.Time{}
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = v.(map[string]any)
	}

	o._restored = true

}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *personWithLockBase) Save(ctx context.Context) {
	if o._restored {
		o.update(ctx)
	} else {
		o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *personWithLockBase) update(ctx context.Context) {
	if !o._restored {
		panic("cannot update a record that was not originally read from the database.")
	}

	var modifiedFields map[string]interface{}
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		// TODO: Perform all reads and consistency checks before saves

		// Save all modified fields to the database
		modifiedFields = o.getModifiedFields()
		if len(modifiedFields) != 0 {
			d.Update(ctx, "person_with_lock", modifiedFields, map[string]any{"id": o._originalPK})
		}

	}) // transaction

	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd", "person_with_lock", o._originalPK, all.SortedKeys(modifiedFields)...)
	}
}

// insert will insert the object into the database. Related items will be saved.
func (o *personWithLockBase) insert(ctx context.Context) {
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		if !o.firstNameIsValid {
			panic("a value for FirstName is required, and there is no default value. Call SetFirstName() before inserting the record.")
		}

		if !o.lastNameIsValid {
			panic("a value for LastName is required, and there is no default value. Call SetLastName() before inserting the record.")
		}

		m := o.getValidFields()

		id := d.Insert(ctx, "person_with_lock", m)
		o.id = id
		o._originalPK = id

	}) // transaction

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "person_with_lock", o.PrimaryKey())
}

// getModifiedFields returns the database columns that have been modified. This
// will determine which specific fields are sent to the database to be changed.
func (o *personWithLockBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.firstNameIsDirty {
		fields["first_name"] = o.firstName
	}
	if o.lastNameIsDirty {
		fields["last_name"] = o.lastName
	}
	if o.sysTimestampIsDirty {
		if o.sysTimestampIsNull {
			fields["sys_timestamp"] = nil
		} else {
			fields["sys_timestamp"] = o.sysTimestamp
		}
	}
	return
}

// getValidFields returns the fields that have valid data in them in a form ready to send to the database.
func (o *personWithLockBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.firstNameIsValid {
		fields["first_name"] = o.firstName
	}
	if o.lastNameIsValid {
		fields["last_name"] = o.lastName
	}
	if o.sysTimestampIsValid {
		if o.sysTimestampIsNull {
			fields["sys_timestamp"] = nil
		} else {
			fields["sys_timestamp"] = o.sysTimestamp
		}
	}
	return
}

// Delete deletes the record from the database.
func (o *personWithLockBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "person_with_lock", map[string]any{"ID": o.id})
	broadcast.Delete(ctx, "goradd", "person_with_lock", fmt.Sprint(o.id))
}

// deletePersonWithLock deletes the associated record from the database.
func deletePersonWithLock(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "person_with_lock", map[string]any{"ID": pk})
	broadcast.Delete(ctx, "goradd", "person_with_lock", fmt.Sprint(pk))
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *personWithLockBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.firstNameIsDirty = false
	o.lastNameIsDirty = false
	o.sysTimestampIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database.
func (o *personWithLockBase) IsDirty() (dirty bool) {
	dirty = o.idIsDirty ||
		o.firstNameIsDirty ||
		o.lastNameIsDirty ||
		o.sysTimestampIsDirty

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *personWithLockBase) Get(key string) interface{} {

	switch key {

	case "ID":
		if !o.idIsValid {
			return nil
		}
		return o.id

	case "FirstName":
		if !o.firstNameIsValid {
			return nil
		}
		return o.firstName

	case "LastName":
		if !o.lastNameIsValid {
			return nil
		}
		return o.lastName

	case "SysTimestamp":
		if !o.sysTimestampIsValid {
			return nil
		}
		return o.sysTimestamp

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *personWithLockBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.id); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.id: %w", err)
	}
	if err := encoder.Encode(o.idIsValid); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.idIsValid: %w", err)
	}
	if err := encoder.Encode(o.idIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.idIsDirty: %w", err)
	}

	if err := encoder.Encode(o.firstName); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.firstName: %w", err)
	}
	if err := encoder.Encode(o.firstNameIsValid); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.firstNameIsValid: %w", err)
	}
	if err := encoder.Encode(o.firstNameIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.firstNameIsDirty: %w", err)
	}

	if err := encoder.Encode(o.lastName); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.lastName: %w", err)
	}
	if err := encoder.Encode(o.lastNameIsValid); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.lastNameIsValid: %w", err)
	}
	if err := encoder.Encode(o.lastNameIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.lastNameIsDirty: %w", err)
	}

	if err := encoder.Encode(o.sysTimestamp); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.sysTimestamp: %w", err)
	}
	if err := encoder.Encode(o.sysTimestampIsNull); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.sysTimestampIsNull: %w", err)
	}
	if err := encoder.Encode(o.sysTimestampIsValid); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.sysTimestampIsValid: %w", err)
	}
	if err := encoder.Encode(o.sysTimestampIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock.sysTimestampIsDirty: %w", err)
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
			return nil, fmt.Errorf("error encoding PersonWithLock._aliases: %w", err)
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock._restored: %w", err)
	}
	if err := encoder.Encode(o._originalPK); err != nil {
		return nil, fmt.Errorf("error encoding PersonWithLock._originalPK: %w", err)
	}

	return buf.Bytes(), nil

	return buf.Bytes(), nil
}

// UnmarshalBinary converts a structure that was created with MarshalBinary into a PersonWithLock object.
func (o *personWithLockBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var isPtr bool

	_ = isPtr
	if err = dec.Decode(&o.id); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.id: %w", err)
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.idIsValid: %w", err)
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.idIsDirty: %w", err)
	}

	if err = dec.Decode(&o.firstName); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.firstName: %w", err)
	}
	if err = dec.Decode(&o.firstNameIsValid); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.firstNameIsValid: %w", err)
	}
	if err = dec.Decode(&o.firstNameIsDirty); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.firstNameIsDirty: %w", err)
	}

	if err = dec.Decode(&o.lastName); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.lastName: %w", err)
	}
	if err = dec.Decode(&o.lastNameIsValid); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.lastNameIsValid: %w", err)
	}
	if err = dec.Decode(&o.lastNameIsDirty); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.lastNameIsDirty: %w", err)
	}

	if err = dec.Decode(&o.sysTimestamp); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.sysTimestamp: %w", err)
	}
	if err = dec.Decode(&o.sysTimestampIsNull); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.sysTimestampIsNull: %w", err)
	}
	if err = dec.Decode(&o.sysTimestampIsValid); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.sysTimestampIsValid: %w", err)
	}
	if err = dec.Decode(&o.sysTimestampIsDirty); err != nil {
		return fmt.Errorf("error decoding PersonWithLock.sysTimestampIsDirty: %w", err)
	}

	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *personWithLockBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *personWithLockBase) MarshalStringMap() map[string]interface{} {
	v := make(map[string]interface{})

	if o.idIsValid {
		v["id"] = o.id
	}

	if o.firstNameIsValid {
		v["firstName"] = o.firstName
	}

	if o.lastNameIsValid {
		v["lastName"] = o.lastName
	}

	if o.sysTimestampIsValid {
		if o.sysTimestampIsNull {
			v["sysTimestamp"] = nil
		} else {
			v["sysTimestamp"] = o.sysTimestamp
		}
	}

	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the PersonWithLock. The PersonWithLock can be a
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
//	"firstName" - string
//	"lastName" - string
//	"sysTimestamp" - time.Time, nullable
func (o *personWithLockBase) UnmarshalJSON(data []byte) (err error) {
	var v map[string]interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	return o.UnmarshalStringMap(v)
}

// UnmarshalStringMap will load the values from the stringmap into the object.
//
// Override this in PersonWithLock to modify the json before sending it here.
func (o *personWithLockBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
	for k, v := range m {
		switch k {

		case "firstName":
			{
				if v == nil {
					return fmt.Errorf("json field %s cannot be null", k)
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetFirstName(s)
				}
			}

		case "lastName":
			{
				if v == nil {
					return fmt.Errorf("json field %s cannot be null", k)
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetLastName(s)
				}
			}

		}
	}
	return
}
