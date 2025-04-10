// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd_unit/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createMinimalSampleDoubleIndex creates an unsaved minimal version of a DoubleIndex object
// for testing.
func createMinimalSampleDoubleIndex() *DoubleIndex {
	obj := NewDoubleIndex()
	updateMinimalSampleDoubleIndex(obj)

	obj.SetID(test.RandomValue[int](32))

	return obj
}

// updateMinimalSampleDoubleIndex sets the values of a minimal sample to new, random values.
func updateMinimalSampleDoubleIndex(obj *DoubleIndex) {

	obj.SetFieldInt(test.RandomValue[int](32))

	obj.SetFieldString(test.RandomValue[string](50))

	obj.SetField2Int(test.RandomValue[int](32))

	obj.SetField2String(test.RandomValue[string](100))

}

// createMaximalSampleDoubleIndex creates an unsaved version of a DoubleIndex object
// for testing that includes references to minimal objects.
func createMaximalSampleDoubleIndex(ctx context.Context) *DoubleIndex {
	obj := NewDoubleIndex()
	obj.SetID(test.RandomValue[int](32))
	updateMaximalSampleDoubleIndex(ctx, obj)
	return obj
}

// updateMaximalSampleDoubleIndex sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleDoubleIndex(ctx context.Context, obj *DoubleIndex) {
	updateMinimalSampleDoubleIndex(obj)

}

// deleteSampleDoubleIndex deletes an object created and saved by one of the sample creator functions.
func deleteSampleDoubleIndex(ctx context.Context, obj *DoubleIndex) {
	if obj == nil {
		return
	}

	_ = obj.Delete(ctx)

}

// assertEqualFieldsDoubleIndex compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsDoubleIndex(t *testing.T, obj1, obj2 *DoubleIndex) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.FieldIntIsLoaded() && obj2.FieldIntIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.FieldInt(), obj2.FieldInt())
	}
	if obj1.FieldStringIsLoaded() && obj2.FieldStringIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.FieldString(), obj2.FieldString())
	}
	if obj1.Field2IntIsLoaded() && obj2.Field2IntIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Field2Int(), obj2.Field2Int())
	}
	if obj1.Field2StringIsLoaded() && obj2.Field2StringIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Field2String(), obj2.Field2String())
	}

}

func TestDoubleIndex_SetID(t *testing.T) {

	obj := NewDoubleIndex()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[int](32)
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID(0)
	assert.EqualValues(t, 0, obj.ID(), "set default")

}
func TestDoubleIndex_SetFieldInt(t *testing.T) {

	obj := NewDoubleIndex()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[int](32)
	obj.SetFieldInt(val)
	assert.Equal(t, val, obj.FieldInt())

	// test default
	obj.SetFieldInt(0)
	assert.EqualValues(t, 0, obj.FieldInt(), "set default")

}
func TestDoubleIndex_SetFieldString(t *testing.T) {

	obj := NewDoubleIndex()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](50)
	obj.SetFieldString(val)
	assert.Equal(t, val, obj.FieldString())

	// test default
	obj.SetFieldString("")
	assert.EqualValues(t, "", obj.FieldString(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetFieldString(val)
	})
}
func TestDoubleIndex_SetField2Int(t *testing.T) {

	obj := NewDoubleIndex()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[int](32)
	obj.SetField2Int(val)
	assert.Equal(t, val, obj.Field2Int())
	assert.False(t, obj.Field2IntIsNull())

	// Test NULL
	obj.SetField2IntToNull()
	assert.EqualValues(t, 0, obj.Field2Int())
	assert.True(t, obj.Field2IntIsNull())

	// test default
	obj.SetField2Int(0)
	assert.EqualValues(t, 0, obj.Field2Int(), "set default")

}
func TestDoubleIndex_SetField2String(t *testing.T) {

	obj := NewDoubleIndex()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](100)
	obj.SetField2String(val)
	assert.Equal(t, val, obj.Field2String())
	assert.False(t, obj.Field2StringIsNull())

	// Test NULL
	obj.SetField2StringToNull()
	assert.EqualValues(t, "", obj.Field2String())
	assert.True(t, obj.Field2StringIsNull())

	// test default
	obj.SetField2String("")
	assert.EqualValues(t, "", obj.Field2String(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](101)
	assert.Panics(t, func() {
		obj.SetField2String(val)
	})
}

func TestDoubleIndex_Copy(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()

	obj2 := obj.Copy()

	assert.Equal(t, obj.ID(), obj2.ID())
	assert.Equal(t, obj.FieldInt(), obj2.FieldInt())
	assert.Equal(t, obj.FieldString(), obj2.FieldString())
	assert.Equal(t, obj.Field2Int(), obj2.Field2Int())
	assert.Equal(t, obj.Field2String(), obj2.Field2String())

}

func TestDoubleIndex_BasicInsert(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleDoubleIndex(ctx, obj)

	// Test retrieval
	obj2, err := LoadDoubleIndex(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)
	assert.NoError(t, err)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsLoaded())
	assert.Panics(t, func() {
		obj2.SetID(obj2.ID())
	})

	assert.True(t, obj2.FieldIntIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.fieldIntIsDirty)
	obj2.SetFieldInt(obj2.FieldInt())
	assert.False(t, obj2.fieldIntIsDirty)

	assert.True(t, obj2.FieldStringIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.fieldStringIsDirty)
	obj2.SetFieldString(obj2.FieldString())
	assert.False(t, obj2.fieldStringIsDirty)

	assert.True(t, obj2.Field2IntIsLoaded())
	assert.False(t, obj2.Field2IntIsNull())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.field2IntIsDirty)
	obj2.SetField2Int(obj2.Field2Int())
	assert.False(t, obj2.field2IntIsDirty)

	assert.True(t, obj2.Field2StringIsLoaded())
	assert.False(t, obj2.Field2StringIsNull())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.field2StringIsDirty)
	obj2.SetField2String(obj2.Field2String())
	assert.False(t, obj2.field2StringIsDirty)

}

func TestDoubleIndex_InsertPanics(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()
	_ = obj
	ctx := db.NewContext(nil)
	_ = ctx

	obj.idIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.idIsLoaded = true

	obj.fieldIntIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.fieldIntIsLoaded = true

	obj.fieldStringIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.fieldStringIsLoaded = true

}

func TestDoubleIndex_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleDoubleIndex(ctx, obj)
	updateMinimalSampleDoubleIndex(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2, err := LoadDoubleIndex(ctx, obj.PrimaryKey())
	assert.NoError(t, err)

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.FieldInt(), obj.FieldInt(), "FieldInt did not update")
	assert.Equal(t, obj2.FieldString(), obj.FieldString(), "FieldString did not update")
	assert.Equal(t, obj2.Field2Int(), obj.Field2Int(), "Field2Int did not update")
	assert.Equal(t, obj2.Field2String(), obj.Field2String(), "Field2String did not update")
}

func TestDoubleIndex_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleDoubleIndex(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleDoubleIndex(ctx, obj)

	// Test that referenced objects were saved and assigned ids

	// Test lazy loading
	obj2, err := LoadDoubleIndex(ctx, obj.PrimaryKey())
	assert.NoError(t, err)
	objPkOnly, err2 := LoadDoubleIndex(ctx, obj.PrimaryKey(), node.DoubleIndex().PrimaryKey())
	assert.NoError(t, err2)
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	// test eager loading
	obj3, _ := LoadDoubleIndex(ctx, obj.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}

func TestDoubleIndex_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleDoubleIndex(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleDoubleIndex(ctx, obj)

	obj2, _ := LoadDoubleIndex(ctx, obj.PrimaryKey())
	updateMaximalSampleDoubleIndex(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleDoubleIndex(ctx, obj2)

	obj3, _ := LoadDoubleIndex(ctx, obj2.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}

func TestDoubleIndex_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleDoubleIndex(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleDoubleIndex(ctx, obj)

	assert.NoError(t, obj.Save(ctx))

	obj2, _ := LoadDoubleIndex(ctx, obj.PrimaryKey())
	_ = obj2 // avoid error if there are no references

}
func TestDoubleIndex_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewDoubleIndex()

	assert.Zero(t, obj.ID())
}

func TestDoubleIndex_Getters(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleDoubleIndex(ctx, obj)

	has, _ := HasDoubleIndex(ctx, obj.PrimaryKey())
	assert.True(t, has)

	obj2, _ := LoadDoubleIndex(ctx, obj.PrimaryKey(), node.DoubleIndex().PrimaryKey())
	assert.Equal(t, obj.ID(), obj2.ID())

	assert.Equal(t, obj.ID(), obj.Get(node.DoubleIndex().ID().Identifier))
	assert.Equal(t, obj.FieldInt(), obj.Get(node.DoubleIndex().FieldInt().Identifier))
	assert.Panics(t, func() { obj2.FieldInt() })
	assert.Nil(t, obj2.Get(node.DoubleIndex().FieldInt().Identifier))
	assert.Equal(t, obj.FieldString(), obj.Get(node.DoubleIndex().FieldString().Identifier))
	assert.Panics(t, func() { obj2.FieldString() })
	assert.Nil(t, obj2.Get(node.DoubleIndex().FieldString().Identifier))
	assert.Equal(t, obj.Field2Int(), obj.Get(node.DoubleIndex().Field2Int().Identifier))
	assert.Panics(t, func() { obj2.Field2Int() })
	assert.Nil(t, obj2.Get(node.DoubleIndex().Field2Int().Identifier))
	assert.Equal(t, obj.Field2String(), obj.Get(node.DoubleIndex().Field2String().Identifier))
	assert.Panics(t, func() { obj2.Field2String() })
	assert.Nil(t, obj2.Get(node.DoubleIndex().Field2String().Identifier))
}

func TestDoubleIndex_QueryLoad(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleDoubleIndex(ctx, obj)

	objs, _ := QueryDoubleIndices(ctx).
		Where(op.Equal(node.DoubleIndex().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.DoubleIndex().PrimaryKey()). // exercise order by
		Limit(1, 0).                              // exercise limit
		Calculation(node.DoubleIndex(), "IsTrue", op.Equal(1, 1)).
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestDoubleIndex_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleDoubleIndex(ctx, obj)

	objs, _ := QueryDoubleIndices(ctx).
		Where(op.Equal(node.DoubleIndex().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestDoubleIndex_QueryCursor(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleDoubleIndex(ctx, obj)

	cursor, _ := QueryDoubleIndices(ctx).
		Where(op.Equal(node.DoubleIndex().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2, err2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.NoError(t, err2)
	obj2, err2 = cursor.Next()
	assert.Nil(t, obj2)
	assert.NoError(t, err2)
	assert.NoError(t, cursor.Close())

	// test empty cursor result
	cursor, err = QueryDoubleIndices(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	obj2, err = cursor.Next()
	assert.Nil(t, obj2)
	assert.NoError(t, err)
	assert.NoError(t, cursor.Close())
}
func TestDoubleIndex_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleDoubleIndex(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleDoubleIndex(ctx, obj)
	// reread in case there are data limitations imposed by the database
	obj2, _ := LoadDoubleIndex(ctx, obj.PrimaryKey())

	assert.Positive(t, func() int { i, _ := CountDoubleIndices(ctx); return i }())
	assert.Positive(t, func() int { i, _ := CountDoubleIndicesByID(ctx, obj2.ID()); return i }())
	assert.Positive(t, func() int { i, _ := CountDoubleIndicesByFieldInt(ctx, obj2.FieldInt()); return i }())
	assert.Positive(t, func() int { i, _ := CountDoubleIndicesByFieldString(ctx, obj2.FieldString()); return i }())
	assert.Positive(t, func() int { i, _ := CountDoubleIndicesByField2Int(ctx, obj2.Field2Int()); return i }())
	assert.Positive(t, func() int { i, _ := CountDoubleIndicesByField2String(ctx, obj2.Field2String()); return i }())

}

func TestDoubleIndex_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewDoubleIndex()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsDoubleIndex(t, obj, obj2)
}

func TestDoubleIndex_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewDoubleIndex()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsDoubleIndex(t, obj, obj2)
}

func TestDoubleIndex_FailingMarshalBinary(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()
	var err error

	for i := 0; i < 20; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
	// do it again with aliases
	obj._aliases = make(map[string]any)
	for i := 0; i < 21; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
}

func TestDoubleIndex_FailingUnmarshalBinary(t *testing.T) {
	obj := createMinimalSampleDoubleIndex()
	b, err := obj.MarshalBinary()
	assert.NoError(t, err)
	obj2 := NewDoubleIndex()
	for i := 0; i < 20; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}

	// do it again with aliases
	obj = createMinimalSampleDoubleIndex()
	obj._aliases = map[string]any{"a": 1}
	b, err = obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 = NewDoubleIndex()
	for i := 0; i < 21; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}
}

func TestDoubleIndex_Indexes(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleDoubleIndex(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleDoubleIndex(ctx, obj)

	var obj2 *DoubleIndex
	obj2, _ = LoadDoubleIndexByID(ctx, obj.ID())
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.True(t, func() bool { h, _ := HasDoubleIndexByID(ctx, obj.ID()); return h }())

	obj2, _ = LoadDoubleIndexByField2IntField2String(ctx, obj.Field2Int(), obj.Field2String())
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.True(t, func() bool {
		h, _ := HasDoubleIndexByField2IntField2String(ctx, obj.Field2Int(), obj.Field2String())
		return h
	}())

	obj2, _ = LoadDoubleIndexByFieldIntFieldString(ctx, obj.FieldInt(), obj.FieldString())
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.True(t, func() bool {
		h, _ := HasDoubleIndexByFieldIntFieldString(ctx, obj.FieldInt(), obj.FieldString())
		return h
	}())

}
