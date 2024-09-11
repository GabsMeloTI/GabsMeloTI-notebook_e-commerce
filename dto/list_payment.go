package dto

import "time"

type PaymentList struct {
	ID            uint      `json:"id"`
	OrderID       uint      `json:"order_id"`
	PaymentMethod string    `json:"payment_method"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaymentBasicList struct {
	ID     uint    `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}
