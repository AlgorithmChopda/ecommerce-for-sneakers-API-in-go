package app

import (
	"database/sql"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/order"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/product"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/seller"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/user"
	repository "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository/postgres"
)

type Dependencies struct {
	UserService    user.Service
	SellerService  seller.Service
	ProductService product.Service
	OrderService   order.Service
}

func NewService(db *sql.DB) Dependencies {
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	sellerRepo := repository.NewSellerRepository(db)
	productRepo := repository.NewProductRepository(db)
	brandRepo := repository.NewBrandRepository(db)
	orderRepo := repository.NewOrderRepository(db)

	userSvc := user.NewService(userRepo, roleRepo)
	sellerSvc := seller.NewService(sellerRepo, userRepo, roleRepo)
	productSvc := product.NewService(productRepo, brandRepo)
	orderSvc := order.NewService(orderRepo)

	return Dependencies{
		UserService:    userSvc,
		SellerService:  sellerSvc,
		ProductService: productSvc,
		OrderService:   orderSvc,
	}
}
