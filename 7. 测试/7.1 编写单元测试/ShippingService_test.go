package domain

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"yourdomain/domain" // Replace with your actual package path
)

func TestCreateShipping(t *testing.T) {
	repo := &InMemoryShippingRepository{
		shipments: make(map[uuid.UUID]*domain.Shipping),
	}

	shippingService := ShippingServiceImpl{shippingRepository: repo}

	ctx := context.Background()

	// Test case 1: Create shipping information for a new order
	orderId := uuid.New()
	shipping := &domain.Shipping{
		OrderId:      orderId,
		LogisticsCompany: "DHL",
		TrackingNumber: "123456789",
		EstimatedDelivery: time.Now().Add(24 * time.Hour),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	err := shippingService.CreateShipping(ctx, orderId, shipping.LogisticsCompany, shipping.TrackingNumber, shipping.EstimatedDelivery)
	assert.NoError(t, err)
	assert.NotNil(t, repo.shipments[shipping.ShippingId])
	assert.Equal(t, shipping.OrderId, repo.shipments[shipping.ShippingId].OrderId)
	assert.Equal(t, shipping.LogisticsCompany, repo.shipments[shipping.ShippingId].LogisticsCompany)

	// Test case 2: Create shipping information for an order that already has shipping information
	err = shippingService.CreateShipping(ctx, orderId, shipping.LogisticsCompany, shipping.TrackingNumber, shipping.EstimatedDelivery)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "shipping already exists for this order")
}

func TestGetShipping(t *testing.T) {
	repo := &InMemoryShippingRepository{
		shipments: make(map[uuid.UUID]*domain.Shipping),
	}

	shippingService := ShippingServiceImpl{shippingRepository: repo}

	ctx := context.Background()

	// Test case 1: Get an existing shipping record by ID
	shipping := &domain.Shipping{
		ShippingId:   uuid.New(),
		OrderId:      uuid.New(),
		LogisticsCompany: "DHL",
		TrackingNumber: "123456789",
		EstimatedDelivery: time.Now().Add(24 * time.Hour),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	repo.shipments[shipping.ShippingId] = shipping

	gotShipping, err := shippingService.GetShipping(ctx, shipping.ShippingId)
	assert.NoError(t, err)
	assert.NotNil(t, gotShipping)
	assert.Equal(t, shipping.OrderId, gotShipping.OrderId)
	assert.Equal(t, shipping.LogisticsCompany, gotShipping.LogisticsCompany)

	// Test case 2: Get a non-existent shipping record
	_, err = shippingService.GetShipping(ctx, uuid.UUID{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "shipping not found with ID")
}

// Additional tests for UpdateShipping and ListShipments can be added here.
// ...