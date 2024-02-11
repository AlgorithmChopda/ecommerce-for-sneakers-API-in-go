package repository

import "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"

type OrderRepository interface {
	Create(userId int) (int, error)
	IsOrderPresent(userId int) (bool, error)
	GetBuyerId(orderId int) (int, error)
	AddProductToOrder(userId, cartId, productDetailId, requiredQuantity int) error
	UpdateOrderItem(userId, cartId, productDetailId, requiredQuantity int) error
	PlaceOrder(userId, orderId int, shippingAddress string) error
	GetOrderItemCount(orderId int) (int, error)
	GetAllOrderItems(orderId int) (any, error)
	CheckOrderValid(userId, orderId int) (bool, error)
	GetPlacedOrderDetails(userId, orderId int) (any, error)
	GetUserPlacedOrders(userId int) ([]dto.UserOrderResponse, error)
}
