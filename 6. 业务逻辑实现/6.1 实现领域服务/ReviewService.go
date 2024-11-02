package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// ReviewService defines the service interface for review management.
type ReviewService interface {
	CreateReview(ctx context.Context, userId, productId int, content string, rating float32) (*Review, error)
	GetReview(ctx context.Context, reviewId uuid.UUID) (*Review, error)
	UpdateReview(ctx context.Context, reviewId uuid.UUID, content string, rating float32) error
	DeleteReview(ctx context.Context, reviewId uuid.UUID) error
	ListReviews(ctx context.Context, productId int) ([]Review, error)
}

// ReviewServiceImpl implements the ReviewService interface.
type ReviewServiceImpl struct {
	reviewRepository ReviewRepository
}

// CreateReview creates a new review for a product.
func (s *ReviewServiceImpl) CreateReview(ctx context.Context, userId, productId int, content string, rating float32) (*Review, error) {
	if err := s.reviewRepository.FindByProductIdAndUserId(ctx, productId, userId); err == nil {
		return nil, errors.New("user has already reviewed this product")
	}
	review := &Review{
		ProductId:   productId,
		UserId:      userId,
		Content:     content,
		Rating:      rating,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	return s.reviewRepository.Save(ctx, review)
}

// GetReview retrieves a review by ID.
func (s *ReviewServiceImpl) GetReview(ctx context.Context, reviewId uuid.UUID) (*Review, error) {
	return s.reviewRepository.FindById(ctx, reviewId)
}

// UpdateReview updates a review.
func (s *ReviewServiceImpl) UpdateReview(ctx context.Context, reviewId uuid.UUID, content string, rating float32) error {
	review, err := s.reviewRepository.FindById(ctx, reviewId)
	if err != nil {
		return err
	}
	review.Content = content
	review.Rating = rating
	return s.reviewRepository.Save(ctx, review)
}

// DeleteReview deletes a review.
func (s *ReviewServiceImpl) DeleteReview(ctx context.Context, reviewId uuid.UUID) error {
	return s.reviewRepository.Delete(ctx, reviewId)
}

// ListReviews retrieves all reviews for a product.
func (s *ReviewServiceImpl) ListReviews(ctx context.Context, productId int) ([]Review, error) {
	return s.reviewRepository.FindByProductId(ctx, productId)
}
