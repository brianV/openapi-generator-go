// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Test
//	Version: 0.1.0
package generatortest

import (
	"encoding/json"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/mitchellh/mapstructure"
)

// GeometryDiscriminator is the discriminator value used during serialization and validation.
// Retrieve this value using  the GeometryDiscriminator() method
type GeometryDiscriminator string

const (
	// GeometryDiscriminatorLine maps Geometry to Line
	GeometryDiscriminatorLine GeometryDiscriminator = "Line"
	// GeometryDiscriminatorShape maps Geometry to Shape
	GeometryDiscriminatorShape GeometryDiscriminator = "Shape"
)

type geometryer interface {
	GeometryDiscriminator() GeometryDiscriminator
	Validate() error
}

// Geometry is a oneOf type.
type Geometry struct {
	data geometryer
}

var EmptyGeometryError = fmt.Errorf("empty data is not an Geometry")
var NotGeometryError = fmt.Errorf("could not convert to type Geometry")

type GeometryNilableRule struct{}

func (nn GeometryNilableRule) Validate(v interface{}) error {
	if m, ok := v.(Geometry); !ok {
		return NotGeometryError
	} else if m.data == nil {
		return nil
	} else {
		return m.Validate()
	}
}

// MarshalJSON implements the json.Marshaller interface
func (m Geometry) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

// UnmarshalJSON implements the json.Unmarshaller interface
func (m *Geometry) UnmarshalJSON(bs []byte) error {
	discriminator := struct {
		Value string `json:"type"`
	}{}
	err := json.Unmarshal(bs, &discriminator)
	if err != nil {
		return err
	}

	if discriminator.Value == "" {
		return validation.Errors{
			"type": fmt.Errorf("can not be empty"),
		}.Filter()
	}

	value := fmt.Sprintf("%v", discriminator.Value)
	switch GeometryDiscriminator(value) {
	case GeometryDiscriminatorLine:
		data := Line{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	case GeometryDiscriminatorShape:
		data := Shape{}
		err := json.Unmarshal(bs, &data)
		if err != nil {
			return err
		}
		m.data = data
	default:
		return validation.Errors{
			"type": fmt.Errorf("invalid value"),
		}.Filter()
	}
	return nil
}

// FromLine sets the Geometry data.
func (m *Geometry) FromLine(data Line) {
	m.data = data
}

// FromShape sets the Geometry data.
func (m *Geometry) FromShape(data Shape) {
	m.data = data
}

// As converts Geometry to a user defined structure.
func (m Geometry) As(target interface{}) error {
	if m.data == nil {
		return EmptyGeometryError
	}
	return mapstructure.Decode(m.data, target)
}

// AsLine converts Geometry to a Line
func (m Geometry) AsLine() (result Line, err error) {
	if m.data == nil {
		return result, EmptyGeometryError
	}
	return result, mapstructure.Decode(m.data, &result)
}

// AsShape converts Geometry to a Shape
func (m Geometry) AsShape() (result Shape, err error) {
	if m.data == nil {
		return result, EmptyGeometryError
	}
	return result, mapstructure.Decode(m.data, &result)
}

// Discriminator returns the GeometryDiscriminator oneOf value
func (m Geometry) Discriminator() GeometryDiscriminator {
	return m.data.GeometryDiscriminator()
}

// Validate implements basic validation for this model
func (m Geometry) Validate() error {
	if m.data == nil {
		return EmptyGeometryError
	}
	discriminator := m.data.GeometryDiscriminator()
	switch discriminator {
	case GeometryDiscriminatorLine:
		return m.data.Validate()
	case GeometryDiscriminatorShape:
		return m.data.Validate()
	default:
		return validation.Errors{
			"type": fmt.Errorf("unknown type value"),
		}.Filter()
	}
}

// IsGeometry tests if data is one of the discriminated sub-types of Geometry.
func IsGeometry(data interface{}) bool {
	t, ok := data.(geometryer)
	if !ok {
		return false
	}

	discriminator := t.GeometryDiscriminator()
	switch discriminator {
	case GeometryDiscriminatorLine:
		return true
	case GeometryDiscriminatorShape:
		return true
	default:
		return false
	}
}

// GeometryDiscriminator implements geometryer and returns the discriminator value.
func (m Line) GeometryDiscriminator() GeometryDiscriminator {
	value := fmt.Sprintf("%v", m.GetType())
	return GeometryDiscriminator(value)
}

// GeometryDiscriminator implements geometryer and returns the discriminator value.
func (m Shape) GeometryDiscriminator() GeometryDiscriminator {
	value := fmt.Sprintf("%v", m.GetType())
	return GeometryDiscriminator(value)
}
