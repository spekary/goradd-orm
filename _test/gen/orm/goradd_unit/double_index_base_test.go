// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"context"
	"testing"

	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDoubleIndex_SetID(t *testing.T) {

	obj := NewDoubleIndex()
	val := test.RandomValue[int](32)
	obj.SetID(val)
	assert.Equal(t, val, obj.ID())

	// test default
	obj.SetID(0)
	assert.EqualValues(t, 0, obj.ID(), "set default")

}
func TestDoubleIndex_SetFieldInt(t *testing.T) {

	obj := NewDoubleIndex()
	val := test.RandomValue[int](32)
	obj.SetFieldInt(val)
	assert.Equal(t, val, obj.FieldInt())

	// test default
	obj.SetFieldInt(0)
	assert.EqualValues(t, 0, obj.FieldInt(), "set default")

}
func TestDoubleIndex_SetFieldString(t *testing.T) {

	obj := NewDoubleIndex()
	val := test.RandomValue[string](50)
	obj.SetFieldString(val)
	assert.Equal(t, val, obj.FieldString())

	// test default
	obj.SetFieldString("")
	assert.EqualValues(t, "", obj.FieldString(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetFieldString(val)
	})
}

// createMinimalSampleDoubleIndex creates and saves a minimal version of a DoubleIndex object
// for testing.
func createMinimalSampleDoubleIndex(ctx context.Context) *DoubleIndex {
	obj := NewDoubleIndex()

	obj.SetID(test.RandomValue[int](32))

	obj.SetFieldInt(test.RandomValue[int](32))

	obj.SetFieldString(test.RandomValue[string](50))

	obj.Save(ctx)
	return obj
}
func TestDoubleIndex_CRUD(t *testing.T) {
	obj := NewDoubleIndex()
	ctx := db.NewContext(nil)

	v_id := test.RandomValue[int](32)
	obj.SetID(v_id)

	v_fieldInt := test.RandomValue[int](32)
	obj.SetFieldInt(v_fieldInt)

	v_fieldString := test.RandomValue[string](50)
	obj.SetFieldString(v_fieldString)

	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer obj.Delete(ctx)

	// Test retrieval
	obj2 := LoadDoubleIndex(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.True(t, obj2.IDIsValid())
	assert.EqualValues(t, v_id, obj2.ID())

	assert.True(t, obj2.FieldIntIsValid())
	assert.EqualValues(t, v_fieldInt, obj2.FieldInt())

	assert.True(t, obj2.FieldStringIsValid())
	assert.EqualValues(t, v_fieldString, obj2.FieldString())

}
