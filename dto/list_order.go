package dto

import "time"

type OrderList struct {
	ID              uint      `json:"id"`
	UserID          uint      `json:"user_id"`
	TotalAmount     float64   `json:"total_amount"`
	Status          string    `json:"status"`
	ShippingAddress string    `json:"shipping_address"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type OrderBasicList struct {
	ID          uint    `json:"id"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`
}
