package product

import (
	"errors"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type service struct {
	productRepo repository.ProductRepository
	brandRepo   repository.BrandRepository
}

type Service interface {
	CreateProduct(product dto.CreateProductRequest, sellerId int) error
	GetProductByID(productId int) (dto.ResponseProduct, error)
	UpdateProduct(req dto.UpdateProductRequest, productId int) error
	GetProductsByFilters(filters map[string]string) ([]dto.ResponseProduct, error)
}

func NewService(productRepoObject repository.ProductRepository, brandRepoObject repository.BrandRepository) Service {
	return &service{
		productRepo: productRepoObject,
		brandRepo:   brandRepoObject,
	}
}

func (productSvc *service) CreateProduct(product dto.CreateProductRequest, sellerId int) error {
	// TODO compute seller id implement transaction
	brandId, err := productSvc.brandRepo.GetBrandId(product.Brand)
	if err != nil {
		return errors.New("No such brand found")
	}

	productInfo := []any{
		product.Name,
		product.Description,
		sellerId,
		brandId,
	}

	productId, err := productSvc.productRepo.CreateProduct(productInfo)
	if err != nil {
		return err
	}

	var productDetailInfo [][]any
	for _, varities := range product.Varieties {
		for _, value := range varities.Details {
			currentProductDetail := []any{
				productId,
			}

			currentProductDetail = append(currentProductDetail, value.Size)
			currentProductDetail = append(currentProductDetail, value.Quantity)
			currentProductDetail = append(currentProductDetail, value.Price)
			currentProductDetail = append(currentProductDetail, varities.Color)
			currentProductDetail = append(currentProductDetail, varities.Image)

			productDetailInfo = append(productDetailInfo, currentProductDetail)
		}
	}

	err = productSvc.productRepo.CreateProductDetail(productDetailInfo)
	if err != nil {
		return err
	}

	return nil
}

func (productSvc *service) GetProductByID(productId int) (dto.ResponseProduct, error) {
	product, err := productSvc.productRepo.GetProductById(productId)
	if err != nil {
		return dto.ResponseProduct{}, err
	}

	return product, err
}

func (productSvc *service) UpdateProduct(req dto.UpdateProductRequest, productId int) error {
	err := productSvc.productRepo.UpdateProduct(productId, req.Name, req.Description)
	if err != nil {
		return err
	}

	return nil
}

func (productSvc *service) GetProductsByFilters(filters map[string]string) ([]dto.ResponseProduct, error) {
	productList, err := productSvc.productRepo.GetProductListWithFilters(filters)
	if err != nil {
		return nil, err
	}

	return productList, err
}
