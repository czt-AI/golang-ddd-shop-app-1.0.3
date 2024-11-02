package api

import (
	"context"
	"yourdomain/domain" // Replace with your actual package path
)

type orderServiceServer struct {
	orderRepository domain.OrderRepository
}

func (s *orderServiceServer) GetOrder(ctx context.Context, req *OrderRequest) (*OrderResponse, error) {
	order, err := s.orderRepository.FindById(ctx, req.OrderId)
	if err != nil {
		return nil, err
	}
	return &OrderResponse{
		OrderId:       order.OrderId,
		UserId:        order.UserId,
		TotalPrice:    order.TotalPrice,
		OrderStatus:   order.OrderStatus,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
		OrderItems:    make([]OrderItemResponse, len(order.OrderItems)),
	}, nil
}

func (s *orderServiceServer) CreateOrder(ctx context.Context, req *OrderRequest) (*OrderResponse, error) {
	order := &domain.Order{
		UserId:        req.UserId,
		TotalPrice:    req.TotalPrice,
		OrderStatus:   req.OrderStatus,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Time{},
		OrderItems:    req.OrderItems,
	}
	if err := s.orderRepository.Save(ctx, order); err != nil {
		return nil, err
	}
	return &OrderResponse{
		OrderId:       order.OrderId,
		UserId:        order.UserId,
		TotalPrice:    order.TotalPrice,
		OrderStatus:   order.OrderStatus,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
		OrderItems:    make([]OrderItemResponse, len(order.OrderItems)),
	}, nil
}

// OrderRequest and OrderResponse are placeholders for the actual request and response types.
type OrderRequest struct {
	OrderId       uuid.UUID
	UserId        int
	TotalPrice    decimal.Decimal
	OrderStatus   OrderStatus
	OrderItems    []OrderItemRequest
}

type OrderResponse struct {
	OrderId       uuid.UUID
	UserId        int
	TotalPrice    decimal.Decimal
	OrderStatus   OrderStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
	OrderItems    []OrderItemResponse
}

type OrderItemRequest struct {
	ProductId    int
	Quantity     int
	Price        decimal.Decimal
}

type OrderItemResponse struct {
	ProductId    int
	Quantity     int
	Price        decimal.Decimal
}

type decimal struct {
	C int64
	D int64
}

func (d decimal) Float64() float64 {
	return float64(d.C) + float64(d.D)/float64(10^int64(len(d.D)))
}

func (d *decimal) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d.%d", d.C, d.D)), nil
}

func (d *decimal) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parts := strings.Split(s, ".")
	if len(parts) != 2 {
		return fmt.Errorf("invalid decimal format: %s", s)
	}
	d.C, _ = strconv.ParseInt(parts[0], 10, 64)
	d.D, _ = strconv.ParseInt(parts[1], 10, 64)
	return nil
}

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusPaid       OrderStatus = "paid"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)
