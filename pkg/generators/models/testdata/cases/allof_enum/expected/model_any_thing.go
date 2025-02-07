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

// AnyThing is an enum.
type AnyThing string

// Validate implements basic validation for this model
func (m AnyThing) Validate() error {
	return InKnownAnyThing.Validate(m)
}

var (
	AnyThingApple  AnyThing = "apple"
	AnyThingBanana AnyThing = "banana"
	AnyThingBlue   AnyThing = "blue"
	AnyThingGreen  AnyThing = "green"
	AnyThingPizza  AnyThing = "pizza"
	AnyThingRed    AnyThing = "red"

	// KnownAnyThing is the list of valid AnyThing
	KnownAnyThing = []AnyThing{
		AnyThingApple,
		AnyThingBanana,
		AnyThingBlue,
		AnyThingGreen,
		AnyThingPizza,
		AnyThingRed,
	}
	// KnownAnyThingString is the list of valid AnyThing as string
	KnownAnyThingString = []string{
		string(AnyThingApple),
		string(AnyThingBanana),
		string(AnyThingBlue),
		string(AnyThingGreen),
		string(AnyThingPizza),
		string(AnyThingRed),
	}

	// InKnownAnyThing is an ozzo-validator for AnyThing
	InKnownAnyThing = validation.In(
		AnyThingApple,
		AnyThingBanana,
		AnyThingBlue,
		AnyThingGreen,
		AnyThingPizza,
		AnyThingRed,
	)
)
