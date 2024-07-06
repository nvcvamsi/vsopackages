// Code generated by go-swagger; DO NOT EDIT.

package network_service

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

// NewListHVNRoutesParams creates a new ListHVNRoutesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListHVNRoutesParams() *ListHVNRoutesParams {
	return &ListHVNRoutesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListHVNRoutesParamsWithTimeout creates a new ListHVNRoutesParams object
// with the ability to set a timeout on a request.
func NewListHVNRoutesParamsWithTimeout(timeout time.Duration) *ListHVNRoutesParams {
	return &ListHVNRoutesParams{
		timeout: timeout,
	}
}

// NewListHVNRoutesParamsWithContext creates a new ListHVNRoutesParams object
// with the ability to set a context for a request.
func NewListHVNRoutesParamsWithContext(ctx context.Context) *ListHVNRoutesParams {
	return &ListHVNRoutesParams{
		Context: ctx,
	}
}

// NewListHVNRoutesParamsWithHTTPClient creates a new ListHVNRoutesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListHVNRoutesParamsWithHTTPClient(client *http.Client) *ListHVNRoutesParams {
	return &ListHVNRoutesParams{
		HTTPClient: client,
	}
}

/*
ListHVNRoutesParams contains all the parameters to send to the API endpoint

	for the list h v n routes operation.

	Typically these are written to a http.Request.
*/
type ListHVNRoutesParams struct {

	/* Destination.

	     destination is an optional parameter that allows returning only routes
	that has a specific destination.
	*/
	Destination *string

	/* HvnDescription.

	     description is a human-friendly description for this link. This is
	used primarily for informational purposes such as error messages.
	*/
	HvnDescription *string

	/* HvnID.

	   id is the identifier for this resource.
	*/
	HvnID string

	/* HvnLocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	HvnLocationOrganizationID string

	/* HvnLocationProjectID.

	   project_id is the projects id.
	*/
	HvnLocationProjectID string

	/* HvnLocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure").
	*/
	HvnLocationRegionProvider *string

	/* HvnLocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1").
	*/
	HvnLocationRegionRegion *string

	/* HvnType.

	     type is the unique type of the resource. Each service publishes a
	unique set of types. The type value is recommended to be formatted
	in "<org>.<type>" such as "hashicorp.hvn". This is to prevent conflicts
	in the future, but any string value will work.
	*/
	HvnType *string

	/* HvnUUID.

	   uuid is the unique UUID for this resource.
	*/
	HvnUUID *string

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	/* TargetID.

	     target_id is an optional parameter that allows returning only routes
	that has a specific target Id.
	*/
	TargetID *string

	/* TargetType.

	     target_type is an optional parameter that allows returning only routes
	that has a specific target Type.
	*/
	TargetType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list h v n routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListHVNRoutesParams) WithDefaults() *ListHVNRoutesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list h v n routes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListHVNRoutesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list h v n routes params
func (o *ListHVNRoutesParams) WithTimeout(timeout time.Duration) *ListHVNRoutesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list h v n routes params
func (o *ListHVNRoutesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list h v n routes params
func (o *ListHVNRoutesParams) WithContext(ctx context.Context) *ListHVNRoutesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list h v n routes params
func (o *ListHVNRoutesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list h v n routes params
func (o *ListHVNRoutesParams) WithHTTPClient(client *http.Client) *ListHVNRoutesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list h v n routes params
func (o *ListHVNRoutesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDestination adds the destination to the list h v n routes params
func (o *ListHVNRoutesParams) WithDestination(destination *string) *ListHVNRoutesParams {
	o.SetDestination(destination)
	return o
}

// SetDestination adds the destination to the list h v n routes params
func (o *ListHVNRoutesParams) SetDestination(destination *string) {
	o.Destination = destination
}

// WithHvnDescription adds the hvnDescription to the list h v n routes params
func (o *ListHVNRoutesParams) WithHvnDescription(hvnDescription *string) *ListHVNRoutesParams {
	o.SetHvnDescription(hvnDescription)
	return o
}

// SetHvnDescription adds the hvnDescription to the list h v n routes params
func (o *ListHVNRoutesParams) SetHvnDescription(hvnDescription *string) {
	o.HvnDescription = hvnDescription
}

// WithHvnID adds the hvnID to the list h v n routes params
func (o *ListHVNRoutesParams) WithHvnID(hvnID string) *ListHVNRoutesParams {
	o.SetHvnID(hvnID)
	return o
}

// SetHvnID adds the hvnId to the list h v n routes params
func (o *ListHVNRoutesParams) SetHvnID(hvnID string) {
	o.HvnID = hvnID
}

// WithHvnLocationOrganizationID adds the hvnLocationOrganizationID to the list h v n routes params
func (o *ListHVNRoutesParams) WithHvnLocationOrganizationID(hvnLocationOrganizationID string) *ListHVNRoutesParams {
	o.SetHvnLocationOrganizationID(hvnLocationOrganizationID)
	return o
}

// SetHvnLocationOrganizationID adds the hvnLocationOrganizationId to the list h v n routes params
func (o *ListHVNRoutesParams) SetHvnLocationOrganizationID(hvnLocationOrganizationID string) {
	o.HvnLocationOrganizationID = hvnLocationOrganizationID
}

// WithHvnLocationProjectID adds the hvnLocationProjectID to the list h v n routes params
func (o *ListHVNRoutesParams) WithHvnLocationProjectID(hvnLocationProjectID string) *ListHVNRoutesParams {
	o.SetHvnLocationProjectID(hvnLocationProjectID)
	return o
}

// SetHvnLocationProjectID adds the hvnLocationProjectId to the list h v n routes params
func (o *ListHVNRoutesParams) SetHvnLocationProjectID(hvnLocationProjectID string) {
	o.HvnLocationProjectID = hvnLocationProjectID
}

// WithHvnLocationRegionProvider adds the hvnLocationRegionProvider to the list h v n routes params
func (o *ListHVNRoutesParams) WithHvnLocationRegionProvider(hvnLocationRegionProvider *string) *ListHVNRoutesParams {
	o.SetHvnLocationRegionProvider(hvnLocationRegionProvider)
	return o
}

// SetHvnLocationRegionProvider adds the hvnLocationRegionProvider to the list h v n routes params
func (o *ListHVNRoutesParams) SetHvnLocationRegionProvider(hvnLocationRegionProvider *string) {
	o.HvnLocationRegionProvider = hvnLocationRegionProvider
}

// WithHvnLocationRegionRegion adds the hvnLocationRegionRegion to the list h v n routes params
func (o *ListHVNRoutesParams) WithHvnLocationRegionRegion(hvnLocationRegionRegion *string) *ListHVNRoutesParams {
	o.SetHvnLocationRegionRegion(hvnLocationRegionRegion)
	return o
}

// SetHvnLocationRegionRegion adds the hvnLocationRegionRegion to the list h v n routes params
func (o *ListHVNRoutesParams) SetHvnLocationRegionRegion(hvnLocationRegionRegion *string) {
	o.HvnLocationRegionRegion = hvnLocationRegionRegion
}

// WithHvnType adds the hvnType to the list h v n routes params
func (o *ListHVNRoutesParams) WithHvnType(hvnType *string) *ListHVNRoutesParams {
	o.SetHvnType(hvnType)
	return o
}

// SetHvnType adds the hvnType to the list h v n routes params
func (o *ListHVNRoutesParams) SetHvnType(hvnType *string) {
	o.HvnType = hvnType
}

// WithHvnUUID adds the hvnUUID to the list h v n routes params
func (o *ListHVNRoutesParams) WithHvnUUID(hvnUUID *string) *ListHVNRoutesParams {
	o.SetHvnUUID(hvnUUID)
	return o
}

// SetHvnUUID adds the hvnUuid to the list h v n routes params
func (o *ListHVNRoutesParams) SetHvnUUID(hvnUUID *string) {
	o.HvnUUID = hvnUUID
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the list h v n routes params
func (o *ListHVNRoutesParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ListHVNRoutesParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the list h v n routes params
func (o *ListHVNRoutesParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the list h v n routes params
func (o *ListHVNRoutesParams) WithPaginationPageSize(paginationPageSize *int64) *ListHVNRoutesParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the list h v n routes params
func (o *ListHVNRoutesParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the list h v n routes params
func (o *ListHVNRoutesParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ListHVNRoutesParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the list h v n routes params
func (o *ListHVNRoutesParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithTargetID adds the targetID to the list h v n routes params
func (o *ListHVNRoutesParams) WithTargetID(targetID *string) *ListHVNRoutesParams {
	o.SetTargetID(targetID)
	return o
}

// SetTargetID adds the targetId to the list h v n routes params
func (o *ListHVNRoutesParams) SetTargetID(targetID *string) {
	o.TargetID = targetID
}

// WithTargetType adds the targetType to the list h v n routes params
func (o *ListHVNRoutesParams) WithTargetType(targetType *string) *ListHVNRoutesParams {
	o.SetTargetType(targetType)
	return o
}

// SetTargetType adds the targetType to the list h v n routes params
func (o *ListHVNRoutesParams) SetTargetType(targetType *string) {
	o.TargetType = targetType
}

// WriteToRequest writes these params to a swagger request
func (o *ListHVNRoutesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Destination != nil {

		// query param destination
		var qrDestination string

		if o.Destination != nil {
			qrDestination = *o.Destination
		}
		qDestination := qrDestination
		if qDestination != "" {

			if err := r.SetQueryParam("destination", qDestination); err != nil {
				return err
			}
		}
	}

	if o.HvnDescription != nil {

		// query param hvn.description
		var qrHvnDescription string

		if o.HvnDescription != nil {
			qrHvnDescription = *o.HvnDescription
		}
		qHvnDescription := qrHvnDescription
		if qHvnDescription != "" {

			if err := r.SetQueryParam("hvn.description", qHvnDescription); err != nil {
				return err
			}
		}
	}

	// path param hvn.id
	if err := r.SetPathParam("hvn.id", o.HvnID); err != nil {
		return err
	}

	// path param hvn.location.organization_id
	if err := r.SetPathParam("hvn.location.organization_id", o.HvnLocationOrganizationID); err != nil {
		return err
	}

	// path param hvn.location.project_id
	if err := r.SetPathParam("hvn.location.project_id", o.HvnLocationProjectID); err != nil {
		return err
	}

	if o.HvnLocationRegionProvider != nil {

		// query param hvn.location.region.provider
		var qrHvnLocationRegionProvider string

		if o.HvnLocationRegionProvider != nil {
			qrHvnLocationRegionProvider = *o.HvnLocationRegionProvider
		}
		qHvnLocationRegionProvider := qrHvnLocationRegionProvider
		if qHvnLocationRegionProvider != "" {

			if err := r.SetQueryParam("hvn.location.region.provider", qHvnLocationRegionProvider); err != nil {
				return err
			}
		}
	}

	if o.HvnLocationRegionRegion != nil {

		// query param hvn.location.region.region
		var qrHvnLocationRegionRegion string

		if o.HvnLocationRegionRegion != nil {
			qrHvnLocationRegionRegion = *o.HvnLocationRegionRegion
		}
		qHvnLocationRegionRegion := qrHvnLocationRegionRegion
		if qHvnLocationRegionRegion != "" {

			if err := r.SetQueryParam("hvn.location.region.region", qHvnLocationRegionRegion); err != nil {
				return err
			}
		}
	}

	if o.HvnType != nil {

		// query param hvn.type
		var qrHvnType string

		if o.HvnType != nil {
			qrHvnType = *o.HvnType
		}
		qHvnType := qrHvnType
		if qHvnType != "" {

			if err := r.SetQueryParam("hvn.type", qHvnType); err != nil {
				return err
			}
		}
	}

	if o.HvnUUID != nil {

		// query param hvn.uuid
		var qrHvnUUID string

		if o.HvnUUID != nil {
			qrHvnUUID = *o.HvnUUID
		}
		qHvnUUID := qrHvnUUID
		if qHvnUUID != "" {

			if err := r.SetQueryParam("hvn.uuid", qHvnUUID); err != nil {
				return err
			}
		}
	}

	if o.PaginationNextPageToken != nil {

		// query param pagination.next_page_token
		var qrPaginationNextPageToken string

		if o.PaginationNextPageToken != nil {
			qrPaginationNextPageToken = *o.PaginationNextPageToken
		}
		qPaginationNextPageToken := qrPaginationNextPageToken
		if qPaginationNextPageToken != "" {

			if err := r.SetQueryParam("pagination.next_page_token", qPaginationNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.PaginationPageSize != nil {

		// query param pagination.page_size
		var qrPaginationPageSize int64

		if o.PaginationPageSize != nil {
			qrPaginationPageSize = *o.PaginationPageSize
		}
		qPaginationPageSize := swag.FormatInt64(qrPaginationPageSize)
		if qPaginationPageSize != "" {

			if err := r.SetQueryParam("pagination.page_size", qPaginationPageSize); err != nil {
				return err
			}
		}
	}

	if o.PaginationPreviousPageToken != nil {

		// query param pagination.previous_page_token
		var qrPaginationPreviousPageToken string

		if o.PaginationPreviousPageToken != nil {
			qrPaginationPreviousPageToken = *o.PaginationPreviousPageToken
		}
		qPaginationPreviousPageToken := qrPaginationPreviousPageToken
		if qPaginationPreviousPageToken != "" {

			if err := r.SetQueryParam("pagination.previous_page_token", qPaginationPreviousPageToken); err != nil {
				return err
			}
		}
	}

	if o.TargetID != nil {

		// query param target_id
		var qrTargetID string

		if o.TargetID != nil {
			qrTargetID = *o.TargetID
		}
		qTargetID := qrTargetID
		if qTargetID != "" {

			if err := r.SetQueryParam("target_id", qTargetID); err != nil {
				return err
			}
		}
	}

	if o.TargetType != nil {

		// query param target_type
		var qrTargetType string

		if o.TargetType != nil {
			qrTargetType = *o.TargetType
		}
		qTargetType := qrTargetType
		if qTargetType != "" {

			if err := r.SetQueryParam("target_type", qTargetType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
