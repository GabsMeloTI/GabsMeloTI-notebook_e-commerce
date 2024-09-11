package dto

import "time"

type CategoryList struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CategoryBasicList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
