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

// Artist is an object.
type Artist struct {
	// LeftHand:
	LeftHand *AnyThing `json:"leftHand,omitempty" mapstructure:"leftHand,omitempty"`
	// RightHand:
	RightHand *Color `json:"rightHand,omitempty" mapstructure:"rightHand,omitempty"`
}

// Validate implements basic validation for this model
func (m Artist) Validate() error {
	return validation.Errors{
		"leftHand": validation.Validate(
			m.LeftHand,
		),
		"rightHand": validation.Validate(
			m.RightHand,
		),
	}.Filter()
}

// GetLeftHand returns the LeftHand property
func (m Artist) GetLeftHand() *AnyThing {
	return m.LeftHand
}

// SetLeftHand sets the LeftHand property
func (m *Artist) SetLeftHand(val *AnyThing) {
	m.LeftHand = val
}

// GetRightHand returns the RightHand property
func (m Artist) GetRightHand() *Color {
	return m.RightHand
}

// SetRightHand sets the RightHand property
func (m *Artist) SetRightHand(val *Color) {
	m.RightHand = val
}
