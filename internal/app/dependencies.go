package app

import (
	"database/sql"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/seller"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/user"
	repository "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository/postgres"
)

type Dependencies struct {
	UserService   user.Service
	SellerService seller.Service
}

func NewService(db *sql.DB) Dependencies {
	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	sellerRepo := repository.NewSellerRepository(db)

	userSvc := user.NewService(userRepo, roleRepo)
	sellerSvc := seller.NewService(sellerRepo, userRepo, roleRepo)

	return Dependencies{
		UserService:   userSvc,
		SellerService: sellerSvc,
	}
}
