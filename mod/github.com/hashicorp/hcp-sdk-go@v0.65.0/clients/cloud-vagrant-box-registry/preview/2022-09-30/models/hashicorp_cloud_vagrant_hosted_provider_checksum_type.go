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

// HashicorpCloudVagrantHostedProviderChecksumType ChecksumTypes describe supported checksum types.
//
// swagger:model hashicorp.cloud.vagrant.HostedProvider.ChecksumType
type HashicorpCloudVagrantHostedProviderChecksumType string

func NewHashicorpCloudVagrantHostedProviderChecksumType(value HashicorpCloudVagrantHostedProviderChecksumType) *HashicorpCloudVagrantHostedProviderChecksumType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudVagrantHostedProviderChecksumType.
func (m HashicorpCloudVagrantHostedProviderChecksumType) Pointer() *HashicorpCloudVagrantHostedProviderChecksumType {
	return &m
}

const (

	// HashicorpCloudVagrantHostedProviderChecksumTypeNONE captures enum value "NONE"
	HashicorpCloudVagrantHostedProviderChecksumTypeNONE HashicorpCloudVagrantHostedProviderChecksumType = "NONE"

	// HashicorpCloudVagrantHostedProviderChecksumTypeMD5 captures enum value "MD5"
	HashicorpCloudVagrantHostedProviderChecksumTypeMD5 HashicorpCloudVagrantHostedProviderChecksumType = "MD5"

	// HashicorpCloudVagrantHostedProviderChecksumTypeSHA1 captures enum value "SHA1"
	HashicorpCloudVagrantHostedProviderChecksumTypeSHA1 HashicorpCloudVagrantHostedProviderChecksumType = "SHA1"

	// HashicorpCloudVagrantHostedProviderChecksumTypeSHA256 captures enum value "SHA256"
	HashicorpCloudVagrantHostedProviderChecksumTypeSHA256 HashicorpCloudVagrantHostedProviderChecksumType = "SHA256"

	// HashicorpCloudVagrantHostedProviderChecksumTypeSHA384 captures enum value "SHA384"
	HashicorpCloudVagrantHostedProviderChecksumTypeSHA384 HashicorpCloudVagrantHostedProviderChecksumType = "SHA384"

	// HashicorpCloudVagrantHostedProviderChecksumTypeSHA512 captures enum value "SHA512"
	HashicorpCloudVagrantHostedProviderChecksumTypeSHA512 HashicorpCloudVagrantHostedProviderChecksumType = "SHA512"
)

// for schema
var hashicorpCloudVagrantHostedProviderChecksumTypeEnum []interface{}

func init() {
	var res []HashicorpCloudVagrantHostedProviderChecksumType
	if err := json.Unmarshal([]byte(`["NONE","MD5","SHA1","SHA256","SHA384","SHA512"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVagrantHostedProviderChecksumTypeEnum = append(hashicorpCloudVagrantHostedProviderChecksumTypeEnum, v)
	}
}

func (m HashicorpCloudVagrantHostedProviderChecksumType) validateHashicorpCloudVagrantHostedProviderChecksumTypeEnum(path, location string, value HashicorpCloudVagrantHostedProviderChecksumType) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVagrantHostedProviderChecksumTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vagrant hosted provider checksum type
func (m HashicorpCloudVagrantHostedProviderChecksumType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVagrantHostedProviderChecksumTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud vagrant hosted provider checksum type based on context it is used
func (m HashicorpCloudVagrantHostedProviderChecksumType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
