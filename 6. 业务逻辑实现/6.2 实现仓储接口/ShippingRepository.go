package repository

import (
	"context"
	"yourdomain/domain" // Replace with your actual package path
	"github.com/google/uuid"
)

// ShippingRepository defines the interface for shipping data access.
type ShippingRepository interface {
	Save(ctx context.Context, shipping *domain.Shipping) error
	FindById(ctx context.Context, shippingId uuid.UUID) (*domain.Shipping, error)
	FindByOrderId(ctx context.Context, orderId uuid.UUID) (*domain.Shipping, error)
	Delete(ctx context.Context, shippingId uuid.UUID) error
}

// InMemoryShippingRepository is an in-memory implementation of ShippingRepository.
type InMemoryShippingRepository struct {
	shipments map[uuid.UUID]*domain.Shipping
}

// Save saves a shipping record to the in-memory store.
func (repo *InMemoryShippingRepository) Save(ctx context.Context, shipping *domain.Shipping) error {
	repo.shipments[shipping.ShippingId] = shipping
	return nil
}

// FindById finds a shipping record by ID.
func (repo *InMemoryShippingRepository) FindById(ctx context.Context, shippingId uuid.UUID) (*domain.Shipping, error) {
	shipping, exists := repo.shipments[shippingId]
	if !exists {
		return nil, fmt.Errorf("shipping not found with ID: %v", shippingId)
	}
	return shipping, nil
}

// FindByOrderId finds shipping records by order ID.
func (repo *InMemoryShippingRepository) FindByOrderId(ctx context.Context, orderId uuid.UUID) (*domain.Shipping, error) {
	for _, shipping := range repo.shipments {
		if shipping.OrderId == orderId {
			return shipping, nil
		}
	}
	return nil, fmt.Errorf("shipping not found with order ID: %v", orderId)
}

// Delete deletes a shipping record from the in-memory store.
func (repo *InMemoryShippingRepository) Delete(ctx context.Context, shippingId uuid.UUID) error {
	if _, exists := repo.shipments[shippingId]; !exists {
		return fmt.Errorf("shipping not found with ID: %v", shippingId)
	}
	delete(repo.shipments, shippingId)
	return nil
}