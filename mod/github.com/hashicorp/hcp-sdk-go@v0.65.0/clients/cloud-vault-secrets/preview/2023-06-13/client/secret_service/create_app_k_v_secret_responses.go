// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-06-13/models"
)

// CreateAppKVSecretReader is a Reader for the CreateAppKVSecret structure.
type CreateAppKVSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAppKVSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateAppKVSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateAppKVSecretDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateAppKVSecretOK creates a CreateAppKVSecretOK with default headers values
func NewCreateAppKVSecretOK() *CreateAppKVSecretOK {
	return &CreateAppKVSecretOK{}
}

/*
CreateAppKVSecretOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateAppKVSecretOK struct {
	Payload *models.Secrets20230613CreateAppKVSecretResponse
}

// IsSuccess returns true when this create app k v secret o k response has a 2xx status code
func (o *CreateAppKVSecretOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create app k v secret o k response has a 3xx status code
func (o *CreateAppKVSecretOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create app k v secret o k response has a 4xx status code
func (o *CreateAppKVSecretOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create app k v secret o k response has a 5xx status code
func (o *CreateAppKVSecretOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create app k v secret o k response a status code equal to that given
func (o *CreateAppKVSecretOK) IsCode(code int) bool {
	return code == 200
}

func (o *CreateAppKVSecretOK) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/kv][%d] createAppKVSecretOK  %+v", 200, o.Payload)
}

func (o *CreateAppKVSecretOK) String() string {
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/kv][%d] createAppKVSecretOK  %+v", 200, o.Payload)
}

func (o *CreateAppKVSecretOK) GetPayload() *models.Secrets20230613CreateAppKVSecretResponse {
	return o.Payload
}

func (o *CreateAppKVSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613CreateAppKVSecretResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAppKVSecretDefault creates a CreateAppKVSecretDefault with default headers values
func NewCreateAppKVSecretDefault(code int) *CreateAppKVSecretDefault {
	return &CreateAppKVSecretDefault{
		_statusCode: code,
	}
}

/*
CreateAppKVSecretDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateAppKVSecretDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the create app k v secret default response
func (o *CreateAppKVSecretDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this create app k v secret default response has a 2xx status code
func (o *CreateAppKVSecretDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create app k v secret default response has a 3xx status code
func (o *CreateAppKVSecretDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create app k v secret default response has a 4xx status code
func (o *CreateAppKVSecretDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create app k v secret default response has a 5xx status code
func (o *CreateAppKVSecretDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create app k v secret default response a status code equal to that given
func (o *CreateAppKVSecretDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CreateAppKVSecretDefault) Error() string {
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/kv][%d] CreateAppKVSecret default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAppKVSecretDefault) String() string {
	return fmt.Sprintf("[POST /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/kv][%d] CreateAppKVSecret default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAppKVSecretDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *CreateAppKVSecretDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateAppKVSecretBody create app k v secret body
swagger:model CreateAppKVSecretBody
*/
type CreateAppKVSecretBody struct {

	// location
	Location *CreateAppKVSecretParamsBodyLocation `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this create app k v secret body
func (o *CreateAppKVSecretBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAppKVSecretBody) validateLocation(formats strfmt.Registry) error {
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

// ContextValidate validate this create app k v secret body based on the context it is used
func (o *CreateAppKVSecretBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAppKVSecretBody) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateAppKVSecretBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAppKVSecretBody) UnmarshalBinary(b []byte) error {
	var res CreateAppKVSecretBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateAppKVSecretParamsBodyLocation Location represents a target for an operation in HCP.
swagger:model CreateAppKVSecretParamsBodyLocation
*/
type CreateAppKVSecretParamsBodyLocation struct {

	// region is the region that the resource is located in. It is
	// optional if the object being referenced is a global object.
	Region *models.LocationRegion `json:"region,omitempty"`
}

// Validate validates this create app k v secret params body location
func (o *CreateAppKVSecretParamsBodyLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAppKVSecretParamsBodyLocation) validateRegion(formats strfmt.Registry) error {
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

// ContextValidate validate this create app k v secret params body location based on the context it is used
func (o *CreateAppKVSecretParamsBodyLocation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateAppKVSecretParamsBodyLocation) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

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
func (o *CreateAppKVSecretParamsBodyLocation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateAppKVSecretParamsBodyLocation) UnmarshalBinary(b []byte) error {
	var res CreateAppKVSecretParamsBodyLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
