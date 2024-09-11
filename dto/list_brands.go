package dto

import "time"

type BrandList struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BrandBasicList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
