// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
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
// This will set new values for references, so save the old values and delete them.
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

// assertEqualFieldsProject compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsProject(t *testing.T, obj1, obj2 *Project) {
	assert.EqualValues(t, obj1.ID(), obj2.ID())

	assert.EqualValues(t, obj1.Num(), obj2.Num())

	assert.EqualValues(t, obj1.Status(), obj2.Status())

	assert.EqualValues(t, obj1.ManagerID(), obj2.ManagerID())

	assert.EqualValues(t, obj1.Name(), obj2.Name())

	assert.EqualValues(t, obj1.Description(), obj2.Description())

	assert.EqualValues(t, obj1.StartDate(), obj2.StartDate())

	assert.EqualValues(t, obj1.EndDate(), obj2.EndDate())

	assert.True(t, test.EqualDecimals(obj1.Budget(), obj2.Budget()))

	assert.True(t, test.EqualDecimals(obj1.Spent(), obj2.Spent()))

	assert.EqualValues(t, obj1.ParentProjectID(), obj2.ParentProjectID())

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

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.numIsDirty)
	obj2.SetNum(obj2.Num())
	assert.False(t, obj2.numIsDirty)

	assert.True(t, obj2.StatusIsValid())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.statusIsDirty)
	obj2.SetStatus(obj2.Status())
	assert.False(t, obj2.statusIsDirty)

	assert.True(t, obj2.NameIsValid())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.nameIsDirty)
	obj2.SetName(obj2.Name())
	assert.False(t, obj2.nameIsDirty)

	assert.True(t, obj2.DescriptionIsValid())
	assert.False(t, obj2.DescriptionIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.descriptionIsDirty)
	obj2.SetDescription(obj2.Description())
	assert.False(t, obj2.descriptionIsDirty)

	assert.True(t, obj2.StartDateIsValid())
	assert.False(t, obj2.StartDateIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.startDateIsDirty)
	obj2.SetStartDate(obj2.StartDate())
	assert.False(t, obj2.startDateIsDirty)

	assert.True(t, obj2.EndDateIsValid())
	assert.False(t, obj2.EndDateIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.endDateIsDirty)
	obj2.SetEndDate(obj2.EndDate())
	assert.False(t, obj2.endDateIsDirty)

	assert.True(t, obj2.BudgetIsValid())
	assert.False(t, obj2.BudgetIsNull())

	assert.True(t, test.EqualDecimals(obj.Budget(), obj2.Budget()))

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.budgetIsDirty)
	obj2.SetBudget(obj2.Budget())
	assert.False(t, obj2.budgetIsDirty)

	assert.True(t, obj2.SpentIsValid())
	assert.False(t, obj2.SpentIsNull())

	assert.True(t, test.EqualDecimals(obj.Spent(), obj2.Spent()))

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

func TestProject_ReferenceLoad(t *testing.T) {
	obj := createMaximalSampleProject()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleProject(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.Manager())
	assert.NotEqual(t, '-', obj.Manager().PrimaryKey()[0])

	assert.NotNil(t, obj.ParentProject())
	assert.NotEqual(t, '-', obj.ParentProject().PrimaryKey()[0])

	// Test lazy loading
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

	assert.Nil(t, obj2.Milestones(), "Milestones is not loaded initially")
	v_Milestones := obj2.LoadMilestones(ctx)
	assert.NotNil(t, v_Milestones)
	assert.Len(t, v_Milestones, 1)
	assert.Nil(t, obj2.ParentProjectProjects(), "ParentProjectProjects is not loaded initially")
	v_ParentProjectProjects := obj2.LoadParentProjectProjects(ctx)
	assert.NotNil(t, v_ParentProjectProjects)
	assert.Len(t, v_ParentProjectProjects, 1)

	assert.Nil(t, obj2.Children(), "Children is not loaded initially")
	v_Children := obj2.LoadChildren(ctx)
	assert.NotNil(t, v_Children)
	assert.Len(t, v_Children, 1)
	assert.Nil(t, obj2.Parents(), "Parents is not loaded initially")
	v_Parents := obj2.LoadParents(ctx)
	assert.NotNil(t, v_Parents)
	assert.Len(t, v_Parents, 1)
	assert.Nil(t, obj2.TeamMembers(), "TeamMembers is not loaded initially")
	v_TeamMembers := obj2.LoadTeamMembers(ctx)
	assert.NotNil(t, v_TeamMembers)
	assert.Len(t, v_TeamMembers, 1)

	// test eager loading
	obj3 := LoadProject(ctx, obj.PrimaryKey(), node.Project().Manager(),
		node.Project().ParentProject(),
		node.Project().Milestones(),
		node.Project().ParentProjectProjects(),
		node.Project().Children(),
		node.Project().Parents(),
		node.Project().TeamMembers(),
	)
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Manager().PrimaryKey(), obj3.Manager().PrimaryKey())
	assert.Equal(t, obj2.ParentProject().PrimaryKey(), obj3.ParentProject().PrimaryKey())
	assert.Equal(t, len(obj2.Milestones()), len(obj3.Milestones()))
	assert.Equal(t, len(obj2.ParentProjectProjects()), len(obj3.ParentProjectProjects()))
	assert.Equal(t, len(obj2.Children()), len(obj3.Children()))
	assert.Equal(t, len(obj2.Parents()), len(obj3.Parents()))
	assert.Equal(t, len(obj2.TeamMembers()), len(obj3.TeamMembers()))

}

func TestProject_ReferenceUpdateNewObjects(t *testing.T) {
	obj := createMaximalSampleProject()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleProject(ctx, obj)

	obj2 := LoadProject(ctx, obj.PrimaryKey())
	updateMaximalSampleProject(obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleProject(ctx, obj2)

	obj3 := LoadProject(ctx, obj2.PrimaryKey(), node.Project().Manager(),
		node.Project().ParentProject(),
		node.Project().Milestones(),
		node.Project().ParentProjectProjects(),
		node.Project().Children(),
		node.Project().Parents(),
		node.Project().TeamMembers(),
	)
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Manager().PrimaryKey(), obj3.Manager().PrimaryKey())
	assert.Equal(t, obj2.ParentProject().PrimaryKey(), obj3.ParentProject().PrimaryKey())

	assert.Equal(t, len(obj2.Milestones()), len(obj3.Milestones()))
	assert.Equal(t, len(obj2.ParentProjectProjects()), len(obj3.ParentProjectProjects()))

	assert.Equal(t, len(obj2.Children()), len(obj3.Children()))
	assert.Equal(t, len(obj2.Parents()), len(obj3.Parents()))
	assert.Equal(t, len(obj2.TeamMembers()), len(obj3.TeamMembers()))

}

func TestProject_ReferenceUpdateOldObjects(t *testing.T) {
	obj := createMaximalSampleProject()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleProject(ctx, obj)

	updateMinimalSamplePerson(obj.Manager())
	updateMinimalSampleProject(obj.ParentProject())
	updateMinimalSampleMilestone(obj.Milestones()[0])
	updateMinimalSampleProject(obj.ParentProjectProjects()[0])
	updateMinimalSampleProject(obj.Children()[0])
	updateMinimalSampleProject(obj.Parents()[0])
	updateMinimalSamplePerson(obj.TeamMembers()[0])

	assert.NoError(t, obj.Save(ctx))

	obj2 := LoadProject(ctx, obj.PrimaryKey(),

		node.Project().Manager(),

		node.Project().ParentProject(),

		node.Project().Milestones(),
		node.Project().ParentProjectProjects(),
		node.Project().Children(),
		node.Project().Parents(),
		node.Project().TeamMembers(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsPerson(t, obj2.Manager(), obj.Manager())
	assertEqualFieldsProject(t, obj2.ParentProject(), obj.ParentProject())

	assertEqualFieldsMilestone(t, obj2.Milestones()[0], obj.Milestones()[0])
	assertEqualFieldsProject(t, obj2.ParentProjectProjects()[0], obj.ParentProjectProjects()[0])

	assertEqualFieldsProject(t, obj2.Children()[0], obj.Children()[0])
	assertEqualFieldsProject(t, obj2.Parents()[0], obj.Parents()[0])
	assertEqualFieldsPerson(t, obj2.TeamMembers()[0], obj.TeamMembers()[0])
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

	assert.Equal(t, obj.ID(), obj.Get(node.Project().ID().Identifier))
	assert.Equal(t, obj.Num(), obj.Get(node.Project().Num().Identifier))
	assert.Panics(t, func() { obj2.Num() })
	assert.Nil(t, obj2.Get(node.Project().Num().Identifier))
	assert.Equal(t, obj.Status(), obj.Get(node.Project().Status().Identifier))
	assert.Panics(t, func() { obj2.Status() })
	assert.Nil(t, obj2.Get(node.Project().Status().Identifier))
	assert.Equal(t, obj.ManagerID(), obj.Get(node.Project().ManagerID().Identifier))
	assert.Panics(t, func() { obj2.ManagerID() })
	assert.Nil(t, obj2.Get(node.Project().ManagerID().Identifier))
	assert.Equal(t, obj.Name(), obj.Get(node.Project().Name().Identifier))
	assert.Panics(t, func() { obj2.Name() })
	assert.Nil(t, obj2.Get(node.Project().Name().Identifier))
	assert.Equal(t, obj.Description(), obj.Get(node.Project().Description().Identifier))
	assert.Panics(t, func() { obj2.Description() })
	assert.Nil(t, obj2.Get(node.Project().Description().Identifier))
	assert.Equal(t, obj.StartDate(), obj.Get(node.Project().StartDate().Identifier))
	assert.Panics(t, func() { obj2.StartDate() })
	assert.Nil(t, obj2.Get(node.Project().StartDate().Identifier))
	assert.Equal(t, obj.EndDate(), obj.Get(node.Project().EndDate().Identifier))
	assert.Panics(t, func() { obj2.EndDate() })
	assert.Nil(t, obj2.Get(node.Project().EndDate().Identifier))
	assert.Equal(t, obj.Budget(), obj.Get(node.Project().Budget().Identifier))
	assert.Panics(t, func() { obj2.Budget() })
	assert.Nil(t, obj2.Get(node.Project().Budget().Identifier))
	assert.Equal(t, obj.Spent(), obj.Get(node.Project().Spent().Identifier))
	assert.Panics(t, func() { obj2.Spent() })
	assert.Nil(t, obj2.Get(node.Project().Spent().Identifier))
	assert.Equal(t, obj.ParentProjectID(), obj.Get(node.Project().ParentProjectID().Identifier))
	assert.Panics(t, func() { obj2.ParentProjectID() })
	assert.Nil(t, obj2.Get(node.Project().ParentProjectID().Identifier))
}

func TestProject_QueryLoad(t *testing.T) {
	obj := createMinimalSampleProject()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleProject(ctx, obj)

	objs := QueryProjects(ctx).
		Where(op.Equal(node.Project().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.Project().PrimaryKey()). // exercise order by
		Limit(1, 0).                          // exercise limit
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
}
func TestProject_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleProject()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleProject(ctx, obj)

	objs := QueryProjects(ctx).
		Where(op.Equal(node.Project().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestProject_QueryCursor(t *testing.T) {
	obj := createMinimalSampleProject()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleProject(ctx, obj)

	cursor := QueryProjects(ctx).
		Where(op.Equal(node.Project().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryProjects(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestProject_Count(t *testing.T) {
	obj := createMaximalSampleProject()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleProject(ctx, obj)

	assert.Less(t, 0, CountProjects(ctx))

	assert.Less(t, 0, CountProjectsByID(ctx, obj.ID()))
	assert.Less(t, 0, CountProjectsByNum(ctx, obj.Num()))
	assert.Less(t, 0, CountProjectsByManagerID(ctx, obj.ManagerID()))
	assert.Less(t, 0, CountProjectsByName(ctx, obj.Name()))
	assert.Less(t, 0, CountProjectsByDescription(ctx, obj.Description()))
	assert.Less(t, 0, CountProjectsByStartDate(ctx, obj.StartDate()))
	assert.Less(t, 0, CountProjectsByEndDate(ctx, obj.EndDate()))
	assert.Less(t, 0, CountProjectsByBudget(ctx, obj.Budget()))
	assert.Less(t, 0, CountProjectsBySpent(ctx, obj.Spent()))
	assert.Less(t, 0, CountProjectsByParentProjectID(ctx, obj.ParentProjectID()))

}
