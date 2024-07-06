// Code generated by go-swagger; DO NOT EDIT.

package consul_service

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

// NewListVersions2Params creates a new ListVersions2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListVersions2Params() *ListVersions2Params {
	return &ListVersions2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListVersions2ParamsWithTimeout creates a new ListVersions2Params object
// with the ability to set a timeout on a request.
func NewListVersions2ParamsWithTimeout(timeout time.Duration) *ListVersions2Params {
	return &ListVersions2Params{
		timeout: timeout,
	}
}

// NewListVersions2ParamsWithContext creates a new ListVersions2Params object
// with the ability to set a context for a request.
func NewListVersions2ParamsWithContext(ctx context.Context) *ListVersions2Params {
	return &ListVersions2Params{
		Context: ctx,
	}
}

// NewListVersions2ParamsWithHTTPClient creates a new ListVersions2Params object
// with the ability to set a custom HTTPClient for a request.
func NewListVersions2ParamsWithHTTPClient(client *http.Client) *ListVersions2Params {
	return &ListVersions2Params{
		HTTPClient: client,
	}
}

/*
ListVersions2Params contains all the parameters to send to the API endpoint

	for the list versions2 operation.

	Typically these are written to a http.Request.
*/
type ListVersions2Params struct {

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID *string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID *string

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

// WithDefaults hydrates default values in the list versions2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVersions2Params) WithDefaults() *ListVersions2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list versions2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListVersions2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list versions2 params
func (o *ListVersions2Params) WithTimeout(timeout time.Duration) *ListVersions2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list versions2 params
func (o *ListVersions2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list versions2 params
func (o *ListVersions2Params) WithContext(ctx context.Context) *ListVersions2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list versions2 params
func (o *ListVersions2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list versions2 params
func (o *ListVersions2Params) WithHTTPClient(client *http.Client) *ListVersions2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list versions2 params
func (o *ListVersions2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationOrganizationID adds the locationOrganizationID to the list versions2 params
func (o *ListVersions2Params) WithLocationOrganizationID(locationOrganizationID *string) *ListVersions2Params {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the list versions2 params
func (o *ListVersions2Params) SetLocationOrganizationID(locationOrganizationID *string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the list versions2 params
func (o *ListVersions2Params) WithLocationProjectID(locationProjectID *string) *ListVersions2Params {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the list versions2 params
func (o *ListVersions2Params) SetLocationProjectID(locationProjectID *string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the list versions2 params
func (o *ListVersions2Params) WithLocationRegionProvider(locationRegionProvider *string) *ListVersions2Params {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the list versions2 params
func (o *ListVersions2Params) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the list versions2 params
func (o *ListVersions2Params) WithLocationRegionRegion(locationRegionRegion *string) *ListVersions2Params {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the list versions2 params
func (o *ListVersions2Params) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *ListVersions2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LocationOrganizationID != nil {

		// query param location.organization_id
		var qrLocationOrganizationID string

		if o.LocationOrganizationID != nil {
			qrLocationOrganizationID = *o.LocationOrganizationID
		}
		qLocationOrganizationID := qrLocationOrganizationID
		if qLocationOrganizationID != "" {

			if err := r.SetQueryParam("location.organization_id", qLocationOrganizationID); err != nil {
				return err
			}
		}
	}

	if o.LocationProjectID != nil {

		// query param location.project_id
		var qrLocationProjectID string

		if o.LocationProjectID != nil {
			qrLocationProjectID = *o.LocationProjectID
		}
		qLocationProjectID := qrLocationProjectID
		if qLocationProjectID != "" {

			if err := r.SetQueryParam("location.project_id", qLocationProjectID); err != nil {
				return err
			}
		}
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
