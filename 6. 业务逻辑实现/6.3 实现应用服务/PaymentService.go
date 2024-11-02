package service

import (
	"context"
	"errors"

	"yourdomain/domain" // Replace with your actual package path
)

// PaymentService defines the service interface for payment management.
type PaymentService interface {
	CreatePayment(ctx context.Context, orderId uuid.UUID, amount decimal.Decimal, paymentMethod string) (*domain.Payment, error)
	GetPayment(ctx context.Context, paymentId uuid.UUID) (*domain.Payment, error)
	UpdatePayment(ctx context.Context, paymentId uuid.UUID, status domain.PaymentStatus) error
	ListPayments(ctx context.Context, orderId uuid.UUID) ([]domain.Payment, error)
}

// PaymentServiceImpl implements the PaymentService interface.
type PaymentServiceImpl struct {
	paymentRepository domain.PaymentRepository
}

// CreatePayment creates a new payment for an order.
func (s *PaymentServiceImpl) CreatePayment(ctx context.Context, orderId uuid.UUID, amount decimal.Decimal, paymentMethod string) (*domain.Payment, error) {
	if err := s.paymentRepository.FindByOrderId(ctx, orderId); err == nil {
		return nil, errors.New("payment already exists for this order")
	}
	payment := &domain.Payment{
		OrderId:     orderId,
		Amount:      amount,
		PaymentMethod: paymentMethod,
		PaymentStatus: domain.PaymentStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	return s.paymentRepository.Save(ctx, payment)
}

// GetPayment retrieves a payment by ID.
func (s *PaymentServiceImpl) GetPayment(ctx context.Context, paymentId uuid.UUID) (*domain.Payment, error) {
	return s.paymentRepository.FindById(ctx, paymentId)
}

// UpdatePayment updates a payment status.
func (s *PaymentServiceImpl) UpdatePayment(ctx context.Context, paymentId uuid.UUID, status domain.PaymentStatus) error {
	payment, err := s.paymentRepository.FindById(ctx, paymentId)
	if err != nil {
		return err
	}
	payment.PaymentStatus = status
	return s.paymentRepository.Save(ctx, payment)
}

// ListPayments retrieves all payments for an order.
func (s *PaymentServiceImpl) ListPayments(ctx context.Context, orderId uuid.UUID) ([]domain.Payment, error) {
	return s.paymentRepository.FindByOrderId(ctx, orderId)
}
