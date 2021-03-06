// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceInterface is an autogenerated mock type for the ResourceInterface type
type ResourceInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, obj, options, subresources
func (_m *ResourceInterface) Create(ctx context.Context, obj *unstructured.Unstructured, options v1.CreateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, obj, options)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(context.Context, *unstructured.Unstructured, v1.CreateOptions, ...string) *unstructured.Unstructured); ok {
		r0 = rf(ctx, obj, options, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *unstructured.Unstructured, v1.CreateOptions, ...string) error); ok {
		r1 = rf(ctx, obj, options, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, name, opts, subresources
func (_m *ResourceInterface) Delete(ctx context.Context, name string, opts v1.DeleteOptions, subresources ...string) error {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.DeleteOptions, ...string) error); ok {
		r0 = rf(ctx, name, opts, subresources...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, name, opts, subresources
func (_m *ResourceInterface) Get(ctx context.Context, name string, opts v1.GetOptions, subresources ...string) (*unstructured.Unstructured, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, opts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.GetOptions, ...string) *unstructured.Unstructured); ok {
		r0 = rf(ctx, name, opts, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, v1.GetOptions, ...string) error); ok {
		r1 = rf(ctx, name, opts, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, obj, options, subresources
func (_m *ResourceInterface) Update(ctx context.Context, obj *unstructured.Unstructured, options v1.UpdateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, obj, options)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *unstructured.Unstructured
	if rf, ok := ret.Get(0).(func(context.Context, *unstructured.Unstructured, v1.UpdateOptions, ...string) *unstructured.Unstructured); ok {
		r0 = rf(ctx, obj, options, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*unstructured.Unstructured)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *unstructured.Unstructured, v1.UpdateOptions, ...string) error); ok {
		r1 = rf(ctx, obj, options, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
