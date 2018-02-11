// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SystemDataUsageOKBody SystemDataUsageResponse
// swagger:model systemDataUsageOKBody
type SystemDataUsageOKBody struct {

	// containers
	Containers SystemDataUsageOKBodyContainers `json:"Containers"`

	// images
	Images SystemDataUsageOKBodyImages `json:"Images"`

	// layers size
	LayersSize int64 `json:"LayersSize,omitempty"`

	// volumes
	Volumes SystemDataUsageOKBodyVolumes `json:"Volumes"`
}

// Validate validates this system data usage o k body
func (m *SystemDataUsageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SystemDataUsageOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemDataUsageOKBody) UnmarshalBinary(b []byte) error {
	var res SystemDataUsageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
