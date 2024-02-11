package api

import (
	"encoding/json"
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/order"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/helpers"
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

func AddOrderHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO take userId from token
		userId := 2
		orderId, err := helpers.GetPathParameter(r, "id")
		if err != nil {
			if err != nil {
				middleware.ErrorResponse(w, http.StatusBadRequest, err)
				return
			}
		}

		productDetailId, err := helpers.GetPathParameter(r, "productDetailId")
		if err != nil {
			if err != nil {
				middleware.ErrorResponse(w, http.StatusBadRequest, err)
				return
			}
		}

		var req dto.ProductCartRequest
		err = json.NewDecoder(r.Body).Decode(&req)
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

		err = orderSvc.AddProductToOrder(userId, orderId, productDetailId, req)

		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, nil, "product added to cart")
	}
}

func UpdateOrderItemHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO take userId from token
		userId := 2
		orderId, err := helpers.GetPathParameter(r, "id")
		if err != nil {
			if err != nil {
				middleware.ErrorResponse(w, http.StatusBadRequest, err)
				return
			}
		}

		productDetailId, err := helpers.GetPathParameter(r, "productDetailId")
		if err != nil {
			if err != nil {
				middleware.ErrorResponse(w, http.StatusBadRequest, err)
				return
			}
		}

		var req dto.ProductCartRequest
		err = json.NewDecoder(r.Body).Decode(&req)
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

		err = orderSvc.UpdateProductInCart(userId, orderId, productDetailId, req)

		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, nil, "product in cart updated")
	}
}

func PlaceOrderHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO take userId from token
		userId := 2
		orderId, err := helpers.GetPathParameter(r, "id")
		if err != nil {
			if err != nil {
				middleware.ErrorResponse(w, http.StatusBadRequest, err)
				return
			}
		}

		var req struct {
			ShippingAddress string `json:"shipping_address"`
		}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		if req.ShippingAddress == "" {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		err = orderSvc.PlaceOrder(userId, orderId, req.ShippingAddress)

		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, nil, "order placed successfully")
	}
}

func GetAllOrderItemsHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO take userId from token
		userId := 2
		orderId, err := helpers.GetPathParameter(r, "id")
		if err != nil {
			if err != nil {
				middleware.ErrorResponse(w, http.StatusBadRequest, err)
				return
			}
		}

		orderItems, err := orderSvc.GetAllOrderItems(userId, orderId)

		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, orderItems, "cart items fetched successfully")
	}
}

func GetPlacedOrderDetailsHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO take userId from token
		userId := 2
		orderId, err := helpers.GetPathParameter(r, "id")
		if err != nil {
			if err != nil {
				middleware.ErrorResponse(w, http.StatusBadRequest, err)
				return
			}
		}

		orderDetails, err := orderSvc.GetPlaceOrderDetails(userId, orderId)

		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, orderDetails, "order fetched successfully")
	}
}
