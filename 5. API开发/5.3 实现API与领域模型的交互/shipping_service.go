package api

import (
	"context"
	"yourdomain/domain" // Replace with your actual package path
)

type shippingServiceServer struct {
	shippingRepository domain.ShippingRepository
}

func (s *shippingServiceServer) GetShipping(ctx context.Context, req *ShippingRequest) (*ShippingResponse, error) {
	shipping, err := s.shippingRepository.FindById(ctx, req.ShippingId)
	if err != nil {
		return nil, err
	}
	return &ShippingResponse{
		ShippingId:   shipping.ShippingId,
		OrderId:      shipping.OrderId,
		LogisticsCompany: shipping.LogisticsCompany,
		TrackingNumber: shipping.TrackingNumber,
		EstimatedDelivery: shipping.EstimatedDelivery,
		CreatedAt:   shipping.CreatedAt,
		UpdatedAt:   shipping.UpdatedAt,
	}, nil
}

func (s *shippingServiceServer) CreateShipping(ctx context.Context, req *ShippingRequest) (*ShippingResponse, error) {
	shipping := &domain.Shipping{
		OrderId:      req.OrderId,
		LogisticsCompany: req.LogisticsCompany,
		TrackingNumber: req.TrackingNumber,
		EstimatedDelivery: req.EstimatedDelivery,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	if err := s.shippingRepository.Save(ctx, shipping); err != nil {
		return nil, err
	}
	return &ShippingResponse{
		ShippingId:   shipping.ShippingId,
		OrderId:      shipping.OrderId,
		LogisticsCompany: shipping.LogisticsCompany,
		TrackingNumber: shipping.TrackingNumber,
		EstimatedDelivery: shipping.EstimatedDelivery,
		CreatedAt:   shipping.CreatedAt,
		UpdatedAt:   shipping.UpdatedAt,
	}, nil
}

// ShippingRequest and ShippingResponse are placeholders for the actual request and response types.
type ShippingRequest struct {
	ShippingId   int
	OrderId      int
	LogisticsCompany string
	TrackingNumber string
	EstimatedDelivery time.Time
}

type ShippingResponse struct {
	ShippingId   int
	OrderId      int
	LogisticsCompany string
	TrackingNumber string
	EstimatedDelivery time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}