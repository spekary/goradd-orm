// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"testing"

	strings2 "github.com/goradd/strings"
	"github.com/stretchr/testify/assert"
)

func TestMilestone_SetName(t *testing.T) {
	obj := NewMilestone()

	name := strings2.RandomString(strings2.AlphaAll, 10)
	obj.SetName(name)
	assert.Equal(t, name, obj.Name())

	obj.SetName("")
	assert.Equal(t, "", obj.Name(), "set empty")

}
