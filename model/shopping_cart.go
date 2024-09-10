package model

import "gorm.io/gorm"

type ShoppingCart struct {
	gorm.Model
	UserID    uint       `gorm:"not null" json:"user_id"`
	User      []User     `gorm:"foreignKey:UserID;references:ID" json:"user"`
	CartItems []CartItem `gorm:"foreignKey:CartID" json:"cart_items"`
}
