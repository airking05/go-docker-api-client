// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateImageInfo create image info
// swagger:model CreateImageInfo
type CreateImageInfo struct {

	// error
	Error string `json:"error,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// progress
	Progress string `json:"progress,omitempty"`

	// progress detail
	ProgressDetail *ProgressDetail `json:"progressDetail,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this create image info
func (m *CreateImageInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProgressDetail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateImageInfo) validateProgressDetail(formats strfmt.Registry) error {

	if swag.IsZero(m.ProgressDetail) { // not required
		return nil
	}

	if m.ProgressDetail != nil {

		if err := m.ProgressDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("progressDetail")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateImageInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateImageInfo) UnmarshalBinary(b []byte) error {
	var res CreateImageInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
