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

// createMinimalSampleLeaf creates an unsaved minimal version of a Leaf object
// for testing.
func createMinimalSampleLeaf() *Leaf {
	obj := NewLeaf()
	updateMinimalSampleLeaf(obj)

	return obj
}

// updateMinimalSampleLeaf sets the values of a minimal sample to new, random values.
func updateMinimalSampleLeaf(obj *Leaf) {

	obj.SetName(test.RandomValue[string](100))

}

// createMaximalSampleLeaf creates an unsaved version of a Leaf object
// for testing that includes references to minimal objects.
func createMaximalSampleLeaf() *Leaf {
	obj := NewLeaf()
	updateMaximalSampleLeaf(obj)
	return obj
}

// updateMaximalSampleLeaf sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleLeaf(obj *Leaf) {
	updateMinimalSampleLeaf(obj)

	obj.SetOptionalLeafRoots(createMinimalSampleRoot())
	obj.SetRequiredLeafRoots(createMinimalSampleRoot())
	obj.SetOptionalLeafUniqueRoot(createMinimalSampleRoot())
	obj.SetRequiredLeafUniqueRoot(createMinimalSampleRoot())
}

// deleteSampleLeaf deletes an object created and saved by one of the sample creator functions.
func deleteSampleLeaf(ctx context.Context, obj *Leaf) {
	if obj == nil {
		return
	}

	for _, item := range obj.OptionalLeafRoots() {
		deleteSampleRoot(ctx, item)
	}
	for _, item := range obj.RequiredLeafRoots() {
		deleteSampleRoot(ctx, item)
	}
	deleteSampleRoot(ctx, obj.OptionalLeafUniqueRoot())
	deleteSampleRoot(ctx, obj.RequiredLeafUniqueRoot())

	obj.Delete(ctx)

}

// assertEqualFieldsLeaf compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsLeaf(t *testing.T, obj1, obj2 *Leaf) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.NameIsLoaded() && obj2.NameIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Name(), obj2.Name())
	}

}

func TestLeaf_SetID(t *testing.T) {

	obj := NewLeaf()

	assert.True(t, obj.IsNew())
	val := test.RandomNumberString()
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID("")
	assert.EqualValues(t, "", obj.ID(), "set default")

}
func TestLeaf_SetName(t *testing.T) {

	obj := NewLeaf()

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

func TestLeaf_Copy(t *testing.T) {
	obj := createMinimalSampleLeaf()

	obj2 := obj.Copy()

	assert.Equal(t, obj.Name(), obj2.Name())

}

func TestLeaf_BasicInsert(t *testing.T) {
	obj := createMinimalSampleLeaf()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeaf(ctx, obj)

	// Test retrieval
	obj2 := LoadLeaf(ctx, obj.PrimaryKey())
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

}

func TestLeaf_InsertPanics(t *testing.T) {
	obj := createMinimalSampleLeaf()
	ctx := db.NewContext(nil)

	obj.nameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsLoaded = true

}

func TestLeaf_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleLeaf()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeaf(ctx, obj)
	updateMinimalSampleLeaf(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadLeaf(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.Name(), obj.Name(), "Name did not update")
}

func TestLeaf_ReferenceLoad(t *testing.T) {
	obj := createMaximalSampleLeaf()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleLeaf(ctx, obj)

	// Test that referenced objects were saved and assigned ids

	// Test lazy loading
	obj2 := LoadLeaf(ctx, obj.PrimaryKey())
	objPkOnly := LoadLeaf(ctx, obj.PrimaryKey(), node.Leaf().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.OptionalLeafRoots(), "OptionalLeafRoots is not loaded initially")
	v_OptionalLeafRoots := obj2.LoadOptionalLeafRoots(ctx)
	assert.NotNil(t, v_OptionalLeafRoots)
	assert.Len(t, v_OptionalLeafRoots, 1)
	assert.Nil(t, obj2.RequiredLeafRoots(), "RequiredLeafRoots is not loaded initially")
	v_RequiredLeafRoots := obj2.LoadRequiredLeafRoots(ctx)
	assert.NotNil(t, v_RequiredLeafRoots)
	assert.Len(t, v_RequiredLeafRoots, 1)
	assert.Nil(t, obj2.OptionalLeafUniqueRoot(), "OptionalLeafUniqueRoot is not loaded initially")
	v_OptionalLeafUniqueRoot := obj2.LoadOptionalLeafUniqueRoot(ctx)
	assert.NotNil(t, v_OptionalLeafUniqueRoot)
	assert.Equal(t, v_OptionalLeafUniqueRoot.PrimaryKey(), obj2.OptionalLeafUniqueRoot().PrimaryKey())
	assert.Equal(t, obj.OptionalLeafUniqueRoot().PrimaryKey(), obj2.OptionalLeafUniqueRoot().PrimaryKey())
	assert.Nil(t, obj2.RequiredLeafUniqueRoot(), "RequiredLeafUniqueRoot is not loaded initially")
	v_RequiredLeafUniqueRoot := obj2.LoadRequiredLeafUniqueRoot(ctx)
	assert.NotNil(t, v_RequiredLeafUniqueRoot)
	assert.Equal(t, v_RequiredLeafUniqueRoot.PrimaryKey(), obj2.RequiredLeafUniqueRoot().PrimaryKey())
	assert.Equal(t, obj.RequiredLeafUniqueRoot().PrimaryKey(), obj2.RequiredLeafUniqueRoot().PrimaryKey())

	// test eager loading
	obj3 := LoadLeaf(ctx, obj.PrimaryKey(), node.Leaf().OptionalLeafRoots(),
		node.Leaf().RequiredLeafRoots(),
		node.Leaf().OptionalLeafUniqueRoot(),
		node.Leaf().RequiredLeafUniqueRoot(),
	)
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, len(obj2.OptionalLeafRoots()), len(obj3.OptionalLeafRoots()))
	assert.Equal(t, len(obj2.RequiredLeafRoots()), len(obj3.RequiredLeafRoots()))
	assert.Equal(t, obj2.OptionalLeafUniqueRoot().PrimaryKey(), obj3.OptionalLeafUniqueRoot().PrimaryKey())
	assert.Equal(t, obj2.RequiredLeafUniqueRoot().PrimaryKey(), obj3.RequiredLeafUniqueRoot().PrimaryKey())

}

func TestLeaf_ReferenceUpdateNewObjects(t *testing.T) {
	obj := createMaximalSampleLeaf()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleLeaf(ctx, obj)

	obj2 := LoadLeaf(ctx, obj.PrimaryKey())
	updateMaximalSampleLeaf(obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleLeaf(ctx, obj2)

	obj3 := LoadLeaf(ctx, obj2.PrimaryKey(), node.Leaf().OptionalLeafRoots(),
		node.Leaf().RequiredLeafRoots(),
		node.Leaf().OptionalLeafUniqueRoot(),
		node.Leaf().RequiredLeafUniqueRoot(),
	)
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, len(obj2.OptionalLeafRoots()), len(obj3.OptionalLeafRoots()))
	assert.Equal(t, len(obj2.RequiredLeafRoots()), len(obj3.RequiredLeafRoots()))
	assert.Equal(t, obj2.OptionalLeafUniqueRoot().PrimaryKey(), obj3.OptionalLeafUniqueRoot().PrimaryKey())
	assert.Equal(t, obj2.RequiredLeafUniqueRoot().PrimaryKey(), obj3.RequiredLeafUniqueRoot().PrimaryKey())

}

func TestLeaf_ReferenceUpdateOldObjects(t *testing.T) {
	obj := createMaximalSampleLeaf()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleLeaf(ctx, obj)

	updateMinimalSampleRoot(obj.OptionalLeafRoots()[0])
	updateMinimalSampleRoot(obj.RequiredLeafRoots()[0])
	updateMinimalSampleRoot(obj.OptionalLeafUniqueRoot())
	updateMinimalSampleRoot(obj.RequiredLeafUniqueRoot())

	assert.NoError(t, obj.Save(ctx))

	obj2 := LoadLeaf(ctx, obj.PrimaryKey(),

		node.Leaf().OptionalLeafRoots(),
		node.Leaf().RequiredLeafRoots(),
		node.Leaf().OptionalLeafUniqueRoot(),
		node.Leaf().RequiredLeafUniqueRoot(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsRoot(t, obj2.OptionalLeafRoots()[0], obj.OptionalLeafRoots()[0])
	assertEqualFieldsRoot(t, obj2.RequiredLeafRoots()[0], obj.RequiredLeafRoots()[0])
	assertEqualFieldsRoot(t, obj2.OptionalLeafUniqueRoot(), obj.OptionalLeafUniqueRoot())
	assertEqualFieldsRoot(t, obj2.RequiredLeafUniqueRoot(), obj.RequiredLeafUniqueRoot())

}
func TestLeaf_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewLeaf()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestLeaf_Getters(t *testing.T) {
	obj := createMinimalSampleLeaf()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleLeaf(ctx, obj)

	assert.True(t, HasLeaf(ctx, obj.PrimaryKey()))

	obj2 := LoadLeaf(ctx, obj.PrimaryKey(), node.Leaf().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.Leaf().ID().Identifier))
	assert.Equal(t, obj.Name(), obj.Get(node.Leaf().Name().Identifier))
	assert.Panics(t, func() { obj2.Name() })
	assert.Nil(t, obj2.Get(node.Leaf().Name().Identifier))
}

func TestLeaf_QueryLoad(t *testing.T) {
	obj := createMinimalSampleLeaf()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeaf(ctx, obj)

	objs := QueryLeafs(ctx).
		Where(op.Equal(node.Leaf().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.Leaf().PrimaryKey()). // exercise order by
		Limit(1, 0).                       // exercise limit
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
}
func TestLeaf_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleLeaf()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeaf(ctx, obj)

	objs := QueryLeafs(ctx).
		Where(op.Equal(node.Leaf().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestLeaf_QueryCursor(t *testing.T) {
	obj := createMinimalSampleLeaf()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeaf(ctx, obj)

	cursor := QueryLeafs(ctx).
		Where(op.Equal(node.Leaf().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryLeafs(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestLeaf_Count(t *testing.T) {
	obj := createMaximalSampleLeaf()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleLeaf(ctx, obj)

	assert.Less(t, 0, CountLeafs(ctx))

	assert.Less(t, 0, CountLeafsByID(ctx, obj.ID()))
	assert.Less(t, 0, CountLeafsByName(ctx, obj.Name()))

}
func TestLeaf_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleLeaf()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewLeaf()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsLeaf(t, obj, obj2)
}

func TestLeaf_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleLeaf()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewLeaf()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsLeaf(t, obj, obj2)
}
