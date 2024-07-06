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

// HashicorpCloudGlobalNetworkManager20220215ServerState hashicorp cloud global network manager 20220215 server state
//
// swagger:model hashicorp.cloud.global_network_manager_20220215.ServerState
type HashicorpCloudGlobalNetworkManager20220215ServerState struct {

	// acl
	ACL *HashicorpCloudGlobalNetworkManager20220215ACLInfo `json:"acl,omitempty"`

	// autopilot
	Autopilot *HashicorpCloudGlobalNetworkManager20220215AutoPilotInfo `json:"autopilot,omitempty"`

	// datacenter
	Datacenter string `json:"datacenter,omitempty"`

	// gossip port
	GossipPort int32 `json:"gossip_port,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// lan address
	LanAddress string `json:"lan_address,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// raft
	Raft *HashicorpCloudGlobalNetworkManager20220215RaftInfo `json:"raft,omitempty"`

	// rpc port
	RPCPort int32 `json:"rpc_port,omitempty"`

	// scada status
	ScadaStatus string `json:"scada_status,omitempty"`

	// tls
	TLS *HashicorpCloudGlobalNetworkManager20220215TLSInfo `json:"tls,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this hashicorp cloud global network manager 20220215 server state
func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateACL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutopilot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRaft(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) validateACL(formats strfmt.Registry) error {
	if swag.IsZero(m.ACL) { // not required
		return nil
	}

	if m.ACL != nil {
		if err := m.ACL.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acl")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) validateAutopilot(formats strfmt.Registry) error {
	if swag.IsZero(m.Autopilot) { // not required
		return nil
	}

	if m.Autopilot != nil {
		if err := m.Autopilot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autopilot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autopilot")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) validateRaft(formats strfmt.Registry) error {
	if swag.IsZero(m.Raft) { // not required
		return nil
	}

	if m.Raft != nil {
		if err := m.Raft.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raft")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("raft")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) validateTLS(formats strfmt.Registry) error {
	if swag.IsZero(m.TLS) { // not required
		return nil
	}

	if m.TLS != nil {
		if err := m.TLS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud global network manager 20220215 server state based on the context it is used
func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateACL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAutopilot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRaft(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) contextValidateACL(ctx context.Context, formats strfmt.Registry) error {

	if m.ACL != nil {
		if err := m.ACL.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acl")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("acl")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) contextValidateAutopilot(ctx context.Context, formats strfmt.Registry) error {

	if m.Autopilot != nil {
		if err := m.Autopilot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autopilot")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autopilot")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) contextValidateRaft(ctx context.Context, formats strfmt.Registry) error {

	if m.Raft != nil {
		if err := m.Raft.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("raft")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("raft")
			}
			return err
		}
	}

	return nil
}

func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) contextValidateTLS(ctx context.Context, formats strfmt.Registry) error {

	if m.TLS != nil {
		if err := m.TLS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudGlobalNetworkManager20220215ServerState) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudGlobalNetworkManager20220215ServerState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}