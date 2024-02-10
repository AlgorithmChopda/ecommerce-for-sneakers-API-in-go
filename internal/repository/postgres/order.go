package repository

import (
	"database/sql"
	"errors"
	"fmt"

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
