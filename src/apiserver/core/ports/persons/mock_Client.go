// Code generated by mockery v2.33.3. DO NOT EDIT.

package persons

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockClient is an autogenerated mock type for the Client type
type MockClient struct {
	mock.Mock
}

type MockClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockClient) EXPECT() *MockClient_Expecter {
	return &MockClient_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *MockClient) Create(_a0 context.Context, _a1 Person) (int32, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int32
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, Person) (int32, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, Person) int32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(context.Context, Person) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockClient_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 Person
func (_e *MockClient_Expecter) Create(_a0 interface{}, _a1 interface{}) *MockClient_Create_Call {
	return &MockClient_Create_Call{Call: _e.mock.On("Create", _a0, _a1)}
}

func (_c *MockClient_Create_Call) Run(run func(_a0 context.Context, _a1 Person)) *MockClient_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(Person))
	})
	return _c
}

func (_c *MockClient_Create_Call) Return(_a0 int32, _a1 error) *MockClient_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_Create_Call) RunAndReturn(run func(context.Context, Person) (int32, error)) *MockClient_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *MockClient) Delete(_a0 context.Context, _a1 int32) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockClient_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockClient_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int32
func (_e *MockClient_Expecter) Delete(_a0 interface{}, _a1 interface{}) *MockClient_Delete_Call {
	return &MockClient_Delete_Call{Call: _e.mock.On("Delete", _a0, _a1)}
}

func (_c *MockClient_Delete_Call) Run(run func(_a0 context.Context, _a1 int32)) *MockClient_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockClient_Delete_Call) Return(_a0 error) *MockClient_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockClient_Delete_Call) RunAndReturn(run func(context.Context, int32) error) *MockClient_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: _a0, _a1
func (_m *MockClient) Read(_a0 context.Context, _a1 int32) (Person, error) {
	ret := _m.Called(_a0, _a1)

	var r0 Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (Person, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) Person); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(Person)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockClient_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int32
func (_e *MockClient_Expecter) Read(_a0 interface{}, _a1 interface{}) *MockClient_Read_Call {
	return &MockClient_Read_Call{Call: _e.mock.On("Read", _a0, _a1)}
}

func (_c *MockClient_Read_Call) Run(run func(_a0 context.Context, _a1 int32)) *MockClient_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockClient_Read_Call) Return(_a0 Person, _a1 error) *MockClient_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_Read_Call) RunAndReturn(run func(context.Context, int32) (Person, error)) *MockClient_Read_Call {
	_c.Call.Return(run)
	return _c
}

// ReadWithinRange provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockClient) ReadWithinRange(_a0 context.Context, _a1 int32, _a2 int32) ([]Person, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) ([]Person, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32, int32) []Person); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Person)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32, int32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_ReadWithinRange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadWithinRange'
type MockClient_ReadWithinRange_Call struct {
	*mock.Call
}

// ReadWithinRange is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 int32
//   - _a2 int32
func (_e *MockClient_Expecter) ReadWithinRange(_a0 interface{}, _a1 interface{}, _a2 interface{}) *MockClient_ReadWithinRange_Call {
	return &MockClient_ReadWithinRange_Call{Call: _e.mock.On("ReadWithinRange", _a0, _a1, _a2)}
}

func (_c *MockClient_ReadWithinRange_Call) Run(run func(_a0 context.Context, _a1 int32, _a2 int32)) *MockClient_ReadWithinRange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32), args[2].(int32))
	})
	return _c
}

func (_c *MockClient_ReadWithinRange_Call) Return(_a0 []Person, _a1 error) *MockClient_ReadWithinRange_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_ReadWithinRange_Call) RunAndReturn(run func(context.Context, int32, int32) ([]Person, error)) *MockClient_ReadWithinRange_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *MockClient) Update(_a0 context.Context, _a1 Person) (Person, error) {
	ret := _m.Called(_a0, _a1)

	var r0 Person
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, Person) (Person, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, Person) Person); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(Person)
	}

	if rf, ok := ret.Get(1).(func(context.Context, Person) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockClient_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockClient_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 Person
func (_e *MockClient_Expecter) Update(_a0 interface{}, _a1 interface{}) *MockClient_Update_Call {
	return &MockClient_Update_Call{Call: _e.mock.On("Update", _a0, _a1)}
}

func (_c *MockClient_Update_Call) Run(run func(_a0 context.Context, _a1 Person)) *MockClient_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(Person))
	})
	return _c
}

func (_c *MockClient_Update_Call) Return(_a0 Person, _a1 error) *MockClient_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockClient_Update_Call) RunAndReturn(run func(context.Context, Person) (Person, error)) *MockClient_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockClient creates a new instance of MockClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockClient {
	mock := &MockClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
