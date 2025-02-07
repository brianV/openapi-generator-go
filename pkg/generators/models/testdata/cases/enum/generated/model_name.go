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

// Name is an enum.
type Name string

// Validate implements basic validation for this model
func (m Name) Validate() error {
	return InKnownName.Validate(m)
}

var (
	NameBar Name = "bar"
	NameFoo Name = "foo"

	// KnownName is the list of valid Name
	KnownName = []Name{
		NameBar,
		NameFoo,
	}
	// KnownNameString is the list of valid Name as string
	KnownNameString = []string{
		string(NameBar),
		string(NameFoo),
	}

	// InKnownName is an ozzo-validator for Name
	InKnownName = validation.In(
		NameBar,
		NameFoo,
	)
)
