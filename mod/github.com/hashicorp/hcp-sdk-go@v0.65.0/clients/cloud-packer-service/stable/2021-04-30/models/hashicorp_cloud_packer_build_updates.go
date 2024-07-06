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

// HashicorpCloudPackerBuildUpdates BuildUpdates is used to group the elements of a Build that are
// allowed to be updated after the build has been created. It is part of the
// UpdateBuildRequest.
//
// swagger:model hashicorp.cloud.packer.BuildUpdates
type HashicorpCloudPackerBuildUpdates struct {

	// The cloud provider that this build produced artifacts for.
	// For example, AWS, GCP, or Azure.
	CloudProvider string `json:"cloud_provider,omitempty"`

	// A list of images to create and associate with this build.
	Images []*HashicorpCloudPackerImageCreateBody `json:"images"`

	// A key:value map for custom, user-settable metadata about your build.
	Labels map[string]string `json:"labels,omitempty"`

	// The UUID specific to this call to Packer build. If you use the manifest
	// post-processor, this UUID will match the UUID present there.
	PackerRunUUID string `json:"packer_run_uuid,omitempty"`

	// The ID of the channel that it was used to fetch the source_iteration_id.
	// When the source channel ID is set, the source iteration ID should also be set.
	SourceChannelID string `json:"source_channel_id,omitempty"`

	// The ID or URL of the remote cloud source image. Used for tracking image
	// dependencies for build pipelines.
	SourceImageID string `json:"source_image_id,omitempty"`

	// The ID of the parent iteration associated with the `source_image_id`.
	// When the source iteration ID is set, the source image ID should also be set.
	SourceIterationID string `json:"source_iteration_id,omitempty"`

	// Status of the build. The status can be RUNNING, DONE, CANCELLED, FAILED,
	// or UNSET.
	Status *HashicorpCloudPackerBuildStatus `json:"status,omitempty"`
}

// Validate validates this hashicorp cloud packer build updates
func (m *HashicorpCloudPackerBuildUpdates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerBuildUpdates) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	for i := 0; i < len(m.Images); i++ {
		if swag.IsZero(m.Images[i]) { // not required
			continue
		}

		if m.Images[i] != nil {
			if err := m.Images[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudPackerBuildUpdates) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp cloud packer build updates based on the context it is used
func (m *HashicorpCloudPackerBuildUpdates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpCloudPackerBuildUpdates) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Images); i++ {

		if m.Images[i] != nil {
			if err := m.Images[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("images" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("images" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HashicorpCloudPackerBuildUpdates) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpCloudPackerBuildUpdates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpCloudPackerBuildUpdates) UnmarshalBinary(b []byte) error {
	var res HashicorpCloudPackerBuildUpdates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
