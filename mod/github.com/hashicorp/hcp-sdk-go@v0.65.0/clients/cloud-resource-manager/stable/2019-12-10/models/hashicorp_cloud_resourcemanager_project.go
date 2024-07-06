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

// HashicorpCloudResourcemanagerProject Project is a container for resources.
//
// swagger:model hashicorp.cloud.resourcemanager.Project
type HashicorpCloudResourcemanagerProject struct {

	// CreatedAt is the time the project was created at.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// Description is the user settable description of the project.
	Description string `json:"description,omitempty"`

	// ID is an UUID identifying the project object.
	ID string `json:"id,omitempty"`

	// Name is the user settable name of the project.
	Name string `json:"name,omitempty"`

	// Parent is the parent resource of the project.
	Parent *HashicorpCloudResourcemanagerResourceID `json:"parent,omitempty"`

	// State is the state of the project.
	State *HashicorpCloudResourcemanagerProjectProjectState `json:"state,omitempty"`
}

// Validate validates this hashicorp cloud resourcemanager project
func (m *HashicorpCloudResourcemanagerProject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudResourcemanagerProject) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudResourcemanagerProject) validateParent(formats strfmt.Registry) error {
	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudResourcemanagerProject) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud resourcemanager project based on the context it is used
func (m *HashicorpCloudResourcemanagerProject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudResourcemanagerProject) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

	if m.Parent != nil {
		if err := m.Parent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudResourcemanagerProject) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerProject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudResourcemanagerProject) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudResourcemanagerProject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
