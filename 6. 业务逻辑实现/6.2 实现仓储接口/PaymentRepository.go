package repository

import (
	"context"
	"yourdomain/domain" // Replace with your actual package path
	"github.com/google/uuid"
)

// PaymentRepository defines the interface for payment data access.
type PaymentRepository interface {
	Save(ctx context.Context, payment *domain.Payment) error
	FindById(ctx context.Context, paymentId uuid.UUID) (*domain.Payment, error)
	FindByOrderId(ctx context.Context, orderId uuid.UUID) (*domain.Payment, error)
	Delete(ctx context.Context, paymentId uuid.UUID) error
}

// InMemoryPaymentRepository is an in-memory implementation of PaymentRepository.
type InMemoryPaymentRepository struct {
	payments map[uuid.UUID]*domain.Payment
}

// Save saves a payment to the in-memory store.
func (repo *InMemoryPaymentRepository) Save(ctx context.Context, payment *domain.Payment) error {
	repo.payments[payment.PaymentId] = payment
	return nil
}

// FindById finds a payment by ID.
func (repo *InMemoryPaymentRepository) FindById(ctx context.Context, paymentId uuid.UUID) (*domain.Payment, error) {
	payment, exists := repo.payments[paymentId]
	if !exists {
		return nil, fmt.Errorf("payment not found with ID: %v", paymentId)
	}
	return payment, nil
}

// FindByOrderId finds a payment by order ID.
func (repo *InMemoryPaymentRepository) FindByOrderId(ctx context.Context, orderId uuid.UUID) (*domain.Payment, error) {
	for _, payment := range repo.payments {
		if payment.OrderId == orderId {
			return payment, nil
		}
	}
	return nil, fmt.Errorf("payment not found with order ID: %v", orderId)
}

// Delete deletes a payment from the in-memory store.
func (repo *InMemoryPaymentRepository) Delete(ctx context.Context, paymentId uuid.UUID) error {
	if _, exists := repo.payments[paymentId]; !exists {
		return fmt.Errorf("payment not found with ID: %v", paymentId)
	}
	delete(repo.payments, paymentId)
	return nil
}
