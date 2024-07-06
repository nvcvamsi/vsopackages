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

// PackerServiceCreateBucketReader is a Reader for the PackerServiceCreateBucket structure.
type PackerServiceCreateBucketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceCreateBucketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceCreateBucketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceCreateBucketDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceCreateBucketOK creates a PackerServiceCreateBucketOK with default headers values
func NewPackerServiceCreateBucketOK() *PackerServiceCreateBucketOK {
	return &PackerServiceCreateBucketOK{}
}

/*
PackerServiceCreateBucketOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceCreateBucketOK struct {
	Payload *models.HashicorpCloudPackerCreateBucketResponse
}

// IsSuccess returns true when this packer service create bucket o k response has a 2xx status code
func (o *PackerServiceCreateBucketOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service create bucket o k response has a 3xx status code
func (o *PackerServiceCreateBucketOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service create bucket o k response has a 4xx status code
func (o *PackerServiceCreateBucketOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service create bucket o k response has a 5xx status code
func (o *PackerServiceCreateBucketOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service create bucket o k response a status code equal to that given
func (o *PackerServiceCreateBucketOK) IsCode(code int) bool {
	return code == 200
}

func (o *PackerServiceCreateBucketOK) Error() string {
	return fmt.Sprintf("[PUT /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images][%d] packerServiceCreateBucketOK  %+v", 200, o.Payload)
}

func (o *PackerServiceCreateBucketOK) String() string {
	return fmt.Sprintf("[PUT /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images][%d] packerServiceCreateBucketOK  %+v", 200, o.Payload)
}

func (o *PackerServiceCreateBucketOK) GetPayload() *models.HashicorpCloudPackerCreateBucketResponse {
	return o.Payload
}

func (o *PackerServiceCreateBucketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerCreateBucketResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceCreateBucketDefault creates a PackerServiceCreateBucketDefault with default headers values
func NewPackerServiceCreateBucketDefault(code int) *PackerServiceCreateBucketDefault {
	return &PackerServiceCreateBucketDefault{
		_statusCode: code,
	}
}

/*
PackerServiceCreateBucketDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceCreateBucketDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the packer service create bucket default response
func (o *PackerServiceCreateBucketDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this packer service create bucket default response has a 2xx status code
func (o *PackerServiceCreateBucketDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service create bucket default response has a 3xx status code
func (o *PackerServiceCreateBucketDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service create bucket default response has a 4xx status code
func (o *PackerServiceCreateBucketDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service create bucket default response has a 5xx status code
func (o *PackerServiceCreateBucketDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service create bucket default response a status code equal to that given
func (o *PackerServiceCreateBucketDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PackerServiceCreateBucketDefault) Error() string {
	return fmt.Sprintf("[PUT /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images][%d] PackerService_CreateBucket default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceCreateBucketDefault) String() string {
	return fmt.Sprintf("[PUT /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images][%d] PackerService_CreateBucket default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceCreateBucketDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceCreateBucketDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PackerServiceCreateBucketBody packer service create bucket body
swagger:model PackerServiceCreateBucketBody
*/
type PackerServiceCreateBucketBody struct {

	// Human-readable name for the bucket.
	BucketSlug string `json:"bucket_slug,omitempty"`

	// A short description of what this bucket's images are for.
	Description string `json:"description,omitempty"`

	// A key:value map for custom, user-settable metadata about your bucket.
	Labels map[string]string `json:"labels,omitempty"`

	// location
	Location *PackerServiceCreateBucketParamsBodyLocation `json:"location,omitempty"`
}

// Validate validates this packer service create bucket body
func (o *PackerServiceCreateBucketBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceCreateBucketBody) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this packer service create bucket body based on the context it is used
func (o *PackerServiceCreateBucketBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceCreateBucketBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (o *PackerServiceCreateBucketBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceCreateBucketBody) UnmarshalBinary(b []byte) error {
	var res PackerServiceCreateBucketBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PackerServiceCreateBucketParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model PackerServiceCreateBucketParamsBodyLocation
*/
type PackerServiceCreateBucketParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this packer service create bucket params body location
func (o *PackerServiceCreateBucketParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceCreateBucketParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this packer service create bucket params body location based on the context it is used
func (o *PackerServiceCreateBucketParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceCreateBucketParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (o *PackerServiceCreateBucketParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceCreateBucketParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res PackerServiceCreateBucketParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}