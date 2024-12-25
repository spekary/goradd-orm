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
	idIsDirty bool

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

// Initialize or re-initialize a EmployeeInfo database object to default values.
// The primary key will get a temporary negative number which will be replaced when the object is saved.
// Multiple calls to Initialize are not guaranteed to create sequential values for the primary key.
func (o *employeeInfoBase) Initialize() {

	newObjectPkCounter = newObjectPkCounter - 1
	o.id = fmt.Sprintf("%d", newObjectPkCounter)

	o.idIsValid = false
	o.idIsDirty = false

	o.personID = ""

	o.personIDIsValid = false
	o.personIDIsDirty = false

	o.employeeNumber = 0

	o.employeeNumberIsValid = false
	o.employeeNumberIsDirty = false

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

// Copy copies all valid fields (except for the primary key) to a new EmployeeInfo object.
// Forward reference ids will be copied, but reverse and many-many references will not.

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

// SetPersonID sets the value of PersonID in the object, to be saved later using the Save() function.
func (o *employeeInfoBase) SetPersonID(personID string) {
	o.personIDIsValid = true
	if o.personID != personID || !o._restored {
		o.personID = personID
		o.personIDIsDirty = true
		o.objPerson = nil
	}
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

// SetEmployeeNumber sets the value of EmployeeNumber in the object, to be saved later using the Save() function.
func (o *employeeInfoBase) SetEmployeeNumber(employeeNumber int) {
	o.employeeNumberIsValid = true
	if o.employeeNumber != employeeNumber || !o._restored {
		o.employeeNumber = employeeNumber
		o.employeeNumberIsDirty = true
	}
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
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See [EmployeeInfosBuilder.Join] and [EmployeeInfosBuilder.Select] for more info.
func LoadEmployeeInfo(ctx context.Context, id string, joinOrSelectNodes ...query.NodeI) *EmployeeInfo {
	return queryEmployeeInfos(ctx).
		Where(op.Equal(node.EmployeeInfo().ID(), id)).
		joinOrSelect(joinOrSelectNodes...).
		Get()
}

// HasEmployeeInfo returns true if a EmployeeInfo with the given primaryKey exists in the database.
// doc: type=EmployeeInfo
func HasEmployeeInfo(ctx context.Context, id string) bool {
	return queryEmployeeInfos(ctx).
		Where(op.Equal(node.EmployeeInfo().ID(), id)).
		Count(false) == 1
}

// LoadEmployeeInfoByPersonID queries for a single EmployeeInfo object by the given unique index values.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See [EmployeeInfosBuilder.Join] and [EmployeeInfosBuilder.Select] for more info.
// If you need a more elaborate query, use QueryEmployeeInfos() to start a query builder.
func LoadEmployeeInfoByPersonID(ctx context.Context, personID string, joinOrSelectNodes ...query.NodeI) *EmployeeInfo {
	q := queryEmployeeInfos(ctx)
	q = q.Where(op.Equal(node.EmployeeInfo().PersonID(), personID))
	return q.joinOrSelect(joinOrSelectNodes...).Get()
}

// HasEmployeeInfoByPersonID returns true if the
// given unique index values exist in the database.
// doc: type=EmployeeInfo
func HasEmployeeInfoByPersonID(ctx context.Context, personID string) bool {
	q := queryEmployeeInfos(ctx)
	q = q.Where(op.Equal(node.EmployeeInfo().PersonID(), personID))
	return q.Count(false) == 1
}

// The EmployeeInfosBuilder uses the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type EmployeeInfosBuilder struct {
	builder query.QueryBuilderI
}

func newEmployeeInfoBuilder(ctx context.Context) *EmployeeInfosBuilder {
	b := &EmployeeInfosBuilder{
		builder: db.GetDatabase("goradd").NewBuilder(ctx),
	}
	return b.Join(node.EmployeeInfo())
}

// Load terminates the query builder, performs the query, and returns a slice of EmployeeInfo objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *EmployeeInfosBuilder) Load() (employeeInfos []*EmployeeInfo) {
	results := b.builder.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(EmployeeInfo)
		o.load(item, o, nil, "")
		employeeInfos = append(employeeInfos, o)
	}
	return
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *EmployeeInfosBuilder) LoadI() (employeeInfos []interface{}) {
	results := b.builder.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(EmployeeInfo)
		o.load(item, o, nil, "")
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
func (b *EmployeeInfosBuilder) LoadCursor() employeeInfosCursor {
	cursor := b.builder.LoadCursor()

	return employeeInfosCursor{cursor}
}

type employeeInfosCursor struct {
	query.CursorI
}

// Next returns the current EmployeeInfo object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c employeeInfosCursor) Next() *EmployeeInfo {
	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(EmployeeInfo)
	o.load(row, o, nil, "")
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
func (b *EmployeeInfosBuilder) Get() *EmployeeInfo {
	results := b.Load()
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *EmployeeInfosBuilder) Expand(n query.NodeI) *EmployeeInfosBuilder {
	b.builder.Expand(n)
	return b
}

// Join adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *EmployeeInfosBuilder) Join(n query.NodeI, conditions ...query.NodeI) *EmployeeInfosBuilder {
	var condition query.NodeI
	if len(conditions) > 1 {
		condition = op.And(conditions)
	} else if len(conditions) == 1 {
		condition = conditions[0]
	}
	b.builder.Join(n, condition)
	return b
}

// Where adds a condition to filter what gets selected.
func (b *EmployeeInfosBuilder) Where(c query.NodeI) *EmployeeInfosBuilder {
	b.builder.Condition(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
func (b *EmployeeInfosBuilder) OrderBy(nodes ...query.NodeI) *EmployeeInfosBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *EmployeeInfosBuilder) Limit(maxRowCount int, offset int) *EmployeeInfosBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *EmployeeInfosBuilder) Select(nodes ...query.NodeI) *EmployeeInfosBuilder {
	b.builder.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using GetAlias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *EmployeeInfosBuilder) Alias(name string, n query.NodeI) *EmployeeInfosBuilder {
	b.builder.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *EmployeeInfosBuilder) Distinct() *EmployeeInfosBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *EmployeeInfosBuilder) GroupBy(nodes ...query.NodeI) *EmployeeInfosBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *EmployeeInfosBuilder) Having(node query.NodeI) *EmployeeInfosBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
//
// distinct wll count the number of distinct items, ignoring duplicates.
//
// nodes will select individual fields, and should be accompanied by a GroupBy.
func (b *EmployeeInfosBuilder) Count(distinct bool, nodes ...query.NodeI) uint {
	return b.builder.Count(distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *EmployeeInfosBuilder) Delete() {
	b.builder.Delete()
	broadcast.BulkChange(b.builder.Context(), "goradd", "employee_info")
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *EmployeeInfosBuilder) Subquery() *query.SubqueryNode {
	return b.builder.Subquery()
}

// joinOrSelect is a private helper function for the Load* functions
func (b *EmployeeInfosBuilder) joinOrSelect(nodes ...query.NodeI) *EmployeeInfosBuilder {
	for _, n := range nodes {
		switch n.(type) {
		case query.TableNodeI:
			b.builder.Join(n, nil)
		case *query.ColumnNode:
			b.Select(n)
		}
	}
	return b
}

// CountEmployeeInfoByID queries the database and returns the number of EmployeeInfo objects that
// have id.
// doc: type=EmployeeInfo
func CountEmployeeInfoByID(ctx context.Context, id string) int {
	return int(queryEmployeeInfos(ctx).Where(op.Equal(node.EmployeeInfo().ID(), id)).Count(false))
}

// CountEmployeeInfoByPersonID queries the database and returns the number of EmployeeInfo objects that
// have personID.
// doc: type=EmployeeInfo
func CountEmployeeInfoByPersonID(ctx context.Context, personID string) int {
	if personID == "" {
		return 0
	}
	return int(queryEmployeeInfos(ctx).Where(op.Equal(node.EmployeeInfo().PersonID(), personID)).Count(false))
}

// CountEmployeeInfoByEmployeeNumber queries the database and returns the number of EmployeeInfo objects that
// have employeeNumber.
// doc: type=EmployeeInfo
func CountEmployeeInfoByEmployeeNumber(ctx context.Context, employeeNumber int) int {
	return int(queryEmployeeInfos(ctx).Where(op.Equal(node.EmployeeInfo().EmployeeNumber(), employeeNumber)).Count(false))
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
func (o *employeeInfoBase) load(m map[string]interface{}, objThis *EmployeeInfo, objParent interface{}, parentKey string) {

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
	}

	if v, ok := m["Person"]; ok {
		if objPerson, ok2 := v.(map[string]interface{}); ok2 {
			o.objPerson = new(Person)
			o.objPerson.load(objPerson, o.objPerson, objThis, "EmployeeInfo")
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
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}

	o._restored = true

}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *employeeInfoBase) Save(ctx context.Context) {
	if o._restored {
		o.update(ctx)
	} else {
		o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *employeeInfoBase) update(ctx context.Context) {
	if !o._restored {
		panic("cannot update a record that was not originally read from the database.")
	}

	var modifiedFields map[string]interface{}
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		// TODO: Perform all reads and consistency checks before saves

		// Save loaded Person object to get its new pk and update it here.
		if o.objPerson != nil {
			o.objPerson.Save(ctx)
			id := o.objPerson.PrimaryKey()
			o.SetPersonID(id)
		}

		// Save all modified fields to the database
		modifiedFields = o.getModifiedFields()
		if len(modifiedFields) != 0 {
			d.Update(ctx, "employee_info", modifiedFields, map[string]any{"id": o._originalPK})
		}

	}) // transaction

	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd", "employee_info", o._originalPK, all.SortedKeys(modifiedFields)...)
	}
}

// insert will insert the object into the database. Related items will be saved.
func (o *employeeInfoBase) insert(ctx context.Context) {
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		if o.objPerson != nil {
			o.objPerson.Save(ctx)
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

	}) // transaction

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "employee_info", o.PrimaryKey())
}

// getModifiedFields returns the database columns that have been modified. This
// will determine which specific fields are sent to the database to be changed.
func (o *employeeInfoBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.personIDIsDirty {
		fields["person_id"] = o.personID
	}
	if o.employeeNumberIsDirty {
		fields["employee_number"] = o.employeeNumber
	}
	return
}

// getValidFields returns the fields that have valid data in them.
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
func (o *employeeInfoBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "employee_info", map[string]any{"ID": o.id})
	broadcast.Delete(ctx, "goradd", "employee_info", fmt.Sprint(o.id))
}

// deleteEmployeeInfo deletes the associated record from the database.
func deleteEmployeeInfo(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "employee_info", map[string]any{"ID": pk})
	broadcast.Delete(ctx, "goradd", "employee_info", fmt.Sprint(pk))
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *employeeInfoBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.personIDIsDirty = false
	o.employeeNumberIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database.
func (o *employeeInfoBase) IsDirty() (dirty bool) {
	dirty = o.idIsDirty ||
		o.personIDIsDirty ||
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
		return nil, err
	}
	if err := encoder.Encode(o.idIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.idIsDirty); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.personID); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.personIDIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.personIDIsDirty); err != nil {
		return nil, err
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
			return nil, err
		}
	}

	if err := encoder.Encode(o.employeeNumber); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.employeeNumberIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.employeeNumberIsDirty); err != nil {
		return nil, err
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
			return nil, err
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o._originalPK); err != nil {
		return nil, err
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
		return
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.personID); err != nil {
		return
	}
	if err = dec.Decode(&o.personIDIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.personIDIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&isPtr); err != nil {
		return
	}
	if isPtr {
		if err = dec.Decode(&o.objPerson); err != nil {
			return
		}
	}
	if err = dec.Decode(&o.employeeNumber); err != nil {
		return
	}
	if err = dec.Decode(&o.employeeNumberIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.employeeNumberIsDirty); err != nil {
		return
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
	if err = json.Unmarshal(data, &v); err != nil {
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
					return fmt.Errorf("json field %s cannot be null", k)
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetPersonID(s)
				}

			}

		case "employeeNumber":
			{
				if v == nil {
					return fmt.Errorf("json field %s cannot be null", k)
				}

				if n, ok := v.(int); ok {
					o.SetEmployeeNumber(int(n))
				} else if n, ok := v.(float64); ok {
					o.SetEmployeeNumber(int(n))
				} else {
					return fmt.Errorf("json field %s must be a number", k)
				}
			}

		}
	}
	return
}
