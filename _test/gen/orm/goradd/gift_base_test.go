// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
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

	obj.SetNumber(test.RandomValue[int](32))

	return obj
}

// updateMinimalSampleGift sets the values of a minimal sample to new, random values.
func updateMinimalSampleGift(obj *Gift) {

	obj.SetName(test.RandomValue[string](50))

}

// createMaximalSampleGift creates an unsaved version of a Gift object
// for testing that includes references to minimal objects.
func createMaximalSampleGift(ctx context.Context) *Gift {
	obj := NewGift()
	obj.SetNumber(test.RandomValue[int](32))
	updateMaximalSampleGift(ctx, obj)
	return obj
}

// updateMaximalSampleGift sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleGift(ctx context.Context, obj *Gift) {
	updateMinimalSampleGift(obj)

}

// deleteSampleGift deletes an object created and saved by one of the sample creator functions.
func deleteSampleGift(ctx context.Context, obj *Gift) {
	if obj == nil {
		return
	}

	_ = obj.Delete(ctx)

}

// assertEqualFieldsGift compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsGift(t *testing.T, obj1, obj2 *Gift) {
	if obj1.NumberIsLoaded() && obj2.NumberIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Number(), obj2.Number())
	}
	if obj1.NameIsLoaded() && obj2.NameIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Name(), obj2.Name())
	}

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
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleGift(ctx, obj)

	// Test retrieval
	obj2, err := LoadGift(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)
	assert.NoError(t, err)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.NumberIsLoaded())
	assert.Panics(t, func() {
		obj2.SetNumber(obj2.Number())
	})

	assert.True(t, obj2.NameIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.nameIsDirty)
	obj2.SetName(obj2.Name())
	assert.False(t, obj2.nameIsDirty)

}

func TestGift_InsertPanics(t *testing.T) {
	obj := createMinimalSampleGift()
	_ = obj
	ctx := db.NewContext(nil)
	_ = ctx

	obj.numberIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.numberIsLoaded = true

	obj.nameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsLoaded = true

}

func TestGift_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleGift()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleGift(ctx, obj)
	updateMinimalSampleGift(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2, err := LoadGift(ctx, obj.PrimaryKey())
	assert.NoError(t, err)

	assert.Equal(t, obj2.Number(), obj.Number(), "Number did not update")
	assert.Equal(t, obj2.Name(), obj.Name(), "Name did not update")
}

func TestGift_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleGift(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleGift(ctx, obj)

	// Test that referenced objects were saved and assigned ids

	// Test lazy loading
	obj2, err := LoadGift(ctx, obj.PrimaryKey())
	assert.NoError(t, err)
	objPkOnly, err2 := LoadGift(ctx, obj.PrimaryKey(), node.Gift().PrimaryKey())
	assert.NoError(t, err2)
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	// test eager loading
	obj3, _ := LoadGift(ctx, obj.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}

func TestGift_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleGift(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleGift(ctx, obj)

	obj2, _ := LoadGift(ctx, obj.PrimaryKey())
	updateMaximalSampleGift(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleGift(ctx, obj2)

	obj3, _ := LoadGift(ctx, obj2.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}

func TestGift_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleGift(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleGift(ctx, obj)

	assert.NoError(t, obj.Save(ctx))

	obj2, _ := LoadGift(ctx, obj.PrimaryKey())
	_ = obj2 // avoid error if there are no references

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

	has, _ := HasGift(ctx, obj.PrimaryKey())
	assert.True(t, has)

	obj2, _ := LoadGift(ctx, obj.PrimaryKey(), node.Gift().PrimaryKey())
	assert.Equal(t, obj.Number(), obj2.Number())

	assert.Equal(t, obj.Number(), obj.Get(node.Gift().Number().Identifier))
	assert.Equal(t, obj.Name(), obj.Get(node.Gift().Name().Identifier))
	assert.Panics(t, func() { obj2.Name() })
	assert.Nil(t, obj2.Get(node.Gift().Name().Identifier))
}

func TestGift_QueryLoad(t *testing.T) {
	obj := createMinimalSampleGift()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleGift(ctx, obj)

	objs, _ := QueryGifts(ctx).
		Where(op.Equal(node.Gift().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.Gift().PrimaryKey()). // exercise order by
		Limit(1, 0).                       // exercise limit
		Calculation(node.Gift(), "IsTrue", op.Equal(1, 1)).
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestGift_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleGift()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleGift(ctx, obj)

	objs, _ := QueryGifts(ctx).
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

	cursor, _ := QueryGifts(ctx).
		Where(op.Equal(node.Gift().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2, err2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.NoError(t, err2)
	obj2, err2 = cursor.Next()
	assert.Nil(t, obj2)
	assert.NoError(t, err2)
	assert.NoError(t, cursor.Close())

	// test empty cursor result
	cursor, err = QueryGifts(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	obj2, err = cursor.Next()
	assert.Nil(t, obj2)
	assert.NoError(t, err)
	assert.NoError(t, cursor.Close())
}
func TestGift_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleGift(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleGift(ctx, obj)
	// reread in case there are data limitations imposed by the database
	obj2, _ := LoadGift(ctx, obj.PrimaryKey())

	assert.Positive(t, func() int { i, _ := CountGifts(ctx); return i }())
	assert.Positive(t, func() int { i, _ := CountGiftsByNumber(ctx, obj2.Number()); return i }())
	assert.Positive(t, func() int { i, _ := CountGiftsByName(ctx, obj2.Name()); return i }())

}

func TestGift_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleGift()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewGift()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsGift(t, obj, obj2)
}

func TestGift_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleGift()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewGift()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsGift(t, obj, obj2)
}

func TestGift_FailingMarshalBinary(t *testing.T) {
	obj := createMinimalSampleGift()
	var err error

	for i := 0; i < 9; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
	// do it again with aliases
	obj._aliases = make(map[string]any)
	for i := 0; i < 10; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
}

func TestGift_FailingUnmarshalBinary(t *testing.T) {
	obj := createMinimalSampleGift()
	b, err := obj.MarshalBinary()
	assert.NoError(t, err)
	obj2 := NewGift()
	for i := 0; i < 9; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}

	// do it again with aliases
	obj = createMinimalSampleGift()
	obj._aliases = map[string]any{"a": 1}
	b, err = obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 = NewGift()
	for i := 0; i < 10; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}
}

func TestGift_Indexes(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleGift(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleGift(ctx, obj)

	var obj2 *Gift
	obj2, _ = LoadGiftByNumber(ctx, obj.Number())
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.True(t, func() bool { h, _ := HasGiftByNumber(ctx, obj.Number()); return h }())

}
