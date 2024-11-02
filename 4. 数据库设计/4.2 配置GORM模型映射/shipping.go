package domain

import (
	"gorm.io/gorm"
	"time"
)

type Shipping struct {
	gorm.Model
	OrderID      int
	LogisticsCompany string
	TrackingNumber string
	EstimatedDelivery time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}