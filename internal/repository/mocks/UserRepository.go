// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CheckUserWithEmailAndPassword provides a mock function with given fields: email, password
func (_m *UserRepository) CheckUserWithEmailAndPassword(email string, password string) error {
	ret := _m.Called(email, password)

	if len(ret) == 0 {
		panic("no return value specified for CheckUserWithEmailAndPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: userInfo
func (_m *UserRepository) CreateUser(userInfo []interface{}) error {
	ret := _m.Called(userInfo)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]interface{}) error); ok {
		r0 = rf(userInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetIdRoleAndPassword provides a mock function with given fields: email
func (_m *UserRepository) GetIdRoleAndPassword(email string) (int, int, string, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetIdRoleAndPassword")
	}

	var r0 int
	var r1 int
	var r2 string
	var r3 error
	if rf, ok := ret.Get(0).(func(string) (int, int, string, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(string) string); ok {
		r2 = rf(email)
	} else {
		r2 = ret.Get(2).(string)
	}

	if rf, ok := ret.Get(3).(func(string) error); ok {
		r3 = rf(email)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// IsUserWithEmailPresent provides a mock function with given fields: email
func (_m *UserRepository) IsUserWithEmailPresent(email string) bool {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for IsUserWithEmailPresent")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}