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

// createMinimalSamplePersonWithLock creates an unsaved minimal version of a PersonWithLock object
// for testing.
func createMinimalSamplePersonWithLock() *PersonWithLock {
	obj := NewPersonWithLock()
	updateMinimalSamplePersonWithLock(obj)
	return obj
}

// updateMinimalSamplePersonWithLock sets the values of a minimal sample to new, random values.
func updateMinimalSamplePersonWithLock(obj *PersonWithLock) {

	obj.SetFirstName(test.RandomValue[string](50))

	obj.SetLastName(test.RandomValue[string](50))

}

// createMaximalSamplePersonWithLock creates an unsaved version of a PersonWithLock object
// for testing that includes references to minimal objects.
func createMaximalSamplePersonWithLock() *PersonWithLock {
	obj := NewPersonWithLock()
	updateMaximalSamplePersonWithLock(obj)
	return obj
}

// updateMaximalSamplePersonWithLock sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSamplePersonWithLock(obj *PersonWithLock) {
	updateMinimalSamplePersonWithLock(obj)

}

// deleteSamplePersonWithLock deletes an object created and saved by one of the sample creator functions.
func deleteSamplePersonWithLock(ctx context.Context, obj *PersonWithLock) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

}

func TestPersonWithLock_SetFirstName(t *testing.T) {

	obj := NewPersonWithLock()

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
func TestPersonWithLock_SetLastName(t *testing.T) {

	obj := NewPersonWithLock()

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

func TestPersonWithLock_Copy(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()

	obj2 := obj.Copy()

	assert.Equal(t, obj.FirstName(), obj2.FirstName())
	assert.Equal(t, obj.LastName(), obj2.LastName())

}

func TestPersonWithLock_BasicInsert(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSamplePersonWithLock(ctx, obj)

	// Test retrieval
	obj2 := LoadPersonWithLock(ctx, obj.PrimaryKey())
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

	assert.True(t, obj2.GroLockIsValid())

	assert.True(t, obj2.GroTimestampIsValid())

}

func TestPersonWithLock_InsertPanics(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()
	ctx := db.NewContext(nil)

	obj.firstNameIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.firstNameIsValid = true

	obj.lastNameIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.lastNameIsValid = true

}

func TestPersonWithLock_BasicUpdate(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSamplePersonWithLock(ctx, obj)
	updateMinimalSamplePersonWithLock(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadPersonWithLock(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.FirstName(), obj.FirstName(), "FirstName did not update")
	assert.Equal(t, obj2.LastName(), obj.LastName(), "LastName did not update")
	assert.Equal(t, obj2.GroLock(), obj.GroLock(), "GroLock did not update")
	assert.Equal(t, obj2.GroTimestamp(), obj.GroTimestamp(), "GroTimestamp did not update")
}

func TestPersonWithLock_References(t *testing.T) {
	obj := createMaximalSamplePersonWithLock()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSamplePersonWithLock(ctx, obj)

	// Test that referenced objects were saved and assigned ids

	// Test lazy loading
	obj2 := LoadPersonWithLock(ctx, obj.PrimaryKey())
	objPkOnly := LoadPersonWithLock(ctx, obj.PrimaryKey(), node.PersonWithLock().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	// test eager loading
	obj3 := LoadPersonWithLock(ctx, obj.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}
func TestPersonWithLock_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewPersonWithLock()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestPersonWithLock_Getters(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSamplePersonWithLock(ctx, obj)

	assert.True(t, HasPersonWithLock(ctx, obj.PrimaryKey()))

	obj2 := LoadPersonWithLock(ctx, obj.PrimaryKey(), node.PersonWithLock().PrimaryKey())

	assert.Panics(t, func() { obj2.FirstName() })
	assert.Panics(t, func() { obj2.LastName() })
	assert.Panics(t, func() { obj2.GroLock() })
	assert.Panics(t, func() { obj2.GroTimestamp() })
}

func TestPersonWithLock_QueryLoad(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSamplePersonWithLock(ctx, obj)

	objs := QueryPersonWithLocks(ctx).
		Where(op.Equal(node.PersonWithLock().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.PersonWithLock().PrimaryKey()). // exercise order by
		Limit(1, 0).                                 // exercise limit
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
}
func TestPersonWithLock_QueryLoadI(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSamplePersonWithLock(ctx, obj)

	objs := QueryPersonWithLocks(ctx).
		Where(op.Equal(node.PersonWithLock().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestPersonWithLock_QueryCursor(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSamplePersonWithLock(ctx, obj)

	cursor := QueryPersonWithLocks(ctx).
		Where(op.Equal(node.PersonWithLock().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryPersonWithLocks(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestPersonWithLock_Count(t *testing.T) {
	obj := createMaximalSamplePersonWithLock()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSamplePersonWithLock(ctx, obj)

	assert.Less(t, 0, CountPersonWithLocksByID(ctx, obj.ID()))
	assert.Less(t, 0, CountPersonWithLocksByFirstName(ctx, obj.FirstName()))
	assert.Less(t, 0, CountPersonWithLocksByLastName(ctx, obj.LastName()))
	assert.Less(t, 0, CountPersonWithLocksByGroLock(ctx, obj.GroLock()))
	assert.Less(t, 0, CountPersonWithLocksByGroTimestamp(ctx, obj.GroTimestamp()))
}
