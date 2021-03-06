// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MountVolumeOptionsDriverConfig Map of driver specific options
// swagger:model mountVolumeOptionsDriverConfig
type MountVolumeOptionsDriverConfig struct {

	// Name of the driver to use to create the volume.
	Name string `json:"Name,omitempty"`

	// key/value map of driver specific options.
	Options map[string]string `json:"Options,omitempty"`
}

// Validate validates this mount volume options driver config
func (m *MountVolumeOptionsDriverConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MountVolumeOptionsDriverConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MountVolumeOptionsDriverConfig) UnmarshalBinary(b []byte) error {
	var res MountVolumeOptionsDriverConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
