// Code generated by GoRADD. DO NOT EDIT.

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

// ForwardNullBase is embedded in a ForwardNull object and provides the ORM access to the database.
// The member variables of the structure are private and should not normally be accessed by the ForwardNull embedder.
// Instead, use the accessor functions.
type forwardNullBase struct {
	id        string
	idIsValid bool
	idIsDirty bool

	name        string
	nameIsValid bool
	nameIsDirty bool

	reverseID        string
	reverseIDIsNull  bool
	reverseIDIsValid bool
	reverseIDIsDirty bool
	objReverse       *Reverse

	// Custom aliases, if specified
	_aliases map[string]any

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update.
	_restored bool

	_originalPK string
}

// IDs used to access the ForwardNull object fields by name using the Get function.
// doc: type=ForwardNull
const (
	ForwardNull_ID        = `ID`
	ForwardNull_Name      = `Name`
	ForwardNull_ReverseID = `ReverseID`
	ForwardNull_Reverse   = `Reverse`
)

const ForwardNullNameMaxLength = 100 // The number of runes the column can hold

// Initialize or re-initialize a ForwardNull database object to default values.
// The primary key will get a temporary negative number which will be replaced when the object is saved.
// Multiple calls to Initialize are not guaranteed to create sequential values for the primary key.
func (o *forwardNullBase) Initialize() {

	newObjectPkCounter = newObjectPkCounter - 1
	o.id = fmt.Sprintf("%d", newObjectPkCounter)

	o.idIsValid = false
	o.idIsDirty = false

	o.name = ""

	o.nameIsValid = false
	o.nameIsDirty = false

	o.reverseID = ""

	o.reverseIDIsNull = true
	o.reverseIDIsValid = true
	o.reverseIDIsDirty = true

	o._restored = false
}

// PrimaryKey returns the current value of the primary key field.
func (o *forwardNullBase) PrimaryKey() string {
	return o.id
}

// OriginalPrimaryKey returns the value of the primary key that was originally loaded into the object when it was
// read from the database.
func (o *forwardNullBase) OriginalPrimaryKey() string {
	return o._originalPK
}

// Copy copies all valid fields (except for the primary key) to a new ForwardNull object.
// Forward reference ids will be copied, but reverse and many-many references will not.

// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *forwardNullBase) Copy() (newObject *ForwardNull) {
	newObject = NewForwardNull()
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
func (o *forwardNullBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *forwardNullBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

// Name returns the loaded value of Name.
func (o *forwardNullBase) Name() string {
	if o._restored && !o.nameIsValid {
		panic("Name was not selected in the last query and has not been set, and so is not valid")
	}
	return o.name
}

// NameIsValid returns true if the value was loaded from the database or has been set.
func (o *forwardNullBase) NameIsValid() bool {
	return o.nameIsValid
}

// SetName sets the value of Name in the object, to be saved later using the Save() function.
func (o *forwardNullBase) SetName(name string) {
	o.nameIsValid = true
	if utf8.RuneCountInString(name) > ForwardNullNameMaxLength {
		panic("attempted to set ForwardNull.Name to a value larger than its maximum length in runes")
	}
	if o.name != name || !o._restored {
		o.name = name
		o.nameIsDirty = true
	}
}

// ReverseID returns the loaded value of ReverseID.
func (o *forwardNullBase) ReverseID() string {
	if o._restored && !o.reverseIDIsValid {
		panic("ReverseID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.reverseID
}

// ReverseIDIsValid returns true if the value was loaded from the database or has been set.
func (o *forwardNullBase) ReverseIDIsValid() bool {
	return o.reverseIDIsValid
}

// ReverseIDIsNull returns true if the related database value is null.
func (o *forwardNullBase) ReverseIDIsNull() bool {
	return o.reverseIDIsNull
}

// ReverseID_I returns the loaded value of ReverseID as an interface.
// If the value in the database is NULL, a nil interface is returned.
func (o *forwardNullBase) ReverseID_I() interface{} {
	if o._restored && !o.reverseIDIsValid {
		panic("reverseID was not selected in the last query and has not been set, and so is not valid")
	} else if o.reverseIDIsNull {
		return nil
	}
	return o.reverseID
}

func (o *forwardNullBase) SetReverseID(i interface{}) {
	o.reverseIDIsValid = true
	if i == nil {
		if !o.reverseIDIsNull {
			o.reverseIDIsNull = true
			o.reverseIDIsDirty = true
			o.reverseID = ""
			o.objReverse = nil
		}
	} else {
		v := i.(string)
		if o.reverseIDIsNull ||
			!o._restored ||
			o.reverseID != v {

			o.reverseIDIsNull = false
			o.reverseID = v
			o.reverseIDIsDirty = true
			o.objReverse = nil
		}
	}
}

// Reverse returns the current value of the loaded Reverse, and nil if its not loaded.
func (o *forwardNullBase) Reverse() *Reverse {
	return o.objReverse
}

// LoadReverse returns the related Reverse. If it is not already loaded,
// it will attempt to load it, provided the ReverseID column has been loaded first.
func (o *forwardNullBase) LoadReverse(ctx context.Context) *Reverse {
	if !o.reverseIDIsValid {
		return nil
	}

	if o.objReverse == nil {
		// Load and cache
		o.objReverse = LoadReverse(ctx, o.reverseID)
	}
	return o.objReverse
}

// SetReverse will set the reference to reverse. The referenced object
// will be saved when ForwardNull is saved. Pass nil to break the connection.
func (o *forwardNullBase) SetReverse(objReverse *Reverse) {
	o.reverseIDIsValid = true
	if objReverse == nil {
		if !o.reverseIDIsNull || !o._restored {
			o.reverseIDIsNull = true
			o.reverseIDIsDirty = true
			o.reverseID = ""
			o.objReverse = nil
		}
	} else {
		o.objReverse = objReverse
		if o.reverseIDIsNull || !o._restored || o.reverseID != objReverse.PrimaryKey() {
			o.reverseIDIsNull = false
			o.reverseID = objReverse.PrimaryKey()
			o.reverseIDIsDirty = true
		}
	}
}

// GetAlias returns the alias for the given key.
func (o *forwardNullBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *forwardNullBase) IsNew() bool {
	return !o._restored
}

// LoadForwardNull returns a ForwardNull from the database.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See [ForwardNullsBuilder.Join] and [ForwardNullsBuilder.Select] for more info.
func LoadForwardNull(ctx context.Context, id string, joinOrSelectNodes ...query.NodeI) *ForwardNull {
	return queryForwardNulls(ctx).
		Where(op.Equal(node.ForwardNull().ID(), id)).
		joinOrSelect(joinOrSelectNodes...).
		Get()
}

// HasForwardNull returns true if a ForwardNull with the given primaryKey exists in the database.
// doc: type=ForwardNull
func HasForwardNull(ctx context.Context, id string) bool {
	return queryForwardNulls(ctx).
		Where(op.Equal(node.ForwardNull().ID(), id)).
		Count(false) == 1
}

// The ForwardNullsBuilder uses the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type ForwardNullsBuilder struct {
	builder query.QueryBuilderI
}

func newForwardNullBuilder(ctx context.Context) *ForwardNullsBuilder {
	b := &ForwardNullsBuilder{
		builder: db.GetDatabase("goradd_unit").NewBuilder(ctx),
	}
	return b.Join(node.ForwardNull())
}

// Load terminates the query builder, performs the query, and returns a slice of ForwardNull objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *ForwardNullsBuilder) Load() (forwardNulls []*ForwardNull) {
	results := b.builder.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(ForwardNull)
		o.load(item, o, nil, "")
		forwardNulls = append(forwardNulls, o)
	}
	return
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *ForwardNullsBuilder) LoadI() (forwardNulls []interface{}) {
	results := b.builder.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(ForwardNull)
		o.load(item, o, nil, "")
		forwardNulls = append(forwardNulls, o)
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
func (b *ForwardNullsBuilder) LoadCursor() forwardNullsCursor {
	cursor := b.builder.LoadCursor()

	return forwardNullsCursor{cursor}
}

type forwardNullsCursor struct {
	query.CursorI
}

// Next returns the current ForwardNull object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c forwardNullsCursor) Next() *ForwardNull {
	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(ForwardNull)
	o.load(row, o, nil, "")
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
func (b *ForwardNullsBuilder) Get() *ForwardNull {
	results := b.Load()
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *ForwardNullsBuilder) Expand(n query.NodeI) *ForwardNullsBuilder {
	b.builder.Expand(n)
	return b
}

// Join adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *ForwardNullsBuilder) Join(n query.NodeI, conditions ...query.NodeI) *ForwardNullsBuilder {
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
func (b *ForwardNullsBuilder) Where(c query.NodeI) *ForwardNullsBuilder {
	b.builder.Condition(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
func (b *ForwardNullsBuilder) OrderBy(nodes ...query.NodeI) *ForwardNullsBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *ForwardNullsBuilder) Limit(maxRowCount int, offset int) *ForwardNullsBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *ForwardNullsBuilder) Select(nodes ...query.NodeI) *ForwardNullsBuilder {
	b.builder.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using GetAlias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *ForwardNullsBuilder) Alias(name string, n query.NodeI) *ForwardNullsBuilder {
	b.builder.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *ForwardNullsBuilder) Distinct() *ForwardNullsBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *ForwardNullsBuilder) GroupBy(nodes ...query.NodeI) *ForwardNullsBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *ForwardNullsBuilder) Having(node query.NodeI) *ForwardNullsBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
//
// distinct wll count the number of distinct items, ignoring duplicates.
//
// nodes will select individual fields, and should be accompanied by a GroupBy.
func (b *ForwardNullsBuilder) Count(distinct bool, nodes ...query.NodeI) uint {
	return b.builder.Count(distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *ForwardNullsBuilder) Delete() {
	b.builder.Delete()
	broadcast.BulkChange(b.builder.Context(), "goradd_unit", "forward_null")
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *ForwardNullsBuilder) Subquery() *query.SubqueryNode {
	return b.builder.Subquery()
}

// joinOrSelect is a private helper function for the Load* functions
func (b *ForwardNullsBuilder) joinOrSelect(nodes ...query.NodeI) *ForwardNullsBuilder {
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

// CountForwardNullByID queries the database and returns the number of ForwardNull objects that
// have id.
// doc: type=ForwardNull
func CountForwardNullByID(ctx context.Context, id string) int {
	return int(queryForwardNulls(ctx).Where(op.Equal(node.ForwardNull().ID(), id)).Count(false))
}

// CountForwardNullByName queries the database and returns the number of ForwardNull objects that
// have name.
// doc: type=ForwardNull
func CountForwardNullByName(ctx context.Context, name string) int {
	return int(queryForwardNulls(ctx).Where(op.Equal(node.ForwardNull().Name(), name)).Count(false))
}

// CountForwardNullByReverseID queries the database and returns the number of ForwardNull objects that
// have reverseID.
// doc: type=ForwardNull
func CountForwardNullByReverseID(ctx context.Context, reverseID string) int {
	if reverseID == "" {
		return 0
	}
	return int(queryForwardNulls(ctx).Where(op.Equal(node.ForwardNull().ReverseID(), reverseID)).Count(false))
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
func (o *forwardNullBase) load(m map[string]interface{}, objThis *ForwardNull, objParent interface{}, parentKey string) {

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

	if v, ok := m["reverse_id"]; ok {
		if v == nil {
			o.reverseID = ""
			o.reverseIDIsNull = true
			o.reverseIDIsValid = true
			o.reverseIDIsDirty = false
		} else if o.reverseID, ok = v.(string); ok {
			o.reverseIDIsNull = false
			o.reverseIDIsValid = true
			o.reverseIDIsDirty = false
		} else {
			panic("Wrong type found for reverse_id.")
		}
	} else {
		o.reverseIDIsValid = false
		o.reverseIDIsNull = true
		o.reverseID = ""
	}

	if v, ok := m["Reverse"]; ok {
		if objReverse, ok2 := v.(map[string]interface{}); ok2 {
			o.objReverse = new(Reverse)
			o.objReverse.load(objReverse, o.objReverse, objThis, "ForwardNulls")
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
func (o *forwardNullBase) Save(ctx context.Context) {
	if o._restored {
		o.update(ctx)
	} else {
		o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *forwardNullBase) update(ctx context.Context) {
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
			d.Update(ctx, "forward_null", modifiedFields, map[string]any{"id": o._originalPK})
		}

	}) // transaction

	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd_unit", "forward_null", o._originalPK, all.SortedKeys(modifiedFields)...)
	}
}

// insert will insert the object into the database. Related items will be saved.
func (o *forwardNullBase) insert(ctx context.Context) {
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

		m := o.getValidFields()

		id := d.Insert(ctx, "forward_null", m)
		o.id = id
		o._originalPK = id

	}) // transaction

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd_unit", "forward_null", o.PrimaryKey())
}

// getModifiedFields returns the database columns that have been modified. This
// will determine which specific fields are sent to the database to be changed.
func (o *forwardNullBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.nameIsDirty {
		fields["name"] = o.name
	}
	if o.reverseIDIsDirty {
		if o.reverseIDIsNull {
			fields["reverse_id"] = nil
		} else {
			fields["reverse_id"] = o.reverseID
		}
	}
	return
}

// getValidFields returns the fields that have valid data in them.
func (o *forwardNullBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}

	if o.nameIsValid {
		fields["name"] = o.name
	}

	if o.reverseIDIsValid {
		if o.reverseIDIsNull {
			fields["reverse_id"] = nil
		} else {
			fields["reverse_id"] = o.reverseID
		}
	}
	return
}

// Delete deletes the record from the database.
func (o *forwardNullBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "forward_null", map[string]any{"ID": o.id})
	broadcast.Delete(ctx, "goradd_unit", "forward_null", fmt.Sprint(o.id))
}

// deleteForwardNull deletes the associated record from the database.
func deleteForwardNull(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd_unit")
	d.Delete(ctx, "forward_null", map[string]any{"ID": pk})
	broadcast.Delete(ctx, "goradd_unit", "forward_null", fmt.Sprint(pk))
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *forwardNullBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.nameIsDirty = false
	o.reverseIDIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database.
func (o *forwardNullBase) IsDirty() (dirty bool) {
	dirty = o.idIsDirty ||
		o.nameIsDirty ||
		o.reverseIDIsDirty ||
		(o.objReverse != nil && o.objReverse.IsDirty())

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *forwardNullBase) Get(key string) interface{} {

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
func (o *forwardNullBase) MarshalBinary() ([]byte, error) {
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

	if err := encoder.Encode(o.name); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.nameIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.nameIsDirty); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.reverseID); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.reverseIDIsNull); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.reverseIDIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.reverseIDIsDirty); err != nil {
		return nil, err
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
			return nil, err
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

// UnmarshalBinary converts a structure that was created with MarshalBinary into a ForwardNull object.
func (o *forwardNullBase) UnmarshalBinary(data []byte) (err error) {

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

	if err = dec.Decode(&o.name); err != nil {
		return
	}
	if err = dec.Decode(&o.nameIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.nameIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.reverseID); err != nil {
		return
	}
	if err = dec.Decode(&o.reverseIDIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.reverseIDIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.reverseIDIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&isPtr); err != nil {
		return
	}
	if isPtr {
		if err = dec.Decode(&o.objReverse); err != nil {
			return
		}
	}
	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *forwardNullBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *forwardNullBase) MarshalStringMap() map[string]interface{} {
	v := make(map[string]interface{})
	if o.idIsValid {
		v["id"] = o.id
	}

	if o.nameIsValid {
		v["name"] = o.name
	}

	if o.reverseIDIsValid {
		if o.reverseIDIsNull {
			v["reverseID"] = nil
		} else {
			v["reverseID"] = o.reverseID
		}
	}

	if val := o.Reverse(); val != nil {
		v["reverse"] = val.MarshalStringMap()
	}
	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the ForwardNull. The ForwardNull can be a
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
//	"reverseID" - string, nullable
func (o *forwardNullBase) UnmarshalJSON(data []byte) (err error) {
	var v map[string]interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	return o.UnmarshalStringMap(v)
}

// UnmarshalStringMap will load the values from the stringmap into the object.
//
// Override this in ForwardNull to modify the json before sending it here.
func (o *forwardNullBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
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
					o.SetReverseID(v)
					continue
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