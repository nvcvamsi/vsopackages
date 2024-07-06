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

// HashicorpCloudGlobalNetworkManager20220215Cluster hashicorp cloud global network manager 20220215 cluster
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.Cluster
type HashicorpCloudGlobalNetworkManager20220215Cluster struct {

	// active_server_count is the number of servers that have checked in and are reported as running
	ActiveServerCount int32 `json:"active_server_count,omitempty"`

	// bootstrap expect
	BootstrapExpect int32 `json:"bootstrap_expect,omitempty"`

	// consul_access_level is the level of access CCM has to the Consul cluster
	ConsulAccessLevel *HashicorpCloudGlobalNetworkManager20220215ClusterConsulAccessLevel `json:"consul_access_level,omitempty"`

	// Consul version of the cluster
	ConsulVersion string `json:"consul_version,omitempty"`

	// `created_at` is the timestamp of when the cluster was first created.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// whether this is an existing self-managed cluster
	ExistingCluster bool `json:"existing_cluster,omitempty"`

	// whether this Consul cluster is managed by HCP
	HcpManaged bool `json:"hcp_managed,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// licensing is the Consul licensing information
	Licensing *HashicorpCloudGlobalNetworkManager20220215Licensing `json:"licensing,omitempty"`

	// location
	Location *cloud.HashicorpCloudLocationLocation `json:"location,omitempty"`

	// resource_id is the unique identifier for the cluster
	ResourceID string `json:"resource_id,omitempty"`

	// resource_link_url is the location.Link encoded as a url including the organization, project, type and ID
	ResourceLinkURL string `json:"resource_link_url,omitempty"`

	// source of the cluster
	Source *HashicorpCloudGlobalNetworkManager20220215ClusterSource `json:"source,omitempty"`

	// Current state of the cluster.
	State *HashicorpCloudGlobalNetworkManager20220215ClusterState `json:"state,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 cluster
func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsulAccessLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicensing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
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

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) validateConsulAccessLevel(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsulAccessLevel) { // not required
		return nil
	}

	if m.ConsulAccessLevel != nil {
		if err := m.ConsulAccessLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consul_access_level")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consul_access_level")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) validateLicensing(formats strfmt.Registry) error {
	if swag.IsZero(m.Licensing) { // not required
		return nil
	}

	if m.Licensing != nil {
		if err := m.Licensing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licensing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licensing")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) validateLocation(formats strfmt.Registry) error {
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

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) validateSource(formats strfmt.Registry) error {
	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) validateState(formats strfmt.Registry) error {
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

// ContextValidate validate this hashicorp cloud global network manager 20220215 cluster based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsulAccessLevel(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLicensing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
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

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) contextValidateConsulAccessLevel(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsulAccessLevel != nil {
		if err := m.ConsulAccessLevel.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consul_access_level")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consul_access_level")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) contextValidateLicensing(ctx context.Context, formats strfmt.Registry) error {

	if m.Licensing != nil {
		if err := m.Licensing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licensing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("licensing")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {
		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

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
func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215Cluster) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}