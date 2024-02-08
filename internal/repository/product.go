package repository

type ProductRepository interface {
	CreateProduct(productInfo []any) (int64, error)
	CreateProductDetail(productDetailInfo [][]any) error
}
