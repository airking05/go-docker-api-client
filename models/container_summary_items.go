// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContainerSummaryItems container summary items
// swagger:model containerSummaryItems
type ContainerSummaryItems struct {

	// Command to run when starting the container
	Command string `json:"Command,omitempty"`

	// When the container was created
	Created int64 `json:"Created,omitempty"`

	// host config
	HostConfig *ContainerSummaryItemsHostConfig `json:"HostConfig,omitempty"`

	// The ID of this container
	ID string `json:"Id,omitempty"`

	// The name of the image used when creating this container
	Image string `json:"Image,omitempty"`

	// The ID of the image that this container was created from
	ImageID string `json:"ImageID,omitempty"`

	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels,omitempty"`

	// mounts
	Mounts ContainerSummaryItemsMounts `json:"Mounts"`

	// The names that this container has been given
	Names []string `json:"Names"`

	// network settings
	NetworkSettings *ContainerSummaryItemsNetworkSettings `json:"NetworkSettings,omitempty"`

	// ports
	Ports ContainerSummaryItemsPorts `json:"Ports"`

	// The total size of all the files in this container
	SizeRootFs int64 `json:"SizeRootFs,omitempty"`

	// The size of files that have been created or changed by this container
	SizeRw int64 `json:"SizeRw,omitempty"`

	// The state of this container (e.g. `Exited`)
	State string `json:"State,omitempty"`

	// Additional human-readable status of this container (e.g. `Exit 0`)
	Status string `json:"Status,omitempty"`
}

// Validate validates this container summary items
func (m *ContainerSummaryItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNames(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworkSettings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContainerSummaryItems) validateHostConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.HostConfig) { // not required
		return nil
	}

	if m.HostConfig != nil {

		if err := m.HostConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HostConfig")
			}
			return err
		}

	}

	return nil
}

func (m *ContainerSummaryItems) validateNames(formats strfmt.Registry) error {

	if swag.IsZero(m.Names) { // not required
		return nil
	}

	return nil
}

func (m *ContainerSummaryItems) validateNetworkSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkSettings) { // not required
		return nil
	}

	if m.NetworkSettings != nil {

		if err := m.NetworkSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("NetworkSettings")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContainerSummaryItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerSummaryItems) UnmarshalBinary(b []byte) error {
	var res ContainerSummaryItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
