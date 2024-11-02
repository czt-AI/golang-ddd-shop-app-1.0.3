package domain

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID     int
	TotalPrice decimal.Decimal
	OrderStatus OrderStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	ProductID int
	Quantity  int
	Price     decimal.Decimal
}

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusPaid       OrderStatus = "paid"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

type decimal struct {
	C int64
	D int64
}

func (d decimal) Float64() float64 {
	return float64(d.C) + float64(d.D)/float64(10^int64(len(d.D)))
}

func (d decimal) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d.%d", d.C, d.D)), nil
}

func (d *decimal) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parts := strings.Split(s, ".")
	if len(parts) != 2 {
		return fmt.Errorf("invalid decimal format: %s", s)
	}
	d.C, _ = strconv.ParseInt(parts[0], 10, 64)
	d.D, _ = strconv.ParseInt(parts[1], 10, 64)
	return nil
}