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

func TestLeaf_SetName(t *testing.T) {

	obj := NewLeaf()
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

// createMinimalSampleLeaf creates and saves a minimal version of a Leaf object
// for testing.
func createMinimalSampleLeaf(ctx context.Context) *Leaf {
	obj := NewLeaf()

	name := test.RandomValue[string](100)
	obj.SetName(name)

	obj.Save(ctx)
	return obj
}
func TestLeaf_CRUD(t *testing.T) {
	obj := NewLeaf()
	ctx := db.NewContext(nil)

	name := test.RandomValue[string](100)
	obj.SetName(name)

	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer obj.Delete(ctx)

	// Test retrieval
	obj2 := LoadLeaf(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.NameIsValid())
	assert.Equal(t, name, obj2.Name())

}
