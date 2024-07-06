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

// HashicorpCloudPackerRegistryBillingDeprovisionReason hashicorp cloud packer registry billing deprovision reason
//
// swagger:model hashicorp.cloud.packer.RegistryBillingDeprovision.Reason
type HashicorpCloudPackerRegistryBillingDeprovisionReason string

func NewHashicorpCloudPackerRegistryBillingDeprovisionReason(value HashicorpCloudPackerRegistryBillingDeprovisionReason) *HashicorpCloudPackerRegistryBillingDeprovisionReason {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudPackerRegistryBillingDeprovisionReason.
func (m HashicorpCloudPackerRegistryBillingDeprovisionReason) Pointer() *HashicorpCloudPackerRegistryBillingDeprovisionReason {
	return &m
}

const (

	// HashicorpCloudPackerRegistryBillingDeprovisionReasonDELINQUENTBILLINGACCOUNT captures enum value "DELINQUENT_BILLING_ACCOUNT"
	HashicorpCloudPackerRegistryBillingDeprovisionReasonDELINQUENTBILLINGACCOUNT HashicorpCloudPackerRegistryBillingDeprovisionReason = "DELINQUENT_BILLING_ACCOUNT"

	// HashicorpCloudPackerRegistryBillingDeprovisionReasonUSERREQUEST captures enum value "USER_REQUEST"
	HashicorpCloudPackerRegistryBillingDeprovisionReasonUSERREQUEST HashicorpCloudPackerRegistryBillingDeprovisionReason = "USER_REQUEST"

	// HashicorpCloudPackerRegistryBillingDeprovisionReasonHASHIADMINREQUEST captures enum value "HASHI_ADMIN_REQUEST"
	HashicorpCloudPackerRegistryBillingDeprovisionReasonHASHIADMINREQUEST HashicorpCloudPackerRegistryBillingDeprovisionReason = "HASHI_ADMIN_REQUEST"
)

// for schema
var hashicorpCloudPackerRegistryBillingDeprovisionReasonEnum []interface{}

func init() {
	var res []HashicorpCloudPackerRegistryBillingDeprovisionReason
	if err := json.Unmarshal([]byte(`["DELINQUENT_BILLING_ACCOUNT","USER_REQUEST","HASHI_ADMIN_REQUEST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudPackerRegistryBillingDeprovisionReasonEnum = append(hashicorpCloudPackerRegistryBillingDeprovisionReasonEnum, v)
	}
}

func (m HashicorpCloudPackerRegistryBillingDeprovisionReason) validateHashicorpCloudPackerRegistryBillingDeprovisionReasonEnum(path, location string, value HashicorpCloudPackerRegistryBillingDeprovisionReason) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudPackerRegistryBillingDeprovisionReasonEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud packer registry billing deprovision reason
func (m HashicorpCloudPackerRegistryBillingDeprovisionReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudPackerRegistryBillingDeprovisionReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud packer registry billing deprovision reason based on context it is used
func (m HashicorpCloudPackerRegistryBillingDeprovisionReason) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}