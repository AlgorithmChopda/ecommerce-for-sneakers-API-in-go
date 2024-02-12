// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	dto "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// SellerRepository is an autogenerated mock type for the SellerRepository type
type SellerRepository struct {
	mock.Mock
}

// CreateCompany provides a mock function with given fields: sellerCompanyInfo
func (_m *SellerRepository) CreateCompany(sellerCompanyInfo []interface{}) (int64, error) {
	ret := _m.Called(sellerCompanyInfo)

	if len(ret) == 0 {
		panic("no return value specified for CreateCompany")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func([]interface{}) (int64, error)); ok {
		return rf(sellerCompanyInfo)
	}
	if rf, ok := ret.Get(0).(func([]interface{}) int64); ok {
		r0 = rf(sellerCompanyInfo)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func([]interface{}) error); ok {
		r1 = rf(sellerCompanyInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSeller provides a mock function with given fields: sellerInfo
func (_m *SellerRepository) CreateSeller(sellerInfo []interface{}) error {
	ret := _m.Called(sellerInfo)

	if len(ret) == 0 {
		panic("no return value specified for CreateSeller")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]interface{}) error); ok {
		r0 = rf(sellerInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSeller provides a mock function with given fields: sellerId
func (_m *SellerRepository) DeleteSeller(sellerId int) error {
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

// GetAllSellers provides a mock function with given fields: roleId
func (_m *SellerRepository) GetAllSellers(roleId int) ([]dto.SellerResponseObject, error) {
	ret := _m.Called(roleId)

	if len(ret) == 0 {
		panic("no return value specified for GetAllSellers")
	}

	var r0 []dto.SellerResponseObject
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]dto.SellerResponseObject, error)); ok {
		return rf(roleId)
	}
	if rf, ok := ret.Get(0).(func(int) []dto.SellerResponseObject); ok {
		r0 = rf(roleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.SellerResponseObject)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(roleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSellerRepository creates a new instance of SellerRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSellerRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *SellerRepository {
	mock := &SellerRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
