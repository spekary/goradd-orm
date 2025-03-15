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

func TestEmployeeInfo_SetPersonID(t *testing.T) {

	obj := NewEmployeeInfo()
	val := test.RandomValue[string](0)
	obj.SetPersonID(val)
	assert.Equal(t, val, obj.PersonID())

	// test default
	obj.SetPersonID("")
	assert.EqualValues(t, "", obj.PersonID(), "set default")

}
func TestEmployeeInfo_SetEmployeeNumber(t *testing.T) {

	obj := NewEmployeeInfo()
	val := test.RandomValue[int](32)
	obj.SetEmployeeNumber(val)
	assert.Equal(t, val, obj.EmployeeNumber())

	// test default
	obj.SetEmployeeNumber(0)
	assert.EqualValues(t, 0, obj.EmployeeNumber(), "set default")

}

// createMinimalSampleEmployeeInfo creates and saves a minimal version of a EmployeeInfo object
// for testing.
func createMinimalSampleEmployeeInfo(ctx context.Context) *EmployeeInfo {
	obj := NewEmployeeInfo()

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	obj.SetPerson(createMinimalSamplePerson(ctx))

	obj.SetEmployeeNumber(test.RandomValue[int](32))

	obj.Save(ctx)
	return obj
}
func TestEmployeeInfo_CRUD(t *testing.T) {
	obj := NewEmployeeInfo()
	ctx := db.NewContext(nil)

	v_objPerson := createMinimalSamplePerson(ctx)
	defer v_objPerson.Delete(ctx)
	obj.SetPerson(v_objPerson)

	v_employeeNumber := test.RandomValue[int](32)
	obj.SetEmployeeNumber(v_employeeNumber)

	err := obj.Save(ctx)
	assert.NoError(t, err)
	defer obj.Delete(ctx)

	// Test retrieval
	obj2 := LoadEmployeeInfo(ctx, obj.PrimaryKey())
	require.NotNil(t, obj2)

	assert.True(t, obj2.IDIsValid())

	assert.True(t, obj2.PersonIDIsValid())
	assert.NotEmpty(t, obj2.PersonID())

	assert.True(t, obj2.EmployeeNumberIsValid())
	assert.EqualValues(t, v_employeeNumber, obj2.EmployeeNumber())

}
