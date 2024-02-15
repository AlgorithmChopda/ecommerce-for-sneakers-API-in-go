package api

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/user/mocks"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/stretchr/testify/mock"
)

func TestLoginUserHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := LoginUserHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name: "Login User",
			input: `
                {
                    "email": "naman@gmail.com",
    				"password": "123"
                }
            `,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("LoginUser", mock.Anything, mock.Anything).Return("token", nil).Once()
			},
			expectedStatusCode: http.StatusAccepted,
		},
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
				"temp": "naman@gmail.com",
				"password": "123"
			}`,
			setup: func(mockSvc *mocks.Service) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Login User",
			input: `
                {
                    "email": "tempuser@gmail.com",
    				"password": "123"
                }
            `,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("LoginUser", mock.Anything, mock.Anything).Return("", apperrors.EmptyError{}).Once()
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("POST", "/user/login", bytes.NewBuffer([]byte(test.input)))
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

func TestRegisterUserHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := RegisterUserHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Invalid json",
			input: `[`,
			setup: func(mockSvc *mocks.Service) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid request",
			input: `{
				"email": "123@com"
			}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Register successfull",
			input: `{
				"first_name": "naman",
				"last_name": "chopda",
				"email": "naman12@gmail.com",
				"password": "kjsdlffkj1H",
				"date_of_birth": "20-12-2024",
				"mobile_no": "8421556465",
				"address": "Baner",
				"city": "Pune",
				"postal_code": 411045
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterUser", mock.Anything, mock.Anything).Return(dto.UserRegisterResponseObject{}, nil).Once()
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name: "Email already present",
			input: `{
				"first_name": "naman",
				"last_name": "chopda",
				"email": "naman12@gmail.com",
				"password": "12dkksdH3",
				"date_of_birth": "20-12-2024",
				"mobile_no": "8421556465",
				"address": "Baner",
				"city": "Pune",
				"postal_code": 411045
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterUser", mock.Anything, mock.Anything).Return(dto.UserRegisterResponseObject{}, apperrors.UserAlreadyPresent{}).Once()
			},
			expectedStatusCode: http.StatusConflict,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer([]byte(test.input)))
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

func TestGetUserListHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	getUserListHandler := GetUserListHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "error getting lst",
			input: "",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetUserList", mock.Anything).Return([]dto.UserResponseObject{}, errors.New("error")).Once()
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:  "success getting lst",
			input: "",
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("GetUserList", mock.Anything).Return([]dto.UserResponseObject{}, nil).Once()
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("GET", "/user", bytes.NewBuffer([]byte(test.input)))
			if err != nil {
				t.Fatal(err)
				return
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getUserListHandler)
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode != test.expectedStatusCode {
				t.Errorf("Expected %d but got %d", test.expectedStatusCode, rr.Result().StatusCode)
			}
		})
	}
}

func TestAdminRegisterHandler(t *testing.T) {
	websitesSvc := mocks.NewService(t)
	addWebsitesHandler := RegisterAdminHandler(websitesSvc)

	tests := []struct {
		name               string
		input              string
		setup              func(mockSvc *mocks.Service)
		expectedStatusCode int
	}{
		{
			name:  "Invalid json",
			input: `[`,
			setup: func(mockSvc *mocks.Service) {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Invalid request",
			input: `{
				"email": "123@com"
			}`,
			setup:              func(mockSvc *mocks.Service) {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "Register successfull",
			input: `{
				"first_name": "naman",
				"last_name": "chopda",
				"email": "naman12@gmail.com",
				"password": "kjsdlffkj1H",
				"date_of_birth": "20-12-2024",
				"mobile_no": "8421556465",
				"address": "Baner",
				"city": "Pune",
				"postal_code": 411045
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterUser", mock.Anything, mock.Anything).Return(dto.UserRegisterResponseObject{}, nil).Once()
			},
			expectedStatusCode: http.StatusAccepted,
		},
		{
			name: "Email already present",
			input: `{
				"first_name": "naman",
				"last_name": "chopda",
				"email": "naman12@gmail.com",
				"password": "12dkksdH3",
				"date_of_birth": "20-12-2024",
				"mobile_no": "8421556465",
				"address": "Baner",
				"city": "Pune",
				"postal_code": 411045
			}`,
			setup: func(mockSvc *mocks.Service) {
				mockSvc.On("RegisterUser", mock.Anything, mock.Anything).Return(dto.UserRegisterResponseObject{}, apperrors.UserAlreadyPresent{}).Once()
			},
			expectedStatusCode: http.StatusConflict,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(websitesSvc)

			req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer([]byte(test.input)))
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
