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

// HashicorpCloudVault20201125GetCurrentMilestoneResponse hashicorp cloud vault 20201125 get current milestone response
//
// swagger:model hashicorp.cloud.vault_20201125.GetCurrentMilestoneResponse
type HashicorpCloudVault20201125GetCurrentMilestoneResponse struct {

	// created_at is the timestamp of when the cluster milestone was first created
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// current milestone
	CurrentMilestone *HashicorpCloudVault20201125GetCurrentMilestoneResponseCurrentMilestone `json:"current_milestone,omitempty"`

	// updated_at is the timestamp of when the cluster milestone was last updated
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this hashicorp cloud vault 20201125 get current milestone response
func (m *HashicorpCloudVault20201125GetCurrentMilestoneResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentMilestone(formats); err != nil {
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

func (m *HashicorpCloudVault20201125GetCurrentMilestoneResponse) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudVault20201125GetCurrentMilestoneResponse) validateCurrentMilestone(formats strfmt.Registry) error {
	if swag.IsZero(m.CurrentMilestone) { // not required
		return nil
	}

	if m.CurrentMilestone != nil {
		if err := m.CurrentMilestone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_milestone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("current_milestone")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudVault20201125GetCurrentMilestoneResponse) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vault 20201125 get current milestone response based on the context it is used
func (m *HashicorpCloudVault20201125GetCurrentMilestoneResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCurrentMilestone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVault20201125GetCurrentMilestoneResponse) contextValidateCurrentMilestone(ctx context.Context, formats strfmt.Registry) error {

	if m.CurrentMilestone != nil {
		if err := m.CurrentMilestone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("current_milestone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("current_milestone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetCurrentMilestoneResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVault20201125GetCurrentMilestoneResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVault20201125GetCurrentMilestoneResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
