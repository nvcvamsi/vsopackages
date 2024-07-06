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

// ServicePrincipalsServiceGetProjectServicePrincipalReader is a Reader for the ServicePrincipalsServiceGetProjectServicePrincipal structure.
type ServicePrincipalsServiceGetProjectServicePrincipalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServicePrincipalsServiceGetProjectServicePrincipalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServicePrincipalsServiceGetProjectServicePrincipalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServicePrincipalsServiceGetProjectServicePrincipalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServicePrincipalsServiceGetProjectServicePrincipalOK creates a ServicePrincipalsServiceGetProjectServicePrincipalOK with default headers values
func NewServicePrincipalsServiceGetProjectServicePrincipalOK() *ServicePrincipalsServiceGetProjectServicePrincipalOK {
	return &ServicePrincipalsServiceGetProjectServicePrincipalOK{}
}

/*
ServicePrincipalsServiceGetProjectServicePrincipalOK describes a response with status code 200, with default header values.

A successful response.
*/
type ServicePrincipalsServiceGetProjectServicePrincipalOK struct {
	Payload *models.HashicorpCloudIamGetProjectServicePrincipalResponse
}

// IsSuccess returns true when this service principals service get project service principal o k response has a 2xx status code
func (o *ServicePrincipalsServiceGetProjectServicePrincipalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service principals service get project service principal o k response has a 3xx status code
func (o *ServicePrincipalsServiceGetProjectServicePrincipalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service principals service get project service principal o k response has a 4xx status code
func (o *ServicePrincipalsServiceGetProjectServicePrincipalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service principals service get project service principal o k response has a 5xx status code
func (o *ServicePrincipalsServiceGetProjectServicePrincipalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service principals service get project service principal o k response a status code equal to that given
func (o *ServicePrincipalsServiceGetProjectServicePrincipalOK) IsCode(code int) bool {
	return code == 200
}

func (o *ServicePrincipalsServiceGetProjectServicePrincipalOK) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] servicePrincipalsServiceGetProjectServicePrincipalOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceGetProjectServicePrincipalOK) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] servicePrincipalsServiceGetProjectServicePrincipalOK  %+v", 200, o.Payload)
}

func (o *ServicePrincipalsServiceGetProjectServicePrincipalOK) GetPayload() *models.HashicorpCloudIamGetProjectServicePrincipalResponse {
	return o.Payload
}

func (o *ServicePrincipalsServiceGetProjectServicePrincipalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamGetProjectServicePrincipalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServicePrincipalsServiceGetProjectServicePrincipalDefault creates a ServicePrincipalsServiceGetProjectServicePrincipalDefault with default headers values
func NewServicePrincipalsServiceGetProjectServicePrincipalDefault(code int) *ServicePrincipalsServiceGetProjectServicePrincipalDefault {
	return &ServicePrincipalsServiceGetProjectServicePrincipalDefault{
		_statusCode: code,
	}
}

/*
ServicePrincipalsServiceGetProjectServicePrincipalDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ServicePrincipalsServiceGetProjectServicePrincipalDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the service principals service get project service principal default response
func (o *ServicePrincipalsServiceGetProjectServicePrincipalDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this service principals service get project service principal default response has a 2xx status code
func (o *ServicePrincipalsServiceGetProjectServicePrincipalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this service principals service get project service principal default response has a 3xx status code
func (o *ServicePrincipalsServiceGetProjectServicePrincipalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this service principals service get project service principal default response has a 4xx status code
func (o *ServicePrincipalsServiceGetProjectServicePrincipalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this service principals service get project service principal default response has a 5xx status code
func (o *ServicePrincipalsServiceGetProjectServicePrincipalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this service principals service get project service principal default response a status code equal to that given
func (o *ServicePrincipalsServiceGetProjectServicePrincipalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ServicePrincipalsServiceGetProjectServicePrincipalDefault) Error() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] ServicePrincipalsService_GetProjectServicePrincipal default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceGetProjectServicePrincipalDefault) String() string {
	return fmt.Sprintf("[GET /iam/2019-12-10/organizations/{organization_id}/projects/{project_id}/service-principals/{principal_id}][%d] ServicePrincipalsService_GetProjectServicePrincipal default  %+v", o._statusCode, o.Payload)
}

func (o *ServicePrincipalsServiceGetProjectServicePrincipalDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ServicePrincipalsServiceGetProjectServicePrincipalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
