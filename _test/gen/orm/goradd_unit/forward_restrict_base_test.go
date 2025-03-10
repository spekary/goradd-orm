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

func TestForwardRestrict_SetName(t *testing.T) {

	obj := NewForwardRestrict()
	name := test.RandomValue[string](100)
	obj.SetName(name)
	assert.Equal(t, name, obj.Name())

	// test default
	obj.SetName("")
	assert.EqualValues(t, "", obj.Name(), "set default")

	// test panic on setting value larger than maximum size allowed
	name = test.RandomValue[string](101)
	assert.Panics(t, func() {
		obj.SetName(name)
	})
}
func TestForwardRestrict_SetReverseID(t *testing.T) {

	obj := NewForwardRestrict()
	reverseID := test.RandomValue[string](0)
	obj.SetReverseID(reverseID)
	assert.Equal(t, reverseID, obj.ReverseID())

	// test default
	obj.SetReverseID("")
	assert.EqualValues(t, "", obj.ReverseID(), "set default")

}

// createMinimalSampleForwardRestrict creates and saves a minimal version of a ForwardRestrict object
// for testing.
func createMinimalSampleForwardRestrict(ctx context.Context) *ForwardRestrict {
	obj := NewForwardRestrict()

	name := test.RandomValue[string](100)
	obj.SetName(name)

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	objReverse := createMinimalSampleReverse(ctx)
	obj.SetReverse(objReverse)

	obj.Save(ctx)
	return obj
}
func TestForwardRestrict_CRUD(t *testing.T) {
	obj := NewForwardRestrict()
	ctx := db.NewContext(nil)

	name := test.RandomValue[string](100)
	obj.SetName(name)

	objReverse := createMinimalSampleReverse(ctx)
	defer objReverse.Delete(ctx)
	obj.SetReverse(objReverse)

	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer obj.Delete(ctx)

	// Test retrieval
	obj2 := LoadForwardRestrict(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.NameIsValid())
	assert.Equal(t, name, obj2.Name())

	assert.True(t, obj2.ReverseIDIsValid())
	assert.NotEmpty(t, obj2.ReverseID())

}
