// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createMinimalSampleMilestone creates an unsaved minimal version of a Milestone object
// for testing.
func createMinimalSampleMilestone() *Milestone {
	obj := NewMilestone()
	updateMinimalSampleMilestone(obj)

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	// If the database is configured so that the referenced object points back here, either directly or through multiple
	// forward references, it possible this could create an endless loop. Such a structure could not be saved anyways.
	obj.SetProject(createMinimalSampleProject())

	return obj
}

// updateMinimalSampleMilestone sets the values of a minimal sample to new, random values.
func updateMinimalSampleMilestone(obj *Milestone) {

	obj.SetName(test.RandomValue[string](50))

}

// createMaximalSampleMilestone creates an unsaved version of a Milestone object
// for testing that includes references to minimal objects.
func createMaximalSampleMilestone(ctx context.Context) *Milestone {
	obj := NewMilestone()
	updateMaximalSampleMilestone(ctx, obj)
	return obj
}

// updateMaximalSampleMilestone sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleMilestone(ctx context.Context, obj *Milestone) {
	updateMinimalSampleMilestone(obj)
	obj.SetProject(createMinimalSampleProject())

}

// deleteSampleMilestone deletes an object created and saved by one of the sample creator functions.
func deleteSampleMilestone(ctx context.Context, obj *Milestone) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

	deleteSampleProject(ctx, obj.Project())

}

// assertEqualFieldsMilestone compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsMilestone(t *testing.T, obj1, obj2 *Milestone) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.ProjectIDIsLoaded() && obj2.ProjectIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ProjectID(), obj2.ProjectID())
	}
	if obj1.NameIsLoaded() && obj2.NameIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Name(), obj2.Name())
	}

}

func TestMilestone_SetID(t *testing.T) {

	obj := NewMilestone()

	assert.True(t, obj.IsNew())
	val := test.RandomNumberString()
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID("")
	assert.EqualValues(t, "", obj.ID(), "set default")

}
func TestMilestone_SetProjectID(t *testing.T) {

	obj := NewMilestone()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetProjectID(val)
	assert.Equal(t, val, obj.ProjectID())

	// test default
	obj.SetProjectID("")
	assert.EqualValues(t, "", obj.ProjectID(), "set default")

}
func TestMilestone_SetName(t *testing.T) {

	obj := NewMilestone()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](50)
	obj.SetName(val)
	assert.Equal(t, val, obj.Name())

	// test default
	obj.SetName("")
	assert.EqualValues(t, "", obj.Name(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetName(val)
	})
}

func TestMilestone_Copy(t *testing.T) {
	obj := createMinimalSampleMilestone()

	obj2 := obj.Copy()

	assert.Equal(t, obj.ProjectID(), obj2.ProjectID())
	assert.Equal(t, obj.Name(), obj2.Name())

}

func TestMilestone_BasicInsert(t *testing.T) {
	obj := createMinimalSampleMilestone()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleMilestone(ctx, obj)

	// Test retrieval
	obj2 := LoadMilestone(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsLoaded())
	assert.Panics(t, func() {
		obj2.SetID(obj2.ID())
	})

	assert.True(t, obj2.NameIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.nameIsDirty)
	obj2.SetName(obj2.Name())
	assert.False(t, obj2.nameIsDirty)

}

func TestMilestone_InsertPanics(t *testing.T) {
	obj := createMinimalSampleMilestone()
	_ = obj
	ctx := db.NewContext(nil)
	_ = ctx

	obj.projectIDIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.projectIDIsLoaded = true

	obj.nameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsLoaded = true

}

func TestMilestone_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleMilestone()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleMilestone(ctx, obj)
	updateMinimalSampleMilestone(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadMilestone(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.Name(), obj.Name(), "Name did not update")
}

func TestMilestone_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleMilestone(ctx)
	obj.Save(ctx)
	defer deleteSampleMilestone(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.Project())
	assert.NotEqual(t, '-', obj.Project().PrimaryKey()[0])

	// Test lazy loading
	obj2 := LoadMilestone(ctx, obj.PrimaryKey())
	objPkOnly := LoadMilestone(ctx, obj.PrimaryKey(), node.Milestone().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.Project(), "Project is not loaded initially")
	v_Project := obj2.LoadProject(ctx)
	assert.NotNil(t, v_Project)
	assert.Equal(t, v_Project.PrimaryKey(), obj2.Project().PrimaryKey())
	assert.Equal(t, obj.Project().PrimaryKey(), obj2.Project().PrimaryKey())
	assert.True(t, obj2.ProjectIDIsLoaded())

	assert.False(t, objPkOnly.ProjectIDIsLoaded())
	assert.Nil(t, objPkOnly.LoadProject(ctx))

	assert.Panics(t, func() {
		objPkOnly.SetProject(nil)
	})

	// test eager loading
	obj3 := LoadMilestone(ctx, obj.PrimaryKey(), node.Milestone().Project())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Project().PrimaryKey(), obj3.Project().PrimaryKey())

}

func TestMilestone_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleMilestone(ctx)
	obj.Save(ctx)
	defer deleteSampleMilestone(ctx, obj)

	obj2 := LoadMilestone(ctx, obj.PrimaryKey())
	updateMaximalSampleMilestone(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleMilestone(ctx, obj2)

	obj3 := LoadMilestone(ctx, obj2.PrimaryKey(), node.Milestone().Project())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Project().PrimaryKey(), obj3.Project().PrimaryKey())

}

func TestMilestone_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleMilestone(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleMilestone(ctx, obj)

	updateMinimalSampleProject(obj.Project())

	assert.NoError(t, obj.Save(ctx))

	obj2 := LoadMilestone(ctx, obj.PrimaryKey(),
		node.Milestone().Project(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsProject(t, obj2.Project(), obj.Project())

}
func TestMilestone_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewMilestone()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestMilestone_Getters(t *testing.T) {
	obj := createMinimalSampleMilestone()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleMilestone(ctx, obj)

	assert.True(t, HasMilestone(ctx, obj.PrimaryKey()))

	obj2 := LoadMilestone(ctx, obj.PrimaryKey(), node.Milestone().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.Milestone().ID().Identifier))
	assert.Panics(t, func() { obj2.ProjectID() })
	assert.Nil(t, obj2.Get(node.Milestone().ProjectID().Identifier))
	assert.Equal(t, obj.Name(), obj.Get(node.Milestone().Name().Identifier))
	assert.Panics(t, func() { obj2.Name() })
	assert.Nil(t, obj2.Get(node.Milestone().Name().Identifier))
}

func TestMilestone_QueryLoad(t *testing.T) {
	obj := createMinimalSampleMilestone()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleMilestone(ctx, obj)

	objs := QueryMilestones(ctx).
		Where(op.Equal(node.Milestone().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.Milestone().PrimaryKey()). // exercise order by
		Limit(1, 0).                            // exercise limit
		Calculation(node.Milestone(), "IsTrue", op.Equal(1, 1)).
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestMilestone_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleMilestone()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleMilestone(ctx, obj)

	objs := QueryMilestones(ctx).
		Where(op.Equal(node.Milestone().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestMilestone_QueryCursor(t *testing.T) {
	obj := createMinimalSampleMilestone()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleMilestone(ctx, obj)

	cursor := QueryMilestones(ctx).
		Where(op.Equal(node.Milestone().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryMilestones(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestMilestone_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleMilestone(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	// reread in case there are data limitations imposed by the database
	obj2 := LoadMilestone(ctx, obj.PrimaryKey())
	defer deleteSampleMilestone(ctx, obj)

	assert.Less(t, 0, CountMilestones(ctx))

	assert.Less(t, 0, CountMilestonesByID(ctx, obj2.ID()))
	assert.Less(t, 0, CountMilestonesByProjectID(ctx, obj2.ProjectID()))
	assert.Less(t, 0, CountMilestonesByName(ctx, obj2.Name()))

}
func TestMilestone_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleMilestone()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewMilestone()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsMilestone(t, obj, obj2)
}

func TestMilestone_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleMilestone()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewMilestone()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsMilestone(t, obj, obj2)
}

func TestMilestone_FailingMarshalBinary(t *testing.T) {
	obj := createMinimalSampleMilestone()
	var err error

	for i := 0; i < 12; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
	// do it again with aliases
	obj._aliases = make(map[string]any)
	for i := 0; i < 13; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
}

func TestMilestone_FailingUnmarshalBinary(t *testing.T) {
	obj := createMinimalSampleMilestone()
	b, err := obj.MarshalBinary()
	assert.NoError(t, err)
	obj2 := NewMilestone()
	for i := 0; i < 12; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}

	// do it again with aliases
	obj = createMinimalSampleMilestone()
	obj._aliases = map[string]any{"a": 1}
	b, err = obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 = NewMilestone()
	for i := 0; i < 13; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}
}
