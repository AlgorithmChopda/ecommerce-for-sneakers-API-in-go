package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/product/mocks"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := CreateProductHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		isToken            bool
		token              interface{}
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:    "invalid json",
			input:   `[]`,
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "invalid request format",
			input:   `{}`,
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "Error adding Product",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `
			{
				"name": "jordan",
				"description": "exclusive shoes",
				"brand": "Puma",
				"variety": [
				  {
					"color": "red",
					"image": "link",
					"detail": [
					  {
						"size": 7,
						"quantity": 50,
						"price": 500
					  },
					  {
						"size": 8,
						"quantity": 100,
						"price": 500
					  }
					]
				  }
				 ]
			  }
		    `,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateProduct", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:    "Success adding Product",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `
			{
				"name": "jordan",
				"description": "exclusive shoes",
				"brand": "Puma",
				"variety": [
				  {
					"color": "red",
					"image": "link",
					"detail": [
					  {
						"size": 7,
						"quantity": 50,
						"price": 500
					  },
					  {
						"size": 8,
						"quantity": 100,
						"price": 500
					  }
					]
				  }
				 ]
			  }
		    `,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("CreateProduct", mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("POST", "/product", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			if test.isToken {
				req = req.WithContext(context.WithValue(req.Context(), "token", test.token))
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(addWebsitesHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetProductHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := GetProductHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:               "invalid parameter",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:  "no product found",
			input: "4",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetProductByID", mock.Anything).Return(dto.ResponseProduct{}, apperrors.NotFoundError{}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:  "success",
			input: "4",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetProductByID", mock.Anything).Return(dto.ResponseProduct{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("GET", "/product", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.input,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(addWebsitesHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestUpdateProductHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := UpdateProductHandler(websitesSvc)

	tests := []struct {
		name               string
		id                 string
		input              string
		setup              func(mockSvc *mocks.Service)
		isToken            bool
		token              dto.JwtToken
		expectedStatusCode int
	}{
		{
			name:               "invalid parameter",
			id:                 "",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
			isToken:            false,
		},
		{
			name:               "invalid parameter",
			id:                 "3",
			input:              "[]",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
			isToken:            false,
		},
		{
			name:               "invalid parameter",
			id:                 "3",
			input:              "{}",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
			isToken:            false,
		},
		{
			name:    "success",
			id:      "4",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `
			{
				"name": "jordan",
				"description": "exclusive shoes",
				"brand": "Puma",
				"variety": [
				  {
					"color": "red",
					"image": "link",
					"detail": [
					  {
						"size": 7,
						"quantity": 50,
						"price": 500
					  },
					  {
						"size": 8,
						"quantity": 100,
						"price": 500
					  }
					]
				  }
				 ]
			  }`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("UpdateProduct", mock.Anything, mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:    "success",
			id:      "4",
			isToken: true,
			token: dto.JwtToken{
				Id:   1,
				Role: 2,
			},
			input: `
			{
				"name": "jordan",
				"description": "exclusive shoes",
				"brand": "Puma",
				"variety": [
				  {
					"color": "red",
					"image": "link",
					"detail": [
					  {
						"size": 7,
						"quantity": 50,
						"price": 500
					  },
					  {
						"size": 8,
						"quantity": 100,
						"price": 500
					  }
					]
				  }
				 ]
			  }`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("UpdateProduct", mock.Anything, mock.Anything).Return(apperrors.NotFoundError{}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("PATCH", "/product", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			if test.isToken {
				req = req.WithContext(context.WithValue(req.Context(), "token", test.token))
			}

			req = mux.SetURLVars(req, map[string]string{
				"id": test.id,
			})

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(addWebsitesHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestGetProductWithFilterHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := GetProductWithFilterHandler(websitesSvc)

	tests := []struct {
		name               string
		id                 string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "success",
			id:    "",
			input: "?color=red&brand=puma&size=7",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetProductsByFilters", mock.Anything).Return([]dto.ResponseProduct{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:  "error",
			id:    "",
			input: "?color=red&brand=puma&size=7",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetProductsByFilters", mock.Anything).Return([]dto.ResponseProduct{}, errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)
			url := "/product" + test.input
			fmt.Println(url)
			req, err := http.NewRequest("GET", "/product?color=red&brand=puma&size=7", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(addWebsitesHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}
