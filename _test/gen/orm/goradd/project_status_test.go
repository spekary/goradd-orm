// Code generated by goradd-orm. DO NOT EDIT.

package goradd

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProjectStatus_String(t *testing.T) {
	assert.Equal(t, ProjectStatusOpen.String(), ProjectStatusOpen.Identifier())
	assert.Equal(t, ProjectStatusCancelled.String(), ProjectStatusCancelled.Identifier())
	assert.Equal(t, ProjectStatusCompleted.String(), ProjectStatusCompleted.Identifier())
	assert.Equal(t, ProjectStatusPlanned.String(), ProjectStatusPlanned.Identifier())
}

func TestProjectStatus_Label(t *testing.T) {
	assert.Equal(t, ProjectStatusOpen.Label(), "open")
	assert.Equal(t, ProjectStatusCancelled.Label(), "cancelled")
	assert.Equal(t, ProjectStatusCompleted.Label(), "completed")
	assert.Equal(t, ProjectStatusPlanned.Label(), "planned")
	assert.Zero(t, ProjectStatus(0).Label())
	assert.Panics(t, func() {
		ProjectStatus(-1).Label()
	})
}

func TestProjectStatus_LabelSlice(t *testing.T) {
	a := ProjectStatusLabels()
	assert.Equal(t, ProjectStatusOpen.Label(), a[0])
	assert.Equal(t, ProjectStatusCancelled.Label(), a[1])
	assert.Equal(t, ProjectStatusCompleted.Label(), a[2])
	assert.Equal(t, ProjectStatusPlanned.Label(), a[3])
}

func TestProjectStatus_Identifier(t *testing.T) {
	assert.Equal(t, ProjectStatusOpen.Identifier(), "open")
	assert.Equal(t, ProjectStatusCancelled.Identifier(), "cancelled")
	assert.Equal(t, ProjectStatusCompleted.Identifier(), "completed")
	assert.Equal(t, ProjectStatusPlanned.Identifier(), "planned")
	assert.Zero(t, ProjectStatus(0).Identifier())
	assert.Panics(t, func() {
		ProjectStatus(-1).Identifier()
	})
}

func TestProjectStatus_IdentifierSlice(t *testing.T) {
	a := ProjectStatusIdentifiers()
	assert.Equal(t, ProjectStatusOpen.Identifier(), a[0])
	assert.Equal(t, ProjectStatusCancelled.Identifier(), a[1])
	assert.Equal(t, ProjectStatusCompleted.Identifier(), a[2])
	assert.Equal(t, ProjectStatusPlanned.Identifier(), a[3])
}

func TestProjectStatus_Description(t *testing.T) {
	assert.Equal(t, ProjectStatusOpen.Description(), "The project is currently active")
	assert.Equal(t, ProjectStatusCancelled.Description(), "The project has been canned")
	assert.Equal(t, ProjectStatusCompleted.Description(), "The project has been completed successfully")
	assert.Equal(t, ProjectStatusPlanned.Description(), "Project is in the planning stages and has not been assigned a manager")
	assert.Zero(t, ProjectStatus(0).Description())
	assert.Panics(t, func() {
		ProjectStatus(-1).Description()
	})
}

func TestProjectStatus_DescriptionSlice(t *testing.T) {
	a := ProjectStatusDescriptions()
	assert.Equal(t, ProjectStatusOpen.Description(), a[0])
	assert.Equal(t, ProjectStatusCancelled.Description(), a[1])
	assert.Equal(t, ProjectStatusCompleted.Description(), a[2])
	assert.Equal(t, ProjectStatusPlanned.Description(), a[3])
}

func TestProjectStatus_Guidelines(t *testing.T) {
	assert.Equal(t, ProjectStatusOpen.Guidelines(), "All projects that we are working on should be in this state")
	assert.Equal(t, ProjectStatusCancelled.Guidelines(), "")
	assert.Equal(t, ProjectStatusCompleted.Guidelines(), "Celebrate successes!")
	assert.Equal(t, ProjectStatusPlanned.Guidelines(), "Get ready")
	assert.Zero(t, ProjectStatus(0).Guidelines())
	assert.Panics(t, func() {
		ProjectStatus(-1).Guidelines()
	})
}

func TestProjectStatus_GuidelinesSlice(t *testing.T) {
	a := ProjectStatusGuidelines()
	assert.Equal(t, ProjectStatusOpen.Guidelines(), a[0])
	assert.Equal(t, ProjectStatusCancelled.Guidelines(), a[1])
	assert.Equal(t, ProjectStatusCompleted.Guidelines(), a[2])
	assert.Equal(t, ProjectStatusPlanned.Guidelines(), a[3])
}

func TestProjectStatus_IsActive(t *testing.T) {
	assert.Equal(t, ProjectStatusOpen.IsActive(), true)
	assert.Equal(t, ProjectStatusCancelled.IsActive(), true)
	assert.Equal(t, ProjectStatusCompleted.IsActive(), true)
	assert.Equal(t, ProjectStatusPlanned.IsActive(), false)
	assert.Zero(t, ProjectStatus(0).IsActive())
	assert.Panics(t, func() {
		ProjectStatus(-1).IsActive()
	})
}

func TestProjectStatus_IsActiveSlice(t *testing.T) {
	a := ProjectStatusIsActives()
	assert.Equal(t, ProjectStatusOpen.IsActive(), a[0])
	assert.Equal(t, ProjectStatusCancelled.IsActive(), a[1])
	assert.Equal(t, ProjectStatusCompleted.IsActive(), a[2])
	assert.Equal(t, ProjectStatusPlanned.IsActive(), a[3])
}

func TestProjectStatus_Keys(t *testing.T) {
	var keys []string

	keys = append(keys, ProjectStatusOpen.Key())
	keys = append(keys, ProjectStatusCancelled.Key())
	keys = append(keys, ProjectStatusCompleted.Key())
	keys = append(keys, ProjectStatusPlanned.Key())
	v := ProjectStatusesFromKeys(keys)
	assert.Equal(t, ProjectStatuses(), v)

	assert.Equal(t, ProjectStatus(0), ProjectStatusFromKey(""))
}

func TestProjectStatus_Values(t *testing.T) {
	a1 := ProjectStatuses()
	a2 := ProjectStatusesI()
	for i, v1 := range a1 {
		assert.Equal(t, v1, a2[i].(ProjectStatus))
		assert.True(t, IsValidProjectStatus(int(v1)))
	}
	assert.False(t, IsValidProjectStatus(0))
}

func TestProjectStatus_FromIdentifier(t *testing.T) {
	var v ProjectStatus
	var err error

	v, err = ProjectStatusFromIdentifier(ProjectStatus(1).Identifier())
	assert.NoError(t, err)
	assert.Equal(t, ProjectStatus(1), v)
	v, err = ProjectStatusFromIdentifier(ProjectStatus(2).Identifier())
	assert.NoError(t, err)
	assert.Equal(t, ProjectStatus(2), v)
	v, err = ProjectStatusFromIdentifier(ProjectStatus(3).Identifier())
	assert.NoError(t, err)
	assert.Equal(t, ProjectStatus(3), v)
	v, err = ProjectStatusFromIdentifier(ProjectStatus(4).Identifier())
	assert.NoError(t, err)
	assert.Equal(t, ProjectStatus(4), v)
}

func TestProjectStatus_FromInterface(t *testing.T) {
	v := ProjectStatusOpen

	v2, err := ProjectStatusFromInterface(int(v))
	assert.NoError(t, err)
	assert.Equal(t, v, v2)

	v2, err = ProjectStatusFromInterface(float64(v))
	assert.NoError(t, err)
	assert.Equal(t, v, v2)

	v2, err = ProjectStatusFromInterface(json.Number(fmt.Sprintf("%d", v)))
	assert.NoError(t, err)
	assert.Equal(t, v, v2)

	v2, err = ProjectStatusFromInterface(fmt.Sprintf("%d", v))
	assert.NoError(t, err)
	assert.Equal(t, v, v2)

	v2, err = ProjectStatusFromInterface(0)
	assert.Error(t, err)
	assert.Equal(t, ProjectStatus(0), v2)

	var t1 time.Time
	v2, err = ProjectStatusFromInterface(t1)
	assert.Error(t, err)
	assert.Equal(t, ProjectStatus(0), v2)
}

func TestProjectStatus_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected ProjectStatus
		wantErr  bool
	}{
		// Integer JSON
		{input: `1`, expected: ProjectStatusOpen, wantErr: false},
		// Float JSON
		{input: `1.0`, expected: ProjectStatusOpen, wantErr: false},
		// Stringified numbers
		{input: `"1"`, expected: ProjectStatusOpen, wantErr: false},
		// Invalid values
		{input: `"1.1.1"`, expected: ProjectStatus(0), wantErr: true},
	}

	for _, tt := range tests {
		var s ProjectStatus
		err := json.Unmarshal([]byte(tt.input), &s)
		if (err != nil) != tt.wantErr {
			t.Errorf("UnmarshalJSON(%s) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if err == nil && s != tt.expected {
			t.Errorf("UnmarshalJSON(%s) = %v, want %v", tt.input, s, tt.expected)
		}
	}
}
