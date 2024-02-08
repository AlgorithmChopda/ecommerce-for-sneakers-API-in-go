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
