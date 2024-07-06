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
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// HashicorpCloudBoundary20211221Cluster Cluster represents a single HCP Boundary cluster.
//
// swagger:model hashicorp.cloud.boundary_20211221.Cluster
type HashicorpCloudBoundary20211221Cluster struct {

	// boundary_version is the version of the HCP Boundary running in this cluster.
	BoundaryVersion string `json:"boundary_version,omitempty"`

	// cluster_id is the id of the cluster set by user on creation.
	ClusterID string `json:"cluster_id,omitempty"`

	// cluster_url is the generated vanity url.
	ClusterURL string `json:"cluster_url,omitempty"`

	// created_at is the timestamp of the cluster creation.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// location is the location of the cluster.
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// marketing_sku is the marketing sku of the cluster [standard, plus]
	MarketingSku *HashicorpCloudBoundary20211221ClusterMarketingSKU `json:"marketing_sku,omitempty"`

	// state is the current state of the cluster.
	State *HashicorpCloudBoundary20211221ClusterState `json:"state,omitempty"`
}

// Validate validates this hashicorp cloud boundary 20211221 cluster
func (m *HashicorpCloudBoundary20211221Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarketingSku(formats); err != nil {
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

func (m *HashicorpCloudBoundary20211221Cluster) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudBoundary20211221Cluster) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudBoundary20211221Cluster) validateMarketingSku(formats strfmt.Registry) error {
	if swag.IsZero(m.MarketingSku) { // not required
		return nil
	}

	if m.MarketingSku != nil {
		if err := m.MarketingSku.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("marketing_sku")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("marketing_sku")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudBoundary20211221Cluster) validateState(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud boundary 20211221 cluster based on the context it is used
func (m *HashicorpCloudBoundary20211221Cluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMarketingSku(ctx, formats); err != nil {
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

func (m *HashicorpCloudBoundary20211221Cluster) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {
		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudBoundary20211221Cluster) contextValidateMarketingSku(ctx context.Context, formats strfmt.Registry) error {

	if m.MarketingSku != nil {
		if err := m.MarketingSku.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("marketing_sku")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("marketing_sku")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudBoundary20211221Cluster) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudBoundary20211221Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudBoundary20211221Cluster) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudBoundary20211221Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
