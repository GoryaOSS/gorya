// Code generated by mockery v2.32.4. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockGoryaServiceHandler is an autogenerated mock type for the GoryaServiceHandler type
type MockGoryaServiceHandler struct {
	mock.Mock
}

type MockGoryaServiceHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGoryaServiceHandler) EXPECT() *MockGoryaServiceHandler_Expecter {
	return &MockGoryaServiceHandler_Expecter{mock: &_m.Mock}
}

// AddPolicy provides a mock function with given fields: ctx
func (_m *MockGoryaServiceHandler) AddPolicy(ctx context.Context) http.Handler {
	ret := _m.Called(ctx)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(context.Context) http.Handler); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_AddPolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddPolicy'
type MockGoryaServiceHandler_AddPolicy_Call struct {
	*mock.Call
}

// AddPolicy is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoryaServiceHandler_Expecter) AddPolicy(ctx interface{}) *MockGoryaServiceHandler_AddPolicy_Call {
	return &MockGoryaServiceHandler_AddPolicy_Call{Call: _e.mock.On("AddPolicy", ctx)}
}

func (_c *MockGoryaServiceHandler_AddPolicy_Call) Run(run func(ctx context.Context)) *MockGoryaServiceHandler_AddPolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoryaServiceHandler_AddPolicy_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_AddPolicy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_AddPolicy_Call) RunAndReturn(run func(context.Context) http.Handler) *MockGoryaServiceHandler_AddPolicy_Call {
	_c.Call.Return(run)
	return _c
}

// AddSchedule provides a mock function with given fields: ctx
func (_m *MockGoryaServiceHandler) AddSchedule(ctx context.Context) http.Handler {
	ret := _m.Called(ctx)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(context.Context) http.Handler); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_AddSchedule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddSchedule'
type MockGoryaServiceHandler_AddSchedule_Call struct {
	*mock.Call
}

// AddSchedule is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoryaServiceHandler_Expecter) AddSchedule(ctx interface{}) *MockGoryaServiceHandler_AddSchedule_Call {
	return &MockGoryaServiceHandler_AddSchedule_Call{Call: _e.mock.On("AddSchedule", ctx)}
}

func (_c *MockGoryaServiceHandler_AddSchedule_Call) Run(run func(ctx context.Context)) *MockGoryaServiceHandler_AddSchedule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoryaServiceHandler_AddSchedule_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_AddSchedule_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_AddSchedule_Call) RunAndReturn(run func(context.Context) http.Handler) *MockGoryaServiceHandler_AddSchedule_Call {
	_c.Call.Return(run)
	return _c
}

// ChangeState provides a mock function with given fields: ctx
func (_m *MockGoryaServiceHandler) ChangeState(ctx context.Context) http.Handler {
	ret := _m.Called(ctx)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(context.Context) http.Handler); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_ChangeState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChangeState'
type MockGoryaServiceHandler_ChangeState_Call struct {
	*mock.Call
}

// ChangeState is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoryaServiceHandler_Expecter) ChangeState(ctx interface{}) *MockGoryaServiceHandler_ChangeState_Call {
	return &MockGoryaServiceHandler_ChangeState_Call{Call: _e.mock.On("ChangeState", ctx)}
}

func (_c *MockGoryaServiceHandler_ChangeState_Call) Run(run func(ctx context.Context)) *MockGoryaServiceHandler_ChangeState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoryaServiceHandler_ChangeState_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_ChangeState_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_ChangeState_Call) RunAndReturn(run func(context.Context) http.Handler) *MockGoryaServiceHandler_ChangeState_Call {
	_c.Call.Return(run)
	return _c
}

// DeletePolicy provides a mock function with given fields: ctx
func (_m *MockGoryaServiceHandler) DeletePolicy(ctx context.Context) http.Handler {
	ret := _m.Called(ctx)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(context.Context) http.Handler); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_DeletePolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeletePolicy'
type MockGoryaServiceHandler_DeletePolicy_Call struct {
	*mock.Call
}

// DeletePolicy is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoryaServiceHandler_Expecter) DeletePolicy(ctx interface{}) *MockGoryaServiceHandler_DeletePolicy_Call {
	return &MockGoryaServiceHandler_DeletePolicy_Call{Call: _e.mock.On("DeletePolicy", ctx)}
}

func (_c *MockGoryaServiceHandler_DeletePolicy_Call) Run(run func(ctx context.Context)) *MockGoryaServiceHandler_DeletePolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoryaServiceHandler_DeletePolicy_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_DeletePolicy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_DeletePolicy_Call) RunAndReturn(run func(context.Context) http.Handler) *MockGoryaServiceHandler_DeletePolicy_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteSchedule provides a mock function with given fields: ctx
func (_m *MockGoryaServiceHandler) DeleteSchedule(ctx context.Context) http.Handler {
	ret := _m.Called(ctx)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(context.Context) http.Handler); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_DeleteSchedule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteSchedule'
type MockGoryaServiceHandler_DeleteSchedule_Call struct {
	*mock.Call
}

// DeleteSchedule is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoryaServiceHandler_Expecter) DeleteSchedule(ctx interface{}) *MockGoryaServiceHandler_DeleteSchedule_Call {
	return &MockGoryaServiceHandler_DeleteSchedule_Call{Call: _e.mock.On("DeleteSchedule", ctx)}
}

func (_c *MockGoryaServiceHandler_DeleteSchedule_Call) Run(run func(ctx context.Context)) *MockGoryaServiceHandler_DeleteSchedule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoryaServiceHandler_DeleteSchedule_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_DeleteSchedule_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_DeleteSchedule_Call) RunAndReturn(run func(context.Context) http.Handler) *MockGoryaServiceHandler_DeleteSchedule_Call {
	_c.Call.Return(run)
	return _c
}

// GetPolicy provides a mock function with given fields: ctx
func (_m *MockGoryaServiceHandler) GetPolicy(ctx context.Context) http.Handler {
	ret := _m.Called(ctx)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(context.Context) http.Handler); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_GetPolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPolicy'
type MockGoryaServiceHandler_GetPolicy_Call struct {
	*mock.Call
}

// GetPolicy is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoryaServiceHandler_Expecter) GetPolicy(ctx interface{}) *MockGoryaServiceHandler_GetPolicy_Call {
	return &MockGoryaServiceHandler_GetPolicy_Call{Call: _e.mock.On("GetPolicy", ctx)}
}

func (_c *MockGoryaServiceHandler_GetPolicy_Call) Run(run func(ctx context.Context)) *MockGoryaServiceHandler_GetPolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoryaServiceHandler_GetPolicy_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_GetPolicy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_GetPolicy_Call) RunAndReturn(run func(context.Context) http.Handler) *MockGoryaServiceHandler_GetPolicy_Call {
	_c.Call.Return(run)
	return _c
}

// GetSchedule provides a mock function with given fields: ctx
func (_m *MockGoryaServiceHandler) GetSchedule(ctx context.Context) http.Handler {
	ret := _m.Called(ctx)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(context.Context) http.Handler); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_GetSchedule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSchedule'
type MockGoryaServiceHandler_GetSchedule_Call struct {
	*mock.Call
}

// GetSchedule is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoryaServiceHandler_Expecter) GetSchedule(ctx interface{}) *MockGoryaServiceHandler_GetSchedule_Call {
	return &MockGoryaServiceHandler_GetSchedule_Call{Call: _e.mock.On("GetSchedule", ctx)}
}

func (_c *MockGoryaServiceHandler_GetSchedule_Call) Run(run func(ctx context.Context)) *MockGoryaServiceHandler_GetSchedule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoryaServiceHandler_GetSchedule_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_GetSchedule_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_GetSchedule_Call) RunAndReturn(run func(context.Context) http.Handler) *MockGoryaServiceHandler_GetSchedule_Call {
	_c.Call.Return(run)
	return _c
}

// GetTimeZone provides a mock function with given fields:
func (_m *MockGoryaServiceHandler) GetTimeZone() http.Handler {
	ret := _m.Called()

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_GetTimeZone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTimeZone'
type MockGoryaServiceHandler_GetTimeZone_Call struct {
	*mock.Call
}

// GetTimeZone is a helper method to define mock.On call
func (_e *MockGoryaServiceHandler_Expecter) GetTimeZone() *MockGoryaServiceHandler_GetTimeZone_Call {
	return &MockGoryaServiceHandler_GetTimeZone_Call{Call: _e.mock.On("GetTimeZone")}
}

func (_c *MockGoryaServiceHandler_GetTimeZone_Call) Run(run func()) *MockGoryaServiceHandler_GetTimeZone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockGoryaServiceHandler_GetTimeZone_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_GetTimeZone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_GetTimeZone_Call) RunAndReturn(run func() http.Handler) *MockGoryaServiceHandler_GetTimeZone_Call {
	_c.Call.Return(run)
	return _c
}

// GetVersionInfo provides a mock function with given fields:
func (_m *MockGoryaServiceHandler) GetVersionInfo() http.Handler {
	ret := _m.Called()

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_GetVersionInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVersionInfo'
type MockGoryaServiceHandler_GetVersionInfo_Call struct {
	*mock.Call
}

// GetVersionInfo is a helper method to define mock.On call
func (_e *MockGoryaServiceHandler_Expecter) GetVersionInfo() *MockGoryaServiceHandler_GetVersionInfo_Call {
	return &MockGoryaServiceHandler_GetVersionInfo_Call{Call: _e.mock.On("GetVersionInfo")}
}

func (_c *MockGoryaServiceHandler_GetVersionInfo_Call) Run(run func()) *MockGoryaServiceHandler_GetVersionInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockGoryaServiceHandler_GetVersionInfo_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_GetVersionInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_GetVersionInfo_Call) RunAndReturn(run func() http.Handler) *MockGoryaServiceHandler_GetVersionInfo_Call {
	_c.Call.Return(run)
	return _c
}

// ListPolicy provides a mock function with given fields: ctx
func (_m *MockGoryaServiceHandler) ListPolicy(ctx context.Context) http.Handler {
	ret := _m.Called(ctx)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(context.Context) http.Handler); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_ListPolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPolicy'
type MockGoryaServiceHandler_ListPolicy_Call struct {
	*mock.Call
}

// ListPolicy is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoryaServiceHandler_Expecter) ListPolicy(ctx interface{}) *MockGoryaServiceHandler_ListPolicy_Call {
	return &MockGoryaServiceHandler_ListPolicy_Call{Call: _e.mock.On("ListPolicy", ctx)}
}

func (_c *MockGoryaServiceHandler_ListPolicy_Call) Run(run func(ctx context.Context)) *MockGoryaServiceHandler_ListPolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoryaServiceHandler_ListPolicy_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_ListPolicy_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_ListPolicy_Call) RunAndReturn(run func(context.Context) http.Handler) *MockGoryaServiceHandler_ListPolicy_Call {
	_c.Call.Return(run)
	return _c
}

// ListSchedule provides a mock function with given fields: ctx
func (_m *MockGoryaServiceHandler) ListSchedule(ctx context.Context) http.Handler {
	ret := _m.Called(ctx)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(context.Context) http.Handler); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_ListSchedule_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSchedule'
type MockGoryaServiceHandler_ListSchedule_Call struct {
	*mock.Call
}

// ListSchedule is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoryaServiceHandler_Expecter) ListSchedule(ctx interface{}) *MockGoryaServiceHandler_ListSchedule_Call {
	return &MockGoryaServiceHandler_ListSchedule_Call{Call: _e.mock.On("ListSchedule", ctx)}
}

func (_c *MockGoryaServiceHandler_ListSchedule_Call) Run(run func(ctx context.Context)) *MockGoryaServiceHandler_ListSchedule_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoryaServiceHandler_ListSchedule_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_ListSchedule_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_ListSchedule_Call) RunAndReturn(run func(context.Context) http.Handler) *MockGoryaServiceHandler_ListSchedule_Call {
	_c.Call.Return(run)
	return _c
}

// ScheduleTask provides a mock function with given fields: ctx
func (_m *MockGoryaServiceHandler) ScheduleTask(ctx context.Context) http.Handler {
	ret := _m.Called(ctx)

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(context.Context) http.Handler); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// MockGoryaServiceHandler_ScheduleTask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ScheduleTask'
type MockGoryaServiceHandler_ScheduleTask_Call struct {
	*mock.Call
}

// ScheduleTask is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockGoryaServiceHandler_Expecter) ScheduleTask(ctx interface{}) *MockGoryaServiceHandler_ScheduleTask_Call {
	return &MockGoryaServiceHandler_ScheduleTask_Call{Call: _e.mock.On("ScheduleTask", ctx)}
}

func (_c *MockGoryaServiceHandler_ScheduleTask_Call) Run(run func(ctx context.Context)) *MockGoryaServiceHandler_ScheduleTask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockGoryaServiceHandler_ScheduleTask_Call) Return(_a0 http.Handler) *MockGoryaServiceHandler_ScheduleTask_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGoryaServiceHandler_ScheduleTask_Call) RunAndReturn(run func(context.Context) http.Handler) *MockGoryaServiceHandler_ScheduleTask_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGoryaServiceHandler creates a new instance of MockGoryaServiceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGoryaServiceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGoryaServiceHandler {
	mock := &MockGoryaServiceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}