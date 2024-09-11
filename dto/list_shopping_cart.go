package dto

import "time"

type ShoppingCartList struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ShoppingCartBasicList struct {
	ID        uint `json:"id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
