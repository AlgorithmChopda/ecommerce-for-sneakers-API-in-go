package repository

type SellerRepository interface {
	CreateSeller(sellerInfo []any) error
	CreateCompany(sellerCompanyInfo []any) (int64, error)
}
