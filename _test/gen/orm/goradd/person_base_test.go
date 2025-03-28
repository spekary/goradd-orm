// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
	"strconv"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createMinimalSamplePerson creates an unsaved minimal version of a Person object
// for testing.
func createMinimalSamplePerson() *Person {
	obj := NewPerson()
	updateMinimalSamplePerson(obj)

	return obj
}

// updateMinimalSamplePerson sets the values of a minimal sample to new, random values.
func updateMinimalSamplePerson(obj *Person) {

	obj.SetFirstName(test.RandomValue[string](50))

	obj.SetLastName(test.RandomValue[string](50))

	obj.SetTypes(test.RandomEnumArray(PersonTypes()))

}

// createMaximalSamplePerson creates an unsaved version of a Person object
// for testing that includes references to minimal objects.
func createMaximalSamplePerson() *Person {
	obj := NewPerson()
	updateMaximalSamplePerson(obj)
	return obj
}

// updateMaximalSamplePerson sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSamplePerson(obj *Person) {
	updateMinimalSamplePerson(obj)

	obj.SetAddresses(createMinimalSampleAddress())
	obj.SetEmployeeInfo(createMinimalSampleEmployeeInfo())
	obj.SetLogin(createMinimalSampleLogin())
	obj.SetManagerProjects(createMinimalSampleProject())
	obj.SetProjects(createMinimalSampleProject())
}

// deleteSamplePerson deletes an object created and saved by one of the sample creator functions.
func deleteSamplePerson(ctx context.Context, obj *Person) {
	if obj == nil {
		return
	}

	for _, item := range obj.Addresses() {
		deleteSampleAddress(ctx, item)
	}
	deleteSampleEmployeeInfo(ctx, obj.EmployeeInfo())
	deleteSampleLogin(ctx, obj.Login())
	for _, item := range obj.ManagerProjects() {
		deleteSampleProject(ctx, item)
	}

	for _, item := range obj.Projects() {
		deleteSampleProject(ctx, item)
	}

	obj.Delete(ctx)

}

// assertEqualFieldsPerson compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsPerson(t *testing.T, obj1, obj2 *Person) {
	assert.EqualValues(t, obj1.ID(), obj2.ID())

	assert.EqualValues(t, obj1.FirstName(), obj2.FirstName())

	assert.EqualValues(t, obj1.LastName(), obj2.LastName())

	assert.True(t, obj1.Types().Equal(obj2.Types()))

}

func TestPerson_SetFirstName(t *testing.T) {

	obj := NewPerson()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](50)
	obj.SetFirstName(val)
	assert.Equal(t, val, obj.FirstName())

	// test default
	obj.SetFirstName("")
	assert.EqualValues(t, "", obj.FirstName(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetFirstName(val)
	})
}
func TestPerson_SetLastName(t *testing.T) {

	obj := NewPerson()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](50)
	obj.SetLastName(val)
	assert.Equal(t, val, obj.LastName())

	// test default
	obj.SetLastName("")
	assert.EqualValues(t, "", obj.LastName(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetLastName(val)
	})
}
func TestPerson_SetTypes(t *testing.T) {

	obj := NewPerson()

	assert.True(t, obj.IsNew())

	val := test.RandomEnumArray(PersonTypes())
	obj.SetTypes(val)

	assert.Equal(t, val, obj.Types())
	assert.False(t, obj.TypesIsNull())

	// Test NULL
	obj.SetTypesToNull()
	assert.Nil(t, obj.Types())
	assert.True(t, obj.TypesIsNull())

	// test default
	obj.SetTypes(nil)
	assert.True(t, obj.Types().Equal(PersonTypeSet(nil)), "set default")

}

func TestPerson_Copy(t *testing.T) {
	obj := createMinimalSamplePerson()

	obj2 := obj.Copy()

	assert.Equal(t, obj.FirstName(), obj2.FirstName())
	assert.Equal(t, obj.LastName(), obj2.LastName())
	assert.Equal(t, obj.Types(), obj2.Types())

}

func TestPerson_BasicInsert(t *testing.T) {
	obj := createMinimalSamplePerson()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSamplePerson(ctx, obj)

	// Test retrieval
	obj2 := LoadPerson(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.FirstNameIsValid())

	assert.EqualValues(t, obj.FirstName(), obj2.FirstName())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.firstNameIsDirty)
	obj2.SetFirstName(obj2.FirstName())
	assert.False(t, obj2.firstNameIsDirty)

	assert.True(t, obj2.LastNameIsValid())

	assert.EqualValues(t, obj.LastName(), obj2.LastName())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.lastNameIsDirty)
	obj2.SetLastName(obj2.LastName())
	assert.False(t, obj2.lastNameIsDirty)

	assert.True(t, obj2.TypesIsValid())
	assert.False(t, obj2.TypesIsNull())
	assert.True(t, obj.Types().Equal(obj2.Types()))
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.typesIsDirty)
	obj2.SetTypes(obj2.Types())
	assert.False(t, obj2.typesIsDirty)

}

func TestPerson_InsertPanics(t *testing.T) {
	obj := createMinimalSamplePerson()
	ctx := db.NewContext(nil)

	obj.firstNameIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.firstNameIsValid = true

	obj.lastNameIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.lastNameIsValid = true

}

func TestPerson_BasicUpdate(t *testing.T) {
	obj := createMinimalSamplePerson()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSamplePerson(ctx, obj)
	updateMinimalSamplePerson(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadPerson(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.FirstName(), obj.FirstName(), "FirstName did not update")
	assert.Equal(t, obj2.LastName(), obj.LastName(), "LastName did not update")
	assert.Equal(t, obj2.Types(), obj.Types(), "Types did not update")
}

func TestPerson_ReferenceLoad(t *testing.T) {
	obj := createMaximalSamplePerson()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSamplePerson(ctx, obj)

	// Test that referenced objects were saved and assigned ids

	// Test lazy loading
	obj2 := LoadPerson(ctx, obj.PrimaryKey())
	objPkOnly := LoadPerson(ctx, obj.PrimaryKey(), node.Person().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.Addresses(), "Addresses is not loaded initially")
	v_Addresses := obj2.LoadAddresses(ctx)
	assert.NotNil(t, v_Addresses)
	assert.Len(t, v_Addresses, 1)
	assert.Nil(t, obj2.EmployeeInfo(), "EmployeeInfo is not loaded initially")
	v_EmployeeInfo := obj2.LoadEmployeeInfo(ctx)
	assert.NotNil(t, v_EmployeeInfo)
	assert.Equal(t, v_EmployeeInfo.PrimaryKey(), obj2.EmployeeInfo().PrimaryKey())
	assert.Equal(t, obj.EmployeeInfo().PrimaryKey(), obj2.EmployeeInfo().PrimaryKey())
	assert.Nil(t, obj2.Login(), "Login is not loaded initially")
	v_Login := obj2.LoadLogin(ctx)
	assert.NotNil(t, v_Login)
	assert.Equal(t, v_Login.PrimaryKey(), obj2.Login().PrimaryKey())
	assert.Equal(t, obj.Login().PrimaryKey(), obj2.Login().PrimaryKey())
	assert.Nil(t, obj2.ManagerProjects(), "ManagerProjects is not loaded initially")
	v_ManagerProjects := obj2.LoadManagerProjects(ctx)
	assert.NotNil(t, v_ManagerProjects)
	assert.Len(t, v_ManagerProjects, 1)

	assert.Nil(t, obj2.Projects(), "Projects is not loaded initially")
	v_Projects := obj2.LoadProjects(ctx)
	assert.NotNil(t, v_Projects)
	assert.Len(t, v_Projects, 1)

	// test eager loading
	obj3 := LoadPerson(ctx, obj.PrimaryKey(), node.Person().Addresses(),
		node.Person().EmployeeInfo(),
		node.Person().Login(),
		node.Person().ManagerProjects(),
		node.Person().Projects(),
	)
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, len(obj2.Addresses()), len(obj3.Addresses()))
	assert.Equal(t, obj2.EmployeeInfo().PrimaryKey(), obj3.EmployeeInfo().PrimaryKey())
	assert.Equal(t, obj2.Login().PrimaryKey(), obj3.Login().PrimaryKey())
	assert.Equal(t, len(obj2.ManagerProjects()), len(obj3.ManagerProjects()))
	assert.Equal(t, len(obj2.Projects()), len(obj3.Projects()))

}

func TestPerson_ReferenceUpdateNewObjects(t *testing.T) {
	obj := createMaximalSamplePerson()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSamplePerson(ctx, obj)

	obj2 := LoadPerson(ctx, obj.PrimaryKey())
	updateMaximalSamplePerson(obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSamplePerson(ctx, obj2)

	obj3 := LoadPerson(ctx, obj2.PrimaryKey(), node.Person().Addresses(),
		node.Person().EmployeeInfo(),
		node.Person().Login(),
		node.Person().ManagerProjects(),
		node.Person().Projects(),
	)
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, len(obj2.Addresses()), len(obj3.Addresses()))
	assert.Equal(t, obj2.EmployeeInfo().PrimaryKey(), obj3.EmployeeInfo().PrimaryKey())
	assert.Equal(t, obj2.Login().PrimaryKey(), obj3.Login().PrimaryKey())
	assert.Equal(t, len(obj2.ManagerProjects()), len(obj3.ManagerProjects()))

	assert.Equal(t, len(obj2.Projects()), len(obj3.Projects()))

}

func TestPerson_ReferenceUpdateOldObjects(t *testing.T) {
	obj := createMaximalSamplePerson()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSamplePerson(ctx, obj)

	updateMinimalSampleAddress(obj.Addresses()[0])
	updateMinimalSampleEmployeeInfo(obj.EmployeeInfo())
	updateMinimalSampleLogin(obj.Login())
	updateMinimalSampleProject(obj.ManagerProjects()[0])
	updateMinimalSampleProject(obj.Projects()[0])

	assert.NoError(t, obj.Save(ctx))

	obj2 := LoadPerson(ctx, obj.PrimaryKey(),

		node.Person().Addresses(),
		node.Person().EmployeeInfo(),
		node.Person().Login(),
		node.Person().ManagerProjects(),
		node.Person().Projects(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsAddress(t, obj2.Addresses()[0], obj.Addresses()[0])
	assertEqualFieldsEmployeeInfo(t, obj2.EmployeeInfo(), obj.EmployeeInfo())
	assertEqualFieldsLogin(t, obj2.Login(), obj.Login())
	assertEqualFieldsProject(t, obj2.ManagerProjects()[0], obj.ManagerProjects()[0])

	assertEqualFieldsProject(t, obj2.Projects()[0], obj.Projects()[0])
}
func TestPerson_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewPerson()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestPerson_Getters(t *testing.T) {
	obj := createMinimalSamplePerson()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSamplePerson(ctx, obj)

	assert.True(t, HasPerson(ctx, obj.PrimaryKey()))

	obj2 := LoadPerson(ctx, obj.PrimaryKey(), node.Person().PrimaryKey())

	assert.Panics(t, func() { obj2.FirstName() })
	assert.Panics(t, func() { obj2.LastName() })
	assert.Panics(t, func() { obj2.Types() })
}

func TestPerson_QueryLoad(t *testing.T) {
	obj := createMinimalSamplePerson()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSamplePerson(ctx, obj)

	objs := QueryPeople(ctx).
		Where(op.Equal(node.Person().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.Person().PrimaryKey()). // exercise order by
		Limit(1, 0).                         // exercise limit
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
}
func TestPerson_QueryLoadI(t *testing.T) {
	obj := createMinimalSamplePerson()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSamplePerson(ctx, obj)

	objs := QueryPeople(ctx).
		Where(op.Equal(node.Person().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestPerson_QueryCursor(t *testing.T) {
	obj := createMinimalSamplePerson()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSamplePerson(ctx, obj)

	cursor := QueryPeople(ctx).
		Where(op.Equal(node.Person().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryPeople(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestPerson_Count(t *testing.T) {
	obj := createMaximalSamplePerson()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSamplePerson(ctx, obj)

	assert.Less(t, 0, CountPeopleByID(ctx, obj.ID()))
	assert.Less(t, 0, CountPeopleByFirstName(ctx, obj.FirstName()))
	assert.Less(t, 0, CountPeopleByLastName(ctx, obj.LastName()))
}
