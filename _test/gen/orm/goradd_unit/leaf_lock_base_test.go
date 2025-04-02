// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd_unit/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createMinimalSampleLeafLock creates an unsaved minimal version of a LeafLock object
// for testing.
func createMinimalSampleLeafLock() *LeafLock {
	obj := NewLeafLock()
	updateMinimalSampleLeafLock(obj)

	return obj
}

// updateMinimalSampleLeafLock sets the values of a minimal sample to new, random values.
func updateMinimalSampleLeafLock(obj *LeafLock) {

	obj.SetName(test.RandomValue[string](100))

}

// createMaximalSampleLeafLock creates an unsaved version of a LeafLock object
// for testing that includes references to minimal objects.
func createMaximalSampleLeafLock(ctx context.Context) *LeafLock {
	obj := NewLeafLock()
	updateMaximalSampleLeafLock(ctx, obj)
	return obj
}

// updateMaximalSampleLeafLock sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleLeafLock(ctx context.Context, obj *LeafLock) {
	updateMinimalSampleLeafLock(obj)

}

// deleteSampleLeafLock deletes an object created and saved by one of the sample creator functions.
func deleteSampleLeafLock(ctx context.Context, obj *LeafLock) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

}

// assertEqualFieldsLeafLock compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsLeafLock(t *testing.T, obj1, obj2 *LeafLock) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.NameIsLoaded() && obj2.NameIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Name(), obj2.Name())
	}
	if obj1.GroLockIsLoaded() && obj2.GroLockIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.GroLock(), obj2.GroLock())
	}

}

func TestLeafLock_SetID(t *testing.T) {

	obj := NewLeafLock()

	assert.True(t, obj.IsNew())
	val := test.RandomNumberString()
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID("")
	assert.EqualValues(t, "", obj.ID(), "set default")

}
func TestLeafLock_SetName(t *testing.T) {

	obj := NewLeafLock()

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

func TestLeafLock_Copy(t *testing.T) {
	obj := createMinimalSampleLeafLock()

	obj2 := obj.Copy()

	assert.Equal(t, obj.Name(), obj2.Name())

}

func TestLeafLock_BasicInsert(t *testing.T) {
	obj := createMinimalSampleLeafLock()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeafLock(ctx, obj)

	// Test retrieval
	obj2 := LoadLeafLock(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsLoaded())
	assert.Panics(t, func() {
		obj2.SetID(obj2.ID())
	})

	assert.True(t, obj2.NameIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.nameIsDirty)
	obj2.SetName(obj2.Name())
	assert.False(t, obj2.nameIsDirty)

	assert.True(t, obj2.GroLockIsLoaded())

}

func TestLeafLock_InsertPanics(t *testing.T) {
	obj := createMinimalSampleLeafLock()
	ctx := db.NewContext(nil)

	obj.nameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsLoaded = true

}

func TestLeafLock_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleLeafLock()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafLock(ctx, obj)
	updateMinimalSampleLeafLock(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadLeafLock(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.Name(), obj.Name(), "Name did not update")
	assert.Equal(t, obj2.GroLock(), obj.GroLock(), "GroLock did not update")
}

func TestLeafLock_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafLock(ctx)
	obj.Save(ctx)
	defer deleteSampleLeafLock(ctx, obj)

	// Test that referenced objects were saved and assigned ids

	// Test lazy loading
	obj2 := LoadLeafLock(ctx, obj.PrimaryKey())
	objPkOnly := LoadLeafLock(ctx, obj.PrimaryKey(), node.LeafLock().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	// test eager loading
	obj3 := LoadLeafLock(ctx, obj.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}

func TestLeafLock_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafLock(ctx)
	obj.Save(ctx)
	defer deleteSampleLeafLock(ctx, obj)

	obj2 := LoadLeafLock(ctx, obj.PrimaryKey())
	updateMaximalSampleLeafLock(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleLeafLock(ctx, obj2)

	obj3 := LoadLeafLock(ctx, obj2.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}

func TestLeafLock_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafLock(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafLock(ctx, obj)

	assert.NoError(t, obj.Save(ctx))

	obj2 := LoadLeafLock(ctx, obj.PrimaryKey())
	_ = obj2 // avoid error if there are no references

}
func TestLeafLock_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewLeafLock()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestLeafLock_Getters(t *testing.T) {
	obj := createMinimalSampleLeafLock()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafLock(ctx, obj)

	assert.True(t, HasLeafLock(ctx, obj.PrimaryKey()))

	obj2 := LoadLeafLock(ctx, obj.PrimaryKey(), node.LeafLock().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.LeafLock().ID().Identifier))
	assert.Equal(t, obj.Name(), obj.Get(node.LeafLock().Name().Identifier))
	assert.Panics(t, func() { obj2.Name() })
	assert.Nil(t, obj2.Get(node.LeafLock().Name().Identifier))
	assert.Equal(t, obj.GroLock(), obj.Get(node.LeafLock().GroLock().Identifier))
	assert.Panics(t, func() { obj2.GroLock() })
	assert.Nil(t, obj2.Get(node.LeafLock().GroLock().Identifier))
}

func TestLeafLock_QueryLoad(t *testing.T) {
	obj := createMinimalSampleLeafLock()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeafLock(ctx, obj)

	objs := QueryLeafLocks(ctx).
		Where(op.Equal(node.LeafLock().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.LeafLock().PrimaryKey()). // exercise order by
		Limit(1, 0).                           // exercise limit
		Calculation(node.LeafLock(), "IsTrue", op.Equal(1, 1)).
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestLeafLock_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleLeafLock()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeafLock(ctx, obj)

	objs := QueryLeafLocks(ctx).
		Where(op.Equal(node.LeafLock().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestLeafLock_QueryCursor(t *testing.T) {
	obj := createMinimalSampleLeafLock()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeafLock(ctx, obj)

	cursor := QueryLeafLocks(ctx).
		Where(op.Equal(node.LeafLock().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryLeafLocks(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestLeafLock_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafLock(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	// reread in case there are data limitations imposed by the database
	obj2 := LoadLeafLock(ctx, obj.PrimaryKey())
	defer deleteSampleLeafLock(ctx, obj)

	assert.Less(t, 0, CountLeafLocks(ctx))

	assert.Less(t, 0, CountLeafLocksByID(ctx, obj2.ID()))
	assert.Less(t, 0, CountLeafLocksByName(ctx, obj2.Name()))
	assert.Less(t, 0, CountLeafLocksByGroLock(ctx, obj2.GroLock()))

}
func TestLeafLock_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleLeafLock()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewLeafLock()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsLeafLock(t, obj, obj2)
}

func TestLeafLock_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleLeafLock()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewLeafLock()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsLeafLock(t, obj, obj2)
}
