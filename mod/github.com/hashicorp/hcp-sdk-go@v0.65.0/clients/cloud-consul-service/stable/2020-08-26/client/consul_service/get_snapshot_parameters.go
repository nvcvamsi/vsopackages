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

// NewGetSnapshotParams creates a new GetSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSnapshotParams() *GetSnapshotParams {
	return &GetSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotParamsWithTimeout creates a new GetSnapshotParams object
// with the ability to set a timeout on a request.
func NewGetSnapshotParamsWithTimeout(timeout time.Duration) *GetSnapshotParams {
	return &GetSnapshotParams{
		timeout: timeout,
	}
}

// NewGetSnapshotParamsWithContext creates a new GetSnapshotParams object
// with the ability to set a context for a request.
func NewGetSnapshotParamsWithContext(ctx context.Context) *GetSnapshotParams {
	return &GetSnapshotParams{
		Context: ctx,
	}
}

// NewGetSnapshotParamsWithHTTPClient creates a new GetSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSnapshotParamsWithHTTPClient(client *http.Client) *GetSnapshotParams {
	return &GetSnapshotParams{
		HTTPClient: client,
	}
}

/*
GetSnapshotParams contains all the parameters to send to the API endpoint

	for the get snapshot operation.

	Typically these are written to a http.Request.
*/
type GetSnapshotParams struct {

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

	/* SnapshotID.

	   snapshot_id represents the snapshot to get.
	*/
	SnapshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotParams) WithDefaults() *GetSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get snapshot params
func (o *GetSnapshotParams) WithTimeout(timeout time.Duration) *GetSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot params
func (o *GetSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot params
func (o *GetSnapshotParams) WithContext(ctx context.Context) *GetSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot params
func (o *GetSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot params
func (o *GetSnapshotParams) WithHTTPClient(client *http.Client) *GetSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot params
func (o *GetSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationOrganizationID adds the locationOrganizationID to the get snapshot params
func (o *GetSnapshotParams) WithLocationOrganizationID(locationOrganizationID string) *GetSnapshotParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get snapshot params
func (o *GetSnapshotParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get snapshot params
func (o *GetSnapshotParams) WithLocationProjectID(locationProjectID string) *GetSnapshotParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get snapshot params
func (o *GetSnapshotParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the get snapshot params
func (o *GetSnapshotParams) WithLocationRegionProvider(locationRegionProvider *string) *GetSnapshotParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the get snapshot params
func (o *GetSnapshotParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the get snapshot params
func (o *GetSnapshotParams) WithLocationRegionRegion(locationRegionRegion *string) *GetSnapshotParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the get snapshot params
func (o *GetSnapshotParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithSnapshotID adds the snapshotID to the get snapshot params
func (o *GetSnapshotParams) WithSnapshotID(snapshotID string) *GetSnapshotParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the get snapshot params
func (o *GetSnapshotParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param snapshot_id
	if err := r.SetPathParam("snapshot_id", o.SnapshotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
