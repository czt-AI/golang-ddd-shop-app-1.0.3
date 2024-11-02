package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// PaymentService defines the service interface for payment management.
type PaymentService interface {
	CreatePayment(ctx context.Context, orderId uuid.UUID, amount decimal.Decimal, paymentMethod string) (*Payment, error)
	GetPayment(ctx context.Context, paymentId uuid.UUID) (*Payment, error)
	UpdatePayment(ctx context.Context, paymentId uuid.UUID, status PaymentStatus) error
	ListPayments(ctx context.Context, orderId uuid.UUID) ([]Payment, error)
}

// PaymentServiceImpl implements the PaymentService interface.
type PaymentServiceImpl struct {
	paymentRepository PaymentRepository
}

// CreatePayment creates a new payment for an order.
func (s *PaymentServiceImpl) CreatePayment(ctx context.Context, orderId uuid.UUID, amount decimal.Decimal, paymentMethod string) (*Payment, error) {
	if err := s.paymentRepository.FindByOrderId(ctx, orderId); err == nil {
		return nil, errors.New("payment already exists for this order")
	}
	payment := &Payment{
		OrderId:     orderId,
		Amount:      amount,
		PaymentMethod: paymentMethod,
		PaymentStatus: PaymentStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	return s.paymentRepository.Save(ctx, payment)
}

// GetPayment retrieves a payment by ID.
func (s *PaymentServiceImpl) GetPayment(ctx context.Context, paymentId uuid.UUID) (*Payment, error) {
	return s.paymentRepository.FindById(ctx, paymentId)
}

// UpdatePayment updates a payment status.
func (s *PaymentServiceImpl) UpdatePayment(ctx context.Context, paymentId uuid.UUID, status PaymentStatus) error {
	payment, err := s.paymentRepository.FindById(ctx, paymentId)
	if err != nil {
		return err
	}
	payment.PaymentStatus = status
	return s.paymentRepository.Save(ctx, payment)
}

// ListPayments retrieves all payments for an order.
func (s *PaymentServiceImpl) ListPayments(ctx context.Context, orderId uuid.UUID) ([]Payment, error) {
	return s.paymentRepository.FindByOrderId(ctx, orderId)
}

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

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSuccess  PaymentStatus = "success"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)
