// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

// Baz is a oneOf type.
type Baz struct {
	data interface{}
}

// MarshalJSON implements the json.Marshaller interface
func (m Baz) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the json.Unmarshaller interface
func (m *Baz) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &m.data)
}

// FromFoo sets the Baz data.
func (m *Baz) FromFoo(data Foo) {
	m.data = data
}

// FromBar sets the Baz data.
func (m *Baz) FromBar(data Bar) {
	m.data = data
}

// FromPerson sets the Baz data.
func (m *Baz) FromPerson(data Person) {
	m.data = data
}

// As converts Baz to a user defined structure.
func (m Baz) As(target interface{}) error {
	return mapstructure.Decode(m.data, target)
}

// AsFoo converts Baz to a Foo
func (m Baz) AsFoo() (result Foo, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsBar converts Baz to a Bar
func (m Baz) AsBar() (result Bar, err error) {
	return result, mapstructure.Decode(m.data, &result)
}

// AsPerson converts Baz to a Person
func (m Baz) AsPerson() (result Person, err error) {
	return result, mapstructure.Decode(m.data, &result)
}
