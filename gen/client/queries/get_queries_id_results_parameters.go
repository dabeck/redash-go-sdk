// Code generated by go-swagger; DO NOT EDIT.

package queries

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

// NewGetQueriesIDResultsParams creates a new GetQueriesIDResultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQueriesIDResultsParams() *GetQueriesIDResultsParams {
	return &GetQueriesIDResultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQueriesIDResultsParamsWithTimeout creates a new GetQueriesIDResultsParams object
// with the ability to set a timeout on a request.
func NewGetQueriesIDResultsParamsWithTimeout(timeout time.Duration) *GetQueriesIDResultsParams {
	return &GetQueriesIDResultsParams{
		timeout: timeout,
	}
}

// NewGetQueriesIDResultsParamsWithContext creates a new GetQueriesIDResultsParams object
// with the ability to set a context for a request.
func NewGetQueriesIDResultsParamsWithContext(ctx context.Context) *GetQueriesIDResultsParams {
	return &GetQueriesIDResultsParams{
		Context: ctx,
	}
}

// NewGetQueriesIDResultsParamsWithHTTPClient creates a new GetQueriesIDResultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQueriesIDResultsParamsWithHTTPClient(client *http.Client) *GetQueriesIDResultsParams {
	return &GetQueriesIDResultsParams{
		HTTPClient: client,
	}
}

/*
GetQueriesIDResultsParams contains all the parameters to send to the API endpoint

	for the get queries ID results operation.

	Typically these are written to a http.Request.
*/
type GetQueriesIDResultsParams struct {

	// ID.
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get queries ID results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueriesIDResultsParams) WithDefaults() *GetQueriesIDResultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get queries ID results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueriesIDResultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get queries ID results params
func (o *GetQueriesIDResultsParams) WithTimeout(timeout time.Duration) *GetQueriesIDResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get queries ID results params
func (o *GetQueriesIDResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get queries ID results params
func (o *GetQueriesIDResultsParams) WithContext(ctx context.Context) *GetQueriesIDResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get queries ID results params
func (o *GetQueriesIDResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get queries ID results params
func (o *GetQueriesIDResultsParams) WithHTTPClient(client *http.Client) *GetQueriesIDResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get queries ID results params
func (o *GetQueriesIDResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get queries ID results params
func (o *GetQueriesIDResultsParams) WithID(id int64) *GetQueriesIDResultsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get queries ID results params
func (o *GetQueriesIDResultsParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetQueriesIDResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
