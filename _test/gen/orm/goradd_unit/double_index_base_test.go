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

	id := test.RandomValue[int](32)
	obj.SetID(id)
	assert.Equal(t, id, obj.ID())

	// test zero
	obj.SetID(0)
	assert.Equal(t, 0, obj.ID(), "set empty")

}
func TestDoubleIndex_SetFieldInt(t *testing.T) {

	obj := NewDoubleIndex()

	fieldInt := test.RandomValue[int](32)
	obj.SetFieldInt(fieldInt)
	assert.Equal(t, fieldInt, obj.FieldInt())

	// test zero
	obj.SetFieldInt(0)
	assert.Equal(t, 0, obj.FieldInt(), "set empty")

}
func TestDoubleIndex_SetFieldString(t *testing.T) {

	obj := NewDoubleIndex()

	fieldString := test.RandomValue[string](50)
	obj.SetFieldString(fieldString)
	assert.Equal(t, fieldString, obj.FieldString())

	// test zero
	obj.SetFieldString("")
	assert.Equal(t, "", obj.FieldString(), "set empty")

	// test panic on setting value larger than maximum size allowed
	fieldString = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetFieldString(fieldString)
	})
}

// createMinimalSampleDoubleIndex creates and saves a minimal version of a DoubleIndex object
// for testing.
func createMinimalSampleDoubleIndex(ctx context.Context) *DoubleIndex {
	obj := NewDoubleIndex()

	id := test.RandomValue[int](32)
	obj.SetID(id)

	fieldInt := test.RandomValue[int](32)
	obj.SetFieldInt(fieldInt)

	fieldString := test.RandomValue[string](50)
	obj.SetFieldString(fieldString)

	obj.Save(ctx)
	return obj
}
func TestDoubleIndex_CRUD(t *testing.T) {
	obj := NewDoubleIndex()
	ctx := db.NewContext(nil)

	id := test.RandomValue[int](32)
	obj.SetID(id)

	fieldInt := test.RandomValue[int](32)
	obj.SetFieldInt(fieldInt)

	fieldString := test.RandomValue[string](50)
	obj.SetFieldString(fieldString)

	// Test retrieval
	obj = LoadDoubleIndex(ctx, obj.PrimaryKey())
	require.NotNil(t, obj)

	assert.True(t, obj.IDIsValid())
	assert.Equal(t, id, obj.ID())

	assert.True(t, obj.FieldIntIsValid())
	assert.Equal(t, fieldInt, obj.FieldInt())

	assert.True(t, obj.FieldStringIsValid())
	assert.Equal(t, fieldString, obj.FieldString())

}
