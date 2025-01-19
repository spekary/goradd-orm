package node

import (
	"testing"

	"github.com/goradd/orm/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestSerializeTableForwardNullUniqueTable(t *testing.T) {
	var n query.Node = ForwardNullUnique()

	assert.Equal(t, "forward_null_unique", n.TableName_())
	assert.Equal(t, query.TableNodeType, n.NodeType_())
	assert.Equal(t, "goradd_unit", n.DatabaseKey_())

	n2 := serNode(t, n)

	assert.Equal(t, "forward_null_unique", n2.TableName_())
	assert.Equal(t, query.TableNodeType, n2.NodeType_())
	assert.Equal(t, "goradd_unit", n2.DatabaseKey_())

	nodes := forwardNullUniqueTable{}.ColumnNodes_()
	for _, cn := range nodes {
		cn2 := serNode(t, cn)
		assert.Equal(t, "forward_null_unique", cn2.TableName_())
		assert.Equal(t, query.TableNodeType, query.NodeParent(cn2).NodeType_())
	}
}

func TestSerializeReferencesForwardNullUniqueTable(t *testing.T) {

	{
		n := ForwardNullUnique().Reverse()
		n2 := serNode(t, n)
		parentNode := query.NodeParent(n2)
		assert.Equal(t, query.TableNodeType, parentNode.NodeType_())
		assert.Equal(t, "forward_null_unique", parentNode.TableName_())

		nodes := n.(query.TableNodeI).ColumnNodes_()
		for _, cn := range nodes {
			cn2 := serNode(t, cn)
			assert.Equal(t, n.TableName_(), cn2.TableName_())
			assert.Equal(t, query.ReferenceNodeType, query.NodeParent(cn2).NodeType_())
		}
	}

}

func TestSerializeReverseReferencesForwardNullUniqueTable(t *testing.T) {
}

func TestSerializeAssociationsForwardNullUniqueTable(t *testing.T) {
}
