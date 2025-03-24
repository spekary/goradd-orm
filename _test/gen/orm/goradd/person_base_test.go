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

func TestPerson_SetFirstName(t *testing.T) {

	obj := NewPerson()
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

func TestPerson_BasicInsert(t *testing.T) {
	obj := NewPerson()
	ctx := db.NewContext(nil)

	v_firstName := test.RandomValue[string](50)
	obj.SetFirstName(v_firstName)

	v_lastName := test.RandomValue[string](50)
	obj.SetLastName(v_lastName)

	v_types := test.RandomEnumArray(PersonTypes())
	obj.SetTypes(v_types)

	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer obj.Delete(ctx)

	// Test retrieval
	obj2 := LoadPerson(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.FirstNameIsValid())
	assert.EqualValues(t, v_firstName, obj2.FirstName())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.firstNameIsDirty)
	obj2.SetFirstName(obj2.FirstName())
	assert.False(t, obj2.firstNameIsDirty)

	assert.True(t, obj2.LastNameIsValid())
	assert.EqualValues(t, v_lastName, obj2.LastName())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.lastNameIsDirty)
	obj2.SetLastName(obj2.LastName())
	assert.False(t, obj2.lastNameIsDirty)

	assert.True(t, obj2.TypesIsValid())
	assert.False(t, obj2.TypesIsNull())
	assert.True(t, v_types.Equal(obj2.Types()))
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

	assert.Equal(t, obj2.ID(), obj.ID())
	assert.Equal(t, obj2.FirstName(), obj.FirstName())
	assert.Equal(t, obj2.LastName(), obj.LastName())
	assert.Equal(t, obj2.Types(), obj.Types())
}

func TestPerson_References(t *testing.T) {
	obj := createMaximalSamplePerson()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSamplePerson(ctx, obj)

	// Test that referenced objects were saved and assigned ids

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

	obj2 := LoadPerson(ctx, obj.PrimaryKey(), node.Person().PrimaryKey())

	assert.Panics(t, func() { obj2.FirstName() })
	assert.Panics(t, func() { obj2.LastName() })
	assert.Panics(t, func() { obj2.Types() })
}
