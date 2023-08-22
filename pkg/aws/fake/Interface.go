// Code generated by mockery v2.32.4. DO NOT EDIT.

package fake

import (
	ec2 "github.com/nduyphuong/gorya/pkg/aws/ec2"
	mock "github.com/stretchr/testify/mock"
)

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// EC2 provides a mock function with given fields:
func (_m *Interface) EC2() ec2.Interface {
	ret := _m.Called()

	var r0 ec2.Interface
	if rf, ok := ret.Get(0).(func() ec2.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ec2.Interface)
		}
	}

	return r0
}

// NewInterface creates a new instance of Interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Interface {
	mock := &Interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
