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
	CreateProduct(product dto.CreateProductRequest) error
}

func NewService(productRepoObject repository.ProductRepository, brandRepoObject repository.BrandRepository) Service {
	return &service{
		productRepo: productRepoObject,
		brandRepo:   brandRepoObject,
	}
}

func (productSvc *service) CreateProduct(product dto.CreateProductRequest) error {
	// TODO compute seller id implement transaction
	sellerId := 1

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
