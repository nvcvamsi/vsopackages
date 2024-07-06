// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudVault20200420GetCORSConfigResponse hashicorp cloud vault 20200420 get c o r s config response
//
// swagger:model hashicorp.cloud.vault_20200420.GetCORSConfigResponse
type HashicorpCloudVault20200420GetCORSConfigResponse struct {

	// allowed_headers are the allowed headers.
	AllowedHeaders []string `json:"allowed_headers"`

	// allowed_origins are the allowed origins.
	AllowedOrigins []string `json:"allowed_origins"`

	// enabled returns true if CORS is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this hashicorp cloud vault 20200420 get c o r s config response
func (m *HashicorpCloudVault20200420GetCORSConfigResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud vault 20200420 get c o r s config response based on context it is used
func (m *HashicorpCloudVault20200420GetCORSConfigResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20200420GetCORSConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20200420GetCORSConfigResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20200420GetCORSConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}