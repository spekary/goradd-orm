// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"context"
	"strconv"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd_unit/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLeafLock_SetName(t *testing.T) {

	obj := NewLeafLock()
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

// createMinimalSampleLeafLock creates an unsaved minimal version of a LeafLock object
// for testing.
func createMinimalSampleLeafLock() *LeafLock {
	obj := NewLeafLock()

	obj.SetName(test.RandomValue[string](100))

	return obj
}

// createMaximalSampleLeafLock creates an unsaved version of a LeafLock object
// for testing that includes references to minimal objects.
func createMaximalSampleLeafLock() *LeafLock {
	obj := createMinimalSampleLeafLock()

	return obj
}

// deleteSampleLeafLock deletes an object created and saved by one of the sample creator functions.
func deleteSampleLeafLock(ctx context.Context, obj *LeafLock) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

}

func TestLeafLock_CRUD(t *testing.T) {
	obj := NewLeafLock()
	ctx := db.NewContext(nil)

	v_name := test.RandomValue[string](100)
	obj.SetName(v_name)

	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer obj.Delete(ctx)

	// Test retrieval
	obj2 := LoadLeafLock(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.NameIsValid())
	assert.EqualValues(t, v_name, obj2.Name())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.nameIsDirty)
	obj2.SetName(obj2.Name())
	assert.False(t, obj2.nameIsDirty)

	assert.True(t, obj2.GroLockIsValid())

}

func TestLeafLock_References(t *testing.T) {
	obj := createMaximalSampleLeafLock()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleLeafLock(ctx, obj)

	// Test that referenced objects were saved and assigned ids

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

	obj2 := LoadLeafLock(ctx, obj.PrimaryKey(), node.LeafLock().PrimaryKey())

	assert.Panics(t, func() { obj2.Name() })
	assert.Panics(t, func() { obj2.GroLock() })
}
