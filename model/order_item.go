package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null" json:"order_id"`
	Order     Order   `gorm:"foreignKey:OrderID;references:ID" json:"order"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID" json:"product"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	UnitPrice float64 `gorm:"not null" json:"unit_price"`
}
