package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type orderStore struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) repository.OrderRepository {
	return &orderStore{
		DB: db,
	}
}

func (order *orderStore) Create(userId int) (int, error) {
	var orderId int
	err := order.DB.QueryRow(CreateOrder, userId).Scan(&orderId)
	if err != nil {
		fmt.Println(err)
		return -1, errors.New("error while creating cart")
	}

	return orderId, nil
}

func (order *orderStore) IsOrderPresent(userId int) (bool, error) {
	rows, err := order.DB.Query(OrderWithUserId, userId)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("error while checking cart for user")
	}

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (order *orderStore) GetBuyerId(orderId int) (int, error) {
	var buyerId int

	row := order.DB.QueryRow(GetBuyerIdOfOrder, orderId)
	err := row.Scan(&buyerId)

	if err != nil {
		return -1, apperrors.NotFoundError{Message: "no cart found"}
	}

	return buyerId, nil
}

func (order *orderStore) CheckOrderValid(userId, orderId int) (bool, error) {
	rows, err := order.DB.Query(CheckOrderValid, userId, orderId)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("error while checking cart for user")
	}

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (order *orderStore) AddProductToOrder(userId, cartId, productDetailId, requiredQuantity int) error {
	// get product price and quantity
	var actualQuantity int
	var price float64
	err := order.DB.QueryRow(GetProductQuantityAndPrice, productDetailId).Scan(&actualQuantity, &price)

	if err != nil {
		return apperrors.NotFoundError{Message: "no such product found"}
	}

	if actualQuantity < requiredQuantity {
		return apperrors.InsufficientProductQuantity{}
	}

	totalProductAmount := float64(requiredQuantity) * price

	rows, err := order.DB.Exec(UpdateOrderAmount, cartId, totalProductAmount)
	if err != nil {
		return errors.New("error while adding product to cart")
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return errors.New("error while adding product to cart")
	}

	if rowsAffected == 0 {
		return apperrors.NotFoundError{Message: "no cart found"}
	}

	// TODO handle if product already present

	// add product to cart
	rows, err = order.DB.Exec(AddProductToOrder, productDetailId, cartId, price, requiredQuantity)
	if err != nil {
		return errors.New("error while adding product to cart")
	}

	rowsAffected, err = rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("error while adding product to cart")
	}
	return nil
}

func (order *orderStore) UpdateOrderItem(userId, cartId, productDetailId, requiredQuantity int) error {
	// execution reaching here says cart belongs to the authorized user

	// get product price and quantity
	var actualQuantity int
	var price float64
	err := order.DB.QueryRow(GetProductQuantityAndPrice, productDetailId).Scan(&actualQuantity, &price)

	if err != nil {
		return apperrors.NotFoundError{Message: "no such product found"}
	}

	if actualQuantity < requiredQuantity {
		return apperrors.InsufficientProductQuantity{}
	}

	// get previous order item price and quantity
	var prevQuantity int
	var prevPrice float64

	order.DB.QueryRow(GetOrderItemPriceAndQuantity, cartId, productDetailId).Scan(&prevQuantity, &prevPrice)
	if err != nil {
		return apperrors.NotFoundError{Message: "no such product found in cart"}
	}
	prevTotalAmountOfItem := float64(prevQuantity * int(prevPrice))

	// update cart total amount
	totalProductAmount := (float64(requiredQuantity) * price) - prevTotalAmountOfItem

	rows, err := order.DB.Exec(UpdateOrderAmount, cartId, totalProductAmount)
	if err != nil {
		return errors.New("error while updating product quantity")
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return errors.New("error while updating product quantity")
	}

	if rowsAffected == 0 {
		return apperrors.NotFoundError{Message: "no cart found"}
	}

	// update product in cart
	rows, err = order.DB.Exec(UpdateOrderItem, productDetailId, cartId, price, requiredQuantity)
	if err != nil {
		return errors.New("error while Updating product to cart")
	}

	rowsAffected, err = rows.RowsAffected()
	if err != nil {
		return errors.New("error while adding product to cart")
	}

	// TODO handle if quantity is set to 0
	return nil
}

func (order *orderStore) PlaceOrder(userId, orderId int, shippingAddress string) error {
	isPresent, err := order.IsOrderPresentWithOrderId(orderId)
	if err != nil {
		return err
	}

	if !isPresent {
		return apperrors.NotFoundError{Message: "no cart found"}
	}

	orderItemCount, err := order.GetOrderItemCount(orderId)
	if err != nil {
		return err
	}

	if orderItemCount == 0 {
		return apperrors.EmptyError{Message: "cart is empty"}
	}

	rows, err := order.DB.Exec(PlaceOrder, orderId, shippingAddress)
	if err != nil {
		return errors.New("error while placing order")
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return errors.New("error while placing order")
	}

	return nil
}

func (order *orderStore) GetOrderItemCount(orderId int) (int, error) {
	var orderItemCount int
	err := order.DB.QueryRow(GetOrderItemCount, orderId).Scan(&orderItemCount)
	if err != nil {
		return -1, apperrors.NotFoundError{Message: "no such cart found"}
	}

	return orderItemCount, nil
}

func (order *orderStore) IsOrderPresentWithOrderId(orderId int) (bool, error) {
	rows, err := order.DB.Query(OrderWithOrderId, orderId)
	if err != nil {
		fmt.Println(err)
		return false, errors.New("error while checking order for user")
	}

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (order *orderStore) GetAllOrderItems(orderId int) (any, error) {
	rows, err := order.DB.Query(GetOrderItems, orderId)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error while fetching cart items")
	}

	var orderItems []dto.OrderItemResponse

	for rows.Next() {
		var currentProduct dto.OrderItemResponse
		err := rows.Scan(
			&currentProduct.Name,
			&currentProduct.Description,
			&currentProduct.Size,
			&currentProduct.Color,
			&currentProduct.Image,
			&currentProduct.Price,
			&currentProduct.Quantity,
		)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("error while fetching cart items")
		}

		orderItems = append(orderItems, currentProduct)
	}

	var totalAmount float64
	err = order.DB.QueryRow(GetOrderAmount, orderId).Scan(&totalAmount)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error while fetching cart items")
	}

	result := struct {
		TotalAmount float64                 `json:"total_amount"`
		OrderItems  []dto.OrderItemResponse `json:"order_items"`
	}{
		TotalAmount: totalAmount,
		OrderItems:  orderItems,
	}

	return result, nil
}
