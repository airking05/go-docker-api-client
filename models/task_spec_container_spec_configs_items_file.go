// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TaskSpecContainerSpecConfigsItemsFile File represents a specific target that is backed by a file.
// swagger:model taskSpecContainerSpecConfigsItemsFile
type TaskSpecContainerSpecConfigsItemsFile struct {

	// GID represents the file GID.
	GID string `json:"GID,omitempty"`

	// Mode represents the FileMode of the file.
	Mode uint32 `json:"Mode,omitempty"`

	// Name represents the final filename in the filesystem.
	Name string `json:"Name,omitempty"`

	// UID represents the file UID.
	UID string `json:"UID,omitempty"`
}

// Validate validates this task spec container spec configs items file
func (m *TaskSpecContainerSpecConfigsItemsFile) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TaskSpecContainerSpecConfigsItemsFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskSpecContainerSpecConfigsItemsFile) UnmarshalBinary(b []byte) error {
	var res TaskSpecContainerSpecConfigsItemsFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
