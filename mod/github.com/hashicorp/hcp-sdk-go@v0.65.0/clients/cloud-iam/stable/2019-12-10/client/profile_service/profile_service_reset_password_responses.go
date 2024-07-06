// Code generated by go-swagger; DO NOT EDIT.

package profile_service

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

// ProfileServiceResetPasswordReader is a Reader for the ProfileServiceResetPassword structure.
type ProfileServiceResetPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfileServiceResetPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProfileServiceResetPasswordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProfileServiceResetPasswordDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProfileServiceResetPasswordOK creates a ProfileServiceResetPasswordOK with default headers values
func NewProfileServiceResetPasswordOK() *ProfileServiceResetPasswordOK {
	return &ProfileServiceResetPasswordOK{}
}

/*
ProfileServiceResetPasswordOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProfileServiceResetPasswordOK struct {
	Payload models.HashicorpCloudIamResetPasswordResponse
}

// IsSuccess returns true when this profile service reset password o k response has a 2xx status code
func (o *ProfileServiceResetPasswordOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this profile service reset password o k response has a 3xx status code
func (o *ProfileServiceResetPasswordOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this profile service reset password o k response has a 4xx status code
func (o *ProfileServiceResetPasswordOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this profile service reset password o k response has a 5xx status code
func (o *ProfileServiceResetPasswordOK) IsServerError() bool {
	return false
}

// IsCode returns true when this profile service reset password o k response a status code equal to that given
func (o *ProfileServiceResetPasswordOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProfileServiceResetPasswordOK) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/me/reset-password][%d] profileServiceResetPasswordOK  %+v", 200, o.Payload)
}

func (o *ProfileServiceResetPasswordOK) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/me/reset-password][%d] profileServiceResetPasswordOK  %+v", 200, o.Payload)
}

func (o *ProfileServiceResetPasswordOK) GetPayload() models.HashicorpCloudIamResetPasswordResponse {
	return o.Payload
}

func (o *ProfileServiceResetPasswordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfileServiceResetPasswordDefault creates a ProfileServiceResetPasswordDefault with default headers values
func NewProfileServiceResetPasswordDefault(code int) *ProfileServiceResetPasswordDefault {
	return &ProfileServiceResetPasswordDefault{
		_statusCode: code,
	}
}

/*
ProfileServiceResetPasswordDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProfileServiceResetPasswordDefault struct {
	_statusCode int

	Payload *cloud.GoogleRPCStatus
}

// Code gets the status code for the profile service reset password default response
func (o *ProfileServiceResetPasswordDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this profile service reset password default response has a 2xx status code
func (o *ProfileServiceResetPasswordDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this profile service reset password default response has a 3xx status code
func (o *ProfileServiceResetPasswordDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this profile service reset password default response has a 4xx status code
func (o *ProfileServiceResetPasswordDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this profile service reset password default response has a 5xx status code
func (o *ProfileServiceResetPasswordDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this profile service reset password default response a status code equal to that given
func (o *ProfileServiceResetPasswordDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProfileServiceResetPasswordDefault) Error() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/me/reset-password][%d] ProfileService_ResetPassword default  %+v", o._statusCode, o.Payload)
}

func (o *ProfileServiceResetPasswordDefault) String() string {
	return fmt.Sprintf("[POST /iam/2019-12-10/me/reset-password][%d] ProfileService_ResetPassword default  %+v", o._statusCode, o.Payload)
}

func (o *ProfileServiceResetPasswordDefault) GetPayload() *cloud.GoogleRPCStatus {
	return o.Payload
}

func (o *ProfileServiceResetPasswordDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(cloud.GoogleRPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
