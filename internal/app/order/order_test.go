package order

import (
	"errors"
	"testing"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository/mocks"
)

func TestGetUpdateProductPriceAndQuantityHandler(t *testing.T) {
	orderRepo := mocks.NewOrderRepository(t)
	productRepo := mocks.NewProductRepository(t)

	service := NewService(orderRepo, productRepo)

	tests := []struct {
		name            string
		setup           func(orderRepo *mocks.OrderRepository, productRepo *mocks.ProductRepository)
		isErrorExpected bool
	}{
		{
			name: "error product id",
			setup: func(orderRepo *mocks.OrderRepository, productRepo *mocks.ProductRepository) {
				orderRepo.On("IsOrderPresent", 1).Return(false, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "error product id",
			setup: func(orderRepo *mocks.OrderRepository, productRepo *mocks.ProductRepository) {
				orderRepo.On("IsOrderPresent", 1).Return(true, nil).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "error product id",
			setup: func(orderRepo *mocks.OrderRepository, productRepo *mocks.ProductRepository) {
				orderRepo.On("IsOrderPresent", 1).Return(false, nil).Once()
				orderRepo.On("Create", 1).Return(1, errors.New("error")).Once()
			},
			isErrorExpected: true,
		},
		{
			name: "error product id",
			setup: func(orderRepo *mocks.OrderRepository, productRepo *mocks.ProductRepository) {
				orderRepo.On("IsOrderPresent", 1).Return(false, nil).Once()
				orderRepo.On("Create", 1).Return(1, nil).Once()
			},
			isErrorExpected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.setup(orderRepo, productRepo)

			// test service
			_, err := service.CreateOrder(1)

			if (err != nil) != test.isErrorExpected {
				t.Errorf("Test Failed, expected error to be %v, but got err %v", test.isErrorExpected, err != nil)
			}
		})
	}
}
