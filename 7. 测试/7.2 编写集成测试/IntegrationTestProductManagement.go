package domain

import (
	"context"
	"testing"

	"yourdomain/domain" // Replace with your actual package path
	"github.com/stretchr/testify/assert"
)

func TestProductManagement(t *testing.T) {
	// Setup
	productRepository := &InMemoryProductRepository{
		products: make(map[int]*domain.Product),
	}

	productService := ProductService{productRepository: productRepository}

	ctx := context.Background()

	// Test case: Create a new product
	product := &domain.Product{
		ProductName:  "Laptop",
		Description:  "A portable computer",
		Price:        decimal.Decimal{C: 1000, D: 0},
		Stock:        10,
		CategoryId:   1,
		Images:       []string{"laptop.jpg"},
	}

	err := productService.CreateProduct(ctx, product.ProductName, product.Description, product.Price, product.Stock, product.CategoryId, product.Images)
	assert.NoError(t, err)

	// Assert
	foundProduct, err := productRepository.FindById(ctx, product.ProductId)
	assert.NoError(t, err)
	assert.NotNil(t, foundProduct)
	assert.Equal(t, product.ProductName, foundProduct.ProductName)
	assert.Equal(t, product.Price, foundProduct.Price)

	// Test case: Update a product
	foundProduct.Price = decimal.Decimal{C: 1100, D: 0}
	err = productService.UpdateProduct(ctx, product.ProductId, foundProduct.ProductName, foundProduct.Description, foundProduct.Price, foundProduct.Stock, foundProduct.CategoryId, foundProduct.Images)
	assert.NoError(t, err)

	// Assert
	updatedProduct, err := productRepository.FindById(ctx, product.ProductId)
	assert.NoError(t, err)
	assert.NotNil(t, updatedProduct)
	assert.Equal(t, foundProduct.Price, updatedProduct.Price)

	// Test case: Delete a product
	err = productService.DeleteProduct(ctx, product.ProductId)
	assert.NoError(t, err)

	// Assert
	_, err = productRepository.FindById(ctx, product.ProductId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "product not found with ID")
}

// Additional test cases for product management can be added here.
// ...