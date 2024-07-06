// Code generated by go-swagger; DO NOT EDIT.

package packer_service

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
	"github.com/go-openapi/swag"
)

// NewPackerServiceDeleteIterationParams creates a new PackerServiceDeleteIterationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceDeleteIterationParams() *PackerServiceDeleteIterationParams {
	return &PackerServiceDeleteIterationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceDeleteIterationParamsWithTimeout creates a new PackerServiceDeleteIterationParams object
// with the ability to set a timeout on a request.
func NewPackerServiceDeleteIterationParamsWithTimeout(timeout time.Duration) *PackerServiceDeleteIterationParams {
	return &PackerServiceDeleteIterationParams{
		timeout: timeout,
	}
}

// NewPackerServiceDeleteIterationParamsWithContext creates a new PackerServiceDeleteIterationParams object
// with the ability to set a context for a request.
func NewPackerServiceDeleteIterationParamsWithContext(ctx context.Context) *PackerServiceDeleteIterationParams {
	return &PackerServiceDeleteIterationParams{
		Context: ctx,
	}
}

// NewPackerServiceDeleteIterationParamsWithHTTPClient creates a new PackerServiceDeleteIterationParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceDeleteIterationParamsWithHTTPClient(client *http.Client) *PackerServiceDeleteIterationParams {
	return &PackerServiceDeleteIterationParams{
		HTTPClient: client,
	}
}

/*
PackerServiceDeleteIterationParams contains all the parameters to send to the API endpoint

	for the packer service delete iteration operation.

	Typically these are written to a http.Request.
*/
type PackerServiceDeleteIterationParams struct {

	/* BucketSlug.

	   Human-readable name for the bucket.
	*/
	BucketSlug *string

	/* IterationID.

	     ULID of the iteration. This was created and set by the
	HCP Packer registry when the iteration was created.
	*/
	IterationID string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* LocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure")
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1")
	*/
	LocationRegionRegion *string

	/* RollbackChannels.

	   When set to true, any user created channels will be rolled back to the last valid iteration they were assigned to.
	*/
	RollbackChannels *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the packer service delete iteration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceDeleteIterationParams) WithDefaults() *PackerServiceDeleteIterationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service delete iteration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceDeleteIterationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) WithTimeout(timeout time.Duration) *PackerServiceDeleteIterationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) WithContext(ctx context.Context) *PackerServiceDeleteIterationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) WithHTTPClient(client *http.Client) *PackerServiceDeleteIterationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlug adds the bucketSlug to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) WithBucketSlug(bucketSlug *string) *PackerServiceDeleteIterationParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) SetBucketSlug(bucketSlug *string) {
	o.BucketSlug = bucketSlug
}

// WithIterationID adds the iterationID to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) WithIterationID(iterationID string) *PackerServiceDeleteIterationParams {
	o.SetIterationID(iterationID)
	return o
}

// SetIterationID adds the iterationId to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) SetIterationID(iterationID string) {
	o.IterationID = iterationID
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceDeleteIterationParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) WithLocationProjectID(locationProjectID string) *PackerServiceDeleteIterationParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) WithLocationRegionProvider(locationRegionProvider *string) *PackerServiceDeleteIterationParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) WithLocationRegionRegion(locationRegionRegion *string) *PackerServiceDeleteIterationParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithRollbackChannels adds the rollbackChannels to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) WithRollbackChannels(rollbackChannels *bool) *PackerServiceDeleteIterationParams {
	o.SetRollbackChannels(rollbackChannels)
	return o
}

// SetRollbackChannels adds the rollbackChannels to the packer service delete iteration params
func (o *PackerServiceDeleteIterationParams) SetRollbackChannels(rollbackChannels *bool) {
	o.RollbackChannels = rollbackChannels
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceDeleteIterationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BucketSlug != nil {

		// query param bucket_slug
		var qrBucketSlug string

		if o.BucketSlug != nil {
			qrBucketSlug = *o.BucketSlug
		}
		qBucketSlug := qrBucketSlug
		if qBucketSlug != "" {

			if err := r.SetQueryParam("bucket_slug", qBucketSlug); err != nil {
				return err
			}
		}
	}

	// path param iteration_id
	if err := r.SetPathParam("iteration_id", o.IterationID); err != nil {
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

	if o.RollbackChannels != nil {

		// query param rollback_channels
		var qrRollbackChannels bool

		if o.RollbackChannels != nil {
			qrRollbackChannels = *o.RollbackChannels
		}
		qRollbackChannels := swag.FormatBool(qrRollbackChannels)
		if qRollbackChannels != "" {

			if err := r.SetQueryParam("rollback_channels", qRollbackChannels); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
