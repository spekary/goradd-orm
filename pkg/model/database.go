package model

import (
	"fmt"
	"github.com/goradd/all"
	"github.com/goradd/maps"
	"github.com/goradd/orm/pkg/query"
	"github.com/goradd/orm/pkg/schema"
	strings2 "github.com/goradd/strings"
	"github.com/kenshaw/snaker"
	"log/slog"
	maps2 "maps"
	"slices"
	"strings"
)

type Model []*Database

func FromSchemas(schemas []*schema.Database) (dbs Model) {
	for _, s := range schemas {
		s.FillDefaults()
		db := FromSchema(s)
		dbs = append(dbs, db)
	}
	return
}

func FromSchema(s *schema.Database) *Database {
	d := Database{
		Key:             s.Key,
		ReferenceSuffix: s.ReferenceSuffix,
		EnumTableSuffix: s.EnumTableSuffix,
		AssnTableSuffix: s.AssnTableSuffix,
	}
	d.importSchema(s)
	return &d
}

// Database is the top level struct that contains a description of a database modeled as objects.
// It is used in code generation and query creation.
type Database struct {
	// The database key corresponding to its key in the global database cluster
	Key string
	// Tables are the tables in the database, keyed by database table name
	Tables map[string]*Table
	// Enums contains a description of the enumerated types linked to the database, keyed by database table name
	Enums map[string]*Enum

	// ReferenceSuffix is the text to strip off the end of foreign key references when converting to names.
	// Defaults to "_id"
	ReferenceSuffix string
	// EnumTableSuffix is the text to string off the end of an enum table when converting it to a type name.
	// Defaults to "_enum".
	EnumTableSuffix string
	// Defaults to _assn
	AssnTableSuffix string
}

// importSchema will convert a database description to a model which generally treats
// tables as objects and columns as member variables.
func (m *Database) importSchema(schema *schema.Database) {
	m.Enums = make(map[string]*Enum)
	m.Tables = make(map[string]*Table)

	// deal with enum tables first
	for _, et := range schema.EnumTables {
		tt := newEnumTable(m.Key, et)
		m.Enums[tt.QueryName] = tt
	}

	// get the regular tables
	for _, table := range schema.Tables {
		t := newTable(m.Key, table)
		m.Tables[t.QueryName] = t
	}

	// import references after the columns are in place
	for _, table := range schema.Tables {
		if t := m.Table(table.Name); t != nil {
			m.importReferences(t, table)
		}
	}

	for _, assn := range schema.AssociationTables {
		m.importAssociation(assn)
	}
}

// Analyzes an association table and creates special virtual columns in the corresponding tables it points to.
// Association tables are used by SQL databases to create many-many relationships. NoSQL databases can define their
// association columns directly and store an array of records on either end of the association.
func (m *Database) importAssociation(schemaAssn *schema.AssociationTable) {
	e1 := m.Enum(schemaAssn.Table1)
	e2 := m.Enum(schemaAssn.Table2)
	t1 := m.Table(schemaAssn.Table1)
	t2 := m.Table(schemaAssn.Table2)

	if t1 != nil && t2 != nil {
		ref1 := makeManyManyRef(schemaAssn.Name, schemaAssn.Column1, schemaAssn.Column2, t1, t2, schemaAssn.Title2, schemaAssn.Title2Plural, schemaAssn.Identifier2, schemaAssn.Identifier2Plural)
		ref2 := makeManyManyRef(schemaAssn.Name, schemaAssn.Column2, schemaAssn.Column1, t2, t1, schemaAssn.Table1, schemaAssn.Title1Plural, schemaAssn.Identifier1, schemaAssn.Identifier1Plural)
		ref1.MM = ref2
		ref2.MM = ref1
	} else if e1 != nil {
		slog.Warn(fmt.Sprintf("Skipped association table %s: Table1 cannot be an enum table.", schemaAssn.Name))
		return
	} else if e2 != nil && t1 != nil {
		makeManyManyEnumRef(schemaAssn.Name, schemaAssn.Column1, schemaAssn.Column2, t1, e2, schemaAssn.Title2, schemaAssn.Title2Plural, schemaAssn.Identifier2, schemaAssn.Identifier2Plural)
	} else {
		slog.Warn(fmt.Sprintf("Skipped association table %s: missing associated table.", schemaAssn.Name))
		return
	}
}

func (m *Database) importReferences(t *Table, schemaTable *schema.Table) {
	for _, col := range schemaTable.Columns {
		m.importReference(t, col)
	}
}

func (m *Database) importReference(table *Table, schemaCol *schema.Column) {
	if schemaCol.Reference != nil {
		refTable := m.Table(schemaCol.Reference.Table)
		et := m.Enum(schemaCol.Reference.Table)
		if refTable == nil && et == nil {
			slog.Error(fmt.Sprintf("Reference skipped, table not found. From %s:%s", table.QueryName, schemaCol.Name))
			return
		}
		f := &Reference{
			Table:                   refTable,
			EnumTable:               et,
			Identifier:              schemaCol.Reference.Identifier,
			Title:                   schemaCol.Reference.Title,
			ReverseTitle:            schemaCol.Reference.ReverseTitle,
			ReverseTitlePlural:      schemaCol.Reference.ReverseTitlePlural,
			ReverseIdentifier:       schemaCol.Reference.ReverseIdentifier,
			ReverseIdentifierPlural: schemaCol.Reference.ReverseIdentifierPlural,
		}
		f.DecapIdentifier = strings2.Decap(f.Identifier)

		var thisCol *Column
		thisCol = table.ColumnByName(schemaCol.Name)
		if thisCol == nil {
			// This should not happen
			slog.Error(fmt.Sprintf("Reference skipped, column not found. From %s:%s", refTable.QueryName, schemaCol.Name))
			return
		}
		thisCol.Reference = f
		if refTable != nil {
			refTable.ReverseReferences = append(refTable.ReverseReferences, thisCol)
		} else {
			// enum table
			thisCol.Type = query.ColTypeInteger // In case this came through unsigned, fix it
		}
	}
}

// Table returns a Table from the database given the table name.
func (m *Database) Table(name string) *Table {
	if v, ok := m.Tables[name]; ok {
		return v
	} else {
		return nil
	}
}

// Enum returns an Enum from the database given the table name.
func (m *Database) Enum(name string) *Enum {
	return m.Enums[name]
}

// IsEnumTable returns true if the given name is the name of an enum table in the database
func (m *Database) IsEnumTable(name string) bool {
	_, ok := m.Enums[name]
	return ok
}

func isReservedIdentifier(s string) bool {
	switch s {
	case "break":
		return true
	case "case":
		return true
	case "chan":
		return true
	case "const":
		return true
	case "continue":
		return true
	case "default":
		return true
	case "defer":
		return true
	case "else":
		return true
	case "fallthrough":
		return true
	case "for":
		return true
	case "func":
		return true
	case "go":
		return true
	case "goto":
		return true
	case "if":
		return true
	case "import":
		return true
	case "interface":
		return true
	case "map":
		return true
	case "package":
		return true
	case "range":
		return true
	case "return":
		return true
	case "select":
		return true
	case "struct":
		return true
	case "switch":
		return true
	case "type":
		return true
	case "var":
		return true
	}
	return false
}

func LowerCaseIdentifier(s string) (i string) {
	if strings.Contains(s, "_") {
		i = snaker.ForceLowerCamelIdentifier(snaker.SnakeToCamelIdentifier(s))
	} else {
		// Not a snake string, but still might need some fixing up
		i = snaker.ForceLowerCamelIdentifier(s)
	}
	i = strings.TrimSpace(i)
	if isReservedIdentifier(i) {
		panic("Cannot use '" + i + "' as an identifier.")
	}
	if i == "" {
		panic("Cannot use blank as an identifier.")
	}
	return
}

func UpperCaseIdentifier(s string) (i string) {
	if strings.Contains(s, "_") {
		i = snaker.ForceCamelIdentifier(snaker.SnakeToCamelIdentifier(s))
	} else {
		// Not a snake string, but still might need some fixing up
		i = snaker.ForceCamelIdentifier(s)
	}
	i = strings.TrimSpace(i)
	if i == "" {
		panic("Cannot use blank as an identifier.")
	}
	return
}

// MarshalOrder returns an array of tables in the order they should be marshaled such that
// if they then get unmarshalled in the same order, there will not be problems with foreign
// keys not existing when the object is saved.
func (m *Database) MarshalOrder() (tables []*Table) {
	var unusedTables maps.Set[*Table]

	unusedTables.Add(slices.Collect(maps2.Values(m.Tables))...)
	// First add the tables that have no forward references
	for { // repeat until unusedTables is empty
		var newTables []*Table
	nexttable:
		for t := range unusedTables.All() {
			for _, col := range t.Columns {
				if col.IsReference() {
					if !slices.Contains(tables, col.Reference.Table) &&
						!slices.Contains(newTables, col.Reference.Table) {
						continue nexttable
					}
				}
			}
			// This has no forward references we care about
			newTables = append(newTables, t)
		}
		slices.SortFunc(newTables, func(a, b *Table) int {
			if a.QueryName < b.QueryName {
				return -1
			} else {
				return all.If(a.QueryName > b.QueryName, 1, 0)
			}
		})
		tables = append(tables, newTables...)
		// Remove the found tables
		for _, t := range newTables {
			unusedTables.Delete(t)
		}
		if unusedTables.Len() == 0 {
			break
		}
	}

	return
}
