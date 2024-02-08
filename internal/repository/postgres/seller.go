package repository

import (
	"database/sql"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type sellerStore struct {
	DB *sql.DB
}

func NewSellerRepository(db *sql.DB) repository.SellerRepository {
	return &sellerStore{
		DB: db,
	}
}

func (seller *sellerStore) CreateSeller(sellerInfo []any) error {
	_, err := seller.DB.Exec(InsertSellerQuery, sellerInfo...)
	if err != nil {
		return err
	}

	return nil
}

func (seller *sellerStore) CreateCompany(sellerCompanyInfo []any) (int64, error) {
	var companyId int64
	err := seller.DB.QueryRow(CreateCompanyQuery, sellerCompanyInfo...).Scan(&companyId)
	// row, err := seller.DB.Exec(CreateCompanyQuery, sellerCompanyInfo...)
	if err != nil {
		return 0, err
	}

	return companyId, nil
}

// func (user *userStore) IsUserWithEmailPresent(email string) bool {
// 	row := user.DB.QueryRow(GetUserWithEmail, email)

// 	var userEmail string
// 	err := row.Scan(&userEmail)

// 	// if user with email does not exist
// 	if err != nil {
// 		return false
// 	}

// 	// if user found
// 	return true
// }

// func (user *userStore) CheckUserWithEmailAndPassword(email, password string) error {
// 	row := user.DB.QueryRow(GetUserWithEmailAndPassword, email, password)

// 	var userEmail, userPassword string
// 	err := row.Scan(&userEmail, &userPassword)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (user *userStore) GetIdRoleAndPassword(email string) (int, int, string, error) {
// 	row := user.DB.QueryRow(GetIdRoleAndPassword, email)

// 	var id, role int
// 	var hashedPassword string
// 	err := row.Scan(&id, &role, &hashedPassword)

// 	if err != nil {
// 		return 0, 0, "", err
// 	}

// 	return id, role, hashedPassword, nil
// }
