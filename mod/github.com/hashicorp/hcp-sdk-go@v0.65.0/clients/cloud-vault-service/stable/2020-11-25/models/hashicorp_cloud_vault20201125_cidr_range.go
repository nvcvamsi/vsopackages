// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20201125CidrRange CidrRange specifies an IP address and prefix in CIDR format.
//
// swagger:model hashicorp.cloud.vault_20201125.CidrRange
type HashicorpCloudVault20201125CidrRange struct {

	// The IPv4 address range including prefixes.
	// E.g. 172.25.16.0/24
	Address string `json:"address,omitempty"`

	// An optional description of the IP address range.
	Description string `json:"description,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 cidr range
func (m *HashicorpCloudVault20201125CidrRange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20201125 cidr range based on context it is used
func (m *HashicorpCloudVault20201125CidrRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125CidrRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125CidrRange) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125CidrRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}