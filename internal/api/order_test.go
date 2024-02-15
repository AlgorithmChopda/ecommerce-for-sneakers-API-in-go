package api

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/order/mocks"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrderHandler(t *testing.T) {
	orderSvc := mocks.NewService(t)
	createWebsiteHandler := CreateOrderHandler(orderSvc)

	tests := []struct {
		name               string
		input              string
		isToken            bool
		token              dto.JwtToken
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:    "succcess creating cart",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: "",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateOrder", mock.Anything).Return(1, nil).Once()
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:    "error creating cart",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: "",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateOrder", mock.Anything).Return(1, errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orderSvc)

			req := httptest.NewRequest("POST", "/cart", bytes.NewBuffer([]byte(test.input)))
			if test.isToken {
				req = req.WithContext(context.WithValue(req.Context(), "token", test.token))
			}

			rr := httptest.NewRecorder()
			createWebsiteHandler(rr, req)
			// handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestAddOrderHandler(t *testing.T) {
	orderSvc := mocks.NewService(t)
	addOrderHandler := AddOrderHandler(orderSvc)

	tests := []struct {
		name               string
		input              string
		cartId             string
		productDetailId    string
		isToken            bool
		token              dto.JwtToken
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:    "invalid id param",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input:              "",
			cartId:             "",
			productDetailId:    "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "invalid product detail param",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input:              "",
			cartId:             "3",
			productDetailId:    "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "invalid json",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input:              "{",
			cartId:             "3",
			productDetailId:    "4",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "invalid json",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `{
				"quantity": -1
			}`,
			cartId:             "3",
			productDetailId:    "4",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "error adding product to cart",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `{
				"quantity": 1
			}`,
			cartId:          "3",
			productDetailId: "4",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("AddProductToOrder", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:    "success adding product to cart",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `{
				"quantity": 1
			}`,
			cartId:          "3",
			productDetailId: "4",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("AddProductToOrder", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orderSvc)

			req := httptest.NewRequest("POST", "/cart/{cartId}/product/{productDetailId}", bytes.NewBuffer([]byte(test.input)))

			req = mux.SetURLVars(req, map[string]string{
				"id":              test.cartId,
				"productDetailId": test.productDetailId,
			})

			if test.isToken {
				req = req.WithContext(context.WithValue(req.Context(), "token", test.token))
			}

			rr := httptest.NewRecorder()
			addOrderHandler(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestUpdateOrderHandler(t *testing.T) {
	orderSvc := mocks.NewService(t)
	addOrderHandler := UpdateOrderItemHandler(orderSvc)

	tests := []struct {
		name               string
		input              string
		cartId             string
		productDetailId    string
		isToken            bool
		token              dto.JwtToken
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:    "invalid id param",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input:              "",
			cartId:             "",
			productDetailId:    "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "invalid product detail param",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input:              "",
			cartId:             "3",
			productDetailId:    "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "invalid json",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input:              "{",
			cartId:             "3",
			productDetailId:    "4",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "invalid quantity",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `{
				"quantity": -1
			}`,
			cartId:             "3",
			productDetailId:    "4",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "success adding product to cart",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `{
				"quantity": 1
			}`,
			cartId:          "3",
			productDetailId: "4",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("UpdateProductInCart", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:    "success adding product to cart",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `{
				"quantity": 1
			}`,
			cartId:          "3",
			productDetailId: "4",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("UpdateProductInCart", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orderSvc)

			req := httptest.NewRequest("PUT", "/cart/{cartId}/product/{productDetailId}", bytes.NewBuffer([]byte(test.input)))

			req = mux.SetURLVars(req, map[string]string{
				"id":              test.cartId,
				"productDetailId": test.productDetailId,
			})

			if test.isToken {
				req = req.WithContext(context.WithValue(req.Context(), "token", test.token))
			}

			rr := httptest.NewRecorder()
			addOrderHandler(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestPlaceOrderHandler(t *testing.T) {
	orderSvc := mocks.NewService(t)
	addOrderHandler := PlaceOrderHandler(orderSvc)

	tests := []struct {
		name               string
		input              string
		cartId             string
		isToken            bool
		token              dto.JwtToken
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:    "invalid id param",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input:              "",
			cartId:             "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "invalid json",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input:              "{",
			cartId:             "2",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "invalid shipping address",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `{
				"shipping_address": ""
			}`,
			cartId:             "2",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "success placing order",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `{
				"shipping_address": "temp"
			}`,
			cartId: "3",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("PlaceOrder", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:    "error placing order",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `{
				"shipping_address": "temp"
			}`,
			cartId: "3",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("PlaceOrder", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orderSvc)

			req := httptest.NewRequest("POST", "/cart/{cartId}/order", bytes.NewBuffer([]byte(test.input)))

			req = mux.SetURLVars(req, map[string]string{
				"id": test.cartId,
			})

			if test.isToken {
				req = req.WithContext(context.WithValue(req.Context(), "token", test.token))
			}

			rr := httptest.NewRecorder()
			addOrderHandler(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
