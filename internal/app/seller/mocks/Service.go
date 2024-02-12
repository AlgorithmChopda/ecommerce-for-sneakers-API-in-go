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

// DeleteSeller provides a mock function with given fields: sellerId
func (_m *Service) DeleteSeller(sellerId int) error {
	ret := _m.Called(sellerId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSeller")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(sellerId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSellerList provides a mock function with given fields:
func (_m *Service) GetSellerList() ([]dto.SellerResponseObject, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSellerList")
	}

	var r0 []dto.SellerResponseObject
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]dto.SellerResponseObject, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []dto.SellerResponseObject); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.SellerResponseObject)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterSeller provides a mock function with given fields: sellerInfo
func (_m *Service) RegisterSeller(sellerInfo dto.RegisterSellerRequest) error {
	ret := _m.Called(sellerInfo)

	if len(ret) == 0 {
		panic("no return value specified for RegisterSeller")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.RegisterSellerRequest) error); ok {
		r0 = rf(sellerInfo)
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
