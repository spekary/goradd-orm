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

func TestPersonWithLock_SetFirstName(t *testing.T) {

	obj := NewPersonWithLock()

	firstName := test.RandomValue[string](50)
	obj.SetFirstName(firstName)
	assert.Equal(t, firstName, obj.FirstName())

	// test zero
	obj.SetFirstName("")
	assert.Equal(t, "", obj.FirstName(), "set empty")

	// test panic on setting value larger than maximum size allowed
	firstName = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetFirstName(firstName)
	})
}
func TestPersonWithLock_SetLastName(t *testing.T) {

	obj := NewPersonWithLock()

	lastName := test.RandomValue[string](50)
	obj.SetLastName(lastName)
	assert.Equal(t, lastName, obj.LastName())

	// test zero
	obj.SetLastName("")
	assert.Equal(t, "", obj.LastName(), "set empty")

	// test panic on setting value larger than maximum size allowed
	lastName = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetLastName(lastName)
	})
}

// createMinimalSamplePersonWithLock creates and saves a minimal version of a PersonWithLock object
// for testing.
func createMinimalSamplePersonWithLock(ctx context.Context) *PersonWithLock {
	obj := NewPersonWithLock()

	firstName := test.RandomValue[string](50)
	obj.SetFirstName(firstName)

	lastName := test.RandomValue[string](50)
	obj.SetLastName(lastName)

	obj.Save(ctx)
	return obj
}
func TestPersonWithLock_CRUD(t *testing.T) {
	obj := NewPersonWithLock()
	ctx := db.NewContext(nil)

	firstName := test.RandomValue[string](50)
	obj.SetFirstName(firstName)

	lastName := test.RandomValue[string](50)
	obj.SetLastName(lastName)

	obj.Save(ctx)

	// Test retrieval
	obj = LoadPersonWithLock(ctx, obj.PrimaryKey())
	require.NotNil(t, obj)

	assert.True(t, obj.IDIsValid())
	assert.NotEmpty(t, obj.ID())

	assert.True(t, obj.FirstNameIsValid())
	assert.Equal(t, firstName, obj.FirstName())

	assert.True(t, obj.LastNameIsValid())
	assert.Equal(t, lastName, obj.LastName())

	assert.True(t, obj.SysTimestampIsValid())
	assert.False(t, obj.SysTimestampIsNull())
	assert.NotEmpty(t, obj.SysTimestamp())

}
