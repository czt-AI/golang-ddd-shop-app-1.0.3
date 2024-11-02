package domain

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"yourdomain/domain" // Replace with your actual package path
)

func TestCreateReview(t *testing.T) {
	repo := &InMemoryReviewRepository{
		reviews: make(map[uuid.UUID]*domain.Review),
	}

	reviewService := ReviewServiceImpl{reviewRepository: repo}

	ctx := context.Background()

	// Test case 1: Create a new review for a product
	review := &domain.Review{
		ProductId:   1,
		UserId:      1,
		Content:     "Great product!",
		Rating:      5,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	err := reviewService.CreateReview(ctx, review.UserId, review.ProductId, review.Content, review.Rating)
	assert.NoError(t, err)
	assert.NotNil(t, repo.reviews[review.ReviewId])
	assert.Equal(t, review.Content, repo.reviews[review.ReviewId].Content)
	assert.Equal(t, review.Rating, repo.reviews[review.ReviewId].Rating)

	// Test case 2: Create a review for a product that has already been reviewed by the same user
	err = reviewService.CreateReview(ctx, review.UserId, review.ProductId, review.Content, review.Rating)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user has already reviewed this product")
}

func TestGetReview(t *testing.T) {
	repo := &InMemoryReviewRepository{
		reviews: make(map[uuid.UUID]*domain.Review),
	}

	reviewService := ReviewServiceImpl{reviewRepository: repo}

	ctx := context.Background()

	// Test case 1: Get an existing review by ID
	review := &domain.Review{
		ProductId:   1,
		UserId:      1,
		Content:     "Great product!",
		Rating:      5,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	repo.reviews[review.ReviewId] = review

	gotReview, err := reviewService.GetReview(ctx, review.ReviewId)
	assert.NoError(t, err)
	assert.NotNil(t, gotReview)
	assert.Equal(t, review.Content, gotReview.Content)
	assert.Equal(t, review.Rating, gotReview.Rating)

	// Test case 2: Get a non-existent review
	_, err = reviewService.GetReview(ctx, uuid.UUID{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "review not found with ID")
}

// Additional tests for UpdateReview, DeleteReview, and ListReviews can be added here.
// ...