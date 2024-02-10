package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
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

func (product *productStore) GetAllProductsAndDetail() (*sql.Rows, error) {
	rows, err := product.DB.Query(GetAllProductsAndDetail)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (product *productStore) GetProductById(productId int) (dto.ResponseProduct, error) {
	rows, err := product.DB.Query(GetProductById, productId)
	if err != nil {
		fmt.Println(err)
		return dto.ResponseProduct{}, errors.New("error while fetching product")
	}

	var productObject dto.ResponseProduct
	isFirstProduct := true
	for rows.Next() {
		var readProduct dto.ReadProduct
		err = rows.Scan(
			&readProduct.ProductID, &readProduct.Name, &readProduct.Description, &readProduct.CreatedAt, &readProduct.UpdatedAt,
			&readProduct.SellerID, &readProduct.BrandID, &readProduct.Size, &readProduct.Color, &readProduct.Image,
			&readProduct.Price, &readProduct.Quantity, &readProduct.BrandName,
		)

		if err != nil {
			fmt.Println(err)
			return dto.ResponseProduct{}, errors.New("error while fetching product")
		}

		if isFirstProduct {
			productObject.Id = readProduct.ProductID
			productObject.BrandID = readProduct.BrandID
			productObject.SellerID = readProduct.SellerID
			productObject.Name = readProduct.Name
			productObject.BrandName = readProduct.BrandName
			productObject.Description = readProduct.Description
			productObject.CreatedAt = readProduct.CreatedAt
			productObject.UpdatedAt = readProduct.UpdatedAt

			isFirstProduct = false
		}

		newVariety := dto.ResponseVarities{
			Color:    readProduct.Color,
			Image:    readProduct.Image,
			Size:     readProduct.Size,
			Price:    readProduct.Price,
			Quantity: readProduct.Quantity,
		}

		productObject.Varieties = append(productObject.Varieties, newVariety)
	}

	// if no product found
	if isFirstProduct {
		return dto.ResponseProduct{}, apperrors.ProductNotFound{}
	}
	return productObject, nil
}

func (product *productStore) UpdateProduct(productId int, name, description string) error {
	res, err := product.DB.Exec(UpdateProduct, productId, name, description)
	if err != nil {
		fmt.Println(err)
		return errors.New("error while updating product")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return errors.New("error while updating product")
	}

	if rowsAffected == 0 {
		return apperrors.ProductNotFound{}
	}

	return nil
}
