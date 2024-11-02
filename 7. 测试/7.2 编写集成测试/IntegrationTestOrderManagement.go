package domain

import (
	"context"
	"testing"
	"time"

	"yourdomain/domain" // Replace with your actual package path
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
)

func TestOrderManagement(t *testing.T) {
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

	userService := UserServiceImpl{userRepository: userRepository}
	productService := ProductService{productRepository: productRepository}
	orderService := OrderServiceImpl{
		orderRepository: orderRepository,
		productRepository: productRepository,
	}

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

	orderItems := []OrderItem{
		{
			ProductId: product.ProductId,
			Quantity:  1,
		},
	}

	err := orderService.CreateOrder(ctx, user.UserId, orderItems)
	assert.NoError(t, err)

	// Assert
	foundOrder, err := orderRepository.FindById(ctx, orderItems[0].ProductId)
	assert.NoError(t, err)
	assert.NotNil(t, foundOrder)
	assert.Equal(t, user.UserId, foundOrder.UserId)
	assert.Equal(t, len(orderItems), len(foundOrder.OrderItems))

	// Test case: Update an order status
	err = orderService.UpdateOrder(ctx, foundOrder.OrderId, OrderStatusPaid)
	assert.NoError(t, err)

	// Assert
	updatedOrder, err := orderRepository.FindById(ctx, foundOrder.OrderId)
	assert.NoError(t, err)
	assert.NotNil(t, updatedOrder)
	assert.Equal(t, OrderStatusPaid, updatedOrder.OrderStatus)

	// Test case: Cancel an order
	err = orderService.CancelOrder(ctx, foundOrder.OrderId)
	assert.NoError(t, err)

	// Assert
	canceledOrder, err := orderRepository.FindById(ctx, foundOrder.OrderId)
	assert.NoError(t, err)
	assert.NotNil(t, canceledOrder)
	assert.Equal(t, OrderStatusCancelled, canceledOrder.OrderStatus)
}

// Additional test cases for order management can be added here.
// ...