package node

import (
	"testing"

	"github.com/goradd/orm/pkg/query"
	"github.com/stretchr/testify/assert"
)

func TestSerializeEmployeeInfoTable(t *testing.T) {
	var n query.NodeI = EmployeeInfo()

	assert.Equal(t, "employee_info", n.TableName_())
	assert.Equal(t, query.TableNodeType, n.NodeType_())
	assert.Equal(t, "goradd", n.DatabaseKey_())

	n2 := serNode(t, n)

	assert.Equal(t, "employee_info", n2.TableName_())
	assert.Equal(t, query.TableNodeType, n2.NodeType_())
	assert.Equal(t, "goradd", n2.DatabaseKey_())

}