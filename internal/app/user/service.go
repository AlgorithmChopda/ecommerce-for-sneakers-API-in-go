package user

import "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"

type service struct {
	userRepo repository.UserRepository
}

type Service interface{}

func NewService(userRepoObject repository.UserRepository) Service {
	return &service{
		userRepo: userRepoObject,
	}
}
