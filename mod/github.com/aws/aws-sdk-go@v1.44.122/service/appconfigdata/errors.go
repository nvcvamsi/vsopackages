// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfigdata

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// The input fails to satisfy the constraints specified by the service.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// There was an internal failure in the service.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The requested resource could not be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request was denied due to request throttling.
	ErrCodeThrottlingException = "ThrottlingException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":       newErrorBadRequestException,
	"InternalServerException":   newErrorInternalServerException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"ThrottlingException":       newErrorThrottlingException,
}
