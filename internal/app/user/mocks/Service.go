// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	dto "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// LoginUser provides a mock function with given fields: email, passsword
func (_m *Service) LoginUser(email string, passsword string) (string, error) {
	ret := _m.Called(email, passsword)

	if len(ret) == 0 {
		panic("no return value specified for LoginUser")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(email, passsword)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, passsword)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, passsword)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: userInfo
func (_m *Service) RegisterUser(userInfo dto.RegisterUserRequest) error {
	ret := _m.Called(userInfo)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.RegisterUserRequest) error); ok {
		r0 = rf(userInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}