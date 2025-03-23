// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
	"strconv"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMilestone_SetProjectID(t *testing.T) {

	obj := NewMilestone()
	val := test.RandomValue[string](0)
	obj.SetProjectID(val)
	assert.Equal(t, val, obj.ProjectID())

	// test default
	obj.SetProjectID("")
	assert.EqualValues(t, "", obj.ProjectID(), "set default")

}
func TestMilestone_SetName(t *testing.T) {

	obj := NewMilestone()
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

// createMinimalSampleMilestone creates an unsaved minimal version of a Milestone object
// for testing.
func createMinimalSampleMilestone() *Milestone {
	obj := NewMilestone()

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	// If the database is configured so that the referenced object points back here, either directly or through multiple
	// forward references, it possible this could create an endless loop. Such a structure could not be saved anyways.
	obj.SetProject(createMinimalSampleProject())

	obj.SetName(test.RandomValue[string](50))

	return obj
}

// createMaximalSampleMilestone creates an unsaved version of a Milestone object
// for testing that includes references to minimal objects.
func createMaximalSampleMilestone() *Milestone {
	obj := createMinimalSampleMilestone()

	return obj
}

// deleteSampleMilestone deletes an object created and saved by one of the sample creator functions.
func deleteSampleMilestone(ctx context.Context, obj *Milestone) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

	deleteSampleProject(ctx, obj.Project())

}

func TestMilestone_CRUD(t *testing.T) {
	obj := NewMilestone()
	ctx := db.NewContext(nil)

	v_objProject := createMinimalSampleProject()
	assert.NoError(t, v_objProject.Save(ctx))
	defer deleteSampleProject(ctx, v_objProject)
	obj.SetProject(v_objProject)

	v_name := test.RandomValue[string](50)
	obj.SetName(v_name)

	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer obj.Delete(ctx)

	// Test retrieval
	obj2 := LoadMilestone(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.ProjectIDIsValid())
	assert.NotEmpty(t, obj2.ProjectID())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.projectIDIsDirty)
	obj2.SetProjectID(obj2.ProjectID())
	assert.False(t, obj2.projectIDIsDirty)

	assert.True(t, obj2.NameIsValid())
	assert.EqualValues(t, v_name, obj2.Name())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.nameIsDirty)
	obj2.SetName(obj2.Name())
	assert.False(t, obj2.nameIsDirty)

}

func TestMilestone_References(t *testing.T) {
	obj := createMaximalSampleMilestone()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleMilestone(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.Project())
	assert.NotEqual(t, '-', obj.Project().PrimaryKey()[0])

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

	obj2 := LoadMilestone(ctx, obj.PrimaryKey(), node.Milestone().PrimaryKey())

	assert.Panics(t, func() { obj2.ProjectID() })
	assert.Panics(t, func() { obj2.Name() })
}
