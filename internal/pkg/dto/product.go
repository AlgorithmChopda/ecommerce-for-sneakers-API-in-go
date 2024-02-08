package dto

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

func (req *CreateProductRequest) Validate() error {
	err := ValidateStruct(*req)
	if err != nil {
		return err
	}

	return nil
}
