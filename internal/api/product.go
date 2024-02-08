package api

import (
	"encoding/json"
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/product"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
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

		err = productSvc.CreateProduct(req)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusAccepted, nil, "Product added")
	}
}
