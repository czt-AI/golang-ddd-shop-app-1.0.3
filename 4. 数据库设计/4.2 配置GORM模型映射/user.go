package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	PasswordHash string
	Email        string
	Phone        string
	Nickname     string
	AvatarUrl    string
	RegistrationTime time.Time
	LastLoginTime time.Time
}

shop-app/4. 数据库设计/4.2 配置GORM模型映射/category.go
package domain

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string
}

shop-app/4. 数据库设计/4.2 配置GORM模型映射/product.go
package domain

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName  string
	Description  string
	Price        decimal.Decimal
	Stock        int
	CategoryID   int
	Images       []string
	Reviews      []Review `gorm:"foreignKey:ProductID"`
}

shop-app/4. 数据库设计/4.2 配置GORM模型映射/order.go
package domain

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID     int
	TotalPrice decimal.Decimal
	OrderStatus OrderStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	ProductID int
	Quantity  int
	Price     decimal.Decimal
}

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusPaid       OrderStatus = "paid"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusCancelled  OrderStatus = "cancelled"
)

shop-app/4. 数据库设计/4.2 配置GORM模型映射/review.go
package domain

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	gorm.Model
	ProductID   int
	Content     string
	Rating      float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

shop-app/4. 数据库设计/4.2 配置GORM模型映射/payment.go
package domain

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	OrderID     int
	Amount      decimal.Decimal
	PaymentMethod string
	PaymentStatus PaymentStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSuccess  PaymentStatus = "success"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)

shop-app/4. 数据库设计/4.2 配置GORM模型映射/shipping.go
package domain

import (
	"gorm.io/gorm"
	"time"
)

type Shipping struct {
	gorm.Model
	OrderID      int
	LogisticsCompany string
	TrackingNumber string
	EstimatedDelivery time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}