package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// OrderService defines the service interface for order management.
type OrderService interface {
	CreateOrder(ctx context.Context, userId int, orderItems []OrderItem) (*Order, error)
	GetOrder(ctx context.Context, orderId uuid.UUID) (*Order, error)
	UpdateOrder(ctx context.Context, orderId uuid.UUID, status OrderStatus) error
	ListOrders(ctx context.Context, userId int) ([]Order, error)
	CancelOrder(ctx context.Context, orderId uuid.UUID) error
}

// OrderServiceImpl implements the OrderService interface.
type OrderServiceImpl struct {
	orderRepository OrderRepository
}

// CreateOrder creates a new order.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, userId int, orderItems []OrderItem) (*Order, error) {
	if len(orderItems) == 0 {
		return nil, errors.New("order items cannot be empty")
	}

	order := &Order{
		UserId:     userId,
		TotalPrice: calculateTotalPrice(orderItems),
		OrderStatus: OrderStatusPending,
		OrderItems:  orderItems,
	}
	return s.orderRepository.Save(ctx, order)
}

// GetOrder retrieves an order by ID.
func (s *OrderServiceImpl) GetOrder(ctx context.Context, orderId uuid.UUID) (*Order, error) {
	return s.orderRepository.FindById(ctx, orderId)
}

// UpdateOrder updates an order status.
func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, orderId uuid.UUID, status OrderStatus) error {
	order, err := s.orderRepository.FindById(ctx, orderId)
	if err != nil {
		return err
	}
	order.OrderStatus = status
	return s.orderRepository.Save(ctx, order)
}

// ListOrders lists all orders for a user.
func (s *OrderServiceImpl) ListOrders(ctx context.Context, userId int) ([]Order, error) {
	return s.orderRepository.FindAllByUserId(ctx, userId)
}

// CancelOrder cancels an order.
func (s *OrderServiceImpl) CancelOrder(ctx context.Context, orderId uuid.UUID) error {
	order, err := s.orderRepository.FindById(ctx, orderId)
	if err != nil {
		return err
	}
	if order.OrderStatus != OrderStatusPending {
		return errors.New("order cannot be cancelled in this status")
	}
	order.OrderStatus = OrderStatusCancelled
	return s.orderRepository.Save(ctx, order)
}

func calculateTotalPrice(orderItems []OrderItem) decimal.Decimal {
	total := decimal.Decimal{}
	for _, item := range orderItems {
		total = total.Add(item.Price.Mul(decimal.Decimal(item.Quantity)))
	}
	return total
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
