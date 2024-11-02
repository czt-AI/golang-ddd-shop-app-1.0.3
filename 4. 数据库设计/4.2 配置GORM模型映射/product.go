package domain

import (
	"gorm.io/gorm"
	"decimal"
)

type Product struct {
	gorm.Model
	ProductName  string
	Description  string
	Price        decimal.Decimal
	Stock        int
	CategoryID   int
	Images       []string
	Reviews      []Review `gorm:"foreignKey:ProductID"`
}

type Review struct {
	gorm.Model
	ProductID   int
	Content     string
	Rating      float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type decimal Decimal // Use a custom type for decimal to avoid precision issues with float64