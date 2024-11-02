package repository

import (
	"context"
	"yourdomain/domain" // Replace with your actual package path
)

// CategoryRepository defines the interface for category data access.
type CategoryRepository interface {
	Save(ctx context.Context, category *domain.Category) error
	FindById(ctx context.Context, categoryId int) (*domain.Category, error)
	FindAll(ctx context.Context) ([]domain.Category, error)
	Delete(ctx context.Context, categoryId int) error
}

// InMemoryCategoryRepository is an in-memory implementation of CategoryRepository.
type InMemoryCategoryRepository struct {
	categories map[int]*domain.Category
}

// Save saves a category to the in-memory store.
func (repo *InMemoryCategoryRepository) Save(ctx context.Context, category *domain.Category) error {
	repo.categories[category.CategoryId] = category
	return nil
}

// FindById finds a category by ID.
func (repo *InMemoryCategoryRepository) FindById(ctx context.Context, categoryId int) (*domain.Category, error) {
	category, exists := repo.categories[categoryId]
	if !exists {
		return nil, fmt.Errorf("category not found with ID: %d", categoryId)
	}
	return category, nil
}

// FindAll finds all categories.
func (repo *InMemoryCategoryRepository) FindAll(ctx context.Context) ([]domain.Category, error) {
	categories := make([]domain.Category, 0, len(repo.categories))
	for _, category := range repo.categories {
		categories = append(categories, *category)
	}
	return categories, nil
}

// Delete deletes a category from the in-memory store.
func (repo *InMemoryCategoryRepository) Delete(ctx context.Context, categoryId int) error {
	if _, exists := repo.categories[categoryId]; !exists {
		return fmt.Errorf("category not found with ID: %d", categoryId)
	}
	delete(repo.categories, categoryId)
	return nil
}
