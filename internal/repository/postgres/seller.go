package repository

import (
	"database/sql"
	"errors"

	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/apperrors"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/pkg/dto"
	"github.com/AlgorithmChopda/ecommerce-for-sneakers-API-in-go/internal/repository"
)

type sellerStore struct {
	DB *sql.DB
}

func NewSellerRepository(db *sql.DB) repository.SellerRepository {
	return &sellerStore{
		DB: db,
	}
}

func (seller *sellerStore) CreateSeller(sellerInfo []any) error {
	_, err := seller.DB.Exec(InsertSellerQuery, sellerInfo...)
	if err != nil {
		return err
	}

	return nil
}

func (seller *sellerStore) CreateCompany(sellerCompanyInfo []any) (int64, error) {
	var companyId int64
	err := seller.DB.QueryRow(CreateCompanyQuery, sellerCompanyInfo...).Scan(&companyId)
	if err != nil {
		return 0, err
	}

	return companyId, nil
}

func (seller *sellerStore) GetAllSellers(roleId int) ([]dto.SellerResponseObject, error) {
	rows, err := seller.DB.Query(GetAllSellers, roleId)
	if err != nil {
		return nil, errors.New("error while fetching seller details")
	}

	var sellerList []dto.SellerResponseObject
	for rows.Next() {
		var currentSeller dto.SellerResponseObject
		err := rows.Scan(
			&currentSeller.FirstName,
			&currentSeller.LastName,
			&currentSeller.Email,
			&currentSeller.DateOfBirth,
			&currentSeller.MobileNumber,
			&currentSeller.Address,
			&currentSeller.City,
			&currentSeller.PostalCode,
			&currentSeller.CreatedAt,
			&currentSeller.UpdatedAt,
			&currentSeller.CompanyName,
			&currentSeller.CompanyAddress,
		)
		if err != nil {
			return nil, errors.New("error while fetching seller details")
		}

		// Append the result to the slice
		sellerList = append(sellerList, currentSeller)
	}

	return sellerList, nil
}

func (seller *sellerStore) DeleteSeller(sellerId int) error {
	rows, err := seller.DB.Exec(DeleteSeller, sellerId)
	if err != nil {
		return errors.New("error while deleting the seller")
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return errors.New("error while deleting the seller")
	}

	if rowsAffected == 0 {
		return apperrors.NotFoundError{Message: "seller not found"}
	}

	return nil
}
