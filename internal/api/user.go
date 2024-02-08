package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/user"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/middleware"
)

func RegisterUserHandler(userSvc user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.RegisterUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println("invalid input request")
			return
		}

		err = req.Validate()
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		err = userSvc.RegisterUser(req)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusAccepted, nil, "request successfull")
	}
}

func LoginUserHandler(userSvc user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.LoginUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Println("invalid input request")
			return
		}

		err = req.Validate()
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		token, err := userSvc.LoginUser(req.Email, req.Password)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusAccepted, struct{ token string }{token}, "login successfull")
	}
}
