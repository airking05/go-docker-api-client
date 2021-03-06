// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SwarmSpecTaskDefaultsLogDriver The log driver to use for tasks created in the orchestrator if
// unspecified by a service.
//
// Updating this value only affects new tasks. Existing tasks continue
// to use their previously configured log driver until recreated.
//
// swagger:model swarmSpecTaskDefaultsLogDriver
type SwarmSpecTaskDefaultsLogDriver struct {

	// The log driver to use as a default for new tasks.
	//
	Name string `json:"Name,omitempty"`

	// Driver-specific options for the selectd log driver, specified
	// as key/value pairs.
	//
	Options map[string]string `json:"Options,omitempty"`
}

// Validate validates this swarm spec task defaults log driver
func (m *SwarmSpecTaskDefaultsLogDriver) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SwarmSpecTaskDefaultsLogDriver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwarmSpecTaskDefaultsLogDriver) UnmarshalBinary(b []byte) error {
	var res SwarmSpecTaskDefaultsLogDriver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
