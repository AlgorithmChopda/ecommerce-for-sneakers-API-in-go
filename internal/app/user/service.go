package user

import (
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type service struct {
	userRepo repository.UserRepository
}

type Service interface {
	RegisterUser(userInfo dto.RegisterUserRequest) error
}

func NewService(userRepoObject repository.UserRepository) Service {
	return &service{
		userRepo: userRepoObject,
	}
}

func (svc *service) RegisterUser(userInfo dto.RegisterUserRequest) error {
	// var values []interface{}
	// v := reflect.ValueOf(userInfo).Elem()

	// for i := 0; i < v.NumField(); i++ {
	// 	values = append(values, v.Field(i).Interface())
	// }
	// values = append(values, time.Now(), time.Now())

	err := svc.userRepo.CreateUser(userInfo)
	if err != nil {
		return err
	}

	return nil
}
