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

// AddProductToOrder provides a mock function with given fields: userId, orderId, productDetailId, product
func (_m *Service) AddProductToOrder(userId int, orderId int, productDetailId int, product dto.ProductCartRequest) error {
	ret := _m.Called(userId, orderId, productDetailId, product)

	if len(ret) == 0 {
		panic("no return value specified for AddProductToOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int, dto.ProductCartRequest) error); ok {
		r0 = rf(userId, orderId, productDetailId, product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateOrder provides a mock function with given fields: userId
func (_m *Service) CreateOrder(userId int) (int, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllOrderItems provides a mock function with given fields: userId, orderId
func (_m *Service) GetAllOrderItems(userId int, orderId int) (interface{}, error) {
	ret := _m.Called(userId, orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllOrderItems")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (interface{}, error)); ok {
		return rf(userId, orderId)
	}
	if rf, ok := ret.Get(0).(func(int, int) interface{}); ok {
		r0 = rf(userId, orderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userId, orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlaceOrderDetails provides a mock function with given fields: userId, orderId
func (_m *Service) GetPlaceOrderDetails(userId int, orderId int) (interface{}, error) {
	ret := _m.Called(userId, orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetPlaceOrderDetails")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (interface{}, error)); ok {
		return rf(userId, orderId)
	}
	if rf, ok := ret.Get(0).(func(int, int) interface{}); ok {
		r0 = rf(userId, orderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userId, orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserPlacedOrders provides a mock function with given fields: userId
func (_m *Service) GetUserPlacedOrders(userId int) ([]dto.UserOrderResponse, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for GetUserPlacedOrders")
	}

	var r0 []dto.UserOrderResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]dto.UserOrderResponse, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(int) []dto.UserOrderResponse); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.UserOrderResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlaceOrder provides a mock function with given fields: userId, orderId, shipping_address
func (_m *Service) PlaceOrder(userId int, orderId int, shipping_address string) error {
	ret := _m.Called(userId, orderId, shipping_address)

	if len(ret) == 0 {
		panic("no return value specified for PlaceOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, string) error); ok {
		r0 = rf(userId, orderId, shipping_address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProductInCart provides a mock function with given fields: userId, orderId, productDetailId, product
func (_m *Service) UpdateProductInCart(userId int, orderId int, productDetailId int, product dto.ProductCartRequest) error {
	ret := _m.Called(userId, orderId, productDetailId, product)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProductInCart")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int, dto.ProductCartRequest) error); ok {
		r0 = rf(userId, orderId, productDetailId, product)
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
