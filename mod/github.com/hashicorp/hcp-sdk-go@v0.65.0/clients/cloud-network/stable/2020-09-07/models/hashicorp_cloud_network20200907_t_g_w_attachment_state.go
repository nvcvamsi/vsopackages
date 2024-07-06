// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HashicorpCloudNetwork20200907TGWAttachmentState state is one of the states the TGW Attachment could be in.
//
// Output only.
//
// swagger:model hashicorp.cloud.network_20200907.TGWAttachment.State
type HashicorpCloudNetwork20200907TGWAttachmentState string

func NewHashicorpCloudNetwork20200907TGWAttachmentState(value HashicorpCloudNetwork20200907TGWAttachmentState) *HashicorpCloudNetwork20200907TGWAttachmentState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudNetwork20200907TGWAttachmentState.
func (m HashicorpCloudNetwork20200907TGWAttachmentState) Pointer() *HashicorpCloudNetwork20200907TGWAttachmentState {
	return &m
}

const (

	// HashicorpCloudNetwork20200907TGWAttachmentStateUNSET captures enum value "UNSET"
	HashicorpCloudNetwork20200907TGWAttachmentStateUNSET HashicorpCloudNetwork20200907TGWAttachmentState = "UNSET"

	// HashicorpCloudNetwork20200907TGWAttachmentStateCREATING captures enum value "CREATING"
	HashicorpCloudNetwork20200907TGWAttachmentStateCREATING HashicorpCloudNetwork20200907TGWAttachmentState = "CREATING"

	// HashicorpCloudNetwork20200907TGWAttachmentStatePENDINGACCEPTANCE captures enum value "PENDING_ACCEPTANCE"
	HashicorpCloudNetwork20200907TGWAttachmentStatePENDINGACCEPTANCE HashicorpCloudNetwork20200907TGWAttachmentState = "PENDING_ACCEPTANCE"

	// HashicorpCloudNetwork20200907TGWAttachmentStateACCEPTED captures enum value "ACCEPTED"
	HashicorpCloudNetwork20200907TGWAttachmentStateACCEPTED HashicorpCloudNetwork20200907TGWAttachmentState = "ACCEPTED"

	// HashicorpCloudNetwork20200907TGWAttachmentStateACTIVE captures enum value "ACTIVE"
	HashicorpCloudNetwork20200907TGWAttachmentStateACTIVE HashicorpCloudNetwork20200907TGWAttachmentState = "ACTIVE"

	// HashicorpCloudNetwork20200907TGWAttachmentStateFAILED captures enum value "FAILED"
	HashicorpCloudNetwork20200907TGWAttachmentStateFAILED HashicorpCloudNetwork20200907TGWAttachmentState = "FAILED"

	// HashicorpCloudNetwork20200907TGWAttachmentStateREJECTED captures enum value "REJECTED"
	HashicorpCloudNetwork20200907TGWAttachmentStateREJECTED HashicorpCloudNetwork20200907TGWAttachmentState = "REJECTED"

	// HashicorpCloudNetwork20200907TGWAttachmentStateDELETING captures enum value "DELETING"
	HashicorpCloudNetwork20200907TGWAttachmentStateDELETING HashicorpCloudNetwork20200907TGWAttachmentState = "DELETING"

	// HashicorpCloudNetwork20200907TGWAttachmentStateEXPIRED captures enum value "EXPIRED"
	HashicorpCloudNetwork20200907TGWAttachmentStateEXPIRED HashicorpCloudNetwork20200907TGWAttachmentState = "EXPIRED"
)

// for schema
var hashicorpCloudNetwork20200907TGWAttachmentStateEnum []interface{}

func init() {
	var res []HashicorpCloudNetwork20200907TGWAttachmentState
	if err := json.Unmarshal([]byte(`["UNSET","CREATING","PENDING_ACCEPTANCE","ACCEPTED","ACTIVE","FAILED","REJECTED","DELETING","EXPIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudNetwork20200907TGWAttachmentStateEnum = append(hashicorpCloudNetwork20200907TGWAttachmentStateEnum, v)
	}
}

func (m HashicorpCloudNetwork20200907TGWAttachmentState) validateHashicorpCloudNetwork20200907TGWAttachmentStateEnum(path, location string, value HashicorpCloudNetwork20200907TGWAttachmentState) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudNetwork20200907TGWAttachmentStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud network 20200907 t g w attachment state
func (m HashicorpCloudNetwork20200907TGWAttachmentState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudNetwork20200907TGWAttachmentStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this hashicorp cloud network 20200907 t g w attachment state based on the context it is used
func (m HashicorpCloudNetwork20200907TGWAttachmentState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := validate.ReadOnly(ctx, "", "body", HashicorpCloudNetwork20200907TGWAttachmentState(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
