package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
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

func (user *userStore) CreateUser(userInfo dto.RegisterUserRequest) error {
	INSERT_USER_QUERY := `
		INSERT INTO "users" (first_name, last_name, email, password, date_of_birth, mobile_no, address, city, postal_code, role_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	parsedDOB, err := parseDate(userInfo.DateOfBirth)
	if err != nil {
		return err
	}
	// to be in service
	values := []interface{}{
		userInfo.FirstName,
		userInfo.LastName,
		userInfo.Email,
		userInfo.Password,
		parsedDOB,
		userInfo.MobileNo,
		userInfo.Address,
		userInfo.City,
		userInfo.PostalCode,
		2,          // temp value for role
		time.Now(), // created_at
		time.Now(), // updated_at
	}

	_, err = user.DB.Exec(INSERT_USER_QUERY, values...)
	if err != nil {
		return err
	}

	fmt.Println("user created")
	return nil
}

func parseDate(dateString string) (time.Time, error) {
	// Parse the user-input date string with the expected format
	parsedDate, err := time.Parse("02-01-2006", dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}
