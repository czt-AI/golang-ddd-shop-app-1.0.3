package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// ReviewRepository defines the interface for review data access.
type ReviewRepository interface {
	Save(ctx context.Context, review *Review) error
	FindById(ctx context.Context, reviewId uuid.UUID) (*Review, error)
	FindByProductId(ctx context.Context, productId int) ([]Review, error)
	Delete(ctx context.Context, reviewId uuid.UUID) error
}

// InMemoryReviewRepository is an in-memory implementation of ReviewRepository.
type InMemoryReviewRepository struct {
	reviews map[uuid.UUID]*Review
}

// Save saves a review to the in-memory store.
func (repo *InMemoryReviewRepository) Save(ctx context.Context, review *Review) error {
	repo.reviews[review.ReviewId] = review
	return nil
}

// FindById finds a review by ID.
func (repo *InMemoryReviewRepository) FindById(ctx context.Context, reviewId uuid.UUID) (*Review, error) {
	review, exists := repo.reviews[reviewId]
	if !exists {
		return nil, fmt.Errorf("review not found with ID: %v", reviewId)
	}
	return review, nil
}

// FindByProductId finds reviews by product ID.
func (repo *InMemoryReviewRepository) FindByProductId(ctx context.Context, productId int) ([]Review, error) {
	reviews := make([]Review, 0)
	for _, review := range repo.reviews {
		if review.ProductId == productId {
			reviews = append(reviews, *review)
		}
	}
	return reviews, nil
}

// Delete deletes a review from the in-memory store.
func (repo *InMemoryReviewRepository) Delete(ctx context.Context, reviewId uuid.UUID) error {
	if _, exists := repo.reviews[reviewId]; !exists {
		return fmt.Errorf("review not found with ID: %v", reviewId)
	}
	delete(repo.reviews, reviewId)
	return nil
}
