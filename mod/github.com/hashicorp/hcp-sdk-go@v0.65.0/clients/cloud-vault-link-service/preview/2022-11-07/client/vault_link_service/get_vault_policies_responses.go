// Code generated by go-swagger; DO NOT EDIT.

package vault_link_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-link-service/preview/2022-11-07/models"
)

// GetVaultPoliciesReader is a Reader for the GetVaultPolicies structure.
type GetVaultPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVaultPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVaultPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVaultPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVaultPoliciesOK creates a GetVaultPoliciesOK with default headers values
func NewGetVaultPoliciesOK() *GetVaultPoliciesOK {
	return &GetVaultPoliciesOK{}
}

/*
GetVaultPoliciesOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetVaultPoliciesOK struct {
	Payload *models.HashicorpCloudVaultLink20221107VaultPolicyResponse
}

// IsSuccess returns true when this get vault policies o k response has a 2xx status code
func (o *GetVaultPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vault policies o k response has a 3xx status code
func (o *GetVaultPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vault policies o k response has a 4xx status code
func (o *GetVaultPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vault policies o k response has a 5xx status code
func (o *GetVaultPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vault policies o k response a status code equal to that given
func (o *GetVaultPoliciesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVaultPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /vault-link/2022-11-07/organizations/{location.organization_id}/projects/{location.project_id}/{cluster_id}/policies][%d] getVaultPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetVaultPoliciesOK) String() string {
	return fmt.Sprintf("[GET /vault-link/2022-11-07/organizations/{location.organization_id}/projects/{location.project_id}/{cluster_id}/policies][%d] getVaultPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetVaultPoliciesOK) GetPayload() *models.HashicorpCloudVaultLink20221107VaultPolicyResponse {
	return o.Payload
}

func (o *GetVaultPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudVaultLink20221107VaultPolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVaultPoliciesDefault creates a GetVaultPoliciesDefault with default headers values
func NewGetVaultPoliciesDefault(code int) *GetVaultPoliciesDefault {
	return &GetVaultPoliciesDefault{
		_statusCode: code,
	}
}

/*
GetVaultPoliciesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetVaultPoliciesDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the get vault policies default response
func (o *GetVaultPoliciesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get vault policies default response has a 2xx status code
func (o *GetVaultPoliciesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get vault policies default response has a 3xx status code
func (o *GetVaultPoliciesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get vault policies default response has a 4xx status code
func (o *GetVaultPoliciesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get vault policies default response has a 5xx status code
func (o *GetVaultPoliciesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get vault policies default response a status code equal to that given
func (o *GetVaultPoliciesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetVaultPoliciesDefault) Error() string {
	return fmt.Sprintf("[GET /vault-link/2022-11-07/organizations/{location.organization_id}/projects/{location.project_id}/{cluster_id}/policies][%d] GetVaultPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetVaultPoliciesDefault) String() string {
	return fmt.Sprintf("[GET /vault-link/2022-11-07/organizations/{location.organization_id}/projects/{location.project_id}/{cluster_id}/policies][%d] GetVaultPolicies default  %+v", o._statusCode, o.Payload)
}

func (o *GetVaultPoliciesDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *GetVaultPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}