package domain

import (
	"context"
	"errors"
)

// ProductService defines the service interface for product management.
type ProductService interface {
	CreateProduct(ctx context.Context, name, description string, price decimal.Decimal, stock int, categoryId int, images []string) (*Product, error)
	GetProduct(ctx context.Context, productId int) (*Product, error)
	UpdateProduct(ctx context.Context, productId int, name, description string, price decimal.Decimal, stock int, categoryId int, images []string) error
	DeleteProduct(ctx context.Context, productId int) error
	ListProducts(ctx context.Context) ([]Product, error)
	ListProductsByCategory(ctx context.Context, categoryId int) ([]Product, error)
}

// ProductServiceImpl implements the ProductService interface.
type ProductServiceImpl struct {
	productRepository ProductRepository
}

// CreateProduct creates a new product.
func (s *ProductServiceImpl) CreateProduct(ctx context.Context, name, description string, price decimal.Decimal, stock int, categoryId int, images []string) (*Product, error) {
	if err := s.productRepository.FindByName(ctx, name); err == nil {
		return nil, errors.New("product name already exists")
	}
	product := &Product{
		ProductName:  name,
		Description:  description,
		Price:        price,
		Stock:        stock,
		CategoryId:   categoryId,
		Images:       images,
	}
	return s.productRepository.Save(ctx, product)
}

// GetProduct retrieves a product by ID.
func (s *ProductServiceImpl) GetProduct(ctx context.Context, productId int) (*Product, error) {
	return s.productRepository.FindById(ctx, productId)
}

// UpdateProduct updates a product.
func (s *ProductServiceImpl) UpdateProduct(ctx context.Context, productId int, name, description string, price decimal.Decimal, stock int, categoryId int, images []string) error {
	product, err := s.productRepository.FindById(ctx, productId)
	if err != nil {
		return err
	}
	product.ProductName = name
	product.Description = description
	product.Price = price
	product.Stock = stock
	product.CategoryId = categoryId
	product.Images = images
	return s.productRepository.Save(ctx, product)
}

// DeleteProduct deletes a product.
func (s *ProductServiceImpl) DeleteProduct(ctx context.Context, productId int) error {
	return s.productRepository.Delete(ctx, productId)
}

// ListProducts lists all products.
func (s *ProductServiceImpl) ListProducts(ctx context.Context) ([]Product, error) {
	return s.productRepository.FindAll(ctx)
}

// ListProductsByCategory lists products by category ID.
func (s *ProductServiceImpl) ListProductsByCategory(ctx context.Context, categoryId int) ([]Product, error) {
	return s.productRepository.FindByCategory(ctx, categoryId)
}
