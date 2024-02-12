package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/seller/mocks"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func TestRegisterSellerHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := RegisterSellerHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Invalid json",
			input: `[]`,
			setup: func(mockSvc *mocks.Service) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid request",
			input: `{
				"first_name": "temp"
			}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Register seller",
			input: `
		        {
		            "first_name": "abc",
					"last_name": "xyz",
					"email": "abc@gmail.com",
					"password": "123",
					"date_of_birth": "20-12-2024",
					"mobile_no": "8421556465",
					"address": "Baner",
					"city": "Pune",
					"postal_code": 411045,
					"company_name": "TEMP",
					"company_address": "silicon valley"
		        }
		    `,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterSeller", mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Register seller",
			input: `
		        {
		            "first_name": "abc",
					"last_name": "xyz",
					"email": "abc@gmail.com",
					"password": "123",
					"date_of_birth": "20-12-2024",
					"mobile_no": "8421556465",
					"address": "Baner",
					"city": "Pune",
					"postal_code": 411045,
					"company_name": "TEMP",
					"company_address": "silicon valley"
		        }
		    `,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterSeller", mock.Anything).Return(apperrors.UserAlreadyPresent{}).Once()
			},
			expectedStatusCode: http.StatusConflict,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("POST", "/seller/register", bytes.NewBuffer([]byte(test.input)))
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

func TestGetAllSellersHandlers(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := GetAllSellersHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "error",
			input: "",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetSellerList").Return(nil, errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:  "success ",
			input: "",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetSellerList").Return([]dto.SellerResponseObject{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("GET", "/seller", bytes.NewBuffer([]byte(test.input)))
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

func TestDeleteSellersHandlers(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := DeleteSellerHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "success",
			input: "4",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteSeller", mock.Anything).Return(nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:  "success",
			input: "4",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("DeleteSeller", mock.Anything).Return(apperrors.NotFoundError{}).Once()
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "success",
			input:              "",
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("DELETE", "/seller", bytes.NewBuffer([]byte(test.input)))
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
