package repository

import "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"

type ProductRepository interface {
	CreateProduct(productInfo []any) (int64, error)
	CreateProductDetail(productDetailInfo [][]any) error
	GetProductById(productId int) (dto.ResponseProduct, error)
	UpdateProduct(productId int, name, description string, sellerId int) error
	UpdateProductDetail(productDetailId, quantity int) error
	GetProductListWithFilters(filters map[string]string, skip, limit int) ([]dto.ResponseProduct, error)
	UpdateProductPriceAndQuantity(sellerId, productDetailId, quantity, price int) error
}
