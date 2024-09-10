package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID       uint    `gorm:"not null" json:"order_id"`
	Order         []Order `gorm:"foreignKey:OrderID;references:ID" json:"order"`
	Amount        float64 `gorm:"not null" json:"amount"`
	PaymentDate   string  `gorm:"type:datetime;not null" json:"payment_date"`
	PaymentMethod string  `gorm:"type:text;not null" json:"payment_method"`
}
