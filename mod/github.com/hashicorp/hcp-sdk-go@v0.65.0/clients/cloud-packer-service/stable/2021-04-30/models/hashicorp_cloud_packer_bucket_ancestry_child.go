// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HashicorpCloudPackerBucketAncestryChild hashicorp cloud packer bucket ancestry child
//
// swagger:model hashicorp.cloud.packer.BucketAncestry.Child
type HashicorpCloudPackerBucketAncestryChild struct {

	// The child image bucket's slug.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// The child iteration's build fingerprint.
	IterationFingerprint string `json:"iteration_fingerprint,omitempty"`

	// The child iteration's ULID.
	IterationID string `json:"iteration_id,omitempty"`

	// The child iteration's incremental version.
	IterationIncrementalVersion int32 `json:"iteration_incremental_version,omitempty"`
}

// Validate validates this hashicorp cloud packer bucket ancestry child
func (m *HashicorpCloudPackerBucketAncestryChild) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this hashicorp cloud packer bucket ancestry child based on context it is used
func (m *HashicorpCloudPackerBucketAncestryChild) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerBucketAncestryChild) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerBucketAncestryChild) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerBucketAncestryChild
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
