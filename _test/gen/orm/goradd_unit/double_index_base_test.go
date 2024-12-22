// Code generated by goradd-orm. DO NOT EDIT.

package goradd_unit

import (
	"testing"

	strings2 "github.com/goradd/strings"
	"github.com/stretchr/testify/assert"
)

func TestDoubleIndex_SetFieldString(t *testing.T) {
	obj := NewDoubleIndex()

	fieldString := strings2.RandomString(strings2.AlphaAll, 10)
	obj.SetFieldString(fieldString)
	assert.Equal(t, fieldString, obj.FieldString())

	obj.SetFieldString("")
	assert.Equal(t, "", obj.FieldString(), "set empty")

}
