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

// HashicorpCloudIamCreateOrganizationServicePrincipalResponse CreateOrganizationServicePrincipalResponse is the response message returned after creating
// a service principal on organization level.
//
// swagger:model hashicorp.cloud.iam.CreateOrganizationServicePrincipalResponse
type HashicorpCloudIamCreateOrganizationServicePrincipalResponse struct {

	// service_principal is the just-created service principal.
	ServicePrincipal *HashicorpCloudIamServicePrincipal `json:"service_principal,omitempty"`
}

// Validate validates this hashicorp cloud iam create organization service principal response
func (m *HashicorpCloudIamCreateOrganizationServicePrincipalResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServicePrincipal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamCreateOrganizationServicePrincipalResponse) validateServicePrincipal(formats strfmt.Registry) error {
	if swag.IsZero(m.ServicePrincipal) { // not required
		return nil
	}

	if m.ServicePrincipal != nil {
		if err := m.ServicePrincipal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_principal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_principal")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud iam create organization service principal response based on the context it is used
func (m *HashicorpCloudIamCreateOrganizationServicePrincipalResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServicePrincipal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamCreateOrganizationServicePrincipalResponse) contextValidateServicePrincipal(ctx context.Context, formats strfmt.Registry) error {

	if m.ServicePrincipal != nil {
		if err := m.ServicePrincipal.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_principal")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("service_principal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamCreateOrganizationServicePrincipalResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamCreateOrganizationServicePrincipalResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamCreateOrganizationServicePrincipalResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
