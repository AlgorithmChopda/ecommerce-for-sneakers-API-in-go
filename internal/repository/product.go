package repository

import "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"

type ProductRepository interface {
	CreateProduct(productInfo []any) (int64, error)
	CreateProductDetail(productDetailInfo [][]any) error
	GetProductById(productId int) (dto.ResponseProduct, error)
}
