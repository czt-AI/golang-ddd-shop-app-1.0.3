package domain

import (
	"context"
	"testing"

	"yourdomain/domain" // Replace with your actual package path
	"github.com/stretchr/testify/assert"
)

func TestCategoryManagement(t *testing.T) {
	// Setup
	categoryRepository := &InMemoryCategoryRepository{
		categories: make(map[int]*domain.Category),
	}

	categoryService := CategoryServiceImpl{categoryRepository: categoryRepository}

	ctx := context.Background()

	// Test case: Create a new category
	category := &domain.Category{
		CategoryName: "Electronics",
	}

	err := categoryService.CreateCategory(ctx, category.CategoryName)
	assert.NoError(t, err)

	// Assert
	foundCategory, err := categoryRepository.FindById(ctx, category.CategoryId)
	assert.NoError(t, err)
	assert.NotNil(t, foundCategory)
	assert.Equal(t, category.CategoryName, foundCategory.CategoryName)

	// Test case: Update a category
	foundCategory.CategoryName = "Gadgets"
	err = categoryService.UpdateCategory(ctx, category.CategoryId, foundCategory.CategoryName)
	assert.NoError(t, err)

	// Assert
	updatedCategory, err := categoryRepository.FindById(ctx, category.CategoryId)
	assert.NoError(t, err)
	assert.NotNil(t, updatedCategory)
	assert.Equal(t, foundCategory.CategoryName, updatedCategory.CategoryName)

	// Test case: Delete a category
	err = categoryService.DeleteCategory(ctx, category.CategoryId)
	assert.NoError(t, err)

	// Assert
	_, err = categoryRepository.FindById(ctx, category.CategoryId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "category not found with ID")
}

// Additional test cases for category management can be added here.
// ...