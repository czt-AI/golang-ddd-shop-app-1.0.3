package domain

import (
	"time"
)

// Order defines the structure for an order entity.
type Order struct {
	OrderId    int
	UserId     int
	OrderItems []OrderItem
	TotalPrice decimal.Decimal
	OrderStatus OrderStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// NewOrder creates a new Order instance.
func NewOrder(orderId, userId int, orderItems []OrderItem, totalPrice decimal.Decimal, orderStatus OrderStatus, createdAt, updatedAt time.Time) *Order {
	return &Order{
		OrderId:     orderId,
		UserId:      userId,
		OrderItems:  orderItems,
		TotalPrice:  totalPrice,
		OrderStatus: orderStatus,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

// OrderItem defines the structure for an order item.
type OrderItem struct {
	ProductId int
	Quantity  int
	Price     decimal.Decimal
}

// NewOrderItem creates a new OrderItem instance.
func NewOrderItem(productId, quantity int, price decimal.Decimal) *OrderItem {
	return &OrderItem{
		ProductId: productId,
		Quantity:  quantity,
		Price:     price,
	}
}

// OrderStatus defines the order status.
type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusPaid       OrderStatus = "paid"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

// OrderDTO defines the data transfer object for an order.
type OrderDTO struct {
	OrderId    int
	UserId     int
	OrderItems []OrderItemDTO
	TotalPrice float64
	OrderStatus OrderStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// ToDTO converts an Order to an OrderDTO.
func (o *Order) ToDTO() *OrderDTO {
	return &OrderDTO{
		OrderId:    o.OrderId,
		UserId:     o.UserId,
		OrderItems: make([]OrderItemDTO, len(o.OrderItems)),
		TotalPrice: o.TotalPrice.Float64(),
		OrderStatus: o.OrderStatus,
		CreatedAt:   o.CreatedAt,
		UpdatedAt:   o.UpdatedAt,
	}
}

// OrderItemDTO defines the data transfer object for an order item.
type OrderItemDTO struct {
	ProductId int
	Quantity  int
	Price     float64
}
