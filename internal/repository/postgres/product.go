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
			&readProduct.Price, &readProduct.Quantity, &readProduct.ProductDetailId, &readProduct.BrandName,
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
			Color:           readProduct.Color,
			Image:           readProduct.Image,
			Size:            readProduct.Size,
			Price:           readProduct.Price,
			Quantity:        readProduct.Quantity,
			ProductDetailId: readProduct.ProductDetailId,
		}

		productObject.Varieties = append(productObject.Varieties, newVariety)
	}

	// if no product found
	if isFirstProduct {
		return dto.ResponseProduct{}, apperrors.NotFoundError{Message: "no such product found"}
	}
	return productObject, nil
}

func (product *productStore) UpdateProduct(productId int, name, description string, sellerId int) error {
	res, err := product.DB.Exec(UpdateProduct, productId, name, description, sellerId)
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
		return apperrors.NotFoundError{Message: "no such product found"}
	}

	return nil
}

func (product *productStore) UpdateProductDetail(productDetailId, quantity int) error {
	res, err := product.DB.Exec(UpdateProductDetail, productDetailId, quantity)
	if err != nil {
		fmt.Println(err)
		return errors.New("error while updating product detail")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return errors.New("error while updating product detail")
	}

	if rowsAffected == 0 {
		return apperrors.NotFoundError{Message: "no such product found"}
	}

	return nil
}

func (product *productStore) UpdateProductPriceAndQuantity(sellerId, productDetailId, quantity, price int) error {
	res, err := product.DB.Exec(UpdateProductPriceAndQuantity, sellerId, productDetailId, quantity, price)
	if err != nil {
		fmt.Println(err)
		return errors.New("error while updating product detail")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return errors.New("error while updating product detail")
	}

	if rowsAffected == 0 {
		return apperrors.NotFoundError{Message: "no such product found"}
	}

	return nil
}

func (product *productStore) GetProductListWithFilters(filters map[string]string, skip, limit int) ([]dto.ResponseProduct, error) {
	rawQuery := getQueryForFilters(filters, skip, limit)
	rows, err := product.DB.Query(rawQuery)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error while fetching products")
	}

	var productObject []dto.ResponseProduct
	isNewProduct := true
	index, prevProductId := -1, -1

	for rows.Next() {
		var readProduct dto.ReadProduct
		err = rows.Scan(
			&readProduct.ProductID, &readProduct.Name, &readProduct.Description, &readProduct.CreatedAt, &readProduct.UpdatedAt,
			&readProduct.SellerID, &readProduct.BrandID, &readProduct.Size, &readProduct.Color, &readProduct.Image,
			&readProduct.Price, &readProduct.Quantity, &readProduct.ProductDetailId, &readProduct.BrandName,
		)

		if err != nil {
			fmt.Println(err)
			return nil, errors.New("error while fetching product")
		}

		if prevProductId != readProduct.ProductID {
			prevProductId = readProduct.ProductID
			isNewProduct = true
		}

		if isNewProduct {
			productObject = append(productObject, dto.ResponseProduct{})
			index++

			productObject[index].Id = readProduct.ProductID
			productObject[index].BrandID = readProduct.BrandID
			productObject[index].SellerID = readProduct.SellerID
			productObject[index].Name = readProduct.Name
			productObject[index].BrandName = readProduct.BrandName
			productObject[index].Description = readProduct.Description
			productObject[index].CreatedAt = readProduct.CreatedAt
			productObject[index].UpdatedAt = readProduct.UpdatedAt

			isNewProduct = false
		}

		newVariety := dto.ResponseVarities{
			Color:           readProduct.Color,
			Image:           readProduct.Image,
			Size:            readProduct.Size,
			Price:           readProduct.Price,
			Quantity:        readProduct.Quantity,
			ProductDetailId: readProduct.ProductDetailId,
		}

		productObject[index].Varieties = append(productObject[index].Varieties, newVariety)
	}

	// low := min(len(),skip * limit)
	// high := min(len(), low + limit)

	// limit = int(math.Min(float64(skip+limit), float64(len(productObject))))
	// return productObject[skip:limit], nil
	return productObject, nil
}
