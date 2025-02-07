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

// ListItem is an object.
type ListItem struct {
	// Next: the next item
	Next *ListItem `json:"next,omitempty" mapstructure:"next,omitempty"`
	// Value:
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}

// Validate implements basic validation for this model
func (m ListItem) Validate() error {
	return validation.Errors{
		"next": validation.Validate(
			m.Next,
		),
	}.Filter()
}

// GetNext returns the Next property
func (m ListItem) GetNext() *ListItem {
	return m.Next
}

// SetNext sets the Next property
func (m *ListItem) SetNext(val *ListItem) {
	m.Next = val
}

// GetValue returns the Value property
func (m ListItem) GetValue() string {
	return m.Value
}

// SetValue sets the Value property
func (m *ListItem) SetValue(val string) {
	m.Value = val
}
