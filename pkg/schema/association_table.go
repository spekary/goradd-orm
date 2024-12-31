package schema

import (
	strings2 "github.com/goradd/strings"
	"github.com/kenshaw/snaker"
	"strings"
)

// AssociationTable describes a table in the database that will be used to create a many-to-many
// relationship between two tables.
type AssociationTable struct {
	// Name is the name of the association table in the database. It should be lower_snake_case.
	Name string `json:"name"`
	// For databases that support schemas, this is the name of the schema of the table.
	// Leave blank for the default schema.
	// Databases that do not support schemas will have this prepended to the name of the table.
	Schema string `json:"schema,omitempty"`

	// Table1 is the name of first table in the association.
	// If it has a schema, should be of the form "schema.table"
	Table1 string `json:"table1"`

	// The name of the column in the association table that will be used to point to table 1.
	Column1 string `json:"column1"`

	// Title1 is the singular description that will be used to describe the objects to human readers.
	// This will be used to create the corresponding reference field name in the database and Go identifier.
	// Note that this isn't necessarily the table name. For example a "person" table might be used in a relationship
	// to describe members of a group. The title in that case could be "Member".
	Title1 string `json:"title1,omitempty"`

	// Title1Plural is the plural description that will be used to describe the objects to human readers.
	// This will be used to create the corresponding Go identifier.
	Title1Plural string `json:"title1_plural,omitempty"`

	// Identifier1 is the singular Go name that will be used for one object pointed at.
	Identifier1 string `json:"identifier1,omitempty"`

	// Identifier1Plural is the plural Go name that will be used for the objects pointed at.
	Identifier1Plural string `json:"identifier1_plural,omitempty"`

	// Table2 is the name of second table in the association.
	// If it has a schema, should be of the form "schema.table"
	Table2 string `json:"table2"`

	// The name of the column in the association table that will be used to point to table 2.
	Column2 string `json:"column2"`

	// Title2 is the singular description that will be used to describe the objects to human readers.
	// This will be used to create the corresponding reference field name in the database and Go identifier.
	Title2 string `json:"title2,omitempty"`

	// Title2Plural is the plural description that will be used to describe the objects to human readers.
	// This will be used to create the corresponding Go identifier.
	Title2Plural string `json:"title2_plural,omitempty"`

	// Identifier2 is the singular Go name that will be used for one object pointed at.
	Identifier2 string `json:"identifier2,omitempty"`

	// Identifier2Plural is the plural Go name that will be used for the objects pointed at.
	Identifier2Plural string `json:"identifier2_plural,omitempty"`
}

func (t *AssociationTable) QualifiedName() string {
	if t.Schema == "" {
		return t.Name
	} else {
		return t.Schema + "." + t.Name
	}
}

// FillDefaults will fill default values where none have been set.
func (t *AssociationTable) FillDefaults(referenceSuffix string) {
	col1 := strings.TrimSuffix(t.Column1, referenceSuffix)
	if t.Title1 == "" {
		t.Title1 = strings2.Title(col1)
	}
	if t.Title1Plural == "" {
		t.Title1Plural = strings2.Plural(t.Title1)
	}
	if t.Identifier1 == "" {
		t.Identifier1 = snaker.SnakeToCamelIdentifier(col1)
	}
	if t.Identifier1Plural == "" {
		t.Identifier1Plural = strings2.Plural(t.Identifier1)
	}

	col2 := strings.TrimSuffix(t.Column2, referenceSuffix)
	if t.Title2 == "" {
		t.Title2 = strings2.Title(col2)
	}
	if t.Title2Plural == "" {
		t.Title2Plural = strings2.Plural(t.Title2)
	}
	if t.Identifier2 == "" {
		t.Identifier2 = snaker.SnakeToCamelIdentifier(col2)
	}
	if t.Identifier2Plural == "" {
		t.Identifier2Plural = strings2.Plural(t.Identifier2)
	}
}
