package dto

import "time"

type StockHistoryList struct {
	ID          uint      `json:"id"`
	ProductID   uint      `json:"product_id"`
	Quantity    int       `json:"quantity"`
	Action      string    `json:"action"` // "added" or "removed"
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type StockHistoryBasicList struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	Action    string `json:"action"`
	Quantity  int    `json:"quantity"`
}
