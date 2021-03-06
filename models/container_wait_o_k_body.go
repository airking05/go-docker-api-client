// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContainerWaitOKBody ContainerWaitResponse
//
// OK response to ContainerWait operation
// swagger:model containerWaitOKBody
type ContainerWaitOKBody struct {

	// error
	Error *ContainerWaitOKBodyError `json:"Error,omitempty"`

	// Exit code of the container
	// Required: true
	StatusCode int64 `json:"StatusCode"`
}

// Validate validates this container wait o k body
func (m *ContainerWaitOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatusCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerWaitOKBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {

		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Error")
			}
			return err
		}

	}

	return nil
}

func (m *ContainerWaitOKBody) validateStatusCode(formats strfmt.Registry) error {

	if err := validate.Required("StatusCode", "body", int64(m.StatusCode)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerWaitOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerWaitOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerWaitOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
