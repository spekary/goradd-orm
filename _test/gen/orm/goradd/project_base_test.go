// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createMinimalSampleProject creates an unsaved minimal version of a Project object
// for testing.
func createMinimalSampleProject() *Project {
	obj := NewProject()
	updateMinimalSampleProject(obj)
	return obj
}

// updateMinimalSampleProject sets the values of a minimal sample to new, random values.
func updateMinimalSampleProject(obj *Project) {

	obj.SetNum(test.RandomValue[int](32))

	obj.SetStatus(test.RandomEnum(ProjectStatuses()))

	obj.SetName(test.RandomValue[string](100))

	obj.SetDescription(test.RandomValue[string](65535))

	obj.SetStartDate(test.RandomValue[time.Time](0))

	obj.SetEndDate(test.RandomValue[time.Time](0))

	obj.SetBudget(test.RandomDecimal(12, 2))

	obj.SetSpent(test.RandomDecimal(12, 2))

}

// createMaximalSampleProject creates an unsaved version of a Project object
// for testing that includes references to minimal objects.
func createMaximalSampleProject() *Project {
	obj := NewProject()
	updateMaximalSampleProject(obj)
	return obj
}

// updateMaximalSampleProject sets all the maximal sample values to new values.
func updateMaximalSampleProject(obj *Project) {
	updateMinimalSampleProject(obj)

	obj.SetManager(createMinimalSamplePerson())

	obj.SetParentProject(createMinimalSampleProject())

	obj.SetMilestones(createMinimalSampleMilestone())
	obj.SetParentProjectProjects(createMinimalSampleProject())
	obj.SetChildren(createMinimalSampleProject())
	obj.SetParents(createMinimalSampleProject())
	obj.SetTeamMembers(createMinimalSamplePerson())
}

// deleteSampleProject deletes an object created and saved by one of the sample creator functions.
func deleteSampleProject(ctx context.Context, obj *Project) {
	if obj == nil {
		return
	}

	for _, item := range obj.Milestones() {
		deleteSampleMilestone(ctx, item)
	}
	for _, item := range obj.ParentProjectProjects() {
		deleteSampleProject(ctx, item)
	}

	for _, item := range obj.Children() {
		deleteSampleProject(ctx, item)
	}
	for _, item := range obj.Parents() {
		deleteSampleProject(ctx, item)
	}
	for _, item := range obj.TeamMembers() {
		deleteSamplePerson(ctx, item)
	}

	obj.Delete(ctx)

	deleteSamplePerson(ctx, obj.Manager())

	deleteSampleProject(ctx, obj.ParentProject())

}

func TestProject_SetNum(t *testing.T) {

	obj := NewProject()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[int](32)
	obj.SetNum(val)
	assert.Equal(t, val, obj.Num())

	// test default
	obj.SetNum(0)
	assert.EqualValues(t, 0, obj.Num(), "set default")

}
func TestProject_SetStatus(t *testing.T) {

	obj := NewProject()

	assert.True(t, obj.IsNew())

	val := test.RandomEnum(ProjectStatuses())
	obj.SetStatus(val)

	assert.Equal(t, val, obj.Status())

	// test default
	obj.SetStatus(0)
	assert.EqualValues(t, 0, obj.Status(), "set default")

}
func TestProject_SetManagerID(t *testing.T) {

	obj := NewProject()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetManagerID(val)
	assert.Equal(t, val, obj.ManagerID())
	assert.False(t, obj.ManagerIDIsNull())

	// Test NULL
	obj.SetManagerIDToNull()
	assert.EqualValues(t, "", obj.ManagerID())
	assert.True(t, obj.ManagerIDIsNull())

	// test default
	obj.SetManagerID("")
	assert.EqualValues(t, "", obj.ManagerID(), "set default")

}
func TestProject_SetName(t *testing.T) {

	obj := NewProject()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](100)
	obj.SetName(val)
	assert.Equal(t, val, obj.Name())

	// test default
	obj.SetName("")
	assert.EqualValues(t, "", obj.Name(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](101)
	assert.Panics(t, func() {
		obj.SetName(val)
	})
}
func TestProject_SetDescription(t *testing.T) {

	obj := NewProject()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](65535)
	obj.SetDescription(val)
	assert.Equal(t, val, obj.Description())
	assert.False(t, obj.DescriptionIsNull())

	// Test NULL
	obj.SetDescriptionToNull()
	assert.EqualValues(t, "", obj.Description())
	assert.True(t, obj.DescriptionIsNull())

	// test default
	obj.SetDescription("")
	assert.EqualValues(t, "", obj.Description(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](65536)
	assert.Panics(t, func() {
		obj.SetDescription(val)
	})
}
func TestProject_SetStartDate(t *testing.T) {

	obj := NewProject()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[time.Time](0)
	obj.SetStartDate(val)
	val = obj.StartDate()
	assert.Zero(t, val.Minute()) // make sure minute part is zero'd
	assert.Zero(t, val.Hour())   // make sure hour part is zero'd
	assert.Zero(t, val.Second()) // make sure second part is zero'd
	assert.False(t, obj.StartDateIsNull())

	// Test NULL
	obj.SetStartDateToNull()
	assert.EqualValues(t, time.Time{}, obj.StartDate())
	assert.True(t, obj.StartDateIsNull())

	// test default
	obj.SetStartDate(time.Time{})
	assert.EqualValues(t, time.Time{}, obj.StartDate(), "set default")

}
func TestProject_SetEndDate(t *testing.T) {

	obj := NewProject()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[time.Time](0)
	obj.SetEndDate(val)
	val = obj.EndDate()
	assert.Zero(t, val.Minute()) // make sure minute part is zero'd
	assert.Zero(t, val.Hour())   // make sure hour part is zero'd
	assert.Zero(t, val.Second()) // make sure second part is zero'd
	assert.False(t, obj.EndDateIsNull())

	// Test NULL
	obj.SetEndDateToNull()
	assert.EqualValues(t, time.Time{}, obj.EndDate())
	assert.True(t, obj.EndDateIsNull())

	// test default
	obj.SetEndDate(time.Time{})
	assert.EqualValues(t, time.Time{}, obj.EndDate(), "set default")

}
func TestProject_SetBudget(t *testing.T) {

	obj := NewProject()

	assert.True(t, obj.IsNew())
	val := test.RandomDecimal(12, 2)
	obj.SetBudget(val)
	assert.Equal(t, val, obj.Budget())
	assert.False(t, obj.BudgetIsNull())

	// Test NULL
	obj.SetBudgetToNull()
	assert.EqualValues(t, "", obj.Budget())
	assert.True(t, obj.BudgetIsNull())

	// test default
	obj.SetBudget("")
	assert.EqualValues(t, "", obj.Budget(), "set default")

}
func TestProject_SetSpent(t *testing.T) {

	obj := NewProject()

	assert.True(t, obj.IsNew())
	val := test.RandomDecimal(12, 2)
	obj.SetSpent(val)
	assert.Equal(t, val, obj.Spent())
	assert.False(t, obj.SpentIsNull())

	// Test NULL
	obj.SetSpentToNull()
	assert.EqualValues(t, "", obj.Spent())
	assert.True(t, obj.SpentIsNull())

	// test default
	obj.SetSpent("")
	assert.EqualValues(t, "", obj.Spent(), "set default")

}
func TestProject_SetParentProjectID(t *testing.T) {

	obj := NewProject()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetParentProjectID(val)
	assert.Equal(t, val, obj.ParentProjectID())
	assert.False(t, obj.ParentProjectIDIsNull())

	// Test NULL
	obj.SetParentProjectIDToNull()
	assert.EqualValues(t, "", obj.ParentProjectID())
	assert.True(t, obj.ParentProjectIDIsNull())

	// test default
	obj.SetParentProjectID("")
	assert.EqualValues(t, "", obj.ParentProjectID(), "set default")

}

func TestProject_Copy(t *testing.T) {
	obj := createMinimalSampleProject()

	obj2 := obj.Copy()

	assert.Equal(t, obj.Num(), obj2.Num())
	assert.Equal(t, obj.Status(), obj2.Status())
	assert.Equal(t, obj.ManagerID(), obj2.ManagerID())
	assert.Equal(t, obj.Name(), obj2.Name())
	assert.Equal(t, obj.Description(), obj2.Description())
	assert.Equal(t, obj.StartDate(), obj2.StartDate())
	assert.Equal(t, obj.EndDate(), obj2.EndDate())
	assert.Equal(t, obj.Budget(), obj2.Budget())
	assert.Equal(t, obj.Spent(), obj2.Spent())
	assert.Equal(t, obj.ParentProjectID(), obj2.ParentProjectID())

}

func TestProject_BasicInsert(t *testing.T) {
	obj := createMinimalSampleProject()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleProject(ctx, obj)

	// Test retrieval
	obj2 := LoadProject(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.NumIsValid())

	assert.EqualValues(t, obj.Num(), obj2.Num())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.numIsDirty)
	obj2.SetNum(obj2.Num())
	assert.False(t, obj2.numIsDirty)

	assert.True(t, obj2.StatusIsValid())

	assert.EqualValues(t, obj.Status(), obj2.Status())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.statusIsDirty)
	obj2.SetStatus(obj2.Status())
	assert.False(t, obj2.statusIsDirty)

	assert.True(t, obj2.NameIsValid())

	assert.EqualValues(t, obj.Name(), obj2.Name())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.nameIsDirty)
	obj2.SetName(obj2.Name())
	assert.False(t, obj2.nameIsDirty)

	assert.True(t, obj2.DescriptionIsValid())
	assert.False(t, obj2.DescriptionIsNull())

	assert.EqualValues(t, obj.Description(), obj2.Description())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.descriptionIsDirty)
	obj2.SetDescription(obj2.Description())
	assert.False(t, obj2.descriptionIsDirty)

	assert.True(t, obj2.StartDateIsValid())
	assert.False(t, obj2.StartDateIsNull())

	assert.EqualValues(t, obj.StartDate(), obj2.StartDate())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.startDateIsDirty)
	obj2.SetStartDate(obj2.StartDate())
	assert.False(t, obj2.startDateIsDirty)

	assert.True(t, obj2.EndDateIsValid())
	assert.False(t, obj2.EndDateIsNull())

	assert.EqualValues(t, obj.EndDate(), obj2.EndDate())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.endDateIsDirty)
	obj2.SetEndDate(obj2.EndDate())
	assert.False(t, obj2.endDateIsDirty)

	assert.True(t, obj2.BudgetIsValid())
	assert.False(t, obj2.BudgetIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.budgetIsDirty)
	obj2.SetBudget(obj2.Budget())
	assert.False(t, obj2.budgetIsDirty)

	assert.True(t, obj2.SpentIsValid())
	assert.False(t, obj2.SpentIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.spentIsDirty)
	obj2.SetSpent(obj2.Spent())
	assert.False(t, obj2.spentIsDirty)

}

func TestProject_InsertPanics(t *testing.T) {
	obj := createMinimalSampleProject()
	ctx := db.NewContext(nil)

	obj.numIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.numIsValid = true

	obj.statusIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.statusIsValid = true

	obj.nameIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsValid = true

}

func TestProject_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleProject()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleProject(ctx, obj)
	updateMinimalSampleProject(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadProject(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.Num(), obj.Num(), "Num did not update")
	assert.Equal(t, obj2.Status(), obj.Status(), "Status did not update")
	assert.Equal(t, obj2.ManagerID(), obj.ManagerID(), "ManagerID did not update")
	assert.Equal(t, obj2.Name(), obj.Name(), "Name did not update")
	assert.Equal(t, obj2.Description(), obj.Description(), "Description did not update")

	assert.WithinDuration(t, obj2.StartDate(), obj.StartDate(), time.Second, "StartDate not within one second")

	assert.WithinDuration(t, obj2.EndDate(), obj.EndDate(), time.Second, "EndDate not within one second")
	assert.Equal(t, obj2.ParentProjectID(), obj.ParentProjectID(), "ParentProjectID did not update")
}

func TestProject_References(t *testing.T) {
	obj := createMaximalSampleProject()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleProject(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.Manager())
	assert.NotEqual(t, '-', obj.Manager().PrimaryKey()[0])

	assert.NotNil(t, obj.ParentProject())
	assert.NotEqual(t, '-', obj.ParentProject().PrimaryKey()[0])

	obj2 := LoadProject(ctx, obj.PrimaryKey())
	objPkOnly := LoadProject(ctx, obj.PrimaryKey(), node.Project().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.Manager(), "Manager is not loaded initially")
	v_Manager := obj2.LoadManager(ctx)
	assert.NotNil(t, v_Manager)
	assert.Equal(t, v_Manager.PrimaryKey(), obj2.Manager().PrimaryKey())
	assert.Equal(t, obj.Manager().PrimaryKey(), obj2.Manager().PrimaryKey())
	assert.True(t, obj2.ManagerIDIsValid())

	assert.False(t, objPkOnly.ManagerIDIsValid())
	assert.Nil(t, objPkOnly.LoadManager(ctx))

	assert.Nil(t, obj2.ParentProject(), "ParentProject is not loaded initially")
	v_ParentProject := obj2.LoadParentProject(ctx)
	assert.NotNil(t, v_ParentProject)
	assert.Equal(t, v_ParentProject.PrimaryKey(), obj2.ParentProject().PrimaryKey())
	assert.Equal(t, obj.ParentProject().PrimaryKey(), obj2.ParentProject().PrimaryKey())
	assert.True(t, obj2.ParentProjectIDIsValid())

	assert.False(t, objPkOnly.ParentProjectIDIsValid())
	assert.Nil(t, objPkOnly.LoadParentProject(ctx))

}
func TestProject_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewProject()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestProject_Getters(t *testing.T) {
	obj := createMinimalSampleProject()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleProject(ctx, obj)

	assert.True(t, HasProject(ctx, obj.PrimaryKey()))

	obj2 := LoadProject(ctx, obj.PrimaryKey(), node.Project().PrimaryKey())

	assert.Panics(t, func() { obj2.Num() })
	assert.Panics(t, func() { obj2.Status() })
	assert.Panics(t, func() { obj2.ManagerID() })
	assert.Panics(t, func() { obj2.Name() })
	assert.Panics(t, func() { obj2.Description() })
	assert.Panics(t, func() { obj2.StartDate() })
	assert.Panics(t, func() { obj2.EndDate() })
	assert.Panics(t, func() { obj2.Budget() })
	assert.Panics(t, func() { obj2.Spent() })
	assert.Panics(t, func() { obj2.ParentProjectID() })
}
