package domain

import (
	"context"
	"testing"
	"time"

	"yourdomain/domain" // Replace with your actual package path
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
)

func TestShippingManagement(t *testing.T) {
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

	shippingRepository := &InMemoryShippingRepository{
		shipments: make(map[uuid.UUID]*domain.Shipping),
	}

	userService := UserServiceImpl{userRepository: userRepository}
	productService := ProductService{productRepository: productRepository}
	orderService := OrderServiceImpl{
		orderRepository: orderRepository,
		productRepository: productRepository,
	}
	shippingService := ShippingServiceImpl{shippingRepository: shippingRepository}

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

	// Test case: Create shipping information for the order
	shipping := &domain.Shipping{
		OrderId:      order.OrderId,
		LogisticsCompany: "DHL",
		TrackingNumber: "123456789",
		EstimatedDelivery: time.Now().Add(24 * time.Hour),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	err := shippingService.CreateShipping(ctx, order.OrderId, shipping.LogisticsCompany, shipping.TrackingNumber, shipping.EstimatedDelivery)
	assert.NoError(t, err)

	// Assert
	foundShipping, err := shippingRepository.FindById(ctx, shipping.ShippingId)
	assert.NoError(t, err)
	assert.NotNil(t, foundShipping)
	assert.Equal(t, shipping.OrderId, foundShipping.OrderId)
	assert.Equal(t, shipping.LogisticsCompany, foundShipping.LogisticsCompany)

	// Test case: Update shipping information
	shipping.LogisticsCompany = "FedEx"
	err = shippingService.UpdateShipping(ctx, shipping.ShippingId, shipping.LogisticsCompany, shipping.TrackingNumber, shipping.EstimatedDelivery)
	assert.NoError(t, err)

	// Assert
	updatedShipping, err := shippingRepository.FindById(ctx, shipping.ShippingId)
	assert.NoError(t, err)
	assert.NotNil(t, updatedShipping)
	assert.Equal(t, shipping.LogisticsCompany, updatedShipping.LogisticsCompany)
}

// Additional test cases for shipping management can be added here.
// ...