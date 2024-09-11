package dto

import "time"

type ShippingDetailsList struct {
	ID            uint      `json:"id"`
	OrderID       uint      `json:"order_id"`
	Carrier       string    `json:"carrier"`
	TrackingCode  string    `json:"tracking_code"`
	EstimatedDate time.Time `json:"estimated_date"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ShippingDetailsBasicList struct {
	ID           uint   `json:"id"`
	TrackingCode string `json:"tracking_code"`
	Status       string `json:"status"`
}
