// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	foo "github.com/vektra/mockery/v2/pkg/fixtures/example_project/bar/foo"
)

// Collision is an autogenerated mock type for the Collision type
type Collision struct {
	mock.Mock
}

// NewClient provides a mock function with given fields:
func (_m *Collision) NewClient() foo.Client {
	ret := _m.Called()

	var r0 foo.Client
	if rf, ok := ret.Get(0).(func() foo.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(foo.Client)
		}
	}

	return r0
}
