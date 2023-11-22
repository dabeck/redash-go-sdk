// Code generated by go-swagger; DO NOT EDIT.

package data_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new data sources API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data sources API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteDataSourcesID(params *DeleteDataSourcesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDataSourcesIDNoContent, error)

	GetDataSourcesID(params *GetDataSourcesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataSourcesIDOK, error)

	PostDataSources(params *PostDataSourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDataSourcesOK, error)

	PostDataSourcesID(params *PostDataSourcesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDataSourcesIDOK, error)

	List(params *ListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteDataSourcesID delete data sources ID API
*/
func (a *Client) DeleteDataSourcesID(params *DeleteDataSourcesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDataSourcesIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDataSourcesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteDataSourcesID",
		Method:             "DELETE",
		PathPattern:        "/data_sources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDataSourcesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDataSourcesIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDataSourcesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDataSourcesID get data sources ID API
*/
func (a *Client) GetDataSourcesID(params *GetDataSourcesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDataSourcesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataSourcesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDataSourcesID",
		Method:             "GET",
		PathPattern:        "/data_sources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDataSourcesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDataSourcesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDataSourcesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostDataSources post data sources API
*/
func (a *Client) PostDataSources(params *PostDataSourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDataSourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDataSourcesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostDataSources",
		Method:             "POST",
		PathPattern:        "/data_sources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostDataSourcesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostDataSourcesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostDataSourcesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostDataSourcesID post data sources ID API
*/
func (a *Client) PostDataSourcesID(params *PostDataSourcesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostDataSourcesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostDataSourcesIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostDataSourcesID",
		Method:             "POST",
		PathPattern:        "/data_sources/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostDataSourcesIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostDataSourcesIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostDataSourcesIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
List list API
*/
func (a *Client) List(params *ListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "list",
		Method:             "GET",
		PathPattern:        "/data_sources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
