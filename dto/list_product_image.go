package dto

import "time"

type ProductImageList struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	URL       string    `json:"url"`
	AltText   string    `json:"alt_text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductImageBasicList struct {
	ID  uint   `json:"id"`
	URL string `json:"url"`
}
