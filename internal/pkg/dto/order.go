package dto

type OrderItemResponse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Size        string  `json:"size"`
	Color       string  `json:"color"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
