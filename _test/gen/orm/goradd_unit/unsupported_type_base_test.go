// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"context"
	"testing"

	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestUnsupportedType_SetTypeSet(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[string](5)
	obj.SetTypeSet(val)
	assert.Equal(t, val, obj.TypeSet())

	// test default
	obj.SetTypeSet("")
	assert.EqualValues(t, "", obj.TypeSet(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](6)
	assert.Panics(t, func() {
		obj.SetTypeSet(val)
	})
}
func TestUnsupportedType_SetTypeEnumerated(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[string](1)
	obj.SetTypeEnumerated(val)
	assert.Equal(t, val, obj.TypeEnumerated())

	// test default
	obj.SetTypeEnumerated("")
	assert.EqualValues(t, "", obj.TypeEnumerated(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](2)
	assert.Panics(t, func() {
		obj.SetTypeEnumerated(val)
	})
}
func TestUnsupportedType_SetTypeDecimal(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[[]byte](13)
	obj.SetTypeDecimal(val)
	assert.Equal(t, val, obj.TypeDecimal())

	// test default
	obj.SetTypeDecimal([]byte{})
	assert.EqualValues(t, []byte{}, obj.TypeDecimal(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[[]byte](14)
	assert.Panics(t, func() {
		obj.SetTypeDecimal(val)
	})
}
func TestUnsupportedType_SetTypeDouble(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[float64](64)
	obj.SetTypeDouble(val)
	assert.Equal(t, val, obj.TypeDouble())

	// test default
	obj.SetTypeDouble(0)
	assert.EqualValues(t, 0, obj.TypeDouble(), "set default")

}
func TestUnsupportedType_SetTypeGeo(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[[]byte](0)
	obj.SetTypeGeo(val)
	assert.Equal(t, val, obj.TypeGeo())

	// test default
	obj.SetTypeGeo([]byte{})
	assert.EqualValues(t, []byte{}, obj.TypeGeo(), "set default")

}
func TestUnsupportedType_SetTypeTinyBlob(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[[]byte](255)
	obj.SetTypeTinyBlob(val)
	assert.Equal(t, val, obj.TypeTinyBlob())

	// test default
	obj.SetTypeTinyBlob([]byte{})
	assert.EqualValues(t, []byte{}, obj.TypeTinyBlob(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[[]byte](256)
	assert.Panics(t, func() {
		obj.SetTypeTinyBlob(val)
	})
}
func TestUnsupportedType_SetTypeMediumBlob(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[[]byte](100000)
	obj.SetTypeMediumBlob(val)
	assert.Equal(t, val, obj.TypeMediumBlob())

	// test default
	obj.SetTypeMediumBlob([]byte{})
	assert.EqualValues(t, []byte{}, obj.TypeMediumBlob(), "set default")

}
func TestUnsupportedType_SetTypeVarbinary(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[[]byte](0)
	obj.SetTypeVarbinary(val)
	assert.Equal(t, val, obj.TypeVarbinary())

	// test default
	obj.SetTypeVarbinary([]byte{})
	assert.EqualValues(t, []byte{}, obj.TypeVarbinary(), "set default")

}
func TestUnsupportedType_SetTypeLongtext(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[string](100000)
	obj.SetTypeLongtext(val)
	assert.Equal(t, val, obj.TypeLongtext())

	// test default
	obj.SetTypeLongtext("")
	assert.EqualValues(t, "", obj.TypeLongtext(), "set default")

}
func TestUnsupportedType_SetTypeBinary(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[[]byte](0)
	obj.SetTypeBinary(val)
	assert.Equal(t, val, obj.TypeBinary())

	// test default
	obj.SetTypeBinary([]byte{})
	assert.EqualValues(t, []byte{}, obj.TypeBinary(), "set default")

}
func TestUnsupportedType_SetTypeSmall(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[int](16)
	obj.SetTypeSmall(val)
	assert.Equal(t, val, obj.TypeSmall())

	// test default
	obj.SetTypeSmall(0)
	assert.EqualValues(t, 0, obj.TypeSmall(), "set default")

}
func TestUnsupportedType_SetTypeMedium(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[int](24)
	obj.SetTypeMedium(val)
	assert.Equal(t, val, obj.TypeMedium())

	// test default
	obj.SetTypeMedium(0)
	assert.EqualValues(t, 0, obj.TypeMedium(), "set default")

}
func TestUnsupportedType_SetTypeBig(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[int64](64)
	obj.SetTypeBig(val)
	assert.Equal(t, val, obj.TypeBig())

	// test default
	obj.SetTypeBig(0)
	assert.EqualValues(t, 0, obj.TypeBig(), "set default")

}
func TestUnsupportedType_SetTypePolygon(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[[]byte](0)
	obj.SetTypePolygon(val)
	assert.Equal(t, val, obj.TypePolygon())

	// test default
	obj.SetTypePolygon([]byte{})
	assert.EqualValues(t, []byte{}, obj.TypePolygon(), "set default")

}
func TestUnsupportedType_SetTypeUnsigned(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[uint](32)
	obj.SetTypeUnsigned(val)
	assert.Equal(t, val, obj.TypeUnsigned())

	// test default
	obj.SetTypeUnsigned(0x0)
	assert.EqualValues(t, 0x0, obj.TypeUnsigned(), "set default")

}
func TestUnsupportedType_SetTypeMultfk1(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[string](50)
	obj.SetTypeMultfk1(val)
	assert.Equal(t, val, obj.TypeMultfk1())

	// test default
	obj.SetTypeMultfk1("")
	assert.EqualValues(t, "", obj.TypeMultfk1(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetTypeMultfk1(val)
	})
}
func TestUnsupportedType_SetTypeMultifk2(t *testing.T) {

	obj := NewUnsupportedType()
	val := test.RandomValue[string](50)
	obj.SetTypeMultifk2(val)
	assert.Equal(t, val, obj.TypeMultifk2())

	// test default
	obj.SetTypeMultifk2("")
	assert.EqualValues(t, "", obj.TypeMultifk2(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetTypeMultifk2(val)
	})
}

// createMinimalSampleUnsupportedType creates and saves a minimal version of a UnsupportedType object
// for testing.
func createMinimalSampleUnsupportedType(ctx context.Context) *UnsupportedType {
	obj := NewUnsupportedType()

	obj.SetTypeSet(test.RandomValue[string](5))

	obj.SetTypeEnumerated(test.RandomValue[string](1))

	obj.SetTypeDouble(test.RandomValue[float64](64))

	obj.SetTypeTinyBlob(test.RandomValue[[]byte](255))

	obj.SetTypeMediumBlob(test.RandomValue[[]byte](16777215))

	obj.SetTypeLongtext(test.RandomValue[string](4294967295))

	obj.SetTypeSmall(test.RandomValue[int](16))

	obj.SetTypeMedium(test.RandomValue[int](24))

	obj.SetTypeBig(test.RandomValue[int64](64))

	obj.SetTypeUnsigned(test.RandomValue[uint](32))

	obj.SetTypeMultfk1(test.RandomValue[string](50))

	obj.SetTypeMultifk2(test.RandomValue[string](50))

	obj.Save(ctx)
	return obj
}
