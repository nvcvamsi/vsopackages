// Code generated by go-swagger; DO NOT EDIT.

package operation_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operation service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operation service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOK, error)

	List(params *ListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOK, error)

	Wait(params *WaitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WaitOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	Get gets returns the latest state of a specific operation

	The Location's organization and project IDs are included in the path in

order to support RBAC checks. We perform RBAC checks early in the request
lifecycle, before loading the resource targeted by the request.
*/
func (a *Client) Get(params *GetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get",
		Method:             "GET",
		PathPattern:        "/operation/2020-05-05/organizations/{location.organization_id}/projects/{location.project_id}/operations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	List lists selects a list of operations that match some filters and then returns them in the response

	The supported filters are the Location's organization_id and project_id.

The project_id supports the special value "-" to allow requesting operations
that match any project within the organization.
*/
func (a *Client) List(params *ListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "List",
		Method:             "GET",
		PathPattern:        "/operation/2020-05-05/organizations/{location.organization_id}/projects/{location.project_id}/operations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	Wait waits blocks until an operation is d o n e or a timeout is reached whichever occurs first it then returns the latest state of the operation that it had at the time of the timeout

	Note that Wait may return at any time. In most cases, the timeout will

be respected but it isn't guaranteed. Therefore, clients should always
check the state of the returned operation. A return from Wait does not
guarantee the operation is DONE.

The Location's organization and project IDs are included in the path in
order to support RBAC checks. We perform RBAC checks early in the request
lifecycle, before loading the resource targeted by the request.
*/
func (a *Client) Wait(params *WaitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*WaitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWaitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Wait",
		Method:             "GET",
		PathPattern:        "/operation/2020-05-05/organizations/{location.organization_id}/projects/{location.project_id}/operations/{id}/wait",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &WaitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WaitOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*WaitDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}