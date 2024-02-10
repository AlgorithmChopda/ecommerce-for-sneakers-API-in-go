package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
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
		return -1, err
	}

	return buyerId, nil
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
		return errors.New("error while updating product quantity")
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return errors.New("error while updating product quantity")
	}

	if rowsAffected == 0 {
		return apperrors.NotFoundError{Message: "cart not found"}
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
