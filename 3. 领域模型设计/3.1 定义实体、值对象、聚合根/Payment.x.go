package domain

// Payment defines the structure for a payment entity.
type Payment struct {
	PaymentId   int
	OrderId     int
	Amount      decimal.Decimal
	PaymentMethod string
	PaymentStatus PaymentStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewPayment creates a new Payment instance.
func NewPayment(paymentId, orderId int, amount decimal.Decimal, paymentMethod, paymentStatus string, createdAt, updatedAt time.Time) *Payment {
	return &Payment{
		PaymentId:   paymentId,
		OrderId:     orderId,
		Amount:      amount,
		PaymentMethod: paymentMethod,
		PaymentStatus: paymentStatus,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

// PaymentStatus defines the status of a payment.
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSuccess  PaymentStatus = "success"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

// PaymentDTO defines the data transfer object for a payment.
type PaymentDTO struct {
	PaymentId   int
	OrderId     int
	Amount      float64
	PaymentMethod string
	PaymentStatus PaymentStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ToDTO converts a Payment to a PaymentDTO.
func (p *Payment) ToDTO() *PaymentDTO {
	return &PaymentDTO{
		PaymentId:   p.PaymentId,
		OrderId:     p.OrderId,
		Amount:      p.Amount.Float64(),
		PaymentMethod: p.PaymentMethod,
		PaymentStatus: p.PaymentStatus,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}
