package repository

import "github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"

type SellerRepository interface {
	CreateSeller(sellerInfo []any) error
	CreateCompany(sellerCompanyInfo []any) (int64, error)
	GetAllSellers(roleId int) ([]dto.SellerResponseObject, error)
}
