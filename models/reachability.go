// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Reachability Reachability represents the reachability of a node.
// swagger:model Reachability
type Reachability string

const (

	// ReachabilityUnknown captures enum value "unknown"
	ReachabilityUnknown Reachability = "unknown"

	// ReachabilityUnreachable captures enum value "unreachable"
	ReachabilityUnreachable Reachability = "unreachable"

	// ReachabilityReachable captures enum value "reachable"
	ReachabilityReachable Reachability = "reachable"
)

// for schema
var reachabilityEnum []interface{}

func init() {
	var res []Reachability
	if err := json.Unmarshal([]byte(`["unknown","unreachable","reachable"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reachabilityEnum = append(reachabilityEnum, v)
	}
}

func (m Reachability) validateReachabilityEnum(path, location string, value Reachability) error {
	if err := validate.Enum(path, location, value, reachabilityEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this reachability
func (m Reachability) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReachabilityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}