// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskSpecContainerSpec Invalid when specified with `PluginSpec`.
// swagger:model taskSpecContainerSpec
type TaskSpecContainerSpec struct {

	// Arguments to the command.
	Args []string `json:"Args"`

	// The command to be run in the image.
	Command []string `json:"Command"`

	// configs
	Configs TaskSpecContainerSpecConfigs `json:"Configs"`

	// DNS config
	DNSConfig *TaskSpecContainerSpecDNSConfig `json:"DNSConfig,omitempty"`

	// The working directory for commands to run in.
	Dir string `json:"Dir,omitempty"`

	// A list of environment variables in the form `VAR=value`.
	Env []string `json:"Env"`

	// A list of additional groups that the container process will run as.
	Groups []string `json:"Groups"`

	// health check
	HealthCheck *HealthConfig `json:"HealthCheck,omitempty"`

	// The hostname to use for the container, as a valid RFC 1123 hostname.
	Hostname string `json:"Hostname,omitempty"`

	// A list of hostname/IP mappings to add to the container's `hosts`
	// file. The format of extra hosts is specified in the
	// [hosts(5)](http://man7.org/linux/man-pages/man5/hosts.5.html)
	// man page:
	//
	//     IP_address canonical_hostname [aliases...]
	//
	Hosts []string `json:"Hosts"`

	// The image name to use for the container
	Image string `json:"Image,omitempty"`

	// Isolation technology of the containers running the service. (Windows only)
	Isolation string `json:"Isolation,omitempty"`

	// User-defined key/value data.
	Labels map[string]string `json:"Labels,omitempty"`

	// mounts
	Mounts TaskSpecContainerSpecMounts `json:"Mounts"`

	// Open `stdin`
	OpenStdin bool `json:"OpenStdin,omitempty"`

	// privileges
	Privileges *TaskSpecContainerSpecPrivileges `json:"Privileges,omitempty"`

	// Mount the container's root filesystem as read only.
	ReadOnly bool `json:"ReadOnly,omitempty"`

	// secrets
	Secrets TaskSpecContainerSpecSecrets `json:"Secrets"`

	// Amount of time to wait for the container to terminate before forcefully killing it.
	StopGracePeriod int64 `json:"StopGracePeriod,omitempty"`

	// Signal to stop the container.
	StopSignal string `json:"StopSignal,omitempty"`

	// Whether a pseudo-TTY should be allocated.
	TTY bool `json:"TTY,omitempty"`

	// The user inside the container.
	User string `json:"User,omitempty"`
}

// Validate validates this task spec container spec
func (m *TaskSpecContainerSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCommand(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDNSConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHealthCheck(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHosts(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIsolation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrivileges(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskSpecContainerSpec) validateArgs(formats strfmt.Registry) error {

	if swag.IsZero(m.Args) { // not required
		return nil
	}

	return nil
}

func (m *TaskSpecContainerSpec) validateCommand(formats strfmt.Registry) error {

	if swag.IsZero(m.Command) { // not required
		return nil
	}

	return nil
}

func (m *TaskSpecContainerSpec) validateDNSConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSConfig) { // not required
		return nil
	}

	if m.DNSConfig != nil {

		if err := m.DNSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("DNSConfig")
			}
			return err
		}

	}

	return nil
}

func (m *TaskSpecContainerSpec) validateEnv(formats strfmt.Registry) error {

	if swag.IsZero(m.Env) { // not required
		return nil
	}

	return nil
}

func (m *TaskSpecContainerSpec) validateGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.Groups) { // not required
		return nil
	}

	return nil
}

func (m *TaskSpecContainerSpec) validateHealthCheck(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthCheck) { // not required
		return nil
	}

	if m.HealthCheck != nil {

		if err := m.HealthCheck.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("HealthCheck")
			}
			return err
		}

	}

	return nil
}

func (m *TaskSpecContainerSpec) validateHosts(formats strfmt.Registry) error {

	if swag.IsZero(m.Hosts) { // not required
		return nil
	}

	return nil
}

var taskSpecContainerSpecTypeIsolationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","process","hyperv"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskSpecContainerSpecTypeIsolationPropEnum = append(taskSpecContainerSpecTypeIsolationPropEnum, v)
	}
}

const (

	// TaskSpecContainerSpecIsolationDefault captures enum value "default"
	TaskSpecContainerSpecIsolationDefault string = "default"

	// TaskSpecContainerSpecIsolationProcess captures enum value "process"
	TaskSpecContainerSpecIsolationProcess string = "process"

	// TaskSpecContainerSpecIsolationHyperv captures enum value "hyperv"
	TaskSpecContainerSpecIsolationHyperv string = "hyperv"
)

// prop value enum
func (m *TaskSpecContainerSpec) validateIsolationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, taskSpecContainerSpecTypeIsolationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TaskSpecContainerSpec) validateIsolation(formats strfmt.Registry) error {

	if swag.IsZero(m.Isolation) { // not required
		return nil
	}

	// value enum
	if err := m.validateIsolationEnum("Isolation", "body", m.Isolation); err != nil {
		return err
	}

	return nil
}

func (m *TaskSpecContainerSpec) validatePrivileges(formats strfmt.Registry) error {

	if swag.IsZero(m.Privileges) { // not required
		return nil
	}

	if m.Privileges != nil {

		if err := m.Privileges.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Privileges")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskSpecContainerSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskSpecContainerSpec) UnmarshalBinary(b []byte) error {
	var res TaskSpecContainerSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
