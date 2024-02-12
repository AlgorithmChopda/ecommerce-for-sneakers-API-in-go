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
		tokenData, err := helpers.GetTokenData(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
		}

		cartId, err := orderSvc.CreateOrder(tokenData.Id)
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
		tokenData, err := helpers.GetTokenData(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
		}

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

		err = orderSvc.AddProductToOrder(tokenData.Id, orderId, productDetailId, req)

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
		tokenData, err := helpers.GetTokenData(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
		}

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

		if req.Quantity < 0 {
			middleware.ErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		err = orderSvc.UpdateProductInCart(tokenData.Id, orderId, productDetailId, req)

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
		tokenData, err := helpers.GetTokenData(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
		}

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

		err = orderSvc.PlaceOrder(tokenData.Id, orderId, req.ShippingAddress)

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
		tokenData, err := helpers.GetTokenData(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
		}
		orderId, err := helpers.GetPathParameter(r, "id")
		if err != nil {
			if err != nil {
				middleware.ErrorResponse(w, http.StatusBadRequest, err)
				return
			}
		}

		orderItems, err := orderSvc.GetAllOrderItems(tokenData.Id, orderId)

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
		tokenData, err := helpers.GetTokenData(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
		}
		orderId, err := helpers.GetPathParameter(r, "id")
		if err != nil {
			if err != nil {
				middleware.ErrorResponse(w, http.StatusBadRequest, err)
				return
			}
		}

		orderDetails, err := orderSvc.GetPlaceOrderDetails(tokenData.Id, orderId)

		if err != nil {
			status, err := apperrors.MapError(err)
			middleware.ErrorResponse(w, status, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, orderDetails, "order fetched successfully")
	}
}

func GetUserPlacedOrderHandler(orderSvc order.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenData, err := helpers.GetTokenData(r.Context())
		if err != nil {
			middleware.ErrorResponse(w, http.StatusUnauthorized, err)
		}
		userOrders, err := orderSvc.GetUserPlacedOrders(tokenData.Id)

		if err != nil {
			middleware.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}

		middleware.SuccessResponse(w, http.StatusOK, userOrders, "user orders fetched successfully")
	}
}
