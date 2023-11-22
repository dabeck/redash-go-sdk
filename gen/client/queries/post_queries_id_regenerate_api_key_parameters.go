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

// NewPostQueriesIDRegenerateAPIKeyParams creates a new PostQueriesIDRegenerateAPIKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostQueriesIDRegenerateAPIKeyParams() *PostQueriesIDRegenerateAPIKeyParams {
	return &PostQueriesIDRegenerateAPIKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostQueriesIDRegenerateAPIKeyParamsWithTimeout creates a new PostQueriesIDRegenerateAPIKeyParams object
// with the ability to set a timeout on a request.
func NewPostQueriesIDRegenerateAPIKeyParamsWithTimeout(timeout time.Duration) *PostQueriesIDRegenerateAPIKeyParams {
	return &PostQueriesIDRegenerateAPIKeyParams{
		timeout: timeout,
	}
}

// NewPostQueriesIDRegenerateAPIKeyParamsWithContext creates a new PostQueriesIDRegenerateAPIKeyParams object
// with the ability to set a context for a request.
func NewPostQueriesIDRegenerateAPIKeyParamsWithContext(ctx context.Context) *PostQueriesIDRegenerateAPIKeyParams {
	return &PostQueriesIDRegenerateAPIKeyParams{
		Context: ctx,
	}
}

// NewPostQueriesIDRegenerateAPIKeyParamsWithHTTPClient creates a new PostQueriesIDRegenerateAPIKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostQueriesIDRegenerateAPIKeyParamsWithHTTPClient(client *http.Client) *PostQueriesIDRegenerateAPIKeyParams {
	return &PostQueriesIDRegenerateAPIKeyParams{
		HTTPClient: client,
	}
}

/*
PostQueriesIDRegenerateAPIKeyParams contains all the parameters to send to the API endpoint

	for the post queries ID regenerate API key operation.

	Typically these are written to a http.Request.
*/
type PostQueriesIDRegenerateAPIKeyParams struct {

	// ID.
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post queries ID regenerate API key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQueriesIDRegenerateAPIKeyParams) WithDefaults() *PostQueriesIDRegenerateAPIKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post queries ID regenerate API key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQueriesIDRegenerateAPIKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post queries ID regenerate API key params
func (o *PostQueriesIDRegenerateAPIKeyParams) WithTimeout(timeout time.Duration) *PostQueriesIDRegenerateAPIKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post queries ID regenerate API key params
func (o *PostQueriesIDRegenerateAPIKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post queries ID regenerate API key params
func (o *PostQueriesIDRegenerateAPIKeyParams) WithContext(ctx context.Context) *PostQueriesIDRegenerateAPIKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post queries ID regenerate API key params
func (o *PostQueriesIDRegenerateAPIKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post queries ID regenerate API key params
func (o *PostQueriesIDRegenerateAPIKeyParams) WithHTTPClient(client *http.Client) *PostQueriesIDRegenerateAPIKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post queries ID regenerate API key params
func (o *PostQueriesIDRegenerateAPIKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post queries ID regenerate API key params
func (o *PostQueriesIDRegenerateAPIKeyParams) WithID(id int64) *PostQueriesIDRegenerateAPIKeyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post queries ID regenerate API key params
func (o *PostQueriesIDRegenerateAPIKeyParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostQueriesIDRegenerateAPIKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
