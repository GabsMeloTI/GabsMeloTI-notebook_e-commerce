package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name              string            `gorm:"type:text;not null" json:"name"`
	Email             string            `gorm:"type:text;not null;unique" json:"email"`
	Password          string            `gorm:"type:text;not null" json:"password"`
	Address           string            `gorm:"type:text;not null" json:"address"`
	PhoneNumber       string            `gorm:"type:text;not null" json:"phone_number"`
	Orders            []Order           `gorm:"foreignKey:UserID" json:"orders"`
	ShoppingCart      ShoppingCart      `gorm:"foreignKey:UserID" json:"shopping_cart"`
	ShippingAddresses []ShippingAddress `gorm:"foreignKey:UserID" json:"shipping_addresses"`
	Reviews           []Review          `gorm:"foreignKey:UserID" json:"reviews"`
}
