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

	personID := test.RandomValue[string](0)
	obj.SetPersonID(personID)
	assert.Equal(t, personID, obj.PersonID())

	// test zero
	obj.SetPersonID("")
	assert.Equal(t, "", obj.PersonID(), "set empty")

}
func TestEmployeeInfo_SetEmployeeNumber(t *testing.T) {

	obj := NewEmployeeInfo()

	employeeNumber := test.RandomValue[int](0)
	obj.SetEmployeeNumber(employeeNumber)
	assert.Equal(t, employeeNumber, obj.EmployeeNumber())

	// test zero
	obj.SetEmployeeNumber(0)
	assert.Equal(t, 0, obj.EmployeeNumber(), "set empty")

}

// createMinimalSampleEmployeeInfo creates and saves a minimal version of a EmployeeInfo object
// for testing.
func createMinimalSampleEmployeeInfo(ctx context.Context) *EmployeeInfo {
	obj := NewEmployeeInfo()

	// A required forward reference will need to be fulfilled just to save the minimal version of this object
	objPerson := createMinimalSamplePerson(ctx)
	obj.SetPerson(objPerson)

	employeeNumber := test.RandomValue[int](0)
	obj.SetEmployeeNumber(employeeNumber)

	obj.Save(ctx)
	return obj
}
func TestEmployeeInfo_CRUD(t *testing.T) {
	obj := NewEmployeeInfo()
	ctx := db.NewContext(nil)

	objPerson := createMinimalSamplePerson(ctx)
	obj.SetPerson(objPerson)

	employeeNumber := test.RandomValue[int](0)
	obj.SetEmployeeNumber(employeeNumber)

	// Test retrieval
	obj = LoadEmployeeInfo(ctx, obj.PrimaryKey())
	require.NotNil(t, obj)

	assert.True(t, obj.IDIsValid())
	assert.NotEmpty(t, obj.ID())

	assert.True(t, obj.PersonIDIsValid())
	assert.NotEmpty(t, obj.PersonID())

	assert.True(t, obj.EmployeeNumberIsValid())
	assert.Equal(t, employeeNumber, obj.EmployeeNumber())

}
