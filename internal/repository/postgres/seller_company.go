package repository

import (
	"database/sql"
)

type sellerCompanyStore struct {
	DB *sql.DB
}

type SellerRepository interface{}

func NewSelleCompanyrRepository(db *sql.DB) SellerRepository {
	return &sellerStore{
		DB: db,
	}
}
