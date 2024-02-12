// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	dto "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: productInfo
func (_m *ProductRepository) CreateProduct(productInfo []interface{}) (int64, error) {
	ret := _m.Called(productInfo)

	if len(ret) == 0 {
		panic("no return value specified for CreateProduct")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func([]interface{}) (int64, error)); ok {
		return rf(productInfo)
	}
	if rf, ok := ret.Get(0).(func([]interface{}) int64); ok {
		r0 = rf(productInfo)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func([]interface{}) error); ok {
		r1 = rf(productInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProductDetail provides a mock function with given fields: productDetailInfo
func (_m *ProductRepository) CreateProductDetail(productDetailInfo [][]interface{}) error {
	ret := _m.Called(productDetailInfo)

	if len(ret) == 0 {
		panic("no return value specified for CreateProductDetail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([][]interface{}) error); ok {
		r0 = rf(productDetailInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProductById provides a mock function with given fields: productId
func (_m *ProductRepository) GetProductById(productId int) (dto.ResponseProduct, error) {
	ret := _m.Called(productId)

	if len(ret) == 0 {
		panic("no return value specified for GetProductById")
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

// GetProductListWithFilters provides a mock function with given fields: filters
func (_m *ProductRepository) GetProductListWithFilters(filters map[string]string) ([]dto.ResponseProduct, error) {
	ret := _m.Called(filters)

	if len(ret) == 0 {
		panic("no return value specified for GetProductListWithFilters")
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

// UpdateProduct provides a mock function with given fields: productId, name, description
func (_m *ProductRepository) UpdateProduct(productId int, name string, description string) error {
	ret := _m.Called(productId, name, description)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, string) error); ok {
		r0 = rf(productId, name, description)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProductDetail provides a mock function with given fields: productDetailId, quantity
func (_m *ProductRepository) UpdateProductDetail(productDetailId int, quantity int) error {
	ret := _m.Called(productDetailId, quantity)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProductDetail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(productDetailId, quantity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}