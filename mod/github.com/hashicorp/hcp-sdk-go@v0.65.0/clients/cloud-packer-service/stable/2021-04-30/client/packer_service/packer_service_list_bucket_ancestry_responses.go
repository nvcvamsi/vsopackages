// Code generated by go-swagger; DO NOT EDIT.

package packer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-packer-service/stable/2021-04-30/models"
	cloud "github.com/hashicorp/hcp-sdk-go/clients/cloud-shared/v1/models"
)

// PackerServiceListBucketAncestryReader is a Reader for the PackerServiceListBucketAncestry structure.
type PackerServiceListBucketAncestryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackerServiceListBucketAncestryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackerServiceListBucketAncestryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPackerServiceListBucketAncestryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPackerServiceListBucketAncestryOK creates a PackerServiceListBucketAncestryOK with default headers values
func NewPackerServiceListBucketAncestryOK() *PackerServiceListBucketAncestryOK {
	return &PackerServiceListBucketAncestryOK{}
}

/*
PackerServiceListBucketAncestryOK describes a response with status code 200, with default header values.

A successful response.
*/
type PackerServiceListBucketAncestryOK struct {
	Payload *models.HashicorpCloudPackerListBucketAncestryResponse
}

// IsSuccess returns true when this packer service list bucket ancestry o k response has a 2xx status code
func (o *PackerServiceListBucketAncestryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packer service list bucket ancestry o k response has a 3xx status code
func (o *PackerServiceListBucketAncestryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packer service list bucket ancestry o k response has a 4xx status code
func (o *PackerServiceListBucketAncestryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packer service list bucket ancestry o k response has a 5xx status code
func (o *PackerServiceListBucketAncestryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packer service list bucket ancestry o k response a status code equal to that given
func (o *PackerServiceListBucketAncestryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PackerServiceListBucketAncestryOK) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/ancestry][%d] packerServiceListBucketAncestryOK  %+v", 200, o.Payload)
}

func (o *PackerServiceListBucketAncestryOK) String() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/ancestry][%d] packerServiceListBucketAncestryOK  %+v", 200, o.Payload)
}

func (o *PackerServiceListBucketAncestryOK) GetPayload() *models.HashicorpCloudPackerListBucketAncestryResponse {
	return o.Payload
}

func (o *PackerServiceListBucketAncestryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HashicorpCloudPackerListBucketAncestryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPackerServiceListBucketAncestryDefault creates a PackerServiceListBucketAncestryDefault with default headers values
func NewPackerServiceListBucketAncestryDefault(code int) *PackerServiceListBucketAncestryDefault {
	return &PackerServiceListBucketAncestryDefault{
		_statusCode: code,
	}
}

/*
PackerServiceListBucketAncestryDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PackerServiceListBucketAncestryDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the packer service list bucket ancestry default response
func (o *PackerServiceListBucketAncestryDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this packer service list bucket ancestry default response has a 2xx status code
func (o *PackerServiceListBucketAncestryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this packer service list bucket ancestry default response has a 3xx status code
func (o *PackerServiceListBucketAncestryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this packer service list bucket ancestry default response has a 4xx status code
func (o *PackerServiceListBucketAncestryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this packer service list bucket ancestry default response has a 5xx status code
func (o *PackerServiceListBucketAncestryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this packer service list bucket ancestry default response a status code equal to that given
func (o *PackerServiceListBucketAncestryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PackerServiceListBucketAncestryDefault) Error() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/ancestry][%d] PackerService_ListBucketAncestry default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceListBucketAncestryDefault) String() string {
	return fmt.Sprintf("[GET /packer/2021-04-30/organizations/{location.organization_id}/projects/{location.project_id}/images/{bucket_slug}/ancestry][%d] PackerService_ListBucketAncestry default  %+v", o._statusCode, o.Payload)
}

func (o *PackerServiceListBucketAncestryDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *PackerServiceListBucketAncestryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
