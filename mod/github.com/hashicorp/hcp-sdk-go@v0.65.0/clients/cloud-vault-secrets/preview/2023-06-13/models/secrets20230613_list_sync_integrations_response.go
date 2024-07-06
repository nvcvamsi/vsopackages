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

// Secrets20230613ListSyncIntegrationsResponse secrets 20230613 list sync integrations response
//
// swagger:model secrets_20230613ListSyncIntegrationsResponse
type Secrets20230613ListSyncIntegrationsResponse struct {

	// sync integrations
	SyncIntegrations []*Secrets20230613SyncIntegration `json:"sync_integrations"`
}

// Validate validates this secrets 20230613 list sync integrations response
func (m *Secrets20230613ListSyncIntegrationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSyncIntegrations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20230613ListSyncIntegrationsResponse) validateSyncIntegrations(formats strfmt.Registry) error {
	if swag.IsZero(m.SyncIntegrations) { // not required
		return nil
	}

	for i := 0; i < len(m.SyncIntegrations); i++ {
		if swag.IsZero(m.SyncIntegrations[i]) { // not required
			continue
		}

		if m.SyncIntegrations[i] != nil {
			if err := m.SyncIntegrations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sync_integrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sync_integrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this secrets 20230613 list sync integrations response based on the context it is used
func (m *Secrets20230613ListSyncIntegrationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSyncIntegrations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Secrets20230613ListSyncIntegrationsResponse) contextValidateSyncIntegrations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SyncIntegrations); i++ {

		if m.SyncIntegrations[i] != nil {
			if err := m.SyncIntegrations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sync_integrations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sync_integrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Secrets20230613ListSyncIntegrationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Secrets20230613ListSyncIntegrationsResponse) UnmarshalBinary(b []byte) error {
	var res Secrets20230613ListSyncIntegrationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}