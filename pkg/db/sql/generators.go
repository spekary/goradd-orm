package sql

import (
	"fmt"
	"github.com/goradd/iter"
	. "github.com/goradd/orm/pkg/query"
	"strings"
)

type operationSqler interface {
	OperationSql(op Operator, operandStrings []string) string
}

type deleteUsesAliaser interface {
	DeleteUsesAlias() bool
}

// selectGenerator is an aid to generating various sql statements.
// SQL dialects are similar, but have small variations. This object
// attempts to handle the major issues, while allowing individual
// implementations of SQL to do their own tweaks.
type selectGenerator struct {
	b       *Builder
	argList []any
}

func newSelectGenerator(builder *Builder) *selectGenerator {
	return &selectGenerator{b: builder}
}

func (g *selectGenerator) iq(v string) string {
	return g.b.db.QuoteIdentifier(v)
}

func (g *selectGenerator) addArg(v any) string {
	g.argList = append(g.argList, v)
	return g.b.db.FormatArgument(len(g.argList))
}

func (g *selectGenerator) generateSelectSql() (sql string) {
	if g.b.IsDistinct {
		sql = "SELECT DISTINCT\n"
	} else {
		sql = "SELECT\n"
	}

	sql += g.generateColumnListWithAliases()
	sql += g.generateFromSql()
	sql += g.generateWhereSql()
	sql += g.generateGroupBySql()
	sql += g.generateHaving()
	sql += g.generateOrderBySql()
	sql += g.generateLimitSql()
	return
}

func (g *selectGenerator) generateDeleteSql() (sql string) {
	if t, ok := g.b.db.(deleteUsesAliaser); ok && t.DeleteUsesAlias() {
		j := g.b.RootJoinTreeItem
		alias := g.iq(j.Alias)
		sql = "DELETE " + alias + "\n"
	} else {
		sql = "DELETE\n"
	}

	sql += g.generateFromSql()
	sql += g.generateWhereSql()
	sql += g.generateOrderBySql()
	sql += g.generateLimitSql()

	return
}

func (g *selectGenerator) generateColumnListWithAliases() (sql string) {
	g.b.ColumnAliases.Range(func(key string, j *JoinTreeItem) bool {
		sql += g.generateColumnNodeSql(j.Parent.Alias, j.Node) + " AS " + g.iq(key) + ",\n"
		return true
	})

	if g.b.AliasNodes != nil {
		g.b.AliasNodes.Range(func(key string, v Aliaser) bool {
			node := v.(NodeI)
			aliaser := v.(Aliaser)
			sql += g.generateNodeSql(node, false)
			alias := aliaser.GetAlias()
			if alias != "" {
				// This happens in a subquery
				sql += " AS " + g.iq(alias)
			}
			sql += ",\n"
			return true
		})
	}

	sql = strings.TrimSuffix(sql, ",\n")
	sql += "\n"
	return
}

// Generate the column node sql.
func (g *selectGenerator) generateColumnNodeSql(parentAlias string, node NodeI) (sql string) {
	return g.iq(parentAlias) + "." + g.iq(ColumnNodeDbName(node.(*ColumnNode)))
}

func (g *selectGenerator) generateNodeSql(n NodeI, useAlias bool) (sql string) {
	switch node := n.(type) {
	case *ValueNode:
		v := ValueNodeGetValue(node)
		if a, ok := v.([]NodeI); ok {
			// value is actually a list of nodes
			var l []string
			for _, o := range a {
				l = append(l, g.generateNodeSql(o, useAlias))
			}
			return strings.Join(l, ",")
		} else {
			return g.addArg(v)
		}
	case *OperationNode:
		return g.generateOperationSql(node, useAlias)
	case *ColumnNode:
		item := g.b.GetItemFromNode(node)
		if useAlias && item.Alias != "" {
			sql = g.generateAlias(item.Alias)
		} else {
			sql = g.generateColumnNodeSql(item.Parent.Alias, node)
		}
	case *AliasNode:
		if useAlias {
			sql = g.iq(node.GetAlias())
		} else {
			n := g.b.GetAliasedNode(node)
			if n != nil {
				sql = g.generateNodeSql(n, useAlias)
			}
		}

	case *SubqueryNode:
		sql = g.generateSubquerySql(node)
	case TableNodeI:
		tj := g.b.GetItemFromNode(node)
		sql = g.generateColumnNodeSql(tj.Alias, node.PrimaryKeyNode())
	default:
		panic("Can't generate sql from node type.")
	}
	return
}

func (g *selectGenerator) generateOperationSql(n *OperationNode, useAlias bool) (sql string) {
	if useAlias && n.GetAlias() != "" {
		sql = g.iq(n.GetAlias())
		return
	}

	var operands []string
	operator := OperationNodeOperator(n)

	for _, o := range OperationNodeOperands(n) {
		operands = append(operands, g.generateNodeSql(o, useAlias))
	}

	if o, ok := g.b.db.(operationSqler); ok {
		sql = o.OperationSql(OperationNodeOperator(n), operands)
		if sql != "" {
			return sql
		}
	}

	switch operator {
	case OpFunc:
		if len(operands) > 0 {
			sql = strings.Join(operands, ",")
		} else {
			if OperationNodeFunction(n) == "COUNT" {
				sql = "*"
			}
		}

		if OperationNodeDistinct(n) {
			sql = "DISTINCT " + sql
		}
		sql = OperationNodeFunction(n) + "(" + sql + ") "

	case OpNull:
		fallthrough
	case OpNotNull:
		s := operands[0]
		sql = s + " IS " + OperationNodeOperator(n).String()
		sql = "(" + sql + ") "

	case OpNot:
		s := operands[0]
		sql = OperationNodeOperator(n).String() + " " + s
		sql = "(" + sql + ") "

	case OpIn:
		fallthrough
	case OpNotIn:
		s := operands[0]
		sql = s + " " + OperationNodeOperator(n).String()
		sql += " (" + operands[1] + ") "

	case OpAll:
		fallthrough
	case OpNone:
		sql = "(" + OperationNodeOperator(n).String() + ") "
	case OpStartsWith:
		// Last argument should be the 2nd operand
		s := g.argList[len(g.argList)-1].(string)
		s += "%"
		g.argList[len(g.argList)-1] = s
		sql = fmt.Sprintf(`(%s LIKE %s)`, operands[0], operands[1])
	case OpEndsWith:
		// SQL supports this with a LIKE operation
		s := g.argList[len(g.argList)-1].(string)
		s = "%" + s
		g.argList[len(g.argList)-1] = s
		sql = fmt.Sprintf(`(%s LIKE %s)`, operands[0], operands[1])
	case OpContains:
		// SQL supports this with a LIKE operation
		s := g.argList[len(g.argList)-1].(string)
		s = "%" + s + "%"
		g.argList[len(g.argList)-1] = s
		sql = fmt.Sprintf(`(%s LIKE %s)`, operands[0], operands[1])
	case OpDateAddSeconds:
		// Modifying a datetime in the query
		// Only works on date, datetime and timestamps. Not times.
		// This is highly SQL dialect dependent. Use the below as an example.
		// Default is to not implement it. To implement it, it must be overridden in the database implementation.

		panic("DateAddSeconds is not implemented in this database")
		/*
			s := operands[0]
			s2 := operands[1]
			sql = fmt.Sprintf(`DATE_ADD(%s, INTERVAL (%s) SECOND)`, s, s2)
		*/
	case OpXor:
		// Some sqls do not have an XOR operator, so we manually implement the code here
		// Override in the database implementation if XOR is implemented
		s := operands[0]
		s2 := operands[1]
		sql = fmt.Sprintf(`(((%[1]s) AND NOT (%[2]s)) OR (NOT (%[1]s) AND (%[2]s)))`, s, s2)

	default:
		sOp := " " + OperationNodeOperator(n).String() + " "
		sql = " (" + strings.Join(operands, sOp) + ") "
	}
	return
}

func (g *selectGenerator) generateAlias(alias string) (sql string) {
	return g.iq(alias)
}

func (g *selectGenerator) generateSubquerySql(node *SubqueryNode) (sql string) {
	// The copy below intentionally reuses the argList and db items
	g2 := *g
	g2.b = SubqueryBuilder(node).(*Builder)
	sql = g2.generateSelectSql()
	sql = "(" + sql + ")"
	return
}

func (g *selectGenerator) generateFromSql() (sql string) {
	sql = "FROM\n"

	j := g.b.RootJoinTreeItem
	sql += g.iq(NodeTableName(j.Node)) + " AS " + g.iq(j.Alias) + "\n"

	for _, child := range j.ChildReferences {
		sql += g.generateJoinSql(child)
	}
	return
}

func (g *selectGenerator) generateJoinSql(j *JoinTreeItem) (sql string) {
	var tn TableNodeI
	var ok bool

	if tn, ok = j.Node.(TableNodeI); !ok {
		return
	}

	switch node := tn.EmbeddedNode_().(type) {
	case *ReferenceNode:
		sql = "LEFT JOIN "
		sql += g.iq(ReferenceNodeRefTable(node)) + " AS " +
			g.iq(j.Alias) + " ON " + g.iq(j.Parent.Alias) + "." +
			g.iq(ReferenceNodeDbColumnName(node)) + " = " + g.iq(j.Alias) + "." + g.iq(ReferenceNodeRefColumn(node))
		if j.JoinCondition != nil {
			s := g.generateNodeSql(j.JoinCondition, false)
			sql += " AND " + s
		}
	case *ReverseReferenceNode:
		if g.b.LimitInfo != nil && ReverseReferenceNodeIsArray(node) {
			panic("We do not currently support limited queries with an array join.")
		}

		sql = "LEFT JOIN "
		sql += g.iq(ReverseReferenceNodeRefTable(node)) + " AS " +
			g.iq(j.Alias) + " ON " + g.iq(j.Parent.Alias) + "." +
			g.iq(ReverseReferenceNodeKeyColumnName(node)) + " = " + g.iq(j.Alias) + "." + g.iq(ReverseReferenceNodeRefColumn(node))
		if j.JoinCondition != nil {
			s := g.generateNodeSql(j.JoinCondition, false)
			sql += " AND " + s
		}
	case *ManyManyNode:
		if g.b.LimitInfo != nil {
			panic("We do not currently support limited queries with an array join.")
		}

		sql = "LEFT JOIN "

		sql += g.iq(ManyManyNodeDbTable(node)) + " AS " + g.iq(j.Alias+"a") + " ON " +
			g.iq(j.Parent.Alias) + "." +
			g.iq(ColumnNodeDbName(ParentNode(node).(TableNodeI).PrimaryKeyNode())) +
			" = " + g.iq(j.Alias+"a") + "." + g.iq(ManyManyNodeDbColumn(node)) + "\n"
		
		sql += "LEFT JOIN " + g.iq(ManyManyNodeRefTable(node)) + " AS " + g.iq(j.Alias) +
			" ON " + g.iq(j.Alias+"a") + "." + g.iq(ManyManyNodeRefColumn(node)) +
			" = " + g.iq(j.Alias) + "." + g.iq(ManyManyNodeRefPk(node))

		if j.JoinCondition != nil {
			s := g.generateNodeSql(j.JoinCondition, false)
			sql += " AND " + s
		}
	default:
		return
	}
	sql += "\n"
	for _, cj := range j.ChildReferences {
		s := g.generateJoinSql(cj)
		sql += s
	}
	return
}

func (g *selectGenerator) generateWhereSql() (sql string) {
	if g.b.ConditionNode != nil {
		sql = "WHERE "
		var s string
		s = g.generateNodeSql(g.b.ConditionNode, false)
		sql += s + "\n"
	}
	return
}

func (g *selectGenerator) generateGroupBySql() (sql string) {
	if g.b.GroupBys != nil && len(g.b.GroupBys) > 0 {
		sql = "GROUP BY "
		for _, n := range g.b.GroupBys {
			s := g.generateNodeSql(n, true)
			sql += s + ","
		}
		sql = strings.TrimSuffix(sql, ",")
		sql += "\n"
	}
	return
}

// Note that some SQLs (MySQL, SqlLite) allow the use of aliases in a having clause,
// and some (Postgres) do not. We might need to check for this at some point.
func (g *selectGenerator) generateHaving() (sql string) {
	if g.b.HavingNode != nil {
		sql = "HAVING "
		var s string
		s = g.generateNodeSql(g.b.HavingNode, false)
		sql += s + "\n"
	}
	return
}

func (g *selectGenerator) generateLimitSql() (sql string) {
	if g.b.LimitInfo == nil {
		return ""
	}

	if g.b.LimitInfo.MaxRowCount > -1 {
		sql += fmt.Sprintf("LIMIT %d ", g.b.LimitInfo.MaxRowCount)
	}

	if g.b.LimitInfo.Offset > 0 {
		sql += fmt.Sprintf("OFFSET %d ", g.b.LimitInfo.Offset)
	}
	return
}

func (g *selectGenerator) generateOrderBySql() (sql string) {
	if g.b.OrderBys != nil && len(g.b.OrderBys) > 0 {
		sql = "ORDER BY "
		for _, n := range g.b.OrderBys {
			s := g.generateNodeSql(n, true)
			if sorter, ok := n.(NodeSorter); ok {
				if NodeSorterSortDesc(sorter) {
					s += " DESC"
				}
			}
			sql += s + ","
		}
		sql = strings.TrimSuffix(sql, ",")
		sql += "\n"
	}
	return
}

// GenerateUpdate is a helper function for database implementations to generate an update statement.
func GenerateUpdate(db DbI, table string, fields map[string]any, where map[string]any) (sql string, args []any) {
	if len(fields) == 0 {
		panic("No fields to set")
	}

	sql = "UPDATE " + db.QuoteIdentifier(table) + "\nSET "

	var sets []string

	// We range on sorted keys to give SQL optimizers a chance to use a prepared
	// statement by making sure the same fields show up in the same order.
	for k, v := range iter.KeySort(fields) {
		args = append(args, v)
		s := fmt.Sprintf("%s=%s", db.QuoteIdentifier(k), db.FormatArgument(len(args)))
		sets = append(sets, s)
	}

	sql += strings.Join(sets, ", ")

	clear(sets)
	for k, v := range iter.KeySort(where) {
		args = append(args, v)
		s := fmt.Sprintf("%s=%s", db.QuoteIdentifier(k), db.FormatArgument(len(args)))
		sets = append(sets, s)
	}

	sql += "\nWHERE "
	sql += strings.Join(sets, ", ")

	return
}

// GenerateInsert is a helper function for database implementations to generate an insert statement.
func GenerateInsert(db DbI, table string, fields map[string]any) (sql string, args []any) {
	if len(fields) == 0 {
		panic("No fields to insert")
	}

	var keys []string
	var values []string

	// We range on sorted keys to give SQL optimizers a chance to use a prepared
	// statement by making sure the same fields show up in the same order.
	for k, v := range iter.KeySort(fields) {
		keys = append(keys, db.QuoteIdentifier(k))
		args = append(args, v)
		values = append(values, db.FormatArgument(len(args)))
	}

	sql = "INSERT INTO " + db.QuoteIdentifier(table)
	sql += "(" + strings.Join(keys, ",") + ")\nVALUES ("
	sql += strings.Join(values, ",") + ")\n"
	return

}

func GenerateDelete(db DbI, table string, where map[string]any) (sql string, args []any) {
	sql = "DELETE FROM " + db.QuoteIdentifier(table) + "\n"
	sql += "WHERE "

	var s string
	s, args = generateWhereClause(db, where)
	sql += s
	return
}

func GenerateSelect(db DbI, table string, fieldNames []string, where map[string]any, orderBy []string) (sql string, args []any) {
	if len(fieldNames) == 0 {
		panic("No fields to select")
	}
	sql = "SELECT "
	sql += strings.Join(fieldNames, ",\n")
	sql += "\nFROM " + db.QuoteIdentifier(table)
	if len(where) > 0 {
		sql += "\nWHERE "
		var s string
		s, args = generateWhereClause(db, where)
		sql += s
	}

	if len(orderBy) > 0 {
		sql += "\nORDER BY "
		sql += strings.Join(orderBy, ", ")
	}
	return
}

func generateWhereClause(db DbI, where map[string]any) (sql string, args []any) {
	var ors []string
	for kOr, vOr := range iter.KeySort(where) {
		if m, ok := vOr.(map[string]any); ok {
			var ands []string
			for kAnd, vAnd := range m {
				args = append(args, vAnd)
				s := fmt.Sprintf("%s=%s", db.QuoteIdentifier(kAnd), db.FormatArgument(len(args)))
				ands = append(ands, s)
			}
			and := strings.Join(ands, " AND ")
			ors = append(ors, and)
		} else {
			args = append(args, vOr)
			s := fmt.Sprintf("%s=%s", db.QuoteIdentifier(kOr), db.FormatArgument(len(args)))
			ors = append(ors, s)
		}
	}
	sql = strings.Join(ors, " OR ")
	return
}
