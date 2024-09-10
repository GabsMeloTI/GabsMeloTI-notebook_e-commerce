package model

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Name        string    `gorm:"type:text;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Products    []Product `gorm:"foreignKey:BrandID" json:"products"`
}
