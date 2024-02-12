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

// CreateProduct provides a mock function with given fields: _a0
func (_m *Service) CreateProduct(_a0 dto.CreateProductRequest) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.CreateProductRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProductByID provides a mock function with given fields: productId
func (_m *Service) GetProductByID(productId int) (dto.ResponseProduct, error) {
	ret := _m.Called(productId)

	if len(ret) == 0 {
		panic("no return value specified for GetProductByID")
	}

	var r0 dto.ResponseProduct
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (dto.ResponseProduct, error)); ok {
		return rf(productId)
	}
	if rf, ok := ret.Get(0).(func(int) dto.ResponseProduct); ok {
		r0 = rf(productId)
	} else {
		r0 = ret.Get(0).(dto.ResponseProduct)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductsByFilters provides a mock function with given fields: filters
func (_m *Service) GetProductsByFilters(filters map[string]string) ([]dto.ResponseProduct, error) {
	ret := _m.Called(filters)

	if len(ret) == 0 {
		panic("no return value specified for GetProductsByFilters")
	}

	var r0 []dto.ResponseProduct
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]string) ([]dto.ResponseProduct, error)); ok {
		return rf(filters)
	}
	if rf, ok := ret.Get(0).(func(map[string]string) []dto.ResponseProduct); ok {
		r0 = rf(filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.ResponseProduct)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]string) error); ok {
		r1 = rf(filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProduct provides a mock function with given fields: req, productId
func (_m *Service) UpdateProduct(req dto.UpdateProductRequest, productId int) error {
	ret := _m.Called(req, productId)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(dto.UpdateProductRequest, int) error); ok {
		r0 = rf(req, productId)
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
