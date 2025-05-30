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

// createMinimalSampleLeafL creates an unsaved minimal version of a LeafL object
// for testing.
func createMinimalSampleLeafL() *LeafL {
	obj := NewLeafL()
	updateMinimalSampleLeafL(obj)

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	// If the database is configured so that the referenced object points back here, either directly or through multiple
	// forward references, it possible this could create an endless loop. Such a structure could not be saved anyways.
	obj.SetRootL(createMinimalSampleRootL())

	return obj
}

// updateMinimalSampleLeafL sets the values of a minimal sample to new, random values.
func updateMinimalSampleLeafL(obj *LeafL) {

	obj.SetName(test.RandomValue[string](100))

}

// createMaximalSampleLeafL creates an unsaved version of a LeafL object
// for testing that includes references to minimal objects.
func createMaximalSampleLeafL(ctx context.Context) *LeafL {
	obj := NewLeafL()
	updateMaximalSampleLeafL(ctx, obj)
	return obj
}

// updateMaximalSampleLeafL sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleLeafL(ctx context.Context, obj *LeafL) {
	updateMinimalSampleLeafL(obj)
	obj.SetRootL(createMinimalSampleRootL())

}

// deleteSampleLeafL deletes an object created and saved by one of the sample creator functions.
func deleteSampleLeafL(ctx context.Context, obj *LeafL) {
	if obj == nil {
		return
	}

	_ = obj.Delete(ctx)

	deleteSampleRootL(ctx, obj.RootL())

}

// assertEqualFieldsLeafL compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsLeafL(t *testing.T, obj1, obj2 *LeafL) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.NameIsLoaded() && obj2.NameIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Name(), obj2.Name())
	}
	if obj1.GroLockIsLoaded() && obj2.GroLockIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.GroLock(), obj2.GroLock())
	}
	if obj1.RootLIDIsLoaded() && obj2.RootLIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.RootLID(), obj2.RootLID())
	}

}

func TestLeafL_SetID(t *testing.T) {

	obj := NewLeafL()

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
func TestLeafL_SetName(t *testing.T) {

	obj := NewLeafL()

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
func TestLeafL_SetRootLID(t *testing.T) {

	obj := NewLeafL()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetRootLID(val)
	assert.Equal(t, val, obj.RootLID())

	// test default
	obj.SetRootLID("")
	assert.EqualValues(t, "", obj.RootLID(), "set default")

}

func TestLeafL_Copy(t *testing.T) {
	obj := createMinimalSampleLeafL()

	obj2 := obj.Copy()

	assert.Equal(t, obj.Name(), obj2.Name())
	assert.Equal(t, obj.RootLID(), obj2.RootLID())

}

func TestLeafL_BasicInsert(t *testing.T) {
	obj := createMinimalSampleLeafL()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafL(ctx, obj)

	// Test retrieval
	obj2, err := LoadLeafL(ctx, obj.PrimaryKey())
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

func TestLeafL_InsertPanics(t *testing.T) {
	obj := createMinimalSampleLeafL()
	_ = obj
	ctx := db.NewContext(nil)
	_ = ctx

	obj.nameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsLoaded = true

	obj.objRootL = nil
	obj.rootLIDIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.rootLIDIsLoaded = true

}

func TestLeafL_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleLeafL()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafL(ctx, obj)
	updateMinimalSampleLeafL(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2, err := LoadLeafL(ctx, obj.PrimaryKey())
	assert.NoError(t, err)

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.Name(), obj.Name(), "Name did not update")
	assert.Equal(t, obj2.GroLock(), obj.GroLock(), "GroLock did not update")
}

func TestLeafL_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafL(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafL(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.RootL())
	assert.NotEqual(t, '-', obj.RootL().PrimaryKey()[0])

	// Test lazy loading
	obj2, err := LoadLeafL(ctx, obj.PrimaryKey())
	assert.NoError(t, err)
	objPkOnly, err2 := LoadLeafL(ctx, obj.PrimaryKey(), node.LeafL().PrimaryKey())
	assert.NoError(t, err2)
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.RootL(), "RootL is not loaded initially")
	v_RootL, _ := obj2.LoadRootL(ctx)
	assert.NotNil(t, v_RootL)
	assert.Equal(t, v_RootL.PrimaryKey(), obj2.RootL().PrimaryKey())
	assert.Equal(t, obj.RootL().PrimaryKey(), obj2.RootL().PrimaryKey())
	assert.True(t, obj2.RootLIDIsLoaded())

	assert.False(t, objPkOnly.RootLIDIsLoaded())
	assert.Panics(t, func() { _, _ = objPkOnly.LoadRootL(ctx) })

	assert.Panics(t, func() {
		objPkOnly.SetRootL(nil)
	})

	// test eager loading
	obj3, _ := LoadLeafL(ctx, obj.PrimaryKey(), node.LeafL().RootL())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.RootL().PrimaryKey(), obj3.RootL().PrimaryKey())

}

func TestLeafL_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafL(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafL(ctx, obj)

	obj2, _ := LoadLeafL(ctx, obj.PrimaryKey())
	updateMaximalSampleLeafL(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleLeafL(ctx, obj2)

	obj3, _ := LoadLeafL(ctx, obj2.PrimaryKey(), node.LeafL().RootL())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.RootL().PrimaryKey(), obj3.RootL().PrimaryKey())

}

func TestLeafL_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafL(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafL(ctx, obj)

	updateMinimalSampleRootL(obj.RootL())

	assert.NoError(t, obj.Save(ctx))

	obj2, _ := LoadLeafL(ctx, obj.PrimaryKey(),

		node.LeafL().RootL(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsRootL(t, obj2.RootL(), obj.RootL())

}
func TestLeafL_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewLeafL()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestLeafL_Getters(t *testing.T) {
	obj := createMinimalSampleLeafL()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafL(ctx, obj)

	has, _ := HasLeafL(ctx, obj.PrimaryKey())
	assert.True(t, has)

	obj2, _ := LoadLeafL(ctx, obj.PrimaryKey(), node.LeafL().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.LeafL().ID().Identifier))
	assert.Equal(t, obj.Name(), obj.Get(node.LeafL().Name().Identifier))
	assert.Panics(t, func() { obj2.Name() })
	assert.Nil(t, obj2.Get(node.LeafL().Name().Identifier))
	assert.Equal(t, obj.GroLock(), obj.Get(node.LeafL().GroLock().Identifier))
	assert.Panics(t, func() { obj2.GroLock() })
	assert.Nil(t, obj2.Get(node.LeafL().GroLock().Identifier))
	assert.Panics(t, func() { obj2.RootLID() })
	assert.Nil(t, obj2.Get(node.LeafL().RootLID().Identifier))
}

func TestLeafL_QueryLoad(t *testing.T) {
	obj := createMinimalSampleLeafL()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafL(ctx, obj)

	objs, err := QueryLeafLs(ctx).
		Where(op.Equal(node.LeafL().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.LeafL().PrimaryKey()). // exercise order by
		Limit(1, 0).                        // exercise limit
		Calculation(node.LeafL(), "IsTrue", op.Equal("A", "A")).
		Load()
	assert.NoError(t, err)
	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestLeafL_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleLeafL()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeafL(ctx, obj)

	objs, _ := QueryLeafLs(ctx).
		Where(op.Equal(node.LeafL().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestLeafL_QueryCursor(t *testing.T) {
	obj := createMinimalSampleLeafL()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeafL(ctx, obj)

	cursor, err := QueryLeafLs(ctx).
		Where(op.Equal(node.LeafL().PrimaryKey(), obj.PrimaryKey())).
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
	cursor, err = QueryLeafLs(ctx).
		Where(op.Equal("B", "A")).
		LoadCursor()
	require.NoError(t, err)

	obj2, err = cursor.Next()
	assert.Nil(t, obj2)
	require.NoError(t, err)
	assert.NoError(t, cursor.Close())
}
func TestLeafL_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleLeafL(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeafL(ctx, obj)
	assert.Positive(t, func() int { i, _ := CountLeafLs(ctx); return i }())

	// reread in case there are data limitations imposed by the database
	obj2, _ := LoadLeafL(ctx, obj.PrimaryKey())
	assert.Positive(t,
		func() int {
			i, _ := CountLeafLsByRootLID(ctx,
				obj2.RootLID())
			return i
		}())

}

func TestLeafL_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleLeafL()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewLeafL()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsLeafL(t, obj, obj2)
}

func TestLeafL_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleLeafL()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewLeafL()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsLeafL(t, obj, obj2)
}

func TestLeafL_FailingMarshalBinary(t *testing.T) {
	obj := createMinimalSampleLeafL()
	var err error

	for i := 0; i < 14; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
	// do it again with aliases
	obj._aliases = make(map[string]any)
	for i := 0; i < 15; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
}

func TestLeafL_FailingUnmarshalBinary(t *testing.T) {
	obj := createMinimalSampleLeafL()
	b, err := obj.MarshalBinary()
	assert.NoError(t, err)
	obj2 := NewLeafL()
	for i := 0; i < 14; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}

	// do it again with aliases
	obj = createMinimalSampleLeafL()
	obj._aliases = map[string]any{"a": 1}
	b, err = obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 = NewLeafL()
	for i := 0; i < 15; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}
}
