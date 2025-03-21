// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"testing"

	"github.com/goradd/orm/pkg/db"
	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLeafLock_SetName(t *testing.T) {

	obj := NewLeafLock()
	val := test.RandomValue[string](100)
	obj.SetName(val)
	assert.Equal(t, val, obj.Name())

	// test default
	obj.SetName("")
	assert.EqualValues(t, "", obj.Name(), "set default")

	// test panic on setting value larger than maximum size allowed
	val = test.RandomValue[string](101)
	assert.Panics(t, func() {
		obj.SetName(val)
	})
}

// createMinimalSampleLeafLock creates an unsaved minimal version of a LeafLock object
// for testing.
func createMinimalSampleLeafLock() *LeafLock {
	obj := NewLeafLock()

	obj.SetName(test.RandomValue[string](100))

	return obj
}

// createMaximalSampleLeafLock creates an unsaved version of a LeafLock object
// for testing that includes references to minimal objects.
func createMaximalSampleLeafLock() *LeafLock {
	obj := createMinimalSampleLeafLock()

	return obj
}

func TestLeafLock_CRUD(t *testing.T) {
	obj := NewLeafLock()
	ctx := db.NewContext(nil)

	v_name := test.RandomValue[string](100)
	obj.SetName(v_name)

	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer obj.Delete(ctx)

	// Test retrieval
	obj2 := LoadLeafLock(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.NameIsValid())
	assert.EqualValues(t, v_name, obj2.Name())

	assert.True(t, obj2.GroLockIsValid())

}
