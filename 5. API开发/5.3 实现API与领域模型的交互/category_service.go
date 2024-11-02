package api

import (
	"context"
	"yourdomain/domain" // Replace with your actual package path
)

type categoryServiceServer struct {
	categoryRepository domain.CategoryRepository
}

func (s *categoryServiceServer) GetCategory(ctx context.Context, req *CategoryRequest) (*CategoryResponse, error) {
	category, err := s.categoryRepository.FindById(ctx, req.CategoryId)
	if err != nil {
		return nil, err
	}
	return &CategoryResponse{
		CategoryId:   category.CategoryId,
		CategoryName: category.CategoryName,
	}, nil
}

func (s *categoryServiceServer) CreateCategory(ctx context.Context, req *CategoryRequest) (*CategoryResponse, error) {
	category := &domain.Category{
		CategoryName: req.CategoryName,
	}
	if err := s.categoryRepository.Save(ctx, category); err != nil {
		return nil, err
	}
	return &CategoryResponse{
		CategoryId:   category.CategoryId,
		CategoryName: category.CategoryName,
	}, nil
}

// CategoryRequest and CategoryResponse are placeholders for the actual request and response types.
type CategoryRequest struct {
	CategoryId int
}

type CategoryResponse struct {
	CategoryId   int
	CategoryName string
}
