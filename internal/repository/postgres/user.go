package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/constants"
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

func (user *userStore) CreateUser(userInfo []any) error {
	_, err := user.DB.Exec(InsertUserQuery, userInfo...)
	if err != nil {
		return err
	}

	fmt.Println("user created")
	return nil
}

func (user *userStore) IsUserWithEmailPresent(email string) bool {
	row := user.DB.QueryRow(GetUserWithEmail, email)

	var userEmail string
	err := row.Scan(&userEmail)

	// if user with email does not exist
	if err != nil {
		return false
	}

	// if user found
	return true
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
		return 0, 0, "", apperrors.EmptyError{Message: "invalid username or password"}
	}

	return id, role, hashedPassword, nil
}

func (user *userStore) GetUserList(roleId int) ([]dto.UserResponseObject, error) {
	// roleId = -1 indicates both type of users
	var rows *sql.Rows
	if roleId == -1 {
		currRows, err := user.DB.Query(GetBuyerAndSellerList, constants.BUYER, constants.SELLER)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("error while fetching user records")
		}
		rows = currRows
	} else {
		currRows, err := user.DB.Query(GetBuyerOrSellerList, roleId)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("error while fetching user records")
		}
		rows = currRows
	}

	var userList []dto.UserResponseObject
	for rows.Next() {
		var currentUser dto.UserResponseObject
		err := rows.Scan(
			&currentUser.FirstName,
			&currentUser.LastName,
			&currentUser.Email,
			&currentUser.DateOfBirth,
			&currentUser.MobileNumber,
			&currentUser.Address,
			&currentUser.City,
			&currentUser.PostalCode,
			&currentUser.CreatedAt,
			&currentUser.UpdatedAt,
			&currentUser.Role,
		)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("error while fetching user records")
		}

		// Append the result to the slice
		userList = append(userList, currentUser)
	}

	return userList, nil
}
