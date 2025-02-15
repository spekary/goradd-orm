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

func TestPerson_SetFirstName(t *testing.T) {

	obj := NewPerson()
	firstName := test.RandomValue[string](50)
	obj.SetFirstName(firstName)
	assert.Equal(t, firstName, obj.FirstName())

	// test default
	obj.SetFirstName("")
	assert.EqualValues(t, "", obj.FirstName(), "set default")

	// test panic on setting value larger than maximum size allowed
	firstName = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetFirstName(firstName)
	})
}
func TestPerson_SetLastName(t *testing.T) {

	obj := NewPerson()
	lastName := test.RandomValue[string](50)
	obj.SetLastName(lastName)
	assert.Equal(t, lastName, obj.LastName())

	// test default
	obj.SetLastName("")
	assert.EqualValues(t, "", obj.LastName(), "set default")

	// test panic on setting value larger than maximum size allowed
	lastName = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetLastName(lastName)
	})
}
func TestPerson_SetTypes(t *testing.T) {

	obj := NewPerson()

	types := test.RandomEnumArray(PersonTypes())
	obj.SetTypes(types)

	assert.Equal(t, types, obj.Types())
	assert.False(t, obj.TypesIsNull())

	// Test NULL
	obj.SetTypesToNull()
	assert.Nil(t, obj.Types())
	assert.True(t, obj.TypesIsNull())

	// test default
	obj.SetTypes(nil)
	assert.True(t, obj.Types().Equal(PersonTypeSet(nil)), "set default")

}

// createMinimalSamplePerson creates and saves a minimal version of a Person object
// for testing.
func createMinimalSamplePerson(ctx context.Context) *Person {
	obj := NewPerson()

	firstName := test.RandomValue[string](50)
	obj.SetFirstName(firstName)

	lastName := test.RandomValue[string](50)
	obj.SetLastName(lastName)

	types := test.RandomEnumArray(PersonTypes())
	obj.SetTypes(types)

	obj.Save(ctx)
	return obj
}
func TestPerson_CRUD(t *testing.T) {
	obj := NewPerson()
	ctx := db.NewContext(nil)

	firstName := test.RandomValue[string](50)
	obj.SetFirstName(firstName)

	lastName := test.RandomValue[string](50)
	obj.SetLastName(lastName)

	types := test.RandomEnumArray(PersonTypes())
	obj.SetTypes(types)

	obj.Save(ctx)

	// Test retrieval
	obj = LoadPerson(ctx, obj.PrimaryKey())
	require.NotNil(t, obj)

	assert.True(t, obj.IDIsValid())

	assert.True(t, obj.FirstNameIsValid())
	assert.Equal(t, firstName, obj.FirstName())

	assert.True(t, obj.LastNameIsValid())
	assert.Equal(t, lastName, obj.LastName())

	assert.True(t, obj.TypesIsValid())
	assert.False(t, obj.TypesIsNull())
	assert.True(t, types.Equal(obj.Types()))

}
