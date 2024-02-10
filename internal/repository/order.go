package repository

type OrderRepository interface {
	Create(userId int) (int, error)
	IsOrderPresent(userId int) (bool, error)
	GetBuyerId(orderId int) (int, error)
	AddProductToOrder(userId, cartId, productDetailId, requiredQuantity int) error
}
