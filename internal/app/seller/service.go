package seller

import (
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/helpers"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	sellerRepo repository.SellerRepository
	userRepo   repository.UserRepository
	roleRepo   repository.RoleRepository
}

type Service interface {
	RegisterSeller(sellerInfo dto.RegisterSellerRequest) error
	GetSellerList() ([]dto.SellerResponseObject, error)
}

func NewService(sellerRepoObject repository.SellerRepository, userRepoObject repository.UserRepository, roleRepoObject repository.RoleRepository) Service {
	return &service{
		sellerRepo: sellerRepoObject,
		userRepo:   userRepoObject,
		roleRepo:   roleRepoObject,
	}
}

func (sellerSvc *service) RegisterSeller(sellerInfo dto.RegisterSellerRequest) error {
	parsedDOB, err := helpers.ParseDate(sellerInfo.DateOfBirth)
	if err != nil {
		return err
	}

	isPresent := sellerSvc.userRepo.IsUserWithEmailPresent(sellerInfo.Email)
	if isPresent {
		return apperrors.UserAlreadyPresent{}
	}

	roleId, err := sellerSvc.roleRepo.GetRoleId("seller")
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(sellerInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	companyDetail := []interface{}{
		sellerInfo.CompanyName,
		sellerInfo.CompanyAddress,
	}

	companyId, err := sellerSvc.sellerRepo.CreateCompany(companyDetail)
	if err != nil {
		return err
	}

	values := []interface{}{
		sellerInfo.FirstName,
		sellerInfo.LastName,
		sellerInfo.Email,
		hashedPassword,
		parsedDOB,
		sellerInfo.MobileNo,
		sellerInfo.Address,
		sellerInfo.City,
		sellerInfo.PostalCode,
		roleId,
		companyId,
	}

	err = sellerSvc.sellerRepo.CreateSeller(values)
	if err != nil {
		return err
	}

	return nil
}

func (sellerSvc *service) GetSellerList() ([]dto.SellerResponseObject, error) {
	roleId, err := sellerSvc.roleRepo.GetRoleId("seller")
	if err != nil {
		return nil, err
	}

	sellerList, err := sellerSvc.sellerRepo.GetAllSellers(roleId)
	if err != nil {
		return nil, err
	}

	return sellerList, nil
}
