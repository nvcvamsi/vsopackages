// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPackerGetRegistryResponse hashicorp cloud packer get registry response
//
// swagger:model hashicorp.cloud.packer.GetRegistryResponse
type HashicorpCloudPackerGetRegistryResponse struct {

	// registry
	Registry *HashicorpCloudPackerRegistry `json:"registry,omitempty"`
}

// Validate validates this hashicorp cloud packer get registry response
func (m *HashicorpCloudPackerGetRegistryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegistry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerGetRegistryResponse) validateRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.Registry) { // not required
		return nil
	}

	if m.Registry != nil {
		if err := m.Registry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer get registry response based on the context it is used
func (m *HashicorpCloudPackerGetRegistryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerGetRegistryResponse) contextValidateRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.Registry != nil {
		if err := m.Registry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerGetRegistryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerGetRegistryResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerGetRegistryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}