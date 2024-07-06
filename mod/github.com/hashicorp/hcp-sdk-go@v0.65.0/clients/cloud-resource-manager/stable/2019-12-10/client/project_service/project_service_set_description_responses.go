// Code generated by go-swagger; DO NOT EDIT.

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ProjectServiceSetDescriptionReader is a Reader for the ProjectServiceSetDescription structure.
type ProjectServiceSetDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectServiceSetDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectServiceSetDescriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectServiceSetDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectServiceSetDescriptionOK creates a ProjectServiceSetDescriptionOK with default headers values
func NewProjectServiceSetDescriptionOK() *ProjectServiceSetDescriptionOK {
	return &ProjectServiceSetDescriptionOK{}
}

/*
ProjectServiceSetDescriptionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProjectServiceSetDescriptionOK struct {
	Payload interface{}
}

// IsSuccess returns true when this project service set description o k response has a 2xx status code
func (o *ProjectServiceSetDescriptionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this project service set description o k response has a 3xx status code
func (o *ProjectServiceSetDescriptionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this project service set description o k response has a 4xx status code
func (o *ProjectServiceSetDescriptionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this project service set description o k response has a 5xx status code
func (o *ProjectServiceSetDescriptionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this project service set description o k response a status code equal to that given
func (o *ProjectServiceSetDescriptionOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectServiceSetDescriptionOK) Error() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/description][%d] projectServiceSetDescriptionOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceSetDescriptionOK) String() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/description][%d] projectServiceSetDescriptionOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceSetDescriptionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ProjectServiceSetDescriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectServiceSetDescriptionDefault creates a ProjectServiceSetDescriptionDefault with default headers values
func NewProjectServiceSetDescriptionDefault(code int) *ProjectServiceSetDescriptionDefault {
	return &ProjectServiceSetDescriptionDefault{
		_statusCode: code,
	}
}

/*
ProjectServiceSetDescriptionDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProjectServiceSetDescriptionDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the project service set description default response
func (o *ProjectServiceSetDescriptionDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this project service set description default response has a 2xx status code
func (o *ProjectServiceSetDescriptionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this project service set description default response has a 3xx status code
func (o *ProjectServiceSetDescriptionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this project service set description default response has a 4xx status code
func (o *ProjectServiceSetDescriptionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this project service set description default response has a 5xx status code
func (o *ProjectServiceSetDescriptionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this project service set description default response a status code equal to that given
func (o *ProjectServiceSetDescriptionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProjectServiceSetDescriptionDefault) Error() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/description][%d] ProjectService_SetDescription default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceSetDescriptionDefault) String() string {
	return fmt.Sprintf("[PUT /resource-manager/2019-12-10/projects/{id}/description][%d] ProjectService_SetDescription default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceSetDescriptionDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ProjectServiceSetDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ProjectServiceSetDescriptionBody ProjectSetDescriptionRequest see ProjectService.SetDescription
swagger:model ProjectServiceSetDescriptionBody
*/
type ProjectServiceSetDescriptionBody struct {

	// description is the value the project's description should be updated to.
	Description string `json:"description,omitempty"`
}

// Validate validates this project service set description body
func (o *ProjectServiceSetDescriptionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this project service set description body based on context it is used
func (o *ProjectServiceSetDescriptionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ProjectServiceSetDescriptionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ProjectServiceSetDescriptionBody) UnmarshalBinary(b []byte) error {
	var res ProjectServiceSetDescriptionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
