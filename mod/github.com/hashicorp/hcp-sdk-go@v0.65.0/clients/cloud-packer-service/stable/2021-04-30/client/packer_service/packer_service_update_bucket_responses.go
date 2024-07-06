// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceUpdateBucketReader is a Reader for the PackerServiceUpdateBucket structure.
type PackerServiceUpdateBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceUpdateBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceUpdateBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceUpdateBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceUpdateBucketOK creates a PackerServiceUpdateBucketOK with default headers values
func NewPackerServiceUpdateBucketOK() *PackerServiceUpdateBucketOK {
	return &PackerServiceUpdateBucketOK{}
}

/*
PackerServiceUpdateBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceUpdateBucketOK struct {
	Payload *models.HashicorpCloudPackerUpdateBucketResponse
}

// IsSuccess returns true when this packer service update bucket o k response has a 2xx status code
func (o *PackerServiceUpdateBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service update bucket o k response has a 3xx status code
func (o *PackerServiceUpdateBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service update bucket o k response has a 4xx status code
func (o *PackerServiceUpdateBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service update bucket o k response has a 5xx status code
func (o *PackerServiceUpdateBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service update bucket o k response a status code equal to that given
func (o *PackerServiceUpdateBucketOK) IsCode(code int) bool {
	return code == 200
}

func (o *PackerServiceUpdateBucketOK) Error() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}][%d] packerServiceUpdateBucketOK  %+v", 200, o.Payload)
}

func (o *PackerServiceUpdateBucketOK) String() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}][%d] packerServiceUpdateBucketOK  %+v", 200, o.Payload)
}

func (o *PackerServiceUpdateBucketOK) GetPayload() *models.HashicorpCloudPackerUpdateBucketResponse {
	return o.Payload
}

func (o *PackerServiceUpdateBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerUpdateBucketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceUpdateBucketDefault creates a PackerServiceUpdateBucketDefault with default headers values
func NewPackerServiceUpdateBucketDefault(code int) *PackerServiceUpdateBucketDefault {
	return &PackerServiceUpdateBucketDefault{
		_statusCode: code,
	}
}

/*
PackerServiceUpdateBucketDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceUpdateBucketDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the packer service update bucket default response
func (o *PackerServiceUpdateBucketDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this packer service update bucket default response has a 2xx status code
func (o *PackerServiceUpdateBucketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service update bucket default response has a 3xx status code
func (o *PackerServiceUpdateBucketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service update bucket default response has a 4xx status code
func (o *PackerServiceUpdateBucketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service update bucket default response has a 5xx status code
func (o *PackerServiceUpdateBucketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service update bucket default response a status code equal to that given
func (o *PackerServiceUpdateBucketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PackerServiceUpdateBucketDefault) Error() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}][%d] PackerService_UpdateBucket default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceUpdateBucketDefault) String() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}][%d] PackerService_UpdateBucket default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceUpdateBucketDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceUpdateBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PackerServiceUpdateBucketBody packer service update bucket body
swagger:model PackerServiceUpdateBucketBody
*/
type PackerServiceUpdateBucketBody struct {

	// A short description of what this bucket's images are for.
	Description string `json:"description,omitempty"`

	// A key:value map for custom, user-settable metadata about your bucket.
	Labels map[string]string `json:"labels,omitempty"`

	// location
	Location *PackerServiceUpdateBucketParamsBodyLocation `json:"location,omitempty"`

	// A list of the cloud providers or other platforms that are included in the latest complete iteration. e.g aws, gcp, or azure.
	Platforms []string `json:"platforms"`
}

// Validate validates this packer service update bucket body
func (o *PackerServiceUpdateBucketBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceUpdateBucketBody) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(o.Location) { // not required
		return nil
	}

	if o.Location != nil {
		if err := o.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this packer service update bucket body based on the context it is used
func (o *PackerServiceUpdateBucketBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceUpdateBucketBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if o.Location != nil {
		if err := o.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PackerServiceUpdateBucketBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceUpdateBucketBody) UnmarshalBinary(b []byte) error {
	var res PackerServiceUpdateBucketBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PackerServiceUpdateBucketParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model PackerServiceUpdateBucketParamsBodyLocation
*/
type PackerServiceUpdateBucketParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this packer service update bucket params body location
func (o *PackerServiceUpdateBucketParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceUpdateBucketParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(o.Region) { // not required
		return nil
	}

	if o.Region != nil {
		if err := o.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this packer service update bucket params body location based on the context it is used
func (o *PackerServiceUpdateBucketParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceUpdateBucketParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {
		if err := o.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "location" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "location" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PackerServiceUpdateBucketParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceUpdateBucketParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res PackerServiceUpdateBucketParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}