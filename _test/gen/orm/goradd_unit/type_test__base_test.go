// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/goradd/orm/_test/gen/orm/goradd_unit/node"
	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeTest_SetDate(t *testing.T) {

	obj := NewTypeTest()
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
	val := test.RandomValue[float64](64)
	obj.SetTestDouble(val)
	assert.Equal(t, val, obj.TestDouble())

	// test default
	obj.SetTestDouble(0)
	assert.EqualValues(t, 0, obj.TestDouble(), "set default")

}
func TestTypeTest_SetTestText(t *testing.T) {

	obj := NewTypeTest()
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

func TestTypeTest_BasicInsert(t *testing.T) {
	obj := NewTypeTest()
	ctx := db.NewContext(nil)

	v_date := test.RandomValue[time.Time](0)
	obj.SetDate(v_date)

	v_date = obj.Date()

	v_time := test.RandomValue[time.Time](0)
	obj.SetTime(v_time)

	v_time = obj.Time()

	v_dateTime := test.RandomValue[time.Time](0)
	obj.SetDateTime(v_dateTime)

	v_dateTime = obj.DateTime()

	v_testInt := test.RandomValue[int](32)
	obj.SetTestInt(v_testInt)

	v_testFloat := test.RandomValue[float32](32)
	obj.SetTestFloat(v_testFloat)

	v_testDouble := test.RandomValue[float64](64)
	obj.SetTestDouble(v_testDouble)

	v_testText := test.RandomValue[string](65535)
	obj.SetTestText(v_testText)

	v_testBit := test.RandomValue[bool](0)
	obj.SetTestBit(v_testBit)

	v_testVarchar := test.RandomValue[string](10)
	obj.SetTestVarchar(v_testVarchar)

	v_testBlob := test.RandomValue[[]byte](65535)
	obj.SetTestBlob(v_testBlob)

	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer obj.Delete(ctx)

	// Test retrieval
	obj2 := LoadTypeTest(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.Equal(t, obj2.PrimaryKey(), obj2.OriginalPrimaryKey())

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.DateIsValid())
	assert.False(t, obj2.DateIsNull())
	assert.EqualValues(t, v_date, obj2.Date())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.dateIsDirty)
	obj2.SetDate(obj2.Date())
	assert.False(t, obj2.dateIsDirty)

	assert.True(t, obj2.TimeIsValid())
	assert.False(t, obj2.TimeIsNull())
	assert.EqualValues(t, v_time, obj2.Time())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.timeIsDirty)
	obj2.SetTime(obj2.Time())
	assert.False(t, obj2.timeIsDirty)

	assert.True(t, obj2.DateTimeIsValid())
	assert.False(t, obj2.DateTimeIsNull())
	assert.EqualValues(t, v_dateTime, obj2.DateTime())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.dateTimeIsDirty)
	obj2.SetDateTime(obj2.DateTime())
	assert.False(t, obj2.dateTimeIsDirty)

	assert.True(t, obj2.TsIsValid())
	assert.False(t, obj2.TsIsNull())

	assert.True(t, obj2.TestIntIsValid())
	assert.False(t, obj2.TestIntIsNull())
	assert.EqualValues(t, v_testInt, obj2.TestInt())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testIntIsDirty)
	obj2.SetTestInt(obj2.TestInt())
	assert.False(t, obj2.testIntIsDirty)

	assert.True(t, obj2.TestFloatIsValid())
	assert.False(t, obj2.TestFloatIsNull())
	assert.EqualValues(t, v_testFloat, obj2.TestFloat())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testFloatIsDirty)
	obj2.SetTestFloat(obj2.TestFloat())
	assert.False(t, obj2.testFloatIsDirty)

	assert.True(t, obj2.TestDoubleIsValid())
	assert.EqualValues(t, v_testDouble, obj2.TestDouble())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testDoubleIsDirty)
	obj2.SetTestDouble(obj2.TestDouble())
	assert.False(t, obj2.testDoubleIsDirty)

	assert.True(t, obj2.TestTextIsValid())
	assert.False(t, obj2.TestTextIsNull())
	assert.EqualValues(t, v_testText, obj2.TestText())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testTextIsDirty)
	obj2.SetTestText(obj2.TestText())
	assert.False(t, obj2.testTextIsDirty)

	assert.True(t, obj2.TestBitIsValid())
	assert.False(t, obj2.TestBitIsNull())
	assert.EqualValues(t, v_testBit, obj2.TestBit())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testBitIsDirty)
	obj2.SetTestBit(obj2.TestBit())
	assert.False(t, obj2.testBitIsDirty)

	assert.True(t, obj2.TestVarcharIsValid())
	assert.False(t, obj2.TestVarcharIsNull())
	assert.EqualValues(t, v_testVarchar, obj2.TestVarchar())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testVarcharIsDirty)
	obj2.SetTestVarchar(obj2.TestVarchar())
	assert.False(t, obj2.testVarcharIsDirty)

	assert.True(t, obj2.TestBlobIsValid())
	assert.EqualValues(t, v_testBlob, obj2.TestBlob())
	// test that setting it to the same value will not change the dirty bit
	assert.False(t, obj2.testBlobIsDirty)
	obj2.SetTestBlob(obj2.TestBlob())
	assert.False(t, obj2.testBlobIsDirty)

}

func TestTypeTest_InsertPanics(t *testing.T) {
	obj := createMinimalSampleTypeTest()
	ctx := db.NewContext(nil)

	obj.testDoubleIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.testDoubleIsValid = true

	obj.testBlobIsValid = false
	assert.Panics(t, func() { obj.Save(ctx) })
	obj.testBlobIsValid = true

}

func TestTypeTest_BasicUpdate(t *testing.T) {
	obj := createMinimalSampleTypeTest()
	ctx := db.NewContext(nil)
	assert.NoError(t, obj.Save(ctx))
	defer deleteSampleTypeTest(ctx, obj)
	updateMinimalSampleTypeTest(obj)
	assert.NoError(t, obj.Save(ctx))
	obj2 := LoadTypeTest(ctx, obj.PrimaryKey())

	assert.Equal(t, obj2.ID(), obj.ID())
	assert.Equal(t, obj2.Date(), obj.Date())
	assert.Equal(t, obj2.Time(), obj.Time())
	assert.Equal(t, obj2.DateTime(), obj.DateTime())
	assert.Equal(t, obj2.Ts(), obj.Ts())
	assert.Equal(t, obj2.TestInt(), obj.TestInt())
	assert.Equal(t, obj2.TestFloat(), obj.TestFloat())
	assert.Equal(t, obj2.TestDouble(), obj.TestDouble())
	assert.Equal(t, obj2.TestText(), obj.TestText())
	assert.Equal(t, obj2.TestBit(), obj.TestBit())
	assert.Equal(t, obj2.TestVarchar(), obj.TestVarchar())
	assert.Equal(t, obj2.TestBlob(), obj.TestBlob())
}

func TestTypeTest_References(t *testing.T) {
	obj := createMaximalSampleTypeTest()
	ctx := db.NewContext(nil)
	obj.Save(ctx)
	defer deleteSampleTypeTest(ctx, obj)

	// Test that referenced objects were saved and assigned ids

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

	obj2 := LoadTypeTest(ctx, obj.PrimaryKey(), node.TypeTest().PrimaryKey())

	assert.Panics(t, func() { obj2.Date() })
	assert.Panics(t, func() { obj2.Time() })
	assert.Panics(t, func() { obj2.DateTime() })
	assert.Panics(t, func() { obj2.Ts() })
	assert.Panics(t, func() { obj2.TestInt() })
	assert.Panics(t, func() { obj2.TestFloat() })
	assert.Panics(t, func() { obj2.TestDouble() })
	assert.Panics(t, func() { obj2.TestText() })
	assert.Panics(t, func() { obj2.TestBit() })
	assert.Panics(t, func() { obj2.TestVarchar() })
	assert.Panics(t, func() { obj2.TestBlob() })
}
