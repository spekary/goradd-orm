// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createMinimalSampleGift creates an unsaved minimal version of a Gift object
// for testing.
func createMinimalSampleGift() *Gift {
	obj := NewGift()
	updateMinimalSampleGift(obj)
	return obj
}

// updateMinimalSampleGift sets the values of a minimal sample to new, random values.
func updateMinimalSampleGift(obj *Gift) {

	obj.SetNumber(test.RandomValue[int](32))

	obj.SetName(test.RandomValue[string](50))

}

// createMaximalSampleGift creates an unsaved version of a Gift object
// for testing that includes references to minimal objects.
func createMaximalSampleGift() *Gift {
	obj := NewGift()
	updateMaximalSampleGift(obj)
	return obj
}

// updateMaximalSampleGift sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleGift(obj *Gift) {
	updateMinimalSampleGift(obj)

}

// deleteSampleGift deletes an object created and saved by one of the sample creator functions.
func deleteSampleGift(ctx context.Context, obj *Gift) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

}

func TestGift_SetNumber(t *testing.T) {

	obj := NewGift()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[int](32)
	obj.SetNumber(val)
	assert.Equal(t, val, obj.Number())

	// test default
	obj.SetNumber(0)
	assert.EqualValues(t, 0, obj.Number(), "set default")

}
func TestGift_SetName(t *testing.T) {

	obj := NewGift()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](50)
	obj.SetName(val)
	assert.Equal(t, val, obj.Name())

	// test default
	obj.SetName("")
	assert.EqualValues(t, "", obj.Name(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetName(val)
	})
}

func TestGift_Copy(t *testing.T) {
	obj := createMinimalSampleGift()

	obj2 := obj.Copy()

	assert.Equal(t, obj.Number(), obj2.Number())
	assert.Equal(t, obj.Name(), obj2.Name())

}

func TestGift_BasicInsert(t *testing.T) {
	obj := createMinimalSampleGift()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleGift(ctx, obj)

	// Test retrieval
	obj2 := LoadGift(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.NumberIsValid())

	assert.EqualValues(t, obj.Number(), obj2.Number())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.numberIsDirty)
	obj2.SetNumber(obj2.Number())
	assert.False(t, obj2.numberIsDirty)

	assert.True(t, obj2.NameIsValid())

	assert.EqualValues(t, obj.Name(), obj2.Name())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.nameIsDirty)
	obj2.SetName(obj2.Name())
	assert.False(t, obj2.nameIsDirty)

}

func TestGift_InsertPanics(t *testing.T) {
	obj := createMinimalSampleGift()
	ctx := db.NewContext(nil)

	obj.numberIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.numberIsValid = true

	obj.nameIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsValid = true

}

func TestGift_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleGift()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleGift(ctx, obj)
	updateMinimalSampleGift(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadGift(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.Number(), obj.Number(), "Number did not update")
	assert.Equal(t, obj2.Name(), obj.Name(), "Name did not update")
}

func TestGift_ReferenceLoad(t *testing.T) {
	obj := createMaximalSampleGift()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleGift(ctx, obj)

	// Test that referenced objects were saved and assigned ids

	// Test lazy loading
	obj2 := LoadGift(ctx, obj.PrimaryKey())
	objPkOnly := LoadGift(ctx, obj.PrimaryKey(), node.Gift().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	// test eager loading
	obj3 := LoadGift(ctx, obj.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}
func TestGift_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewGift()

	assert.Zero(t, obj.Number())
}

func TestGift_Getters(t *testing.T) {
	obj := createMinimalSampleGift()

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleGift(ctx, obj)

	assert.True(t, HasGift(ctx, obj.PrimaryKey()))

	obj2 := LoadGift(ctx, obj.PrimaryKey(), node.Gift().PrimaryKey())
	assert.Equal(t, obj.Number(), obj2.Number())

	assert.Panics(t, func() { obj2.Name() })
}

func TestGift_QueryLoad(t *testing.T) {
	obj := createMinimalSampleGift()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleGift(ctx, obj)

	objs := QueryGifts(ctx).
		Where(op.Equal(node.Gift().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.Gift().PrimaryKey()). // exercise order by
		Limit(1, 0).                       // exercise limit
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
}
func TestGift_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleGift()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleGift(ctx, obj)

	objs := QueryGifts(ctx).
		Where(op.Equal(node.Gift().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("Number"))
}
func TestGift_QueryCursor(t *testing.T) {
	obj := createMinimalSampleGift()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleGift(ctx, obj)

	cursor := QueryGifts(ctx).
		Where(op.Equal(node.Gift().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryGifts(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestGift_Count(t *testing.T) {
	obj := createMaximalSampleGift()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleGift(ctx, obj)

	assert.Less(t, 0, CountGiftsByNumber(ctx, obj.Number()))
	assert.Less(t, 0, CountGiftsByName(ctx, obj.Name()))
}
