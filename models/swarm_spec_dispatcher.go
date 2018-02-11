// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SwarmSpecDispatcher Dispatcher configuration.
// swagger:model swarmSpecDispatcher
type SwarmSpecDispatcher struct {

	// The delay for an agent to send a heartbeat to the dispatcher.
	HeartbeatPeriod int64 `json:"HeartbeatPeriod,omitempty"`
}

// Validate validates this swarm spec dispatcher
func (m *SwarmSpecDispatcher) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SwarmSpecDispatcher) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwarmSpecDispatcher) UnmarshalBinary(b []byte) error {
	var res SwarmSpecDispatcher
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}