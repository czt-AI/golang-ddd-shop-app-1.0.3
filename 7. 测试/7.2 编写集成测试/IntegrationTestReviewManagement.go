package domain

import (
	"context"
	"testing"
	"time"

	"yourdomain/domain" // Replace with your actual package path
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
)

func TestReviewManagement(t *testing.T) {
	// Setup
	userRepository := &InMemoryUserRepository{
		users: make(map[int]*domain.User),
	}

	productRepository := &InMemoryProductRepository{
		products: make(map[int]*domain.Product),
	}

	reviewRepository := &InMemoryReviewRepository{
		reviews: make(map[uuid.UUID]*domain.Review),
	}

	userService := UserServiceImpl{userRepository: userRepository}
	productService := ProductService{productRepository: productRepository}
	reviewService := ReviewServiceImpl{reviewRepository: reviewRepository}

	ctx := context.Background()

	// Test case: Create a new review for a product
	user := &domain.User{
		Username: "testuser",
	}
	userRepository.Save(ctx, user)

	product := &domain.Product{
		ProductName:  "Laptop",
	}
	productRepository.Save(ctx, product)

	review := &domain.Review{
		ProductId:   product.ProductId,
		UserId:      user.UserId,
		Content:     "Great product!",
		Rating:      5,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}

	err := reviewService.CreateReview(ctx, user.UserId, product.ProductId, review.Content, review.Rating)
	assert.NoError(t, err)

	// Assert
	foundReview, err := reviewRepository.FindById(ctx, review.ReviewId)
	assert.NoError(t, err)
	assert.NotNil(t, foundReview)
	assert.Equal(t, review.Content, foundReview.Content)
	assert.Equal(t, review.Rating, foundReview.Rating)

	// Test case: Update a review
	review.Content = "Excellent product!"
	err = reviewService.UpdateReview(ctx, review.ReviewId, review.Content, review.Rating)
	assert.NoError(t, err)

	// Assert
	updatedReview, err := reviewRepository.FindById(ctx, review.ReviewId)
	assert.NoError(t, err)
	assert.NotNil(t, updatedReview)
	assert.Equal(t, review.Content, updatedReview.Content)

	// Test case: Delete a review
	err = reviewService.DeleteReview(ctx, review.ReviewId)
	assert.NoError(t, err)

	// Assert
	_, err = reviewRepository.FindById(ctx, review.ReviewId)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "review not found with ID")
}

// Additional test cases for review management can be added here.
// ...