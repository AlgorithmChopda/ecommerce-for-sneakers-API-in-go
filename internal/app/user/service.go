package user

import (
	"errors"
	"fmt"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/helpers"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepo repository.UserRepository
	roleRepo repository.RoleRepository
}

type Service interface {
	RegisterUser(userInfo dto.RegisterUserRequest) error
	LoginUser(email, passsword string) (string, error)
}

func NewService(userRepoObject repository.UserRepository, roleRepoObject repository.RoleRepository) Service {
	return &service{
		userRepo: userRepoObject,
		roleRepo: roleRepoObject,
	}
}

func (svc *service) RegisterUser(userInfo dto.RegisterUserRequest) error {
	parsedDOB, err := helpers.ParseDate(userInfo.DateOfBirth)
	if err != nil {
		return err
	}

	isPresent := svc.userRepo.IsUserWithEmailPresent(userInfo.Email)
	if isPresent {
		return apperrors.UserAlreadyPresent{}
	}

	roleId, err := svc.roleRepo.GetRoleId("buyer")
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	values := []interface{}{
		userInfo.FirstName,
		userInfo.LastName,
		userInfo.Email,
		hashedPassword,
		parsedDOB,
		userInfo.MobileNo,
		userInfo.Address,
		userInfo.City,
		userInfo.PostalCode,
		roleId,
	}

	err = svc.userRepo.CreateUser(values)
	if err != nil {
		return err
	}

	return nil
}

// TODO: handle token, register use all case - if email already exists and error handling
func (svc *service) LoginUser(email, passsword string) (string, error) {
	id, role, hashedPassword, err := svc.userRepo.GetIdRoleAndPassword(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passsword))
	if err != nil {
		return "", errors.New("Invalid login details")
	}

	token, err := helpers.CreateToken(id, role)
	if err != nil {
		fmt.Println("Error :", err)
		return "", errors.New("error while creating token : ")
	}

	return token, nil
}
