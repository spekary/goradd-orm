//** This file was code generated by GoT. DO NOT EDIT. ***

package template

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/goradd/orm/pkg/codegen"
	"github.com/goradd/orm/pkg/model"
)

func init() {
	codegen.RegisterTemplate(new(NodeTestTemplate))
	codegen.RegisterTemplate(new(NodeTestUtilsTemplate))
}

// NodeTestUtilsTemplate generates test code for the node package.
type NodeTestUtilsTemplate struct{}

func (n NodeTestUtilsTemplate) FileName(dbKey string) string {
	return filepath.Join("orm", dbKey, "node", "utils_test.go")
}

func (n NodeTestUtilsTemplate) GenerateDatabase(database *model.Database, _w io.Writer) (err error) {

	//*** testutils.tmpl

	if _, err = io.WriteString(_w, `package node

func serNode(t *testing.T, n query.NodeI) query.NodeI {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(&n)
	assert.NoError(t, err)

	var n2 query.NodeI
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&n2)
	assert.NoError(t, err)
	return n2
}

`); err != nil {
		return
	}

	return
}

func (n NodeTestUtilsTemplate) Overwrite() bool {
	return true
}

// NodeTestTemplate generates test code for table nodes.
type NodeTestTemplate struct {
}

func (n *NodeTestTemplate) FileName(table *model.Table) string {
	return filepath.Join("orm", table.DbKey, "node", table.FileName()+"_test.go")
}

func (n *NodeTestTemplate) GenerateTable(table *model.Table, _w io.Writer, importPath string) (err error) {

	//*** test.tmpl

	if _, err = io.WriteString(_w, `package node

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSerializeTable`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Table(t *testing.T) {
	var n query.NodeI = `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `()

    assert.Equal(t, "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `", n.TableName_())
    assert.Equal(t, query.TableNodeType, n.NodeType_())
    assert.Equal(t, "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DbKey); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `", n.DatabaseKey_())

	n2 := serNode(t, n)

    assert.Equal(t, "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `", n2.TableName_())
    assert.Equal(t, query.TableNodeType, n2.NodeType_())
    assert.Equal(t, "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DbKey); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `", n2.DatabaseKey_())

    nodes := `); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.DecapIdentifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Table{}.ColumnNodes_()
    for _,cn := range nodes {
        cn2 := serNode(t, cn)
        assert.Equal(t, "`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.QueryName); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `", cn2.TableName_())
        assert.Implements(t, (*query.NodeLinker)(nil), cn2)
        assert.Equal(t, query.TableNodeType, cn2.(query.NodeLinker).Parent().NodeType_())
    }
}

func TestSerializeReferences`); err != nil {
		return
	}

	if _, err = io.WriteString(_w, table.Identifier); err != nil {
		return
	}

	if _, err = io.WriteString(_w, `Table(t *testing.T) {
`); err != nil {
		return
	}

	for _, col := range table.Columns {

		if col.IsReference() {

			if _, err = io.WriteString(_w, `
{
    n := `); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.Identifier); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `().`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, fmt.Sprint(col.ReferenceIdentifier())); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `()
    n2 := serNode(t, n)
    parentNode := n2.(query.NodeLinker).Parent()
    assert.Equal(t, query.TableNodeType, parentNode.NodeType_())
    assert.Equal(t, "`); err != nil {
				return
			}

			if _, err = io.WriteString(_w, table.QueryName); err != nil {
				return
			}

			if _, err = io.WriteString(_w, `", parentNode.TableName_())

    nodes := n.(query.TableNodeI).ColumnNodes_()
    for _,cn := range nodes {
        cn2 := serNode(t, cn)
        assert.Equal(t, n.TableName_(), cn2.TableName_())
        assert.Implements(t, (*query.NodeLinker)(nil), cn2)
        assert.Equal(t, query.ReferenceNodeType, cn2.(query.NodeLinker).Parent().NodeType_())
    }
}

`); err != nil {
				return
			}

		}

	}

	if _, err = io.WriteString(_w, `}

`); err != nil {
		return
	}

	return
}

func (n *NodeTestTemplate) Overwrite() bool {
	return true
}
