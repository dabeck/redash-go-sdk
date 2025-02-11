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

// NewPostQueriesIDParams creates a new PostQueriesIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostQueriesIDParams() *PostQueriesIDParams {
	return &PostQueriesIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostQueriesIDParamsWithTimeout creates a new PostQueriesIDParams object
// with the ability to set a timeout on a request.
func NewPostQueriesIDParamsWithTimeout(timeout time.Duration) *PostQueriesIDParams {
	return &PostQueriesIDParams{
		timeout: timeout,
	}
}

// NewPostQueriesIDParamsWithContext creates a new PostQueriesIDParams object
// with the ability to set a context for a request.
func NewPostQueriesIDParamsWithContext(ctx context.Context) *PostQueriesIDParams {
	return &PostQueriesIDParams{
		Context: ctx,
	}
}

// NewPostQueriesIDParamsWithHTTPClient creates a new PostQueriesIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostQueriesIDParamsWithHTTPClient(client *http.Client) *PostQueriesIDParams {
	return &PostQueriesIDParams{
		HTTPClient: client,
	}
}

/*
PostQueriesIDParams contains all the parameters to send to the API endpoint

	for the post queries ID operation.

	Typically these are written to a http.Request.
*/
type PostQueriesIDParams struct {

	// Body.
	Body PostQueriesIDBody

	// ID.
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post queries ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQueriesIDParams) WithDefaults() *PostQueriesIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post queries ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostQueriesIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post queries ID params
func (o *PostQueriesIDParams) WithTimeout(timeout time.Duration) *PostQueriesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post queries ID params
func (o *PostQueriesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post queries ID params
func (o *PostQueriesIDParams) WithContext(ctx context.Context) *PostQueriesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post queries ID params
func (o *PostQueriesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post queries ID params
func (o *PostQueriesIDParams) WithHTTPClient(client *http.Client) *PostQueriesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post queries ID params
func (o *PostQueriesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post queries ID params
func (o *PostQueriesIDParams) WithBody(body PostQueriesIDBody) *PostQueriesIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post queries ID params
func (o *PostQueriesIDParams) SetBody(body PostQueriesIDBody) {
	o.Body = body
}

// WithID adds the id to the post queries ID params
func (o *PostQueriesIDParams) WithID(id int64) *PostQueriesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post queries ID params
func (o *PostQueriesIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostQueriesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
