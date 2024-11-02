package domain

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"yourdomain/domain" // Replace with your actual package path
)

func TestCreateProduct(t *testing.T) {
	repo := &InMemoryProductRepository{
		products: make(map[int]*domain.Product),
	}

	productService := ProductService{productRepository: repo}

	ctx := context.Background()

	// Test case 1: Create a new product
	product := &domain.Product{
		ProductName:  "Test Product",
		Description:  "A test product",
		Price:        decimal.Decimal{C: 100, D: 0},
		Stock:        10,
		CategoryId:   1,
		Images:       []string{"image1.jpg", "image2.jpg"},
	}

	err := productService.CreateProduct(ctx, product.ProductName, product.Description, product.Price, product.Stock, product.CategoryId, product.Images)
	assert.NoError(t, err)
	assert.NotNil(t, repo.products[product.ProductId])
	assert.Equal(t, product.ProductName, repo.products[product.ProductId].ProductName)
	assert.Equal(t, product.Price, repo.products[product.ProductId].Price)

	// Test case 2: Create a product with a name that already exists
	product2 := &domain.Product{
		ProductName:  "Test Product",
		Description:  "Another test product",
		Price:        decimal.Decimal{C: 200, D: 0},
		Stock:        20,
		CategoryId:   2,
		Images:       []string{"image3.jpg"},
	}

	err = productService.CreateProduct(ctx, product2.ProductName, product2.Description, product2.Price, product2.Stock, product2.CategoryId, product2.Images)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "product name already exists")
}

func TestGetProduct(t *testing.T) {
	repo := &InMemoryProductRepository{
		products: make(map[int]*domain.Product),
	}

	productService := ProductService{productRepository: repo}

	ctx := context.Background()

	// Test case 1: Get a product by ID
	product := &domain.Product{
		ProductName:  "Test Product",
		Description:  "A test product",
		Price:        decimal.Decimal{C: 100, D: 0},
		Stock:        10,
		CategoryId:   1,
		Images:       []string{"image1.jpg", "image2.jpg"},
	}

	repo.products[product.ProductId] = product

	gotProduct, err := productService.GetProduct(ctx, product.ProductId)
	assert.NoError(t, err)
	assert.NotNil(t, gotProduct)
	assert.Equal(t, product.ProductName, gotProduct.ProductName)
	assert.Equal(t, product.Price, gotProduct.Price)

	// Test case 2: Get a non-existent product
	_, err = productService.GetProduct(ctx, 999)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "product not found with ID")
}

// Additional tests for UpdateProduct, DeleteProduct, ListProducts, and ListProductsByCategory can be added here.
// ...