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

// Gender is an enum.
type Gender string

// Validate implements basic validation for this model
func (m Gender) Validate() error {
	return InKnownGender.Validate(m)
}

var (
	GenderDefault Gender = "default"
	GenderFemale  Gender = "female"
	GenderMale    Gender = "male"

	// KnownGender is the list of valid Gender
	KnownGender = []Gender{
		GenderDefault,
		GenderFemale,
		GenderMale,
	}
	// KnownGenderString is the list of valid Gender as string
	KnownGenderString = []string{
		string(GenderDefault),
		string(GenderFemale),
		string(GenderMale),
	}

	// InKnownGender is an ozzo-validator for Gender
	InKnownGender = validation.In(
		GenderDefault,
		GenderFemale,
		GenderMale,
	)
)
