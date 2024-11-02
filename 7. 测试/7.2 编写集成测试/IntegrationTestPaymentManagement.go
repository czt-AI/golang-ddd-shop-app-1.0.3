package domain

import (
	"context"
	"testing"
	"time"

	"yourdomain/domain" // Replace with your actual package path
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
)

func TestPaymentManagement(t *testing.T) {
	// Setup
	userRepository := &InMemoryUserRepository{
		users: make(map[int]*domain.User),
	}

	productRepository := &InMemoryProductRepository{
		products: make(map[int]*domain.Product),
	}

	orderRepository := &InMemoryOrderRepository{
		orders: make(map[uuid.UUID]*domain.Order),
	}

	paymentRepository := &InMemoryPaymentRepository{
		payments: make(map[uuid.UUID]*domain.Payment),
	}

	userService := UserServiceImpl{userRepository: userRepository}
	productService := ProductService{productRepository: productRepository}
	orderService := OrderServiceImpl{
		orderRepository: orderRepository,
		productRepository: productRepository,
	}
	paymentService := PaymentServiceImpl{paymentRepository: paymentRepository}

	ctx := context.Background()

	// Test case: Create a new order
	user := &domain.User{
		Username: "testuser",
	}
	userRepository.Save(ctx, user)

	product := &domain.Product{
		ProductName:  "Laptop",
		Price:        decimal.Decimal{C: 1000, D: 0},
		Stock:        10,
		CategoryId:   1,
	}
	productRepository.Save(ctx, product)

	order := &domain.Order{
		UserId:     user.UserId,
		TotalPrice: product.Price,
		OrderItems: []OrderItem{
			{
				ProductId: product.ProductId,
				Quantity:  1,
			},
		},
	}
	orderRepository.Save(ctx, order)

	// Test case: Create a new payment for the order
	payment := &domain.Payment{
		OrderId:     order.OrderId,
		Amount:      order.TotalPrice,
		PaymentMethod: "Credit Card",
		PaymentStatus: domain.PaymentStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	err := paymentService.CreatePayment(ctx, order.OrderId, payment.Amount, payment.PaymentMethod)
	assert.NoError(t, err)

	// Assert
	foundPayment, err := paymentRepository.FindById(ctx, payment.PaymentId)
	assert.NoError(t, err)
	assert.NotNil(t, foundPayment)
	assert.Equal(t, payment.OrderId, foundPayment.OrderId)
	assert.Equal(t, payment.Amount, foundPayment.Amount)

	// Test case: Update a payment status
	payment.PaymentStatus = domain.PaymentStatusSuccess
	err = paymentService.UpdatePayment(ctx, payment.PaymentId, payment.PaymentStatus)
	assert.NoError(t, err)

	// Assert
	updatedPayment, err := paymentRepository.FindById(ctx, payment.PaymentId)
	assert.NoError(t, err)
	assert.NotNil(t, updatedPayment)
	assert.Equal(t, payment.PaymentStatus, updatedPayment.PaymentStatus)
}

// Additional test cases for payment management can be added here.
// ...