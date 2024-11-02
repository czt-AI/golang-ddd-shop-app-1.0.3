package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// ShippingService defines the service interface for shipping management.
type ShippingService interface {
	CreateShipping(ctx context.Context, orderId uuid.UUID, logisticsCompany, trackingNumber string, estimatedDelivery time.Time) (*Shipping, error)
	GetShipping(ctx context.Context, shippingId uuid.UUID) (*Shipping, error)
	UpdateShipping(ctx context.Context, shippingId uuid.UUID, logisticsCompany, trackingNumber string, estimatedDelivery time.Time) error
	ListShipments(ctx context.Context, orderId uuid.UUID) ([]Shipping, error)
}

// ShippingServiceImpl implements the ShippingService interface.
type ShippingServiceImpl struct {
	shippingRepository ShippingRepository
}

// CreateShipping creates shipping information for an order.
func (s *ShippingServiceImpl) CreateShipping(ctx context.Context, orderId uuid.UUID, logisticsCompany, trackingNumber string, estimatedDelivery time.Time) (*Shipping, error) {
	if _, err := s.shippingRepository.FindByOrderId(ctx, orderId); err == nil {
		return nil, errors.New("shipping already exists for this order")
	}
	shipping := &Shipping{
		OrderId:      orderId,
		LogisticsCompany: logisticsCompany,
		TrackingNumber: trackingNumber,
		EstimatedDelivery: estimatedDelivery,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	return s.shippingRepository.Save(ctx, shipping)
}

// GetShipping retrieves shipping information by ID.
func (s *ShippingServiceImpl) GetShipping(ctx context.Context, shippingId uuid.UUID) (*Shipping, error) {
	return s.shippingRepository.FindById(ctx, shippingId)
}

// UpdateShipping updates shipping information.
func (s *ShippingServiceImpl) UpdateShipping(ctx context.Context, shippingId uuid.UUID, logisticsCompany, trackingNumber string, estimatedDelivery time.Time) error {
	shipping, err := s.shippingRepository.FindById(ctx, shippingId)
	if err != nil {
		return err
	}
	shipping.LogisticsCompany = logisticsCompany
	shipping.TrackingNumber = trackingNumber
	shipping.EstimatedDelivery = estimatedDelivery
	return s.shippingRepository.Save(ctx, shipping)
}

// ListShipments retrieves all shipments for an order.
func (s *ShippingServiceImpl) ListShipments(ctx context.Context, orderId uuid.UUID) ([]Shipping, error) {
	return s.shippingRepository.FindByOrderId(ctx, orderId)
}