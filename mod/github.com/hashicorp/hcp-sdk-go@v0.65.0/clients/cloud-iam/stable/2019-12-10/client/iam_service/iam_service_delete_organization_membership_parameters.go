// Code generated by go-swagger; DO NOT EDIT.

package iam_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewIamServiceDeleteOrganizationMembershipParams creates a new IamServiceDeleteOrganizationMembershipParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIamServiceDeleteOrganizationMembershipParams() *IamServiceDeleteOrganizationMembershipParams {
	return &IamServiceDeleteOrganizationMembershipParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIamServiceDeleteOrganizationMembershipParamsWithTimeout creates a new IamServiceDeleteOrganizationMembershipParams object
// with the ability to set a timeout on a request.
func NewIamServiceDeleteOrganizationMembershipParamsWithTimeout(timeout time.Duration) *IamServiceDeleteOrganizationMembershipParams {
	return &IamServiceDeleteOrganizationMembershipParams{
		timeout: timeout,
	}
}

// NewIamServiceDeleteOrganizationMembershipParamsWithContext creates a new IamServiceDeleteOrganizationMembershipParams object
// with the ability to set a context for a request.
func NewIamServiceDeleteOrganizationMembershipParamsWithContext(ctx context.Context) *IamServiceDeleteOrganizationMembershipParams {
	return &IamServiceDeleteOrganizationMembershipParams{
		Context: ctx,
	}
}

// NewIamServiceDeleteOrganizationMembershipParamsWithHTTPClient creates a new IamServiceDeleteOrganizationMembershipParams object
// with the ability to set a custom HTTPClient for a request.
func NewIamServiceDeleteOrganizationMembershipParamsWithHTTPClient(client *http.Client) *IamServiceDeleteOrganizationMembershipParams {
	return &IamServiceDeleteOrganizationMembershipParams{
		HTTPClient: client,
	}
}

/*
IamServiceDeleteOrganizationMembershipParams contains all the parameters to send to the API endpoint

	for the iam service delete organization membership operation.

	Typically these are written to a http.Request.
*/
type IamServiceDeleteOrganizationMembershipParams struct {

	/* OrganizationID.

	   organization_id is the ID of the organization the user principal is being deleted from.
	*/
	OrganizationID string

	/* UserPrincipalID.

	   user_principal_id is the ID of the user principal to delete from the organization.
	*/
	UserPrincipalID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iam service delete organization membership params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamServiceDeleteOrganizationMembershipParams) WithDefaults() *IamServiceDeleteOrganizationMembershipParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iam service delete organization membership params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamServiceDeleteOrganizationMembershipParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the iam service delete organization membership params
func (o *IamServiceDeleteOrganizationMembershipParams) WithTimeout(timeout time.Duration) *IamServiceDeleteOrganizationMembershipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iam service delete organization membership params
func (o *IamServiceDeleteOrganizationMembershipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iam service delete organization membership params
func (o *IamServiceDeleteOrganizationMembershipParams) WithContext(ctx context.Context) *IamServiceDeleteOrganizationMembershipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iam service delete organization membership params
func (o *IamServiceDeleteOrganizationMembershipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iam service delete organization membership params
func (o *IamServiceDeleteOrganizationMembershipParams) WithHTTPClient(client *http.Client) *IamServiceDeleteOrganizationMembershipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iam service delete organization membership params
func (o *IamServiceDeleteOrganizationMembershipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the iam service delete organization membership params
func (o *IamServiceDeleteOrganizationMembershipParams) WithOrganizationID(organizationID string) *IamServiceDeleteOrganizationMembershipParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the iam service delete organization membership params
func (o *IamServiceDeleteOrganizationMembershipParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithUserPrincipalID adds the userPrincipalID to the iam service delete organization membership params
func (o *IamServiceDeleteOrganizationMembershipParams) WithUserPrincipalID(userPrincipalID string) *IamServiceDeleteOrganizationMembershipParams {
	o.SetUserPrincipalID(userPrincipalID)
	return o
}

// SetUserPrincipalID adds the userPrincipalId to the iam service delete organization membership params
func (o *IamServiceDeleteOrganizationMembershipParams) SetUserPrincipalID(userPrincipalID string) {
	o.UserPrincipalID = userPrincipalID
}

// WriteToRequest writes these params to a swagger request
func (o *IamServiceDeleteOrganizationMembershipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	// path param user_principal_id
	if err := r.SetPathParam("user_principal_id", o.UserPrincipalID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}