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

// HashicorpCloudPackerRegistryConfigTier - PRIVATE_BETA: Deprecated
//
// swagger:model hashicorp.cloud.packer.RegistryConfig.Tier
type HashicorpCloudPackerRegistryConfigTier string

func NewHashicorpCloudPackerRegistryConfigTier(value HashicorpCloudPackerRegistryConfigTier) *HashicorpCloudPackerRegistryConfigTier {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudPackerRegistryConfigTier.
func (m HashicorpCloudPackerRegistryConfigTier) Pointer() *HashicorpCloudPackerRegistryConfigTier {
	return &m
}

const (

	// HashicorpCloudPackerRegistryConfigTierUNSET captures enum value "UNSET"
	HashicorpCloudPackerRegistryConfigTierUNSET HashicorpCloudPackerRegistryConfigTier = "UNSET"

	// HashicorpCloudPackerRegistryConfigTierPRIVATEBETA captures enum value "PRIVATE_BETA"
	HashicorpCloudPackerRegistryConfigTierPRIVATEBETA HashicorpCloudPackerRegistryConfigTier = "PRIVATE_BETA"

	// HashicorpCloudPackerRegistryConfigTierSTANDARD captures enum value "STANDARD"
	HashicorpCloudPackerRegistryConfigTierSTANDARD HashicorpCloudPackerRegistryConfigTier = "STANDARD"

	// HashicorpCloudPackerRegistryConfigTierPLUS captures enum value "PLUS"
	HashicorpCloudPackerRegistryConfigTierPLUS HashicorpCloudPackerRegistryConfigTier = "PLUS"
)

// for schema
var hashicorpCloudPackerRegistryConfigTierEnum []interface{}

func init() {
	var res []HashicorpCloudPackerRegistryConfigTier
	if err := json.Unmarshal([]byte(`["UNSET","PRIVATE_BETA","STANDARD","PLUS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudPackerRegistryConfigTierEnum = append(hashicorpCloudPackerRegistryConfigTierEnum, v)
	}
}

func (m HashicorpCloudPackerRegistryConfigTier) validateHashicorpCloudPackerRegistryConfigTierEnum(path, location string, value HashicorpCloudPackerRegistryConfigTier) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudPackerRegistryConfigTierEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud packer registry config tier
func (m HashicorpCloudPackerRegistryConfigTier) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudPackerRegistryConfigTierEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud packer registry config tier based on context it is used
func (m HashicorpCloudPackerRegistryConfigTier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}