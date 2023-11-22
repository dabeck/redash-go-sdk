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
)

// NewPostQueriesParams creates a new PostQueriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostQueriesParams() *PostQueriesParams {
	return &PostQueriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostQueriesParamsWithTimeout creates a new PostQueriesParams object
// with the ability to set a timeout on a request.
func NewPostQueriesParamsWithTimeout(timeout time.Duration) *PostQueriesParams {
	return &PostQueriesParams{
		timeout: timeout,
	}
}

// NewPostQueriesParamsWithContext creates a new PostQueriesParams object
// with the ability to set a context for a request.
func NewPostQueriesParamsWithContext(ctx context.Context) *PostQueriesParams {
	return &PostQueriesParams{
		Context: ctx,
	}
}

// NewPostQueriesParamsWithHTTPClient creates a new PostQueriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostQueriesParamsWithHTTPClient(client *http.Client) *PostQueriesParams {
	return &PostQueriesParams{
		HTTPClient: client,
	}
}

/*
PostQueriesParams contains all the parameters to send to the API endpoint

	for the post queries operation.

	Typically these are written to a http.Request.
*/
type PostQueriesParams struct {

	// Body.
	Body PostQueriesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQueriesParams) WithDefaults() *PostQueriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post queries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQueriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post queries params
func (o *PostQueriesParams) WithTimeout(timeout time.Duration) *PostQueriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post queries params
func (o *PostQueriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post queries params
func (o *PostQueriesParams) WithContext(ctx context.Context) *PostQueriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post queries params
func (o *PostQueriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post queries params
func (o *PostQueriesParams) WithHTTPClient(client *http.Client) *PostQueriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post queries params
func (o *PostQueriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post queries params
func (o *PostQueriesParams) WithBody(body PostQueriesBody) *PostQueriesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post queries params
func (o *PostQueriesParams) SetBody(body PostQueriesBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostQueriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
