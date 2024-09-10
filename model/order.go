package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint        `gorm:"not null" json:"user_id"`
	User        User        `gorm:"foreignKey:UserID;references:ID" json:"user"`
	OrderDate   string      `gorm:"type:datetime;not null" json:"order_date"`
	Status      string      `gorm:"type:text;not null" json:"status"`
	TotalAmount float64     `gorm:"not null" json:"total_amount"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
	Payment     Payment     `gorm:"foreignKey:OrderID" json:"payment"`
}
