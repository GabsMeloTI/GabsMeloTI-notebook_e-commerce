package model

import "gorm.io/gorm"

type ShippingAddress struct {
	gorm.Model
	UserID     uint   `gorm:"not null" json:"user_id"`
	User       User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Address    string `gorm:"type:text;not null" json:"address"`
	City       string `gorm:"type:text;not null" json:"city"`
	State      string `gorm:"type:text;not null" json:"state"`
	PostalCode string `gorm:"type:text;not null" json:"postal_code"`
}
