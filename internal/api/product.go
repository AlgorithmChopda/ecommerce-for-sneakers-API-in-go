package api

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/product"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/helpers"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/middleware"
)

func CreateProductHandler(productSvc product.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var req dto.CreateProductRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		// TODO validate not working for nested object
		err = req.Validate()
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		tokenData, err := helpers.GetTokenData(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
		}

		err = productSvc.CreateProduct(req, tokenData.Id)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusCreated, nil, "Product added")
	}
}

func GetProductHandler(productSvc product.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		productId, err := helpers.GetPathParameter(r, "id")
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		product, err := productSvc.GetProductByID(productId)
		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, product, "Product Fetched")
	}
}

func UpdateProductHandler(productSvc product.Service) func(w http.ResponseWriter, r *http.Request) {
	// TODO handle updated brand
	return func(w http.ResponseWriter, r *http.Request) {
		productId, err := helpers.GetPathParameter(r, "id")
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		var req dto.UpdateProductRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		err = req.Validate()
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		tokenData, err := helpers.GetTokenData(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
		}

		err = productSvc.UpdateProduct(req, productId, tokenData.Id)
		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, nil, "product updated")
	}
}

func GetProductWithFilterHandler(productSvc product.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// order, color, size, brand
		var filters = map[string]string{}
		color := r.URL.Query().Get("color")
		if color != "" {
			filters["color"] = color
		}

		size := r.URL.Query().Get("size")
		if size != "" {
			filters["size"] = size
		}

		brand := r.URL.Query().Get("brand")
		if brand != "" {
			filters["name"] = brand
		}

		// TODO add pagination
		skip, err := strconv.Atoi(r.URL.Query().Get("skip"))
		if err != nil {
			skip = 0
		}

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			limit = 10 // default limit
		}
		limit = int(math.Min(float64(limit), 100))

		product, err := productSvc.GetProductsByFilters(filters, skip, limit)
		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, product, "Product Fetched")
	}
}
