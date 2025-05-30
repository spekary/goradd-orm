// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"unicode/utf8"

	"github.com/goradd/anyutil"
	"github.com/goradd/orm/_test/gen/orm/goradd_unit/node"
	"github.com/goradd/orm/pkg/broadcast"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/query"
)

// LeafUnBase is embedded in a LeafUn object and provides the ORM access to the database.
// The member variables of the structure are private and should not normally be accessed by the LeafUn embedder.
// Instead, use the accessor functions.
type leafUnBase struct {
	id               string
	idIsLoaded       bool
	idIsDirty        bool
	name             string
	nameIsLoaded     bool
	nameIsDirty      bool
	rootUnID         string
	rootUnIDIsNull   bool
	rootUnIDIsLoaded bool
	rootUnIDIsDirty  bool
	objRootUn        *RootUn

	// Custom aliases, if specified
	_aliases map[string]any

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update.
	_restored bool

	_originalPK string
}

// IDs used to access the LeafUn object fields by name using the Get function.
// doc: type=LeafUn
const (
	LeafUn_ID       = `ID`
	LeafUn_Name     = `Name`
	LeafUn_RootUnID = `RootUnID`
	LeafUn_RootUn   = `RootUn`
)

const LeafUnIDMaxLength = 32    // The number of runes the column can hold
const LeafUnNameMaxLength = 100 // The number of runes the column can hold

// Initialize or re-initialize a LeafUn database object to default values.
// The primary key will get a temporary negative number which will be replaced when the object is saved.
// Multiple calls to Initialize are not guaranteed to create sequential values for the primary key.
func (o *leafUnBase) Initialize() {
	o.id = db.TemporaryPrimaryKey()
	o.idIsLoaded = true
	o.idIsDirty = false

	o.name = ""
	o.nameIsLoaded = false
	o.nameIsDirty = false

	o.rootUnID = ""
	o.rootUnIDIsNull = true
	o.rootUnIDIsLoaded = false
	o.rootUnIDIsDirty = false

	o._aliases = nil
	o._restored = false
}

// PrimaryKey returns the current value of the primary key field.
func (o *leafUnBase) PrimaryKey() string {
	return o.id
}

// OriginalPrimaryKey returns the value of the primary key that was originally loaded into the object when it was
// read from the database.
func (o *leafUnBase) OriginalPrimaryKey() string {
	return o._originalPK
}

// Copy copies most fields to a new LeafUn object.
// Forward reference ids will be copied, but reverse and many-many references will not.
// Attached objects will not be included in the copy.
// Automatically generated fields will not be included in the copy.
// The primary key field will not be copied, since it is normally auto-generated.
// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *leafUnBase) Copy() (newObject *LeafUn) {
	newObject = NewLeafUn()
	if o.idIsLoaded {
		newObject.SetID(o.id)
	}
	if o.nameIsLoaded {
		newObject.SetName(o.name)
	}
	if o.rootUnIDIsLoaded {
		newObject.SetRootUnID(o.rootUnID)
	}
	return
}

// ID returns the value of ID.
func (o *leafUnBase) ID() string {
	if o._restored && !o.idIsLoaded {
		panic("ID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.id
}

// IDIsLoaded returns true if the value was loaded from the database or has been set.
func (o *leafUnBase) IDIsLoaded() bool {
	return o.idIsLoaded
}

// SetID sets the value of ID in the object, to be saved later in the database using the Save() function.
// Normally you will not need to call this function, since the ID value is automatically generated by the
// database driver. Exceptions might include importing data to a new database, or correcting primary key conflicts when
// merging data.
// You cannot change a primary key for a record that has been written to the database. While SQL databases will
// allow it, NoSql databases will not. Save a copy and delete this one instead.
func (o *leafUnBase) SetID(v string) {
	if o._restored {
		panic("error: Do not change a primary key for a record that has been saved. Instead, save a copy and delete the original.")
	}
	if utf8.RuneCountInString(v) > LeafUnIDMaxLength {
		panic("attempted to set LeafUn.ID to a value larger than its maximum length in runes")
	}

	o.idIsLoaded = true
	o.id = v
	o.idIsDirty = true
}

// Name returns the value of Name.
func (o *leafUnBase) Name() string {
	if o._restored && !o.nameIsLoaded {
		panic("Name was not selected in the last query and has not been set, and so is not valid")
	}
	return o.name
}

// NameIsLoaded returns true if the value was loaded from the database or has been set.
func (o *leafUnBase) NameIsLoaded() bool {
	return o.nameIsLoaded
}

// SetName sets the value of Name in the object, to be saved later in the database using the Save() function.
func (o *leafUnBase) SetName(v string) {
	if utf8.RuneCountInString(v) > LeafUnNameMaxLength {
		panic("attempted to set LeafUn.Name to a value larger than its maximum length in runes")
	}
	if o._restored &&
		o.nameIsLoaded && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.name == v {
		// no change
		return
	}

	o.nameIsLoaded = true
	o.name = v
	o.nameIsDirty = true
}

// RootUnID returns the value of RootUnID.
func (o *leafUnBase) RootUnID() string {
	if o._restored && !o.rootUnIDIsLoaded {
		panic("RootUnID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.rootUnID
}

// RootUnIDIsLoaded returns true if the value was loaded from the database or has been set.
func (o *leafUnBase) RootUnIDIsLoaded() bool {
	return o.rootUnIDIsLoaded
}

// RootUnIDIsNull returns true if the related database value is null.
func (o *leafUnBase) RootUnIDIsNull() bool {
	return o.rootUnIDIsNull
}

// SetRootUnID sets the value of RootUnID in the object, to be saved later in the database using the Save() function.
func (o *leafUnBase) SetRootUnID(v string) {
	if o._restored &&
		o.rootUnIDIsLoaded && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		!o.rootUnIDIsNull && // if the db value is null, force a set of value
		o.rootUnID == v {
		// no change
		return
	}

	o.rootUnIDIsLoaded = true
	o.rootUnID = v
	o.rootUnIDIsDirty = true
	o.rootUnIDIsNull = false
	if o.objRootUn != nil &&
		o.rootUnID != o.objRootUn.PrimaryKey() {
		o.objRootUn = nil
	}
}

// SetRootUnIDToNull() will set the root_un_id value in the database to NULL.
// RootUnID() will return the column's default value after this.
func (o *leafUnBase) SetRootUnIDToNull() {
	if !o.rootUnIDIsLoaded || !o.rootUnIDIsNull {
		// If we know it is null in the database, don't save it
		o.rootUnIDIsDirty = true
	}
	o.rootUnIDIsLoaded = true
	o.rootUnIDIsNull = true
	o.rootUnID = ""
	o.objRootUn = nil
}

// RootUn returns the current value of the loaded RootUn, and nil if its not loaded.
func (o *leafUnBase) RootUn() *RootUn {
	return o.objRootUn
}

// LoadRootUn returns the related RootUn. If it is not already loaded,
// it will attempt to load it, provided the RootUnID column has been loaded first.
func (o *leafUnBase) LoadRootUn(ctx context.Context) (*RootUn, error) {
	var err error

	if o.objRootUn == nil {
		if !o.rootUnIDIsLoaded {
			panic("RootUnID must be selected in the previous query")
		}
		// Load and cache
		o.objRootUn, err = LoadRootUn(ctx, o.rootUnID)
	}
	return o.objRootUn, err
}

// SetRootUn will set the reference to rootUn. The referenced object
// will be saved when LeafUn is saved. Pass nil to break the connection.
func (o *leafUnBase) SetRootUn(objRootUn *RootUn) {
	o.rootUnIDIsLoaded = true
	if objRootUn == nil {
		if !o.rootUnIDIsNull || !o._restored {
			o.rootUnIDIsNull = true
			o.rootUnIDIsDirty = true
			o.rootUnID = ""
			o.objRootUn = nil
		}
	} else {
		o.objRootUn = objRootUn
		if o.rootUnIDIsNull || !o._restored || o.rootUnID != objRootUn.PrimaryKey() {
			o.rootUnIDIsNull = false
			o.rootUnID = objRootUn.PrimaryKey()
			o.rootUnIDIsDirty = true
		}
	}
}

// GetAlias returns the value for the Alias node aliasKey that was returned in the most
// recent query.
func (o *leafUnBase) GetAlias(aliasKey string) query.AliasValue {
	if a, ok := o._aliases[aliasKey]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + aliasKey + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *leafUnBase) IsNew() bool {
	return !o._restored
}

// LoadLeafUn returns a LeafUn from the database.
// selectNodes lets you provide nodes for selecting specific fields or additional fields from related tables.
// See [LeafUnsBuilder.Select] for more info.
func LoadLeafUn(ctx context.Context, id string, selectNodes ...query.Node) (*LeafUn, error) {
	return queryLeafUns(ctx).
		Where(op.Equal(node.LeafUn().ID(), id)).
		Select(selectNodes...).
		Get()
}

// HasLeafUn returns true if a LeafUn with the given primary key exists in the database.
// doc: type=LeafUn
func HasLeafUn(ctx context.Context, id string) (bool, error) {
	v, err := queryLeafUns(ctx).
		Where(op.Equal(node.LeafUn().ID(), id)).
		Count()
	return v > 0, err
}

// LoadLeafUnByRootUnID queries for a single LeafUn object by the given unique index values.
// selectNodes optionally let you provide nodes for joining to other tables or selecting specific fields.
// See [LeafUnsBuilder.Select].
// If you need a more elaborate query, use QueryLeafUns() to start a query builder.
func LoadLeafUnByRootUnID(ctx context.Context, rootUnID interface{}, selectNodes ...query.Node) (*LeafUn, error) {
	q := queryLeafUns(ctx)
	if rootUnID == nil {
		q = q.Where(op.IsNull(node.LeafUn().RootUnID()))
	} else {
		q = q.Where(op.Equal(node.LeafUn().RootUnID(), rootUnID))
	}
	return q.Select(selectNodes...).Get()
}

// HasLeafUnByRootUnID returns true if the
// given unique index values exist in the database.
// doc: type=LeafUn
func HasLeafUnByRootUnID(ctx context.Context, rootUnID interface{}) (bool, error) {
	q := queryLeafUns(ctx)
	if rootUnID == nil {
		q = q.Where(op.IsNull(node.LeafUn().RootUnID()))
	} else {
		q = q.Where(op.Equal(node.LeafUn().RootUnID(), rootUnID))
	}
	v, err := q.Count()
	return v > 0, err
}

// The LeafUnBuilder uses a builder pattern to create a query on the database.
// Create a LeafUnBuilder by calling QueryLeafUns, which will select all
// the LeafUn object in the database. Then filter and arrange those objects
// by calling Where, Select, etc.
// End a query by calling either Load, LoadI, LoadCursor, Get, or Count.
// A LeafUnBuilder stores the context it will use to perform the query, and thus is
// meant to be a short-lived object. You should not save it for later use.
type LeafUnBuilder struct {
	builder *query.Builder
	ctx     context.Context
}

func newLeafUnBuilder(ctx context.Context) *LeafUnBuilder {
	b := LeafUnBuilder{
		builder: query.NewBuilder(node.LeafUn()),
		ctx:     ctx,
	}
	return &b
}

// Load terminates the query builder, performs the query, and returns a slice of LeafUn objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *LeafUnBuilder) Load() (leafUns []*LeafUn, err error) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd_unit")
	var results any

	ctx := b.ctx
	results, err = database.BuilderQuery(ctx, b.builder)
	if results == nil || err != nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(LeafUn)
		o.unpack(item, o)
		leafUns = append(leafUns, o)
	}
	return
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces.
// This can then satisfy a variety of interfaces that load arrays of objects, including KeyLabeler.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *LeafUnBuilder) LoadI() (leafUns []query.OrmObj, err error) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd_unit")
	var results any

	ctx := b.ctx
	results, err = database.BuilderQuery(ctx, b.builder)
	if results == nil || err != nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(LeafUn)
		o.unpack(item, o)
		leafUns = append(leafUns, o)
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
func (b *LeafUnBuilder) LoadCursor() (leafUnsCursor, error) {
	b.builder.Command = query.BuilderCommandLoadCursor
	database := db.GetDatabase("goradd_unit")
	result, err := database.BuilderQuery(b.ctx, b.builder)
	var cursor query.CursorI
	if result != nil {
		cursor = result.(query.CursorI)
	}
	return leafUnsCursor{cursor}, err
}

type leafUnsCursor struct {
	query.CursorI
}

// Next returns the current LeafUn object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c leafUnsCursor) Next() (*LeafUn, error) {
	if c.CursorI == nil {
		return nil, nil
	}

	row, err := c.CursorI.Next()
	if row == nil || err != nil {
		return nil, err
	}
	o := new(LeafUn)
	o.unpack(row, o)
	return o, nil
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
// If an error occurs, or no results are found, a nil is returned.
func (b *LeafUnBuilder) Get() (*LeafUn, error) {
	results, err := b.Load()
	if err != nil || len(results) == 0 {
		return nil, err
	}
	return results[0], nil
}

// Where adds a condition to filter what gets selected.
// Calling Where multiple times will AND the conditions together.
func (b *LeafUnBuilder) Where(c query.Node) *LeafUnBuilder {
	b.builder.Where(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
// By default, the given nodes are sorted in ascending order.
// Add Descending() to the node to specify that it should be sorted in descending order.
func (b *LeafUnBuilder) OrderBy(nodes ...query.Sorter) *LeafUnBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified.
// For large data sets and specific types of queries, this can be slow, because it will perform
// the entire query before computing the limit.
// You cannot limit a query that has embedded arrays.
func (b *LeafUnBuilder) Limit(maxRowCount int, offset int) *LeafUnBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select specifies what specific columns will be loaded with data.
// By default, all the columns of the leaf_un table will be queried and loaded.
// If nodes contains columns from the leaf_un table, that will limit the columns queried and loaded to only those columns.
// If related tables are specified, then all the columns from those tables are queried, selected and joined to the result.
// If columns in related tables are specified, then only those columns will be queried and loaded.
// Depending on the query, additional columns may automatically be added to the query. In particular, primary key columns
// will be added in most situations. The exception to this would be in distinct queries, group by queries, or subqueries.
func (b *LeafUnBuilder) Select(nodes ...query.Node) *LeafUnBuilder {
	b.builder.Select(nodes...)
	return b
}

// Calculation adds operation as an aliased value onto base.
// After the query, you can read the data by passing alias to GetAlias on the returned object.
func (b *LeafUnBuilder) Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) *LeafUnBuilder {
	b.builder.Calculation(base, alias, operation)
	return b
}

// Distinct removes duplicates from the results of the query.
// Adding a Select() is usually required.
func (b *LeafUnBuilder) Distinct() *LeafUnBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions with Calculation.
func (b *LeafUnBuilder) GroupBy(nodes ...query.Node) *LeafUnBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query after the query is performed.
func (b *LeafUnBuilder) Having(node query.Node) *LeafUnBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items in the result.
// If you have Select or Calculation columns in the query, it will count NULL results as well.
// To not count NULL values, use Where in the builder with a NotNull operation.
// To count distinct combinations of items, call Distinct() on the builder.
func (b *LeafUnBuilder) Count() (int, error) {
	b.builder.Command = query.BuilderCommandCount
	database := db.GetDatabase("goradd_unit")

	ctx := b.ctx
	results, err := database.BuilderQuery(ctx, b.builder)
	if results == nil || err != nil {
		return 0, err
	}
	return results.(int), nil
}

// CountLeafUns returns the total number of items in the leaf_un table.
func CountLeafUns(ctx context.Context) (int, error) {
	return QueryLeafUns(ctx).Count()
}

// CountLeafUnsByRootUnID queries the database and returns the number of LeafUn objects that
// have rootUnID.
// doc: type=LeafUn
func CountLeafUnsByRootUnID(ctx context.Context, rootUnID string) (int, error) {
	v_rootUnID := rootUnID
	return QueryLeafUns(ctx).
		Where(op.Equal(node.LeafUn().RootUnID(), v_rootUnID)).
		Count()
}

// unpack recursively transforms data coming from the database into ORM objects.
func (o *leafUnBase) unpack(m map[string]interface{}, objThis *LeafUn) {

	if v, ok := m["id"]; ok && v != nil {
		if o.id, ok = v.(string); ok {
			o.idIsLoaded = true
			o.idIsDirty = false

			o._originalPK = o.id

		} else {
			panic("Wrong type found for id.")
		}
	} else {
		o.idIsLoaded = false
		o.id = ""
		o.idIsDirty = false
	}

	if v, ok := m["name"]; ok && v != nil {
		if o.name, ok = v.(string); ok {
			o.nameIsLoaded = true
			o.nameIsDirty = false

		} else {
			panic("Wrong type found for name.")
		}
	} else {
		o.nameIsLoaded = false
		o.name = ""
		o.nameIsDirty = false
	}

	if v, ok := m["root_un_id"]; ok {
		if v == nil {
			o.rootUnID = ""
			o.rootUnIDIsNull = true
			o.rootUnIDIsLoaded = true
			o.rootUnIDIsDirty = false
		} else if o.rootUnID, ok = v.(string); ok {
			o.rootUnIDIsNull = false
			o.rootUnIDIsLoaded = true
			o.rootUnIDIsDirty = false
		} else {
			panic("Wrong type found for root_un_id.")
		}
	} else {
		o.rootUnIDIsLoaded = false
		o.rootUnIDIsNull = true
		o.rootUnID = ""
		o.rootUnIDIsDirty = false
	}

	if v, ok := m["RootUn"]; ok {
		if objRootUn, ok2 := v.(map[string]any); ok2 {
			o.objRootUn = new(RootUn)
			o.objRootUn.unpack(objRootUn, o.objRootUn)
			o.rootUnIDIsLoaded = true
			o.rootUnIDIsDirty = false
		} else {
			panic("Wrong type found for RootUn object.")
		}
	} else {
		o.objRootUn = nil
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = v.(map[string]any)
	}

	o._restored = true

}

// save will update or insert the object, depending on the state of the object.
func (o *leafUnBase) save(ctx context.Context) error {
	if o._restored {
		return o.update(ctx)
	} else {
		return o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
// If the table has auto-generated values, those will be updated automatically.
func (o *leafUnBase) update(ctx context.Context) error {
	if !o._restored {
		panic("cannot update a record that was not originally read from the database.")
	}
	if !o.IsDirty() {
		return nil // nothing to save
	}

	var modifiedFields map[string]interface{}

	d := Database()
	err := db.ExecuteTransaction(ctx, d, func() error {

		// Save loaded RootUn object to get its new pk and update it here.
		if o.objRootUn != nil {
			if err := o.objRootUn.Save(ctx); err != nil {
				return err
			}
			o.SetRootUnID(o.objRootUn.PrimaryKey())
		}

		if o.rootUnIDIsDirty &&
			!o.rootUnIDIsNull {
			if obj, err := LoadLeafUnByRootUnID(ctx, o.rootUnID); err != nil {
				return err
			} else if obj != nil {
				return db.NewUniqueValueError("leaf_un", map[string]any{"root_un_id": o.rootUnID}, nil)
			}
		}

		modifiedFields = getLeafUnUpdateFields(o)
		if len(modifiedFields) != 0 {
			var err2 error

			_, err2 = d.Update(ctx, "leaf_un", "id", o._originalPK, modifiedFields, "", 0)
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
		broadcast.Update(ctx, "goradd_unit", "leaf_un", o._originalPK, anyutil.SortedKeys(modifiedFields)...)
	}

	return nil
}

// insert will insert the object into the database. Related items will be saved.
func (o *leafUnBase) insert(ctx context.Context) (err error) {
	var insertFields map[string]interface{}
	d := Database()
	err = db.ExecuteTransaction(ctx, d, func() error {

		// Save loaded RootUn object to get its new pk and update it here.
		if o.objRootUn != nil {
			if err := o.objRootUn.Save(ctx); err != nil {
				return err
			}
			o.SetRootUnID(o.objRootUn.PrimaryKey())
		}

		if !o.nameIsLoaded {
			panic("a value for Name is required, and there is no default value. Call SetName() before inserting the record.")
		}
		if o.rootUnIDIsDirty &&
			!o.rootUnIDIsNull {
			if obj, err := LoadLeafUnByRootUnID(ctx, o.rootUnID); err != nil {
				return err
			} else if obj != nil {
				return db.NewUniqueValueError("leaf_un", map[string]any{"root_un_id": o.rootUnID}, nil)
			}
		}
		insertFields = getLeafUnInsertFields(o)
		var newPK string
		newPK, err = d.Insert(ctx, "leaf_un", "id", insertFields)
		if err != nil {
			return err
		}
		o.id = newPK
		o._originalPK = newPK
		o.idIsLoaded = true

		return nil

	}) // transaction

	if err != nil {
		return
	}

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd_unit", "leaf_un", o.PrimaryKey())
	return
}

// getUpdateFields returns the database columns that will be sent to the update process.
// This will include timestamp fields only if some other column has changed.
func (o *leafUnBase) getUpdateFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.nameIsDirty {
		fields["name"] = o.name
	}
	if o.rootUnIDIsDirty {
		if o.rootUnIDIsNull {
			fields["root_un_id"] = nil
		} else {
			fields["root_un_id"] = o.rootUnID
		}
	}
	return
}

// getInsertFields returns the fields that will be specified in an insert operation.
// Optional fields that have not been set and have no default will be returned as nil.
// NoSql databases should interpret this as no value. Sql databases should interpret this as
// explicitly setting a NULL value, which would override any database specific default value.
// Auto-generated fields will be returned with their generated values, except AutoPK fields, which are generated by the
// database driver and updated after the insert.
func (o *leafUnBase) getInsertFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}

	fields["name"] = o.name
	if o.rootUnIDIsNull {
		fields["root_un_id"] = nil
	} else {
		fields["root_un_id"] = o.rootUnID
	}
	return
}

// Delete deletes the record from the database.
func (o *leafUnBase) Delete(ctx context.Context) (err error) {
	if o == nil {
		return // allow deleting of a nil object to be a noop
	}
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	err = d.Delete(ctx, "leaf_un", "id", o.id, "", 0)
	if err != nil {
		return err
	}
	broadcast.Delete(ctx, "goradd_unit", "leaf_un", fmt.Sprint(o.id))
	return
}

// deleteLeafUn deletes the LeafUn with primary key pk from the database
// and handles associated records.
func deleteLeafUn(ctx context.Context, pk string) error {
	d := db.GetDatabase("goradd_unit")
	err := d.Delete(ctx, "leaf_un", "id", pk, "", 0)
	if err != nil {
		return err
	}
	broadcast.Delete(ctx, "goradd_unit", "leaf_un", fmt.Sprint(pk))
	return err
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *leafUnBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.nameIsDirty = false
	o.rootUnIDIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database or created.
// However, a new object that has a column with a default value will be automatically marked as dirty upon creation.
func (o *leafUnBase) IsDirty() (dirty bool) {
	dirty = o.idIsDirty ||
		o.nameIsDirty ||
		o.rootUnIDIsDirty ||
		(o.objRootUn != nil && o.objRootUn.IsDirty())

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil.
// Get can be used to retrieve a value by using the Identifier of a node.
func (o *leafUnBase) Get(key string) interface{} {

	switch key {

	case "ID":
		if !o.idIsLoaded {
			return nil
		}
		return o.id

	case "Name":
		if !o.nameIsLoaded {
			return nil
		}
		return o.name

	case "RootUnID":
		if !o.rootUnIDIsLoaded {
			return nil
		}
		return o.rootUnID

	case "RootUn":
		return o.RootUn()

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *leafUnBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := o.encodeTo(enc); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (o *leafUnBase) encodeTo(enc db.Encoder) error {

	if err := enc.Encode(o.id); err != nil {
		return fmt.Errorf("error encoding LeafUn.id: %w", err)
	}
	if err := enc.Encode(o.idIsLoaded); err != nil {
		return fmt.Errorf("error encoding LeafUn.idIsLoaded: %w", err)
	}
	if err := enc.Encode(o.idIsDirty); err != nil {
		return fmt.Errorf("error encoding LeafUn.idIsDirty: %w", err)
	}

	if err := enc.Encode(o.name); err != nil {
		return fmt.Errorf("error encoding LeafUn.name: %w", err)
	}
	if err := enc.Encode(o.nameIsLoaded); err != nil {
		return fmt.Errorf("error encoding LeafUn.nameIsLoaded: %w", err)
	}
	if err := enc.Encode(o.nameIsDirty); err != nil {
		return fmt.Errorf("error encoding LeafUn.nameIsDirty: %w", err)
	}

	if err := enc.Encode(o.rootUnID); err != nil {
		return fmt.Errorf("error encoding LeafUn.rootUnID: %w", err)
	}
	if err := enc.Encode(o.rootUnIDIsNull); err != nil {
		return fmt.Errorf("error encoding LeafUn.rootUnIDIsNull: %w", err)
	}
	if err := enc.Encode(o.rootUnIDIsLoaded); err != nil {
		return fmt.Errorf("error encoding LeafUn.rootUnIDIsLoaded: %w", err)
	}
	if err := enc.Encode(o.rootUnIDIsDirty); err != nil {
		return fmt.Errorf("error encoding LeafUn.rootUnIDIsDirty: %w", err)
	}

	if o.objRootUn == nil {
		if err := enc.Encode(false); err != nil {
			return err
		}
	} else {
		if err := enc.Encode(true); err != nil {
			return err
		}
		if err := enc.Encode(o.objRootUn); err != nil {
			return fmt.Errorf("error encoding LeafUn.objRootUn: %w", err)
		}
	}

	if o._aliases == nil {
		if err := enc.Encode(false); err != nil {
			return err
		}
	} else {
		if err := enc.Encode(true); err != nil {
			return err
		}
		if err := enc.Encode(o._aliases); err != nil {
			return fmt.Errorf("error encoding LeafUn._aliases: %w", err)
		}
	}

	if err := enc.Encode(o._restored); err != nil {
		return fmt.Errorf("error encoding LeafUn._restored: %w", err)
	}
	if err := enc.Encode(o._originalPK); err != nil {
		return fmt.Errorf("error encoding LeafUn._originalPK: %w", err)
	}
	return nil
}

// UnmarshalBinary converts a structure that was created with MarshalBinary into a LeafUn object.
func (o *leafUnBase) UnmarshalBinary(data []byte) (err error) {
	buf := bytes.NewReader(data)
	dec := gob.NewDecoder(buf)
	return o.decodeFrom(dec)
}

func (o *leafUnBase) decodeFrom(dec db.Decoder) (err error) {
	var isPtr bool

	_ = isPtr
	if err = dec.Decode(&o.id); err != nil {
		return fmt.Errorf("error decoding LeafUn.id: %w", err)
	}
	if err = dec.Decode(&o.idIsLoaded); err != nil {
		return fmt.Errorf("error decoding LeafUn.idIsLoaded: %w", err)
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return fmt.Errorf("error decoding LeafUn.idIsDirty: %w", err)
	}

	if err = dec.Decode(&o.name); err != nil {
		return fmt.Errorf("error decoding LeafUn.name: %w", err)
	}
	if err = dec.Decode(&o.nameIsLoaded); err != nil {
		return fmt.Errorf("error decoding LeafUn.nameIsLoaded: %w", err)
	}
	if err = dec.Decode(&o.nameIsDirty); err != nil {
		return fmt.Errorf("error decoding LeafUn.nameIsDirty: %w", err)
	}

	if err = dec.Decode(&o.rootUnID); err != nil {
		return fmt.Errorf("error decoding LeafUn.rootUnID: %w", err)
	}
	if err = dec.Decode(&o.rootUnIDIsNull); err != nil {
		return fmt.Errorf("error decoding LeafUn.rootUnIDIsNull: %w", err)
	}
	if err = dec.Decode(&o.rootUnIDIsLoaded); err != nil {
		return fmt.Errorf("error decoding LeafUn.rootUnIDIsLoaded: %w", err)
	}
	if err = dec.Decode(&o.rootUnIDIsDirty); err != nil {
		return fmt.Errorf("error decoding LeafUn.rootUnIDIsDirty: %w", err)
	}

	if err = dec.Decode(&isPtr); err != nil {
		return fmt.Errorf("error decoding LeafUn.objRootUn isPtr: %w", err)
	}
	if isPtr {
		if err = dec.Decode(&o.objRootUn); err != nil {
			return fmt.Errorf("error decoding LeafUn.objRootUn: %w", err)
		}
	}
	if err = dec.Decode(&isPtr); err != nil {
		return fmt.Errorf("error decoding LeafUn._aliases isPtr: %w", err)
	}
	if isPtr {
		if err = dec.Decode(&o._aliases); err != nil {
			return fmt.Errorf("error decoding LeafUn._aliases: %w", err)
		}
	}

	if err = dec.Decode(&o._restored); err != nil {
		return fmt.Errorf("error decoding LeafUn._restored: %w", err)
	}
	if err = dec.Decode(&o._originalPK); err != nil {
		return fmt.Errorf("error decoding LeafUn._originalPK: %w", err)
	}
	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *leafUnBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *leafUnBase) MarshalStringMap() map[string]interface{} {
	v := make(map[string]interface{})

	if o.idIsLoaded {
		v["id"] = o.id
	}

	if o.nameIsLoaded {
		v["name"] = o.name
	}

	if val := o.objRootUn; val != nil {
		v["rootUn"] = val.MarshalStringMap()
	} else if o.rootUnIDIsLoaded {
		if o.rootUnIDIsNull {
			v["rootUnID"] = nil
		} else {
			v["rootUnID"] = o.rootUnID
		}
	}

	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the LeafUn. The LeafUn can be a
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
//	"rootUnID" - string, nullable
func (o *leafUnBase) UnmarshalJSON(data []byte) (err error) {
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
// Override this in LeafUn to modify the json before sending it here.
func (o *leafUnBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
	for k, v := range m {
		switch k {

		case "id":
			{
				if v == nil {
					return fmt.Errorf("field %s cannot be null", k)
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetID(s)
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

		case "rootUnID":
			{
				if v == nil {
					o.SetRootUnIDToNull()
					continue
				}

				if _, ok := m["rootUn"]; ok {
					continue // importing the foreign key will remove the object
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetRootUnID(s)
				}
			}

		case "rootUn":
			v2 := NewRootUn()
			m2, ok := v.(map[string]any)
			if !ok {
				return fmt.Errorf("json field %s must be a map", k)
			}
			err = v2.UnmarshalStringMap(m2)
			if err != nil {
				return
			}
			o.SetRootUn(v2)

		}
	}
	return
}
