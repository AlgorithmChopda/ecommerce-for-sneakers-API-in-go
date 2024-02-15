package product

import (
	"errors"
	"testing"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {
	productRepo := mocks.NewProductRepository(t)
	brandRepo := mocks.NewBrandRepository(t)

	service := NewService(productRepo, brandRepo)

	tests := []struct {
		name            string
		product         dto.CreateProductRequest
		sellerId        int
		setup           func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository)
		isErrorExpected bool
	}{
		{
			name:     "error brand id",
			product:  dto.CreateProductRequest{},
			sellerId: 1,
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				brandRepo.On("GetBrandId", mock.Anything).Return(1, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name:     "error create product",
			product:  dto.CreateProductRequest{},
			sellerId: 1,
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				brandRepo.On("GetBrandId", mock.Anything).Return(1, nil).Once()
				productRepo.On("CreateProduct", mock.Anything).Return(int64(1), errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name:     "error create product",
			product:  dto.CreateProductRequest{},
			sellerId: 1,
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				brandRepo.On("GetBrandId", mock.Anything).Return(1, nil).Once()
				productRepo.On("CreateProduct", mock.Anything).Return(int64(1), nil).Once()
				productRepo.On("CreateProductDetail", mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name:     "success",
			product:  dto.CreateProductRequest{},
			sellerId: 1,
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				brandRepo.On("GetBrandId", mock.Anything).Return(1, nil)
				productRepo.On("CreateProduct", mock.Anything).Return(int64(1), nil)
				productRepo.On("CreateProductDetail", mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(productRepo, brandRepo)

			// test service
			err := service.CreateProduct(dto.CreateProductRequest{
				Name:        "abc",
				Description: "temp",
			}, 1)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestGetProductByIdHandler(t *testing.T) {
	productRepo := mocks.NewProductRepository(t)
	brandRepo := mocks.NewBrandRepository(t)

	service := NewService(productRepo, brandRepo)

	tests := []struct {
		name            string
		setup           func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository)
		isErrorExpected bool
	}{
		{
			name: "error product id",
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				productRepo.On("GetProductById", 1).Return(dto.ResponseProduct{}, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "success product id",
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				productRepo.On("GetProductById", 1).Return(dto.ResponseProduct{}, nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(productRepo, brandRepo)

			// test service
			_, err := service.GetProductByID(1)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestUpdateProductHandler(t *testing.T) {
	productRepo := mocks.NewProductRepository(t)
	brandRepo := mocks.NewBrandRepository(t)

	service := NewService(productRepo, brandRepo)

	tests := []struct {
		name            string
		setup           func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository)
		isErrorExpected bool
	}{
		{
			name: "error product id",
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				productRepo.On("UpdateProduct", 1, "", "", 1).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "success product id",
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				productRepo.On("UpdateProduct", 1, "", "", 1).Return(nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(productRepo, brandRepo)

			// test service
			err := service.UpdateProduct(dto.UpdateProductRequest{}, 1, 1)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestGetProductByFilterHandler(t *testing.T) {
	productRepo := mocks.NewProductRepository(t)
	brandRepo := mocks.NewBrandRepository(t)

	service := NewService(productRepo, brandRepo)

	tests := []struct {
		name            string
		setup           func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository)
		isErrorExpected bool
	}{
		{
			name: "error product id",
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				productRepo.On("GetProductListWithFilters", mock.Anything, 1, 1).Return([]dto.ResponseProduct{}, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "success product id",
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				productRepo.On("GetProductListWithFilters", mock.Anything, 1, 1).Return([]dto.ResponseProduct{}, nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(productRepo, brandRepo)

			// test service
			_, err := service.GetProductsByFilters(map[string]string{}, 1, 1)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}

func TestGetUpdateProductPriceAndQuantityHandler(t *testing.T) {
	productRepo := mocks.NewProductRepository(t)
	brandRepo := mocks.NewBrandRepository(t)

	service := NewService(productRepo, brandRepo)

	tests := []struct {
		name            string
		setup           func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository)
		isErrorExpected bool
	}{
		{
			name: "error product id",
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				productRepo.On("UpdateProductPriceAndQuantity", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "success product id",
			setup: func(productRepo *mocks.ProductRepository, brandRepo *mocks.BrandRepository) {
				productRepo.On("UpdateProductPriceAndQuantity", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(productRepo, brandRepo)

			// test service
			err := service.UpdateProductPriceAndQuantity(dto.UpdateProductDetailRequest{}, 1, 1)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
