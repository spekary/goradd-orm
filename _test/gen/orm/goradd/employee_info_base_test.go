// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/goradd/orm/_test/gen/orm/goradd/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createMinimalSampleEmployeeInfo creates an unsaved minimal version of a EmployeeInfo object
// for testing.
func createMinimalSampleEmployeeInfo() *EmployeeInfo {
	obj := NewEmployeeInfo()
	updateMinimalSampleEmployeeInfo(obj)

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	// If the database is configured so that the referenced object points back here, either directly or through multiple
	// forward references, it possible this could create an endless loop. Such a structure could not be saved anyways.
	obj.SetPerson(createMinimalSamplePerson())

	return obj
}

// updateMinimalSampleEmployeeInfo sets the values of a minimal sample to new, random values.
func updateMinimalSampleEmployeeInfo(obj *EmployeeInfo) {

	obj.SetEmployeeNumber(test.RandomValue[int](32))

}

// createMaximalSampleEmployeeInfo creates an unsaved version of a EmployeeInfo object
// for testing that includes references to minimal objects.
func createMaximalSampleEmployeeInfo(ctx context.Context) *EmployeeInfo {
	obj := NewEmployeeInfo()
	updateMaximalSampleEmployeeInfo(ctx, obj)
	return obj
}

// updateMaximalSampleEmployeeInfo sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleEmployeeInfo(ctx context.Context, obj *EmployeeInfo) {
	updateMinimalSampleEmployeeInfo(obj)
	obj.SetPerson(createMinimalSamplePerson())

}

// deleteSampleEmployeeInfo deletes an object created and saved by one of the sample creator functions.
func deleteSampleEmployeeInfo(ctx context.Context, obj *EmployeeInfo) {
	if obj == nil {
		return
	}

	_ = obj.Delete(ctx)

	deleteSamplePerson(ctx, obj.Person())

}

// assertEqualFieldsEmployeeInfo compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsEmployeeInfo(t *testing.T, obj1, obj2 *EmployeeInfo) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.PersonIDIsLoaded() && obj2.PersonIDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.PersonID(), obj2.PersonID())
	}
	if obj1.EmployeeNumberIsLoaded() && obj2.EmployeeNumberIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.EmployeeNumber(), obj2.EmployeeNumber())
	}

}

func TestEmployeeInfo_SetID(t *testing.T) {

	obj := NewEmployeeInfo()

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
func TestEmployeeInfo_SetPersonID(t *testing.T) {

	obj := NewEmployeeInfo()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](0)
	obj.SetPersonID(val)
	assert.Equal(t, val, obj.PersonID())

	// test default
	obj.SetPersonID("")
	assert.EqualValues(t, "", obj.PersonID(), "set default")

}
func TestEmployeeInfo_SetEmployeeNumber(t *testing.T) {

	obj := NewEmployeeInfo()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[int](32)
	obj.SetEmployeeNumber(val)
	assert.Equal(t, val, obj.EmployeeNumber())

	// test default
	obj.SetEmployeeNumber(0)
	assert.EqualValues(t, 0, obj.EmployeeNumber(), "set default")

}

func TestEmployeeInfo_Copy(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()

	obj2 := obj.Copy()

	assert.Equal(t, obj.PersonID(), obj2.PersonID())
	assert.Equal(t, obj.EmployeeNumber(), obj2.EmployeeNumber())

}

func TestEmployeeInfo_BasicInsert(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleEmployeeInfo(ctx, obj)

	// Test retrieval
	obj2, err := LoadEmployeeInfo(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)
	assert.NoError(t, err)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsLoaded())
	assert.Panics(t, func() {
		obj2.SetID(obj2.ID())
	})

	assert.True(t, obj2.EmployeeNumberIsLoaded())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.employeeNumberIsDirty)
	obj2.SetEmployeeNumber(obj2.EmployeeNumber())
	assert.False(t, obj2.employeeNumberIsDirty)

}

func TestEmployeeInfo_InsertPanics(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()
	_ = obj
	ctx := db.NewContext(nil)
	_ = ctx

	obj.objPerson = nil
	obj.personIDIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.personIDIsLoaded = true

	obj.employeeNumberIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.employeeNumberIsLoaded = true

}

func TestEmployeeInfo_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleEmployeeInfo(ctx, obj)
	updateMinimalSampleEmployeeInfo(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2, err := LoadEmployeeInfo(ctx, obj.PrimaryKey())
	assert.NoError(t, err)

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")
	assert.Equal(t, obj2.EmployeeNumber(), obj.EmployeeNumber(), "EmployeeNumber did not update")
}

func TestEmployeeInfo_ReferenceLoad(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleEmployeeInfo(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleEmployeeInfo(ctx, obj)

	// Test that referenced objects were saved and assigned ids
	assert.NotNil(t, obj.Person())
	assert.NotEqual(t, '-', obj.Person().PrimaryKey()[0])

	// Test lazy loading
	obj2, err := LoadEmployeeInfo(ctx, obj.PrimaryKey())
	assert.NoError(t, err)
	objPkOnly, err2 := LoadEmployeeInfo(ctx, obj.PrimaryKey(), node.EmployeeInfo().PrimaryKey())
	assert.NoError(t, err2)
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	assert.Nil(t, obj2.Person(), "Person is not loaded initially")
	v_Person, _ := obj2.LoadPerson(ctx)
	assert.NotNil(t, v_Person)
	assert.Equal(t, v_Person.PrimaryKey(), obj2.Person().PrimaryKey())
	assert.Equal(t, obj.Person().PrimaryKey(), obj2.Person().PrimaryKey())
	assert.True(t, obj2.PersonIDIsLoaded())

	assert.False(t, objPkOnly.PersonIDIsLoaded())
	assert.Panics(t, func() { _, _ = objPkOnly.LoadPerson(ctx) })

	assert.Panics(t, func() {
		objPkOnly.SetPerson(nil)
	})

	// test eager loading
	obj3, _ := LoadEmployeeInfo(ctx, obj.PrimaryKey(), node.EmployeeInfo().Person())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Person().PrimaryKey(), obj3.Person().PrimaryKey())

}

func TestEmployeeInfo_ReferenceUpdateNewObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleEmployeeInfo(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleEmployeeInfo(ctx, obj)

	obj2, _ := LoadEmployeeInfo(ctx, obj.PrimaryKey())
	updateMaximalSampleEmployeeInfo(ctx, obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleEmployeeInfo(ctx, obj2)

	obj3, _ := LoadEmployeeInfo(ctx, obj2.PrimaryKey(), node.EmployeeInfo().Person())
	_ = obj3 // avoid error if there are no references

	assert.Equal(t, obj2.Person().PrimaryKey(), obj3.Person().PrimaryKey())

}

func TestEmployeeInfo_ReferenceUpdateOldObjects(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleEmployeeInfo(ctx)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleEmployeeInfo(ctx, obj)

	updateMinimalSamplePerson(obj.Person())

	assert.NoError(t, obj.Save(ctx))

	obj2, _ := LoadEmployeeInfo(ctx, obj.PrimaryKey(),
		node.EmployeeInfo().Person(),
	)
	_ = obj2 // avoid error if there are no references

	assertEqualFieldsPerson(t, obj2.Person(), obj.Person())

}
func TestEmployeeInfo_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewEmployeeInfo()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestEmployeeInfo_Getters(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleEmployeeInfo(ctx, obj)

	has, _ := HasEmployeeInfo(ctx, obj.PrimaryKey())
	assert.True(t, has)

	obj2, _ := LoadEmployeeInfo(ctx, obj.PrimaryKey(), node.EmployeeInfo().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.EmployeeInfo().ID().Identifier))
	assert.Panics(t, func() { obj2.PersonID() })
	assert.Nil(t, obj2.Get(node.EmployeeInfo().PersonID().Identifier))
	assert.Equal(t, obj.EmployeeNumber(), obj.Get(node.EmployeeInfo().EmployeeNumber().Identifier))
	assert.Panics(t, func() { obj2.EmployeeNumber() })
	assert.Nil(t, obj2.Get(node.EmployeeInfo().EmployeeNumber().Identifier))
}

func TestEmployeeInfo_QueryLoad(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleEmployeeInfo(ctx, obj)

	objs, err := QueryEmployeeInfos(ctx).
		Where(op.Equal(node.EmployeeInfo().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.EmployeeInfo().PrimaryKey()). // exercise order by
		Limit(1, 0).                               // exercise limit
		Calculation(node.EmployeeInfo(), "IsTrue", op.Equal("A", "A")).
		Load()
	assert.NoError(t, err)
	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
	assert.True(t, objs[0].GetAlias("IsTrue").Bool())
}
func TestEmployeeInfo_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleEmployeeInfo(ctx, obj)

	objs, _ := QueryEmployeeInfos(ctx).
		Where(op.Equal(node.EmployeeInfo().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestEmployeeInfo_QueryCursor(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleEmployeeInfo(ctx, obj)

	cursor, err := QueryEmployeeInfos(ctx).
		Where(op.Equal(node.EmployeeInfo().PrimaryKey(), obj.PrimaryKey())).
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
	cursor, err = QueryEmployeeInfos(ctx).
		Where(op.Equal("B", "A")).
		LoadCursor()
	require.NoError(t, err)

	obj2, err = cursor.Next()
	assert.Nil(t, obj2)
	require.NoError(t, err)
	assert.NoError(t, cursor.Close())
}
func TestEmployeeInfo_Count(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleEmployeeInfo(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleEmployeeInfo(ctx, obj)
	assert.Positive(t, func() int { i, _ := CountEmployeeInfos(ctx); return i }())

	// reread in case there are data limitations imposed by the database
	obj2, _ := LoadEmployeeInfo(ctx, obj.PrimaryKey())
	assert.Positive(t,
		func() int {
			i, _ := CountEmployeeInfosByPersonID(ctx,
				obj2.PersonID())
			return i
		}())

}

func TestEmployeeInfo_MarshalJSON(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()

	b, err := json.Marshal(obj)
	assert.NoError(t, err)

	obj2 := NewEmployeeInfo()
	err = json.Unmarshal(b, &obj2)
	assert.NoError(t, err)

	assertEqualFieldsEmployeeInfo(t, obj, obj2)
}

func TestEmployeeInfo_MarshalBinary(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()

	b, err := obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 := NewEmployeeInfo()
	err = obj2.UnmarshalBinary(b)
	assert.NoError(t, err)

	assertEqualFieldsEmployeeInfo(t, obj, obj2)
}

func TestEmployeeInfo_FailingMarshalBinary(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()
	var err error

	for i := 0; i < 12; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
	// do it again with aliases
	obj._aliases = make(map[string]any)
	for i := 0; i < 13; i++ {
		enc := &test.GobEncoder{Count: i}
		err = obj.encodeTo(enc)
		assert.Error(t, err)
	}
}

func TestEmployeeInfo_FailingUnmarshalBinary(t *testing.T) {
	obj := createMinimalSampleEmployeeInfo()
	b, err := obj.MarshalBinary()
	assert.NoError(t, err)
	obj2 := NewEmployeeInfo()
	for i := 0; i < 12; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}

	// do it again with aliases
	obj = createMinimalSampleEmployeeInfo()
	obj._aliases = map[string]any{"a": 1}
	b, err = obj.MarshalBinary()
	assert.NoError(t, err)

	obj2 = NewEmployeeInfo()
	for i := 0; i < 13; i++ {
		buf := bytes.NewReader(b)
		dec := &test.GobDecoder{Decoder: gob.NewDecoder(buf), Count: i}
		err = obj2.decodeFrom(dec)
		assert.Error(t, err)
	}
}

func TestEmployeeInfo_Indexes(t *testing.T) {
	ctx := db.NewContext(nil)
	obj := createMaximalSampleEmployeeInfo(ctx)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleEmployeeInfo(ctx, obj)

	var obj2 *EmployeeInfo
	obj2, _ = LoadEmployeeInfoByPersonID(ctx, obj.PersonID())
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.True(t, func() bool { h, _ := HasEmployeeInfoByPersonID(ctx, obj.PersonID()); return h }())

}
