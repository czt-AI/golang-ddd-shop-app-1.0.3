package api

import (
	"context"
	"yourdomain/domain" // Replace with your actual package path
)

type productServiceServer struct {
	productRepository domain.ProductRepository
}

func (s *productServiceServer) GetProduct(ctx context.Context, req *ProductRequest) (*ProductResponse, error) {
	product, err := s.productRepository.FindById(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}
	return &ProductResponse{
		ProductId:    product.ProductId,
		ProductName:  product.ProductName,
		Description:  product.Description,
		Price:        product.Price,
		Stock:        product.Stock,
		CategoryId:   product.CategoryId,
		Images:       product.Images,
	}, nil
}

func (s *productServiceServer) CreateProduct(ctx context.Context, req *ProductRequest) (*ProductResponse, error) {
	product := &domain.Product{
		ProductName:  req.ProductName,
		Description:  req.Description,
		Price:        req.Price,
		Stock:        req.Stock,
		CategoryId:   req.CategoryId,
		Images:       req.Images,
	}
	if err := s.productRepository.Save(ctx, product); err != nil {
		return nil, err
	}
	return &ProductResponse{
		ProductId:    product.ProductId,
		ProductName:  product.ProductName,
		Description:  product.Description,
		Price:        product.Price,
		Stock:        product.Stock,
		CategoryId:   product.CategoryId,
		Images:       product.Images,
	}, nil
}

// ProductRequest and ProductResponse are placeholders for the actual request and response types.
type ProductRequest struct {
	ProductId    int
	ProductName  string
	Description  string
	Price        decimal.Decimal
	Stock        int
	CategoryId   int
	Images       []string
}

type ProductResponse struct {
	ProductId    int
	ProductName  string
	Description  string
	Price        decimal.Decimal
	Stock        int
	CategoryId   int
	Images       []string
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
