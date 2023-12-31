// Code generated by mockery v2.32.4. DO NOT EDIT.

package gcp

import (
	cloudsql "github.com/nduyphuong/gorya/pkg/gcp/cloudsql"
	gce "github.com/nduyphuong/gorya/pkg/gcp/gce"

	mock "github.com/stretchr/testify/mock"
)

// MockInterface is an autogenerated mock type for the Interface type
type MockInterface struct {
	mock.Mock
}

type MockInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterface) EXPECT() *MockInterface_Expecter {
	return &MockInterface_Expecter{mock: &_m.Mock}
}

// CloudSQL provides a mock function with given fields:
func (_m *MockInterface) CloudSQL() cloudsql.Interface {
	ret := _m.Called()

	var r0 cloudsql.Interface
	if rf, ok := ret.Get(0).(func() cloudsql.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cloudsql.Interface)
		}
	}

	return r0
}

// MockInterface_CloudSQL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CloudSQL'
type MockInterface_CloudSQL_Call struct {
	*mock.Call
}

// CloudSQL is a helper method to define mock.On call
func (_e *MockInterface_Expecter) CloudSQL() *MockInterface_CloudSQL_Call {
	return &MockInterface_CloudSQL_Call{Call: _e.mock.On("CloudSQL")}
}

func (_c *MockInterface_CloudSQL_Call) Run(run func()) *MockInterface_CloudSQL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_CloudSQL_Call) Return(_a0 cloudsql.Interface) *MockInterface_CloudSQL_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_CloudSQL_Call) RunAndReturn(run func() cloudsql.Interface) *MockInterface_CloudSQL_Call {
	_c.Call.Return(run)
	return _c
}

// GCE provides a mock function with given fields:
func (_m *MockInterface) GCE() gce.Interface {
	ret := _m.Called()

	var r0 gce.Interface
	if rf, ok := ret.Get(0).(func() gce.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gce.Interface)
		}
	}

	return r0
}

// MockInterface_GCE_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GCE'
type MockInterface_GCE_Call struct {
	*mock.Call
}

// GCE is a helper method to define mock.On call
func (_e *MockInterface_Expecter) GCE() *MockInterface_GCE_Call {
	return &MockInterface_GCE_Call{Call: _e.mock.On("GCE")}
}

func (_c *MockInterface_GCE_Call) Run(run func()) *MockInterface_GCE_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterface_GCE_Call) Return(_a0 gce.Interface) *MockInterface_GCE_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockInterface_GCE_Call) RunAndReturn(run func() gce.Interface) *MockInterface_GCE_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInterface creates a new instance of MockInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterface {
	mock := &MockInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
