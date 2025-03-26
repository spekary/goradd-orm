package model

import (
	"fmt"
	. "github.com/goradd/orm/pkg/query"
	"github.com/goradd/orm/pkg/schema"
	strings2 "github.com/goradd/strings"
	"time"
)

const (
	CreatedTime  = "now"
	ModifiedTime = "update"
)

// Column describes a database column. Most of the information is either
// gleaned from the structure of the database, or is taken from a file that describes the relationships between
// different record types. Some information is filled in after analysis. Some information can be
// provided through information embedded in database comments.
type Column struct {
	// QueryName is the name of the column in the database.
	QueryName string
	// Table is a pointer back to the table that this column is part of
	Table *Table
	// Identifier is the name of the column in go code.
	Identifier string
	// DecapIdentifier is a cache for the lower case identifier for the column.
	DecapIdentifier string
	// SchemaType is the type info specified in the schema description
	SchemaType schema.ColumnType
	// SchemaSubType further describes the schema type.
	SchemaSubType schema.ColumnSubType
	//  ReceiverType indicates the Go type that matches the column.
	ReceiverType ReceiverType
	// Size is the maximum length of runes to allow in the column if a string type column.
	// If a byte array, it is the number of bytes permitted.
	// If an attempt at entering more than this amount occurs, it is considered a programming bug
	// and we will panic.
	// If an integer or float type, it is the number of bits in the data type.
	Size uint64
	// DefaultValue is the default value as specified by the database. We will initialize new ORM objects
	// with this value.
	DefaultValue interface{}
	// IsAutoId is true if this column represents a unique identifier automatically generated by the database or database driver.
	IsAutoId bool
	// IsPrimaryKey is true if this is the primary key column.
	IsPrimaryKey bool
	// IsNullable is true if the column can be given a NULL value.
	IsNullable bool
	// IsUnique is true if the column's table has a single unique index on the column.
	IsUnique bool
	// Reference is additional information describing a foreign key relationship or enum table
	Reference *Reference
	// Options are the options extracted from the comments string
	Options map[string]interface{}
}

// DefaultConstantName returns the name of the default value constant that will be used to refer to the default value
func (cd *Column) DefaultConstantName() string {
	title := cd.Table.Identifier + cd.Identifier + "Default"
	return title
}

func (cd *Column) VariableIdentifier() string {
	return cd.DecapIdentifier
}

func (cd *Column) VariableIdentifierPlural() string {
	return strings2.Plural(cd.DecapIdentifier)
}

// DefaultValueAsValue returns the default value of the column as a GO value
func (cd *Column) DefaultValueAsValue() string {
	if cd.DefaultValue == nil {
		if cd.IsAutoId || cd.IsReference() {
			return `""`
		} else if cd.IsEnum() {
			if cd.IsEnumArray() {
				return "nil"
			}
			return "0"
		}
		return cd.ReceiverType.DefaultValueString()
	}

	if cd.ReceiverType == ColTypeTime {
		if cd.DefaultValue == CreatedTime || cd.DefaultValue == ModifiedTime {
			return "time.Time{}" // These times will be updated when the object is saved.
		} else {
			t := cd.DefaultValue.(time.Time)
			if t.IsZero() {
				return "time.Time{}"
			}
			return fmt.Sprintf("time2.NewDateTime(%d, %d, %d, %d, %d, %d, %d)", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
		}
	}

	if cd.IsEnum() {
		if cd.IsEnumArray() {
			// not supported yet
		} else {
			return fmt.Sprintf("%s(%d)", cd.ReferenceType(), cd.DefaultValue) // should be casting an int to an enum type
		}
	}

	return fmt.Sprintf("%#v", cd.DefaultValue)
}

/*
// DefaultValueAsConstant returns the default value of the column as a Go constant
func (cd *Column) DefaultValueAsConstant() string {
	if cd.ReceiverType == ColTypeTime {
		if cd.DefaultValue == CreatedTime || cd.DefaultValue == ModifiedTime {
			return `time2.Current`
		} else if cd.DefaultValue == nil {
			return `time2.Zero`
		} else {
			d := cd.DefaultValue.(time.Time)
			if b, _ := d.MarshalText(); b == nil {
				return `time2.Zero`
			} else {
				s := string(b[:])
				return fmt.Sprintf("%#v", s)
			}
		}
	} else if cd.DefaultValue == nil || cd.IsAutoId {
		v := cd.ReceiverType.DefaultValueString()
		if v == "nil" {
			return ""
		}
		return cd.ReceiverType.DefaultValueString()
	} else {
		return fmt.Sprintf("%#v", cd.DefaultValue)
	}
}
*/

// JsonKey returns the key used for the column when outputting JSON.
func (cd *Column) JsonKey() string {
	return cd.DecapIdentifier
}

// IsReference returns true if the column is a reference to an object in another table.
func (cd *Column) IsReference() bool {
	return cd.Reference != nil && cd.Reference.Table != nil
}

// IsEnum returns true if the column contains a type defined by a enum table.
// This could be a ColTypeEnum or ColTypeEnumArray
func (cd *Column) IsEnum() bool {
	return cd.Reference != nil && cd.Reference.EnumTable != nil
}

// IsEnumArray returns true if the column contains a an enum array
func (cd *Column) IsEnumArray() bool {
	return cd.SchemaType == schema.ColTypeEnumArray
}

// ReferenceIdentifier returns the capitalized name that should be used to refer to the object
// in a forward reference. This is not the object's type.
func (cd *Column) ReferenceIdentifier() string {
	if cd.Reference == nil {
		panic(fmt.Sprintf("column %s.%s is not a reference", cd.Table.Identifier, cd.Identifier))
	}

	return cd.Reference.Identifier
}

// ReferenceType returns the name of the Go struct type in a forward reference.
func (cd *Column) ReferenceType() string {
	if cd.Reference == nil {
		panic(fmt.Sprintf("column %s.%s is not a reference", cd.Table.Identifier, cd.Identifier))
	}
	if cd.Reference.Table != nil {
		return cd.Reference.Table.Identifier
	}
	if cd.Reference.EnumTable != nil {
		return cd.Reference.EnumTable.Identifier
	}
	panic("reference does not have a Table or an Enum")
}

// ReferenceVariableIdentifier returns the name of the local variable that will
// hold the object loaded in the reference.
func (cd *Column) ReferenceVariableIdentifier() string {
	if cd.Reference == nil {
		panic(fmt.Sprintf("column %s.%s is not a reference", cd.Table.Identifier, cd.Identifier))
	}
	return "obj" + cd.Reference.Identifier
}

// ReverseIdentifier returns the function name that should be used to refer to the object
// that is referred to by the reverse reference.
func (cd *Column) ReverseIdentifier() string {
	if cd.Reference == nil {
		panic(fmt.Sprintf("column %s.%s is not a reference", cd.Table.Identifier, cd.Identifier))
	} else if cd.IsUnique {
		return cd.Reference.ReverseIdentifier
	} else {
		return cd.Reference.ReverseIdentifierPlural
	}
}

// ReverseVariableIdentifier returns the name of the local variable that will
// hold the object(s) loaded in the reverse reference.
func (cd *Column) ReverseVariableIdentifier() string {
	if cd.Reference == nil {
		panic(fmt.Sprintf("column %s.%s is not a reference", cd.Table.Identifier, cd.Identifier))
	} else {
		if cd.IsUnique {
			return "rev" + cd.Reference.ReverseIdentifier
		} else {
			return "rev" + cd.Reference.ReverseIdentifierPlural
		}
	}
}

// ReferenceJsonKey returns the key that will be used for the referenced object in JSON.
func (cd *Column) ReferenceJsonKey() string {
	if cd.Reference == nil {
		panic(fmt.Sprintf("column %s.%s is not a reference", cd.Table.Identifier, cd.Identifier))
	}
	return cd.Reference.DecapIdentifier
}

// ReverseJsonKey returns the key that will be used for the reverse referenced object in JSON.
func (cd *Column) ReverseJsonKey() string {
	if cd.Reference == nil {
		panic(fmt.Sprintf("column %s.%s is not a reference", cd.Table.Identifier, cd.Identifier))
	}
	if cd.IsUnique {
		return LowerCaseIdentifier(cd.Reference.ReverseIdentifier)
	} else {
		return LowerCaseIdentifier(cd.Reference.ReverseIdentifierPlural)
	}
}

// GoType return the Go type of the internal member variable corresponding to the column.
func (cd *Column) GoType() string {
	if ref := cd.Reference; ref != nil {
		if table := cd.Reference.Table; table != nil {
			if col := table.PrimaryKeyColumn(); col != nil {
				if col.IsAutoId {
					return "string" // auto id keys are always strings
				}
			}
		} else if enumTable := cd.Reference.EnumTable; enumTable != nil {
			// enum tables are an enumerated type
			if cd.SchemaType == schema.ColTypeEnum {
				return enumTable.Identifier
			} else if cd.SchemaType == schema.ColTypeEnumArray {
				return enumTable.Identifier + "Set"
			} else {
				panic("unknown column type for enum table")
			}
		} else {
			panic("reference is missing either a Table or Enum entry")
		}
	}
	return cd.ReceiverType.GoType()
}

// HasSetter returns true if the column should be allowed to be set by the programmer. Some columns should not be alterable,
// including AutoID columns, and time based columns that automatically set or update their times.
func (cd *Column) HasSetter() bool {
	if cd.IsAutoId {
		return false
	}
	if cd.ReceiverType == ColTypeTime {
		if cd.DefaultValue == CreatedTime || cd.DefaultValue == ModifiedTime {
			return false
		}
	}
	if cd.SchemaSubType == schema.ColSubTypeTimestamp ||
		cd.SchemaSubType == schema.ColSubTypeLock {
		return false
	}
	return true
}

// MaxInt returns the maximum integer that the column can hold if it is an integer type.
// Returns 0 if not.
func (cd *Column) MaxInt() int64 {
	if cd.ReceiverType == ColTypeInteger {
		switch cd.Size {
		case 8:
			return 127
		case 16:
			return 32767
		case 24:
			return 8388607
		case 32:
			return 2147483647
		}
	} else if cd.ReceiverType == ColTypeUnsigned {
		switch cd.Size {
		case 8:
			return 255
		case 16:
			return 65535
		case 24:
			return 16777215
		case 32:
			return 4294967295
		}
	}
	return 0
}

// MinInt returns the minimum integer that the column can hold if it is an integer type.
// Returns 0 if not.
func (cd *Column) MinInt() int64 {
	if cd.ReceiverType == ColTypeInteger {
		switch cd.Size {
		case 8:
			return -128
		case 16:
			return -32768
		case 24:
			return -8388608
		case 32:
			return -2147483648
		}
	}
	return 0
}

// MaxLength returns the maximum length of the column, which normally is Column.Size, but
// in certain situations might be something else.
func (cd *Column) MaxLength() uint64 {
	if cd.SchemaSubType == schema.ColSubTypeNumeric {
		return cd.Size&0x0000ffff + 2 // allow for +/- and decimal point
	}
	return cd.Size
}

// DecimalPrecision returns the precision value of a decimal number that is packed into the Size value.
func (cd *Column) DecimalPrecision() uint64 {
	if cd.SchemaSubType == schema.ColSubTypeNumeric {
		return cd.Size & 0x0000ffff
	}
	return 0
}

// DecimalScale returns the scale value of a decimal number when it is packed into the Size value.
func (cd *Column) DecimalScale() uint64 {
	if cd.SchemaSubType == schema.ColSubTypeNumeric {
		return cd.Size >> 16
	}
	return 0
}

func newColumn(schemaCol *schema.Column) *Column {
	var recType ReceiverType
	if schemaCol.Type == schema.ColTypeReference {
		recType = ReceiverTypeFromSchema(schema.ColTypeString, schemaCol.Size) // default. Will be fixed up later.
	} else {
		recType = ReceiverTypeFromSchema(schemaCol.Type, schemaCol.Size)
	}
	col := &Column{
		QueryName:       schemaCol.Name,
		Identifier:      schemaCol.Identifier,
		DecapIdentifier: strings2.Decap(schemaCol.Identifier),
		SchemaType:      schemaCol.Type,
		SchemaSubType:   schemaCol.SubType,
		ReceiverType:    recType,
		Size:            schemaCol.Size,
		DefaultValue:    schemaCol.DefaultValue,
		IsAutoId:        schemaCol.Type == schema.ColTypeAutoPrimaryKey,
		IsPrimaryKey:    schemaCol.Type == schema.ColTypeAutoPrimaryKey || schemaCol.IndexLevel == schema.IndexLevelManualPrimaryKey,
		IsNullable:      schemaCol.IsNullable,
		IsUnique: schemaCol.Type == schema.ColTypeAutoPrimaryKey ||
			schemaCol.IndexLevel == schema.IndexLevelManualPrimaryKey ||
			schemaCol.IndexLevel == schema.IndexLevelUnique,
	}

	if col.IsAutoId {
		col.ReceiverType = ColTypeString // We treat auto-generated ids as strings for cross database compatibility.
	}

	return col
}
