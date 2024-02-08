package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/seller"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/middleware"
)

func RegisterSellerHandler(sellerSvc seller.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.RegisterSellerRequest
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

		err = sellerSvc.RegisterSeller(req)
		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusAccepted, nil, "Seller Created")
	}
}
