package schema

import "encoding/json"

// ColumnType is a general specifier for the type of column in a database.
//
// Mapping data types in Go to data types in databases is tricky, and necessarily
// puts limitations on what kinds of data can be represented and how it is stored.
// The ORM provides some flexibility by allowing each column to specify its database
// type, while also specifying its Go type by using a combination of the DatabaseType,
// VariableType and Size fields as described below.
//
// # ColTypeUnknown
//
// This is a column type that is unknown to Go and that will be provided to Go as a []byte array
// or a string depending on the database driver.
// Many databases have non-standard or custom data types or types that simply don't fit well in Go.
// For example, the NUMERIC or DECIMAL type in many databases is an exact precision decimal
// number for which there is no equivalent in Go, since Go uses floating point numbers to represent
// decimal numbers. To provide data of this type to your application, you can create a custom
// type in Go, and then create accessor functions that translate between the database representation
// and your custom type.
//
// # ColTypeBytes
//
// This is known as a BLOB type in many SQL databases, and is represented as a []byte slice in Go.
// If there is a MaxSize value on the column, the data definition will use this in the schema definition.
// Different databases treat this differently, and may or may not truncate the data internally.
// The ORM will truncate the data and issue a warning to the log file if an attempt to insert
// too big of a byte blob is attempted.
//
// # ColTypeString
//
// ColTypeString is always represented as a string in Go, but may have a variety of representations in
// the database depending on the MaxSize value and the database. If a Column.Size is set,
// the generated code will truncate strings to this number of runes before sending it to the database.
//
// Postgres doc states the TEXT type is the fastest and most efficient text storage type,
// and by default this will always be the corresponding data type in Postgres.
//
// MySQL will use a VarChar(MaxSize) if a maximum size is present, and a Text if not. Mysql stores VarChar
// variables inside the table, and Text data outside of the table with a reference, so VarChar is more efficient.
//
// By default, collation will be by UTF8 case-sensitive rules. Set the Sort to case-insensitive to
// change this to case-insensitive. Other collations are outside the capabilities of the ORM and you
// should directly edit the database to set those up.
//
// # ColTypeInt and ColTypeUint
//
// These integer column types will be represented in Go as int or uint, and in the database as a
// 32-bit integer or unsigned integer by default. To specifically set the storage size, specify either
// 8, 16, 32 or 64 in the Size value, and the corresponding Go type will be used, as well as
// the corresponding type in the database. Not all databases support 8-bit integers. Check
// your database vendor to be sure.
//
// # ColTypeTime
//
// This corresponds to a datetime in most databases. The value is stored in UTC time
// in the database, and when read back from the database, will initially be in UTC time.
// Usually this is what is wanted, since it allows times from different timezones to be correctly
// sorted, and javascript and other libraries are capable of converting UTC time to local
// time in the client locale for display purposes.
// MySQL will store the value as a DateTime and not a Timestamp, since Timestamps are assumed
// to be in server local time and not UTC time and get time shifted
// in transit. Postgres uses Timestamp without timezone (which is the default).
// When specifying a default time, enter a Go time.Time type, or the string "now" to have
// the ORM set it to the time the object was created.
//
// # ColTypeFloat
//
// By default, this will be a float64 value in Go, and double precision float in the database.
// Specify a MaxSize of 32 to make this a float32 and single precision float in the database (if supported).
//
// # ColTypeBool
//
// ColTypeBool is a bool in Go, and is database dependent in the database. MySQL uses a BIT column,
// and Postgres uses a boolean.
//
// # ColTypeAutoPrimaryKey
//
// These are auto-generated primary keys generated by the database or database driver,
// and are represented in Go as a string, even though databases can use a variety of types for
// a primary key internally. This is primarily for portability, so that if you change databases, you do not have to change
// your code. MySQL will use an auto incremented integer internally, and Postgres will use a serial type, which is an
// auto incremented integer. It is up to the database driver to ensure that the primary key is unique. The primary
// key is used in reference columns to create associations between tables.
//
// The ORM allows multi-column primary keys in the database, however such keys are not ColTypePrimaryKey, but rather
// whatever the type is they are assigned in the database. They are treated as multi-column unique indexes by the ORM,
// and cannot be auto-generated by the database driver, but rather must be set by the application.
//
// # ColTypeJSON
//
// This special type is supported by many databases and allows querying and data retrieval from
// within JSON documents stored in a field in a database. In Go, querying the entire field will
// result in a string type value.
//
// # ColTypeReference
//
// This is a column that contains the value of a primary key of a table, creating a reference to the
// object in that table. This is known as a foreign key in SQL databases or an edge in graph databases.
type ColumnType int

const (
	ColTypeUnknown ColumnType = iota
	ColTypeBytes
	ColTypeString
	ColTypeInt
	ColTypeUint
	ColTypeTime
	ColTypeFloat
	ColTypeBool
	ColTypeAutoPrimaryKey
	ColTypeJSON
	ColTypeReference
)

// String returns the string representation of a ColumnType.
func (ct ColumnType) String() string {
	switch ct {
	case ColTypeBytes:
		return "Bytes"
	case ColTypeString:
		return "String"
	case ColTypeInt:
		return "Int"
	case ColTypeUint:
		return "Uint"
	case ColTypeTime:
		return "Time"
	case ColTypeFloat:
		return "Float"
	case ColTypeBool:
		return "Bool"
	case ColTypeAutoPrimaryKey:
		return "AutoPrimaryKey"
	case ColTypeJSON:
		return "JSON"
	case ColTypeReference:
		return "Reference"
	default:
		return "Unknown"
	}
}

// MarshalJSON customizes how ColumnType is serialized to JSON.
func (ct ColumnType) MarshalJSON() ([]byte, error) {
	// Return the string representation of the Type
	return json.Marshal(ct.String())
}

// UnmarshalJSON customizes how ColumnType is deserialized from JSON.
func (ct *ColumnType) UnmarshalJSON(data []byte) error {
	var ctStr string
	if err := json.Unmarshal(data, &ctStr); err != nil {
		return err
	}

	// Match the string representation and assign the corresponding Type value
	switch ctStr {
	case "Bytes":
		*ct = ColTypeBytes
	case "String":
		*ct = ColTypeString
	case "Int":
		*ct = ColTypeInt
	case "Uint":
		*ct = ColTypeUint
	case "Time":
		*ct = ColTypeTime
	case "Float":
		*ct = ColTypeFloat
	case "Bool":
		*ct = ColTypeBool
	case "AutoPrimaryKey":
		*ct = ColTypeAutoPrimaryKey
	case "JSON":
		*ct = ColTypeJSON
	case "Reference":
		*ct = ColTypeReference
	default:
		*ct = ColTypeUnknown
	}
	return nil
}
