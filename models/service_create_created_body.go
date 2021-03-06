// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ServiceCreateCreatedBody ServiceCreateResponse
// swagger:model serviceCreateCreatedBody
type ServiceCreateCreatedBody struct {

	// The ID of the created service.
	ID string `json:"ID,omitempty"`

	// Optional warning message
	Warning string `json:"Warning,omitempty"`
}

// Validate validates this service create created body
func (m *ServiceCreateCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceCreateCreatedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceCreateCreatedBody) UnmarshalBinary(b []byte) error {
	var res ServiceCreateCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
