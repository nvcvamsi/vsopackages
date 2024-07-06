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

// HashicorpCloudGlobalNetworkManager20220215GetTrialResponse hashicorp cloud global network manager 20220215 get trial response
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.GetTrialResponse
type HashicorpCloudGlobalNetworkManager20220215GetTrialResponse struct {

	// trial
	Trial *HashicorpCloudGlobalNetworkManager20220215Trial `json:"trial,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 get trial response
func (m *HashicorpCloudGlobalNetworkManager20220215GetTrialResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrial(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215GetTrialResponse) validateTrial(formats strfmt.Registry) error {
	if swag.IsZero(m.Trial) { // not required
		return nil
	}

	if m.Trial != nil {
		if err := m.Trial.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trial")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trial")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 get trial response based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215GetTrialResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTrial(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215GetTrialResponse) contextValidateTrial(ctx context.Context, formats strfmt.Registry) error {

	if m.Trial != nil {
		if err := m.Trial.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trial")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trial")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215GetTrialResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215GetTrialResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215GetTrialResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
