package user

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/constants"
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
	RegisterUser(userInfo dto.RegisterUserRequest, userRole string) (dto.UserRegisterResponseObject, error)
	LoginUser(email, passsword string) (string, error)
	GetUserList(r *http.Request) ([]dto.UserResponseObject, error)
	GetUserProfile(userId int) (dto.UserResponseObject, error)
}

func NewService(userRepoObject repository.UserRepository, roleRepoObject repository.RoleRepository) Service {
	return &service{
		userRepo: userRepoObject,
		roleRepo: roleRepoObject,
	}
}

func (svc *service) RegisterUser(userInfo dto.RegisterUserRequest, userRole string) (dto.UserRegisterResponseObject, error) {
	fmt.Println(userInfo.DateOfBirth)
	parsedDOB, err := helpers.ParseDate(userInfo.DateOfBirth)
	if err != nil {
		return dto.UserRegisterResponseObject{}, err
	}

	isPresent := svc.userRepo.IsUserWithEmailPresent(userInfo.Email)
	if isPresent {
		return dto.UserRegisterResponseObject{}, apperrors.UserAlreadyPresent{}
	}

	roleId, err := svc.roleRepo.GetRoleId(userRole)
	if err != nil {
		return dto.UserRegisterResponseObject{}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.UserRegisterResponseObject{}, err
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
		return dto.UserRegisterResponseObject{}, err
	}

	userResponse := dto.UserRegisterResponseObject{
		FirstName:    userInfo.FirstName,
		LastName:     userInfo.LastName,
		Email:        userInfo.Email,
		DateOfBirth:  userInfo.DateOfBirth,
		MobileNumber: userInfo.MobileNo,
		Address:      userInfo.Address,
		City:         userInfo.City,
		PostalCode:   userInfo.PostalCode,
	}
	return userResponse, nil
}

func (svc *service) LoginUser(email, passsword string) (string, error) {
	id, role, hashedPassword, err := svc.userRepo.GetIdRoleAndPassword(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passsword))
	if err != nil {
		fmt.Println("in here")
		return "", apperrors.EmptyError{Message: "invalid login details"}
	}

	token, err := helpers.CreateToken(id, role)
	if err != nil {
		fmt.Println("Error :", err)
		return "", errors.New("error while creating token : ")
	}

	return token, nil
}

func (svc *service) GetUserList(r *http.Request) ([]dto.UserResponseObject, error) {
	role := r.URL.Query().Get("type")
	roleType := -1

	switch role {
	case "seller":
		roleType = constants.SELLER
	case "buyer":
		roleType = constants.BUYER
	default:
		if role == "" {
			roleType = -1
		} else {
			return nil, apperrors.EmptyError{Message: "type not found"}
		}
	}

	userList, err := svc.userRepo.GetUserList(roleType)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (svc *service) GetUserProfile(userId int) (dto.UserResponseObject, error) {
	userDetail, err := svc.userRepo.GetUserProfile(userId)
	if err != nil {
		return dto.UserResponseObject{}, err
	}

	return userDetail, nil
}
