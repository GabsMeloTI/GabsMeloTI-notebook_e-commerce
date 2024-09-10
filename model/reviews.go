package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	UserID    uint    `gorm:"not null" json:"user_id"`
	User      User    `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Rating    int     `gorm:"not null" json:"rating"`
	Comment   string  `gorm:"type:text" json:"comment"`
}
