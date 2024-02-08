package repository

import (
	"database/sql"
	"fmt"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type userStore struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userStore{
		DB: db,
	}
}

func (user *userStore) CreateUser(userInfo []any) error {
	_, err := user.DB.Exec(InsertUserQuery, userInfo...)
	if err != nil {
		return err
	}

	fmt.Println("user created")
	return nil
}
