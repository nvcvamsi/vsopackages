// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

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

// NewServicePrincipalsServiceCreateWorkloadIdentityProviderParams creates a new ServicePrincipalsServiceCreateWorkloadIdentityProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceCreateWorkloadIdentityProviderParams() *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceCreateWorkloadIdentityProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceCreateWorkloadIdentityProviderParamsWithTimeout creates a new ServicePrincipalsServiceCreateWorkloadIdentityProviderParams object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceCreateWorkloadIdentityProviderParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceCreateWorkloadIdentityProviderParams{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceCreateWorkloadIdentityProviderParamsWithContext creates a new ServicePrincipalsServiceCreateWorkloadIdentityProviderParams object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceCreateWorkloadIdentityProviderParamsWithContext(ctx context.Context) *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceCreateWorkloadIdentityProviderParams{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceCreateWorkloadIdentityProviderParamsWithHTTPClient creates a new ServicePrincipalsServiceCreateWorkloadIdentityProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceCreateWorkloadIdentityProviderParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams {
	return &ServicePrincipalsServiceCreateWorkloadIdentityProviderParams{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceCreateWorkloadIdentityProviderParams contains all the parameters to send to the API endpoint

	for the service principals service create workload identity provider operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceCreateWorkloadIdentityProviderParams struct {

	// Body.
	Body *models.HashicorpCloudIamCreateWorkloadIdentityProviderRequest

	/* ParentResourceName.

	     parent_resource_name is the name of the service principal the workload
	identity provider will be created under.
	*/
	ParentResourceName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service create workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) WithDefaults() *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service create workload identity provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service create workload identity provider params
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service create workload identity provider params
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service create workload identity provider params
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) WithContext(ctx context.Context) *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service create workload identity provider params
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service create workload identity provider params
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service create workload identity provider params
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the service principals service create workload identity provider params
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) WithBody(body *models.HashicorpCloudIamCreateWorkloadIdentityProviderRequest) *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service principals service create workload identity provider params
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) SetBody(body *models.HashicorpCloudIamCreateWorkloadIdentityProviderRequest) {
	o.Body = body
}

// WithParentResourceName adds the parentResourceName to the service principals service create workload identity provider params
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) WithParentResourceName(parentResourceName string) *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams {
	o.SetParentResourceName(parentResourceName)
	return o
}

// SetParentResourceName adds the parentResourceName to the service principals service create workload identity provider params
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) SetParentResourceName(parentResourceName string) {
	o.ParentResourceName = parentResourceName
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceCreateWorkloadIdentityProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param parent_resource_name
	if err := r.SetPathParam("parent_resource_name", o.ParentResourceName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}