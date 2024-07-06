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
)

// HashicorpCloudIamCreateOrganizationInvitationsRequest CreateOrganizationInvitationsRequest is a request to create an organization invitation.
// Inviting multiple people at the same time is supported.
//
// swagger:model hashicorp.cloud.iam.CreateOrganizationInvitationsRequest
type HashicorpCloudIamCreateOrganizationInvitationsRequest struct {

	// invitations is a list of the data for the invitations that should be created.
	Invitations []*HashicorpCloudIamNewOrganizationInvitation `json:"invitations"`

	// organization_id is the unique identifier (UUID) of the organization
	// to which the invitee is being invited.
	OrganizationID string `json:"organization_id,omitempty"`
}

// Validate validates this hashicorp cloud iam create organization invitations request
func (m *HashicorpCloudIamCreateOrganizationInvitationsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInvitations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamCreateOrganizationInvitationsRequest) validateInvitations(formats strfmt.Registry) error {
	if swag.IsZero(m.Invitations) { // not required
		return nil
	}

	for i := 0; i < len(m.Invitations); i++ {
		if swag.IsZero(m.Invitations[i]) { // not required
			continue
		}

		if m.Invitations[i] != nil {
			if err := m.Invitations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("invitations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("invitations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this hashicorp cloud iam create organization invitations request based on the context it is used
func (m *HashicorpCloudIamCreateOrganizationInvitationsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInvitations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudIamCreateOrganizationInvitationsRequest) contextValidateInvitations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Invitations); i++ {

		if m.Invitations[i] != nil {
			if err := m.Invitations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("invitations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("invitations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudIamCreateOrganizationInvitationsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudIamCreateOrganizationInvitationsRequest) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudIamCreateOrganizationInvitationsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
