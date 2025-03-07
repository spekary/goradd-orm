// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"github.com/goradd/all"
	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/broadcast"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/query"
)

// EmployeeInfoBase is embedded in a EmployeeInfo object and provides the ORM access to the database.
// The member variables of the structure are private and should not normally be accessed by the EmployeeInfo embedder.
// Instead, use the accessor functions.
type employeeInfoBase struct {
	id        string
	idIsValid bool

	personID        string
	personIDIsValid bool
	personIDIsDirty bool
	objPerson       *Person

	employeeNumber        int
	employeeNumberIsValid bool
	employeeNumberIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]any

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update.
	_restored bool

	_originalPK string
}

// IDs used to access the EmployeeInfo object fields by name using the Get function.
// doc: type=EmployeeInfo
const (
	EmployeeInfo_ID             = `ID`
	EmployeeInfo_PersonID       = `PersonID`
	EmployeeInfo_Person         = `Person`
	EmployeeInfo_EmployeeNumber = `EmployeeNumber`
)

const EmployeeInfoEmployeeNumberMax = 2147483647
const EmployeeInfoEmployeeNumberMin = -2147483648

// Initialize or re-initialize a EmployeeInfo database object to default values.
// The primary key will get a temporary negative number which will be replaced when the object is saved.
// Multiple calls to Initialize are not guaranteed to create sequential values for the primary key.
func (o *employeeInfoBase) Initialize() {

	o.id = db.TemporaryPrimaryKey()

	o.idIsValid = false

	o.personID = ""

	o.personIDIsValid = false
	o.personIDIsDirty = false

	o.employeeNumber = 0

	o.employeeNumberIsValid = false
	o.employeeNumberIsDirty = false

	o._aliases = nil

	o._restored = false
}

// PrimaryKey returns the current value of the primary key field.
func (o *employeeInfoBase) PrimaryKey() string {
	return o.id
}

// OriginalPrimaryKey returns the value of the primary key that was originally loaded into the object when it was
// read from the database.
func (o *employeeInfoBase) OriginalPrimaryKey() string {
	return o._originalPK
}

// Copy copies all valid fields to a new EmployeeInfo object.
// Forward reference ids will be copied, but reverse and many-many references will not.
// Attached objects will not be included in the copy.
// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *employeeInfoBase) Copy() (newObject *EmployeeInfo) {
	newObject = NewEmployeeInfo()
	if o.personIDIsValid {
		newObject.SetPersonID(o.personID)
	}
	if o.employeeNumberIsValid {
		newObject.SetEmployeeNumber(o.employeeNumber)
	}
	return
}

// ID returns the loaded value of ID or
// the zero value if not loaded. Call IDIsValid() to determine
// if it is loaded.
func (o *employeeInfoBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *employeeInfoBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

// PersonID returns the loaded value of PersonID.
func (o *employeeInfoBase) PersonID() string {
	if o._restored && !o.personIDIsValid {
		panic("PersonID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.personID
}

// PersonIDIsValid returns true if the value was loaded from the database or has been set.
func (o *employeeInfoBase) PersonIDIsValid() bool {
	return o.personIDIsValid
}

// SetPersonID sets the value of PersonID in the object, to be saved later in the database using the Save() function.
func (o *employeeInfoBase) SetPersonID(v string) {
	if o._restored &&
		o.personIDIsValid && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.personID == v {
		// no change
		return
	}

	o.personIDIsValid = true
	o.personID = v
	o.personIDIsDirty = true
	o.objPerson = nil
}

// Person returns the current value of the loaded Person, and nil if its not loaded.
func (o *employeeInfoBase) Person() *Person {
	return o.objPerson
}

// LoadPerson returns the related Person. If it is not already loaded,
// it will attempt to load it, provided the PersonID column has been loaded first.
func (o *employeeInfoBase) LoadPerson(ctx context.Context) *Person {
	if !o.personIDIsValid {
		return nil
	}

	if o.objPerson == nil {
		// Load and cache
		o.objPerson = LoadPerson(ctx, o.personID)
	}
	return o.objPerson
}

// SetPerson sets the value of Person in the object, to be saved later using the Save() function.
func (o *employeeInfoBase) SetPerson(objPerson *Person) {
	if objPerson == nil {
		panic("Cannot set Person to a null value.")
	} else {
		o.objPerson = objPerson
		o.personIDIsValid = true
		if o.personID != objPerson.PrimaryKey() {
			o.personID = objPerson.PrimaryKey()
			o.personIDIsDirty = true
		}
	}
}

// EmployeeNumber returns the loaded value of EmployeeNumber.
func (o *employeeInfoBase) EmployeeNumber() int {
	if o._restored && !o.employeeNumberIsValid {
		panic("EmployeeNumber was not selected in the last query and has not been set, and so is not valid")
	}
	return o.employeeNumber
}

// EmployeeNumberIsValid returns true if the value was loaded from the database or has been set.
func (o *employeeInfoBase) EmployeeNumberIsValid() bool {
	return o.employeeNumberIsValid
}

// SetEmployeeNumber sets the value of EmployeeNumber in the object, to be saved later in the database using the Save() function.
func (o *employeeInfoBase) SetEmployeeNumber(v int) {
	if o._restored &&
		o.employeeNumberIsValid && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.employeeNumber == v {
		// no change
		return
	}

	o.employeeNumberIsValid = true
	o.employeeNumber = v
	o.employeeNumberIsDirty = true
}

// GetAlias returns the alias for the given key.
func (o *employeeInfoBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *employeeInfoBase) IsNew() bool {
	return !o._restored
}

// LoadEmployeeInfo returns a EmployeeInfo from the database.
// selectNodes lets you provide nodes for selecting specific fields or additional fields from related tables.
// See [EmployeeInfosBuilder.Select] for more info.
func LoadEmployeeInfo(ctx context.Context, id string, selectNodes ...query.Node) *EmployeeInfo {
	return queryEmployeeInfos(ctx).
		Where(op.Equal(node.EmployeeInfo().ID(), id)).
		Select(selectNodes...).
		Get()
}

// HasEmployeeInfo returns true if a EmployeeInfo with the given primaryKey exists in the database.
// doc: type=EmployeeInfo
func HasEmployeeInfo(ctx context.Context, id string) bool {
	return queryEmployeeInfos(ctx).
		Where(op.Equal(node.EmployeeInfo().ID(), id)).
		Count() == 1
}

// LoadEmployeeInfoByPersonID queries for a single EmployeeInfo object by the given unique index values.
// selectNodes optionally let you provide nodes for joining to other tables or selecting specific fields.
// See [EmployeeInfosBuilder.Select].
// If you need a more elaborate query, use QueryEmployeeInfos() to start a query builder.
func LoadEmployeeInfoByPersonID(ctx context.Context, personID string, selectNodes ...query.Node) *EmployeeInfo {
	q := queryEmployeeInfos(ctx)
	q = q.Where(op.Equal(node.EmployeeInfo().PersonID(), personID))
	return q.Select(selectNodes...).Get()
}

// HasEmployeeInfoByPersonID returns true if the
// given unique index values exist in the database.
// doc: type=EmployeeInfo
func HasEmployeeInfoByPersonID(ctx context.Context, personID string) bool {
	q := queryEmployeeInfos(ctx)
	q = q.Where(op.Equal(node.EmployeeInfo().PersonID(), personID))
	return q.Count() == 1
}

// The EmployeeInfoBuilder uses the query.BuilderI interface to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, LoadCursor, Get, Count, or Delete
type EmployeeInfoBuilder interface {
	// Join(alias string, joinedTable query.Node, condition query.Node) EmployeeInfoBuilder

	// Where adds a condition to filter what gets selected.
	// Calling Where multiple times will AND the conditions together.
	Where(c query.Node) EmployeeInfoBuilder

	// OrderBy specifies how the resulting data should be sorted.
	// By default, the given nodes are sorted in ascending order.
	// Add Descending() to the node to specify that it should be sorted in descending order.
	OrderBy(nodes ...query.Sorter) EmployeeInfoBuilder

	// Limit will return a subset of the data, limited to the offset and number of rows specified.
	// For large data sets and specific types of queries, this can be slow, because it will perform
	// the entire query before computing the limit.
	// You cannot limit a query that has selected a "many" relationship".
	Limit(maxRowCount int, offset int) EmployeeInfoBuilder

	// Select performs two functions:
	//  - Passing a table type node will join the object or objects from that table to this object.
	//  - Passing a column node will optimize the query to only return the specified fields.
	// Once you select at least one column, you must select all the columns that you want in the result.
	// Some fields, like primary keys, are always selected.
	// If you are using a GroupBy, you must select the fields in the GroupBy.
	Select(nodes ...query.Node) EmployeeInfoBuilder

	// Calculation adds a calculation described by operation with the name alias.
	// After the query, you can read the data using GetAlias() on the object identified by base.
	Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) EmployeeInfoBuilder

	// Distinct removes duplicates from the results of the query.
	// Adding a Select() is required.
	Distinct() EmployeeInfoBuilder

	// GroupBy controls how results are grouped when using aggregate functions with Calculation.
	GroupBy(nodes ...query.Node) EmployeeInfoBuilder

	// Having does additional filtering on the results of the query after the query is performed.
	Having(node query.Node) EmployeeInfoBuilder

	// Load terminates the query builder, performs the query, and returns a slice of EmployeeInfo objects.
	// If there are any errors, nil is returned and the specific error is stored in the context.
	// If no results come back from the query, it will return a non-nil empty slice.
	Load() []*EmployeeInfo
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
	LoadCursor() employeeInfosCursor

	// Get is a convenience method to return only the first item found in a query.
	// The entire query is performed, so you should generally use this only if you know
	// you are selecting on one or very few items.
	//
	// If an error occurs, or no results are found, a nil is returned.
	// In the case of an error, the error is returned in the context.
	Get() *EmployeeInfo

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

type employeeInfoQueryBuilder struct {
	builder *query.Builder
}

func newEmployeeInfoBuilder(ctx context.Context) EmployeeInfoBuilder {
	b := employeeInfoQueryBuilder{
		builder: query.NewBuilder(ctx, node.EmployeeInfo()),
	}
	return &b
}

// Load terminates the query builder, performs the query, and returns a slice of EmployeeInfo objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *employeeInfoQueryBuilder) Load() (employeeInfos []*EmployeeInfo) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(EmployeeInfo)
		o.load(item, o)
		employeeInfos = append(employeeInfos, o)
	}
	return
}

// Load terminates the query builder, performs the query, and returns a slice of interfaces.
// This can then satisfy a general interface that loads arrays of objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *employeeInfoQueryBuilder) LoadI() (employeeInfos []any) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(EmployeeInfo)
		o.load(item, o)
		employeeInfos = append(employeeInfos, o)
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
func (b *employeeInfoQueryBuilder) LoadCursor() employeeInfosCursor {
	b.builder.Command = query.BuilderCommandLoadCursor
	database := db.GetDatabase("goradd")
	result := database.BuilderQuery(b.builder)
	if result == nil {
		return employeeInfosCursor{}
	}
	cursor := result.(query.CursorI)

	return employeeInfosCursor{cursor}
}

type employeeInfosCursor struct {
	query.CursorI
}

// Next returns the current EmployeeInfo object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c employeeInfosCursor) Next() *EmployeeInfo {
	if c.CursorI == nil {
		return nil
	}

	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(EmployeeInfo)
	o.load(row, o)
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
//
// If an error occurs, or no results are found, a nil is returned.
// In the case of an error, the error is returned in the context.
func (b *employeeInfoQueryBuilder) Get() *EmployeeInfo {
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
func (b *employeeInfoQueryBuilder) Join(alias string, joinedTable query.Node, condition query.Node) EmployeeInfoBuilder {
    if query.RootNode(n).TableName_() != "employee_info" {
        panic("you can only join a node that is rooted at node.EmployeeInfo()")
    }
    // TODO: make sure joinedTable is a table node
	b.builder.Join(alias, joinedTable, condition)
	return b
}
*/

// Where adds a condition to filter what gets selected.
// Calling Where multiple times will AND the conditions together.
func (b *employeeInfoQueryBuilder) Where(c query.Node) EmployeeInfoBuilder {
	b.builder.Where(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
// By default, the given nodes are sorted in ascending order.
// Add Descending() to the node to specify that it should be sorted in descending order.
func (b *employeeInfoQueryBuilder) OrderBy(nodes ...query.Sorter) EmployeeInfoBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified.
// For large data sets and specific types of queries, this can be slow, because it will perform
// the entire query before computing the limit.
// You cannot limit a query that has embedded arrays.
func (b *employeeInfoQueryBuilder) Limit(maxRowCount int, offset int) EmployeeInfoBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select specifies what specific columns will be loaded with data.
// By default, all the columns of the employee_info table will be queried and loaded.
// If nodes contains columns from the employee_info table, that will limit the columns queried and loaded to only those columns.
// If related tables are specified, then all the columns from those tables are queried, selected and joined to the result.
// If columns in related tables are specified, then only those columns will be queried and loaded.
// Depending on the query, additional columns may automatically be added to the query. In particular, primary key columns
// will be added in most situations. The exception to this would be in distinct queries, group by queries, or subqueries.
func (b *employeeInfoQueryBuilder) Select(nodes ...query.Node) EmployeeInfoBuilder {
	b.builder.Select(nodes...)
	return b
}

// Calculation adds operation as an aliased value onto base.
// After the query, you can read the data by passing alias to GetAlias on the returned object.
func (b *employeeInfoQueryBuilder) Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) EmployeeInfoBuilder {
	b.builder.Calculation(base, alias, operation)
	return b
}

// Distinct removes duplicates from the results of the query.
// Adding a Select() is usually required.
func (b *employeeInfoQueryBuilder) Distinct() EmployeeInfoBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions with Calculation.
func (b *employeeInfoQueryBuilder) GroupBy(nodes ...query.Node) EmployeeInfoBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query after the query is performed.
func (b *employeeInfoQueryBuilder) Having(node query.Node) EmployeeInfoBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items in the result.
// If you have Select or Calculation columns in the query, it will count NULL results as well.
// To not count NULL values, use Where in the builder with a NotNull operation.
// To count distinct combinations of items, call Distinct() on the builder.
func (b *employeeInfoQueryBuilder) Count() int {
	b.builder.Command = query.BuilderCommandCount
	database := db.GetDatabase("goradd")
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
func (b *employeeInfoQueryBuilder)  Subquery() *query.SubqueryNode {
	 return b.builder.Subquery()
}
*/

// CountEmployeeInfoByID queries the database and returns the number of EmployeeInfo objects that
// have id.
// doc: type=EmployeeInfo
func CountEmployeeInfoByID(ctx context.Context, id string) int {
	return queryEmployeeInfos(ctx).Where(op.Equal(node.EmployeeInfo().ID(), id)).Count()
}

// CountEmployeeInfoByPersonID queries the database and returns the number of EmployeeInfo objects that
// have personID.
// doc: type=EmployeeInfo
func CountEmployeeInfoByPersonID(ctx context.Context, personID string) int {
	if personID == "" {
		return 0
	}
	return queryEmployeeInfos(ctx).Where(op.Equal(node.EmployeeInfo().PersonID(), personID)).Count()
}

// CountEmployeeInfoByEmployeeNumber queries the database and returns the number of EmployeeInfo objects that
// have employeeNumber.
// doc: type=EmployeeInfo
func CountEmployeeInfoByEmployeeNumber(ctx context.Context, employeeNumber int) int {
	return queryEmployeeInfos(ctx).Where(op.Equal(node.EmployeeInfo().EmployeeNumber(), employeeNumber)).Count()
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
func (o *employeeInfoBase) load(m map[string]interface{}, objThis *EmployeeInfo) {

	if v, ok := m["id"]; ok && v != nil {
		if o.id, ok = v.(string); ok {
			o.idIsValid = true

			o._originalPK = o.id

		} else {
			panic("Wrong type found for id.")
		}
	} else {
		o.idIsValid = false
		o.id = ""
	}

	if v, ok := m["person_id"]; ok && v != nil {
		if o.personID, ok = v.(string); ok {
			o.personIDIsValid = true
			o.personIDIsDirty = false

		} else {
			panic("Wrong type found for person_id.")
		}
	} else {
		o.personIDIsValid = false
		o.personID = ""
		o.personIDIsDirty = false
	}

	if v, ok := m["Person"]; ok {
		if objPerson, ok2 := v.(map[string]any); ok2 {
			o.objPerson = new(Person)
			o.objPerson.load(objPerson, o.objPerson)
			o.personIDIsValid = true
			o.personIDIsDirty = false
		} else {
			panic("Wrong type found for Person object.")
		}
	} else {
		o.objPerson = nil
	}

	if v, ok := m["employee_number"]; ok && v != nil {
		if o.employeeNumber, ok = v.(int); ok {
			o.employeeNumberIsValid = true
			o.employeeNumberIsDirty = false

		} else {
			panic("Wrong type found for employee_number.")
		}
	} else {
		o.employeeNumberIsValid = false
		o.employeeNumber = 0
		o.employeeNumberIsDirty = false
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
func (o *employeeInfoBase) Save(ctx context.Context) error {
	if o._restored {
		return o.update(ctx)
	} else {
		return o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *employeeInfoBase) update(ctx context.Context) error {
	if !o._restored {
		panic("cannot update a record that was not originally read from the database.")
	}

	var modifiedFields map[string]interface{}
	d := Database()
	err := db.ExecuteTransaction(ctx, d, func() error {

		// Save loaded Person object to get its new pk and update it here.
		if o.objPerson != nil {
			if err := o.objPerson.Save(ctx); err != nil {
				return err
			}
			id := o.objPerson.PrimaryKey()
			o.SetPersonID(id)
		}

		// Save all modified fields to the database
		modifiedFields = o.getModifiedFields()
		if len(modifiedFields) != 0 {
			err := d.Update(ctx, "employee_info", "id", o._originalPK, modifiedFields, "", 0)
			if err != nil {
				return err
			}
		}

		return nil
	}) // transaction
	if err != nil {
		return err
	}

	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd", "employee_info", o._originalPK, all.SortedKeys(modifiedFields)...)
	}

	return nil
}

// insert will insert the object into the database. Related items will be saved.
func (o *employeeInfoBase) insert(ctx context.Context) (err error) {
	d := Database()
	err = db.ExecuteTransaction(ctx, d, func() error {

		if o.objPerson != nil {
			if err = o.objPerson.Save(ctx); err != nil {
				return err
			}
			id := o.objPerson.PrimaryKey()
			o.SetPersonID(id)
		}

		if !o.personIDIsValid {
			panic("a value for PersonID is required, and there is no default value. Call SetPersonID() before inserting the record.")
		}
		if !o.employeeNumberIsValid {
			panic("a value for EmployeeNumber is required, and there is no default value. Call SetEmployeeNumber() before inserting the record.")
		}

		m := o.getValidFields()

		id := d.Insert(ctx, "employee_info", m)
		o.id = id
		o._originalPK = id

		return nil

	}) // transaction

	if err != nil {
		return
	}

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "employee_info", o.PrimaryKey())
	return
}

// getModifiedFields returns the database columns that have been modified. This
// will determine which specific fields are sent to the database to be changed.
func (o *employeeInfoBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.personIDIsDirty {
		fields["person_id"] = o.personID
	}
	if o.employeeNumberIsDirty {
		fields["employee_number"] = o.employeeNumber
	}
	return
}

// getValidFields returns the fields that have valid data in them in a form ready to send to the database.
func (o *employeeInfoBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.personIDIsValid {
		fields["person_id"] = o.personID
	}
	if o.employeeNumberIsValid {
		fields["employee_number"] = o.employeeNumber
	}
	return
}

// Delete deletes the record from the database.
func (o *employeeInfoBase) Delete(ctx context.Context) (err error) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "employee_info", map[string]any{"ID": o.id})
	return nil
	broadcast.Delete(ctx, "goradd", "employee_info", fmt.Sprint(o.id))
	return
}

// deleteEmployeeInfo deletes the EmployeeInfo with primary key pk from the database
// and handles associated records.
func deleteEmployeeInfo(ctx context.Context, pk string) error {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "employee_info", map[string]any{"ID": pk})
	broadcast.Delete(ctx, "goradd", "employee_info", fmt.Sprint(pk))
	return nil
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *employeeInfoBase) resetDirtyStatus() {
	o.personIDIsDirty = false
	o.employeeNumberIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database.
func (o *employeeInfoBase) IsDirty() (dirty bool) {
	dirty = o.personIDIsDirty ||
		(o.objPerson != nil && o.objPerson.IsDirty()) ||
		o.employeeNumberIsDirty

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *employeeInfoBase) Get(key string) interface{} {

	switch key {

	case "ID":
		if !o.idIsValid {
			return nil
		}
		return o.id

	case "PersonID":
		if !o.personIDIsValid {
			return nil
		}
		return o.personID

	case "Person":
		return o.Person()

	case "EmployeeNumber":
		if !o.employeeNumberIsValid {
			return nil
		}
		return o.employeeNumber

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *employeeInfoBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.id); err != nil {
		return nil, fmt.Errorf("error encoding EmployeeInfo.id: %w", err)
	}
	if err := encoder.Encode(o.idIsValid); err != nil {
		return nil, fmt.Errorf("error encoding EmployeeInfo.idIsValid: %w", err)
	}

	if err := encoder.Encode(o.personID); err != nil {
		return nil, fmt.Errorf("error encoding EmployeeInfo.personID: %w", err)
	}
	if err := encoder.Encode(o.personIDIsValid); err != nil {
		return nil, fmt.Errorf("error encoding EmployeeInfo.personIDIsValid: %w", err)
	}
	if err := encoder.Encode(o.personIDIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding EmployeeInfo.personIDIsDirty: %w", err)
	}

	if o.objPerson == nil {
		if err := encoder.Encode(false); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(true); err != nil {
			return nil, err
		}
		if err := encoder.Encode(o.objPerson); err != nil {
			return nil, fmt.Errorf("error encoding EmployeeInfo.objPerson: %w", err)
		}
	}

	if err := encoder.Encode(o.employeeNumber); err != nil {
		return nil, fmt.Errorf("error encoding EmployeeInfo.employeeNumber: %w", err)
	}
	if err := encoder.Encode(o.employeeNumberIsValid); err != nil {
		return nil, fmt.Errorf("error encoding EmployeeInfo.employeeNumberIsValid: %w", err)
	}
	if err := encoder.Encode(o.employeeNumberIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding EmployeeInfo.employeeNumberIsDirty: %w", err)
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
			return nil, fmt.Errorf("error encoding EmployeeInfo._aliases: %w", err)
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, fmt.Errorf("error encoding EmployeeInfo._restored: %w", err)
	}
	if err := encoder.Encode(o._originalPK); err != nil {
		return nil, fmt.Errorf("error encoding EmployeeInfo._originalPK: %w", err)
	}

	return buf.Bytes(), nil

	return buf.Bytes(), nil
}

// UnmarshalBinary converts a structure that was created with MarshalBinary into a EmployeeInfo object.
func (o *employeeInfoBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var isPtr bool

	_ = isPtr
	if err = dec.Decode(&o.id); err != nil {
		return fmt.Errorf("error decoding EmployeeInfo.id: %w", err)
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return fmt.Errorf("error decoding EmployeeInfo.idIsValid: %w", err)
	}

	if err = dec.Decode(&o.personID); err != nil {
		return fmt.Errorf("error decoding EmployeeInfo.personID: %w", err)
	}
	if err = dec.Decode(&o.personIDIsValid); err != nil {
		return fmt.Errorf("error decoding EmployeeInfo.personIDIsValid: %w", err)
	}
	if err = dec.Decode(&o.personIDIsDirty); err != nil {
		return fmt.Errorf("error decoding EmployeeInfo.personIDIsDirty: %w", err)
	}

	if err = dec.Decode(&isPtr); err != nil {
		return fmt.Errorf("error decoding EmployeeInfo.objPerson isPtr: %w", err)
	}
	if isPtr {
		if err = dec.Decode(&o.objPerson); err != nil {
			return fmt.Errorf("error decoding EmployeeInfo.objPerson: %w", err)
		}
	}
	if err = dec.Decode(&o.employeeNumber); err != nil {
		return fmt.Errorf("error decoding EmployeeInfo.employeeNumber: %w", err)
	}
	if err = dec.Decode(&o.employeeNumberIsValid); err != nil {
		return fmt.Errorf("error decoding EmployeeInfo.employeeNumberIsValid: %w", err)
	}
	if err = dec.Decode(&o.employeeNumberIsDirty); err != nil {
		return fmt.Errorf("error decoding EmployeeInfo.employeeNumberIsDirty: %w", err)
	}

	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *employeeInfoBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *employeeInfoBase) MarshalStringMap() map[string]interface{} {
	v := make(map[string]interface{})

	if o.idIsValid {
		v["id"] = o.id
	}

	if o.personIDIsValid {
		v["personID"] = o.personID
	}

	if val := o.Person(); val != nil {
		v["person"] = val.MarshalStringMap()
	}

	if o.employeeNumberIsValid {
		v["employeeNumber"] = o.employeeNumber
	}

	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the EmployeeInfo. The EmployeeInfo can be a
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
//	"personID" - string
//	"employeeNumber" - int
func (o *employeeInfoBase) UnmarshalJSON(data []byte) (err error) {
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
// Override this in EmployeeInfo to modify the json before sending it here.
func (o *employeeInfoBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
	for k, v := range m {
		switch k {

		case "personID":
			{
				if v == nil {
					return fmt.Errorf("field %s cannot be null", k)
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("field %s must be a string", k)
				} else {
					o.SetPersonID(s)
				}

			}

		case "employeeNumber":
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
					o.SetEmployeeNumber(int(n2))
				case int:
					o.SetEmployeeNumber(n)
				case float64:
					o.SetEmployeeNumber(int(n))
				default:
					return fmt.Errorf("field %s must be a number", k)
				}
			}

		}
	}
	return
}
