// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/goradd/orm/_test/gen/orm/goradd_unit/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/op"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// createMinimalSampleTypeTest creates an unsaved minimal version of a TypeTest object
// for testing.
func createMinimalSampleTypeTest() *TypeTest {
	obj := NewTypeTest()
	updateMinimalSampleTypeTest(obj)

	return obj
}

// updateMinimalSampleTypeTest sets the values of a minimal sample to new, random values.
func updateMinimalSampleTypeTest(obj *TypeTest) {

	obj.SetDate(test.RandomValue[time.Time](0))

	obj.SetTime(test.RandomValue[time.Time](0))

	obj.SetDateTime(test.RandomValue[time.Time](0))

	obj.SetTestInt(test.RandomValue[int](32))

	obj.SetTestFloat(test.RandomValue[float32](32))

	obj.SetTestDouble(test.RandomValue[float64](64))

	obj.SetTestText(test.RandomValue[string](65535))

	obj.SetTestBit(test.RandomValue[bool](0))

	obj.SetTestVarchar(test.RandomValue[string](10))

	obj.SetTestBlob(test.RandomValue[[]byte](65535))

}

// createMaximalSampleTypeTest creates an unsaved version of a TypeTest object
// for testing that includes references to minimal objects.
func createMaximalSampleTypeTest() *TypeTest {
	obj := NewTypeTest()
	updateMaximalSampleTypeTest(obj)
	return obj
}

// updateMaximalSampleTypeTest sets all the maximal sample values to new values.
// This will set new values for references, so save the old values and delete them.
func updateMaximalSampleTypeTest(obj *TypeTest) {
	updateMinimalSampleTypeTest(obj)

}

// deleteSampleTypeTest deletes an object created and saved by one of the sample creator functions.
func deleteSampleTypeTest(ctx context.Context, obj *TypeTest) {
	if obj == nil {
		return
	}

	obj.Delete(ctx)

}

// assertEqualFieldsTypeTest compares two objects and asserts that the basic fields are equal.
func assertEqualFieldsTypeTest(t *testing.T, obj1, obj2 *TypeTest) {
	if obj1.IDIsLoaded() && obj2.IDIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.ID(), obj2.ID())
	}
	if obj1.DateIsLoaded() && obj2.DateIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Date(), obj2.Date())
	}
	if obj1.TimeIsLoaded() && obj2.TimeIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Time(), obj2.Time())
	}
	if obj1.DateTimeIsLoaded() && obj2.DateTimeIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.DateTime(), obj2.DateTime())
	}
	if obj1.TsIsLoaded() && obj2.TsIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.Ts(), obj2.Ts())
	}
	if obj1.TestIntIsLoaded() && obj2.TestIntIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.TestInt(), obj2.TestInt())
	}
	if obj1.TestFloatIsLoaded() && obj2.TestFloatIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.TestFloat(), obj2.TestFloat())
	}
	if obj1.TestDoubleIsLoaded() && obj2.TestDoubleIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.TestDouble(), obj2.TestDouble())
	}
	if obj1.TestTextIsLoaded() && obj2.TestTextIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.TestText(), obj2.TestText())
	}
	if obj1.TestBitIsLoaded() && obj2.TestBitIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.TestBit(), obj2.TestBit())
	}
	if obj1.TestVarcharIsLoaded() && obj2.TestVarcharIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.TestVarchar(), obj2.TestVarchar())
	}
	if obj1.TestBlobIsLoaded() && obj2.TestBlobIsLoaded() { // only check loaded values
		assert.EqualValues(t, obj1.TestBlob(), obj2.TestBlob())
	}

}

func TestTypeTest_SetID(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomNumberString()
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID("")
	assert.EqualValues(t, "", obj.ID(), "set default")

}
func TestTypeTest_SetDate(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[time.Time](0)
	obj.SetDate(val)
	val = obj.Date()
	assert.Zero(t, val.Minute()) // make sure minute part is zero'd
	assert.Zero(t, val.Hour())   // make sure hour part is zero'd
	assert.Zero(t, val.Second()) // make sure second part is zero'd
	assert.False(t, obj.DateIsNull())

	// Test NULL
	obj.SetDateToNull()
	assert.EqualValues(t, time.Time{}, obj.Date())
	assert.True(t, obj.DateIsNull())

	// test default
	obj.SetDate(time.Time{})
	assert.EqualValues(t, time.Time{}, obj.Date(), "set default")

}
func TestTypeTest_SetTime(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[time.Time](0)
	obj.SetTime(val)
	val = obj.Time()
	assert.EqualValues(t, 1, val.Year())  // make sure year part is zero'd. The zero value of time.Time has a year of 1.
	assert.EqualValues(t, 1, val.Month()) // make sure month part is zero'd
	assert.EqualValues(t, 1, val.Day())   // make sure day part is zero'd
	assert.False(t, obj.TimeIsNull())

	// Test NULL
	obj.SetTimeToNull()
	assert.EqualValues(t, time.Time{}, obj.Time())
	assert.True(t, obj.TimeIsNull())

	// test default
	obj.SetTime(time.Time{})
	assert.EqualValues(t, time.Time{}, obj.Time(), "set default")

}
func TestTypeTest_SetDateTime(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[time.Time](0)
	obj.SetDateTime(val)
	assert.Equal(t, val, obj.DateTime())
	assert.False(t, obj.DateTimeIsNull())

	// Test NULL
	obj.SetDateTimeToNull()
	assert.EqualValues(t, time.Time{}, obj.DateTime())
	assert.True(t, obj.DateTimeIsNull())

	// test default
	obj.SetDateTime(time.Time{})
	assert.EqualValues(t, time.Time{}, obj.DateTime(), "set default")

}
func TestTypeTest_SetTestInt(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[int](32)
	obj.SetTestInt(val)
	assert.Equal(t, val, obj.TestInt())
	assert.False(t, obj.TestIntIsNull())

	// Test NULL
	obj.SetTestIntToNull()
	assert.EqualValues(t, 5, obj.TestInt())
	assert.True(t, obj.TestIntIsNull())

	// test default
	obj.SetTestInt(5)
	assert.EqualValues(t, 5, obj.TestInt(), "set default")

}
func TestTypeTest_SetTestFloat(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[float32](32)
	obj.SetTestFloat(val)
	assert.Equal(t, val, obj.TestFloat())
	assert.False(t, obj.TestFloatIsNull())

	// Test NULL
	obj.SetTestFloatToNull()
	assert.EqualValues(t, 0, obj.TestFloat())
	assert.True(t, obj.TestFloatIsNull())

	// test default
	obj.SetTestFloat(0)
	assert.EqualValues(t, 0, obj.TestFloat(), "set default")

}
func TestTypeTest_SetTestDouble(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[float64](64)
	obj.SetTestDouble(val)
	assert.Equal(t, val, obj.TestDouble())

	// test default
	obj.SetTestDouble(0)
	assert.EqualValues(t, 0, obj.TestDouble(), "set default")

}
func TestTypeTest_SetTestText(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](65535)
	obj.SetTestText(val)
	assert.Equal(t, val, obj.TestText())
	assert.False(t, obj.TestTextIsNull())

	// Test NULL
	obj.SetTestTextToNull()
	assert.EqualValues(t, "", obj.TestText())
	assert.True(t, obj.TestTextIsNull())

	// test default
	obj.SetTestText("")
	assert.EqualValues(t, "", obj.TestText(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](65536)
	assert.Panics(t, func() {
		obj.SetTestText(val)
	})
}
func TestTypeTest_SetTestBit(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[bool](0)
	obj.SetTestBit(val)
	assert.Equal(t, val, obj.TestBit())
	assert.False(t, obj.TestBitIsNull())

	// Test NULL
	obj.SetTestBitToNull()
	assert.EqualValues(t, false, obj.TestBit())
	assert.True(t, obj.TestBitIsNull())

	// test default
	obj.SetTestBit(false)
	assert.EqualValues(t, false, obj.TestBit(), "set default")

}
func TestTypeTest_SetTestVarchar(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[string](10)
	obj.SetTestVarchar(val)
	assert.Equal(t, val, obj.TestVarchar())
	assert.False(t, obj.TestVarcharIsNull())

	// Test NULL
	obj.SetTestVarcharToNull()
	assert.EqualValues(t, "", obj.TestVarchar())
	assert.True(t, obj.TestVarcharIsNull())

	// test default
	obj.SetTestVarchar("")
	assert.EqualValues(t, "", obj.TestVarchar(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](11)
	assert.Panics(t, func() {
		obj.SetTestVarchar(val)
	})
}
func TestTypeTest_SetTestBlob(t *testing.T) {

	obj := NewTypeTest()

	assert.True(t, obj.IsNew())
	val := test.RandomValue[[]byte](65535)
	obj.SetTestBlob(val)
	assert.Equal(t, val, obj.TestBlob())

	// test default
	obj.SetTestBlob([]byte{})
	assert.EqualValues(t, []byte{}, obj.TestBlob(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[[]byte](65536)
	assert.Panics(t, func() {
		obj.SetTestBlob(val)
	})
}

func TestTypeTest_Copy(t *testing.T) {
	obj := createMinimalSampleTypeTest()

	obj2 := obj.Copy()

	assert.Equal(t, obj.Date(), obj2.Date())
	assert.Equal(t, obj.Time(), obj2.Time())
	assert.Equal(t, obj.DateTime(), obj2.DateTime())
	assert.Equal(t, obj.TestInt(), obj2.TestInt())
	assert.Equal(t, obj.TestFloat(), obj2.TestFloat())
	assert.Equal(t, obj.TestDouble(), obj2.TestDouble())
	assert.Equal(t, obj.TestText(), obj2.TestText())
	assert.Equal(t, obj.TestBit(), obj2.TestBit())
	assert.Equal(t, obj.TestVarchar(), obj2.TestVarchar())
	assert.Equal(t, obj.TestBlob(), obj2.TestBlob())

}

func TestTypeTest_BasicInsert(t *testing.T) {
	obj := createMinimalSampleTypeTest()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleTypeTest(ctx, obj)

	// Test retrieval
	obj2 := LoadTypeTest(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsLoaded())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.idIsDirty)
	obj2.SetID(obj2.ID())
	assert.False(t, obj2.idIsDirty)

	assert.True(t, obj2.DateIsLoaded())
	assert.False(t, obj2.DateIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.dateIsDirty)
	obj2.SetDate(obj2.Date())
	assert.False(t, obj2.dateIsDirty)

	assert.True(t, obj2.TimeIsLoaded())
	assert.False(t, obj2.TimeIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.timeIsDirty)
	obj2.SetTime(obj2.Time())
	assert.False(t, obj2.timeIsDirty)

	assert.True(t, obj2.DateTimeIsLoaded())
	assert.False(t, obj2.DateTimeIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.dateTimeIsDirty)
	obj2.SetDateTime(obj2.DateTime())
	assert.False(t, obj2.dateTimeIsDirty)

	assert.True(t, obj2.TsIsLoaded())
	assert.False(t, obj2.TsIsNull())

	assert.True(t, obj2.TestIntIsLoaded())
	assert.False(t, obj2.TestIntIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testIntIsDirty)
	obj2.SetTestInt(obj2.TestInt())
	assert.False(t, obj2.testIntIsDirty)

	assert.True(t, obj2.TestFloatIsLoaded())
	assert.False(t, obj2.TestFloatIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testFloatIsDirty)
	obj2.SetTestFloat(obj2.TestFloat())
	assert.False(t, obj2.testFloatIsDirty)

	assert.True(t, obj2.TestDoubleIsLoaded())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testDoubleIsDirty)
	obj2.SetTestDouble(obj2.TestDouble())
	assert.False(t, obj2.testDoubleIsDirty)

	assert.True(t, obj2.TestTextIsLoaded())
	assert.False(t, obj2.TestTextIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testTextIsDirty)
	obj2.SetTestText(obj2.TestText())
	assert.False(t, obj2.testTextIsDirty)

	assert.True(t, obj2.TestBitIsLoaded())
	assert.False(t, obj2.TestBitIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testBitIsDirty)
	obj2.SetTestBit(obj2.TestBit())
	assert.False(t, obj2.testBitIsDirty)

	assert.True(t, obj2.TestVarcharIsLoaded())
	assert.False(t, obj2.TestVarcharIsNull())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testVarcharIsDirty)
	obj2.SetTestVarchar(obj2.TestVarchar())
	assert.False(t, obj2.testVarcharIsDirty)

	assert.True(t, obj2.TestBlobIsLoaded())

	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testBlobIsDirty)
	obj2.SetTestBlob(obj2.TestBlob())
	assert.False(t, obj2.testBlobIsDirty)

}

func TestTypeTest_InsertPanics(t *testing.T) {
	obj := createMinimalSampleTypeTest()
	ctx := db.NewContext(nil)

	obj.testDoubleIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.testDoubleIsLoaded = true

	obj.testBlobIsLoaded = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.testBlobIsLoaded = true

}

func TestTypeTest_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleTypeTest()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleTypeTest(ctx, obj)
	updateMinimalSampleTypeTest(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadTypeTest(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID(), "ID did not update")

	assert.WithinDuration(t, obj2.Date(), obj.Date(), time.Second, "Date not within one second")

	assert.WithinDuration(t, obj2.Time(), obj.Time(), time.Second, "Time not within one second")

	assert.WithinDuration(t, obj2.DateTime(), obj.DateTime(), time.Second, "DateTime not within one second")

	assert.WithinDuration(t, obj2.Ts(), obj.Ts(), time.Second, "Ts not within one second")
	assert.Equal(t, obj2.TestInt(), obj.TestInt(), "TestInt did not update")
	assert.Equal(t, obj2.TestFloat(), obj.TestFloat(), "TestFloat did not update")
	assert.Equal(t, obj2.TestDouble(), obj.TestDouble(), "TestDouble did not update")
	assert.Equal(t, obj2.TestText(), obj.TestText(), "TestText did not update")
	assert.Equal(t, obj2.TestBit(), obj.TestBit(), "TestBit did not update")
	assert.Equal(t, obj2.TestVarchar(), obj.TestVarchar(), "TestVarchar did not update")
	assert.Equal(t, obj2.TestBlob(), obj.TestBlob(), "TestBlob did not update")
}

func TestTypeTest_ReferenceLoad(t *testing.T) {
	obj := createMaximalSampleTypeTest()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleTypeTest(ctx, obj)

	// Test that referenced objects were saved and assigned ids

	// Test lazy loading
	obj2 := LoadTypeTest(ctx, obj.PrimaryKey())
	objPkOnly := LoadTypeTest(ctx, obj.PrimaryKey(), node.TypeTest().PrimaryKey())
	_ = obj2 // avoid error if there are no references
	_ = objPkOnly

	// test eager loading
	obj3 := LoadTypeTest(ctx, obj.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}

func TestTypeTest_ReferenceUpdateNewObjects(t *testing.T) {
	obj := createMaximalSampleTypeTest()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleTypeTest(ctx, obj)

	obj2 := LoadTypeTest(ctx, obj.PrimaryKey())
	updateMaximalSampleTypeTest(obj2)
	assert.NoError(t, obj2.Save(ctx))
	defer deleteSampleTypeTest(ctx, obj2)

	obj3 := LoadTypeTest(ctx, obj2.PrimaryKey())
	_ = obj3 // avoid error if there are no references

}

func TestTypeTest_ReferenceUpdateOldObjects(t *testing.T) {
	obj := createMaximalSampleTypeTest()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleTypeTest(ctx, obj)

	assert.NoError(t, obj.Save(ctx))

	obj2 := LoadTypeTest(ctx, obj.PrimaryKey())
	_ = obj2 // avoid error if there are no references

}
func TestTypeTest_EmptyPrimaryKeyGetter(t *testing.T) {
	obj := NewTypeTest()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)
}

func TestTypeTest_Getters(t *testing.T) {
	obj := createMinimalSampleTypeTest()

	i, err := strconv.Atoi(obj.ID())
	assert.NoError(t, err)
	assert.True(t, i < 0)

	ctx := db.NewContext(nil)
	require.NoError(t, obj.Save(ctx))
	defer deleteSampleTypeTest(ctx, obj)

	assert.True(t, HasTypeTest(ctx, obj.PrimaryKey()))

	obj2 := LoadTypeTest(ctx, obj.PrimaryKey(), node.TypeTest().PrimaryKey())

	assert.Equal(t, obj.ID(), obj.Get(node.TypeTest().ID().Identifier))
	assert.Equal(t, obj.Date(), obj.Get(node.TypeTest().Date().Identifier))
	assert.Panics(t, func() { obj2.Date() })
	assert.Nil(t, obj2.Get(node.TypeTest().Date().Identifier))
	assert.Equal(t, obj.Time(), obj.Get(node.TypeTest().Time().Identifier))
	assert.Panics(t, func() { obj2.Time() })
	assert.Nil(t, obj2.Get(node.TypeTest().Time().Identifier))
	assert.Equal(t, obj.DateTime(), obj.Get(node.TypeTest().DateTime().Identifier))
	assert.Panics(t, func() { obj2.DateTime() })
	assert.Nil(t, obj2.Get(node.TypeTest().DateTime().Identifier))
	assert.Equal(t, obj.Ts(), obj.Get(node.TypeTest().Ts().Identifier))
	assert.Panics(t, func() { obj2.Ts() })
	assert.Nil(t, obj2.Get(node.TypeTest().Ts().Identifier))
	assert.Equal(t, obj.TestInt(), obj.Get(node.TypeTest().TestInt().Identifier))
	assert.Panics(t, func() { obj2.TestInt() })
	assert.Nil(t, obj2.Get(node.TypeTest().TestInt().Identifier))
	assert.Equal(t, obj.TestFloat(), obj.Get(node.TypeTest().TestFloat().Identifier))
	assert.Panics(t, func() { obj2.TestFloat() })
	assert.Nil(t, obj2.Get(node.TypeTest().TestFloat().Identifier))
	assert.Equal(t, obj.TestDouble(), obj.Get(node.TypeTest().TestDouble().Identifier))
	assert.Panics(t, func() { obj2.TestDouble() })
	assert.Nil(t, obj2.Get(node.TypeTest().TestDouble().Identifier))
	assert.Equal(t, obj.TestText(), obj.Get(node.TypeTest().TestText().Identifier))
	assert.Panics(t, func() { obj2.TestText() })
	assert.Nil(t, obj2.Get(node.TypeTest().TestText().Identifier))
	assert.Equal(t, obj.TestBit(), obj.Get(node.TypeTest().TestBit().Identifier))
	assert.Panics(t, func() { obj2.TestBit() })
	assert.Nil(t, obj2.Get(node.TypeTest().TestBit().Identifier))
	assert.Equal(t, obj.TestVarchar(), obj.Get(node.TypeTest().TestVarchar().Identifier))
	assert.Panics(t, func() { obj2.TestVarchar() })
	assert.Nil(t, obj2.Get(node.TypeTest().TestVarchar().Identifier))
	assert.Equal(t, obj.TestBlob(), obj.Get(node.TypeTest().TestBlob().Identifier))
	assert.Panics(t, func() { obj2.TestBlob() })
	assert.Nil(t, obj2.Get(node.TypeTest().TestBlob().Identifier))
}

func TestTypeTest_QueryLoad(t *testing.T) {
	obj := createMinimalSampleTypeTest()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleTypeTest(ctx, obj)

	objs := QueryTypeTests(ctx).
		Where(op.Equal(node.TypeTest().PrimaryKey(), obj.PrimaryKey())).
		OrderBy(node.TypeTest().PrimaryKey()). // exercise order by
		Limit(1, 0).                           // exercise limit
		Load()

	assert.Equal(t, obj.PrimaryKey(), objs[0].PrimaryKey())
}
func TestTypeTest_QueryLoadI(t *testing.T) {
	obj := createMinimalSampleTypeTest()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleTypeTest(ctx, obj)

	objs := QueryTypeTests(ctx).
		Where(op.Equal(node.TypeTest().PrimaryKey(), obj.PrimaryKey())).
		LoadI()

	assert.Equal(t, obj.PrimaryKey(), objs[0].Get("ID"))
}
func TestTypeTest_QueryCursor(t *testing.T) {
	obj := createMinimalSampleTypeTest()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleTypeTest(ctx, obj)

	cursor := QueryTypeTests(ctx).
		Where(op.Equal(node.TypeTest().PrimaryKey(), obj.PrimaryKey())).
		LoadCursor()

	obj2 := cursor.Next()
	assert.Equal(t, obj.PrimaryKey(), obj2.PrimaryKey())
	assert.Nil(t, cursor.Next())

	// test empty cursor result
	cursor = QueryTypeTests(ctx).
		Where(op.Equal(1, 0)).
		LoadCursor()
	assert.Nil(t, cursor.Next())

}
func TestTypeTest_Count(t *testing.T) {
	obj := createMaximalSampleTypeTest()
	ctx := db.NewContext(nil)
	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer deleteSampleTypeTest(ctx, obj)

	assert.Less(t, 0, CountTypeTests(ctx))

	assert.Less(t, 0, CountTypeTestsByID(ctx, obj.ID()))
	assert.Less(t, 0, CountTypeTestsByDate(ctx, obj.Date()))
	assert.Less(t, 0, CountTypeTestsByTime(ctx, obj.Time()))
	assert.Less(t, 0, CountTypeTestsByDateTime(ctx, obj.DateTime()))
	assert.Less(t, 0, CountTypeTestsByTs(ctx, obj.Ts()))
	assert.Less(t, 0, CountTypeTestsByTestInt(ctx, obj.TestInt()))
	assert.Less(t, 0, CountTypeTestsByTestFloat(ctx, obj.TestFloat()))
	assert.Less(t, 0, CountTypeTestsByTestDouble(ctx, obj.TestDouble()))
	assert.Less(t, 0, CountTypeTestsByTestText(ctx, obj.TestText()))
	assert.Less(t, 0, CountTypeTestsByTestBit(ctx, obj.TestBit()))
	assert.Less(t, 0, CountTypeTestsByTestVarchar(ctx, obj.TestVarchar()))
	assert.Less(t, 0, CountTypeTestsByTestBlob(ctx, obj.TestBlob()))

}
