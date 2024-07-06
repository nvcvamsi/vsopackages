// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HashicorpCloudVault20201125GetUtilizationResponseValue hashicorp cloud vault 20201125 get utilization response value
//
// swagger:model hashicorp.cloud.vault_20201125.GetUtilizationResponse.Value
type HashicorpCloudVault20201125GetUtilizationResponseValue struct {

	// max value
	MaxValue float32 `json:"max_value,omitempty"`

	// unit
	Unit *HashicorpCloudVault20201125GetUtilizationResponseValueUnit `json:"unit,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// value
	Value float32 `json:"value,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 get utilization response value
func (m *HashicorpCloudVault20201125GetUtilizationResponseValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125GetUtilizationResponseValue) validateUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.Unit) { // not required
		return nil
	}

	if m.Unit != nil {
		if err := m.Unit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125GetUtilizationResponseValue) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20201125 get utilization response value based on the context it is used
func (m *HashicorpCloudVault20201125GetUtilizationResponseValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125GetUtilizationResponseValue) contextValidateUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.Unit != nil {
		if err := m.Unit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetUtilizationResponseValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetUtilizationResponseValue) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125GetUtilizationResponseValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
