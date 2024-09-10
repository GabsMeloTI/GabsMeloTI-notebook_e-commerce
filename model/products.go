package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string         `gorm:"type:text;not null" json:"name"`
	Description   string         `gorm:"type:text;not null" json:"description"`
	Price         float64        `gorm:"not null" json:"price"`
	StockQuantity int            `gorm:"not null" json:"stock_quantity"`
	CategoryID    uint           `gorm:"not null" json:"category_id"`
	Category      Category       `gorm:"foreignKey:CategoryID;references:ID" json:"category"`
	BrandID       uint           `gorm:"not null" json:"brand_id"`
	Brand         Brand          `gorm:"foreignKey:BrandID;references:ID" json:"brand"`
	Images        []ProductImage `gorm:"foreignKey:ProductID" json:"images"`
	Reviews       []Review       `gorm:"foreignKey:ProductID" json:"reviews"`
	OrderItems    []OrderItem    `gorm:"foreignKey:ProductID" json:"order_items"`
	StockHistory  []StockHistory `gorm:"foreignKey:ProductID" json:"stock_history"`
}
