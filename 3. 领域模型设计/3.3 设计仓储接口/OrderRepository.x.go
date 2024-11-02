package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// OrderRepository defines the interface for order data access.
type OrderRepository interface {
	Save(ctx context.Context, order *Order) error
	FindById(ctx context.Context, orderId uuid.UUID) (*Order, error)
	FindAllByUserId(ctx context.Context, userId int) ([]Order, error)
	Delete(ctx context.Context, orderId uuid.UUID) error
}

// InMemoryOrderRepository is an in-memory implementation of OrderRepository.
type InMemoryOrderRepository struct {
	orders map[uuid.UUID]*Order
}

// Save saves an order to the in-memory store.
func (repo *InMemoryOrderRepository) Save(ctx context.Context, order *Order) error {
	repo.orders[order.OrderId] = order
	return nil
}

// FindById finds an order by ID.
func (repo *InMemoryOrderRepository) FindById(ctx context.Context, orderId uuid.UUID) (*Order, error) {
	order, exists := repo.orders[orderId]
	if !exists {
		return nil, fmt.Errorf("order not found with ID: %v", orderId)
	}
	return order, nil
}

// FindAllByUserId finds all orders for a user.
func (repo *InMemoryOrderRepository) FindAllByUserId(ctx context.Context, userId int) ([]Order, error) {
	orders := make([]Order, 0)
	for _, order := range repo.orders {
		if order.UserId == userId {
			orders = append(orders, *order)
		}
	}
	return orders, nil
}

// Delete deletes an order from the in-memory store.
func (repo *InMemoryOrderRepository) Delete(ctx context.Context, orderId uuid.UUID) error {
	if _, exists := repo.orders[orderId]; !exists {
		return fmt.Errorf("order not found with ID: %v", orderId)
	}
	delete(repo.orders, orderId)
	return nil
}
