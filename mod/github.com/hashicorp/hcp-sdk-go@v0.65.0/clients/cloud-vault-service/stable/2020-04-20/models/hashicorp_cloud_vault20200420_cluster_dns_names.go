// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20200420ClusterDNSNames DNSNames holds all of the cluster's DNS names.
//
// swagger:model hashicorp.cloud.vault_20200420.Cluster.DNSNames
type HashicorpCloudVault20200420ClusterDNSNames struct {

	// private is the DNS name pointing to the cluster's private IP addresses.
	Private string `json:"private,omitempty"`

	// public is the DNS name pointing to the cluster's public IP addresses.
	Public string `json:"public,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 cluster DNS names
func (m *HashicorpCloudVault20200420ClusterDNSNames) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20200420 cluster DNS names based on context it is used
func (m *HashicorpCloudVault20200420ClusterDNSNames) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420ClusterDNSNames) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420ClusterDNSNames) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420ClusterDNSNames
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}