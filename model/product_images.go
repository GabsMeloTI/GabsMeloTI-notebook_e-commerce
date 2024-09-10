package model

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	ImageURL  string  `gorm:"type:text;not null" json:"image_url"`
}
