package model

import "gorm.io/gorm"

type StockHistory struct {
	gorm.Model
	ProductID      uint    `gorm:"not null" json:"product_id"`
	Product        Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	QuantityChange int     `gorm:"not null" json:"quantity_change"`
	ChangeDate     string  `gorm:"type:datetime;not null" json:"change_date"`
	Reason         string  `gorm:"type:text;not null" json:"reason"`
}
