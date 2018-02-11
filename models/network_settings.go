// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NetworkSettings NetworkSettings exposes the network settings in the API
// swagger:model NetworkSettings
type NetworkSettings struct {

	// Name of the network'a bridge (for example, `docker0`).
	Bridge string `json:"Bridge,omitempty"`

	// EndpointID uniquely represents a service endpoint in a Sandbox.
	//
	// <p><br /></p>
	//
	// > **Deprecated**: This field is only propagated when attached to the
	// > default "bridge" network. Use the information from the "bridge"
	// > network inside the `Networks` map instead, which contains the same
	// > information. This field was deprecated in Docker 1.9 and is scheduled
	// > to be removed in Docker 17.12.0
	//
	EndpointID string `json:"EndpointID,omitempty"`

	// Gateway address for the default "bridge" network.
	//
	// <p><br /></p>
	//
	// > **Deprecated**: This field is only propagated when attached to the
	// > default "bridge" network. Use the information from the "bridge"
	// > network inside the `Networks` map instead, which contains the same
	// > information. This field was deprecated in Docker 1.9 and is scheduled
	// > to be removed in Docker 17.12.0
	//
	Gateway string `json:"Gateway,omitempty"`

	// Global IPv6 address for the default "bridge" network.
	//
	// <p><br /></p>
	//
	// > **Deprecated**: This field is only propagated when attached to the
	// > default "bridge" network. Use the information from the "bridge"
	// > network inside the `Networks` map instead, which contains the same
	// > information. This field was deprecated in Docker 1.9 and is scheduled
	// > to be removed in Docker 17.12.0
	//
	GlobalIPV6Address string `json:"GlobalIPv6Address,omitempty"`

	// Mask length of the global IPv6 address.
	//
	// <p><br /></p>
	//
	// > **Deprecated**: This field is only propagated when attached to the
	// > default "bridge" network. Use the information from the "bridge"
	// > network inside the `Networks` map instead, which contains the same
	// > information. This field was deprecated in Docker 1.9 and is scheduled
	// > to be removed in Docker 17.12.0
	//
	GlobalIPV6PrefixLen int64 `json:"GlobalIPv6PrefixLen,omitempty"`

	// Indicates if hairpin NAT should be enabled on the virtual interface.
	//
	HairpinMode bool `json:"HairpinMode,omitempty"`

	// IPv4 address for the default "bridge" network.
	//
	// <p><br /></p>
	//
	// > **Deprecated**: This field is only propagated when attached to the
	// > default "bridge" network. Use the information from the "bridge"
	// > network inside the `Networks` map instead, which contains the same
	// > information. This field was deprecated in Docker 1.9 and is scheduled
	// > to be removed in Docker 17.12.0
	//
	IPAddress string `json:"IPAddress,omitempty"`

	// Mask length of the IPv4 address.
	//
	// <p><br /></p>
	//
	// > **Deprecated**: This field is only propagated when attached to the
	// > default "bridge" network. Use the information from the "bridge"
	// > network inside the `Networks` map instead, which contains the same
	// > information. This field was deprecated in Docker 1.9 and is scheduled
	// > to be removed in Docker 17.12.0
	//
	IPPrefixLen int64 `json:"IPPrefixLen,omitempty"`

	// IPv6 gateway address for this network.
	//
	// <p><br /></p>
	//
	// > **Deprecated**: This field is only propagated when attached to the
	// > default "bridge" network. Use the information from the "bridge"
	// > network inside the `Networks` map instead, which contains the same
	// > information. This field was deprecated in Docker 1.9 and is scheduled
	// > to be removed in Docker 17.12.0
	//
	IPV6Gateway string `json:"IPv6Gateway,omitempty"`

	// IPv6 unicast address using the link-local prefix.
	LinkLocalIPV6Address string `json:"LinkLocalIPv6Address,omitempty"`

	// Prefix length of the IPv6 unicast address.
	LinkLocalIPV6PrefixLen int64 `json:"LinkLocalIPv6PrefixLen,omitempty"`

	// MAC address for the container on the default "bridge" network.
	//
	// <p><br /></p>
	//
	// > **Deprecated**: This field is only propagated when attached to the
	// > default "bridge" network. Use the information from the "bridge"
	// > network inside the `Networks` map instead, which contains the same
	// > information. This field was deprecated in Docker 1.9 and is scheduled
	// > to be removed in Docker 17.12.0
	//
	MacAddress string `json:"MacAddress,omitempty"`

	// networks
	Networks NetworkSettingsNetworks `json:"Networks,omitempty"`

	// ports
	Ports PortMap `json:"Ports,omitempty"`

	// SandboxID uniquely represents a container's network stack.
	SandboxID string `json:"SandboxID,omitempty"`

	// SandboxKey identifies the sandbox
	SandboxKey string `json:"SandboxKey,omitempty"`

	// secondary IP addresses
	SecondaryIPAddresses *NetworkSettingsSecondaryIPAddresses `json:"SecondaryIPAddresses"`

	// secondary ipv6 addresses
	SecondaryIPV6Addresses *NetworkSettingsSecondaryIPV6Addresses `json:"SecondaryIPv6Addresses"`
}

// Validate validates this network settings
func (m *NetworkSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSettings) UnmarshalBinary(b []byte) error {
	var res NetworkSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}