package repository

type BrandRepository interface {
	GetBrandId(brand string) (int, error)
}
