// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
	"testing"

	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMilestone_SetProjectID(t *testing.T) {

	obj := NewMilestone()
	projectID := test.RandomValue[string](0)
	obj.SetProjectID(projectID)
	assert.Equal(t, projectID, obj.ProjectID())

	// test default
	obj.SetProjectID("")
	assert.EqualValues(t, "", obj.ProjectID(), "set default")

}
func TestMilestone_SetName(t *testing.T) {

	obj := NewMilestone()
	name := test.RandomValue[string](50)
	obj.SetName(name)
	assert.Equal(t, name, obj.Name())

	// test default
	obj.SetName("")
	assert.EqualValues(t, "", obj.Name(), "set default")

	// test panic on setting value larger than maximum size allowed
	name = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetName(name)
	})
}

// createMinimalSampleMilestone creates and saves a minimal version of a Milestone object
// for testing.
func createMinimalSampleMilestone(ctx context.Context) *Milestone {
	obj := NewMilestone()

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	objProject := createMinimalSampleProject(ctx)
	obj.SetProject(objProject)

	name := test.RandomValue[string](50)
	obj.SetName(name)

	obj.Save(ctx)
	return obj
}
func TestMilestone_CRUD(t *testing.T) {
	obj := NewMilestone()
	ctx := db.NewContext(nil)

	objProject := createMinimalSampleProject(ctx)
	defer objProject.Delete(ctx)
	obj.SetProject(objProject)

	name := test.RandomValue[string](50)
	obj.SetName(name)

	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer obj.Delete(ctx)

	// Test retrieval
	obj2 := LoadMilestone(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.ProjectIDIsValid())
	assert.NotEmpty(t, obj2.ProjectID())

	assert.True(t, obj2.NameIsValid())
	assert.Equal(t, name, obj2.Name())

}
