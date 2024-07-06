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

// PackerServiceUpdateBuildReader is a Reader for the PackerServiceUpdateBuild structure.
type PackerServiceUpdateBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceUpdateBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceUpdateBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceUpdateBuildDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceUpdateBuildOK creates a PackerServiceUpdateBuildOK with default headers values
func NewPackerServiceUpdateBuildOK() *PackerServiceUpdateBuildOK {
	return &PackerServiceUpdateBuildOK{}
}

/*
PackerServiceUpdateBuildOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceUpdateBuildOK struct {
	Payload *models.HashicorpCloudPackerUpdateBuildResponse
}

// IsSuccess returns true when this packer service update build o k response has a 2xx status code
func (o *PackerServiceUpdateBuildOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service update build o k response has a 3xx status code
func (o *PackerServiceUpdateBuildOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service update build o k response has a 4xx status code
func (o *PackerServiceUpdateBuildOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service update build o k response has a 5xx status code
func (o *PackerServiceUpdateBuildOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service update build o k response a status code equal to that given
func (o *PackerServiceUpdateBuildOK) IsCode(code int) bool {
	return code == 200
}

func (o *PackerServiceUpdateBuildOK) Error() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] packerServiceUpdateBuildOK  %+v", 200, o.Payload)
}

func (o *PackerServiceUpdateBuildOK) String() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] packerServiceUpdateBuildOK  %+v", 200, o.Payload)
}

func (o *PackerServiceUpdateBuildOK) GetPayload() *models.HashicorpCloudPackerUpdateBuildResponse {
	return o.Payload
}

func (o *PackerServiceUpdateBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerUpdateBuildResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceUpdateBuildDefault creates a PackerServiceUpdateBuildDefault with default headers values
func NewPackerServiceUpdateBuildDefault(code int) *PackerServiceUpdateBuildDefault {
	return &PackerServiceUpdateBuildDefault{
		_statusCode: code,
	}
}

/*
PackerServiceUpdateBuildDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceUpdateBuildDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the packer service update build default response
func (o *PackerServiceUpdateBuildDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this packer service update build default response has a 2xx status code
func (o *PackerServiceUpdateBuildDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service update build default response has a 3xx status code
func (o *PackerServiceUpdateBuildDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service update build default response has a 4xx status code
func (o *PackerServiceUpdateBuildDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service update build default response has a 5xx status code
func (o *PackerServiceUpdateBuildDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service update build default response a status code equal to that given
func (o *PackerServiceUpdateBuildDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PackerServiceUpdateBuildDefault) Error() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] PackerService_UpdateBuild default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceUpdateBuildDefault) String() string {
	return fmt.Sprintf("[PATCH /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/builds/{build_id}][%d] PackerService_UpdateBuild default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceUpdateBuildDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceUpdateBuildDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PackerServiceUpdateBuildBody packer service update build body
swagger:model PackerServiceUpdateBuildBody
*/
type PackerServiceUpdateBuildBody struct {

	// location
	Location *PackerServiceUpdateBuildParamsBodyLocation `json:"location,omitempty"`

	// Information about the build you are updating.
	Updates *models.HashicorpCloudPackerBuildUpdates `json:"updates,omitempty"`
}

// Validate validates this packer service update build body
func (o *PackerServiceUpdateBuildBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceUpdateBuildBody) validateLocation(formats strfmt.Registry) error {
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

func (o *PackerServiceUpdateBuildBody) validateUpdates(formats strfmt.Registry) error {
	if swag.IsZero(o.Updates) { // not required
		return nil
	}

	if o.Updates != nil {
		if err := o.Updates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "updates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "updates")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this packer service update build body based on the context it is used
func (o *PackerServiceUpdateBuildBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUpdates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceUpdateBuildBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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

func (o *PackerServiceUpdateBuildBody) contextValidateUpdates(ctx context.Context, formats strfmt.Registry) error {

	if o.Updates != nil {
		if err := o.Updates.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "updates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "updates")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PackerServiceUpdateBuildBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceUpdateBuildBody) UnmarshalBinary(b []byte) error {
	var res PackerServiceUpdateBuildBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PackerServiceUpdateBuildParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model PackerServiceUpdateBuildParamsBodyLocation
*/
type PackerServiceUpdateBuildParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *cloud.HashicorpCloudLocationRegion `json:"region,omitempty"`
}

// Validate validates this packer service update build params body location
func (o *PackerServiceUpdateBuildParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceUpdateBuildParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this packer service update build params body location based on the context it is used
func (o *PackerServiceUpdateBuildParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PackerServiceUpdateBuildParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (o *PackerServiceUpdateBuildParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PackerServiceUpdateBuildParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res PackerServiceUpdateBuildParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
