package domain

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"yourdomain/domain" // Replace with your actual package path
)

func TestCreateOrder(t *testing.T) {
	repo := &InMemoryOrderRepository{
		orders: make(map[uuid.UUID]*domain.Order),
	}

	productRepo := &InMemoryProductRepository{
		products: make(map[int]*domain.Product),
	}

	orderService := OrderServiceImpl{
		orderRepository: repo,
		productRepository: productRepo,
	}

	ctx := context.Background()

	// Test case 1: Create a new order with valid items
	orderItems := []OrderItem{
		{
			ProductId: 1,
			Quantity:  2,
		},
	}

	err := orderService.CreateOrder(ctx, 1, orderItems)
	assert.NoError(t, err)
	assert.NotNil(t, repo.orders[orderItems[0].ProductId])
	assert.Equal(t, 2, len(repo.orders[orderItems[0].ProductId].OrderItems))

	// Test case 2: Create a new order with an invalid product ID
	orderItems = []OrderItem{
		{
			ProductId: 999, // Non-existent product ID
			Quantity:  1,
		},
	}

	err = orderService.CreateOrder(ctx, 1, orderItems)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "insufficient stock for product")
}

func TestGetOrder(t *testing.T) {
	repo := &InMemoryOrderRepository{
		orders: make(map[uuid.UUID]*domain.Order),
	}

	productRepo := &InMemoryProductRepository{
		products: make(map[int]*domain.Product),
	}

	orderService := OrderServiceImpl{
		orderRepository: repo,
		productRepository: productRepo,
	}

	ctx := context.Background()

	// Test case 1: Get an existing order by ID
	order := &domain.Order{
		UserId:     1,
		TotalPrice: decimal.Decimal{C: 100, D: 0},
		OrderStatus: OrderStatusPending,
		OrderItems:  []OrderItem{
			{
				ProductId: 1,
				Quantity:  1,
			},
		},
	}

	repo.orders[order.OrderId] = order

	gotOrder, err := orderService.GetOrder(ctx, order.OrderId)
	assert.NoError(t, err)
	assert.NotNil(t, gotOrder)
	assert.Equal(t, order.TotalPrice, gotOrder.TotalPrice)
	assert.Equal(t, order.OrderStatus, gotOrder.OrderStatus)

	// Test case 2: Get a non-existent order
	_, err = orderService.GetOrder(ctx, uuid.UUID{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "order not found with ID")
}

// Additional tests for UpdateOrder, ListOrders, and CancelOrder can be added here.
// ...