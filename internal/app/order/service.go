package order

import (
	"fmt"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type service struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
}

type Service interface {
	CreateOrder(userId int) (int, error)
	AddProductToOrder(userId, orderId, productDetailId int, product dto.ProductCartRequest) error
	UpdateProductInCart(userId, orderId, productDetailId int, product dto.ProductCartRequest) error
	PlaceOrder(userId, orderId int, shipping_address string) error
	GetAllOrderItems(userId, orderId int) (any, error)
	GetPlaceOrderDetails(userId, orderId int) (any, error)
	GetUserPlacedOrders(userId int) ([]dto.UserOrderResponse, error)
}

func NewService(orderRepoObject repository.OrderRepository, productRepoObject repository.ProductRepository) Service {
	return &service{
		orderRepo:   orderRepoObject,
		productRepo: productRepoObject,
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

	isCartPresent, err := orderSvc.orderRepo.CheckOrderValid(userId, orderId)
	if err != nil {
		return err
	}

	if !isCartPresent {
		return apperrors.NotFoundError{Message: "no cart found"}
	}

	err = orderSvc.orderRepo.AddProductToOrder(userId, orderId, productDetailId, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (orderSvc *service) UpdateProductInCart(userId, orderId, productDetailId int, product dto.ProductCartRequest) error {
	buyerId, err := orderSvc.orderRepo.GetBuyerId(orderId)
	if err != nil {
		return err
	}

	if buyerId != userId {
		return apperrors.UnauthorizedAccess{Message: "Unauthorized access"}
	}

	isCartPresent, err := orderSvc.orderRepo.CheckOrderValid(userId, orderId)
	if err != nil {
		return err
	}

	if !isCartPresent {
		return apperrors.NotFoundError{Message: "no cart found"}
	}

	err = orderSvc.orderRepo.UpdateOrderItem(userId, orderId, productDetailId, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (orderSvc *service) PlaceOrder(userId, orderId int, shipping_address string) error {
	buyerId, err := orderSvc.orderRepo.GetBuyerId(orderId)
	if err != nil {
		return err
	}

	if buyerId != userId {
		return apperrors.UnauthorizedAccess{Message: "Unauthorized access"}
	}

	productDetailIdList, resultQuantity, err := orderSvc.orderRepo.GetUpdateItemsList(orderId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for i := 0; i < len(productDetailIdList); i++ {
		err := orderSvc.productRepo.UpdateProductDetail(productDetailIdList[i], resultQuantity[i])
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	err = orderSvc.orderRepo.PlaceOrder(userId, orderId, shipping_address)
	if err != nil {
		return err
	}

	return nil
}

func (orderSvc *service) GetAllOrderItems(userId, orderId int) (any, error) {
	buyerId, err := orderSvc.orderRepo.GetBuyerId(orderId)
	if err != nil {
		return nil, err
	}

	if buyerId != userId {
		return nil, apperrors.UnauthorizedAccess{Message: "Unauthorized access"}
	}

	isCartPresent, err := orderSvc.orderRepo.CheckOrderValid(userId, orderId)
	if err != nil {
		return nil, err
	}

	if !isCartPresent {
		return nil, apperrors.NotFoundError{Message: "no cart found"}
	}

	orderItems, err := orderSvc.orderRepo.GetAllOrderItems(orderId)

	if err != nil {
		return nil, err
	}

	return orderItems, err
}

func (orderSvc *service) GetPlaceOrderDetails(userId, orderId int) (any, error) {
	orderDetais, err := orderSvc.orderRepo.GetPlacedOrderDetails(userId, orderId)
	if err != nil {
		return nil, err
	}

	return orderDetais, nil
}

func (orderSvc *service) GetUserPlacedOrders(userId int) ([]dto.UserOrderResponse, error) {
	userOrders, err := orderSvc.orderRepo.GetUserPlacedOrders(userId)
	if err != nil {
		return nil, err
	}

	return userOrders, nil
}
