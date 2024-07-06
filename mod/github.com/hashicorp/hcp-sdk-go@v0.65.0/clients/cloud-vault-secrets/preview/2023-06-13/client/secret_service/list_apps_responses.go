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

// ListAppsReader is a Reader for the ListApps structure.
type ListAppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListAppsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListAppsOK creates a ListAppsOK with default headers values
func NewListAppsOK() *ListAppsOK {
	return &ListAppsOK{}
}

/*
ListAppsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListAppsOK struct {
	Payload *models.Secrets20230613ListAppsResponse
}

// IsSuccess returns true when this list apps o k response has a 2xx status code
func (o *ListAppsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list apps o k response has a 3xx status code
func (o *ListAppsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list apps o k response has a 4xx status code
func (o *ListAppsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list apps o k response has a 5xx status code
func (o *ListAppsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list apps o k response a status code equal to that given
func (o *ListAppsOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListAppsOK) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps][%d] listAppsOK  %+v", 200, o.Payload)
}

func (o *ListAppsOK) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps][%d] listAppsOK  %+v", 200, o.Payload)
}

func (o *ListAppsOK) GetPayload() *models.Secrets20230613ListAppsResponse {
	return o.Payload
}

func (o *ListAppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Secrets20230613ListAppsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAppsDefault creates a ListAppsDefault with default headers values
func NewListAppsDefault(code int) *ListAppsDefault {
	return &ListAppsDefault{
		_statusCode: code,
	}
}

/*
ListAppsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListAppsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the list apps default response
func (o *ListAppsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list apps default response has a 2xx status code
func (o *ListAppsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list apps default response has a 3xx status code
func (o *ListAppsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list apps default response has a 4xx status code
func (o *ListAppsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list apps default response has a 5xx status code
func (o *ListAppsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list apps default response a status code equal to that given
func (o *ListAppsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListAppsDefault) Error() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps][%d] ListApps default  %+v", o._statusCode, o.Payload)
}

func (o *ListAppsDefault) String() string {
	return fmt.Sprintf("[GET /secrets/2023-06-13/organizations/{location.organization_id}/projects/{location.project_id}/apps][%d] ListApps default  %+v", o._statusCode, o.Payload)
}

func (o *ListAppsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ListAppsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
