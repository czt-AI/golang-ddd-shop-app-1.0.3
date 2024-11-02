package domain

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	OrderID     int
	Amount      decimal.Decimal
	PaymentMethod string
	PaymentStatus PaymentStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSuccess  PaymentStatus = "success"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

type decimal struct {
	C int64
	D int64
}

func (d decimal) Float64() float64 {
	return float64(d.C) + float64(d.D)/float64(10^int64(len(d.D)))
}

func (d *decimal) MarshalJSON() ([]byte, error) {
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