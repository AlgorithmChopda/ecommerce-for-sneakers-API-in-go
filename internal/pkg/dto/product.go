package dto

import (
	"errors"
	"time"
)

type CreateProductRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Brand       string    `json:"brand"`
	Varieties   []Variety `json:"variety"`
}

type Variety struct {
	Color   string   `json:"color"`
	Image   string   `json:"image"`
	Details []Detail `json:"detail"`
}

type Detail struct {
	Size     int     `json:"size"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type ReadProduct struct {
	ProductID   int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	SellerID    int
	BrandName   string
	BrandID     int
	Size        int
	Color       string
	Image       string
	Price       float64
	Quantity    int
}

type ResponseProduct struct {
	Id          int                `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	SellerID    int                `json:"seller_id"`
	BrandID     int                `json:"brand_id"`
	BrandName   string             `json:"brand_name"`
	Varieties   []ResponseVarities `json:"variety"`
}

type ResponseVarities struct {
	Color    string  `json:"color"`
	Image    string  `json:"image"`
	Size     int     `json:"size"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type UpdateProductRequest struct {
	Name        string
	Description string
}

func (req *CreateProductRequest) Validate() error {
	err := ValidateStruct(*req)
	if err != nil {
		return err
	}

	return nil
}

func (req *UpdateProductRequest) Validate() error {
	err := ValidateStruct(*req)
	if err != nil {
		return err
	}

	return nil
}

type ProductCartRequest struct {
	Quantity int
}

func (req ProductCartRequest) Validate() error {
	if req.Quantity <= 0 {
		return errors.New("invalid product quantity")
	}

	return nil
}
