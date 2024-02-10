package order

import (
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type service struct {
	orderRepo repository.OrderRepository
}

type Service interface {
	CreateOrder(userId int) (int, error)
	AddProductToOrder(userId, orderId, productDetailId int, product dto.ProductCartRequest) error
}

func NewService(orderRepoObject repository.OrderRepository) Service {
	return &service{
		orderRepo: orderRepoObject,
	}
}

func (orderSvc *service) CreateOrder(userId int) (int, error) {
	isPresent, err := orderSvc.orderRepo.IsOrderPresent(userId)
	if err != nil {
		return -1, err
	}
	if isPresent {
		return -1, apperrors.CartAlreadyPresent{}
	}

	cartId, err := orderSvc.orderRepo.Create(userId)
	if err != nil {
		return -1, err
	}

	return cartId, nil
}

func (orderSvc *service) AddProductToOrder(userId, orderId, productDetailId int, product dto.ProductCartRequest) error {
	buyerId, err := orderSvc.orderRepo.GetBuyerId(orderId)
	if err != nil {
		return err
	}

	if buyerId != userId {
		return apperrors.UnauthorizedAccess{Message: "Unauthorized access"}
	}

	err = orderSvc.orderRepo.AddProductToOrder(userId, orderId, productDetailId, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}
