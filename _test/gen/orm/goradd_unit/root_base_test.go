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

// createMinimalSampleRoot creates an unsaved minimal version of a Root object
// for testing.
func createMinimalSampleRoot() *Root {
	obj := NewRoot()
	updateMinimalSampleRoot(obj)

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	// If the database is configured so that the referenced object points back here, either directly or through multiple
	// forward references, it possible this could create an endless loop. Such a structure could not be saved anyways.
	obj.SetRequiredLeaf(createMinimalSampleLeaf())

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	// If the database is configured so that the referenced object points back here, either directly or through multiple
	// forward references, it possible this could create an endless loop. Such a structure could not be saved anyways.
	obj.SetOptionalLeafUnique(createMinimalSampleLeaf())

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	// If the database is configured so that the referenced object points back here, either directly or through multiple
	// forward references, it possible this could create an endless loop. Such a structure could not be saved anyways.
	obj.SetRequiredLeafUnique(createMinimalSampleLeaf())

	return obj
}

// updateMinimalSampleRoot sets the values of a minimal sample to new, random values.
func updateMinimalSampleRoot(obj *Root) {

	obj.SetName(test.RandomValue[string](100))

}

// createMaximalSampleRoot creates an unsaved version of a Root object
// for testing that includes references to minimal objects.
func createMaximalSampleRoot(ctx context.Context) *Root {
	obj := NewRoot()
	updateMaximalSampleRoot(ctx, obj)
	return obj
}

// updateMaximalSampleRoot sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleRoot(ctx context.Context, obj *Root) {
	updateMinimalSampleRoot(obj)
	obj.SetOptionalLeaf(createMinimalSampleLeaf())
	obj.SetRequiredLeaf(createMinimalSampleLeaf())
	obj.SetOptionalLeafUnique(createMinimalSampleLeaf())
	obj.SetRequiredLeafUnique(createMinimalSampleLeaf())
	obj.SetParent(createMinimalSampleRoot())

	obj.SetParentRoots(createMinimalSampleRoot())
}

// deleteSampleRoot deletes an object created and saved by one of the sample creator functions.
func deleteSampleRoot(ctx context.Context, obj *Root) {
	if obj == nil {
		return
	}

	for _, item := range obj.ParentRoots() {
		deleteSampleRoot(ctx, item)
	}

	obj.Delete(ctx)

	deleteSampleLeaf(ctx, obj.OptionalLeaf())

	deleteSampleLeaf(ctx, obj.RequiredLeaf())

	deleteSampleLeaf(ctx, obj.OptionalLeafUnique())

	deleteSampleLeaf(ctx, obj.RequiredLeafUnique())

	deleteSampleRoot(ctx, obj.Parent())

}

// assertEqualFieldsRoot compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsRoot(t *testing.T, obj1, obj2 *Root) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.NameIsLoaded() && obj2.NameIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Name(), obj2.Name())
	}
	if obj1.OptionalLeafIDIsLoaded() && obj2.OptionalLeafIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.OptionalLeafID(), obj2.OptionalLeafID())
	}
	if obj1.RequiredLeafIDIsLoaded() && obj2.RequiredLeafIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.RequiredLeafID(), obj2.RequiredLeafID())
	}
	if obj1.OptionalLeafUniqueIDIsLoaded() && obj2.OptionalLeafUniqueIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.OptionalLeafUniqueID(), obj2.OptionalLeafUniqueID())
	}
	if obj1.RequiredLeafUniqueIDIsLoaded() && obj2.RequiredLeafUniqueIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.RequiredLeafUniqueID(), obj2.RequiredLeafUniqueID())
	}
	if obj1.ParentIDIsLoaded() && obj2.ParentIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ParentID(), obj2.ParentID())
	}

}

func TestRoot_SetID(t *testing.T) {

	obj := NewRoot()

	assert.True(t, obj.IsNew())
	val := test.RandomNumberString()
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID("")
	assert.EqualValues(t, "", obj.ID(), "set default")

}
func TestRoot_SetName(t *testing.T) {

	obj := NewRoot()

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
func TestRoot_SetOptionalLeafID(t *testing.T) {

	obj := NewRoot()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetOptionalLeafID(val)
	assert.Equal(t, val, obj.OptionalLeafID())
	assert.False(t, obj.OptionalLeafIDIsNull())

	// Test NULL
	obj.SetOptionalLeafIDToNull()
	assert.EqualValues(t, "", obj.OptionalLeafID())
	assert.True(t, obj.OptionalLeafIDIsNull())

	// test default
	obj.SetOptionalLeafID("")
	assert.EqualValues(t, "", obj.OptionalLeafID(), "set default")

}
func TestRoot_SetRequiredLeafID(t *testing.T) {

	obj := NewRoot()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetRequiredLeafID(val)
	assert.Equal(t, val, obj.RequiredLeafID())

	// test default
	obj.SetRequiredLeafID("")
	assert.EqualValues(t, "", obj.RequiredLeafID(), "set default")

}
func TestRoot_SetOptionalLeafUniqueID(t *testing.T) {

	obj := NewRoot()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetOptionalLeafUniqueID(val)
	assert.Equal(t, val, obj.OptionalLeafUniqueID())

	// test default
	obj.SetOptionalLeafUniqueID("")
	assert.EqualValues(t, "", obj.OptionalLeafUniqueID(), "set default")

}
func TestRoot_SetRequiredLeafUniqueID(t *testing.T) {

	obj := NewRoot()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetRequiredLeafUniqueID(val)
	assert.Equal(t, val, obj.RequiredLeafUniqueID())

	// test default
	obj.SetRequiredLeafUniqueID("")
	assert.EqualValues(t, "", obj.RequiredLeafUniqueID(), "set default")

}
func TestRoot_SetParentID(t *testing.T) {

	obj := NewRoot()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetParentID(val)
	assert.Equal(t, val, obj.ParentID())
	assert.False(t, obj.ParentIDIsNull())

	// Test NULL
	obj.SetParentIDToNull()
	assert.EqualValues(t, "", obj.ParentID())
	assert.True(t, obj.ParentIDIsNull())

	// test default
	obj.SetParentID("")
	assert.EqualValues(t, "", obj.ParentID(), "set default")

}

func TestRoot_Copy(t *testing.T) {
	obj := createMinimalSampleRoot()

	obj2 := obj.Copy()

	assert.Equal(t, obj.Name(), obj2.Name())
	assert.Equal(t, obj.OptionalLeafID(), obj2.OptionalLeafID())
	assert.Equal(t, obj.RequiredLeafID(), obj2.RequiredLeafID())
	assert.Equal(t, obj.OptionalLeafUniqueID(), obj2.OptionalLeafUniqueID())
	assert.Equal(t, obj.RequiredLeafUniqueID(), obj2.RequiredLeafUniqueID())
	assert.Equal(t, obj.ParentID(), obj2.ParentID())

}

func TestRoot_BasicInsert(t *testing.T) {
	obj := createMinimalSampleRoot()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleRoot(ctx, obj)

	// Test retrieval
	obj2 := LoadRoot(ctx, obj.PrimaryKey())
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

func TestRoot_InsertPanics(t *testing.T) {
	obj := createMinimalSampleRoot()
	_ = obj
	ctx := db.NewContext(nil)
	_ = ctx

	obj.nameIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.nameIsLoaded = true

	obj.requiredLeafIDIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.requiredLeafIDIsLoaded = true

	obj.optionalLeafUniqueIDIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.optionalLeafUniqueIDIsLoaded = true

	obj.requiredLeafUniqueIDIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.requiredLeafUniqueIDIsLoaded = true

}

func TestRoot_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleRoot()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleRoot(ctx, obj)
	updateMinimalSampleRoot(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadRoot(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.Name(), obj.Name(), "Name did not update")
}

func TestRoot_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleRoot(ctx)
	obj.Save(ctx)
	defer deleteSampleRoot(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.OptionalLeaf())
	assert.NotEqual(t, '-', obj.OptionalLeaf().PrimaryKey()[0])

	assert.NotNil(t, obj.RequiredLeaf())
	assert.NotEqual(t, '-', obj.RequiredLeaf().PrimaryKey()[0])

	assert.NotNil(t, obj.OptionalLeafUnique())
	assert.NotEqual(t, '-', obj.OptionalLeafUnique().PrimaryKey()[0])

	assert.NotNil(t, obj.RequiredLeafUnique())
	assert.NotEqual(t, '-', obj.RequiredLeafUnique().PrimaryKey()[0])

	assert.NotNil(t, obj.Parent())
	assert.NotEqual(t, '-', obj.Parent().PrimaryKey()[0])

	// Test lazy loading
	obj2 := LoadRoot(ctx, obj.PrimaryKey())
	objPkOnly := LoadRoot(ctx, obj.PrimaryKey(), node.Root().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.OptionalLeaf(), "OptionalLeaf is not loaded initially")
	v_OptionalLeaf := obj2.LoadOptionalLeaf(ctx)
	assert.NotNil(t, v_OptionalLeaf)
	assert.Equal(t, v_OptionalLeaf.PrimaryKey(), obj2.OptionalLeaf().PrimaryKey())
	assert.Equal(t, obj.OptionalLeaf().PrimaryKey(), obj2.OptionalLeaf().PrimaryKey())
	assert.True(t, obj2.OptionalLeafIDIsLoaded())

	assert.False(t, objPkOnly.OptionalLeafIDIsLoaded())
	assert.Nil(t, objPkOnly.LoadOptionalLeaf(ctx))

	assert.Nil(t, obj2.RequiredLeaf(), "RequiredLeaf is not loaded initially")
	v_RequiredLeaf := obj2.LoadRequiredLeaf(ctx)
	assert.NotNil(t, v_RequiredLeaf)
	assert.Equal(t, v_RequiredLeaf.PrimaryKey(), obj2.RequiredLeaf().PrimaryKey())
	assert.Equal(t, obj.RequiredLeaf().PrimaryKey(), obj2.RequiredLeaf().PrimaryKey())
	assert.True(t, obj2.RequiredLeafIDIsLoaded())

	assert.False(t, objPkOnly.RequiredLeafIDIsLoaded())
	assert.Nil(t, objPkOnly.LoadRequiredLeaf(ctx))

	assert.Panics(t, func() {
		objPkOnly.SetRequiredLeaf(nil)
	})

	assert.Nil(t, obj2.OptionalLeafUnique(), "OptionalLeafUnique is not loaded initially")
	v_OptionalLeafUnique := obj2.LoadOptionalLeafUnique(ctx)
	assert.NotNil(t, v_OptionalLeafUnique)
	assert.Equal(t, v_OptionalLeafUnique.PrimaryKey(), obj2.OptionalLeafUnique().PrimaryKey())
	assert.Equal(t, obj.OptionalLeafUnique().PrimaryKey(), obj2.OptionalLeafUnique().PrimaryKey())
	assert.True(t, obj2.OptionalLeafUniqueIDIsLoaded())

	assert.False(t, objPkOnly.OptionalLeafUniqueIDIsLoaded())
	assert.Nil(t, objPkOnly.LoadOptionalLeafUnique(ctx))

	assert.Panics(t, func() {
		objPkOnly.SetOptionalLeafUnique(nil)
	})

	assert.Nil(t, obj2.RequiredLeafUnique(), "RequiredLeafUnique is not loaded initially")
	v_RequiredLeafUnique := obj2.LoadRequiredLeafUnique(ctx)
	assert.NotNil(t, v_RequiredLeafUnique)
	assert.Equal(t, v_RequiredLeafUnique.PrimaryKey(), obj2.RequiredLeafUnique().PrimaryKey())
	assert.Equal(t, obj.RequiredLeafUnique().PrimaryKey(), obj2.RequiredLeafUnique().PrimaryKey())
	assert.True(t, obj2.RequiredLeafUniqueIDIsLoaded())

	assert.False(t, objPkOnly.RequiredLeafUniqueIDIsLoaded())
	assert.Nil(t, objPkOnly.LoadRequiredLeafUnique(ctx))

	assert.Panics(t, func() {
		objPkOnly.SetRequiredLeafUnique(nil)
	})

	assert.Nil(t, obj2.Parent(), "Parent is not loaded initially")
	v_Parent := obj2.LoadParent(ctx)
	assert.NotNil(t, v_Parent)
	assert.Equal(t, v_Parent.PrimaryKey(), obj2.Parent().PrimaryKey())
	assert.Equal(t, obj.Parent().PrimaryKey(), obj2.Parent().PrimaryKey())
	assert.True(t, obj2.ParentIDIsLoaded())

	assert.False(t, objPkOnly.ParentIDIsLoaded())
	assert.Nil(t, objPkOnly.LoadParent(ctx))

	assert.Nil(t, obj2.ParentRoots(), "ParentRoots is not loaded initially")
	v_ParentRoots := obj2.LoadParentRoots(ctx)
	assert.NotNil(t, v_ParentRoots)
	assert.Len(t, v_ParentRoots, 1)

	// test eager loading
	obj3 := LoadRoot(ctx, obj.PrimaryKey(), node.Root().OptionalLeaf(),
		node.Root().RequiredLeaf(),
		node.Root().OptionalLeafUnique(),
		node.Root().RequiredLeafUnique(),
		node.Root().Parent(),
		node.Root().ParentRoots(),
	)
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.OptionalLeaf().PrimaryKey(), obj3.OptionalLeaf().PrimaryKey())
	assert.Equal(t, obj2.RequiredLeaf().PrimaryKey(), obj3.RequiredLeaf().PrimaryKey())
	assert.Equal(t, obj2.OptionalLeafUnique().PrimaryKey(), obj3.OptionalLeafUnique().PrimaryKey())
	assert.Equal(t, obj2.RequiredLeafUnique().PrimaryKey(), obj3.RequiredLeafUnique().PrimaryKey())
	assert.Equal(t, obj2.Parent().PrimaryKey(), obj3.Parent().PrimaryKey())
	assert.Equal(t, len(obj2.ParentRoots()), len(obj3.ParentRoots()))

}

func TestRoot_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleRoot(ctx)
	obj.Save(ctx)
	defer deleteSampleRoot(ctx, obj)

	obj2 := LoadRoot(ctx, obj.PrimaryKey())
	updateMaximalSampleRoot(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleRoot(ctx, obj2)

	obj3 := LoadRoot(ctx, obj2.PrimaryKey(), node.Root().OptionalLeaf(),
		node.Root().RequiredLeaf(),
		node.Root().OptionalLeafUnique(),
		node.Root().RequiredLeafUnique(),
		node.Root().Parent(),
		node.Root().ParentRoots(),
	)
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.OptionalLeaf().PrimaryKey(), obj3.OptionalLeaf().PrimaryKey())
	assert.Equal(t, obj2.RequiredLeaf().PrimaryKey(), obj3.RequiredLeaf().PrimaryKey())
	assert.Equal(t, obj2.OptionalLeafUnique().PrimaryKey(), obj3.OptionalLeafUnique().PrimaryKey())
	assert.Equal(t, obj2.RequiredLeafUnique().PrimaryKey(), obj3.RequiredLeafUnique().PrimaryKey())
	assert.Equal(t, obj2.Parent().PrimaryKey(), obj3.Parent().PrimaryKey())

	assert.Equal(t, len(obj2.ParentRoots()), len(obj3.ParentRoots()))

}

func TestRoot_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleRoot(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleRoot(ctx, obj)

	updateMinimalSampleLeaf(obj.OptionalLeaf())
	updateMinimalSampleLeaf(obj.RequiredLeaf())
	updateMinimalSampleLeaf(obj.OptionalLeafUnique())
	updateMinimalSampleLeaf(obj.RequiredLeafUnique())
	updateMinimalSampleRoot(obj.Parent())
	updateMinimalSampleRoot(obj.ParentRoots()[0])

	assert.NoError(t, obj.Save(ctx))

	obj2 := LoadRoot(ctx, obj.PrimaryKey(),

		node.Root().OptionalLeaf(),

		node.Root().RequiredLeaf(),

		node.Root().OptionalLeafUnique(),

		node.Root().RequiredLeafUnique(),

		node.Root().Parent(),

		node.Root().ParentRoots(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsLeaf(t, obj2.OptionalLeaf(), obj.OptionalLeaf())
	assertEqualFieldsLeaf(t, obj2.RequiredLeaf(), obj.RequiredLeaf())
	assertEqualFieldsLeaf(t, obj2.OptionalLeafUnique(), obj.OptionalLeafUnique())
	assertEqualFieldsLeaf(t, obj2.RequiredLeafUnique(), obj.RequiredLeafUnique())
	assertEqualFieldsRoot(t, obj2.Parent(), obj.Parent())

	assertEqualFieldsRoot(t, obj2.ParentRoots()[0], obj.ParentRoots()[0])

}
func TestRoot_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewRoot()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestRoot_Getters(t *testing.T) {
	obj := createMinimalSampleRoot()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleRoot(ctx, obj)

	assert.True(t, HasRoot(ctx, obj.PrimaryKey()))

	obj2 := LoadRoot(ctx, obj.PrimaryKey(), node.Root().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.Root().ID().Identifier))
	assert.Equal(t, obj.Name(), obj.Get(node.Root().Name().Identifier))
	assert.Panics(t, func() { obj2.Name() })
	assert.Nil(t, obj2.Get(node.Root().Name().Identifier))
	assert.Panics(t, func() { obj2.OptionalLeafID() })
	assert.Nil(t, obj2.Get(node.Root().OptionalLeafID().Identifier))
	assert.Panics(t, func() { obj2.RequiredLeafID() })
	assert.Nil(t, obj2.Get(node.Root().RequiredLeafID().Identifier))
	assert.Panics(t, func() { obj2.OptionalLeafUniqueID() })
	assert.Nil(t, obj2.Get(node.Root().OptionalLeafUniqueID().Identifier))
	assert.Panics(t, func() { obj2.RequiredLeafUniqueID() })
	assert.Nil(t, obj2.Get(node.Root().RequiredLeafUniqueID().Identifier))
	assert.Panics(t, func() { obj2.ParentID() })
	assert.Nil(t, obj2.Get(node.Root().ParentID().Identifier))
}

func TestRoot_QueryLoad(t *testing.T) {
	obj := createMinimalSampleRoot()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleRoot(ctx, obj)

	objs := QueryRoots(ctx).
		Where(op.Equal(node.Root().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.Root().PrimaryKey()). // exercise order by
		Limit(1, 0).                       // exercise limit
		Calculation(node.Root(), "IsTrue", op.Equal(1, 1)).
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestRoot_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleRoot()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleRoot(ctx, obj)

	objs := QueryRoots(ctx).
		Where(op.Equal(node.Root().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestRoot_QueryCursor(t *testing.T) {
	obj := createMinimalSampleRoot()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleRoot(ctx, obj)

	cursor := QueryRoots(ctx).
		Where(op.Equal(node.Root().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryRoots(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestRoot_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleRoot(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	// reread in case there are data limitations imposed by the database
	obj2 := LoadRoot(ctx, obj.PrimaryKey())
	defer deleteSampleRoot(ctx, obj)

	assert.Less(t, 0, CountRoots(ctx))

	assert.Less(t, 0, CountRootsByID(ctx, obj2.ID()))
	assert.Less(t, 0, CountRootsByName(ctx, obj2.Name()))
	assert.Less(t, 0, CountRootsByOptionalLeafID(ctx, obj2.OptionalLeafID()))
	assert.Less(t, 0, CountRootsByRequiredLeafID(ctx, obj2.RequiredLeafID()))
	assert.Less(t, 0, CountRootsByOptionalLeafUniqueID(ctx, obj2.OptionalLeafUniqueID()))
	assert.Less(t, 0, CountRootsByRequiredLeafUniqueID(ctx, obj2.RequiredLeafUniqueID()))
	assert.Less(t, 0, CountRootsByParentID(ctx, obj2.ParentID()))

}
func TestRoot_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleRoot()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewRoot()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsRoot(t, obj, obj2)
}

func TestRoot_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleRoot()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewRoot()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsRoot(t, obj, obj2)
}

func TestRoot_Indexes(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleRoot(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleRoot(ctx, obj)

	var obj2 *Root
	obj2 = LoadRootByOptionalLeafUniqueID(ctx, obj.OptionalLeafUniqueID())
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.True(t, HasRootByOptionalLeafUniqueID(ctx, obj.OptionalLeafUniqueID()))

	obj2 = LoadRootByRequiredLeafUniqueID(ctx, obj.RequiredLeafUniqueID())
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.True(t, HasRootByRequiredLeafUniqueID(ctx, obj.RequiredLeafUniqueID()))

}
