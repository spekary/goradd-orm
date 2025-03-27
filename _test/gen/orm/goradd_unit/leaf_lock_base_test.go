// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"context"
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
func createMaximalSampleLeafLock() *LeafLock {
	obj := NewLeafLock()
	updateMaximalSampleLeafLock(obj)
	return obj
}

// updateMaximalSampleLeafLock sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleLeafLock(obj *LeafLock) {
	updateMinimalSampleLeafLock(obj)

}

// deleteSampleLeafLock deletes an object created and saved by one of the sample creator functions.
func deleteSampleLeafLock(ctx context.Context, obj *LeafLock) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

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

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.NameIsValid())

	assert.EqualValues(t, obj.Name(), obj2.Name())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.nameIsDirty)
	obj2.SetName(obj2.Name())
	assert.False(t, obj2.nameIsDirty)

	assert.True(t, obj2.GroLockIsValid())

}

func TestLeafLock_InsertPanics(t *testing.T) {
	obj := createMinimalSampleLeafLock()
	ctx := db.NewContext(nil)

	obj.nameIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsValid = true

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

func TestLeafLock_References(t *testing.T) {
	obj := createMaximalSampleLeafLock()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleLeafLock(ctx, obj)

	// Test that referenced objects were saved and assigned ids

	obj2 := LoadLeafLock(ctx, obj.PrimaryKey())
	objPkOnly := LoadLeafLock(ctx, obj.PrimaryKey(), node.LeafLock().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

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

	assert.Panics(t, func() { obj2.Name() })
	assert.Panics(t, func() { obj2.GroLock() })
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
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
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
	obj := createMaximalSampleLeafLock()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeafLock(ctx, obj)

	assert.Less(t, 0, CountLeafLocksByID(ctx, obj.ID()))
	assert.Less(t, 0, CountLeafLocksByName(ctx, obj.Name()))
	assert.Less(t, 0, CountLeafLocksByGroLock(ctx, obj.GroLock()))
}
