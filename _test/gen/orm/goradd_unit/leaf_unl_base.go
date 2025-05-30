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

// LeafUnlBase is embedded in a LeafUnl object and provides the ORM access to the database.
// The member variables of the structure are private and should not normally be accessed by the LeafUnl embedder.
// Instead, use the accessor functions.
type leafUnlBase struct {
	id                string
	idIsLoaded        bool
	idIsDirty         bool
	name              string
	nameIsLoaded      bool
	nameIsDirty       bool
	rootUnlID         string
	rootUnlIDIsNull   bool
	rootUnlIDIsLoaded bool
	rootUnlIDIsDirty  bool
	objRootUnl        *RootUnl
	groLock           int64
	groLockIsLoaded   bool

	// Custom aliases, if specified
	_aliases map[string]any

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update.
	_restored bool

	_originalPK string
}

// IDs used to access the LeafUnl object fields by name using the Get function.
// doc: type=LeafUnl
const (
	LeafUnl_ID        = `ID`
	LeafUnl_Name      = `Name`
	LeafUnl_RootUnlID = `RootUnlID`
	LeafUnl_RootUnl   = `RootUnl`
	LeafUnl_GroLock   = `GroLock`
)

const LeafUnlIDMaxLength = 32    // The number of runes the column can hold
const LeafUnlNameMaxLength = 100 // The number of runes the column can hold

// Initialize or re-initialize a LeafUnl database object to default values.
// The primary key will get a temporary negative number which will be replaced when the object is saved.
// Multiple calls to Initialize are not guaranteed to create sequential values for the primary key.
func (o *leafUnlBase) Initialize() {
	o.id = db.TemporaryPrimaryKey()
	o.idIsLoaded = true
	o.idIsDirty = false

	o.name = ""
	o.nameIsLoaded = false
	o.nameIsDirty = false

	o.rootUnlID = ""
	o.rootUnlIDIsNull = true
	o.rootUnlIDIsLoaded = false
	o.rootUnlIDIsDirty = false

	o.groLock = 0
	o.groLockIsLoaded = false

	o._aliases = nil
	o._restored = false
}

// PrimaryKey returns the current value of the primary key field.
func (o *leafUnlBase) PrimaryKey() string {
	return o.id
}

// OriginalPrimaryKey returns the value of the primary key that was originally loaded into the object when it was
// read from the database.
func (o *leafUnlBase) OriginalPrimaryKey() string {
	return o._originalPK
}

// Copy copies most fields to a new LeafUnl object.
// Forward reference ids will be copied, but reverse and many-many references will not.
// Attached objects will not be included in the copy.
// Automatically generated fields will not be included in the copy.
// The primary key field will not be copied, since it is normally auto-generated.
// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *leafUnlBase) Copy() (newObject *LeafUnl) {
	newObject = NewLeafUnl()
	if o.idIsLoaded {
		newObject.SetID(o.id)
	}
	if o.nameIsLoaded {
		newObject.SetName(o.name)
	}
	if o.rootUnlIDIsLoaded {
		newObject.SetRootUnlID(o.rootUnlID)
	}
	return
}

// ID returns the value of ID.
func (o *leafUnlBase) ID() string {
	if o._restored && !o.idIsLoaded {
		panic("ID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.id
}

// IDIsLoaded returns true if the value was loaded from the database or has been set.
func (o *leafUnlBase) IDIsLoaded() bool {
	return o.idIsLoaded
}

// SetID sets the value of ID in the object, to be saved later in the database using the Save() function.
// Normally you will not need to call this function, since the ID value is automatically generated by the
// database driver. Exceptions might include importing data to a new database, or correcting primary key conflicts when
// merging data.
// You cannot change a primary key for a record that has been written to the database. While SQL databases will
// allow it, NoSql databases will not. Save a copy and delete this one instead.
func (o *leafUnlBase) SetID(v string) {
	if o._restored {
		panic("error: Do not change a primary key for a record that has been saved. Instead, save a copy and delete the original.")
	}
	if utf8.RuneCountInString(v) > LeafUnlIDMaxLength {
		panic("attempted to set LeafUnl.ID to a value larger than its maximum length in runes")
	}

	o.idIsLoaded = true
	o.id = v
	o.idIsDirty = true
}

// Name returns the value of Name.
func (o *leafUnlBase) Name() string {
	if o._restored && !o.nameIsLoaded {
		panic("Name was not selected in the last query and has not been set, and so is not valid")
	}
	return o.name
}

// NameIsLoaded returns true if the value was loaded from the database or has been set.
func (o *leafUnlBase) NameIsLoaded() bool {
	return o.nameIsLoaded
}

// SetName sets the value of Name in the object, to be saved later in the database using the Save() function.
func (o *leafUnlBase) SetName(v string) {
	if utf8.RuneCountInString(v) > LeafUnlNameMaxLength {
		panic("attempted to set LeafUnl.Name to a value larger than its maximum length in runes")
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

// RootUnlID returns the value of RootUnlID.
func (o *leafUnlBase) RootUnlID() string {
	if o._restored && !o.rootUnlIDIsLoaded {
		panic("RootUnlID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.rootUnlID
}

// RootUnlIDIsLoaded returns true if the value was loaded from the database or has been set.
func (o *leafUnlBase) RootUnlIDIsLoaded() bool {
	return o.rootUnlIDIsLoaded
}

// RootUnlIDIsNull returns true if the related database value is null.
func (o *leafUnlBase) RootUnlIDIsNull() bool {
	return o.rootUnlIDIsNull
}

// SetRootUnlID sets the value of RootUnlID in the object, to be saved later in the database using the Save() function.
func (o *leafUnlBase) SetRootUnlID(v string) {
	if o._restored &&
		o.rootUnlIDIsLoaded && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		!o.rootUnlIDIsNull && // if the db value is null, force a set of value
		o.rootUnlID == v {
		// no change
		return
	}

	o.rootUnlIDIsLoaded = true
	o.rootUnlID = v
	o.rootUnlIDIsDirty = true
	o.rootUnlIDIsNull = false
	if o.objRootUnl != nil &&
		o.rootUnlID != o.objRootUnl.PrimaryKey() {
		o.objRootUnl = nil
	}
}

// SetRootUnlIDToNull() will set the root_unl_id value in the database to NULL.
// RootUnlID() will return the column's default value after this.
func (o *leafUnlBase) SetRootUnlIDToNull() {
	if !o.rootUnlIDIsLoaded || !o.rootUnlIDIsNull {
		// If we know it is null in the database, don't save it
		o.rootUnlIDIsDirty = true
	}
	o.rootUnlIDIsLoaded = true
	o.rootUnlIDIsNull = true
	o.rootUnlID = ""
	o.objRootUnl = nil
}

// RootUnl returns the current value of the loaded RootUnl, and nil if its not loaded.
func (o *leafUnlBase) RootUnl() *RootUnl {
	return o.objRootUnl
}

// LoadRootUnl returns the related RootUnl. If it is not already loaded,
// it will attempt to load it, provided the RootUnlID column has been loaded first.
func (o *leafUnlBase) LoadRootUnl(ctx context.Context) (*RootUnl, error) {
	var err error

	if o.objRootUnl == nil {
		if !o.rootUnlIDIsLoaded {
			panic("RootUnlID must be selected in the previous query")
		}
		// Load and cache
		o.objRootUnl, err = LoadRootUnl(ctx, o.rootUnlID)
	}
	return o.objRootUnl, err
}

// SetRootUnl will set the reference to rootUnl. The referenced object
// will be saved when LeafUnl is saved. Pass nil to break the connection.
func (o *leafUnlBase) SetRootUnl(objRootUnl *RootUnl) {
	o.rootUnlIDIsLoaded = true
	if objRootUnl == nil {
		if !o.rootUnlIDIsNull || !o._restored {
			o.rootUnlIDIsNull = true
			o.rootUnlIDIsDirty = true
			o.rootUnlID = ""
			o.objRootUnl = nil
		}
	} else {
		o.objRootUnl = objRootUnl
		if o.rootUnlIDIsNull || !o._restored || o.rootUnlID != objRootUnl.PrimaryKey() {
			o.rootUnlIDIsNull = false
			o.rootUnlID = objRootUnl.PrimaryKey()
			o.rootUnlIDIsDirty = true
		}
	}
}

// GroLock returns the value of GroLock.
func (o *leafUnlBase) GroLock() int64 {
	if o._restored && !o.groLockIsLoaded {
		panic("GroLock was not selected in the last query and has not been set, and so is not valid")
	}
	return o.groLock
}

// GroLockIsLoaded returns true if the value was loaded from the database or has been set.
func (o *leafUnlBase) GroLockIsLoaded() bool {
	return o.groLockIsLoaded
}

// GetAlias returns the value for the Alias node aliasKey that was returned in the most
// recent query.
func (o *leafUnlBase) GetAlias(aliasKey string) query.AliasValue {
	if a, ok := o._aliases[aliasKey]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + aliasKey + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *leafUnlBase) IsNew() bool {
	return !o._restored
}

// LoadLeafUnl returns a LeafUnl from the database.
// selectNodes lets you provide nodes for selecting specific fields or additional fields from related tables.
// See [LeafUnlsBuilder.Select] for more info.
func LoadLeafUnl(ctx context.Context, id string, selectNodes ...query.Node) (*LeafUnl, error) {
	return queryLeafUnls(ctx).
		Where(op.Equal(node.LeafUnl().ID(), id)).
		Select(selectNodes...).
		Get()
}

// HasLeafUnl returns true if a LeafUnl with the given primary key exists in the database.
// doc: type=LeafUnl
func HasLeafUnl(ctx context.Context, id string) (bool, error) {
	v, err := queryLeafUnls(ctx).
		Where(op.Equal(node.LeafUnl().ID(), id)).
		Count()
	return v > 0, err
}

// LoadLeafUnlByRootUnlID queries for a single LeafUnl object by the given unique index values.
// selectNodes optionally let you provide nodes for joining to other tables or selecting specific fields.
// See [LeafUnlsBuilder.Select].
// If you need a more elaborate query, use QueryLeafUnls() to start a query builder.
func LoadLeafUnlByRootUnlID(ctx context.Context, rootUnlID interface{}, selectNodes ...query.Node) (*LeafUnl, error) {
	q := queryLeafUnls(ctx)
	if rootUnlID == nil {
		q = q.Where(op.IsNull(node.LeafUnl().RootUnlID()))
	} else {
		q = q.Where(op.Equal(node.LeafUnl().RootUnlID(), rootUnlID))
	}
	return q.Select(selectNodes...).Get()
}

// HasLeafUnlByRootUnlID returns true if the
// given unique index values exist in the database.
// doc: type=LeafUnl
func HasLeafUnlByRootUnlID(ctx context.Context, rootUnlID interface{}) (bool, error) {
	q := queryLeafUnls(ctx)
	if rootUnlID == nil {
		q = q.Where(op.IsNull(node.LeafUnl().RootUnlID()))
	} else {
		q = q.Where(op.Equal(node.LeafUnl().RootUnlID(), rootUnlID))
	}
	v, err := q.Count()
	return v > 0, err
}

// The LeafUnlBuilder uses a builder pattern to create a query on the database.
// Create a LeafUnlBuilder by calling QueryLeafUnls, which will select all
// the LeafUnl object in the database. Then filter and arrange those objects
// by calling Where, Select, etc.
// End a query by calling either Load, LoadI, LoadCursor, Get, or Count.
// A LeafUnlBuilder stores the context it will use to perform the query, and thus is
// meant to be a short-lived object. You should not save it for later use.
type LeafUnlBuilder struct {
	builder *query.Builder
	ctx     context.Context
}

func newLeafUnlBuilder(ctx context.Context) *LeafUnlBuilder {
	b := LeafUnlBuilder{
		builder: query.NewBuilder(node.LeafUnl()),
		ctx:     ctx,
	}
	return &b
}

// Load terminates the query builder, performs the query, and returns a slice of LeafUnl objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *LeafUnlBuilder) Load() (leafUnls []*LeafUnl, err error) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd_unit")
	var results any

	ctx := b.ctx
	results, err = database.BuilderQuery(ctx, b.builder)
	if results == nil || err != nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(LeafUnl)
		o.unpack(item, o)
		leafUnls = append(leafUnls, o)
	}
	return
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces.
// This can then satisfy a variety of interfaces that load arrays of objects, including KeyLabeler.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *LeafUnlBuilder) LoadI() (leafUnls []query.OrmObj, err error) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd_unit")
	var results any

	ctx := b.ctx
	results, err = database.BuilderQuery(ctx, b.builder)
	if results == nil || err != nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(LeafUnl)
		o.unpack(item, o)
		leafUnls = append(leafUnls, o)
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
func (b *LeafUnlBuilder) LoadCursor() (leafUnlsCursor, error) {
	b.builder.Command = query.BuilderCommandLoadCursor
	database := db.GetDatabase("goradd_unit")
	result, err := database.BuilderQuery(b.ctx, b.builder)
	var cursor query.CursorI
	if result != nil {
		cursor = result.(query.CursorI)
	}
	return leafUnlsCursor{cursor}, err
}

type leafUnlsCursor struct {
	query.CursorI
}

// Next returns the current LeafUnl object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c leafUnlsCursor) Next() (*LeafUnl, error) {
	if c.CursorI == nil {
		return nil, nil
	}

	row, err := c.CursorI.Next()
	if row == nil || err != nil {
		return nil, err
	}
	o := new(LeafUnl)
	o.unpack(row, o)
	return o, nil
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
// If an error occurs, or no results are found, a nil is returned.
func (b *LeafUnlBuilder) Get() (*LeafUnl, error) {
	results, err := b.Load()
	if err != nil || len(results) == 0 {
		return nil, err
	}
	return results[0], nil
}

// Where adds a condition to filter what gets selected.
// Calling Where multiple times will AND the conditions together.
func (b *LeafUnlBuilder) Where(c query.Node) *LeafUnlBuilder {
	b.builder.Where(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
// By default, the given nodes are sorted in ascending order.
// Add Descending() to the node to specify that it should be sorted in descending order.
func (b *LeafUnlBuilder) OrderBy(nodes ...query.Sorter) *LeafUnlBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified.
// For large data sets and specific types of queries, this can be slow, because it will perform
// the entire query before computing the limit.
// You cannot limit a query that has embedded arrays.
func (b *LeafUnlBuilder) Limit(maxRowCount int, offset int) *LeafUnlBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select specifies what specific columns will be loaded with data.
// By default, all the columns of the leaf_unl table will be queried and loaded.
// If nodes contains columns from the leaf_unl table, that will limit the columns queried and loaded to only those columns.
// If related tables are specified, then all the columns from those tables are queried, selected and joined to the result.
// If columns in related tables are specified, then only those columns will be queried and loaded.
// Depending on the query, additional columns may automatically be added to the query. In particular, primary key columns
// will be added in most situations. The exception to this would be in distinct queries, group by queries, or subqueries.
func (b *LeafUnlBuilder) Select(nodes ...query.Node) *LeafUnlBuilder {
	b.builder.Select(nodes...)
	return b
}

// Calculation adds operation as an aliased value onto base.
// After the query, you can read the data by passing alias to GetAlias on the returned object.
func (b *LeafUnlBuilder) Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) *LeafUnlBuilder {
	b.builder.Calculation(base, alias, operation)
	return b
}

// Distinct removes duplicates from the results of the query.
// Adding a Select() is usually required.
func (b *LeafUnlBuilder) Distinct() *LeafUnlBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions with Calculation.
func (b *LeafUnlBuilder) GroupBy(nodes ...query.Node) *LeafUnlBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query after the query is performed.
func (b *LeafUnlBuilder) Having(node query.Node) *LeafUnlBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items in the result.
// If you have Select or Calculation columns in the query, it will count NULL results as well.
// To not count NULL values, use Where in the builder with a NotNull operation.
// To count distinct combinations of items, call Distinct() on the builder.
func (b *LeafUnlBuilder) Count() (int, error) {
	b.builder.Command = query.BuilderCommandCount
	database := db.GetDatabase("goradd_unit")

	ctx := b.ctx
	results, err := database.BuilderQuery(ctx, b.builder)
	if results == nil || err != nil {
		return 0, err
	}
	return results.(int), nil
}

// CountLeafUnls returns the total number of items in the leaf_unl table.
func CountLeafUnls(ctx context.Context) (int, error) {
	return QueryLeafUnls(ctx).Count()
}

// CountLeafUnlsByRootUnlID queries the database and returns the number of LeafUnl objects that
// have rootUnlID.
// doc: type=LeafUnl
func CountLeafUnlsByRootUnlID(ctx context.Context, rootUnlID string) (int, error) {
	v_rootUnlID := rootUnlID
	return QueryLeafUnls(ctx).
		Where(op.Equal(node.LeafUnl().RootUnlID(), v_rootUnlID)).
		Count()
}

// unpack recursively transforms data coming from the database into ORM objects.
func (o *leafUnlBase) unpack(m map[string]interface{}, objThis *LeafUnl) {

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

	if v, ok := m["root_unl_id"]; ok {
		if v == nil {
			o.rootUnlID = ""
			o.rootUnlIDIsNull = true
			o.rootUnlIDIsLoaded = true
			o.rootUnlIDIsDirty = false
		} else if o.rootUnlID, ok = v.(string); ok {
			o.rootUnlIDIsNull = false
			o.rootUnlIDIsLoaded = true
			o.rootUnlIDIsDirty = false
		} else {
			panic("Wrong type found for root_unl_id.")
		}
	} else {
		o.rootUnlIDIsLoaded = false
		o.rootUnlIDIsNull = true
		o.rootUnlID = ""
		o.rootUnlIDIsDirty = false
	}

	if v, ok := m["RootUnl"]; ok {
		if objRootUnl, ok2 := v.(map[string]any); ok2 {
			o.objRootUnl = new(RootUnl)
			o.objRootUnl.unpack(objRootUnl, o.objRootUnl)
			o.rootUnlIDIsLoaded = true
			o.rootUnlIDIsDirty = false
		} else {
			panic("Wrong type found for RootUnl object.")
		}
	} else {
		o.objRootUnl = nil
	}

	if v, ok := m["gro_lock"]; ok && v != nil {
		if o.groLock, ok = v.(int64); ok {
			o.groLockIsLoaded = true

		} else {
			panic("Wrong type found for gro_lock.")
		}
	} else {
		o.groLockIsLoaded = false
		o.groLock = 0
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = v.(map[string]any)
	}

	o._restored = true

}

// save will update or insert the object, depending on the state of the object.
func (o *leafUnlBase) save(ctx context.Context) error {
	if o._restored {
		return o.update(ctx)
	} else {
		return o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
// If the table has auto-generated values, those will be updated automatically.
func (o *leafUnlBase) update(ctx context.Context) error {
	if !o._restored {
		panic("cannot update a record that was not originally read from the database.")
	}
	if !o.IsDirty() {
		return nil // nothing to save
	}

	var modifiedFields map[string]interface{}
	var newLock int64

	d := Database()
	err := db.ExecuteTransaction(ctx, d, func() error {

		// Save loaded RootUnl object to get its new pk and update it here.
		if o.objRootUnl != nil {
			if err := o.objRootUnl.Save(ctx); err != nil {
				return err
			}
			o.SetRootUnlID(o.objRootUnl.PrimaryKey())
		}

		if o.rootUnlIDIsDirty &&
			!o.rootUnlIDIsNull {
			if obj, err := LoadLeafUnlByRootUnlID(ctx, o.rootUnlID); err != nil {
				return err
			} else if obj != nil {
				return db.NewUniqueValueError("leaf_unl", map[string]any{"root_unl_id": o.rootUnlID}, nil)
			}
		}

		modifiedFields = getLeafUnlUpdateFields(o)
		if len(modifiedFields) != 0 {
			var err2 error

			// If this panics with an invalid GroLock value, then the GroLock field was not selected in a prior query. Be sure to include it in any Select statements.
			newLock, err2 = d.Update(ctx, "leaf_unl", "id", o._originalPK, modifiedFields, "gro_lock", o.GroLock())
			if err2 != nil {
				return err2
			}
		}

		return nil
	}) // transaction
	if err != nil {
		return err
	}
	if newLock != 0 {
		o.groLock = newLock
		o.groLockIsLoaded = true
	}

	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd_unit", "leaf_unl", o._originalPK, anyutil.SortedKeys(modifiedFields)...)
	}

	return nil
}

// insert will insert the object into the database. Related items will be saved.
func (o *leafUnlBase) insert(ctx context.Context) (err error) {
	var insertFields map[string]interface{}
	d := Database()
	err = db.ExecuteTransaction(ctx, d, func() error {

		// Save loaded RootUnl object to get its new pk and update it here.
		if o.objRootUnl != nil {
			if err := o.objRootUnl.Save(ctx); err != nil {
				return err
			}
			o.SetRootUnlID(o.objRootUnl.PrimaryKey())
		}

		if !o.nameIsLoaded {
			panic("a value for Name is required, and there is no default value. Call SetName() before inserting the record.")
		}
		if o.rootUnlIDIsDirty &&
			!o.rootUnlIDIsNull {
			if obj, err := LoadLeafUnlByRootUnlID(ctx, o.rootUnlID); err != nil {
				return err
			} else if obj != nil {
				return db.NewUniqueValueError("leaf_unl", map[string]any{"root_unl_id": o.rootUnlID}, nil)
			}
		}
		insertFields = getLeafUnlInsertFields(o)
		var newPK string
		newPK, err = d.Insert(ctx, "leaf_unl", "id", insertFields)
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
	if t, ok := insertFields["gro_lock"]; ok {
		o.groLock = t.(int64)
		o.groLockIsLoaded = true
	}

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd_unit", "leaf_unl", o.PrimaryKey())
	return
}

// getUpdateFields returns the database columns that will be sent to the update process.
// This will include timestamp fields only if some other column has changed.
func (o *leafUnlBase) getUpdateFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.nameIsDirty {
		fields["name"] = o.name
	}
	if o.rootUnlIDIsDirty {
		if o.rootUnlIDIsNull {
			fields["root_unl_id"] = nil
		} else {
			fields["root_unl_id"] = o.rootUnlID
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
func (o *leafUnlBase) getInsertFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}

	fields["name"] = o.name
	if o.rootUnlIDIsNull {
		fields["root_unl_id"] = nil
	} else {
		fields["root_unl_id"] = o.rootUnlID
	}
	fields["gro_lock"] = db.RecordVersion(0)
	return
}

// Delete deletes the record from the database.
func (o *leafUnlBase) Delete(ctx context.Context) (err error) {
	if o == nil {
		return // allow deleting of a nil object to be a noop
	}
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	err = d.Delete(ctx, "leaf_unl", "id", o.id, "gro_lock", o.GroLock())
	if err != nil {
		return err
	}
	broadcast.Delete(ctx, "goradd_unit", "leaf_unl", fmt.Sprint(o.id))
	return
}

// deleteLeafUnl deletes the LeafUnl with primary key pk from the database
// and handles associated records.
func deleteLeafUnl(ctx context.Context, pk string) error {
	d := db.GetDatabase("goradd_unit")
	err := d.Delete(ctx, "leaf_unl", "id", pk, "", 0)
	if err != nil {
		return err
	}
	broadcast.Delete(ctx, "goradd_unit", "leaf_unl", fmt.Sprint(pk))
	return err
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *leafUnlBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.nameIsDirty = false
	o.rootUnlIDIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database or created.
// However, a new object that has a column with a default value will be automatically marked as dirty upon creation.
func (o *leafUnlBase) IsDirty() (dirty bool) {
	dirty = o.idIsDirty ||
		o.nameIsDirty ||
		o.rootUnlIDIsDirty ||
		(o.objRootUnl != nil && o.objRootUnl.IsDirty())

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil.
// Get can be used to retrieve a value by using the Identifier of a node.
func (o *leafUnlBase) Get(key string) interface{} {

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

	case "RootUnlID":
		if !o.rootUnlIDIsLoaded {
			return nil
		}
		return o.rootUnlID

	case "RootUnl":
		return o.RootUnl()

	case "GroLock":
		if !o.groLockIsLoaded {
			return nil
		}
		return o.groLock

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *leafUnlBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := o.encodeTo(enc); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (o *leafUnlBase) encodeTo(enc db.Encoder) error {

	if err := enc.Encode(o.id); err != nil {
		return fmt.Errorf("error encoding LeafUnl.id: %w", err)
	}
	if err := enc.Encode(o.idIsLoaded); err != nil {
		return fmt.Errorf("error encoding LeafUnl.idIsLoaded: %w", err)
	}
	if err := enc.Encode(o.idIsDirty); err != nil {
		return fmt.Errorf("error encoding LeafUnl.idIsDirty: %w", err)
	}

	if err := enc.Encode(o.name); err != nil {
		return fmt.Errorf("error encoding LeafUnl.name: %w", err)
	}
	if err := enc.Encode(o.nameIsLoaded); err != nil {
		return fmt.Errorf("error encoding LeafUnl.nameIsLoaded: %w", err)
	}
	if err := enc.Encode(o.nameIsDirty); err != nil {
		return fmt.Errorf("error encoding LeafUnl.nameIsDirty: %w", err)
	}

	if err := enc.Encode(o.rootUnlID); err != nil {
		return fmt.Errorf("error encoding LeafUnl.rootUnlID: %w", err)
	}
	if err := enc.Encode(o.rootUnlIDIsNull); err != nil {
		return fmt.Errorf("error encoding LeafUnl.rootUnlIDIsNull: %w", err)
	}
	if err := enc.Encode(o.rootUnlIDIsLoaded); err != nil {
		return fmt.Errorf("error encoding LeafUnl.rootUnlIDIsLoaded: %w", err)
	}
	if err := enc.Encode(o.rootUnlIDIsDirty); err != nil {
		return fmt.Errorf("error encoding LeafUnl.rootUnlIDIsDirty: %w", err)
	}

	if o.objRootUnl == nil {
		if err := enc.Encode(false); err != nil {
			return err
		}
	} else {
		if err := enc.Encode(true); err != nil {
			return err
		}
		if err := enc.Encode(o.objRootUnl); err != nil {
			return fmt.Errorf("error encoding LeafUnl.objRootUnl: %w", err)
		}
	}

	if err := enc.Encode(o.groLock); err != nil {
		return fmt.Errorf("error encoding LeafUnl.groLock: %w", err)
	}
	if err := enc.Encode(o.groLockIsLoaded); err != nil {
		return fmt.Errorf("error encoding LeafUnl.groLockIsLoaded: %w", err)
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
			return fmt.Errorf("error encoding LeafUnl._aliases: %w", err)
		}
	}

	if err := enc.Encode(o._restored); err != nil {
		return fmt.Errorf("error encoding LeafUnl._restored: %w", err)
	}
	if err := enc.Encode(o._originalPK); err != nil {
		return fmt.Errorf("error encoding LeafUnl._originalPK: %w", err)
	}
	return nil
}

// UnmarshalBinary converts a structure that was created with MarshalBinary into a LeafUnl object.
func (o *leafUnlBase) UnmarshalBinary(data []byte) (err error) {
	buf := bytes.NewReader(data)
	dec := gob.NewDecoder(buf)
	return o.decodeFrom(dec)
}

func (o *leafUnlBase) decodeFrom(dec db.Decoder) (err error) {
	var isPtr bool

	_ = isPtr
	if err = dec.Decode(&o.id); err != nil {
		return fmt.Errorf("error decoding LeafUnl.id: %w", err)
	}
	if err = dec.Decode(&o.idIsLoaded); err != nil {
		return fmt.Errorf("error decoding LeafUnl.idIsLoaded: %w", err)
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return fmt.Errorf("error decoding LeafUnl.idIsDirty: %w", err)
	}

	if err = dec.Decode(&o.name); err != nil {
		return fmt.Errorf("error decoding LeafUnl.name: %w", err)
	}
	if err = dec.Decode(&o.nameIsLoaded); err != nil {
		return fmt.Errorf("error decoding LeafUnl.nameIsLoaded: %w", err)
	}
	if err = dec.Decode(&o.nameIsDirty); err != nil {
		return fmt.Errorf("error decoding LeafUnl.nameIsDirty: %w", err)
	}

	if err = dec.Decode(&o.rootUnlID); err != nil {
		return fmt.Errorf("error decoding LeafUnl.rootUnlID: %w", err)
	}
	if err = dec.Decode(&o.rootUnlIDIsNull); err != nil {
		return fmt.Errorf("error decoding LeafUnl.rootUnlIDIsNull: %w", err)
	}
	if err = dec.Decode(&o.rootUnlIDIsLoaded); err != nil {
		return fmt.Errorf("error decoding LeafUnl.rootUnlIDIsLoaded: %w", err)
	}
	if err = dec.Decode(&o.rootUnlIDIsDirty); err != nil {
		return fmt.Errorf("error decoding LeafUnl.rootUnlIDIsDirty: %w", err)
	}

	if err = dec.Decode(&isPtr); err != nil {
		return fmt.Errorf("error decoding LeafUnl.objRootUnl isPtr: %w", err)
	}
	if isPtr {
		if err = dec.Decode(&o.objRootUnl); err != nil {
			return fmt.Errorf("error decoding LeafUnl.objRootUnl: %w", err)
		}
	}
	if err = dec.Decode(&o.groLock); err != nil {
		return fmt.Errorf("error decoding LeafUnl.groLock: %w", err)
	}
	if err = dec.Decode(&o.groLockIsLoaded); err != nil {
		return fmt.Errorf("error decoding LeafUnl.groLockIsLoaded: %w", err)
	}

	if err = dec.Decode(&isPtr); err != nil {
		return fmt.Errorf("error decoding LeafUnl._aliases isPtr: %w", err)
	}
	if isPtr {
		if err = dec.Decode(&o._aliases); err != nil {
			return fmt.Errorf("error decoding LeafUnl._aliases: %w", err)
		}
	}

	if err = dec.Decode(&o._restored); err != nil {
		return fmt.Errorf("error decoding LeafUnl._restored: %w", err)
	}
	if err = dec.Decode(&o._originalPK); err != nil {
		return fmt.Errorf("error decoding LeafUnl._originalPK: %w", err)
	}
	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *leafUnlBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *leafUnlBase) MarshalStringMap() map[string]interface{} {
	v := make(map[string]interface{})

	if o.idIsLoaded {
		v["id"] = o.id
	}

	if o.nameIsLoaded {
		v["name"] = o.name
	}

	if val := o.objRootUnl; val != nil {
		v["rootUnl"] = val.MarshalStringMap()
	} else if o.rootUnlIDIsLoaded {
		if o.rootUnlIDIsNull {
			v["rootUnlID"] = nil
		} else {
			v["rootUnlID"] = o.rootUnlID
		}
	}

	if o.groLockIsLoaded {
		v["groLock"] = o.groLock
	}

	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the LeafUnl. The LeafUnl can be a
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
//	"rootUnlID" - string, nullable
//	"groLock" - int64
func (o *leafUnlBase) UnmarshalJSON(data []byte) (err error) {
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
// Override this in LeafUnl to modify the json before sending it here.
func (o *leafUnlBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
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

		case "rootUnlID":
			{
				if v == nil {
					o.SetRootUnlIDToNull()
					continue
				}

				if _, ok := m["rootUnl"]; ok {
					continue // importing the foreign key will remove the object
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetRootUnlID(s)
				}
			}

		case "rootUnl":
			v2 := NewRootUnl()
			m2, ok := v.(map[string]any)
			if !ok {
				return fmt.Errorf("json field %s must be a map", k)
			}
			err = v2.UnmarshalStringMap(m2)
			if err != nil {
				return
			}
			o.SetRootUnl(v2)

		}
	}
	return
}
