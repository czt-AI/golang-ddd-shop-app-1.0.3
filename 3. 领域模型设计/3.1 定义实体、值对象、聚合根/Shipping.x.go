package domain

// Shipping defines the structure for a shipping entity.
type Shipping struct {
	ShippingId   int
	OrderId      int
	LogisticsCompany string
	TrackingNumber string
	EstimatedDelivery time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewShipping creates a new Shipping instance.
func NewShipping(shippingId, orderId int, logisticsCompany, trackingNumber string, estimatedDelivery time.Time, createdAt, updatedAt time.Time) *Shipping {
	return &Shipping{
		ShippingId:   shippingId,
		OrderId:      orderId,
		LogisticsCompany: logisticsCompany,
		TrackingNumber: trackingNumber,
		EstimatedDelivery: estimatedDelivery,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

// ShippingDTO defines the data transfer object for a shipping.
type ShippingDTO struct {
	ShippingId   int
	OrderId      int
	LogisticsCompany string
	TrackingNumber string
	EstimatedDelivery time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ToDTO converts a Shipping to a ShippingDTO.
func (s *Shipping) ToDTO() *ShippingDTO {
	return &ShippingDTO{
		ShippingId:   s.ShippingId,
		OrderId:      s.OrderId,
		LogisticsCompany: s.LogisticsCompany,
		TrackingNumber: s.TrackingNumber,
		EstimatedDelivery: s.EstimatedDelivery,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}