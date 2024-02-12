// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	dto "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// OrderRepository is an autogenerated mock type for the OrderRepository type
type OrderRepository struct {
	mock.Mock
}

// AddProductToOrder provides a mock function with given fields: userId, cartId, productDetailId, requiredQuantity
func (_m *OrderRepository) AddProductToOrder(userId int, cartId int, productDetailId int, requiredQuantity int) error {
	ret := _m.Called(userId, cartId, productDetailId, requiredQuantity)

	if len(ret) == 0 {
		panic("no return value specified for AddProductToOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int, int) error); ok {
		r0 = rf(userId, cartId, productDetailId, requiredQuantity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckOrderValid provides a mock function with given fields: userId, orderId
func (_m *OrderRepository) CheckOrderValid(userId int, orderId int) (bool, error) {
	ret := _m.Called(userId, orderId)

	if len(ret) == 0 {
		panic("no return value specified for CheckOrderValid")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (bool, error)); ok {
		return rf(userId, orderId)
	}
	if rf, ok := ret.Get(0).(func(int, int) bool); ok {
		r0 = rf(userId, orderId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(userId, orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: userId
func (_m *OrderRepository) Create(userId int) (int, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for Create")
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

// GetAllOrderItems provides a mock function with given fields: orderId
func (_m *OrderRepository) GetAllOrderItems(orderId int) (interface{}, error) {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllOrderItems")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (interface{}, error)); ok {
		return rf(orderId)
	}
	if rf, ok := ret.Get(0).(func(int) interface{}); ok {
		r0 = rf(orderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBuyerId provides a mock function with given fields: orderId
func (_m *OrderRepository) GetBuyerId(orderId int) (int, error) {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetBuyerId")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int, error)); ok {
		return rf(orderId)
	}
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(orderId)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderItemCount provides a mock function with given fields: orderId
func (_m *OrderRepository) GetOrderItemCount(orderId int) (int, error) {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderItemCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int, error)); ok {
		return rf(orderId)
	}
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(orderId)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlacedOrderDetails provides a mock function with given fields: userId, orderId
func (_m *OrderRepository) GetPlacedOrderDetails(userId int, orderId int) (interface{}, error) {
	ret := _m.Called(userId, orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetPlacedOrderDetails")
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

// GetUpdateItemsList provides a mock function with given fields: orderId
func (_m *OrderRepository) GetUpdateItemsList(orderId int) ([]int, []int, error) {
	ret := _m.Called(orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetUpdateItemsList")
	}

	var r0 []int
	var r1 []int
	var r2 error
	if rf, ok := ret.Get(0).(func(int) ([]int, []int, error)); ok {
		return rf(orderId)
	}
	if rf, ok := ret.Get(0).(func(int) []int); ok {
		r0 = rf(orderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	if rf, ok := ret.Get(1).(func(int) []int); ok {
		r1 = rf(orderId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]int)
		}
	}

	if rf, ok := ret.Get(2).(func(int) error); ok {
		r2 = rf(orderId)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetUserPlacedOrders provides a mock function with given fields: userId
func (_m *OrderRepository) GetUserPlacedOrders(userId int) ([]dto.UserOrderResponse, error) {
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

// IsOrderPresent provides a mock function with given fields: userId
func (_m *OrderRepository) IsOrderPresent(userId int) (bool, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for IsOrderPresent")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (bool, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PlaceOrder provides a mock function with given fields: userId, orderId, shippingAddress
func (_m *OrderRepository) PlaceOrder(userId int, orderId int, shippingAddress string) error {
	ret := _m.Called(userId, orderId, shippingAddress)

	if len(ret) == 0 {
		panic("no return value specified for PlaceOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, string) error); ok {
		r0 = rf(userId, orderId, shippingAddress)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrderItem provides a mock function with given fields: userId, cartId, productDetailId, requiredQuantity
func (_m *OrderRepository) UpdateOrderItem(userId int, cartId int, productDetailId int, requiredQuantity int) error {
	ret := _m.Called(userId, cartId, productDetailId, requiredQuantity)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrderItem")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int, int) error); ok {
		r0 = rf(userId, cartId, productDetailId, requiredQuantity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrderRepository creates a new instance of OrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderRepository {
	mock := &OrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}