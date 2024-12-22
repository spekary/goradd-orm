// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"strconv"
)

type ProjectStatus int

const (
	ProjectStatusOpen      ProjectStatus = 1
	ProjectStatusCancelled ProjectStatus = 2
	ProjectStatusCompleted ProjectStatus = 3
	ProjectStatusPlanned   ProjectStatus = 4
)

// ProjectStatusMaxValue is the maximum enumerated value of ProjectStatus
// doc: type=ProjectStatus
const ProjectStatusMaxValue = 4

// String returns the name value of the type and satisfies the fmt.Stringer interface
func (e ProjectStatus) String() string {
	switch e {
	case 0:
		return ""
	case ProjectStatusOpen:
		return `Open`
	case ProjectStatusCancelled:
		return `Cancelled`
	case ProjectStatusCompleted:
		return `Completed`
	case ProjectStatusPlanned:
		return `Planned`
	default:
		panic("index out of range")
	}
	return "" // prevent warning
}

// ID returns a string representation of the id and satisfies the IDer interface
func (e ProjectStatus) ID() string {
	return strconv.Itoa(int(e))
}

// ProjectStatusFromID converts a ProjectStatus ID to a ProjectStatus
func ProjectStatusFromID(id string) ProjectStatus {
	switch id {
	case `1`:
		return ProjectStatusOpen
	case `2`:
		return ProjectStatusCancelled
	case `3`:
		return ProjectStatusCompleted
	case `4`:
		return ProjectStatusPlanned
	}
	return ProjectStatus(0)
}

// ProjectStatusesFromIDs converts a slice of ProjectStatus IDs to a slice of ProjectStatus
func ProjectStatusesFromIDs(ids []string) (values []ProjectStatus) {
	values = make([]ProjectStatus, 0, len(ids))
	for _, id := range ids {
		values = append(values, ProjectStatusFromID(id))
	}
	return
}

// ProjectStatusFromName converts a ProjectStatus name to a ProjectStatus
func ProjectStatusFromName(name string) ProjectStatus {
	switch name {
	case `Open`:
		return ProjectStatusOpen
	case `Cancelled`:
		return ProjectStatusCancelled
	case `Completed`:
		return ProjectStatusCompleted
	case `Planned`:
		return ProjectStatusPlanned
	}
	return ProjectStatus(0)
}

// ProjectStatuses returns a slice of all the ProjectStatus values
// in key order.
func ProjectStatuses() []ProjectStatus {
	return []ProjectStatus{
		ProjectStatusOpen,
		ProjectStatusCancelled,
		ProjectStatusCompleted,
		ProjectStatusPlanned,
	}
}

// ProjectStatusesI returns a slice of all the ProjectStatus values as generic interfaces.
// doc: type=ProjectStatus
func ProjectStatusesI() (values []any) {
	return []any{
		ProjectStatusOpen,
		ProjectStatusCancelled,
		ProjectStatusCompleted,
		ProjectStatusPlanned,
	}
}

// Label returns the string that will be displayed to a user for this item.
func (e ProjectStatus) Label() string {
	return e.String()
}

// Value returns the value as an interface. It satisfies goradd's Valuer interface.
func (e ProjectStatus) Value() any {
	return e.ID()
}

func (e ProjectStatus) Description() string {
	switch e {
	case 0:
		return ""
	case ProjectStatusOpen:
		return "The project is currently active"
	case ProjectStatusCancelled:
		return "The project has been canned"
	case ProjectStatusCompleted:
		return "The project has been completed successfully"
	case ProjectStatusPlanned:
		return "Project is in the planning stages and has not been assigned a manager"
	default:
		panic("Index out of range")
	}
	return "" // prevent warning
}

// ProjectStatusDescriptions returns a slice of all the Descriptions associated with ProjectStatus values.
// doc: type=ProjectStatus
func ProjectStatusDescriptions() []string {
	return []string{
		// 0 item will be a zero value
		"",
		"The project is currently active",
		"The project has been canned",
		"The project has been completed successfully",
		"Project is in the planning stages and has not been assigned a manager",
	}
}

func (e ProjectStatus) Guidelines() string {
	switch e {
	case 0:
		return ""
	case ProjectStatusOpen:
		return "All projects that we are working on should be in this state"
	case ProjectStatusCancelled:
		return ""
	case ProjectStatusCompleted:
		return "Celebrate successes!"
	case ProjectStatusPlanned:
		return "Get ready"
	default:
		panic("Index out of range")
	}
	return "" // prevent warning
}

// ProjectStatusGuidelines returns a slice of all the Guidelines associated with ProjectStatus values.
// doc: type=ProjectStatus
func ProjectStatusGuidelines() []string {
	return []string{
		// 0 item will be a zero value
		"",
		"All projects that we are working on should be in this state",
		"",
		"Celebrate successes!",
		"Get ready",
	}
}

func (e ProjectStatus) IsActive() bool {
	switch e {
	case 0:
		return false
	case ProjectStatusOpen:
		return true
	case ProjectStatusCancelled:
		return true
	case ProjectStatusCompleted:
		return true
	case ProjectStatusPlanned:
		return false
	default:
		panic("Index out of range")
	}
	return false // prevent warning
}

// ProjectStatusIsActives returns a slice of all the Is Actives associated with ProjectStatus values.
// doc: type=ProjectStatus
func ProjectStatusIsActives() []bool {
	return []bool{
		// 0 item will be a zero value
		false,
		true,
		true,
		true,
		false,
	}
}