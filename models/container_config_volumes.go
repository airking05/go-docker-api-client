// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContainerConfigVolumes An object mapping mount point paths inside the container to empty objects.
// swagger:model containerConfigVolumes
type ContainerConfigVolumes struct {

	// additional properties
	AdditionalProperties interface{} `json:"additionalProperties,omitempty"`
}

// Validate validates this container config volumes
func (m *ContainerConfigVolumes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var containerConfigVolumesTypeAdditionalPropertiesPropEnum []interface{}

func init() {
	var res []interface{}
	if err := json.Unmarshal([]byte(`[{}]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerConfigVolumesTypeAdditionalPropertiesPropEnum = append(containerConfigVolumesTypeAdditionalPropertiesPropEnum, v)
	}
}

// prop value enum
func (m *ContainerConfigVolumes) validateAdditionalPropertiesEnum(path, location string, value interface{}) error {
	if err := validate.Enum(path, location, value, containerConfigVolumesTypeAdditionalPropertiesPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ContainerConfigVolumes) validateAdditionalProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.AdditionalProperties) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerConfigVolumes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerConfigVolumes) UnmarshalBinary(b []byte) error {
	var res ContainerConfigVolumes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
