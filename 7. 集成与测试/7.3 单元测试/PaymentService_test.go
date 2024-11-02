package domain

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"yourdomain/domain" // Replace with your actual package path
)

func TestCreatePayment(t *testing.T) {
	repo := &InMemoryPaymentRepository{
		payments: make(map[uuid.UUID]*domain.Payment),
	}

	paymentService := PaymentServiceImpl{paymentRepository: repo}

	ctx := context.Background()

	// Test case 1: Create a new payment for an order
	orderId := uuid.New()
	payment := &domain.Payment{
		OrderId:     orderId,
		Amount:      decimal.Decimal{C: 100, D: 0},
		PaymentMethod: "Credit Card",
		PaymentStatus: domain.PaymentStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	err := paymentService.CreatePayment(ctx, orderId, payment.Amount, payment.PaymentMethod)
	assert.NoError(t, err)
	assert.NotNil(t, repo.payments[payment.PaymentId])
	assert.Equal(t, payment.OrderId, repo.payments[payment.PaymentId].OrderId)
	assert.Equal(t, payment.Amount, repo.payments[payment.PaymentId].Amount)

	// Test case 2: Create a payment for an order that already has a payment
	err = paymentService.CreatePayment(ctx, orderId, payment.Amount, payment.PaymentMethod)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "payment already exists for this order")
}

func TestGetPayment(t *testing.T) {
	repo := &InMemoryPaymentRepository{
		payments: make(map[uuid.UUID]*domain.Payment),
	}

	paymentService := PaymentServiceImpl{paymentRepository: repo}

	ctx := context.Background()

	// Test case 1: Get an existing payment by ID
	payment := &domain.Payment{
		PaymentId:   uuid.New(),
		OrderId:     uuid.New(),
		Amount:      decimal.Decimal{C: 100, D: 0},
		PaymentMethod: "Credit Card",
		PaymentStatus: domain.PaymentStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	repo.payments[payment.PaymentId] = payment

	gotPayment, err := paymentService.GetPayment(ctx, payment.PaymentId)
	assert.NoError(t, err)
	assert.NotNil(t, gotPayment)
	assert.Equal(t, payment.Amount, gotPayment.Amount)
	assert.Equal(t, payment.PaymentStatus, gotPayment.PaymentStatus)

	// Test case 2: Get a non-existent payment
	_, err = paymentService.GetPayment(ctx, uuid.UUID{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "payment not found with ID")
}

// Additional tests for UpdatePayment and ListPayments can be added here.
// ...