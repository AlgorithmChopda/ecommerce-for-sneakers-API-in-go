package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

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
	// check if product already present in cart or not
	var prevQuantity int
	var prevPrice float64

	err := order.DB.QueryRow(GetOrderItemPriceAndQuantity, cartId, productDetailId).Scan(&prevQuantity, &prevPrice)
	if err == nil {
		return apperrors.EmptyError{Message: "product already present in cart"}
	}

	// get product price and quantity
	var actualQuantity int
	var price float64
	err = order.DB.QueryRow(GetProductQuantityAndPrice, productDetailId).Scan(&actualQuantity, &price)

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
	if requiredQuantity == 0 {
		rows, err = order.DB.Exec(DeleteItemFromOrder, cartId, productDetailId)
	} else {
		rows, err = order.DB.Exec(UpdateOrderItem, productDetailId, cartId, price, requiredQuantity)
	}

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
	rows, err := order.DB.Exec(PlaceOrder, orderId, shippingAddress)
	if err != nil {
		fmt.Println(err)
		return errors.New("error while placing order")
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil || rowsAffected == 0 {
		fmt.Println(err)
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

func (order *orderStore) GetPlacedOrderDetails(userId, orderId int) (any, error) {
	var totalAmount float64
	var orderDate time.Time
	var shippingAddress string

	row := order.DB.QueryRow(GetPlacedOrderDetails, orderId, userId)
	err := row.Scan(&totalAmount, &orderDate, &shippingAddress)

	if err != nil {
		fmt.Println(err)
		return nil, apperrors.NotFoundError{Message: "no such order found"}
	}

	rows, err := order.DB.Query(GetOrderItems, orderId)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error while fetching order items")
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

	result := struct {
		TotalAmount     float64                 `json:"total_amount"`
		OrderDate       time.Time               `json:"order_date"`
		ShippingAddress string                  `json:"shipping_address"`
		OrderItems      []dto.OrderItemResponse `json:"order_items"`
	}{
		TotalAmount:     totalAmount,
		OrderDate:       orderDate,
		ShippingAddress: shippingAddress,
		OrderItems:      orderItems,
	}

	return result, nil
}

func (order *orderStore) GetUserPlacedOrders(userId int) ([]dto.UserOrderResponse, error) {
	var userOrders []dto.UserOrderResponse

	rows, err := order.DB.Query(GetUserOrders, userId)
	if err != nil {
		return nil, errors.New("error while fetching user orders")
	}

	for rows.Next() {
		// get order details
		var currentOrder dto.UserOrderResponse
		var currentOrderId int

		rows.Scan(&currentOrderId, &currentOrder.TotalAmount, &currentOrder.OrderDate, &currentOrder.ShippingAddress)

		// get order items
		rows, err := order.DB.Query(GetOrderItems, currentOrderId)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("error while fetching order items")
		}

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

			currentOrder.OrderItems = append(currentOrder.OrderItems, currentProduct)
		}

		userOrders = append(userOrders, currentOrder)
	}

	return userOrders, nil
}

func (order *orderStore) GetUpdateItemsList(orderId int) ([]int, []int, error) {
	// check if cart or not
	isPresent, err := order.IsOrderPresentWithOrderId(orderId)
	if err != nil {
		return nil, nil, err
	}

	if !isPresent {
		return nil, nil, apperrors.NotFoundError{Message: "no cart found"}
	}

	// check if cart empty or not
	orderItemCount, err := order.GetOrderItemCount(orderId)
	if err != nil {
		return nil, nil, err
	}

	if orderItemCount == 0 {
		return nil, nil, apperrors.EmptyError{Message: "cart is empty"}
	}

	// check if the items desired quantity available or not
	rows, err := order.DB.Query(GetOrderItemProductIdAndQuantity, orderId)
	if err != nil {
		fmt.Println(err)
		return nil, nil, errors.New("error while placing order")
	}

	var productDetailIdList, resultQuantity []int
	for rows.Next() {
		var curProductDetailId, curQuantity int
		err = rows.Scan(&curProductDetailId, &curQuantity)

		if err != nil {
			fmt.Println(err)
			return nil, nil, errors.New("error while placing o/rder")
		}

		var actualQuantity int
		err := order.DB.QueryRow(GetProductQuantity, curProductDetailId).Scan(&actualQuantity)
		if err != nil {
			fmt.Println(err)
			return nil, nil, errors.New("error while placing o/rder")
		}

		if actualQuantity < curQuantity {
			fmt.Println("id:", curProductDetailId, actualQuantity, curQuantity)
			return nil, nil, apperrors.InsufficientProductQuantity{}
		}

		productDetailIdList = append(productDetailIdList, curProductDetailId)
		resultQuantity = append(resultQuantity, actualQuantity-curQuantity)
	}

	return productDetailIdList, resultQuantity, nil
}

func (order *orderStore) DeleteItemFromOrder(orderId, productDetailId int) error {
	rows, err := order.DB.Exec(DeleteItemFromOrder, orderId, productDetailId)
	if err != nil {
		return errors.New("errror while updating cart item")
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return errors.New("errror while updating cart item")
	}

	if rowsAffected == 0 {
		return apperrors.NotFoundError{Message: "item not found in cart"}
	}

	return nil
}
