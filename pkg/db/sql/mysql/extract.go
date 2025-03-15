package mysql

import (
	"database/sql"
	"fmt"
	"github.com/goradd/iter"
	"github.com/goradd/maps"
	"github.com/goradd/orm/pkg/db"
	sql2 "github.com/goradd/orm/pkg/db/sql"
	. "github.com/goradd/orm/pkg/query"
	"github.com/goradd/orm/pkg/schema"
	strings2 "github.com/goradd/strings"
	"log"
	log2 "log/slog"
	"math"
	"slices"
	"sort"
	"strings"
)

/*
This file contains the code that parses the data structure found in a MySQL database into
our own cross-platform internal database description object.
*/

type mysqlTable struct {
	name                string
	columns             []mysqlColumn
	indexes             []mysqlIndex
	fkMap               map[string]mysqlForeignKey
	comment             string
	options             map[string]interface{}
	supportsForeignKeys bool
}

type mysqlColumn struct {
	name            string
	defaultValue    sql2.SqlReceiver
	isNullable      string
	dataType        string
	dataLen         int
	characterMaxLen sql.NullInt64
	columnType      string
	collation       sql.NullString
	key             string
	extra           string
	comment         string
	options         map[string]any
}

type mysqlIndex struct {
	name       string
	nonUnique  bool
	tableName  string
	columnName string
}

type mysqlForeignKey struct {
	constraintName            string
	tableName                 string
	columnName                string
	referencedTableName       sql.NullString
	referencedColumnName      sql.NullString
	referencedColumnIndexName sql.NullString
	updateRule                sql.NullString
	deleteRule                sql.NullString
}

func (m *DB) ExtractSchema(options map[string]any) schema.Database {
	rawTables := m.getRawTables()
	return m.schemaFromRawTables(rawTables, options)
}

func (m *DB) getRawTables() map[string]mysqlTable {
	var tableMap = make(map[string]mysqlTable)

	indexes, err := m.getIndexes()
	if err != nil {
		return nil
	}

	foreignKeys, err := m.getForeignKeys()
	if err != nil {
		return nil
	}

	tables := m.getTables()
	for _, table := range tables {
		// Do some processing on the foreign keys
		for _, fk := range foreignKeys[table.name] {
			if fk.referencedColumnName.Valid && fk.referencedTableName.Valid {
				if _, ok := table.fkMap[fk.columnName]; ok {
					log2.Warn(fmt.Sprintf("Column %s:%s multi-column foreign keys are not supported.", table.name, fk.columnName))
					delete(table.fkMap, fk.columnName)
				} else {
					table.fkMap[fk.columnName] = fk
				}
			}
		}

		columns, err2 := m.getColumns(table.name)
		if err2 != nil {
			return nil
		}

		table.indexes = indexes[table.name]
		table.columns = columns
		tableMap[table.name] = table
	}

	return tableMap

}

// Gets information for a table
func (m *DB) getTables() []mysqlTable {
	var tableName, tableComment, tableEngine string
	var tables []mysqlTable

	// Use the MySQL5 Information Schema to get a list of all the tables in this database
	// (excluding views, etc.)
	dbName := m.databaseName

	rows, err := m.SqlDb().Query(fmt.Sprintf(`
	SELECT
	table_name,
	table_comment,
	engine
	FROM
	information_schema.tables
	WHERE
	table_type <> 'VIEW' AND
	table_schema = '%s';
	`, dbName))

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&tableName, &tableComment, &tableEngine)
		var supportsForeignKeys bool
		if tableEngine == "InnoDB" {
			supportsForeignKeys = true
		}
		if err != nil {
			log.Fatal(err)
		}
		log2.Info("Importing schema for table " + tableName)
		table := mysqlTable{
			name:                tableName,
			comment:             tableComment,
			columns:             []mysqlColumn{},
			fkMap:               make(map[string]mysqlForeignKey),
			indexes:             []mysqlIndex{},
			supportsForeignKeys: supportsForeignKeys,
		}
		if table.options, table.comment, err = sql2.ExtractOptions(table.comment); err != nil {
			log2.Warn("Error in comment options for table " + table.name + " - " + err.Error())
		}

		tables = append(tables, table)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return tables
}

func (m *DB) getColumns(table string) (columns []mysqlColumn, err error) {
	dbName := m.databaseName

	rows, err := m.SqlDb().Query(fmt.Sprintf(`
	SELECT
	column_name,
	column_default,
	is_nullable,
	data_type,
	character_maximum_length,
	column_type,
	column_key,
	extra,
	column_comment,
	collation_name
	FROM
	information_schema.columns
	WHERE
	table_name = '%s' AND
	table_schema = '%s'
	ORDER BY
	ordinal_position;
	`, table, dbName))

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var col mysqlColumn

	for rows.Next() {
		col = mysqlColumn{}
		err = rows.Scan(&col.name, &col.defaultValue.R, &col.isNullable, &col.dataType, &col.characterMaxLen, &col.columnType, &col.key, &col.extra, &col.comment, &col.collation)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		if col.options, col.comment, err = sql2.ExtractOptions(col.comment); err != nil {
			log2.Warn("Error in table comment options for table " + table + ":" + col.name + " - " + err.Error())
		}
		columns = append(columns, col)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return columns, err
}

func (m *DB) getIndexes() (indexes map[string][]mysqlIndex, err error) {

	dbName := m.databaseName
	indexes = make(map[string][]mysqlIndex)

	rows, err := m.SqlDb().Query(fmt.Sprintf(`
	SELECT
	index_name,
	non_unique,
	table_name,
	column_name
	FROM
	information_schema.statistics
	WHERE
	table_schema = '%s';
	`, dbName))

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var index mysqlIndex

	for rows.Next() {
		index = mysqlIndex{}
		err = rows.Scan(&index.name, &index.nonUnique, &index.tableName, &index.columnName)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		tableIndexes := indexes[index.tableName]
		tableIndexes = append(tableIndexes, index)
		indexes[index.tableName] = tableIndexes
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return indexes, err
}

// getForeignKeys gets information on the foreign keys.
//
// Note that querying the information_schema database is SLOW, so we want to do it as few times as possible.
func (m *DB) getForeignKeys() (foreignKeys map[string][]mysqlForeignKey, err error) {
	dbName := m.databaseName
	fkMap := make(map[string]mysqlForeignKey)

	rows, err := m.SqlDb().Query(fmt.Sprintf(`
SELECT
    t1.constraint_name,
    t1.table_name,
    t1.column_name,
    t1.referenced_table_name,
    t1.referenced_column_name,
    t2.index_name
FROM
    information_schema.key_column_usage AS t1
JOIN
    information_schema.statistics AS t2
    ON t2.table_name = t1.referenced_table_name
    AND t2.column_name = t1.referenced_column_name
WHERE
    t1.constraint_schema = '%s'
ORDER BY
    t1.ordinal_position
`, dbName))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		fk := mysqlForeignKey{}
		err = rows.Scan(&fk.constraintName, &fk.tableName, &fk.columnName, &fk.referencedTableName, &fk.referencedColumnName, &fk.referencedColumnIndexName)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		if fk.referencedColumnName.Valid {
			fkMap[fk.constraintName] = fk
		}
	}

	rows.Close()

	rows, err = m.SqlDb().Query(fmt.Sprintf(`
	SELECT
	constraint_name,
	update_rule,
	delete_rule
	FROM
	information_schema.referential_constraints
	WHERE
	constraint_schema = '%s';
	`, dbName))
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var constraintName string
		var updateRule, deleteRule sql.NullString
		err = rows.Scan(&constraintName, &updateRule, &deleteRule)
		if err != nil {
			log.Fatal(err)
		}
		if fk, ok := fkMap[constraintName]; ok {
			fk.updateRule = updateRule
			fk.deleteRule = deleteRule
			fkMap[constraintName] = fk
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	foreignKeys = make(map[string][]mysqlForeignKey)
	for _, fk := range iter.KeySort(fkMap) {
		tableKeys := foreignKeys[fk.tableName]
		tableKeys = append(tableKeys, fk)
		foreignKeys[fk.tableName] = tableKeys
	}
	return foreignKeys, err
}

// Convert the database native type to a more generic sql type, and a go table type.
func processTypeInfo(column mysqlColumn) (
	typ schema.ColumnType,
	subType schema.ColumnSubType,
	maxLength uint64,
	defaultValue interface{},
	extra map[string]interface{},
	err error) {
	dataLen := sql2.GetDataDefLength(column.columnType)
	isUnsigned := strings.Contains(column.columnType, "unsigned")

	switch column.dataType {
	case "time":
		typ = schema.ColTypeTime
		subType = schema.ColSubTypeTimeOnly
	case "date":
		typ = schema.ColTypeTime
		subType = schema.ColSubTypeDateOnly
	case "timestamp":
		typ = schema.ColTypeTime
	case "datetime":
		typ = schema.ColTypeTime
	case "tinyint":
		if dataLen == 1 {
			typ = schema.ColTypeBool
		} else {
			maxLength = 8
			if isUnsigned {
				typ = schema.ColTypeUint
			} else {
				typ = schema.ColTypeInt
			}
		}

	case "int":
		maxLength = 32 // mysql standard int has a 32-bit limit even in 64-bit implementations
		if isUnsigned {
			typ = schema.ColTypeUint
		} else {
			typ = schema.ColTypeInt
		}

	case "smallint":
		maxLength = 16
		if isUnsigned {
			typ = schema.ColTypeUint
		} else {
			typ = schema.ColTypeInt
		}

	case "mediumint":
		maxLength = 24
		if isUnsigned {
			typ = schema.ColTypeUint
		} else {
			typ = schema.ColTypeInt
		}

	case "bigint":
		maxLength = 64
		if isUnsigned {
			typ = schema.ColTypeUint
		} else {
			typ = schema.ColTypeInt
			if column.name == schema.GroTimestampColumnName {
				subType = schema.ColSubTypeTimestamp
			} else if column.name == schema.GroLockColumnName {
				subType = schema.ColSubTypeLock
			}
		}

	case "float":
		typ = schema.ColTypeFloat
		maxLength = 32
	case "double":
		typ = schema.ColTypeFloat
		maxLength = 64

	case "varchar":
		fallthrough
	case "char":
		typ = schema.ColTypeString
		maxLength = uint64(dataLen)

	case "blob":
		typ = schema.ColTypeBytes
		maxLength = 65535
	case "tinyblob":
		typ = schema.ColTypeBytes
		maxLength = 255
	case "mediumblob":
		typ = schema.ColTypeBytes
		maxLength = 16777215
	case "longblob":
		typ = schema.ColTypeBytes
		maxLength = math.MaxUint32

	case "text":
		typ = schema.ColTypeString
		maxLength = 65535
	case "tinytext":
		typ = schema.ColTypeString
		maxLength = 255
	case "mediumtext":
		typ = schema.ColTypeString
		maxLength = 16777215
	case "longtext":
		typ = schema.ColTypeString
		maxLength = math.MaxUint32

	case "decimal":
		// No native equivalent in Go.
		// See the shopspring/decimal package for possible support.
		// You will need to shepherd numbers into and out of []byte format to move data to the database.
		typ = schema.ColTypeUnknown
		maxLength = uint64(dataLen) + 3
		extra = map[string]interface{}{"type": column.columnType}

	case "year":
		typ = schema.ColTypeInt
		extra = map[string]interface{}{"type": column.columnType}

	case "set":
		typ = schema.ColTypeUnknown
		maxLength = uint64(column.characterMaxLen.Int64)
		extra = map[string]interface{}{"type": column.columnType}

	case "enum":
		typ = schema.ColTypeUnknown
		maxLength = uint64(column.characterMaxLen.Int64)
		extra = map[string]interface{}{"type": column.columnType}

	default:
		typ = schema.ColTypeUnknown
		extra = map[string]interface{}{"type": column.columnType}
	}

	defaultValue = column.defaultValue.UnpackDefaultValue(typ, int(maxLength))
	if typ == schema.ColTypeTime && strings.Contains(strings.ToUpper(column.extra), "ON UPDATE") {
		defaultValue = "update"
	}
	return
}

func (m *DB) schemaFromRawTables(rawTables map[string]mysqlTable, options map[string]any) schema.Database {

	dd := schema.Database{
		ReferenceSuffix: options["reference_suffix"].(string),
		EnumTableSuffix: options["enum_table_suffix"].(string),
		AssnTableSuffix: options["assn_table_suffix"].(string),
		Key:             m.DbKey(),
	}

	for tableName, rawTable := range iter.KeySort(rawTables) {
		if rawTable.options["skip"] != nil {
			continue
		}

		if strings2.EndsWith(tableName, dd.EnumTableSuffix) {
			if t, err := m.getEnumTableSchema(rawTable); err != nil {
				log2.Error("Enum rawTable " + tableName + " skipped: " + err.Error())
			} else {
				dd.EnumTables = append(dd.EnumTables, &t)
			}
		} else if strings2.EndsWith(tableName, dd.AssnTableSuffix) {
			if mm, err := m.getAssociationSchema(rawTable, dd.EnumTableSuffix); err != nil {
				log2.Error("Association rawTable " + tableName + " skipped: " + err.Error())
			} else {
				dd.AssociationTables = append(dd.AssociationTables, &mm)
			}
		} else {
			t := m.getTableSchema(rawTable, dd.EnumTableSuffix)
			dd.Tables = append(dd.Tables, &t)
		}
	}
	return dd
}

func (m *DB) getTableSchema(t mysqlTable, enumTableSuffix string) schema.Table {
	var columnSchemas []*schema.Column

	// Build the indexes
	indexes := make(map[string]*schema.MultiColumnIndex)
	pkColumns := maps.NewSet[string]()
	uniqueColumns := maps.NewSet[string]()
	singleIndexes := maps.NewSet[string]()

	// Fill pkColumns map with the column names of all the pk columns
	// Also fill the indexes map with a list of columns for each index keyed by index name
	for _, idx := range t.indexes {
		if idx.name == "PRIMARY" {
			pkColumns.Add(idx.columnName)
		} else if i, ok2 := indexes[idx.name]; ok2 {
			i.Columns = append(i.Columns, idx.columnName)
			sort.Strings(i.Columns) // make sure this list stays in a predictable order each time
		} else {
			i = &schema.MultiColumnIndex{IsUnique: !idx.nonUnique, Columns: []string{idx.columnName}}
			indexes[idx.name] = i
		}
	}

	// Fill the uniqueColumns set with all the columns that have a single unique index,
	// including any PK columns. Single indexes are used to determine 1 to 1 relationships.
	// Also fill the singleIndexes set with columns that have a single index.
	for _, idx := range indexes {
		if len(idx.Columns) == 1 {
			singleIndexes.Add(idx.Columns[0])
			if idx.IsUnique {
				uniqueColumns.Add(idx.Columns[0])
			}
		}
	}

	// If there is only one primary key column in the table, add it to the unique columns set
	if pkColumns.Len() == 1 {
		uniqueColumns.Add(pkColumns.Values()[0])
	}

	var pkCount int
	for _, col := range t.columns {
		cd := m.getColumnSchema(t, col, singleIndexes.Has(col.name), pkColumns.Has(col.name), uniqueColumns.Has(col.name), enumTableSuffix)

		if cd.Type == schema.ColTypeAutoPrimaryKey || cd.IndexLevel == schema.IndexLevelManualPrimaryKey {
			// private keys go first
			columnSchemas = slices.Insert(columnSchemas, pkCount, &cd)
			pkCount++
		} else {
			columnSchemas = append(columnSchemas, &cd)
		}
	}

	schem := ""
	tableName := t.name
	parts := strings.Split(t.name, ".")
	if len(parts) == 2 {
		schem = parts[0]
		tableName = parts[1]
	}
	td := schema.Table{
		Name:    tableName,
		Schema:  schem,
		Columns: columnSchemas,
		Key:     t.name,
	}

	// Create the multi-column index array
	for _, idx := range indexes {
		if len(idx.Columns) > 1 {
			slices.Sort(idx.Columns)
			td.MultiColumnIndexes = append(td.MultiColumnIndexes, *idx)
		}
	}
	if pkColumns.Len() == 2 {
		mc := schema.MultiColumnIndex{
			IsUnique: true,
			Columns:  pkColumns.Values(),
		}
		slices.Sort(mc.Columns)
		td.MultiColumnIndexes = append(td.MultiColumnIndexes, mc)
	}

	// Keep the MultiColumnIndexes in a predictable order
	slices.SortFunc(td.MultiColumnIndexes, func(m1 schema.MultiColumnIndex, m2 schema.MultiColumnIndex) int {
		return slices.Compare(m1.Columns, m2.Columns)
	})

	return td
}

func (m *DB) getEnumTableSchema(t mysqlTable) (ed schema.EnumTable, err error) {
	td := m.getTableSchema(t, "")

	var columnNames []string
	var receiverTypes []ReceiverType

	if td.Columns[0].Type == schema.ColTypeAutoPrimaryKey ||
		(td.Columns[0].IndexLevel == schema.IndexLevelManualPrimaryKey &&
			(td.Columns[0].Type == schema.ColTypeInt || td.Columns[0].Type == schema.ColTypeUint)) {
	} else {
		err = fmt.Errorf("error: An enum table must have a single primary key that is an integer column")
		return
	}

	if td.Columns[1].Type == schema.ColTypeAutoPrimaryKey ||
		td.Columns[1].IndexLevel == schema.IndexLevelManualPrimaryKey {
		err = fmt.Errorf("error: An enum table cannot have more than one primary key column")
		return
	}

	ed.Name = td.Name

	for i, c := range td.Columns {
		if c.Type == schema.ColTypeReference {
			err = fmt.Errorf("cannot have a reference column in an enum table")
			return
		}
		columnNames = append(columnNames, c.Name)
		recType := ReceiverTypeFromSchema(c.Type, c.Size)
		typ := c.Type
		if i == 0 {
			recType = ColTypeInteger // Force first value to be treated like an integer
			typ = schema.ColTypeInt
		}
		if c.Type == schema.ColTypeUnknown {
			recType = ColTypeBytes
			typ = schema.ColTypeBytes
		}

		receiverTypes = append(receiverTypes, recType)
		ft := schema.EnumField{
			Name: c.Name,
			Type: typ,
		}
		if ed.Name == "name" {
			if len(ed.Fields) == 0 {
				panic("1st field should be the id primary key field")
			}
			slices.Insert(ed.Fields, 1, &ft)
		} else {
			ed.Fields = append(ed.Fields, &ft)
		}
	}

	var result *sql.Rows
	result, err = m.SqlDb().Query(`
	SELECT ` +
		"`" + strings.Join(columnNames, "`,`") + "`" +
		`
	FROM ` +
		"`" + td.Name + "`" +
		` ORDER BY ` + "`" + columnNames[0] + "`")

	if err != nil {
		return
	}

	receiver := sql2.ReceiveRows(result, receiverTypes, columnNames, nil)
	for _, row := range receiver {
		var values []interface{}
		for _, field := range ed.Fields {
			values = append(values, row[field.Name])
		}
		ed.Values = append(ed.Values, values)
	}
	return
}

func (m *DB) getColumnSchema(table mysqlTable,
	column mysqlColumn,
	isIndexed bool,
	isPk bool,
	isUnique bool,
	enumTableSuffix string) schema.Column {

	cd := schema.Column{
		Name: column.name,
	}
	var err error
	var extra map[string]any
	cd.Type, cd.SubType, cd.Size, cd.DefaultValue, extra, err = processTypeInfo(column)
	if err != nil {
		log2.Warn(err.Error() + ". Table = " + table.name + "; Column = " + column.name)
	}
	if extra != nil {
		if cd.DatabaseColumnInfo == nil {
			cd.DatabaseColumnInfo = make(map[string]map[string]interface{})
		}
		cd.DatabaseColumnInfo[db.DriverTypeMysql] = extra
	}

	isAuto := strings.Contains(column.extra, "auto_increment")
	if isAuto {
		cd.Type = schema.ColTypeAutoPrimaryKey
		cd.Size = 0
	} else if isPk {
		cd.IndexLevel = schema.IndexLevelManualPrimaryKey
	} else if isUnique {
		cd.IndexLevel = schema.IndexLevelUnique
	} else if isIndexed {
		cd.IndexLevel = schema.IndexLevelIndexed
	}

	cd.IsNullable = column.isNullable == "YES"

	if enumTableSuffix != "" && strings.HasSuffix(column.name, enumTableSuffix) { // handle enum type
		var refTable string
		if v, ok := column.options["enum_table"]; ok {
			if s, ok2 := v.(string); ok2 {
				refTable = s
			}
		} // otherwise we will infer the reftable from the name of the column later

		if cd.Type == schema.ColTypeInt || cd.Type == schema.ColTypeUint {
			cd.Type = schema.ColTypeEnum
			if fk, ok := table.fkMap[column.name]; ok { // if it is a foreign key to the enum table, use that
				cd.Reference = &schema.Reference{
					Table: fk.referencedTableName.String,
				}
			} else if refTable != "" {
				cd.Reference = &schema.Reference{
					Table: refTable,
				}
			}
		} else if cd.Type == schema.ColTypeString {
			cd.Type = schema.ColTypeEnumArray
			if refTable != "" {
				cd.Reference = &schema.Reference{
					Table: refTable,
				}
			}
		}
	} else if fk, ok2 := table.fkMap[column.name]; ok2 { // handle forward reference
		if fk.referencedColumnIndexName.String != "PRIMARY" {
			log2.Warn(fmt.Sprintf("Foregin key for %s:%s appears to not be pointing to a primary key. Only primary key foreign keys are supported.", table.name, column.name))
		}
		cd.Reference = &schema.Reference{
			Table: fk.referencedTableName.String,
		}
		cd.Type = schema.ColTypeReference
		cd.Size = 0
	}

	if strings.HasSuffix(column.collation.String, "_ci") {
		cd.CaseInsensitive = true
	}

	return cd
}

func (m *DB) getAssociationSchema(t mysqlTable, enumTableSuffix string) (mm schema.AssociationTable, err error) {
	td := m.getTableSchema(t, enumTableSuffix)
	if len(td.Columns) != 2 {
		err = fmt.Errorf("association table must have only 2 columns")
		return
	}
	if len(td.MultiColumnIndexes) != 1 ||
		!td.MultiColumnIndexes[0].IsUnique {
		err = fmt.Errorf("association table must have one multi-column index that is unique")
	}
	var typeIndex = -1
	for i, cd := range td.Columns {
		if cd.Reference == nil {
			err = fmt.Errorf("column " + cd.Name + " must be a foreign key.")
			return
		}

		if cd.IsNullable {
			err = fmt.Errorf("column " + cd.Name + " cannot be nullable.")
			return
		}

		if cd.Type == schema.ColTypeEnum {
			if typeIndex != -1 {
				err = fmt.Errorf("column " + cd.Name + " cannot have two foreign keys to enum tables.")
				return
			}
			typeIndex = i
		}
	}
	mm.Name = td.Name
	mm.Table1 = td.Columns[0].Reference.Table
	mm.Column1 = td.Columns[0].Name
	mm.Table2 = td.Columns[1].Reference.Table
	mm.Column2 = td.Columns[1].Name
	return
}
