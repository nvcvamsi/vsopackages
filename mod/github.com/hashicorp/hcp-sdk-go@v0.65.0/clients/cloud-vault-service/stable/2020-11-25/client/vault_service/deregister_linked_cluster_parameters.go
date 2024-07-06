// Code generated by go-swagger; DO NOT EDIT.

package vault_service

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

// NewDeregisterLinkedClusterParams creates a new DeregisterLinkedClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeregisterLinkedClusterParams() *DeregisterLinkedClusterParams {
	return &DeregisterLinkedClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeregisterLinkedClusterParamsWithTimeout creates a new DeregisterLinkedClusterParams object
// with the ability to set a timeout on a request.
func NewDeregisterLinkedClusterParamsWithTimeout(timeout time.Duration) *DeregisterLinkedClusterParams {
	return &DeregisterLinkedClusterParams{
		timeout: timeout,
	}
}

// NewDeregisterLinkedClusterParamsWithContext creates a new DeregisterLinkedClusterParams object
// with the ability to set a context for a request.
func NewDeregisterLinkedClusterParamsWithContext(ctx context.Context) *DeregisterLinkedClusterParams {
	return &DeregisterLinkedClusterParams{
		Context: ctx,
	}
}

// NewDeregisterLinkedClusterParamsWithHTTPClient creates a new DeregisterLinkedClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeregisterLinkedClusterParamsWithHTTPClient(client *http.Client) *DeregisterLinkedClusterParams {
	return &DeregisterLinkedClusterParams{
		HTTPClient: client,
	}
}

/*
DeregisterLinkedClusterParams contains all the parameters to send to the API endpoint

	for the deregister linked cluster operation.

	Typically these are written to a http.Request.
*/
type DeregisterLinkedClusterParams struct {

	// ClusterID.
	ClusterID string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* LocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure").
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1").
	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the deregister linked cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeregisterLinkedClusterParams) WithDefaults() *DeregisterLinkedClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the deregister linked cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeregisterLinkedClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) WithTimeout(timeout time.Duration) *DeregisterLinkedClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) WithContext(ctx context.Context) *DeregisterLinkedClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) WithHTTPClient(client *http.Client) *DeregisterLinkedClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) WithClusterID(clusterID string) *DeregisterLinkedClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) WithLocationOrganizationID(locationOrganizationID string) *DeregisterLinkedClusterParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) WithLocationProjectID(locationProjectID string) *DeregisterLinkedClusterParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) WithLocationRegionProvider(locationRegionProvider *string) *DeregisterLinkedClusterParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) WithLocationRegionRegion(locationRegionRegion *string) *DeregisterLinkedClusterParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the deregister linked cluster params
func (o *DeregisterLinkedClusterParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *DeregisterLinkedClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if o.LocationRegionProvider != nil {

		// query param location.region.provider
		var qrLocationRegionProvider string

		if o.LocationRegionProvider != nil {
			qrLocationRegionProvider = *o.LocationRegionProvider
		}
		qLocationRegionProvider := qrLocationRegionProvider
		if qLocationRegionProvider != "" {

			if err := r.SetQueryParam("location.region.provider", qLocationRegionProvider); err != nil {
				return err
			}
		}
	}

	if o.LocationRegionRegion != nil {

		// query param location.region.region
		var qrLocationRegionRegion string

		if o.LocationRegionRegion != nil {
			qrLocationRegionRegion = *o.LocationRegionRegion
		}
		qLocationRegionRegion := qrLocationRegionRegion
		if qLocationRegionRegion != "" {

			if err := r.SetQueryParam("location.region.region", qLocationRegionRegion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
