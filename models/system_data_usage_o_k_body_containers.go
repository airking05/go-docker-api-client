// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// SystemDataUsageOKBodyContainers system data usage o k body containers
// swagger:model systemDataUsageOKBodyContainers
type SystemDataUsageOKBodyContainers []ContainerSummary

// Validate validates this system data usage o k body containers
func (m SystemDataUsageOKBodyContainers) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if err := m[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName(strconv.Itoa(i))
			}
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
