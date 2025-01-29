// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"context"
	"testing"
	"time"

	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTypeTest_SetDate(t *testing.T) {

	obj := NewTypeTest()
	date := test.RandomValue[time.Time](0)
	obj.SetDate(date)
	date = obj.Date()
	assert.Zero(t, date.Minute()) // make sure time part is zero'd
	assert.Zero(t, date.Hour())   // make sure time part is zero'd
	assert.Zero(t, date.Second()) // make sure time part is zero'd
	assert.False(t, obj.DateIsNull())

	// Test nil
	obj.SetDate(nil)
	assert.Equal(t, time.Time{}, obj.Date(), "set nil")
	assert.True(t, obj.DateIsNull())

	// test default
	obj.SetDate(time.Time{})
	assert.EqualValues(t, time.Time{}, obj.Date(), "set default")

	assert.False(t, obj.DateIsNull())

}
func TestTypeTest_SetTime(t *testing.T) {

	obj := NewTypeTest()
	time := test.RandomValue[time.Time](0)
	obj.SetTime(time)
	time = obj.Time()
	assert.Zero(t, time.Year())      // make sure minute part is zero'd
	assert.Equal(t, 1, time.Month()) // make sure minute part is zero'd
	assert.Equal(t, 1, time.Day())   // make sure minute part is zero'd
	assert.False(t, obj.TimeIsNull())

	// Test nil
	obj.SetTime(nil)
	assert.Equal(t, time.Time{}, obj.Time(), "set nil")
	assert.True(t, obj.TimeIsNull())

	// test default
	obj.SetTime(time.Time{})
	assert.EqualValues(t, time.Time{}, obj.Time(), "set default")

	assert.False(t, obj.TimeIsNull())

}
func TestTypeTest_SetDateTime(t *testing.T) {

	obj := NewTypeTest()
	dateTime := test.RandomValue[time.Time](0)
	obj.SetDateTime(dateTime)
	assert.Equal(t, dateTime, obj.DateTime())
	assert.False(t, obj.DateTimeIsNull())

	// Test nil
	obj.SetDateTime(nil)
	assert.Equal(t, time.Time{}, obj.DateTime(), "set nil")
	assert.True(t, obj.DateTimeIsNull())

	// test default
	obj.SetDateTime(time.Time{})
	assert.EqualValues(t, time.Time{}, obj.DateTime(), "set default")

	assert.False(t, obj.DateTimeIsNull())

}
func TestTypeTest_SetTestInt(t *testing.T) {

	obj := NewTypeTest()
	testInt := test.RandomValue[int](32)
	obj.SetTestInt(testInt)
	assert.Equal(t, testInt, obj.TestInt())
	assert.False(t, obj.TestIntIsNull())

	// Test nil
	obj.SetTestInt(nil)
	assert.Equal(t, 5, obj.TestInt(), "set nil")
	assert.True(t, obj.TestIntIsNull())

	// test default
	obj.SetTestInt(5)
	assert.EqualValues(t, 5, obj.TestInt(), "set default")

	assert.False(t, obj.TestIntIsNull())

}
func TestTypeTest_SetTestFloat(t *testing.T) {

	obj := NewTypeTest()
	testFloat := test.RandomValue[float32](32)
	obj.SetTestFloat(testFloat)
	assert.Equal(t, testFloat, obj.TestFloat())
	assert.False(t, obj.TestFloatIsNull())

	// Test nil
	obj.SetTestFloat(nil)
	assert.Equal(t, 0, obj.TestFloat(), "set nil")
	assert.True(t, obj.TestFloatIsNull())

	// test default
	obj.SetTestFloat(0)
	assert.EqualValues(t, 0, obj.TestFloat(), "set default")

	assert.False(t, obj.TestFloatIsNull())

}
func TestTypeTest_SetTestDouble(t *testing.T) {

	obj := NewTypeTest()
	testDouble := test.RandomValue[float64](64)
	obj.SetTestDouble(testDouble)
	assert.Equal(t, testDouble, obj.TestDouble())

	// test default
	obj.SetTestDouble(0)
	assert.EqualValues(t, 0, obj.TestDouble(), "set default")

}
func TestTypeTest_SetTestText(t *testing.T) {

	obj := NewTypeTest()
	testText := test.RandomValue[string](65535)
	obj.SetTestText(testText)
	assert.Equal(t, testText, obj.TestText())
	assert.False(t, obj.TestTextIsNull())

	// Test nil
	obj.SetTestText(nil)
	assert.Equal(t, "", obj.TestText(), "set nil")
	assert.True(t, obj.TestTextIsNull())

	// test default
	obj.SetTestText("")
	assert.EqualValues(t, "", obj.TestText(), "set default")

	assert.False(t, obj.TestTextIsNull())

	// test panic on setting value larger than maximum size allowed
	testText = test.RandomValue[string](65536)
	assert.Panics(t, func() {
		obj.SetTestText(testText)
	})
}
func TestTypeTest_SetTestBit(t *testing.T) {

	obj := NewTypeTest()
	testBit := test.RandomValue[bool](0)
	obj.SetTestBit(testBit)
	assert.Equal(t, testBit, obj.TestBit())
	assert.False(t, obj.TestBitIsNull())

	// Test nil
	obj.SetTestBit(nil)
	assert.Equal(t, false, obj.TestBit(), "set nil")
	assert.True(t, obj.TestBitIsNull())

	// test default
	obj.SetTestBit(false)
	assert.EqualValues(t, false, obj.TestBit(), "set default")

	assert.False(t, obj.TestBitIsNull())

}
func TestTypeTest_SetTestVarchar(t *testing.T) {

	obj := NewTypeTest()
	testVarchar := test.RandomValue[string](10)
	obj.SetTestVarchar(testVarchar)
	assert.Equal(t, testVarchar, obj.TestVarchar())
	assert.False(t, obj.TestVarcharIsNull())

	// Test nil
	obj.SetTestVarchar(nil)
	assert.Equal(t, "", obj.TestVarchar(), "set nil")
	assert.True(t, obj.TestVarcharIsNull())

	// test default
	obj.SetTestVarchar("")
	assert.EqualValues(t, "", obj.TestVarchar(), "set default")

	assert.False(t, obj.TestVarcharIsNull())

	// test panic on setting value larger than maximum size allowed
	testVarchar = test.RandomValue[string](11)
	assert.Panics(t, func() {
		obj.SetTestVarchar(testVarchar)
	})
}
func TestTypeTest_SetTestBlob(t *testing.T) {

	obj := NewTypeTest()
	testBlob := test.RandomValue[[]byte](65535)
	obj.SetTestBlob(testBlob)
	assert.Equal(t, testBlob, obj.TestBlob())

	// test default
	obj.SetTestBlob([]byte(nil))
	assert.EqualValues(t, []byte(nil), obj.TestBlob(), "set default")

	// test panic on setting value larger than maximum size allowed
	testBlob = test.RandomValue[[]byte](65536)
	assert.Panics(t, func() {
		obj.SetTestBlob(testBlob)
	})
}

// createMinimalSampleTypeTest creates and saves a minimal version of a TypeTest object
// for testing.
func createMinimalSampleTypeTest(ctx context.Context) *TypeTest {
	obj := NewTypeTest()

	date := test.RandomValue[time.Time](0)
	obj.SetDate(date)

	time := test.RandomValue[time.Time](0)
	obj.SetTime(time)

	dateTime := test.RandomValue[time.Time](0)
	obj.SetDateTime(dateTime)

	testInt := test.RandomValue[int](32)
	obj.SetTestInt(testInt)

	testFloat := test.RandomValue[float32](32)
	obj.SetTestFloat(testFloat)

	testDouble := test.RandomValue[float64](64)
	obj.SetTestDouble(testDouble)

	testText := test.RandomValue[string](65535)
	obj.SetTestText(testText)

	testBit := test.RandomValue[bool](0)
	obj.SetTestBit(testBit)

	testVarchar := test.RandomValue[string](10)
	obj.SetTestVarchar(testVarchar)

	testBlob := test.RandomValue[[]byte](65535)
	obj.SetTestBlob(testBlob)

	obj.Save(ctx)
	return obj
}
func TestTypeTest_CRUD(t *testing.T) {
	obj := NewTypeTest()
	ctx := db.NewContext(nil)

	date := test.RandomValue[time.Time](0)
	obj.SetDate(date)

	date = obj.Date()

	time := test.RandomValue[time.Time](0)
	obj.SetTime(time)

	time = obj.Time()

	dateTime := test.RandomValue[time.Time](0)
	obj.SetDateTime(dateTime)

	dateTime = obj.DateTime()

	testInt := test.RandomValue[int](32)
	obj.SetTestInt(testInt)

	testFloat := test.RandomValue[float32](32)
	obj.SetTestFloat(testFloat)

	testDouble := test.RandomValue[float64](64)
	obj.SetTestDouble(testDouble)

	testText := test.RandomValue[string](65535)
	obj.SetTestText(testText)

	testBit := test.RandomValue[bool](0)
	obj.SetTestBit(testBit)

	testVarchar := test.RandomValue[string](10)
	obj.SetTestVarchar(testVarchar)

	testBlob := test.RandomValue[[]byte](65535)
	obj.SetTestBlob(testBlob)

	obj.Save(ctx)

	// Test retrieval
	obj = LoadTypeTest(ctx, obj.PrimaryKey())
	require.NotNil(t, obj)

	assert.True(t, obj.IDIsValid())

	assert.True(t, obj.DateIsValid())
	assert.False(t, obj.DateIsNull())
	assert.Equal(t, date, obj.Date())

	assert.True(t, obj.TimeIsValid())
	assert.False(t, obj.TimeIsNull())
	assert.Equal(t, time, obj.Time())

	assert.True(t, obj.DateTimeIsValid())
	assert.False(t, obj.DateTimeIsNull())
	assert.Equal(t, dateTime, obj.DateTime())

	assert.True(t, obj.TsIsValid())
	assert.False(t, obj.TsIsNull())

	assert.True(t, obj.TestIntIsValid())
	assert.False(t, obj.TestIntIsNull())
	assert.Equal(t, testInt, obj.TestInt())

	assert.True(t, obj.TestFloatIsValid())
	assert.False(t, obj.TestFloatIsNull())
	assert.Equal(t, testFloat, obj.TestFloat())

	assert.True(t, obj.TestDoubleIsValid())
	assert.Equal(t, testDouble, obj.TestDouble())

	assert.True(t, obj.TestTextIsValid())
	assert.False(t, obj.TestTextIsNull())
	assert.Equal(t, testText, obj.TestText())

	assert.True(t, obj.TestBitIsValid())
	assert.False(t, obj.TestBitIsNull())
	assert.Equal(t, testBit, obj.TestBit())

	assert.True(t, obj.TestVarcharIsValid())
	assert.False(t, obj.TestVarcharIsNull())
	assert.Equal(t, testVarchar, obj.TestVarchar())

	assert.True(t, obj.TestBlobIsValid())
	assert.Equal(t, testBlob, obj.TestBlob())

}
