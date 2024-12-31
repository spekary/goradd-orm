// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"unicode/utf8"

	"github.com/goradd/all"
	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/broadcast"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/query"
)

// GiftBase is embedded in a Gift object and provides the ORM access to the database.
// The member variables of the structure are private and should not normally be accessed by the Gift embedder.
// Instead, use the accessor functions.
type giftBase struct {
	number        int
	numberIsValid bool
	numberIsDirty bool

	name        string
	nameIsValid bool
	nameIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]any

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update.
	_restored bool

	_originalPK int
}

// IDs used to access the Gift object fields by name using the Get function.
// doc: type=Gift
const (
	Gift_Number = `Number`
	Gift_Name   = `Name`
)

const GiftNameMaxLength = 50 // The number of runes the column can hold

// Initialize or re-initialize a Gift database object to default values.
func (o *giftBase) Initialize() {

	o.number = 0

	o.numberIsValid = false
	o.numberIsDirty = false

	o.name = ""

	o.nameIsValid = false
	o.nameIsDirty = false

	o._restored = false
}

// PrimaryKey returns the current value of the primary key field.
func (o *giftBase) PrimaryKey() int {
	return o.number
}

// OriginalPrimaryKey returns the value of the primary key that was originally loaded into the object when it was
// read from the database.
func (o *giftBase) OriginalPrimaryKey() int {
	return o._originalPK
}

// Copy copies all valid fields to a new Gift object.
// Forward reference ids will be copied, but reverse and many-many references will not.
// Attached objects will not be included in the copy.
// You will need to manually set the primary key field before saving.
// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *giftBase) Copy() (newObject *Gift) {
	newObject = NewGift()
	if o.numberIsValid {
		newObject.SetNumber(o.number)
	}
	if o.nameIsValid {
		newObject.SetName(o.name)
	}
	return
}

// Number returns the loaded value of Number.
func (o *giftBase) Number() int {
	if o._restored && !o.numberIsValid {
		panic("Number was not selected in the last query and has not been set, and so is not valid")
	}
	return o.number
}

// NumberIsValid returns true if the value was loaded from the database or has been set.
func (o *giftBase) NumberIsValid() bool {
	return o.numberIsValid
}

// SetNumber sets the value of Number in the object, to be saved later using the Save() function.
func (o *giftBase) SetNumber(number int) {
	o.numberIsValid = true

	if o.number != number || !o._restored {
		o.number = number
		o.numberIsDirty = true
	}
}

// Name returns the loaded value of Name.
func (o *giftBase) Name() string {
	if o._restored && !o.nameIsValid {
		panic("Name was not selected in the last query and has not been set, and so is not valid")
	}
	return o.name
}

// NameIsValid returns true if the value was loaded from the database or has been set.
func (o *giftBase) NameIsValid() bool {
	return o.nameIsValid
}

// SetName sets the value of Name in the object, to be saved later using the Save() function.
func (o *giftBase) SetName(name string) {
	o.nameIsValid = true

	if utf8.RuneCountInString(name) > GiftNameMaxLength {
		panic("attempted to set Gift.Name to a value larger than its maximum length in runes")
	}
	if o.name != name || !o._restored {
		o.name = name
		o.nameIsDirty = true
	}
}

// GetAlias returns the alias for the given key.
func (o *giftBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *giftBase) IsNew() bool {
	return !o._restored
}

// LoadGift returns a Gift from the database.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See [GiftsBuilder.Join] and [GiftsBuilder.Select] for more info.
func LoadGift(ctx context.Context, number int, joinOrSelectNodes ...query.NodeI) *Gift {
	return queryGifts(ctx).
		Where(op.Equal(node.Gift().Number(), number)).
		joinOrSelect(joinOrSelectNodes...).
		Get()
}

// HasGift returns true if a Gift with the given primaryKey exists in the database.
// doc: type=Gift
func HasGift(ctx context.Context, number int) bool {
	return queryGifts(ctx).
		Where(op.Equal(node.Gift().Number(), number)).
		Count(false) == 1
}

// LoadGiftByNumber queries for a single Gift object by the given unique index values.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See [GiftsBuilder.Join] and [GiftsBuilder.Select] for more info.
// If you need a more elaborate query, use QueryGifts() to start a query builder.
func LoadGiftByNumber(ctx context.Context, number int, joinOrSelectNodes ...query.NodeI) *Gift {
	q := queryGifts(ctx)
	q = q.Where(op.Equal(node.Gift().Number(), number))
	return q.joinOrSelect(joinOrSelectNodes...).Get()
}

// HasGiftByNumber returns true if the
// given unique index values exist in the database.
// doc: type=Gift
func HasGiftByNumber(ctx context.Context, number int) bool {
	q := queryGifts(ctx)
	q = q.Where(op.Equal(node.Gift().Number(), number))
	return q.Count(false) == 1
}

// The GiftsBuilder uses the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type GiftsBuilder struct {
	builder query.QueryBuilderI
}

func newGiftBuilder(ctx context.Context) *GiftsBuilder {
	b := &GiftsBuilder{
		builder: db.GetDatabase("goradd").NewBuilder(ctx),
	}
	return b.Join(node.Gift())
}

// Load terminates the query builder, performs the query, and returns a slice of Gift objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *GiftsBuilder) Load() (gifts []*Gift) {
	results := b.builder.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Gift)
		o.load(item, o, nil, "")
		gifts = append(gifts, o)
	}
	return
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *GiftsBuilder) LoadI() (gifts []interface{}) {
	results := b.builder.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Gift)
		o.load(item, o, nil, "")
		gifts = append(gifts, o)
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
func (b *GiftsBuilder) LoadCursor() giftsCursor {
	cursor := b.builder.LoadCursor()

	return giftsCursor{cursor}
}

type giftsCursor struct {
	query.CursorI
}

// Next returns the current Gift object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c giftsCursor) Next() *Gift {
	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(Gift)
	o.load(row, o, nil, "")
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
func (b *GiftsBuilder) Get() *Gift {
	results := b.Load()
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *GiftsBuilder) Expand(n query.NodeI) *GiftsBuilder {
	b.builder.Expand(n)
	return b
}

// Join adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *GiftsBuilder) Join(n query.NodeI, conditions ...query.NodeI) *GiftsBuilder {
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
func (b *GiftsBuilder) Where(c query.NodeI) *GiftsBuilder {
	b.builder.Condition(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
func (b *GiftsBuilder) OrderBy(nodes ...query.NodeI) *GiftsBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *GiftsBuilder) Limit(maxRowCount int, offset int) *GiftsBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *GiftsBuilder) Select(nodes ...query.NodeI) *GiftsBuilder {
	b.builder.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using GetAlias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *GiftsBuilder) Alias(name string, n query.NodeI) *GiftsBuilder {
	b.builder.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *GiftsBuilder) Distinct() *GiftsBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *GiftsBuilder) GroupBy(nodes ...query.NodeI) *GiftsBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *GiftsBuilder) Having(node query.NodeI) *GiftsBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
//
// distinct wll count the number of distinct items, ignoring duplicates.
//
// nodes will select individual fields, and should be accompanied by a GroupBy.
func (b *GiftsBuilder) Count(distinct bool, nodes ...query.NodeI) uint {
	return b.builder.Count(distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *GiftsBuilder) Delete() {
	b.builder.Delete()
	broadcast.BulkChange(b.builder.Context(), "goradd", "gift")
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *GiftsBuilder) Subquery() *query.SubqueryNode {
	return b.builder.Subquery()
}

// joinOrSelect is a private helper function for the Load* functions
func (b *GiftsBuilder) joinOrSelect(nodes ...query.NodeI) *GiftsBuilder {
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

// CountGiftByNumber queries the database and returns the number of Gift objects that
// have number.
// doc: type=Gift
func CountGiftByNumber(ctx context.Context, number int) int {
	return int(queryGifts(ctx).Where(op.Equal(node.Gift().Number(), number)).Count(false))
}

// CountGiftByName queries the database and returns the number of Gift objects that
// have name.
// doc: type=Gift
func CountGiftByName(ctx context.Context, name string) int {
	return int(queryGifts(ctx).Where(op.Equal(node.Gift().Name(), name)).Count(false))
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
func (o *giftBase) load(m map[string]interface{}, objThis *Gift, objParent interface{}, parentKey string) {

	if v, ok := m["number"]; ok && v != nil {
		if o.number, ok = v.(int); ok {
			o.numberIsValid = true
			o.numberIsDirty = false

			o._originalPK = o.number

		} else {
			panic("Wrong type found for number.")
		}
	} else {
		o.numberIsValid = false
		o.number = 0
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

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}

	o._restored = true

}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *giftBase) Save(ctx context.Context) {
	if o._restored {
		o.update(ctx)
	} else {
		o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *giftBase) update(ctx context.Context) {
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
			d.Update(ctx, "gift", modifiedFields, map[string]any{"number": o._originalPK})
		}

	}) // transaction

	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd", "gift", o._originalPK, all.SortedKeys(modifiedFields)...)
	}
}

// insert will insert the object into the database. Related items will be saved.
func (o *giftBase) insert(ctx context.Context) {
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		if !o.numberIsValid {
			panic("a value for Number is required, and there is no default value. Call SetNumber() before inserting the record.")
		}

		if !o.nameIsValid {
			panic("a value for Name is required, and there is no default value. Call SetName() before inserting the record.")
		}

		m := o.getValidFields()

		d.Insert(ctx, "gift", m)
		id := o.PrimaryKey()
		o._originalPK = id

	}) // transaction

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "gift", o.PrimaryKey())
}

// getModifiedFields returns the database columns that have been modified. This
// will determine which specific fields are sent to the database to be changed.
func (o *giftBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.numberIsDirty {
		fields["number"] = o.number
	}
	if o.nameIsDirty {
		fields["name"] = o.name
	}
	return
}

// getValidFields returns the fields that have valid data in them.
func (o *giftBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}

	if o.numberIsValid {
		fields["number"] = o.number
	}

	if o.nameIsValid {
		fields["name"] = o.name
	}
	return
}

// Delete deletes the record from the database.
func (o *giftBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "gift", map[string]any{"Number": o.number})
	broadcast.Delete(ctx, "goradd", "gift", fmt.Sprint(o.number))
}

// deleteGift deletes the associated record from the database.
func deleteGift(ctx context.Context, pk int) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "gift", map[string]any{"Number": pk})
	broadcast.Delete(ctx, "goradd", "gift", fmt.Sprint(pk))
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *giftBase) resetDirtyStatus() {
	o.numberIsDirty = false
	o.nameIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database.
func (o *giftBase) IsDirty() (dirty bool) {
	dirty = o.numberIsDirty ||
		o.nameIsDirty

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *giftBase) Get(key string) interface{} {

	switch key {

	case "Number":
		if !o.numberIsValid {
			return nil
		}
		return o.number

	case "Name":
		if !o.nameIsValid {
			return nil
		}
		return o.name

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *giftBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.number); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.numberIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.numberIsDirty); err != nil {
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

// UnmarshalBinary converts a structure that was created with MarshalBinary into a Gift object.
func (o *giftBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var isPtr bool

	_ = isPtr
	if err = dec.Decode(&o.number); err != nil {
		return
	}
	if err = dec.Decode(&o.numberIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.numberIsDirty); err != nil {
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

	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *giftBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *giftBase) MarshalStringMap() map[string]interface{} {
	v := make(map[string]interface{})

	if o.numberIsValid {
		v["number"] = o.number
	}

	if o.nameIsValid {
		v["name"] = o.name
	}

	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the Gift. The Gift can be a
// newly created object, or one loaded from the database.
//
// After unmarshalling, the object is not  saved. You must call Save to insert it into the database
// or update it.
//
// Unmarshalling of sub-objects, as in objects linked via foreign keys, is not currently supported.
//
// The fields it expects are:
//
//	"number" - int
//	"name" - string
func (o *giftBase) UnmarshalJSON(data []byte) (err error) {
	var v map[string]interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	return o.UnmarshalStringMap(v)
}

// UnmarshalStringMap will load the values from the stringmap into the object.
//
// Override this in Gift to modify the json before sending it here.
func (o *giftBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
	for k, v := range m {
		switch k {

		case "number":
			{
				if v == nil {
					return fmt.Errorf("json field %s cannot be null", k)
				}

				if n, ok := v.(int); ok {
					o.SetNumber(int(n))
				} else if n, ok := v.(float64); ok {
					o.SetNumber(int(n))
				} else {
					return fmt.Errorf("json field %s must be a number", k)
				}
			}

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

		}
	}
	return
}
