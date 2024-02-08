package repository

import (
	"database/sql"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type brandStore struct {
	DB *sql.DB
}

func NewBrandRepository(db *sql.DB) repository.BrandRepository {
	return &brandStore{
		DB: db,
	}
}

func (brand *brandStore) GetBrandId(brandName string) (int, error) {
	q := "select id from brand where name = $1"
	row := brand.DB.QueryRow(q, brandName)

	var id int
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
