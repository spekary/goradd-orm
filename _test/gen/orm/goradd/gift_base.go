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

const GiftNumberMax = 2147483647
const GiftNumberMin = -2147483648
const GiftNameMaxLength = 50 // The number of runes the column can hold

// Initialize or re-initialize a Gift database object to default values.
func (o *giftBase) Initialize() {

	o.number = 0

	o.numberIsValid = false
	o.numberIsDirty = false

	o.name = ""

	o.nameIsValid = false
	o.nameIsDirty = false

	o._aliases = nil

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

// SetNumber sets the value of Number in the object, to be saved later in the database using the Save() function.
func (o *giftBase) SetNumber(v int) {
	if o._restored &&
		o.numberIsValid && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.number == v {
		// no change
		return
	}

	o.numberIsValid = true
	o.number = v
	o.numberIsDirty = true
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

// SetName sets the value of Name in the object, to be saved later in the database using the Save() function.
func (o *giftBase) SetName(v string) {
	if utf8.RuneCountInString(v) > GiftNameMaxLength {
		panic("attempted to set Gift.Name to a value larger than its maximum length in runes")
	}
	if o._restored &&
		o.nameIsValid && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.name == v {
		// no change
		return
	}

	o.nameIsValid = true
	o.name = v
	o.nameIsDirty = true
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
// selectNodes lets you provide nodes for selecting specific fields or additional fields from related tables.
// See [GiftsBuilder.Select] for more info.
func LoadGift(ctx context.Context, number int, selectNodes ...query.Node) *Gift {
	return queryGifts(ctx).
		Where(op.Equal(node.Gift().Number(), number)).
		Select(selectNodes...).
		Get()
}

// HasGift returns true if a Gift with the given primaryKey exists in the database.
// doc: type=Gift
func HasGift(ctx context.Context, number int) bool {
	return queryGifts(ctx).
		Where(op.Equal(node.Gift().Number(), number)).
		Count() == 1
}

// LoadGiftByNumber queries for a single Gift object by the given unique index values.
// selectNodes optionally let you provide nodes for joining to other tables or selecting specific fields.
// See [GiftsBuilder.Select].
// If you need a more elaborate query, use QueryGifts() to start a query builder.
func LoadGiftByNumber(ctx context.Context, number int, selectNodes ...query.Node) *Gift {
	q := queryGifts(ctx)
	q = q.Where(op.Equal(node.Gift().Number(), number))
	return q.Select(selectNodes...).Get()
}

// HasGiftByNumber returns true if the
// given unique index values exist in the database.
// doc: type=Gift
func HasGiftByNumber(ctx context.Context, number int) bool {
	q := queryGifts(ctx)
	q = q.Where(op.Equal(node.Gift().Number(), number))
	return q.Count() == 1
}

// The GiftBuilder uses the query.BuilderI interface to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, LoadCursor, Get, Count, or Delete
type GiftBuilder interface {
	// Join(alias string, joinedTable query.Node, condition query.Node) GiftBuilder

	// Where adds a condition to filter what gets selected.
	// Calling Where multiple times will AND the conditions together.
	Where(c query.Node) GiftBuilder

	// OrderBy specifies how the resulting data should be sorted.
	// By default, the given nodes are sorted in ascending order.
	// Add Descending() to the node to specify that it should be sorted in descending order.
	OrderBy(nodes ...query.Sorter) GiftBuilder

	// Limit will return a subset of the data, limited to the offset and number of rows specified.
	// For large data sets and specific types of queries, this can be slow, because it will perform
	// the entire query before computing the limit.
	// You cannot limit a query that has selected a "many" relationship".
	Limit(maxRowCount int, offset int) GiftBuilder

	// Select performs two functions:
	//  - Passing a table type node will join the object or objects from that table to this object.
	//  - Passing a column node will optimize the query to only return the specified fields.
	// Once you select at least one column, you must select all the columns that you want in the result.
	// Some fields, like primary keys, are always selected.
	// If you are using a GroupBy, you must select the fields in the GroupBy.
	Select(nodes ...query.Node) GiftBuilder

	// Calculation adds a calculation described by operation with the name alias.
	// After the query, you can read the data using GetAlias() on the object identified by base.
	Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) GiftBuilder

	// Distinct removes duplicates from the results of the query.
	// Adding a Select() is required.
	Distinct() GiftBuilder

	// GroupBy controls how results are grouped when using aggregate functions with Calculation.
	GroupBy(nodes ...query.Node) GiftBuilder

	// Having does additional filtering on the results of the query after the query is performed.
	Having(node query.Node) GiftBuilder

	// Load terminates the query builder, performs the query, and returns a slice of Gift objects.
	// If there are any errors, nil is returned and the specific error is stored in the context.
	// If no results come back from the query, it will return a non-nil empty slice.
	Load() []*Gift
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
	LoadCursor() giftsCursor

	// Get is a convenience method to return only the first item found in a query.
	// The entire query is performed, so you should generally use this only if you know
	// you are selecting on one or very few items.
	//
	// If an error occurs, or no results are found, a nil is returned.
	// In the case of an error, the error is returned in the context.
	Get() *Gift

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

type giftQueryBuilder struct {
	builder *query.Builder
}

func newGiftBuilder(ctx context.Context) GiftBuilder {
	b := giftQueryBuilder{
		builder: query.NewBuilder(ctx, node.Gift()),
	}
	return &b
}

// Load terminates the query builder, performs the query, and returns a slice of Gift objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *giftQueryBuilder) Load() (gifts []*Gift) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(Gift)
		o.load(item, o)
		gifts = append(gifts, o)
	}
	return
}

// Load terminates the query builder, performs the query, and returns a slice of interfaces.
// This can then satisfy a general interface that loads arrays of objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *giftQueryBuilder) LoadI() (gifts []any) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(Gift)
		o.load(item, o)
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
func (b *giftQueryBuilder) LoadCursor() giftsCursor {
	b.builder.Command = query.BuilderCommandLoadCursor
	database := db.GetDatabase("goradd")
	result := database.BuilderQuery(b.builder)
	if result == nil {
		return giftsCursor{}
	}
	cursor := result.(query.CursorI)

	return giftsCursor{cursor}
}

type giftsCursor struct {
	query.CursorI
}

// Next returns the current Gift object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c giftsCursor) Next() *Gift {
	if c.CursorI == nil {
		return nil
	}

	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(Gift)
	o.load(row, o)
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
//
// If an error occurs, or no results are found, a nil is returned.
// In the case of an error, the error is returned in the context.
func (b *giftQueryBuilder) Get() *Gift {
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
func (b *giftQueryBuilder) Join(alias string, joinedTable query.Node, condition query.Node) GiftBuilder {
    if query.RootNode(n).TableName_() != "gift" {
        panic("you can only join a node that is rooted at node.Gift()")
    }
    // TODO: make sure joinedTable is a table node
	b.builder.Join(alias, joinedTable, condition)
	return b
}
*/

// Where adds a condition to filter what gets selected.
// Calling Where multiple times will AND the conditions together.
func (b *giftQueryBuilder) Where(c query.Node) GiftBuilder {
	b.builder.Where(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
// By default, the given nodes are sorted in ascending order.
// Add Descending() to the node to specify that it should be sorted in descending order.
func (b *giftQueryBuilder) OrderBy(nodes ...query.Sorter) GiftBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified.
// For large data sets and specific types of queries, this can be slow, because it will perform
// the entire query before computing the limit.
// You cannot limit a query that has embedded arrays.
func (b *giftQueryBuilder) Limit(maxRowCount int, offset int) GiftBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select specifies what specific columns will be loaded with data.
// By default, all the columns of the gift table will be queried and loaded.
// If nodes contains columns from the gift table, that will limit the columns queried and loaded to only those columns.
// If related tables are specified, then all the columns from those tables are queried, selected and joined to the result.
// If columns in related tables are specified, then only those columns will be queried and loaded.
// Depending on the query, additional columns may automatically be added to the query. In particular, primary key columns
// will be added in most situations. The exception to this would be in distinct queries, group by queries, or subqueries.
func (b *giftQueryBuilder) Select(nodes ...query.Node) GiftBuilder {
	b.builder.Select(nodes...)
	return b
}

// Calculation adds operation as an aliased value onto base.
// After the query, you can read the data by passing alias to GetAlias on the returned object.
func (b *giftQueryBuilder) Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) GiftBuilder {
	b.builder.Calculation(base, alias, operation)
	return b
}

// Distinct removes duplicates from the results of the query.
// Adding a Select() is usually required.
func (b *giftQueryBuilder) Distinct() GiftBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions with Calculation.
func (b *giftQueryBuilder) GroupBy(nodes ...query.Node) GiftBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query after the query is performed.
func (b *giftQueryBuilder) Having(node query.Node) GiftBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items in the result.
// If you have Select or Calculation columns in the query, it will count NULL results as well.
// To not count NULL values, use Where in the builder with a NotNull operation.
// To count distinct combinations of items, call Distinct() on the builder.
func (b *giftQueryBuilder) Count() int {
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
func (b *giftQueryBuilder)  Subquery() *query.SubqueryNode {
	 return b.builder.Subquery()
}
*/

// CountGiftsByNumber queries the database and returns the number of Gift objects that
// have number.
// doc: type=Gift
func CountGiftsByNumber(ctx context.Context, number int) int {
	return queryGifts(ctx).Where(op.Equal(node.Gift().Number(), number)).Count()
}

// CountGiftsByName queries the database and returns the number of Gift objects that
// have name.
// doc: type=Gift
func CountGiftsByName(ctx context.Context, name string) int {
	return queryGifts(ctx).Where(op.Equal(node.Gift().Name(), name)).Count()
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
func (o *giftBase) load(m map[string]interface{}, objThis *Gift) {

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
		o.numberIsDirty = false
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
		o.nameIsDirty = false
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
func (o *giftBase) Save(ctx context.Context) error {
	if o._restored {
		return o.update(ctx)
	} else {
		return o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *giftBase) update(ctx context.Context) error {
	if !o._restored {
		panic("cannot update a record that was not originally read from the database.")
	}

	var modifiedFields map[string]interface{}
	d := Database()
	err := db.ExecuteTransaction(ctx, d, func() error {

		// Save all modified fields to the database
		modifiedFields = o.getModifiedFields()
		if len(modifiedFields) != 0 {
			err := d.Update(ctx, "gift", "number", o._originalPK, modifiedFields, "", 0)
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
		broadcast.Update(ctx, "goradd", "gift", o._originalPK, all.SortedKeys(modifiedFields)...)
	}

	return nil
}

// insert will insert the object into the database. Related items will be saved.
func (o *giftBase) insert(ctx context.Context) (err error) {
	d := Database()
	err = db.ExecuteTransaction(ctx, d, func() error {

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

		return nil

	}) // transaction

	if err != nil {
		return
	}

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "gift", o.PrimaryKey())
	return
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

// getValidFields returns the fields that have valid data in them in a form ready to send to the database.
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
func (o *giftBase) Delete(ctx context.Context) (err error) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "gift", map[string]any{"Number": o.number})
	return nil
	broadcast.Delete(ctx, "goradd", "gift", fmt.Sprint(o.number))
	return
}

// deleteGift deletes the Gift with primary key pk from the database
// and handles associated records.
func deleteGift(ctx context.Context, pk int) error {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "gift", map[string]any{"Number": pk})
	broadcast.Delete(ctx, "goradd", "gift", fmt.Sprint(pk))
	return nil
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
		return nil, fmt.Errorf("error encoding Gift.number: %w", err)
	}
	if err := encoder.Encode(o.numberIsValid); err != nil {
		return nil, fmt.Errorf("error encoding Gift.numberIsValid: %w", err)
	}
	if err := encoder.Encode(o.numberIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding Gift.numberIsDirty: %w", err)
	}

	if err := encoder.Encode(o.name); err != nil {
		return nil, fmt.Errorf("error encoding Gift.name: %w", err)
	}
	if err := encoder.Encode(o.nameIsValid); err != nil {
		return nil, fmt.Errorf("error encoding Gift.nameIsValid: %w", err)
	}
	if err := encoder.Encode(o.nameIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding Gift.nameIsDirty: %w", err)
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
			return nil, fmt.Errorf("error encoding Gift._aliases: %w", err)
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, fmt.Errorf("error encoding Gift._restored: %w", err)
	}
	if err := encoder.Encode(o._originalPK); err != nil {
		return nil, fmt.Errorf("error encoding Gift._originalPK: %w", err)
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
		return fmt.Errorf("error decoding Gift.number: %w", err)
	}
	if err = dec.Decode(&o.numberIsValid); err != nil {
		return fmt.Errorf("error decoding Gift.numberIsValid: %w", err)
	}
	if err = dec.Decode(&o.numberIsDirty); err != nil {
		return fmt.Errorf("error decoding Gift.numberIsDirty: %w", err)
	}

	if err = dec.Decode(&o.name); err != nil {
		return fmt.Errorf("error decoding Gift.name: %w", err)
	}
	if err = dec.Decode(&o.nameIsValid); err != nil {
		return fmt.Errorf("error decoding Gift.nameIsValid: %w", err)
	}
	if err = dec.Decode(&o.nameIsDirty); err != nil {
		return fmt.Errorf("error decoding Gift.nameIsDirty: %w", err)
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
// Override this in Gift to modify the json before sending it here.
func (o *giftBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
	for k, v := range m {
		switch k {

		case "number":
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
					o.SetNumber(int(n2))
				case int:
					o.SetNumber(n)
				case float64:
					o.SetNumber(int(n))
				default:
					return fmt.Errorf("field %s must be a number", k)
				}
			}

		case "name":
			{
				if v == nil {
					return fmt.Errorf("field %s cannot be null", k)
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
