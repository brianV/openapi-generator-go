// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// User is an object.
type User struct {
	// DateOfBirth:
	DateOfBirth string `json:"date_of_birth,omitempty" mapstructure:"date_of_birth,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// UserId:
	UserId string `json:"user_id,omitempty" mapstructure:"user_id,omitempty"`
}

// Validate implements basic validation for this model
func (m User) Validate() error {
	return validation.Errors{}.Filter()
}

// GetDateOfBirth returns the DateOfBirth property
func (m User) GetDateOfBirth() string {
	return m.DateOfBirth
}

// SetDateOfBirth sets the DateOfBirth property
func (m *User) SetDateOfBirth(val string) {
	m.DateOfBirth = val
}

// GetName returns the Name property
func (m User) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *User) SetName(val string) {
	m.Name = val
}

// GetUserId returns the UserId property
func (m User) GetUserId() string {
	return m.UserId
}

// SetUserId sets the UserId property
func (m *User) SetUserId(val string) {
	m.UserId = val
}
