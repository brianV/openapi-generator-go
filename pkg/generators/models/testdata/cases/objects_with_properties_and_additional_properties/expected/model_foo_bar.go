// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	"encoding/json"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// FooBar is an object.
type FooBar struct {
	FooBarProperties
	AdditionalProperties map[string][]struct {
		Test string `json:"test,omitempty" mapstructure:"test,omitempty"`
	}
}

type FooBarProperties struct {
	// Baz:
	Baz string `json:"baz,omitempty" mapstructure:"baz,omitempty"`
}

// Unmarshal all named properties into FooBarProperties and
// the rest into the AdditionalProperties map
func (obj *FooBar) UnmarshalJSON(data []byte) error {
	var generic map[string]json.RawMessage
	if err := json.Unmarshal(data, &generic); err != nil {
		return err
	}

	obj.FooBarProperties = FooBarProperties{}

	var additionalProperties = make(map[string][]struct {
		Test string `json:"test,omitempty" mapstructure:"test,omitempty"`
	})
	for k, v := range generic {
		if k == "baz" {
			if err := json.Unmarshal(v, &(obj.FooBarProperties.Baz)); err != nil {
				return err
			}
			continue
		}
		var prop []struct {
			Test string `json:"test,omitempty" mapstructure:"test,omitempty"`
		}
		if err := json.Unmarshal(v, &prop); err != nil {
			return err
		}
		additionalProperties[k] = prop
	}

	obj.AdditionalProperties = additionalProperties
	return nil
}

// Marshal FooBar by combining the AdditionalProperties with the
// named properties in a single JSON object. Named properties take
// precedence.
func (obj FooBar) MarshalJSON() ([]byte, error) {
	props := make(map[string]json.RawMessage)

	// start with additional properties so regular properties overwrite them
	for k, v := range obj.AdditionalProperties {
		if propData, err := json.Marshal(v); err == nil {
			props[k] = propData
		} else {
			return nil, err
		}
	}
	if propData, err := json.Marshal(obj.FooBarProperties.Baz); err == nil {
		props["baz"] = propData
	} else {
		return nil, err
	}

	data, err := json.Marshal(props)
	return data, err
}

// Validate implements basic validation for this model
func (m FooBar) Validate() error {
	return validation.Errors{}.Filter()
}

// GetBaz returns the Baz property
func (m FooBar) GetBaz() string {
	return m.Baz
}

// SetBaz sets the Baz property
func (m *FooBar) SetBaz(val string) {
	m.Baz = val
}

func (m *FooBar) GetAdditionalProperties() map[string][]struct {
	Test string `json:"test,omitempty" mapstructure:"test,omitempty"`
} {
	return m.AdditionalProperties
}

func (m *FooBar) SetAdditionalProperties(val map[string][]struct {
	Test string `json:"test,omitempty" mapstructure:"test,omitempty"`
}) {
	m.AdditionalProperties = val
}
