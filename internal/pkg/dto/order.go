package dto

import "time"

type OrderItemResponse struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Size        string  `json:"size"`
	Color       string  `json:"color"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type UserOrderResponse struct {
	TotalAmount     float64             `json:"total_amount"`
	OrderDate       time.Time           `json:"order_date"`
	ShippingAddress string              `json:"shipping_address"`
	OrderItems      []OrderItemResponse `json:"order_items"`
}
