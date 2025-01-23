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

// AddressBase is embedded in a Address object and provides the ORM access to the database.
// The member variables of the structure are private and should not normally be accessed by the Address embedder.
// Instead, use the accessor functions.
type addressBase struct {
	id        string
	idIsValid bool
	idIsDirty bool

	personID        string
	personIDIsValid bool
	personIDIsDirty bool
	objPerson       *Person

	street        string
	streetIsValid bool
	streetIsDirty bool

	city        string
	cityIsNull  bool
	cityIsValid bool
	cityIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]any

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update.
	_restored bool

	_originalPK string
}

// IDs used to access the Address object fields by name using the Get function.
// doc: type=Address
const (
	Address_ID       = `ID`
	Address_PersonID = `PersonID`
	Address_Person   = `Person`
	Address_Street   = `Street`
	Address_City     = `City`
)

const AddressStreetMaxLength = 100 // The number of runes the column can hold
const AddressCityMaxLength = 100   // The number of runes the column can hold

// Initialize or re-initialize a Address database object to default values.
// The primary key will get a temporary negative number which will be replaced when the object is saved.
// Multiple calls to Initialize are not guaranteed to create sequential values for the primary key.
func (o *addressBase) Initialize() {

	newObjectPkCounter = newObjectPkCounter - 1
	o.id = fmt.Sprintf("%d", newObjectPkCounter)

	o.idIsValid = false
	o.idIsDirty = false

	o.personID = ""

	o.personIDIsValid = false
	o.personIDIsDirty = false

	o.street = ""

	o.streetIsValid = false
	o.streetIsDirty = false

	o.city = "BOB"

	o.cityIsNull = false
	o.cityIsValid = true
	o.cityIsDirty = true

	o._restored = false
}

// PrimaryKey returns the current value of the primary key field.
func (o *addressBase) PrimaryKey() string {
	return o.id
}

// OriginalPrimaryKey returns the value of the primary key that was originally loaded into the object when it was
// read from the database.
func (o *addressBase) OriginalPrimaryKey() string {
	return o._originalPK
}

// Copy copies all valid fields to a new Address object.
// Forward reference ids will be copied, but reverse and many-many references will not.
// Attached objects will not be included in the copy.
// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *addressBase) Copy() (newObject *Address) {
	newObject = NewAddress()
	if o.personIDIsValid {
		newObject.SetPersonID(o.personID)
	}
	if o.streetIsValid {
		newObject.SetStreet(o.street)
	}
	if o.cityIsValid {
		newObject.SetCity(o.city)
	}
	return
}

// ID returns the loaded value of ID or
// the zero value if not loaded. Call IDIsValid() to determine
// if it is loaded.
func (o *addressBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *addressBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

// PersonID returns the loaded value of PersonID.
func (o *addressBase) PersonID() string {
	if o._restored && !o.personIDIsValid {
		panic("PersonID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.personID
}

// PersonIDIsValid returns true if the value was loaded from the database or has been set.
func (o *addressBase) PersonIDIsValid() bool {
	return o.personIDIsValid
}

// SetPersonID sets the value of PersonID in the object, to be saved later using the Save() function.
func (o *addressBase) SetPersonID(personID string) {
	o.personIDIsValid = true
	if o.personID != personID || !o._restored {
		o.personID = personID
		o.personIDIsDirty = true
		o.objPerson = nil
	}

}

// Person returns the current value of the loaded Person, and nil if its not loaded.
func (o *addressBase) Person() *Person {
	return o.objPerson
}

// LoadPerson returns the related Person. If it is not already loaded,
// it will attempt to load it, provided the PersonID column has been loaded first.
func (o *addressBase) LoadPerson(ctx context.Context) *Person {
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
func (o *addressBase) SetPerson(objPerson *Person) {
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

// Street returns the loaded value of Street.
func (o *addressBase) Street() string {
	if o._restored && !o.streetIsValid {
		panic("Street was not selected in the last query and has not been set, and so is not valid")
	}
	return o.street
}

// StreetIsValid returns true if the value was loaded from the database or has been set.
func (o *addressBase) StreetIsValid() bool {
	return o.streetIsValid
}

// SetStreet sets the value of Street in the object, to be saved later using the Save() function.
func (o *addressBase) SetStreet(street string) {
	o.streetIsValid = true
	if utf8.RuneCountInString(street) > AddressStreetMaxLength {
		panic("attempted to set Address.Street to a value larger than its maximum length in runes")
	}
	if o.street != street || !o._restored {
		o.street = street
		o.streetIsDirty = true
	}

}

// City returns the loaded value of City.
func (o *addressBase) City() string {
	if o._restored && !o.cityIsValid {
		panic("City was not selected in the last query and has not been set, and so is not valid")
	}
	return o.city
}

// CityIsValid returns true if the value was loaded from the database or has been set.
func (o *addressBase) CityIsValid() bool {
	return o.cityIsValid
}

// CityIsNull returns true if the related database value is null.
func (o *addressBase) CityIsNull() bool {
	return o.cityIsNull
}

// City_I returns the loaded value of City as an interface.
// If the value in the database is NULL, a nil interface is returned.
func (o *addressBase) City_I() interface{} {
	if o._restored && !o.cityIsValid {
		panic("city was not selected in the last query and has not been set, and so is not valid")
	} else if o.cityIsNull {
		return nil
	}
	return o.city
}

func (o *addressBase) SetCity(i interface{}) {
	o.cityIsValid = true
	if i == nil {
		if !o.cityIsNull {
			o.cityIsNull = true
			o.cityIsDirty = true
			o.city = "BOB"
		}
	} else {
		v := i.(string)

		if utf8.RuneCountInString(v) > AddressCityMaxLength {
			panic("attempted to set Address.City to a value larger than its maximum length in runes")
		}
		if o.cityIsNull ||
			!o._restored ||
			o.city != v {
			o.cityIsNull = false
			o.city = v
			o.cityIsDirty = true
		}
	}
}

// GetAlias returns the alias for the given key.
func (o *addressBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *addressBase) IsNew() bool {
	return !o._restored
}

// LoadAddress returns a Address from the database.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields.
// Table nodes will be considered Join nodes, and column nodes will be Select nodes.
// See [AddressBuilder.Join] and [AddressesBuilder.Select] for more info.
func LoadAddress(ctx context.Context, id string, joinOrSelectNodes ...query.Node) *Address {
	return queryAddresses(ctx).
		Where(op.Equal(node.Address().ID(), id)).
		joinOrSelect(joinOrSelectNodes...).
		Get()
}

// HasAddress returns true if a Address with the given primaryKey exists in the database.
// doc: type=Address
func HasAddress(ctx context.Context, id string) bool {
	return queryAddresses(ctx).
		Where(op.Equal(node.Address().ID(), id)).
		Count(false) == 1
}

// The AddressBuilder uses the query.BuilderI interface to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, LoadCursor, Get, Count, or Delete
type AddressBuilder interface {
	// Join adds node n to the node tree so that its fields will appear in the query.
	// Optionally add conditions to filter what gets included. Multiple conditions are anded.
	Join(n query.Node, conditions ...query.Node) AddressBuilder

	// Where adds a condition to filter what gets selected.
	// Calling Where multiple times will AND the conditions together.
	Where(c query.Node) AddressBuilder

	// OrderBy specifies how the resulting data should be sorted.
	// By default, the given nodes are sorted in ascending order.
	// Add Descending() to the node to specify that it should be sorted in descending order.
	OrderBy(nodes ...query.Sorter) AddressBuilder

	// Limit will return a subset of the data, limited to the offset and number of rows specified.
	// For large data sets and specific types of queries, this can be slow, because it will perform
	// the entire query before computing the limit.
	// You cannot limit a query that has embedded arrays.
	Limit(maxRowCount int, offset int) AddressBuilder

	// Select optimizes the query to only return the specified fields.
	// Once you put a Select in your query, you must specify all the fields that you will eventually read out.
	// Some fields, like primary keys, are always selected.
	// If you are using a GroupBy, most database drivers will only allow selecting on fields in the GroupBy, and
	// doing otherwise will result in an error.
	Select(nodes ...query.Node) AddressBuilder

	// Calculation adds a calculation node with an aliased name.
	// After the query, you can read the data using GetAlias() on a returned object.
	Calculation(name string, n query.Aliaser) AddressBuilder

	// Distinct removes duplicates from the results of the query.
	// Adding a Select() is usually required.
	Distinct() AddressBuilder

	// GroupBy controls how results are grouped when using aggregate functions with Calculation.
	GroupBy(nodes ...query.Node) AddressBuilder

	// Having does additional filtering on the results of the query after the query is performed.
	Having(node query.Node) AddressBuilder

	// Load terminates the query builder, performs the query, and returns a slice of Address objects.
	// If there are any errors, nil is returned and the specific error is stored in the context.
	// If no results come back from the query, it will return a non-nil empty slice.
	Load() []*Address
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
	LoadCursor() addressesCursor

	// Get is a convenience method to return only the first item found in a query.
	// The entire query is performed, so you should generally use this only if you know
	// you are selecting on one or very few items.
	//
	// If an error occurs, or no results are found, a nil is returned.
	// In the case of an error, the error is returned in the context.
	Get() *Address

	// Count terminates a query and returns just the number of items selected.
	// distinct wll count the number of distinct items, ignoring duplicates.
	// nodes will select individual fields, and should be accompanied by a GroupBy.
	Count(distinct bool, nodes ...query.Node) int

	// Delete uses the query builder to delete a group of records that match the criteria
	Delete()

	// Subquery terminates the query builder and tags it as a subquery within a larger query.
	// You MUST include what you are selecting by adding Calculation or Select functions on the subquery builder.
	// Generally you would use this as a node to a Calculation function on the surrounding query builder.
	Subquery() *query.SubqueryNode

	joinOrSelect(nodes ...query.Node) AddressBuilder
}

type addressQueryBuilder struct {
	builder *query.Builder
}

func newAddressBuilder(ctx context.Context) AddressBuilder {
	b := addressQueryBuilder{
		builder: query.NewBuilder(ctx),
	}
	return b.Join(node.Address()) // seed builder with the top table
}

// Load terminates the query builder, performs the query, and returns a slice of Address objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *addressQueryBuilder) Load() (addresses []*Address) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder.Ctx, b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(Address)
		o.load(item, o, nil, "")
		addresses = append(addresses, o)
	}
	return
}

// Load terminates the query builder, performs the query, and returns a slice of interfaces.
// This can then satisfy a general interface that loads arrays of objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *addressQueryBuilder) LoadI() (addresses []any) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder.Ctx, b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(Address)
		o.load(item, o, nil, "")
		addresses = append(addresses, o)
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
func (b *addressQueryBuilder) LoadCursor() addressesCursor {
	b.builder.Command = query.BuilderCommandLoadCursor
	database := db.GetDatabase("goradd")
	result := database.BuilderQuery(b.builder.Ctx, b.builder)
	if result == nil {
		return addressesCursor{}
	}
	cursor := result.(query.CursorI)

	return addressesCursor{cursor}
}

type addressesCursor struct {
	query.CursorI
}

// Next returns the current Address object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c addressesCursor) Next() *Address {
	if c.CursorI == nil {
		return nil
	}

	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(Address)
	o.load(row, o, nil, "")
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
//
// If an error occurs, or no results are found, a nil is returned.
// In the case of an error, the error is returned in the context.
func (b *addressQueryBuilder) Get() *Address {
	results := b.Load()
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Join adds node n to the node tree so that its fields will appear in the query.
// Optionally add conditions to filter what gets included. Multiple conditions are anded.
func (b *addressQueryBuilder) Join(n query.Node, conditions ...query.Node) AddressBuilder {
	if query.RootNode(n).TableName_() != "address" {
		panic("you can only join a node that is rooted at node.Address()")
	}

	var condition query.Node
	if len(conditions) > 1 {
		condition = op.And(conditions)
	} else if len(conditions) == 1 {
		condition = conditions[0]
	}
	b.builder.Join(n, condition)
	return b
}

// Where adds a condition to filter what gets selected.
// Calling Where multiple times will AND the conditions together.
func (b *addressQueryBuilder) Where(c query.Node) AddressBuilder {
	b.builder.Where(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
// By default, the given nodes are sorted in ascending order.
// Add Descending() to the node to specify that it should be sorted in descending order.
func (b *addressQueryBuilder) OrderBy(nodes ...query.Sorter) AddressBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified.
// For large data sets and specific types of queries, this can be slow, because it will perform
// the entire query before computing the limit.
// You cannot limit a query that has embedded arrays.
func (b *addressQueryBuilder) Limit(maxRowCount int, offset int) AddressBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields.
// Once you put a Select in your query, you must specify all the fields that you will eventually read out.
func (b *addressQueryBuilder) Select(nodes ...query.Node) AddressBuilder {
	b.builder.Select(nodes...)
	return b
}

// Calculation adds a calculation node with an aliased name.
// After the query, you can read the data using GetAlias() on a returned object.
func (b *addressQueryBuilder) Calculation(name string, n query.Aliaser) AddressBuilder {
	b.builder.Calculation(name, n)
	return b
}

// Distinct removes duplicates from the results of the query.
// Adding a Select() is usually required.
func (b *addressQueryBuilder) Distinct() AddressBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions with Calculation.
func (b *addressQueryBuilder) GroupBy(nodes ...query.Node) AddressBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query after the query is performed.
func (b *addressQueryBuilder) Having(node query.Node) AddressBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
// distinct wll count the number of distinct items, ignoring duplicates.
// nodes will select individual fields, and should be accompanied by a GroupBy.
func (b *addressQueryBuilder) Count(distinct bool, nodes ...query.Node) int {
	b.builder.Command = query.BuilderCommandCount
	if distinct {
		b.builder.Distinct()
	}
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder.Ctx, b.builder)
	if results == nil {
		return 0
	}
	return results.(int)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *addressQueryBuilder) Delete() {
	b.builder.Command = query.BuilderCommandDelete
	database := db.GetDatabase("goradd")
	database.BuilderQuery(b.builder.Ctx, b.builder)
	broadcast.BulkChange(b.builder.Context(), "goradd", "address")
}

// Subquery terminates the query builder and tags it as a subquery within a larger query.
// You MUST include what you are selecting by adding Calculation or Select functions on the subquery builder.
// Generally you would use this as a node to a Calculation function on the surrounding query builder.
func (b *addressQueryBuilder) Subquery() *query.SubqueryNode {
	return b.builder.Subquery()
}

// joinOrSelect is a private helper function for the Load* functions
func (b *addressQueryBuilder) joinOrSelect(nodes ...query.Node) AddressBuilder {
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

// CountAddressByID queries the database and returns the number of Address objects that
// have id.
// doc: type=Address
func CountAddressByID(ctx context.Context, id string) int {
	return int(queryAddresses(ctx).Where(op.Equal(node.Address().ID(), id)).Count(false))
}

// CountAddressByPersonID queries the database and returns the number of Address objects that
// have personID.
// doc: type=Address
func CountAddressByPersonID(ctx context.Context, personID string) int {
	if personID == "" {
		return 0
	}
	return int(queryAddresses(ctx).Where(op.Equal(node.Address().PersonID(), personID)).Count(false))
}

// CountAddressByStreet queries the database and returns the number of Address objects that
// have street.
// doc: type=Address
func CountAddressByStreet(ctx context.Context, street string) int {
	return int(queryAddresses(ctx).Where(op.Equal(node.Address().Street(), street)).Count(false))
}

// CountAddressByCity queries the database and returns the number of Address objects that
// have city.
// doc: type=Address
func CountAddressByCity(ctx context.Context, city string) int {
	return int(queryAddresses(ctx).Where(op.Equal(node.Address().City(), city)).Count(false))
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
func (o *addressBase) load(m map[string]interface{}, objThis *Address, objParent interface{}, parentKey string) {

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
			o.objPerson.load(objPerson, o.objPerson, objThis, "Addresses")
			o.personIDIsValid = true
			o.personIDIsDirty = false
		} else {
			panic("Wrong type found for Person object.")
		}
	} else {
		o.objPerson = nil
	}

	if v, ok := m["street"]; ok && v != nil {
		if o.street, ok = v.(string); ok {
			o.streetIsValid = true
			o.streetIsDirty = false

		} else {
			panic("Wrong type found for street.")
		}
	} else {
		o.streetIsValid = false
		o.street = ""
	}

	if v, ok := m["city"]; ok {
		if v == nil {
			o.city = "BOB"
			o.cityIsNull = true
			o.cityIsValid = true
			o.cityIsDirty = false
		} else if o.city, ok = v.(string); ok {
			o.cityIsNull = false
			o.cityIsValid = true
			o.cityIsDirty = false
		} else {
			panic("Wrong type found for city.")
		}
	} else {
		o.cityIsValid = false
		o.cityIsNull = true
		o.city = "BOB"
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}

	o._restored = true

}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *addressBase) Save(ctx context.Context) {
	if o._restored {
		o.update(ctx)
	} else {
		o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *addressBase) update(ctx context.Context) {
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
			d.Update(ctx, "address", modifiedFields, map[string]any{"id": o._originalPK})
		}

	}) // transaction

	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd", "address", o._originalPK, all.SortedKeys(modifiedFields)...)
	}
}

// insert will insert the object into the database. Related items will be saved.
func (o *addressBase) insert(ctx context.Context) {
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

		if !o.streetIsValid {
			panic("a value for Street is required, and there is no default value. Call SetStreet() before inserting the record.")
		}

		m := o.getValidFields()

		id := d.Insert(ctx, "address", m)
		o.id = id
		o._originalPK = id

	}) // transaction

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "address", o.PrimaryKey())
}

// getModifiedFields returns the database columns that have been modified. This
// will determine which specific fields are sent to the database to be changed.
func (o *addressBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.personIDIsDirty {
		fields["person_id"] = o.personID
	}
	if o.streetIsDirty {
		fields["street"] = o.street
	}
	if o.cityIsDirty {
		if o.cityIsNull {
			fields["city"] = nil
		} else {
			fields["city"] = o.city
		}
	}
	return
}

// getValidFields returns the fields that have valid data in them in a form ready to send to the database.
func (o *addressBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.personIDIsValid {
		fields["person_id"] = o.personID
	}
	if o.streetIsValid {
		fields["street"] = o.street
	}
	if o.cityIsValid {
		if o.cityIsNull {
			fields["city"] = nil
		} else {
			fields["city"] = o.city
		}
	}
	return
}

// Delete deletes the record from the database.
func (o *addressBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "address", map[string]any{"ID": o.id})
	broadcast.Delete(ctx, "goradd", "address", fmt.Sprint(o.id))
}

// deleteAddress deletes the associated record from the database.
func deleteAddress(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "address", map[string]any{"ID": pk})
	broadcast.Delete(ctx, "goradd", "address", fmt.Sprint(pk))
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *addressBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.personIDIsDirty = false
	o.streetIsDirty = false
	o.cityIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database.
func (o *addressBase) IsDirty() (dirty bool) {
	dirty = o.idIsDirty ||
		o.personIDIsDirty ||
		(o.objPerson != nil && o.objPerson.IsDirty()) ||
		o.streetIsDirty ||
		o.cityIsDirty

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *addressBase) Get(key string) interface{} {

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

	case "Street":
		if !o.streetIsValid {
			return nil
		}
		return o.street

	case "City":
		if !o.cityIsValid {
			return nil
		}
		return o.city

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *addressBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.id); err != nil {
		return nil, fmt.Errorf("error encoding Address.id: %w", err)
	}
	if err := encoder.Encode(o.idIsValid); err != nil {
		return nil, fmt.Errorf("error encoding Address.idIsValid: %w", err)
	}
	if err := encoder.Encode(o.idIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding Address.idIsDirty: %w", err)
	}

	if err := encoder.Encode(o.personID); err != nil {
		return nil, fmt.Errorf("error encoding Address.personID: %w", err)
	}
	if err := encoder.Encode(o.personIDIsValid); err != nil {
		return nil, fmt.Errorf("error encoding Address.personIDIsValid: %w", err)
	}
	if err := encoder.Encode(o.personIDIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding Address.personIDIsDirty: %w", err)
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
			return nil, fmt.Errorf("error encoding Address.objPerson: %w", err)
		}
	}

	if err := encoder.Encode(o.street); err != nil {
		return nil, fmt.Errorf("error encoding Address.street: %w", err)
	}
	if err := encoder.Encode(o.streetIsValid); err != nil {
		return nil, fmt.Errorf("error encoding Address.streetIsValid: %w", err)
	}
	if err := encoder.Encode(o.streetIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding Address.streetIsDirty: %w", err)
	}

	if err := encoder.Encode(o.city); err != nil {
		return nil, fmt.Errorf("error encoding Address.city: %w", err)
	}
	if err := encoder.Encode(o.cityIsNull); err != nil {
		return nil, fmt.Errorf("error encoding Address.cityIsNull: %w", err)
	}
	if err := encoder.Encode(o.cityIsValid); err != nil {
		return nil, fmt.Errorf("error encoding Address.cityIsValid: %w", err)
	}
	if err := encoder.Encode(o.cityIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding Address.cityIsDirty: %w", err)
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
			return nil, fmt.Errorf("error encoding Address._aliases: %w", err)
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, fmt.Errorf("error encoding Address._restored: %w", err)
	}
	if err := encoder.Encode(o._originalPK); err != nil {
		return nil, fmt.Errorf("error encoding Address._originalPK: %w", err)
	}

	return buf.Bytes(), nil

	return buf.Bytes(), nil
}

// UnmarshalBinary converts a structure that was created with MarshalBinary into a Address object.
func (o *addressBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var isPtr bool

	_ = isPtr
	if err = dec.Decode(&o.id); err != nil {
		return fmt.Errorf("error decoding Address.id: %w", err)
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return fmt.Errorf("error decoding Address.idIsValid: %w", err)
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return fmt.Errorf("error decoding Address.idIsDirty: %w", err)
	}

	if err = dec.Decode(&o.personID); err != nil {
		return fmt.Errorf("error decoding Address.personID: %w", err)
	}
	if err = dec.Decode(&o.personIDIsValid); err != nil {
		return fmt.Errorf("error decoding Address.personIDIsValid: %w", err)
	}
	if err = dec.Decode(&o.personIDIsDirty); err != nil {
		return fmt.Errorf("error decoding Address.personIDIsDirty: %w", err)
	}

	if err = dec.Decode(&isPtr); err != nil {
		return fmt.Errorf("error decoding Address.objPerson isPtr: %w", err)
	}
	if isPtr {
		if err = dec.Decode(&o.objPerson); err != nil {
			return fmt.Errorf("error decoding Address.objPerson: %w", err)
		}
	}
	if err = dec.Decode(&o.street); err != nil {
		return fmt.Errorf("error decoding Address.street: %w", err)
	}
	if err = dec.Decode(&o.streetIsValid); err != nil {
		return fmt.Errorf("error decoding Address.streetIsValid: %w", err)
	}
	if err = dec.Decode(&o.streetIsDirty); err != nil {
		return fmt.Errorf("error decoding Address.streetIsDirty: %w", err)
	}

	if err = dec.Decode(&o.city); err != nil {
		return fmt.Errorf("error decoding Address.city: %w", err)
	}
	if err = dec.Decode(&o.cityIsNull); err != nil {
		return fmt.Errorf("error decoding Address.cityIsNull: %w", err)
	}
	if err = dec.Decode(&o.cityIsValid); err != nil {
		return fmt.Errorf("error decoding Address.cityIsValid: %w", err)
	}
	if err = dec.Decode(&o.cityIsDirty); err != nil {
		return fmt.Errorf("error decoding Address.cityIsDirty: %w", err)
	}

	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *addressBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *addressBase) MarshalStringMap() map[string]interface{} {
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

	if o.streetIsValid {
		v["street"] = o.street
	}

	if o.cityIsValid {
		if o.cityIsNull {
			v["city"] = nil
		} else {
			v["city"] = o.city
		}
	}

	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the Address. The Address can be a
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
//	"street" - string
//	"city" - string, nullable
func (o *addressBase) UnmarshalJSON(data []byte) (err error) {
	var v map[string]interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	return o.UnmarshalStringMap(v)
}

// UnmarshalStringMap will load the values from the stringmap into the object.
//
// Override this in Address to modify the json before sending it here.
func (o *addressBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
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

		case "street":
			{
				if v == nil {
					return fmt.Errorf("json field %s cannot be null", k)
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetStreet(s)
				}
			}

		case "city":
			{
				if v == nil {
					o.SetCity(v)
					continue
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetCity(s)
				}
			}

		}
	}
	return
}
