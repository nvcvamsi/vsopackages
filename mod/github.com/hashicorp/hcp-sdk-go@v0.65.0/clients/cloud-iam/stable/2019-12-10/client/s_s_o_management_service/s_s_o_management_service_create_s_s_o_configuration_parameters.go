// Code generated by go-swagger; DO NOT EDIT.

package s_s_o_management_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
)

// NewSSOManagementServiceCreateSSOConfigurationParams creates a new SSOManagementServiceCreateSSOConfigurationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSSOManagementServiceCreateSSOConfigurationParams() *SSOManagementServiceCreateSSOConfigurationParams {
	return &SSOManagementServiceCreateSSOConfigurationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSSOManagementServiceCreateSSOConfigurationParamsWithTimeout creates a new SSOManagementServiceCreateSSOConfigurationParams object
// with the ability to set a timeout on a request.
func NewSSOManagementServiceCreateSSOConfigurationParamsWithTimeout(timeout time.Duration) *SSOManagementServiceCreateSSOConfigurationParams {
	return &SSOManagementServiceCreateSSOConfigurationParams{
		timeout: timeout,
	}
}

// NewSSOManagementServiceCreateSSOConfigurationParamsWithContext creates a new SSOManagementServiceCreateSSOConfigurationParams object
// with the ability to set a context for a request.
func NewSSOManagementServiceCreateSSOConfigurationParamsWithContext(ctx context.Context) *SSOManagementServiceCreateSSOConfigurationParams {
	return &SSOManagementServiceCreateSSOConfigurationParams{
		Context: ctx,
	}
}

// NewSSOManagementServiceCreateSSOConfigurationParamsWithHTTPClient creates a new SSOManagementServiceCreateSSOConfigurationParams object
// with the ability to set a custom HTTPClient for a request.
func NewSSOManagementServiceCreateSSOConfigurationParamsWithHTTPClient(client *http.Client) *SSOManagementServiceCreateSSOConfigurationParams {
	return &SSOManagementServiceCreateSSOConfigurationParams{
		HTTPClient: client,
	}
}

/*
SSOManagementServiceCreateSSOConfigurationParams contains all the parameters to send to the API endpoint

	for the s s o management service create s s o configuration operation.

	Typically these are written to a http.Request.
*/
type SSOManagementServiceCreateSSOConfigurationParams struct {

	// Body.
	Body *models.HashicorpCloudIamCreateSSOConfigurationRequest

	/* OrganizationID.

	   OrganizationId is the identifier of the organization.
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s s o management service create s s o configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSOManagementServiceCreateSSOConfigurationParams) WithDefaults() *SSOManagementServiceCreateSSOConfigurationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s s o management service create s s o configuration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SSOManagementServiceCreateSSOConfigurationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s s o management service create s s o configuration params
func (o *SSOManagementServiceCreateSSOConfigurationParams) WithTimeout(timeout time.Duration) *SSOManagementServiceCreateSSOConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s s o management service create s s o configuration params
func (o *SSOManagementServiceCreateSSOConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s s o management service create s s o configuration params
func (o *SSOManagementServiceCreateSSOConfigurationParams) WithContext(ctx context.Context) *SSOManagementServiceCreateSSOConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s s o management service create s s o configuration params
func (o *SSOManagementServiceCreateSSOConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s s o management service create s s o configuration params
func (o *SSOManagementServiceCreateSSOConfigurationParams) WithHTTPClient(client *http.Client) *SSOManagementServiceCreateSSOConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s s o management service create s s o configuration params
func (o *SSOManagementServiceCreateSSOConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the s s o management service create s s o configuration params
func (o *SSOManagementServiceCreateSSOConfigurationParams) WithBody(body *models.HashicorpCloudIamCreateSSOConfigurationRequest) *SSOManagementServiceCreateSSOConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the s s o management service create s s o configuration params
func (o *SSOManagementServiceCreateSSOConfigurationParams) SetBody(body *models.HashicorpCloudIamCreateSSOConfigurationRequest) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the s s o management service create s s o configuration params
func (o *SSOManagementServiceCreateSSOConfigurationParams) WithOrganizationID(organizationID string) *SSOManagementServiceCreateSSOConfigurationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the s s o management service create s s o configuration params
func (o *SSOManagementServiceCreateSSOConfigurationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *SSOManagementServiceCreateSSOConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}