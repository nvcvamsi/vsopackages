// Code generated by go-swagger; DO NOT EDIT.

package box_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// NewCreateBoxParams creates a new CreateBoxParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBoxParams() *CreateBoxParams {
	return &CreateBoxParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBoxParamsWithTimeout creates a new CreateBoxParams object
// with the ability to set a timeout on a request.
func NewCreateBoxParamsWithTimeout(timeout time.Duration) *CreateBoxParams {
	return &CreateBoxParams{
		timeout: timeout,
	}
}

// NewCreateBoxParamsWithContext creates a new CreateBoxParams object
// with the ability to set a context for a request.
func NewCreateBoxParamsWithContext(ctx context.Context) *CreateBoxParams {
	return &CreateBoxParams{
		Context: ctx,
	}
}

// NewCreateBoxParamsWithHTTPClient creates a new CreateBoxParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBoxParamsWithHTTPClient(client *http.Client) *CreateBoxParams {
	return &CreateBoxParams{
		HTTPClient: client,
	}
}

/*
CreateBoxParams contains all the parameters to send to the API endpoint

	for the create box operation.

	Typically these are written to a http.Request.
*/
type CreateBoxParams struct {

	// Body.
	Body *models.HashicorpCloudVagrantCreateBoxRequest

	/* Registry.

	     The Registry segment of the Box. As an example, this field would represent
	the "hashicorp" in "hashicorp/vagrant".
	*/
	Registry string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create box params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBoxParams) WithDefaults() *CreateBoxParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create box params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBoxParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create box params
func (o *CreateBoxParams) WithTimeout(timeout time.Duration) *CreateBoxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create box params
func (o *CreateBoxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create box params
func (o *CreateBoxParams) WithContext(ctx context.Context) *CreateBoxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create box params
func (o *CreateBoxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create box params
func (o *CreateBoxParams) WithHTTPClient(client *http.Client) *CreateBoxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create box params
func (o *CreateBoxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create box params
func (o *CreateBoxParams) WithBody(body *models.HashicorpCloudVagrantCreateBoxRequest) *CreateBoxParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create box params
func (o *CreateBoxParams) SetBody(body *models.HashicorpCloudVagrantCreateBoxRequest) {
	o.Body = body
}

// WithRegistry adds the registry to the create box params
func (o *CreateBoxParams) WithRegistry(registry string) *CreateBoxParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the create box params
func (o *CreateBoxParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBoxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param registry
	if err := r.SetPathParam("registry", o.Registry); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
