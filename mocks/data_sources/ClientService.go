// Code generated by mockery v2.14.0. DO NOT EDIT.

package data_sourcesmock

import (
	data_sources "github.com/dabeck/redash-go-sdk/gen/client/data_sources"
	mock "github.com/stretchr/testify/mock"

	runtime "github.com/go-openapi/runtime"
)

// ClientService is an autogenerated mock type for the ClientService type
type ClientService struct {
	mock.Mock
}

// DeleteDataSourcesID provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) DeleteDataSourcesID(params *data_sources.DeleteDataSourcesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...data_sources.ClientOption) (*data_sources.DeleteDataSourcesIDNoContent, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *data_sources.DeleteDataSourcesIDNoContent
	if rf, ok := ret.Get(0).(func(*data_sources.DeleteDataSourcesIDParams, runtime.ClientAuthInfoWriter, ...data_sources.ClientOption) *data_sources.DeleteDataSourcesIDNoContent); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data_sources.DeleteDataSourcesIDNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*data_sources.DeleteDataSourcesIDParams, runtime.ClientAuthInfoWriter, ...data_sources.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataSourcesID provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) GetDataSourcesID(params *data_sources.GetDataSourcesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...data_sources.ClientOption) (*data_sources.GetDataSourcesIDOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *data_sources.GetDataSourcesIDOK
	if rf, ok := ret.Get(0).(func(*data_sources.GetDataSourcesIDParams, runtime.ClientAuthInfoWriter, ...data_sources.ClientOption) *data_sources.GetDataSourcesIDOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data_sources.GetDataSourcesIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*data_sources.GetDataSourcesIDParams, runtime.ClientAuthInfoWriter, ...data_sources.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) List(params *data_sources.ListParams, authInfo runtime.ClientAuthInfoWriter, opts ...data_sources.ClientOption) (*data_sources.ListOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *data_sources.ListOK
	if rf, ok := ret.Get(0).(func(*data_sources.ListParams, runtime.ClientAuthInfoWriter, ...data_sources.ClientOption) *data_sources.ListOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data_sources.ListOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*data_sources.ListParams, runtime.ClientAuthInfoWriter, ...data_sources.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostDataSources provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostDataSources(params *data_sources.PostDataSourcesParams, authInfo runtime.ClientAuthInfoWriter, opts ...data_sources.ClientOption) (*data_sources.PostDataSourcesOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *data_sources.PostDataSourcesOK
	if rf, ok := ret.Get(0).(func(*data_sources.PostDataSourcesParams, runtime.ClientAuthInfoWriter, ...data_sources.ClientOption) *data_sources.PostDataSourcesOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data_sources.PostDataSourcesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*data_sources.PostDataSourcesParams, runtime.ClientAuthInfoWriter, ...data_sources.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PostDataSourcesID provides a mock function with given fields: params, authInfo, opts
func (_m *ClientService) PostDataSourcesID(params *data_sources.PostDataSourcesIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...data_sources.ClientOption) (*data_sources.PostDataSourcesIDOK, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, params, authInfo)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *data_sources.PostDataSourcesIDOK
	if rf, ok := ret.Get(0).(func(*data_sources.PostDataSourcesIDParams, runtime.ClientAuthInfoWriter, ...data_sources.ClientOption) *data_sources.PostDataSourcesIDOK); ok {
		r0 = rf(params, authInfo, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*data_sources.PostDataSourcesIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*data_sources.PostDataSourcesIDParams, runtime.ClientAuthInfoWriter, ...data_sources.ClientOption) error); ok {
		r1 = rf(params, authInfo, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *ClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

type mockConstructorTestingTNewClientService interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientService creates a new instance of ClientService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientService(t mockConstructorTestingTNewClientService) *ClientService {
	mock := &ClientService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
