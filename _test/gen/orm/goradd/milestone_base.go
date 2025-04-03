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

// MilestoneBase is embedded in a Milestone object and provides the ORM access to the database.
// The member variables of the structure are private and should not normally be accessed by the Milestone embedder.
// Instead, use the accessor functions.
type milestoneBase struct {
	id                string
	idIsLoaded        bool
	idIsDirty         bool
	projectID         string
	projectIDIsLoaded bool
	projectIDIsDirty  bool
	objProject        *Project
	name              string
	nameIsLoaded      bool
	nameIsDirty       bool

	// Custom aliases, if specified
	_aliases map[string]any

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update.
	_restored bool

	_originalPK string
}

// IDs used to access the Milestone object fields by name using the Get function.
// doc: type=Milestone
const (
	Milestone_ID        = `ID`
	Milestone_ProjectID = `ProjectID`
	Milestone_Project   = `Project`
	Milestone_Name      = `Name`
)

const MilestoneNameMaxLength = 50 // The number of runes the column can hold

// Initialize or re-initialize a Milestone database object to default values.
// The primary key will get a temporary negative number which will be replaced when the object is saved.
// Multiple calls to Initialize are not guaranteed to create sequential values for the primary key.
func (o *milestoneBase) Initialize() {
	o.id = db.TemporaryPrimaryKey()
	o.idIsLoaded = true
	o.idIsDirty = false

	o.projectID = ""
	o.projectIDIsLoaded = false
	o.projectIDIsDirty = false

	o.name = ""
	o.nameIsLoaded = false
	o.nameIsDirty = false

	o._aliases = nil
	o._restored = false
}

// PrimaryKey returns the current value of the primary key field.
func (o *milestoneBase) PrimaryKey() string {
	return o.id
}

// OriginalPrimaryKey returns the value of the primary key that was originally loaded into the object when it was
// read from the database.
func (o *milestoneBase) OriginalPrimaryKey() string {
	return o._originalPK
}

// Copy copies most fields to a new Milestone object.
// Forward reference ids will be copied, but reverse and many-many references will not.
// Attached objects will not be included in the copy.
// Automatically generated fields will not be included in the copy.
// The primary key field will not be copied, since it is normally auto-generated.
// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *milestoneBase) Copy() (newObject *Milestone) {
	newObject = NewMilestone()
	if o.idIsLoaded {
		newObject.SetID(o.id)
	}
	if o.projectIDIsLoaded {
		newObject.SetProjectID(o.projectID)
	}
	if o.nameIsLoaded {
		newObject.SetName(o.name)
	}
	return
}

// ID returns the value of ID.
func (o *milestoneBase) ID() string {
	if o._restored && !o.idIsLoaded {
		panic("ID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.id
}

// IDIsLoaded returns true if the value was loaded from the database or has been set.
func (o *milestoneBase) IDIsLoaded() bool {
	return o.idIsLoaded
}

// SetID sets the value of ID in the object, to be saved later in the database using the Save() function.
// Normally you will not need to call this function, since the ID value is automatically generated by the
// database driver. Exceptions might include importing data to a new database, or correcting primary key conflicts when
// merging data. You cannot change a primary key for a record that has been written to the database.
// (While SQL databases allow it, most NoSQL databases do not)
func (o *milestoneBase) SetID(v string) {
	if o._restored {
		panic("you cannot change an auto-generated primary key")
	}

	o.idIsLoaded = true
	o.id = v
	o.idIsDirty = true
}

// ProjectID returns the value of ProjectID.
func (o *milestoneBase) ProjectID() string {
	if o._restored && !o.projectIDIsLoaded {
		panic("ProjectID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.projectID
}

// ProjectIDIsLoaded returns true if the value was loaded from the database or has been set.
func (o *milestoneBase) ProjectIDIsLoaded() bool {
	return o.projectIDIsLoaded
}

// SetProjectID sets the value of ProjectID in the object, to be saved later in the database using the Save() function.
func (o *milestoneBase) SetProjectID(v string) {
	if o._restored &&
		o.projectIDIsLoaded && // if it was not selected, then make sure it gets set, since our end comparison won't be valid
		o.projectID == v {
		// no change
		return
	}

	o.projectIDIsLoaded = true
	o.projectID = v
	o.projectIDIsDirty = true
	o.objProject = nil
}

// Project returns the current value of the loaded Project, and nil if its not loaded.
func (o *milestoneBase) Project() *Project {
	return o.objProject
}

// LoadProject returns the related Project. If it is not already loaded,
// it will attempt to load it, provided the ProjectID column has been loaded first.
func (o *milestoneBase) LoadProject(ctx context.Context) *Project {
	if !o.projectIDIsLoaded {
		return nil
	}

	if o.objProject == nil {
		// Load and cache
		o.objProject = LoadProject(ctx, o.projectID)
	}
	return o.objProject
}

// SetProject sets the value of Project in the object, to be saved later using the Save() function.
func (o *milestoneBase) SetProject(objProject *Project) {
	if objProject == nil {
		panic("Cannot set Project to a nil value since ProjectID is not nullable.")
	} else {
		o.objProject = objProject
		o.projectIDIsLoaded = true
		if o.projectID != objProject.PrimaryKey() {
			o.projectID = objProject.PrimaryKey()
			o.projectIDIsDirty = true
		}
	}
}

// Name returns the value of Name.
func (o *milestoneBase) Name() string {
	if o._restored && !o.nameIsLoaded {
		panic("Name was not selected in the last query and has not been set, and so is not valid")
	}
	return o.name
}

// NameIsLoaded returns true if the value was loaded from the database or has been set.
func (o *milestoneBase) NameIsLoaded() bool {
	return o.nameIsLoaded
}

// SetName sets the value of Name in the object, to be saved later in the database using the Save() function.
func (o *milestoneBase) SetName(v string) {
	if utf8.RuneCountInString(v) > MilestoneNameMaxLength {
		panic("attempted to set Milestone.Name to a value larger than its maximum length in runes")
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

// GetAlias returns the value for the Alias node aliasKey that was returned in the most
// recent query.
func (o *milestoneBase) GetAlias(aliasKey string) query.AliasValue {
	if a, ok := o._aliases[aliasKey]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + aliasKey + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *milestoneBase) IsNew() bool {
	return !o._restored
}

// LoadMilestone returns a Milestone from the database.
// selectNodes lets you provide nodes for selecting specific fields or additional fields from related tables.
// See [MilestonesBuilder.Select] for more info.
func LoadMilestone(ctx context.Context, id string, selectNodes ...query.Node) *Milestone {
	return queryMilestones(ctx).
		Where(op.Equal(node.Milestone().ID(), id)).
		Select(selectNodes...).
		Get()
}

// HasMilestone returns true if a Milestone with the given primary key exists in the database.
// doc: type=Milestone
func HasMilestone(ctx context.Context, id string) bool {
	return queryMilestones(ctx).
		Where(op.Equal(node.Milestone().ID(), id)).
		Count() == 1
}

// The MilestoneBuilder uses the query.BuilderI interface to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, LoadCursor, Get, Count, or Delete
type MilestoneBuilder interface {
	// Join(alias string, joinedTable query.Node, condition query.Node) MilestoneBuilder

	// Where adds a condition to filter what gets selected.
	// Calling Where multiple times will AND the conditions together.
	Where(c query.Node) MilestoneBuilder

	// OrderBy specifies how the resulting data should be sorted.
	// By default, the given nodes are sorted in ascending order.
	// Add Descending() to the node to specify that it should be sorted in descending order.
	OrderBy(nodes ...query.Sorter) MilestoneBuilder

	// Limit will return a subset of the data, limited to the offset and number of rows specified.
	// For large data sets and specific types of queries, this can be slow, because it will perform
	// the entire query before computing the limit.
	// You cannot limit a query that has selected a "many" relationship".
	Limit(maxRowCount int, offset int) MilestoneBuilder

	// Select performs two functions:
	//  - Passing a table type node will join the object or objects from that table to this object.
	//  - Passing a column node will optimize the query to only return the specified fields.
	// Once you select at least one column, you must select all the columns that you want in the result.
	// Some fields, like primary keys, are always selected.
	// If you are using a GroupBy, you must select the fields in the GroupBy.
	Select(nodes ...query.Node) MilestoneBuilder

	// Calculation adds a calculation described by operation with the name alias.
	// After the query, you can read the data using GetAlias() on the object identified by base.
	Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) MilestoneBuilder

	// Distinct removes duplicates from the results of the query.
	// Adding a Select() is required.
	Distinct() MilestoneBuilder

	// GroupBy controls how results are grouped when using aggregate functions with Calculation.
	GroupBy(nodes ...query.Node) MilestoneBuilder

	// Having does additional filtering on the results of the query after the query is performed.
	Having(node query.Node) MilestoneBuilder

	// Load terminates the query builder, performs the query, and returns a slice of Milestone objects.
	// If there are any errors, nil is returned and the specific error is stored in the context.
	// If no results come back from the query, it will return a non-nil empty slice.
	Load() []*Milestone
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
	LoadCursor() milestonesCursor

	// Get is a convenience method to return only the first item found in a query.
	// The entire query is performed, so you should generally use this only if you know
	// you are selecting on one or very few items.
	//
	// If an error occurs, or no results are found, a nil is returned.
	// In the case of an error, the error is returned in the context.
	Get() *Milestone

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

type milestoneQueryBuilder struct {
	builder *query.Builder
}

func newMilestoneBuilder(ctx context.Context) MilestoneBuilder {
	b := milestoneQueryBuilder{
		builder: query.NewBuilder(ctx, node.Milestone()),
	}
	return &b
}

// Load terminates the query builder, performs the query, and returns a slice of Milestone objects.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *milestoneQueryBuilder) Load() (milestones []*Milestone) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(Milestone)
		o.load(item, o)
		milestones = append(milestones, o)
	}
	return
}

// Load terminates the query builder, performs the query, and returns a slice of interfaces.
// This can then satisfy a variety of interfaces that load arrays of objects, including KeyLabeler.
// If there are any errors, nil is returned and the specific error is stored in the context.
// If no results come back from the query, it will return a non-nil empty slice.
func (b *milestoneQueryBuilder) LoadI() (milestones []query.OrmObj) {
	b.builder.Command = query.BuilderCommandLoad
	database := db.GetDatabase("goradd")
	results := database.BuilderQuery(b.builder)
	if results == nil {
		return
	}
	for _, item := range results.([]map[string]any) {
		o := new(Milestone)
		o.load(item, o)
		milestones = append(milestones, o)
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
func (b *milestoneQueryBuilder) LoadCursor() milestonesCursor {
	b.builder.Command = query.BuilderCommandLoadCursor
	database := db.GetDatabase("goradd")
	result := database.BuilderQuery(b.builder)
	cursor := result.(query.CursorI)

	return milestonesCursor{cursor}
}

type milestonesCursor struct {
	query.CursorI
}

// Next returns the current Milestone object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c milestonesCursor) Next() *Milestone {
	if c.CursorI == nil {
		return nil
	}

	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(Milestone)
	o.load(row, o)
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
//
// If an error occurs, or no results are found, a nil is returned.
// In the case of an error, the error is returned in the context.
func (b *milestoneQueryBuilder) Get() *Milestone {
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
func (b *milestoneQueryBuilder) Join(alias string, joinedTable query.Node, condition query.Node) MilestoneBuilder {
    if query.RootNode(n).TableName_() != "milestone" {
        panic("you can only join a node that is rooted at node.Milestone()")
    }
    // TODO: make sure joinedTable is a table node
	b.builder.Join(alias, joinedTable, condition)
	return b
}
*/

// Where adds a condition to filter what gets selected.
// Calling Where multiple times will AND the conditions together.
func (b *milestoneQueryBuilder) Where(c query.Node) MilestoneBuilder {
	b.builder.Where(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
// By default, the given nodes are sorted in ascending order.
// Add Descending() to the node to specify that it should be sorted in descending order.
func (b *milestoneQueryBuilder) OrderBy(nodes ...query.Sorter) MilestoneBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified.
// For large data sets and specific types of queries, this can be slow, because it will perform
// the entire query before computing the limit.
// You cannot limit a query that has embedded arrays.
func (b *milestoneQueryBuilder) Limit(maxRowCount int, offset int) MilestoneBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select specifies what specific columns will be loaded with data.
// By default, all the columns of the milestone table will be queried and loaded.
// If nodes contains columns from the milestone table, that will limit the columns queried and loaded to only those columns.
// If related tables are specified, then all the columns from those tables are queried, selected and joined to the result.
// If columns in related tables are specified, then only those columns will be queried and loaded.
// Depending on the query, additional columns may automatically be added to the query. In particular, primary key columns
// will be added in most situations. The exception to this would be in distinct queries, group by queries, or subqueries.
func (b *milestoneQueryBuilder) Select(nodes ...query.Node) MilestoneBuilder {
	b.builder.Select(nodes...)
	return b
}

// Calculation adds operation as an aliased value onto base.
// After the query, you can read the data by passing alias to GetAlias on the returned object.
func (b *milestoneQueryBuilder) Calculation(base query.TableNodeI, alias string, operation query.OperationNodeI) MilestoneBuilder {
	b.builder.Calculation(base, alias, operation)
	return b
}

// Distinct removes duplicates from the results of the query.
// Adding a Select() is usually required.
func (b *milestoneQueryBuilder) Distinct() MilestoneBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions with Calculation.
func (b *milestoneQueryBuilder) GroupBy(nodes ...query.Node) MilestoneBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query after the query is performed.
func (b *milestoneQueryBuilder) Having(node query.Node) MilestoneBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items in the result.
// If you have Select or Calculation columns in the query, it will count NULL results as well.
// To not count NULL values, use Where in the builder with a NotNull operation.
// To count distinct combinations of items, call Distinct() on the builder.
func (b *milestoneQueryBuilder) Count() int {
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
func (b *milestoneQueryBuilder)  Subquery() *query.SubqueryNode {
	 return b.builder.Subquery()
}
*/

func CountMilestones(ctx context.Context) int {
	return QueryMilestones(ctx).Count()
}

// CountMilestonesByID queries the database and returns the number of Milestone objects that
// have id.
// doc: type=Milestone
func CountMilestonesByID(ctx context.Context, id string) int {
	return QueryMilestones(ctx).Where(op.Equal(node.Milestone().ID(), id)).Count()
}

// CountMilestonesByProjectID queries the database and returns the number of Milestone objects that
// have projectID.
// doc: type=Milestone
func CountMilestonesByProjectID(ctx context.Context, projectID string) int {
	if projectID == "" {
		return 0
	}
	return QueryMilestones(ctx).Where(op.Equal(node.Milestone().ProjectID(), projectID)).Count()
}

// CountMilestonesByName queries the database and returns the number of Milestone objects that
// have name.
// doc: type=Milestone
func CountMilestonesByName(ctx context.Context, name string) int {
	return QueryMilestones(ctx).Where(op.Equal(node.Milestone().Name(), name)).Count()
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
func (o *milestoneBase) load(m map[string]interface{}, objThis *Milestone) {

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

	if v, ok := m["project_id"]; ok && v != nil {
		if o.projectID, ok = v.(string); ok {
			o.projectIDIsLoaded = true
			o.projectIDIsDirty = false

		} else {
			panic("Wrong type found for project_id.")
		}
	} else {
		o.projectIDIsLoaded = false
		o.projectID = ""
		o.projectIDIsDirty = false
	}

	if v, ok := m["Project"]; ok {
		if objProject, ok2 := v.(map[string]any); ok2 {
			o.objProject = new(Project)
			o.objProject.load(objProject, o.objProject)
			o.projectIDIsLoaded = true
			o.projectIDIsDirty = false
		} else {
			panic("Wrong type found for Project object.")
		}
	} else {
		o.objProject = nil
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

	if v, ok := m["aliases_"]; ok {
		o._aliases = v.(map[string]any)
	}

	o._restored = true

}

// save will update or insert the object, depending on the state of the object.
func (o *milestoneBase) save(ctx context.Context) error {
	if o._restored {
		return o.update(ctx)
	} else {
		return o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
// If the table has auto-generated values, those will be updated automatically.
func (o *milestoneBase) update(ctx context.Context) error {
	if !o._restored {
		panic("cannot update a record that was not originally read from the database.")
	}
	if !o.IsDirty() {
		return nil // nothing to save
	}

	var modifiedFields map[string]interface{}

	d := Database()
	err := db.ExecuteTransaction(ctx, d, func() error {

		// Save loaded Project object to get its new pk and update it here.
		if o.objProject != nil {
			if err := o.objProject.Save(ctx); err != nil {
				return err
			}
			o.projectID = o.objProject.PrimaryKey()
		}

		modifiedFields = o.getUpdateFields()
		if len(modifiedFields) != 0 {
			var err2 error

			_, err2 = d.Update(ctx, "milestone", "id", o._originalPK, modifiedFields, "", 0)
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
		broadcast.Update(ctx, "goradd", "milestone", o._originalPK, all.SortedKeys(modifiedFields)...)
	}

	return nil
}

// insert will insert the object into the database. Related items will be saved.
func (o *milestoneBase) insert(ctx context.Context) (err error) {
	var insertFields map[string]interface{}
	d := Database()
	err = db.ExecuteTransaction(ctx, d, func() error {

		// Save loaded Project object to get its new pk and update it here.
		if o.objProject != nil {
			if err := o.objProject.Save(ctx); err != nil {
				return err
			}
			o.projectID = o.objProject.PrimaryKey()
		}

		if !o.projectIDIsLoaded {
			panic("a value for ProjectID is required, and there is no default value. Call SetProjectID() before inserting the record.")
		}
		if !o.nameIsLoaded {
			panic("a value for Name is required, and there is no default value. Call SetName() before inserting the record.")
		}

		insertFields = o.getInsertFields()
		var newPk string

		newPk, err = d.Insert(ctx, "milestone", "id", insertFields)
		if err != nil {
			return err
		}
		o.id = newPk
		o._originalPK = newPk
		o.idIsLoaded = true

		return nil

	}) // transaction

	if err != nil {
		return
	}

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "milestone", o.PrimaryKey())
	return
}

// getUpdateFields returns the database columns that will be sent to the update process.
// This will include timestamp fields only if some other column has changed.
func (o *milestoneBase) getUpdateFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.projectIDIsDirty {
		fields["project_id"] = o.projectID
	}
	if o.nameIsDirty {
		fields["name"] = o.name
	}
	return
}

// getInsertFields returns the fields that will be specified in an insert operation.
// Optional fields that have not been set and have no default will be returned as nil.
// NoSql databases should interpret this as no value. Sql databases should interpret this as
// explicitly setting a NULL value, which would override any database specific default value.
// Auto-generated fields will be returned with their generated values, except AutoPK fields, which are generated by the
// database driver and updated after the insert.
func (o *milestoneBase) getInsertFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}

	fields["project_id"] = o.projectID

	fields["name"] = o.name
	return
}

// Delete deletes the record from the database.
func (o *milestoneBase) Delete(ctx context.Context) (err error) {
	if o == nil {
		return // allow deleting of a nil object to be a noop
	}
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "milestone", map[string]any{"ID": o.id})
	return nil
	broadcast.Delete(ctx, "goradd", "milestone", fmt.Sprint(o.id))
	return
}

// deleteMilestone deletes the Milestone with primary key pk from the database
// and handles associated records.
func deleteMilestone(ctx context.Context, pk string) error {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "milestone", map[string]any{"ID": pk})
	broadcast.Delete(ctx, "goradd", "milestone", fmt.Sprint(pk))
	return nil
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *milestoneBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.projectIDIsDirty = false
	o.nameIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database or created.
// However, a new object that has a column with a default value will be automatically marked as dirty upon creation.
func (o *milestoneBase) IsDirty() (dirty bool) {
	dirty = o.idIsDirty ||
		o.projectIDIsDirty ||
		(o.objProject != nil && o.objProject.IsDirty()) ||
		o.nameIsDirty

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil.
// Get can be used to retrieve a value by using the Identifier of a node.
func (o *milestoneBase) Get(key string) interface{} {

	switch key {

	case "ID":
		if !o.idIsLoaded {
			return nil
		}
		return o.id

	case "ProjectID":
		if !o.projectIDIsLoaded {
			return nil
		}
		return o.projectID

	case "Project":
		return o.Project()

	case "Name":
		if !o.nameIsLoaded {
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
func (o *milestoneBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.id); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.id: %w", err)
	}
	if err := encoder.Encode(o.idIsLoaded); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.idIsLoaded: %w", err)
	}
	if err := encoder.Encode(o.idIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.idIsDirty: %w", err)
	}

	if err := encoder.Encode(o.projectID); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.projectID: %w", err)
	}
	if err := encoder.Encode(o.projectIDIsLoaded); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.projectIDIsLoaded: %w", err)
	}
	if err := encoder.Encode(o.projectIDIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.projectIDIsDirty: %w", err)
	}

	if o.objProject == nil {
		if err := encoder.Encode(false); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(true); err != nil {
			return nil, err
		}
		if err := encoder.Encode(o.objProject); err != nil {
			return nil, fmt.Errorf("error encoding Milestone.objProject: %w", err)
		}
	}

	if err := encoder.Encode(o.name); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.name: %w", err)
	}
	if err := encoder.Encode(o.nameIsLoaded); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.nameIsLoaded: %w", err)
	}
	if err := encoder.Encode(o.nameIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.nameIsDirty: %w", err)
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
			return nil, fmt.Errorf("error encoding Milestone._aliases: %w", err)
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, fmt.Errorf("error encoding Milestone._restored: %w", err)
	}
	if err := encoder.Encode(o._originalPK); err != nil {
		return nil, fmt.Errorf("error encoding Milestone._originalPK: %w", err)
	}

	return buf.Bytes(), nil

	return buf.Bytes(), nil
}

// UnmarshalBinary converts a structure that was created with MarshalBinary into a Milestone object.
func (o *milestoneBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var isPtr bool

	_ = isPtr
	if err = dec.Decode(&o.id); err != nil {
		return fmt.Errorf("error decoding Milestone.id: %w", err)
	}
	if err = dec.Decode(&o.idIsLoaded); err != nil {
		return fmt.Errorf("error decoding Milestone.idIsLoaded: %w", err)
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return fmt.Errorf("error decoding Milestone.idIsDirty: %w", err)
	}

	if err = dec.Decode(&o.projectID); err != nil {
		return fmt.Errorf("error decoding Milestone.projectID: %w", err)
	}
	if err = dec.Decode(&o.projectIDIsLoaded); err != nil {
		return fmt.Errorf("error decoding Milestone.projectIDIsLoaded: %w", err)
	}
	if err = dec.Decode(&o.projectIDIsDirty); err != nil {
		return fmt.Errorf("error decoding Milestone.projectIDIsDirty: %w", err)
	}

	if err = dec.Decode(&isPtr); err != nil {
		return fmt.Errorf("error decoding Milestone.objProject isPtr: %w", err)
	}
	if isPtr {
		if err = dec.Decode(&o.objProject); err != nil {
			return fmt.Errorf("error decoding Milestone.objProject: %w", err)
		}
	}
	if err = dec.Decode(&o.name); err != nil {
		return fmt.Errorf("error decoding Milestone.name: %w", err)
	}
	if err = dec.Decode(&o.nameIsLoaded); err != nil {
		return fmt.Errorf("error decoding Milestone.nameIsLoaded: %w", err)
	}
	if err = dec.Decode(&o.nameIsDirty); err != nil {
		return fmt.Errorf("error decoding Milestone.nameIsDirty: %w", err)
	}

	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. Another way to control the output
// is to call MarshalStringMap, modify the map, then encode the map.
func (o *milestoneBase) MarshalJSON() (data []byte, err error) {
	v := o.MarshalStringMap()
	return json.Marshal(v)
}

// MarshalStringMap serializes the object into a string map of interfaces.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object. The keys are the same as the json keys.
func (o *milestoneBase) MarshalStringMap() map[string]interface{} {
	v := make(map[string]interface{})

	if o.idIsLoaded {
		v["id"] = o.id
	}

	if val := o.objProject; val != nil {
		v["project"] = val.MarshalStringMap()
	} else if o.projectIDIsLoaded {
		v["projectID"] = o.projectID
	}

	if o.nameIsLoaded {
		v["name"] = o.name
	}

	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return v
}

// UnmarshalJSON unmarshalls the given json data into the Milestone. The Milestone can be a
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
//	"projectID" - string
//	"name" - string
func (o *milestoneBase) UnmarshalJSON(data []byte) (err error) {
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
// Override this in Milestone to modify the json before sending it here.
func (o *milestoneBase) UnmarshalStringMap(m map[string]interface{}) (err error) {
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

		case "projectID":
			{
				if v == nil {
					return fmt.Errorf("field %s cannot be null", k)
				}

				if _, ok := m["project"]; ok {
					continue // importing the foreign key will remove the object
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetProjectID(s)
				}
			}

		case "project":
			v2 := NewProject()
			m2, ok := v.(map[string]any)
			if !ok {
				return fmt.Errorf("json field %s must be a map", k)
			}
			err = v2.UnmarshalStringMap(m2)
			if err != nil {
				return
			}
			o.SetProject(v2)

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
