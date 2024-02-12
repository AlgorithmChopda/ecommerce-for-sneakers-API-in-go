package api

import (
	"bytes"
	"errors"
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
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:               "invalid json",
			input:              `[]`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "invalid request format",
			input:              `{}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Error adding Product",
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
			name: "Success adding Product",
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
				mockSvc.On("CreateProduct", mock.Anything).Return(nil).Once()
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
		expectedStatusCode int
	}{
		{
			name:               "invalid parameter",
			id:                 "",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "invalid parameter",
			id:                 "3",
			input:              "[]",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "invalid parameter",
			id:                 "3",
			input:              "{}",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "success",
			id:   "4",
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
			name: "success",
			id:   "4",
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
