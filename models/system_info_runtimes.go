// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SystemInfoRuntimes List of [OCI compliant](https://github.com/opencontainers/runtime-spec)
// runtimes configured on the daemon. Keys hold the "name" used to
// reference the runtime.
//
// The Docker daemon relies on an OCI compliant runtime (invoked via the
// `containerd` daemon) as its interface to the Linux kernel namespaces,
// cgroups, and SELinux.
//
// The default runtime is `runc`, and automatically configured. Additional
// runtimes can be configured by the user and will be listed here.
//
// swagger:model systemInfoRuntimes
type SystemInfoRuntimes map[string]Runtime

// Validate validates this system info runtimes
func (m SystemInfoRuntimes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", SystemInfoRuntimes(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}