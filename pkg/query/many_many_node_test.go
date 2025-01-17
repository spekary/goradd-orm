package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManyManyNodeInterfaces(t *testing.T) {
	n := &ManyManyNode{
		AssnTableQueryName:       "table",
		ParentColumnQueryName:    "col1",
		ParentColumnReceiverType: ColTypeString,
		Identifier:               "Field1",
		RefColumnQueryName:       "col2",
		RefColumnReceiverType:    ColTypeInteger,
	}

	assert.Implements(t, (*Conditioner)(nil), n)
	assert.Implements(t, (*Linker)(nil), n)
	assert.Implements(t, (*Expander)(nil), n)
}
