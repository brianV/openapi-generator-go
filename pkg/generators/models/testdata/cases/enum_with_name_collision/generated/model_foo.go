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

// Foo is an enum.
type Foo string

// Validate implements basic validation for this model
func (m Foo) Validate() error {
	return InKnownFoo.Validate(m)
}

var (
	FooItem0 Foo = "{{{"

	// KnownFoo is the list of valid Foo
	KnownFoo = []Foo{
		FooItem0,
	}
	// KnownFooString is the list of valid Foo as string
	KnownFooString = []string{
		string(FooItem0),
	}

	// InKnownFoo is an ozzo-validator for Foo
	InKnownFoo = validation.In(
		FooItem0,
	)
)
