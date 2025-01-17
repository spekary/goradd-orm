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
	id        string
	idIsValid bool
	idIsDirty bool

	projectID        string
	projectIDIsValid bool
	projectIDIsDirty bool
	objProject       *Project

	name        string
	nameIsValid bool
	nameIsDirty bool

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

	newObjectPkCounter = newObjectPkCounter - 1
	o.id = fmt.Sprintf("%d", newObjectPkCounter)

	o.idIsValid = false
	o.idIsDirty = false

	o.projectID = ""

	o.projectIDIsValid = false
	o.projectIDIsDirty = false

	o.name = ""

	o.nameIsValid = false
	o.nameIsDirty = false

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

// Copy copies all valid fields to a new Milestone object.
// Forward reference ids will be copied, but reverse and many-many references will not.
// Attached objects will not be included in the copy.
// Call Save() on the new object to save it into the database.
// Copy might panic if any fields in the database were set to a size larger than the
// maximum size through a process that accessed the database outside of the ORM.
func (o *milestoneBase) Copy() (newObject *Milestone) {
	newObject = NewMilestone()
	if o.projectIDIsValid {
		newObject.SetProjectID(o.projectID)
	}
	if o.nameIsValid {
		newObject.SetName(o.name)
	}
	return
}

// ID returns the loaded value of ID or
// the zero value if not loaded. Call IDIsValid() to determine
// if it is loaded.
func (o *milestoneBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *milestoneBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

// ProjectID returns the loaded value of ProjectID.
func (o *milestoneBase) ProjectID() string {
	if o._restored && !o.projectIDIsValid {
		panic("ProjectID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.projectID
}

// ProjectIDIsValid returns true if the value was loaded from the database or has been set.
func (o *milestoneBase) ProjectIDIsValid() bool {
	return o.projectIDIsValid
}

// SetProjectID sets the value of ProjectID in the object, to be saved later using the Save() function.
func (o *milestoneBase) SetProjectID(projectID string) {
	o.projectIDIsValid = true

	if o.projectID != projectID || !o._restored {
		o.projectID = projectID
		o.projectIDIsDirty = true
		o.objProject = nil
	}
}

// Project returns the current value of the loaded Project, and nil if its not loaded.
func (o *milestoneBase) Project() *Project {
	return o.objProject
}

// LoadProject returns the related Project. If it is not already loaded,
// it will attempt to load it, provided the ProjectID column has been loaded first.
func (o *milestoneBase) LoadProject(ctx context.Context) *Project {
	if !o.projectIDIsValid {
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
		panic("Cannot set Project to a null value.")
	} else {
		o.objProject = objProject
		o.projectIDIsValid = true
		if o.projectID != objProject.PrimaryKey() {
			o.projectID = objProject.PrimaryKey()
			o.projectIDIsDirty = true
		}
	}
}

// Name returns the loaded value of Name.
func (o *milestoneBase) Name() string {
	if o._restored && !o.nameIsValid {
		panic("Name was not selected in the last query and has not been set, and so is not valid")
	}
	return o.name
}

// NameIsValid returns true if the value was loaded from the database or has been set.
func (o *milestoneBase) NameIsValid() bool {
	return o.nameIsValid
}

// SetName sets the value of Name in the object, to be saved later using the Save() function.
func (o *milestoneBase) SetName(name string) {
	o.nameIsValid = true

	if utf8.RuneCountInString(name) > MilestoneNameMaxLength {
		panic("attempted to set Milestone.Name to a value larger than its maximum length in runes")
	}
	if o.name != name || !o._restored {
		o.name = name
		o.nameIsDirty = true
	}
}

// GetAlias returns the alias for the given key.
func (o *milestoneBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
	}
}

// IsNew returns true if the object will create a new record when saved.
func (o *milestoneBase) IsNew() bool {
	return !o._restored
}

// LoadMilestone returns a Milestone from the database.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See [MilestonesBuilder.Join] and [MilestonesBuilder.Select] for more info.
func LoadMilestone(ctx context.Context, id string, joinOrSelectNodes ...query.Node) *Milestone {
	return queryMilestones(ctx).
		Where(op.Equal(node.Milestone().ID(), id)).
		joinOrSelect(joinOrSelectNodes...).
		Get()
}

// HasMilestone returns true if a Milestone with the given primaryKey exists in the database.
// doc: type=Milestone
func HasMilestone(ctx context.Context, id string) bool {
	return queryMilestones(ctx).
		Where(op.Equal(node.Milestone().ID(), id)).
		Count(false) == 1
}

// The MilestonesBuilder uses the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type MilestonesBuilder struct {
	builder query.QueryBuilderI
}

func newMilestoneBuilder(ctx context.Context) *MilestonesBuilder {
	b := &MilestonesBuilder{
		builder: db.GetDatabase("goradd").NewBuilder(ctx),
	}
	return b.Join(node.Milestone())
}

// Load terminates the query builder, performs the query, and returns a slice of Milestone objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *MilestonesBuilder) Load() (milestones []*Milestone) {
	results := b.builder.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Milestone)
		o.load(item, o, nil, "")
		milestones = append(milestones, o)
	}
	return
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *MilestonesBuilder) LoadI() (milestones []interface{}) {
	results := b.builder.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Milestone)
		o.load(item, o, nil, "")
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
func (b *MilestonesBuilder) LoadCursor() milestonesCursor {
	cursor := b.builder.LoadCursor()

	return milestonesCursor{cursor}
}

type milestonesCursor struct {
	query.CursorI
}

// Next returns the current Milestone object and moves the cursor to the next one.
//
// If there are no more records, it returns nil.
func (c milestonesCursor) Next() *Milestone {
	row := c.CursorI.Next()
	if row == nil {
		return nil
	}
	o := new(Milestone)
	o.load(row, o, nil, "")
	return o
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
func (b *MilestonesBuilder) Get() *Milestone {
	results := b.Load()
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Join adds node n to the node tree so that its fields will appear in the query.
// Optionally add conditions to filter what gets included.
func (b *MilestonesBuilder) Join(n query.Node, conditions ...query.Node) *MilestonesBuilder {
	if !query.NodeIsTableNodeI(n) {
		panic("you can only join Table, Reference, ReverseReference and ManyManyReference nodes")
	}

	if query.NodeTableName(query.RootNode(n)) != "milestone" {
		panic("you can only join a node that is rooted at node.Milestone()")
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
func (b *MilestonesBuilder) Where(c query.Node) *MilestonesBuilder {
	b.builder.Condition(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
func (b *MilestonesBuilder) OrderBy(nodes ...query.Node) *MilestonesBuilder {
	b.builder.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *MilestonesBuilder) Limit(maxRowCount int, offset int) *MilestonesBuilder {
	b.builder.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *MilestonesBuilder) Select(nodes ...query.Node) *MilestonesBuilder {
	b.builder.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using Alias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *MilestonesBuilder) Alias(name string, n query.Node) *MilestonesBuilder {
	b.builder.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *MilestonesBuilder) Distinct() *MilestonesBuilder {
	b.builder.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *MilestonesBuilder) GroupBy(nodes ...query.Node) *MilestonesBuilder {
	b.builder.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *MilestonesBuilder) Having(node query.Node) *MilestonesBuilder {
	b.builder.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
//
// distinct wll count the number of distinct items, ignoring duplicates.
//
// nodes will select individual fields, and should be accompanied by a GroupBy.
func (b *MilestonesBuilder) Count(distinct bool, nodes ...query.Node) uint {
	return b.builder.Count(distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *MilestonesBuilder) Delete() {
	b.builder.Delete()
	broadcast.BulkChange(b.builder.Context(), "goradd", "milestone")
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *MilestonesBuilder) Subquery() *query.SubqueryNode {
	return b.builder.Subquery()
}

// joinOrSelect is a private helper function for the Load* functions
func (b *MilestonesBuilder) joinOrSelect(nodes ...query.Node) *MilestonesBuilder {
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

// CountMilestoneByID queries the database and returns the number of Milestone objects that
// have id.
// doc: type=Milestone
func CountMilestoneByID(ctx context.Context, id string) int {
	return int(queryMilestones(ctx).Where(op.Equal(node.Milestone().ID(), id)).Count(false))
}

// CountMilestoneByProjectID queries the database and returns the number of Milestone objects that
// have projectID.
// doc: type=Milestone
func CountMilestoneByProjectID(ctx context.Context, projectID string) int {
	if projectID == "" {
		return 0
	}
	return int(queryMilestones(ctx).Where(op.Equal(node.Milestone().ProjectID(), projectID)).Count(false))
}

// CountMilestoneByName queries the database and returns the number of Milestone objects that
// have name.
// doc: type=Milestone
func CountMilestoneByName(ctx context.Context, name string) int {
	return int(queryMilestones(ctx).Where(op.Equal(node.Milestone().Name(), name)).Count(false))
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
func (o *milestoneBase) load(m map[string]interface{}, objThis *Milestone, objParent interface{}, parentKey string) {

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

	if v, ok := m["project_id"]; ok && v != nil {
		if o.projectID, ok = v.(string); ok {
			o.projectIDIsValid = true
			o.projectIDIsDirty = false

		} else {
			panic("Wrong type found for project_id.")
		}
	} else {
		o.projectIDIsValid = false
		o.projectID = ""
	}

	if v, ok := m["Project"]; ok {
		if objProject, ok2 := v.(map[string]interface{}); ok2 {
			o.objProject = new(Project)
			o.objProject.load(objProject, o.objProject, objThis, "Milestones")
			o.projectIDIsValid = true
			o.projectIDIsDirty = false
		} else {
			panic("Wrong type found for Project object.")
		}
	} else {
		o.objProject = nil
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
func (o *milestoneBase) Save(ctx context.Context) {
	if o._restored {
		o.update(ctx)
	} else {
		o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *milestoneBase) update(ctx context.Context) {
	if !o._restored {
		panic("cannot update a record that was not originally read from the database.")
	}

	var modifiedFields map[string]interface{}
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		// TODO: Perform all reads and consistency checks before saves

		// Save loaded Project object to get its new pk and update it here.
		if o.objProject != nil {
			o.objProject.Save(ctx)
			id := o.objProject.PrimaryKey()
			o.SetProjectID(id)
		}

		// Save all modified fields to the database
		modifiedFields = o.getModifiedFields()
		if len(modifiedFields) != 0 {
			d.Update(ctx, "milestone", modifiedFields, map[string]any{"id": o._originalPK})
		}

	}) // transaction

	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd", "milestone", o._originalPK, all.SortedKeys(modifiedFields)...)
	}
}

// insert will insert the object into the database. Related items will be saved.
func (o *milestoneBase) insert(ctx context.Context) {
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		if o.objProject != nil {
			o.objProject.Save(ctx)
			id := o.objProject.PrimaryKey()
			o.SetProjectID(id)
		}

		if !o.projectIDIsValid {
			panic("a value for ProjectID is required, and there is no default value. Call SetProjectID() before inserting the record.")
		}

		if !o.nameIsValid {
			panic("a value for Name is required, and there is no default value. Call SetName() before inserting the record.")
		}

		m := o.getValidFields()

		id := d.Insert(ctx, "milestone", m)
		o.id = id
		o._originalPK = id

	}) // transaction

	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "milestone", o.PrimaryKey())
}

// getModifiedFields returns the database columns that have been modified. This
// will determine which specific fields are sent to the database to be changed.
func (o *milestoneBase) getModifiedFields() (fields map[string]interface{}) {
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

// getValidFields returns the fields that have valid data in them.
func (o *milestoneBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}

	if o.projectIDIsValid {
		fields["project_id"] = o.projectID
	}

	if o.nameIsValid {
		fields["name"] = o.name
	}
	return
}

// Delete deletes the record from the database.
func (o *milestoneBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "milestone", map[string]any{"ID": o.id})
	broadcast.Delete(ctx, "goradd", "milestone", fmt.Sprint(o.id))
}

// deleteMilestone deletes the associated record from the database.
func deleteMilestone(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "milestone", map[string]any{"ID": pk})
	broadcast.Delete(ctx, "goradd", "milestone", fmt.Sprint(pk))
}

// resetDirtyStatus resets the dirty status of every field in the object.
func (o *milestoneBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.projectIDIsDirty = false
	o.nameIsDirty = false

}

// IsDirty returns true if the object has been changed since it was read from the database.
func (o *milestoneBase) IsDirty() (dirty bool) {
	dirty = o.idIsDirty ||
		o.projectIDIsDirty ||
		(o.objProject != nil && o.objProject.IsDirty()) ||
		o.nameIsDirty

	return
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *milestoneBase) Get(key string) interface{} {

	switch key {

	case "ID":
		if !o.idIsValid {
			return nil
		}
		return o.id

	case "ProjectID":
		if !o.projectIDIsValid {
			return nil
		}
		return o.projectID

	case "Project":
		return o.Project()

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
func (o *milestoneBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.id); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.id: %w", err)
	}
	if err := encoder.Encode(o.idIsValid); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.idIsValid: %w", err)
	}
	if err := encoder.Encode(o.idIsDirty); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.idIsDirty: %w", err)
	}

	if err := encoder.Encode(o.projectID); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.projectID: %w", err)
	}
	if err := encoder.Encode(o.projectIDIsValid); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.projectIDIsValid: %w", err)
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
	if err := encoder.Encode(o.nameIsValid); err != nil {
		return nil, fmt.Errorf("error encoding Milestone.nameIsValid: %w", err)
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
	if err = dec.Decode(&o.idIsValid); err != nil {
		return fmt.Errorf("error decoding Milestone.idIsValid: %w", err)
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return fmt.Errorf("error decoding Milestone.idIsDirty: %w", err)
	}

	if err = dec.Decode(&o.projectID); err != nil {
		return fmt.Errorf("error decoding Milestone.projectID: %w", err)
	}
	if err = dec.Decode(&o.projectIDIsValid); err != nil {
		return fmt.Errorf("error decoding Milestone.projectIDIsValid: %w", err)
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
	if err = dec.Decode(&o.nameIsValid); err != nil {
		return fmt.Errorf("error decoding Milestone.nameIsValid: %w", err)
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

	if o.idIsValid {
		v["id"] = o.id
	}

	if o.projectIDIsValid {
		v["projectID"] = o.projectID
	}

	if val := o.Project(); val != nil {
		v["project"] = val.MarshalStringMap()
	}

	if o.nameIsValid {
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
	if err = json.Unmarshal(data, &v); err != nil {
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

		case "projectID":
			{
				if v == nil {
					return fmt.Errorf("json field %s cannot be null", k)
				}

				if s, ok := v.(string); !ok {
					return fmt.Errorf("json field %s must be a string", k)
				} else {
					o.SetProjectID(s)
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
