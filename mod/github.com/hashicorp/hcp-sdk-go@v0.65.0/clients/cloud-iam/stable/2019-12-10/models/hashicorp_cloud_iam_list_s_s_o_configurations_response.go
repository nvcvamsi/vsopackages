// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudIamListSSOConfigurationsResponse hashicorp cloud iam list s s o configurations response
//
// swagger:model hashicorp.cloud.iam.ListSSOConfigurationsResponse
type HashicorpCloudIamListSSOConfigurationsResponse struct {

	// Config holds the SSO configurations for this organization.
	Configurations []*HashicorpCloudIamSSOConfig `json:"configurations"`

	// Pagination contains the pagination tokens for a subsequent request.
	Pagination *cloud.HashicorpCloudCommonPaginationResponse `json:"pagination,omitempty"`
}

// Validate validates this hashicorp cloud iam list s s o configurations response
func (m *HashicorpCloudIamListSSOConfigurationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfigurations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamListSSOConfigurationsResponse) validateConfigurations(formats strfmt.Registry) error {
	if swag.IsZero(m.Configurations) { // not required
		return nil
	}

	for i := 0; i < len(m.Configurations); i++ {
		if swag.IsZero(m.Configurations[i]) { // not required
			continue
		}

		if m.Configurations[i] != nil {
			if err := m.Configurations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configurations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudIamListSSOConfigurationsResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud iam list s s o configurations response based on the context it is used
func (m *HashicorpCloudIamListSSOConfigurationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfigurations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamListSSOConfigurationsResponse) contextValidateConfigurations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Configurations); i++ {

		if m.Configurations[i] != nil {
			if err := m.Configurations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("configurations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("configurations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudIamListSSOConfigurationsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamListSSOConfigurationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamListSSOConfigurationsResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamListSSOConfigurationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}