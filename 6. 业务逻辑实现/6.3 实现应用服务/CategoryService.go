package service

import (
	"context"
	"errors"

	"yourdomain/domain" // Replace with your actual package path
)

// CategoryService defines the service interface for category management.
type CategoryService interface {
	CreateCategory(ctx context.Context, name string) (*domain.Category, error)
	GetCategory(ctx context.Context, categoryId int) (*domain.Category, error)
	UpdateCategory(ctx context.Context, categoryId int, name string) error
	DeleteCategory(ctx context.Context, categoryId int) error
	ListCategories(ctx context.Context) ([]domain.Category, error)
}

// CategoryServiceImpl implements the CategoryService interface.
type CategoryServiceImpl struct {
	categoryRepository domain.CategoryRepository
}

// CreateCategory creates a new category.
func (s *CategoryServiceImpl) CreateCategory(ctx context.Context, name string) (*domain.Category, error) {
	if err := s.categoryRepository.FindByName(ctx, name); err == nil {
		return nil, errors.New("category name already exists")
	}
	category := &domain.Category{
		CategoryName: name,
	}
	return s.categoryRepository.Save(ctx, category)
}

// GetCategory retrieves a category by ID.
func (s *CategoryServiceImpl) GetCategory(ctx context.Context, categoryId int) (*domain.Category, error) {
	return s.categoryRepository.FindById(ctx, categoryId)
}

// UpdateCategory updates a category.
func (s *CategoryServiceImpl) UpdateCategory(ctx context.Context, categoryId int, name string) error {
	category, err := s.categoryRepository.FindById(ctx, categoryId)
	if err != nil {
		return err
	}
	category.CategoryName = name
	return s.categoryRepository.Save(ctx, category)
}

// DeleteCategory deletes a category.
func (s *CategoryServiceImpl) DeleteCategory(ctx context.Context, categoryId int) error {
	return s.categoryRepository.Delete(ctx, categoryId)
}

// ListCategories lists all categories.
func (s *CategoryServiceImpl) ListCategories(ctx context.Context) ([]domain.Category, error) {
	return s.categoryRepository.FindAll(ctx)
}
