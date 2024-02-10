package api

import (
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/order"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/middleware"
)

func CreateOrderHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO take userId from token
		userId := 2
		cartId, err := orderSvc.CreateOrder(userId)
		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusAccepted, cartId, "Cart created")
	}
}
