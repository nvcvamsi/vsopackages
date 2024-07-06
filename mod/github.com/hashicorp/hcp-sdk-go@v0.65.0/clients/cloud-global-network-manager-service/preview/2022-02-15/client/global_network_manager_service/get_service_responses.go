// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2022-02-15/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// GetServiceReader is a Reader for the GetService structure.
type GetServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetServiceOK creates a GetServiceOK with default headers values
func NewGetServiceOK() *GetServiceOK {
	return &GetServiceOK{}
}

/*
GetServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetServiceOK struct {
	Payload *models.HashicorpCloudGlobalNetworkManager20220215GetServiceResponse
}

// IsSuccess returns true when this get service o k response has a 2xx status code
func (o *GetServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get service o k response has a 3xx status code
func (o *GetServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get service o k response has a 4xx status code
func (o *GetServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get service o k response has a 5xx status code
func (o *GetServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get service o k response a status code equal to that given
func (o *GetServiceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetServiceOK) Error() string {
	return fmt.Sprintf("[GET /2022-02-15/global-network-manager/{cluster_resource_name}/service/{service_name}][%d] getServiceOK  %+v", 200, o.Payload)
}

func (o *GetServiceOK) String() string {
	return fmt.Sprintf("[GET /2022-02-15/global-network-manager/{cluster_resource_name}/service/{service_name}][%d] getServiceOK  %+v", 200, o.Payload)
}

func (o *GetServiceOK) GetPayload() *models.HashicorpCloudGlobalNetworkManager20220215GetServiceResponse {
	return o.Payload
}

func (o *GetServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudGlobalNetworkManager20220215GetServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServiceDefault creates a GetServiceDefault with default headers values
func NewGetServiceDefault(code int) *GetServiceDefault {
	return &GetServiceDefault{
		_statusCode: code,
	}
}

/*
GetServiceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetServiceDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the get service default response
func (o *GetServiceDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get service default response has a 2xx status code
func (o *GetServiceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get service default response has a 3xx status code
func (o *GetServiceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get service default response has a 4xx status code
func (o *GetServiceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get service default response has a 5xx status code
func (o *GetServiceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get service default response a status code equal to that given
func (o *GetServiceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetServiceDefault) Error() string {
	return fmt.Sprintf("[GET /2022-02-15/global-network-manager/{cluster_resource_name}/service/{service_name}][%d] GetService default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceDefault) String() string {
	return fmt.Sprintf("[GET /2022-02-15/global-network-manager/{cluster_resource_name}/service/{service_name}][%d] GetService default  %+v", o._statusCode, o.Payload)
}

func (o *GetServiceDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *GetServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
