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

// HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus  - VERSION_UP_TO_DATE: VERSION_UP_TO_DATE is used when node is running the latest Vault Version.
//   - UPGRADE_AVAILABLE: UPGRADE_AVAILABLE is used when node is running the latest minor release of a Vault Version, but there is a new major Vault version available for upgrade.
//   - UPGRADE_RECOMMENDED: UPGRADE_RECOMMENDED is used when node is running an outdated but still supported version, but there is a new minor or major Vault versions available for upgrade.
//   - UPGRADE_REQUIRED: UPGRADE_REQUIRED is used when node is running a no longer supported version and there are minor and major versions available for upgrade.
//
// swagger:model hashicorp.cloud.vault_link_20221107.LinkedClusterNode.VersionStatus
type HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus string

func NewHashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus(value HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus) *HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus.
func (m HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus) Pointer() *HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus {
	return &m
}

const (

	// HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusLINKEDCLUSTERNODEVERSIONSTATUSINVALID captures enum value "LINKED_CLUSTER_NODE_VERSION_STATUS_INVALID"
	HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusLINKEDCLUSTERNODEVERSIONSTATUSINVALID HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus = "LINKED_CLUSTER_NODE_VERSION_STATUS_INVALID"

	// HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusVERSIONUPTODATE captures enum value "VERSION_UP_TO_DATE"
	HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusVERSIONUPTODATE HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus = "VERSION_UP_TO_DATE"

	// HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusUPGRADEAVAILABLE captures enum value "UPGRADE_AVAILABLE"
	HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusUPGRADEAVAILABLE HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus = "UPGRADE_AVAILABLE"

	// HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusUPGRADERECOMMENDED captures enum value "UPGRADE_RECOMMENDED"
	HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusUPGRADERECOMMENDED HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus = "UPGRADE_RECOMMENDED"

	// HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusUPGRADEREQUIRED captures enum value "UPGRADE_REQUIRED"
	HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusUPGRADEREQUIRED HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus = "UPGRADE_REQUIRED"
)

// for schema
var hashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusEnum []interface{}

func init() {
	var res []HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus
	if err := json.Unmarshal([]byte(`["LINKED_CLUSTER_NODE_VERSION_STATUS_INVALID","VERSION_UP_TO_DATE","UPGRADE_AVAILABLE","UPGRADE_RECOMMENDED","UPGRADE_REQUIRED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusEnum = append(hashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusEnum, v)
	}
}

func (m HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus) validateHashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusEnum(path, location string, value HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus) error {
	if err := validate.EnumCase(path, location, value, hashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this hashicorp cloud vault link 20221107 linked cluster node version status
func (m HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this hashicorp cloud vault link 20221107 linked cluster node version status based on context it is used
func (m HashicorpCloudVaultLink20221107LinkedClusterNodeVersionStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}