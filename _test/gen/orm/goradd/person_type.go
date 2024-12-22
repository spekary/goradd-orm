// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"strconv"
)

type PersonType int

const (
	PersonTypeContractor    PersonType = 1
	PersonTypeManager       PersonType = 2
	PersonTypeInactive      PersonType = 3
	PersonTypeCompanyCar    PersonType = 4
	PersonTypeWorksFromHome PersonType = 5
)

// PersonTypeMaxValue is the maximum enumerated value of PersonType
// doc: type=PersonType
const PersonTypeMaxValue = 5

// String returns the name value of the type and satisfies the fmt.Stringer interface
func (e PersonType) String() string {
	switch e {
	case 0:
		return ""
	case PersonTypeContractor:
		return `Contractor`
	case PersonTypeManager:
		return `Manager`
	case PersonTypeInactive:
		return `Inactive`
	case PersonTypeCompanyCar:
		return `Company Car`
	case PersonTypeWorksFromHome:
		return `Works From Home`
	default:
		panic("index out of range")
	}
	return "" // prevent warning
}

// ID returns a string representation of the id and satisfies the IDer interface
func (e PersonType) ID() string {
	return strconv.Itoa(int(e))
}

// PersonTypeFromID converts a PersonType ID to a PersonType
func PersonTypeFromID(id string) PersonType {
	switch id {
	case `1`:
		return PersonTypeContractor
	case `2`:
		return PersonTypeManager
	case `3`:
		return PersonTypeInactive
	case `4`:
		return PersonTypeCompanyCar
	case `5`:
		return PersonTypeWorksFromHome
	}
	return PersonType(0)
}

// PersonTypesFromIDs converts a slice of PersonType IDs to a slice of PersonType
func PersonTypesFromIDs(ids []string) (values []PersonType) {
	values = make([]PersonType, 0, len(ids))
	for _, id := range ids {
		values = append(values, PersonTypeFromID(id))
	}
	return
}

// PersonTypeFromName converts a PersonType name to a PersonType
func PersonTypeFromName(name string) PersonType {
	switch name {
	case `Contractor`:
		return PersonTypeContractor
	case `Manager`:
		return PersonTypeManager
	case `Inactive`:
		return PersonTypeInactive
	case `Company Car`:
		return PersonTypeCompanyCar
	case `Works From Home`:
		return PersonTypeWorksFromHome
	}
	return PersonType(0)
}

// PersonTypes returns a slice of all the PersonType values
// in key order.
func PersonTypes() []PersonType {
	return []PersonType{
		PersonTypeContractor,
		PersonTypeManager,
		PersonTypeInactive,
		PersonTypeCompanyCar,
		PersonTypeWorksFromHome,
	}
}

// PersonTypesI returns a slice of all the PersonType values as generic interfaces.
// doc: type=PersonType
func PersonTypesI() (values []any) {
	return []any{
		PersonTypeContractor,
		PersonTypeManager,
		PersonTypeInactive,
		PersonTypeCompanyCar,
		PersonTypeWorksFromHome,
	}
}

// Label returns the string that will be displayed to a user for this item.
func (e PersonType) Label() string {
	return e.String()
}

// Value returns the value as an interface. It satisfies goradd's Valuer interface.
func (e PersonType) Value() any {
	return e.ID()
}
