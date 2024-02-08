package repository

import (
	"database/sql"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type productStore struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) repository.ProductRepository {
	return &productStore{
		DB: db,
	}
}

func (product *productStore) CreateProduct(productInfo []any) (int64, error) {
	var productId int64
	err := product.DB.QueryRow(CreateProductQuery, productInfo...).Scan(&productId)
	if err != nil {
		return 0, err
	}
	return productId, nil
}

func (product *productStore) CreateProductDetail(productDetailInfo [][]any) error {
	for _, row := range productDetailInfo {
		_, err := product.DB.Exec(CreateProductDetailQuery, row...)
		if err != nil {
			return err
		}
	}

	return nil
}
