// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_management_service

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

// SSOManagementServiceVerifySSODomainOwnershipReader is a Reader for the SSOManagementServiceVerifySSODomainOwnership structure.
type SSOManagementServiceVerifySSODomainOwnershipReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SSOManagementServiceVerifySSODomainOwnershipReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSSOManagementServiceVerifySSODomainOwnershipOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSSOManagementServiceVerifySSODomainOwnershipDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSSOManagementServiceVerifySSODomainOwnershipOK creates a SSOManagementServiceVerifySSODomainOwnershipOK with default headers values
func NewSSOManagementServiceVerifySSODomainOwnershipOK() *SSOManagementServiceVerifySSODomainOwnershipOK {
	return &SSOManagementServiceVerifySSODomainOwnershipOK{}
}

/*
SSOManagementServiceVerifySSODomainOwnershipOK describes a response with status code 200, with default header values.

A successful response.
*/
type SSOManagementServiceVerifySSODomainOwnershipOK struct {
	Payload *models.HashicorpCloudIamVerifyDomainOwnershipResponse
}

// IsSuccess returns true when this s s o management service verify s s o domain ownership o k response has a 2xx status code
func (o *SSOManagementServiceVerifySSODomainOwnershipOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s s o management service verify s s o domain ownership o k response has a 3xx status code
func (o *SSOManagementServiceVerifySSODomainOwnershipOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s s o management service verify s s o domain ownership o k response has a 4xx status code
func (o *SSOManagementServiceVerifySSODomainOwnershipOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s s o management service verify s s o domain ownership o k response has a 5xx status code
func (o *SSOManagementServiceVerifySSODomainOwnershipOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s s o management service verify s s o domain ownership o k response a status code equal to that given
func (o *SSOManagementServiceVerifySSODomainOwnershipOK) IsCode(code int) bool {
	return code == 200
}

func (o *SSOManagementServiceVerifySSODomainOwnershipOK) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/verify-sso-domain-ownership][%d] sSOManagementServiceVerifySSODomainOwnershipOK  %+v", 200, o.Payload)
}

func (o *SSOManagementServiceVerifySSODomainOwnershipOK) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/verify-sso-domain-ownership][%d] sSOManagementServiceVerifySSODomainOwnershipOK  %+v", 200, o.Payload)
}

func (o *SSOManagementServiceVerifySSODomainOwnershipOK) GetPayload() *models.HashicorpCloudIamVerifyDomainOwnershipResponse {
	return o.Payload
}

func (o *SSOManagementServiceVerifySSODomainOwnershipOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudIamVerifyDomainOwnershipResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSSOManagementServiceVerifySSODomainOwnershipDefault creates a SSOManagementServiceVerifySSODomainOwnershipDefault with default headers values
func NewSSOManagementServiceVerifySSODomainOwnershipDefault(code int) *SSOManagementServiceVerifySSODomainOwnershipDefault {
	return &SSOManagementServiceVerifySSODomainOwnershipDefault{
		_statusCode: code,
	}
}

/*
SSOManagementServiceVerifySSODomainOwnershipDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SSOManagementServiceVerifySSODomainOwnershipDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the s s o management service verify s s o domain ownership default response
func (o *SSOManagementServiceVerifySSODomainOwnershipDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this s s o management service verify s s o domain ownership default response has a 2xx status code
func (o *SSOManagementServiceVerifySSODomainOwnershipDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s s o management service verify s s o domain ownership default response has a 3xx status code
func (o *SSOManagementServiceVerifySSODomainOwnershipDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s s o management service verify s s o domain ownership default response has a 4xx status code
func (o *SSOManagementServiceVerifySSODomainOwnershipDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s s o management service verify s s o domain ownership default response has a 5xx status code
func (o *SSOManagementServiceVerifySSODomainOwnershipDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s s o management service verify s s o domain ownership default response a status code equal to that given
func (o *SSOManagementServiceVerifySSODomainOwnershipDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SSOManagementServiceVerifySSODomainOwnershipDefault) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/verify-sso-domain-ownership][%d] SSOManagementService_VerifySSODomainOwnership default  %+v", o._statusCode, o.Payload)
}

func (o *SSOManagementServiceVerifySSODomainOwnershipDefault) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/organizations/{organization_id}/verify-sso-domain-ownership][%d] SSOManagementService_VerifySSODomainOwnership default  %+v", o._statusCode, o.Payload)
}

func (o *SSOManagementServiceVerifySSODomainOwnershipDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *SSOManagementServiceVerifySSODomainOwnershipDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
