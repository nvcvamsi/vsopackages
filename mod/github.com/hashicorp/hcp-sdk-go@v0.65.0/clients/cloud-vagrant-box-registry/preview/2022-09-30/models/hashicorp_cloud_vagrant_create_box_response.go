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

// HashicorpCloudVagrantCreateBoxResponse hashicorp cloud vagrant create box response
//
// swagger:model hashicorp.cloud.vagrant.CreateBoxResponse
type HashicorpCloudVagrantCreateBoxResponse struct {

	// The created Box.
	Box *HashicorpCloudVagrantBox `json:"box,omitempty"`
}

// Validate validates this hashicorp cloud vagrant create box response
func (m *HashicorpCloudVagrantCreateBoxResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBox(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrantCreateBoxResponse) validateBox(formats strfmt.Registry) error {
	if swag.IsZero(m.Box) { // not required
		return nil
	}

	if m.Box != nil {
		if err := m.Box.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("box")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("box")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud vagrant create box response based on the context it is used
func (m *HashicorpCloudVagrantCreateBoxResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBox(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudVagrantCreateBoxResponse) contextValidateBox(ctx context.Context, formats strfmt.Registry) error {

	if m.Box != nil {
		if err := m.Box.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("box")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("box")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudVagrantCreateBoxResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudVagrantCreateBoxResponse) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudVagrantCreateBoxResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
