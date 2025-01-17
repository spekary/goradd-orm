package node

import (
	"testing"

	"github.com/goradd/orm/pkg/query"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSerializeTableUnsupportedTypeTable(t *testing.T) {
	var n query.Node = UnsupportedType()

	assert.Equal(t, "unsupported_type", n.TableName_())
	assert.Equal(t, query.TableNodeType, n.NodeType_())
	assert.Equal(t, "goradd_unit", n.DatabaseKey_())

	n2 := serNode(t, n)

	assert.Equal(t, "unsupported_type", n2.TableName_())
	assert.Equal(t, query.TableNodeType, n2.NodeType_())
	assert.Equal(t, "goradd_unit", n2.DatabaseKey_())

	nodes := unsupportedTypeTable{}.ColumnNodes_()
	for _, cn := range nodes {
		cn2 := serNode(t, cn)
		assert.Equal(t, "unsupported_type", cn2.TableName_())
		require.Implements(t, (*query.Linker)(nil), cn2)
		assert.Equal(t, query.TableNodeType, cn2.(query.Linker).Parent().NodeType_())
	}
}

func TestSerializeReferencesUnsupportedTypeTable(t *testing.T) {
}

func TestSerializeReverseReferencesUnsupportedTypeTable(t *testing.T) {
}

func TestSerializeAssociationsUnsupportedTypeTable(t *testing.T) {
}
