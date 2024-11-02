package domain

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"yourdomain/domain" // Replace with your actual package path
)

func TestCreateCategory(t *testing.T) {
	repo := &InMemoryCategoryRepository{
		categories: make(map[int]*domain.Category),
	}

	categoryService := CategoryServiceImpl{categoryRepository: repo}

	ctx := context.Background()

	// Test case 1: Create a new category
	category := &domain.Category{
		CategoryName: "Test Category",
	}

	err := categoryService.CreateCategory(ctx, category.CategoryName)
	assert.NoError(t, err)
	assert.NotNil(t, repo.categories[category.CategoryId])
	assert.Equal(t, category.CategoryName, repo.categories[category.CategoryId].CategoryName)

	// Test case 2: Create a category with a name that already exists
	err = categoryService.CreateCategory(ctx, category.CategoryName)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "category name already exists")
}

func TestGetCategory(t *testing.T) {
	repo := &InMemoryCategoryRepository{
		categories: make(map[int]*domain.Category),
	}

	categoryService := CategoryServiceImpl{categoryRepository: repo}

	ctx := context.Background()

	// Test case 1: Get a category by ID
	category := &domain.Category{
		CategoryName: "Test Category",
	}

	repo.categories[category.CategoryId] = category

	gotCategory, err := categoryService.GetCategory(ctx, category.CategoryId)
	assert.NoError(t, err)
	assert.NotNil(t, gotCategory)
	assert.Equal(t, category.CategoryName, gotCategory.CategoryName)

	// Test case 2: Get a non-existent category
	_, err = categoryService.GetCategory(ctx, 999)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "category not found with ID")
}

// Additional tests for UpdateCategory, DeleteCategory, and ListCategories can be added here.
// ...