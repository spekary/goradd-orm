package model

import (
	"cmp"
	any2 "github.com/goradd/any"
	strings2 "github.com/goradd/strings"
	"github.com/kenshaw/snaker"
	"log/slog"
	"slices"
	. "spekary/goradd/orm/pkg/query"
	"spekary/goradd/orm/pkg/schema"
)

type ConstVal struct {
	Value int
	Const string
}

// EnumTable describes an table that represents a fixed, enumerated type that will
// not change during the application's use.
type EnumTable struct {
	// DbKey is the key used to find the database in the global database cluster
	DbKey string
	// QueryName is the name of the table to use in querying the database.
	QueryName string
	// Title is the english name of the object when describing it to the world.
	Title string
	// TitlePlural is the plural english name of the object.
	TitlePlural string
	// Identifier is the name of the item as a go type name.
	Identifier string
	// IdentifierPlural is the plural of the go type name.
	IdentifierPlural string
	// DecapIdentifier is the Identifier with the first letter lower case.
	DecapIdentifier string
	// Fields are the names of the fields defined in the table. The first field name MUST be the name of the id field, and 2nd MUST be the name of the name field, the others are optional extra fields.
	Fields []EnumField
	// Values are the go values that will be hardcoded and returned in accessor functions.
	// The map is keyed by row id, and then by field query name
	Values map[int]map[string]any
	// Constants are the constant identifiers that will be used for each enumerated value.
	// These are in ascending order by keys.
	Constants []ConstVal
}

// PkQueryName returns the name of the primary key field as used in database queries.
func (tt *EnumTable) PkQueryName() string {
	return tt.FieldQueryName(0)
}

func (tt *EnumTable) FieldQueryName(i int) string {
	return tt.Fields[i].QueryName
}

func (tt *EnumTable) FieldValue(row int, fieldNum int) any {
	name := tt.FieldQueryName(fieldNum)
	v := tt.Values[row][name]
	if any2.IsNil(v) {
		v = tt.Fields[fieldNum].Type.DefaultValue()
	}
	return v
}

// FieldIdentifier returns the go name corresponding to the given field offset
func (tt *EnumTable) FieldIdentifier(i int) string {
	return tt.Fields[i].Identifier
}

// FieldIdentifierPlural returns the go plural name corresponding to the given field offset
func (tt *EnumTable) FieldIdentifierPlural(i int) string {
	return tt.Fields[i].IdentifierPlural
}

// FieldReceiverType returns the ReceiverType corresponding to the given field offset
func (tt *EnumTable) FieldReceiverType(i int) ReceiverType {
	return tt.Fields[i].Type
}

// FileName returns the default file name corresponding to the enum table.
func (tt *EnumTable) FileName() string {
	return snaker.CamelToSnake(tt.Identifier)
}

// newEnumTable will import the enum table from ets
func newEnumTable(dbKey string, ets *schema.EnumTable) *EnumTable {
	t := &EnumTable{
		DbKey:            dbKey,
		QueryName:        ets.QualifiedName(),
		Title:            ets.Title,
		TitlePlural:      ets.TitlePlural,
		Identifier:       ets.Identifier,
		IdentifierPlural: ets.IdentifierPlural,
		DecapIdentifier:  strings2.Decap(ets.Identifier),
		Values:           make(map[int]map[string]any),
	}
	if len(ets.Values) == 0 {
		slog.Warn("enum table " + t.QueryName + " has no data entries. Specify constants by adding entries to this table.")
	}

	for _, field := range ets.Fields {
		f := EnumField{
			QueryName:        field.Name,
			Title:            field.Title,
			TitlePlural:      field.TitlePlural,
			Identifier:       field.Identifier,
			IdentifierPlural: field.IdentifierPlural,
			Type:             ReceiverTypeFromSchema(field.Type, 0),
		}
		t.Fields = append(t.Fields, f)
	}

	for _, row := range ets.Values {
		valueMap := make(map[string]interface{})
		key, ok := row[0].(int)
		if !ok {
			panic("first column of enum table must be an integer")
		}
		valueMap[t.Fields[0].QueryName] = key
		value, ok := row[1].(string)
		if !ok {
			panic("second column of enum table must be a string")
		}
		valueMap[t.Fields[1].QueryName] = value
		t.Constants = append(t.Constants, ConstVal{key, enumValueToConstant(t.Identifier, value)})
		for i, val := range row[2:] {
			valueMap[t.Fields[i+2].QueryName] = val
		}
		t.Values[key] = valueMap
	}
	slices.SortFunc(t.Constants, func(a, b ConstVal) int {
		return cmp.Compare(a.Value, b.Value)
	})
	return t
}

func enumValueToConstant(prefix string, v string) string {
	v = snaker.ForceCamelIdentifier(v)
	return prefix + v
}

type EnumField struct {
	// QueryName is the name of the field in the database.
	// The name of the first field is typically "id" by convention.
	// The name of the second field must be "name".
	// The name of the following fields is up to you, but should be lower_snake_case.
	QueryName string
	// Title is the title of the data stored in the field.
	Title string
	// TitlePlural is the plural form of the Title.
	TitlePlural string
	// Identifier is the name used in Go code to access the data.
	Identifier string
	// IdentifierPlural is the plural form of Identifier.
	IdentifierPlural string
	// Type is the type of the column.
	// The first column must be type ColTypeInt.
	// The second column must be type ColTypeString.
	// Other columns can be one of the other types, but not ColTypeReference.
	Type ReceiverType
}
