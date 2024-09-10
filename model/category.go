package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `gorm:"type:text;not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Products    []Product `gorm:"foreignKey:CategoryID" json:"products"`
}
