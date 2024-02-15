package api

import (
	"encoding/json"
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/seller"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/helpers"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/middleware"
)

func RegisterSellerHandler(sellerSvc seller.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.RegisterSellerRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
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

		middleware.SuccessResponse(w, http.StatusCreated, nil, "Seller Created")
	}
}

func GetAllSellersHandler(sellerSvc seller.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sellerList, err := sellerSvc.GetSellerList()

		if err != nil {
			middleware.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, sellerList, "Sellers fetched successfully")
	}
}

func DeleteSellerHandler(sellerSvc seller.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sellerId, err := helpers.GetPathParameter(r, "id")
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		err = sellerSvc.DeleteSeller(sellerId)
		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, nil, "seller deleted successfully")
	}
}

func GetSellerMonthlySaleHandler(sellerSvc seller.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sellerId, err := helpers.GetTokenData(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		sellerSale, err := sellerSvc.GetSellerMonthlySale(sellerId.Id)
		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, sellerSale, "seller monthly sale fetched successfully")
	}
}
