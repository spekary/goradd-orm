// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"testing"
	"time"

	"github.com/goradd/orm/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestPersonWithLock_SetFirstName(t *testing.T) {

	obj := NewPersonWithLock()

	firstName := test.RandomValue[string](50)
	obj.SetFirstName(firstName)
	assert.Equal(t, firstName, obj.FirstName())

	// test zero
	obj.SetFirstName("")
	assert.Equal(t, "", obj.FirstName(), "set empty")

	// test panic on setting value larger than maximum size allowed
	firstName = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetFirstName(firstName)
	})
}
func TestPersonWithLock_SetLastName(t *testing.T) {

	obj := NewPersonWithLock()

	lastName := test.RandomValue[string](50)
	obj.SetLastName(lastName)
	assert.Equal(t, lastName, obj.LastName())

	// test zero
	obj.SetLastName("")
	assert.Equal(t, "", obj.LastName(), "set empty")

	// test panic on setting value larger than maximum size allowed
	lastName = test.RandomValue[string](51)
	assert.Panics(t, func() {
		obj.SetLastName(lastName)
	})
}
func TestPersonWithLock_SetSysTimestamp(t *testing.T) {

	obj := NewPersonWithLock()

	sysTimestamp := test.RandomValue[time.Time](0)
	obj.SetSysTimestamp(sysTimestamp)
	assert.Equal(t, sysTimestamp, obj.SysTimestamp())
	assert.False(t, obj.SysTimestampIsNull())

	// Test nil
	obj.SetSysTimestamp(nil)
	assert.Equal(t, time.Time{}, obj.SysTimestamp(), "set nil")
	assert.True(t, obj.SysTimestampIsNull())

	// test zero
	obj.SetSysTimestamp(time.Time{})
	assert.Equal(t, time.Time{}, obj.SysTimestamp(), "set empty")
	assert.False(t, obj.SysTimestampIsNull())

}
