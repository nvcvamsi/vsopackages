// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// ServicePrincipalsServiceListServicePrincipals2Reader is a Reader for the ServicePrincipalsServiceListServicePrincipals2 structure.
type ServicePrincipalsServiceListServicePrincipals2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceListServicePrincipals2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceListServicePrincipals2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceListServicePrincipals2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceListServicePrincipals2OK creates a ServicePrincipalsServiceListServicePrincipals2OK with default headers values
func NewServicePrincipalsServiceListServicePrincipals2OK() *ServicePrincipalsServiceListServicePrincipals2OK {
	return &ServicePrincipalsServiceListServicePrincipals2OK{}
}

/*
ServicePrincipalsServiceListServicePrincipals2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceListServicePrincipals2OK struct {
	Payload *models.HashicorpCloudIamListServicePrincipalsResponse
}

// IsSuccess returns true when this service principals service list service principals2 o k response has a 2xx status code
func (o *ServicePrincipalsServiceListServicePrincipals2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service list service principals2 o k response has a 3xx status code
func (o *ServicePrincipalsServiceListServicePrincipals2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service list service principals2 o k response has a 4xx status code
func (o *ServicePrincipalsServiceListServicePrincipals2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service list service principals2 o k response has a 5xx status code
func (o *ServicePrincipalsServiceListServicePrincipals2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service list service principals2 o k response a status code equal to that given
func (o *ServicePrincipalsServiceListServicePrincipals2OK) IsCode(code int) bool {
	return code == 200
}

func (o *ServicePrincipalsServiceListServicePrincipals2OK) Error() string {
	return fmt.Sprintf("[GET /2019-12-10/iam/{parent_resource_name}/service-principals][%d] servicePrincipalsServiceListServicePrincipals2OK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceListServicePrincipals2OK) String() string {
	return fmt.Sprintf("[GET /2019-12-10/iam/{parent_resource_name}/service-principals][%d] servicePrincipalsServiceListServicePrincipals2OK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceListServicePrincipals2OK) GetPayload() *models.HashicorpCloudIamListServicePrincipalsResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceListServicePrincipals2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamListServicePrincipalsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceListServicePrincipals2Default creates a ServicePrincipalsServiceListServicePrincipals2Default with default headers values
func NewServicePrincipalsServiceListServicePrincipals2Default(code int) *ServicePrincipalsServiceListServicePrincipals2Default {
	return &ServicePrincipalsServiceListServicePrincipals2Default{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceListServicePrincipals2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceListServicePrincipals2Default struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the service principals service list service principals2 default response
func (o *ServicePrincipalsServiceListServicePrincipals2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this service principals service list service principals2 default response has a 2xx status code
func (o *ServicePrincipalsServiceListServicePrincipals2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service list service principals2 default response has a 3xx status code
func (o *ServicePrincipalsServiceListServicePrincipals2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service list service principals2 default response has a 4xx status code
func (o *ServicePrincipalsServiceListServicePrincipals2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service list service principals2 default response has a 5xx status code
func (o *ServicePrincipalsServiceListServicePrincipals2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service list service principals2 default response a status code equal to that given
func (o *ServicePrincipalsServiceListServicePrincipals2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServicePrincipalsServiceListServicePrincipals2Default) Error() string {
	return fmt.Sprintf("[GET /2019-12-10/iam/{parent_resource_name}/service-principals][%d] ServicePrincipalsService_ListServicePrincipals2 default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceListServicePrincipals2Default) String() string {
	return fmt.Sprintf("[GET /2019-12-10/iam/{parent_resource_name}/service-principals][%d] ServicePrincipalsService_ListServicePrincipals2 default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceListServicePrincipals2Default) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceListServicePrincipals2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
