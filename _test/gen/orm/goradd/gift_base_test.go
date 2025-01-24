// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"context"
	"testing"

	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGift_SetNumber(t *testing.T) {

	obj := NewGift()

	number := test.RandomValue[int](32)
	obj.SetNumber(number)
	assert.Equal(t, number, obj.Number())

	// test zero
	obj.SetNumber(0)
	assert.Equal(t, 0, obj.Number(), "set empty")

}
func TestGift_SetName(t *testing.T) {

	obj := NewGift()

	name := test.RandomValue[string](50)
	obj.SetName(name)
	assert.Equal(t, name, obj.Name())

	// test zero
	obj.SetName("")
	assert.Equal(t, "", obj.Name(), "set empty")

	// test panic on setting value larger than maximum size allowed
	name = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetName(name)
	})
}

// createMinimalSampleGift creates and saves a minimal version of a Gift object
// for testing.
func createMinimalSampleGift(ctx context.Context) *Gift {
	obj := NewGift()

	number := test.RandomValue[int](32)
	obj.SetNumber(number)

	name := test.RandomValue[string](50)
	obj.SetName(name)

	obj.Save(ctx)
	return obj
}
func TestGift_CRUD(t *testing.T) {
	obj := NewGift()
	ctx := db.NewContext(nil)

	number := test.RandomValue[int](32)
	obj.SetNumber(number)

	name := test.RandomValue[string](50)
	obj.SetName(name)

	obj.Save(ctx)

	// Test retrieval
	obj = LoadGift(ctx, obj.PrimaryKey())
	require.NotNil(t, obj)

	assert.True(t, obj.NumberIsValid())
	assert.Equal(t, number, obj.Number())

	assert.True(t, obj.NameIsValid())
	assert.Equal(t, name, obj.Name())

}
