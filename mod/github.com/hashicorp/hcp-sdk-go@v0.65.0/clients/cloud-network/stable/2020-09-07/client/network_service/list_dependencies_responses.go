// Code generated by go-swagger; DO NOT EDIT.

package network_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-network/stable/2020-09-07/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ListDependenciesReader is a Reader for the ListDependencies structure.
type ListDependenciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDependenciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDependenciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListDependenciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDependenciesOK creates a ListDependenciesOK with default headers values
func NewListDependenciesOK() *ListDependenciesOK {
	return &ListDependenciesOK{}
}

/*
ListDependenciesOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListDependenciesOK struct {
	Payload *models.HashicorpCloudNetwork20200907ListDependenciesResponse
}

// IsSuccess returns true when this list dependencies o k response has a 2xx status code
func (o *ListDependenciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list dependencies o k response has a 3xx status code
func (o *ListDependenciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list dependencies o k response has a 4xx status code
func (o *ListDependenciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list dependencies o k response has a 5xx status code
func (o *ListDependenciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list dependencies o k response a status code equal to that given
func (o *ListDependenciesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListDependenciesOK) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{id}/dependencies][%d] listDependenciesOK  %+v", 200, o.Payload)
}

func (o *ListDependenciesOK) String() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{id}/dependencies][%d] listDependenciesOK  %+v", 200, o.Payload)
}

func (o *ListDependenciesOK) GetPayload() *models.HashicorpCloudNetwork20200907ListDependenciesResponse {
	return o.Payload
}

func (o *ListDependenciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudNetwork20200907ListDependenciesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDependenciesDefault creates a ListDependenciesDefault with default headers values
func NewListDependenciesDefault(code int) *ListDependenciesDefault {
	return &ListDependenciesDefault{
		_statusCode: code,
	}
}

/*
ListDependenciesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListDependenciesDefault struct {
	_statusCode int

	Payload *cloud.GrpcGatewayRuntimeError
}

// Code gets the status code for the list dependencies default response
func (o *ListDependenciesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this list dependencies default response has a 2xx status code
func (o *ListDependenciesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list dependencies default response has a 3xx status code
func (o *ListDependenciesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list dependencies default response has a 4xx status code
func (o *ListDependenciesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list dependencies default response has a 5xx status code
func (o *ListDependenciesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list dependencies default response a status code equal to that given
func (o *ListDependenciesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ListDependenciesDefault) Error() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{id}/dependencies][%d] ListDependencies default  %+v", o._statusCode, o.Payload)
}

func (o *ListDependenciesDefault) String() string {
	return fmt.Sprintf("[GET /network/2020-09-07/organizations/{location.organization_id}/projects/{location.project_id}/networks/{id}/dependencies][%d] ListDependencies default  %+v", o._statusCode, o.Payload)
}

func (o *ListDependenciesDefault) GetPayload() *cloud.GrpcGatewayRuntimeError {
	return o.Payload
}

func (o *ListDependenciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GrpcGatewayRuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
