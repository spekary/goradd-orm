// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"testing"

	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestForwardCascadeUnique_SetName(t *testing.T) {

	obj := NewForwardCascadeUnique()

	name := test.RandomValue[string](100)
	obj.SetName(name)
	assert.Equal(t, name, obj.Name())

	// test zero
	obj.SetName("")
	assert.Equal(t, "", obj.Name(), "set empty")

	// test panic on setting value larger than maximum size allowed
	name = test.RandomValue[string](101)
	assert.Panics(t, func() {
		obj.SetName(name)
	})
}
func TestForwardCascadeUnique_SetReverseID(t *testing.T) {

	obj := NewForwardCascadeUnique()

	reverseID := test.RandomValue[string](0)
	obj.SetReverseID(reverseID)
	assert.Equal(t, reverseID, obj.ReverseID())
	assert.False(t, obj.ReverseIDIsNull())

	// Test nil
	obj.SetReverseID(nil)
	assert.Equal(t, "", obj.ReverseID(), "set nil")
	assert.True(t, obj.ReverseIDIsNull())

	// test zero
	obj.SetReverseID("")
	assert.Equal(t, "", obj.ReverseID(), "set empty")
	assert.False(t, obj.ReverseIDIsNull())

}
