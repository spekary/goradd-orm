// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
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
func createMaximalSamplePersonWithLock(ctx context.Context) *PersonWithLock {
	obj := NewPersonWithLock()
	updateMaximalSamplePersonWithLock(ctx, obj)
	return obj
}

// updateMaximalSamplePersonWithLock sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSamplePersonWithLock(ctx context.Context, obj *PersonWithLock) {
	updateMinimalSamplePersonWithLock(obj)

}

// deleteSamplePersonWithLock deletes an object created and saved by one of the sample creator functions.
func deleteSamplePersonWithLock(ctx context.Context, obj *PersonWithLock) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

}

// assertEqualFieldsPersonWithLock compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsPersonWithLock(t *testing.T, obj1, obj2 *PersonWithLock) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.FirstNameIsLoaded() && obj2.FirstNameIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.FirstName(), obj2.FirstName())
	}
	if obj1.LastNameIsLoaded() && obj2.LastNameIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.LastName(), obj2.LastName())
	}
	if obj1.GroLockIsLoaded() && obj2.GroLockIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.GroLock(), obj2.GroLock())
	}
	if obj1.GroTimestampIsLoaded() && obj2.GroTimestampIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.GroTimestamp(), obj2.GroTimestamp())
	}

}

func TestPersonWithLock_SetID(t *testing.T) {

	obj := NewPersonWithLock()

	assert.True(t, obj.IsNew())
	val := test.RandomNumberString()
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID("")
	assert.EqualValues(t, "", obj.ID(), "set default")

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

	assert.True(t, obj2.IDIsLoaded())
	assert.Panics(t, func() {
		obj2.SetID(obj2.ID())
	})

	assert.True(t, obj2.FirstNameIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.firstNameIsDirty)
	obj2.SetFirstName(obj2.FirstName())
	assert.False(t, obj2.firstNameIsDirty)

	assert.True(t, obj2.LastNameIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.lastNameIsDirty)
	obj2.SetLastName(obj2.LastName())
	assert.False(t, obj2.lastNameIsDirty)

	assert.True(t, obj2.GroLockIsLoaded())

	assert.True(t, obj2.GroTimestampIsLoaded())

}

func TestPersonWithLock_InsertPanics(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()
	ctx := db.NewContext(nil)

	obj.firstNameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.firstNameIsLoaded = true

	obj.lastNameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.lastNameIsLoaded = true

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

func TestPersonWithLock_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSamplePersonWithLock(ctx)
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

func TestPersonWithLock_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSamplePersonWithLock(ctx)
	obj.Save(ctx)
	defer deleteSamplePersonWithLock(ctx, obj)

	obj2 := LoadPersonWithLock(ctx, obj.PrimaryKey())
	updateMaximalSamplePersonWithLock(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSamplePersonWithLock(ctx, obj2)

	obj3 := LoadPersonWithLock(ctx, obj2.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}

func TestPersonWithLock_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSamplePersonWithLock(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSamplePersonWithLock(ctx, obj)

	assert.NoError(t, obj.Save(ctx))

	obj2 := LoadPersonWithLock(ctx, obj.PrimaryKey())
	_ = obj2 // avoid error if there are no references

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

	assert.Equal(t, obj.ID(), obj.Get(node.PersonWithLock().ID().Identifier))
	assert.Equal(t, obj.FirstName(), obj.Get(node.PersonWithLock().FirstName().Identifier))
	assert.Panics(t, func() { obj2.FirstName() })
	assert.Nil(t, obj2.Get(node.PersonWithLock().FirstName().Identifier))
	assert.Equal(t, obj.LastName(), obj.Get(node.PersonWithLock().LastName().Identifier))
	assert.Panics(t, func() { obj2.LastName() })
	assert.Nil(t, obj2.Get(node.PersonWithLock().LastName().Identifier))
	assert.Equal(t, obj.GroLock(), obj.Get(node.PersonWithLock().GroLock().Identifier))
	assert.Panics(t, func() { obj2.GroLock() })
	assert.Nil(t, obj2.Get(node.PersonWithLock().GroLock().Identifier))
	assert.Equal(t, obj.GroTimestamp(), obj.Get(node.PersonWithLock().GroTimestamp().Identifier))
	assert.Panics(t, func() { obj2.GroTimestamp() })
	assert.Nil(t, obj2.Get(node.PersonWithLock().GroTimestamp().Identifier))
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
		Calculation(node.PersonWithLock(), "IsTrue", op.Equal(1, 1)).
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
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
	ctx := db.NewContext(nil)
	obj := createMaximalSamplePersonWithLock(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	// reread in case there are data limitations imposed by the database
	obj2 := LoadPersonWithLock(ctx, obj.PrimaryKey())
	defer deleteSamplePersonWithLock(ctx, obj)

	assert.Less(t, 0, CountPersonWithLocks(ctx))

	assert.Less(t, 0, CountPersonWithLocksByID(ctx, obj2.ID()))
	assert.Less(t, 0, CountPersonWithLocksByFirstName(ctx, obj2.FirstName()))
	assert.Less(t, 0, CountPersonWithLocksByLastName(ctx, obj2.LastName()))
	assert.Less(t, 0, CountPersonWithLocksByGroLock(ctx, obj2.GroLock()))
	assert.Less(t, 0, CountPersonWithLocksByGroTimestamp(ctx, obj2.GroTimestamp()))

}
func TestPersonWithLock_MarshalJSON(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewPersonWithLock()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsPersonWithLock(t, obj, obj2)
}

func TestPersonWithLock_MarshalBinary(t *testing.T) {
	obj := createMinimalSamplePersonWithLock()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewPersonWithLock()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsPersonWithLock(t, obj, obj2)
}
