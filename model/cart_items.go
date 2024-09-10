package model

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID       uint         `gorm:"not null" json:"cart_id"`
	ShoppingCart ShoppingCart `gorm:"foreignKey:CartID;references:ID" json:"shopping_cart"`
	ProductID    uint         `gorm:"not null" json:"product_id"`
	Product      Product      `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	Quantity     int          `gorm:"not null" json:"quantity"`
}
