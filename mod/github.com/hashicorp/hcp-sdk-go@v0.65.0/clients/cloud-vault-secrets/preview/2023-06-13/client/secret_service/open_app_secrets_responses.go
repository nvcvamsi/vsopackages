// Code generated by go-swagger; DO NOT EDIT.

package secret_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-06-13/models"
)

// OpenAppSecretsReader is a Reader for the OpenAppSecrets structure.
type OpenAppSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OpenAppSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOpenAppSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOpenAppSecretsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOpenAppSecretsOK creates a OpenAppSecretsOK with default headers values
func NewOpenAppSecretsOK() *OpenAppSecretsOK {
	return &OpenAppSecretsOK{}
}

/*
OpenAppSecretsOK describes a response with status code 200, with default header values.

A successful response.
*/
type OpenAppSecretsOK struct {
	Payload *models.Secrets20230613OpenAppSecretsResponse
}

// IsSuccess returns true when this open app secrets o k response has a 2xx status code
func (o *OpenAppSecretsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this open app secrets o k response has a 3xx status code
func (o *OpenAppSecretsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this open app secrets o k response has a 4xx status code
func (o *OpenAppSecretsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this open app secrets o k response has a 5xx status code
func (o *OpenAppSecretsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this open app secrets o k response a status code equal to that given
func (o *OpenAppSecretsOK) IsCode(code int) bool {
	return code == 200
}

func (o *OpenAppSecretsOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/open][%d] openAppSecretsOK  %+v", 200, o.Payload)
}

func (o *OpenAppSecretsOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/open][%d] openAppSecretsOK  %+v", 200, o.Payload)
}

func (o *OpenAppSecretsOK) GetPayload() *models.Secrets20230613OpenAppSecretsResponse {
	return o.Payload
}

func (o *OpenAppSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613OpenAppSecretsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOpenAppSecretsDefault creates a OpenAppSecretsDefault with default headers values
func NewOpenAppSecretsDefault(code int) *OpenAppSecretsDefault {
	return &OpenAppSecretsDefault{
		_statusCode: code,
	}
}

/*
OpenAppSecretsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type OpenAppSecretsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the open app secrets default response
func (o *OpenAppSecretsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this open app secrets default response has a 2xx status code
func (o *OpenAppSecretsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this open app secrets default response has a 3xx status code
func (o *OpenAppSecretsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this open app secrets default response has a 4xx status code
func (o *OpenAppSecretsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this open app secrets default response has a 5xx status code
func (o *OpenAppSecretsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this open app secrets default response a status code equal to that given
func (o *OpenAppSecretsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *OpenAppSecretsDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/open][%d] OpenAppSecrets default  %+v", o._statusCode, o.Payload)
}

func (o *OpenAppSecretsDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps/{app_name}/open][%d] OpenAppSecrets default  %+v", o._statusCode, o.Payload)
}

func (o *OpenAppSecretsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *OpenAppSecretsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}