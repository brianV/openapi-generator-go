// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test oneOf discriminator support
//	Version: 0.1.0
package generatortest

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// NonEmptyContainer is an object.
type NonEmptyContainer struct {
	// Error:
	Error Error `json:"error" mapstructure:"error"`
}

// Validate implements basic validation for this model
func (m NonEmptyContainer) Validate() error {
	return validation.Errors{
		"error": validation.Validate(
			m.Error, validation.NotNil,
		),
	}.Filter()
}

// GetError returns the Error property
func (m NonEmptyContainer) GetError() Error {
	return m.Error
}

// SetError sets the Error property
func (m *NonEmptyContainer) SetError(val Error) {
	m.Error = val
}
