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

// createMinimalSampleLeafN creates an unsaved minimal version of a LeafN object
// for testing.
func createMinimalSampleLeafN() *LeafN {
	obj := NewLeafN()
	updateMinimalSampleLeafN(obj)

	return obj
}

// updateMinimalSampleLeafN sets the values of a minimal sample to new, random values.
func updateMinimalSampleLeafN(obj *LeafN) {

	obj.SetName(test.RandomValue[string](100))

}

// createMaximalSampleLeafN creates an unsaved version of a LeafN object
// for testing that includes references to minimal objects.
func createMaximalSampleLeafN(ctx context.Context) *LeafN {
	obj := NewLeafN()
	updateMaximalSampleLeafN(ctx, obj)
	return obj
}

// updateMaximalSampleLeafN sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleLeafN(ctx context.Context, obj *LeafN) {
	updateMinimalSampleLeafN(obj)
	obj.SetRootN(createMinimalSampleRootN())

}

// deleteSampleLeafN deletes an object created and saved by one of the sample creator functions.
func deleteSampleLeafN(ctx context.Context, obj *LeafN) {
	if obj == nil {
		return
	}

	_ = obj.Delete(ctx)

	deleteSampleRootN(ctx, obj.RootN())

}

// assertEqualFieldsLeafN compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsLeafN(t *testing.T, obj1, obj2 *LeafN) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.NameIsLoaded() && obj2.NameIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Name(), obj2.Name())
	}
	if obj1.RootNIDIsLoaded() && obj2.RootNIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.RootNID(), obj2.RootNID())
	}

}

func TestLeafN_SetID(t *testing.T) {

	obj := NewLeafN()

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
func TestLeafN_SetName(t *testing.T) {

	obj := NewLeafN()

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
func TestLeafN_SetRootNID(t *testing.T) {

	obj := NewLeafN()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetRootNID(val)
	assert.Equal(t, val, obj.RootNID())
	assert.False(t, obj.RootNIDIsNull())

	// Test NULL
	obj.SetRootNIDToNull()
	assert.EqualValues(t, "", obj.RootNID())
	assert.True(t, obj.RootNIDIsNull())

	// test default
	obj.SetRootNID("")
	assert.EqualValues(t, "", obj.RootNID(), "set default")

}

func TestLeafN_Copy(t *testing.T) {
	obj := createMinimalSampleLeafN()

	obj2 := obj.Copy()

	assert.Equal(t, obj.Name(), obj2.Name())
	assert.Equal(t, obj.RootNID(), obj2.RootNID())

}

func TestLeafN_BasicInsert(t *testing.T) {
	obj := createMinimalSampleLeafN()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafN(ctx, obj)

	// Test retrieval
	obj2, err := LoadLeafN(ctx, obj.PrimaryKey())
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

}

func TestLeafN_InsertPanics(t *testing.T) {
	obj := createMinimalSampleLeafN()
	_ = obj
	ctx := db.NewContext(nil)
	_ = ctx

	obj.nameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsLoaded = true

}

func TestLeafN_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleLeafN()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafN(ctx, obj)
	updateMinimalSampleLeafN(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2, err := LoadLeafN(ctx, obj.PrimaryKey())
	assert.NoError(t, err)

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.Name(), obj.Name(), "Name did not update")
}

func TestLeafN_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafN(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafN(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.RootN())
	assert.NotEqual(t, '-', obj.RootN().PrimaryKey()[0])

	// Test lazy loading
	obj2, err := LoadLeafN(ctx, obj.PrimaryKey())
	assert.NoError(t, err)
	objPkOnly, err2 := LoadLeafN(ctx, obj.PrimaryKey(), node.LeafN().PrimaryKey())
	assert.NoError(t, err2)
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.RootN(), "RootN is not loaded initially")
	v_RootN, _ := obj2.LoadRootN(ctx)
	assert.NotNil(t, v_RootN)
	assert.Equal(t, v_RootN.PrimaryKey(), obj2.RootN().PrimaryKey())
	assert.Equal(t, obj.RootN().PrimaryKey(), obj2.RootN().PrimaryKey())
	assert.True(t, obj2.RootNIDIsLoaded())

	assert.False(t, objPkOnly.RootNIDIsLoaded())
	assert.Panics(t, func() { _, _ = objPkOnly.LoadRootN(ctx) })

	// test eager loading
	obj3, _ := LoadLeafN(ctx, obj.PrimaryKey(), node.LeafN().RootN())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.RootN().PrimaryKey(), obj3.RootN().PrimaryKey())

}

func TestLeafN_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafN(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafN(ctx, obj)

	obj2, _ := LoadLeafN(ctx, obj.PrimaryKey())
	updateMaximalSampleLeafN(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleLeafN(ctx, obj2)

	obj3, _ := LoadLeafN(ctx, obj2.PrimaryKey(), node.LeafN().RootN())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.RootN().PrimaryKey(), obj3.RootN().PrimaryKey())

}

func TestLeafN_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafN(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafN(ctx, obj)

	updateMinimalSampleRootN(obj.RootN())

	assert.NoError(t, obj.Save(ctx))

	obj2, _ := LoadLeafN(ctx, obj.PrimaryKey(),

		node.LeafN().RootN(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsRootN(t, obj2.RootN(), obj.RootN())

}
func TestLeafN_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewLeafN()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestLeafN_Getters(t *testing.T) {
	obj := createMinimalSampleLeafN()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafN(ctx, obj)

	has, _ := HasLeafN(ctx, obj.PrimaryKey())
	assert.True(t, has)

	obj2, _ := LoadLeafN(ctx, obj.PrimaryKey(), node.LeafN().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.LeafN().ID().Identifier))
	assert.Equal(t, obj.Name(), obj.Get(node.LeafN().Name().Identifier))
	assert.Panics(t, func() { obj2.Name() })
	assert.Nil(t, obj2.Get(node.LeafN().Name().Identifier))
	assert.Panics(t, func() { obj2.RootNID() })
	assert.Nil(t, obj2.Get(node.LeafN().RootNID().Identifier))
}

func TestLeafN_QueryLoad(t *testing.T) {
	obj := createMinimalSampleLeafN()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafN(ctx, obj)

	objs, err := QueryLeafNs(ctx).
		Where(op.Equal(node.LeafN().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.LeafN().PrimaryKey()). // exercise order by
		Limit(1, 0).                        // exercise limit
		Calculation(node.LeafN(), "IsTrue", op.Equal("A", "A")).
		Load()
	assert.NoError(t, err)
	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestLeafN_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleLeafN()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeafN(ctx, obj)

	objs, _ := QueryLeafNs(ctx).
		Where(op.Equal(node.LeafN().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestLeafN_QueryCursor(t *testing.T) {
	obj := createMinimalSampleLeafN()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafN(ctx, obj)

	cursor, err := QueryLeafNs(ctx).
		Where(op.Equal(node.LeafN().PrimaryKey(), obj.PrimaryKey())).
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
	cursor, err = QueryLeafNs(ctx).
		Where(op.Equal("B", "A")).
		LoadCursor()
	require.NoError(t, err)

	obj2, err = cursor.Next()
	assert.Nil(t, obj2)
	require.NoError(t, err)
	assert.NoError(t, cursor.Close())
}
func TestLeafN_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafN(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeafN(ctx, obj)
	assert.Positive(t, func() int { i, _ := CountLeafNs(ctx); return i }())

	// reread in case there are data limitations imposed by the database
	obj2, _ := LoadLeafN(ctx, obj.PrimaryKey())
	assert.Positive(t,
		func() int {
			i, _ := CountLeafNsByRootNID(ctx,
				obj2.RootNID())
			return i
		}())

}

func TestLeafN_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleLeafN()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewLeafN()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsLeafN(t, obj, obj2)
}

func TestLeafN_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleLeafN()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewLeafN()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsLeafN(t, obj, obj2)
}

func TestLeafN_FailingMarshalBinary(t *testing.T) {
	obj := createMinimalSampleLeafN()
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

func TestLeafN_FailingUnmarshalBinary(t *testing.T) {
	obj := createMinimalSampleLeafN()
	b, err := obj.MarshalBinary()
	assert.NoError(t, err)
	obj2 := NewLeafN()
	for i := 0; i < 13; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}

	// do it again with aliases
	obj = createMinimalSampleLeafN()
	obj._aliases = map[string]any{"a": 1}
	b, err = obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 = NewLeafN()
	for i := 0; i < 14; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}
}
