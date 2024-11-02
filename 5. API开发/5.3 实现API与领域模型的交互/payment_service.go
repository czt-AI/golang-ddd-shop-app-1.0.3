package api

import (
	"context"
	"yourdomain/domain" // Replace with your actual package path
)

type paymentServiceServer struct {
	paymentRepository domain.PaymentRepository
}

func (s *paymentServiceServer) GetPayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	payment, err := s.paymentRepository.FindById(ctx, req.PaymentId)
	if err != nil {
		return nil, err
	}
	return &PaymentResponse{
		PaymentId:   payment.PaymentId,
		OrderId:     payment.OrderId,
		Amount:      payment.Amount,
		PaymentMethod: payment.PaymentMethod,
		PaymentStatus: payment.PaymentStatus,
		CreatedAt:   payment.CreatedAt,
		UpdatedAt:   payment.UpdatedAt,
	}, nil
}

func (s *paymentServiceServer) CreatePayment(ctx context.Context, req *PaymentRequest) (*PaymentResponse, error) {
	payment := &domain.Payment{
		OrderId:     req.OrderId,
		Amount:      req.Amount,
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: req.PaymentStatus,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	if err := s.paymentRepository.Save(ctx, payment); err != nil {
		return nil, err
	}
	return &PaymentResponse{
		PaymentId:   payment.PaymentId,
		OrderId:     payment.OrderId,
		Amount:      payment.Amount,
		PaymentMethod: payment.PaymentMethod,
		PaymentStatus: payment.PaymentStatus,
		CreatedAt:   payment.CreatedAt,
		UpdatedAt:   payment.UpdatedAt,
	}, nil
}

// PaymentRequest and PaymentResponse are placeholders for the actual request and response types.
type PaymentRequest struct {
	PaymentId   int
	OrderId     int
	Amount      decimal.Decimal
	PaymentMethod string
	PaymentStatus PaymentStatus
}

type PaymentResponse struct {
	PaymentId   int
	OrderId     int
	Amount      decimal.Decimal
	PaymentMethod string
	PaymentStatus PaymentStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
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

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusSuccess  PaymentStatus = "success"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusCancelled PaymentStatus = "cancelled"
)
