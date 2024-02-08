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

func (user *userStore) CheckUserWithEmailAndPassword(email, password string) error {
	row := user.DB.QueryRow(GetUserWithEmailAndPassword, email, password)

	var userEmail, userPassword string
	err := row.Scan(&userEmail, &userPassword)

	if err != nil {
		return err
	}

	return nil
}

func (user *userStore) GetIdRoleAndPassword(email string) (int, int, string, error) {
	row := user.DB.QueryRow(GetIdRoleAndPassword, email)

	var id, role int
	var hashedPassword string
	err := row.Scan(&id, &role, &hashedPassword)

	if err != nil {
		return 0, 0, "", err
	}

	return id, role, hashedPassword, nil
}
