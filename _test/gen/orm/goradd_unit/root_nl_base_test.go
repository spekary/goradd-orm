// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"bytes"
	"context"
	"encoding/gob"
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

// createMinimalSampleRootNl creates an unsaved minimal version of a RootNl object
// for testing.
func createMinimalSampleRootNl() *RootNl {
	obj := NewRootNl()
	updateMinimalSampleRootNl(obj)

	return obj
}

// updateMinimalSampleRootNl sets the values of a minimal sample to new, random values.
func updateMinimalSampleRootNl(obj *RootNl) {

	obj.SetName(test.RandomValue[string](100))

}

// createMaximalSampleRootNl creates an unsaved version of a RootNl object
// for testing that includes references to minimal objects.
func createMaximalSampleRootNl(ctx context.Context) *RootNl {
	obj := NewRootNl()
	updateMaximalSampleRootNl(ctx, obj)
	return obj
}

// updateMaximalSampleRootNl sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleRootNl(ctx context.Context, obj *RootNl) {
	updateMinimalSampleRootNl(obj)

	obj.SetLeafNls(createMinimalSampleLeafNl())
}

// deleteSampleRootNl deletes an object created and saved by one of the sample creator functions.
func deleteSampleRootNl(ctx context.Context, obj *RootNl) {
	if obj == nil {
		return
	}

	for _, item := range obj.LeafNls() {
		deleteSampleLeafNl(ctx, item)
	}

	_ = obj.Delete(ctx)

}

// assertEqualFieldsRootNl compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsRootNl(t *testing.T, obj1, obj2 *RootNl) {
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

func TestRootNl_SetID(t *testing.T) {

	obj := NewRootNl()

	assert.True(t, obj.IsNew())
	val := test.RandomNumberString()
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID("")
	assert.EqualValues(t, "", obj.ID(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](33)
	assert.Panics(t, func() {
		obj.SetID(val)
	})
}
func TestRootNl_SetName(t *testing.T) {

	obj := NewRootNl()

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

func TestRootNl_Copy(t *testing.T) {
	obj := createMinimalSampleRootNl()

	obj2 := obj.Copy()

	assert.Equal(t, obj.Name(), obj2.Name())

}

func TestRootNl_BasicInsert(t *testing.T) {
	obj := createMinimalSampleRootNl()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleRootNl(ctx, obj)

	// Test retrieval
	obj2, err := LoadRootNl(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)
	assert.NoError(t, err)

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

func TestRootNl_InsertPanics(t *testing.T) {
	obj := createMinimalSampleRootNl()
	_ = obj
	ctx := db.NewContext(nil)
	_ = ctx

	obj.nameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsLoaded = true

}

func TestRootNl_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleRootNl()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleRootNl(ctx, obj)
	updateMinimalSampleRootNl(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2, err := LoadRootNl(ctx, obj.PrimaryKey())
	assert.NoError(t, err)

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.Name(), obj.Name(), "Name did not update")
	assert.Equal(t, obj2.GroLock(), obj.GroLock(), "GroLock did not update")
}

func TestRootNl_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleRootNl(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleRootNl(ctx, obj)

	// Test that referenced objects were saved and assigned ids

	// Test lazy loading
	obj2, err := LoadRootNl(ctx, obj.PrimaryKey())
	assert.NoError(t, err)
	objPkOnly, err2 := LoadRootNl(ctx, obj.PrimaryKey(), node.RootNl().PrimaryKey())
	assert.NoError(t, err2)
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.LeafNls(), "LeafNls is not loaded initially")
	v_LeafNls, _ := obj2.LoadLeafNls(ctx)
	assert.NotNil(t, v_LeafNls)
	assert.Len(t, v_LeafNls, 1)

	// test eager loading
	obj3, _ := LoadRootNl(ctx, obj.PrimaryKey(), node.RootNl().LeafNls())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, len(obj2.LeafNls()), len(obj3.LeafNls()))

}

func TestRootNl_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleRootNl(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleRootNl(ctx, obj)

	obj2, _ := LoadRootNl(ctx, obj.PrimaryKey())
	updateMaximalSampleRootNl(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleRootNl(ctx, obj2)

	obj3, _ := LoadRootNl(ctx, obj2.PrimaryKey(), node.RootNl().LeafNls())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, len(obj2.LeafNls()), len(obj3.LeafNls()))

}

func TestRootNl_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleRootNl(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleRootNl(ctx, obj)

	updateMinimalSampleLeafNl(obj.LeafNls()[0])

	assert.NoError(t, obj.Save(ctx))

	obj2, _ := LoadRootNl(ctx, obj.PrimaryKey(),

		node.RootNl().LeafNls(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsLeafNl(t, obj2.LeafNls()[0], obj.LeafNls()[0])

}
func TestRootNl_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewRootNl()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestRootNl_Getters(t *testing.T) {
	obj := createMinimalSampleRootNl()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleRootNl(ctx, obj)

	has, _ := HasRootNl(ctx, obj.PrimaryKey())
	assert.True(t, has)

	obj2, _ := LoadRootNl(ctx, obj.PrimaryKey(), node.RootNl().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.RootNl().ID().Identifier))
	assert.Equal(t, obj.Name(), obj.Get(node.RootNl().Name().Identifier))
	assert.Panics(t, func() { obj2.Name() })
	assert.Nil(t, obj2.Get(node.RootNl().Name().Identifier))
	assert.Equal(t, obj.GroLock(), obj.Get(node.RootNl().GroLock().Identifier))
	assert.Panics(t, func() { obj2.GroLock() })
	assert.Nil(t, obj2.Get(node.RootNl().GroLock().Identifier))
}

func TestRootNl_QueryLoad(t *testing.T) {
	obj := createMinimalSampleRootNl()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleRootNl(ctx, obj)

	objs, err := QueryRootNls(ctx).
		Where(op.Equal(node.RootNl().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.RootNl().PrimaryKey()). // exercise order by
		Limit(1, 0).                         // exercise limit
		Calculation(node.RootNl(), "IsTrue", op.Equal("A", "A")).
		Load()
	assert.NoError(t, err)
	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestRootNl_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleRootNl()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleRootNl(ctx, obj)

	objs, _ := QueryRootNls(ctx).
		Where(op.Equal(node.RootNl().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestRootNl_QueryCursor(t *testing.T) {
	obj := createMinimalSampleRootNl()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleRootNl(ctx, obj)

	cursor, err := QueryRootNls(ctx).
		Where(op.Equal(node.RootNl().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()
	require.NoError(t, err)
	obj2, err2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	require.NoError(t, err2)
	obj2, err2 = cursor.Next()
	assert.Nil(t, obj2)
	require.NoError(t, err2)
	assert.NoError(t, cursor.Close())

	// test empty cursor result
	cursor, err = QueryRootNls(ctx).
		Where(op.Equal("B", "A")).
		LoadCursor()
	require.NoError(t, err)

	obj2, err = cursor.Next()
	assert.Nil(t, obj2)
	require.NoError(t, err)
	assert.NoError(t, cursor.Close())
}
func TestRootNl_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleRootNl(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleRootNl(ctx, obj)
	assert.Positive(t, func() int { i, _ := CountRootNls(ctx); return i }())

}

func TestRootNl_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleRootNl()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewRootNl()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsRootNl(t, obj, obj2)
}

func TestRootNl_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleRootNl()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewRootNl()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsRootNl(t, obj, obj2)
}

func TestRootNl_FailingMarshalBinary(t *testing.T) {
	obj := createMinimalSampleRootNl()
	var err error

	for i := 0; i < 13; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
	// do it again with aliases
	obj._aliases = make(map[string]any)
	for i := 0; i < 14; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
}

func TestRootNl_FailingUnmarshalBinary(t *testing.T) {
	obj := createMinimalSampleRootNl()
	b, err := obj.MarshalBinary()
	assert.NoError(t, err)
	obj2 := NewRootNl()
	for i := 0; i < 13; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}

	// do it again with aliases
	obj = createMinimalSampleRootNl()
	obj._aliases = map[string]any{"a": 1}
	b, err = obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 = NewRootNl()
	for i := 0; i < 14; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}
}
