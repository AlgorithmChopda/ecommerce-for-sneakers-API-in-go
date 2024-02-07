package app

import (
	"database/sql"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/app/user"
	repository "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository/postgres"
)

type Dependencies struct {
	UserService user.Service
}

func NewService(db *sql.DB) Dependencies {
	userRepo := repository.NewUserRepository(db)
	userSvc := user.NewService(userRepo)

	return Dependencies{
		UserService: userSvc,
	}
}
