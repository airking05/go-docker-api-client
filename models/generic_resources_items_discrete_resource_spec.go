// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GenericResourcesItemsDiscreteResourceSpec generic resources items discrete resource spec
// swagger:model genericResourcesItemsDiscreteResourceSpec
type GenericResourcesItemsDiscreteResourceSpec struct {

	// kind
	Kind string `json:"Kind,omitempty"`

	// value
	Value int64 `json:"Value,omitempty"`
}

// Validate validates this generic resources items discrete resource spec
func (m *GenericResourcesItemsDiscreteResourceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GenericResourcesItemsDiscreteResourceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericResourcesItemsDiscreteResourceSpec) UnmarshalBinary(b []byte) error {
	var res GenericResourcesItemsDiscreteResourceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}