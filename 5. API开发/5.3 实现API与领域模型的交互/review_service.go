package api

import (
	"context"
	"yourdomain/domain" // Replace with your actual package path
)

type reviewServiceServer struct {
	reviewRepository domain.ReviewRepository
}

func (s *reviewServiceServer) GetReview(ctx context.Context, req *ReviewRequest) (*ReviewResponse, error) {
	review, err := s.reviewRepository.FindById(ctx, req.ReviewId)
	if err != nil {
		return nil, err
	}
	return &ReviewResponse{
		ReviewId:    review.ReviewId,
		ProductId:   review.ProductId,
		UserId:      review.UserId,
		Content:     review.Content,
		Rating:      review.Rating,
		CreatedAt:   review.CreatedAt,
		UpdatedAt:   review.UpdatedAt,
	}, nil
}

func (s *reviewServiceServer) CreateReview(ctx context.Context, req *ReviewRequest) (*ReviewResponse, error) {
	review := &domain.Review{
		ProductId:   req.ProductId,
		UserId:      req.UserId,
		Content:     req.Content,
		Rating:      req.Rating,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Time{},
	}
	if err := s.reviewRepository.Save(ctx, review); err != nil {
		return nil, err
	}
	return &ReviewResponse{
		ReviewId:    review.ReviewId,
		ProductId:   review.ProductId,
		UserId:      review.UserId,
		Content:     review.Content,
		Rating:      review.Rating,
		CreatedAt:   review.CreatedAt,
		UpdatedAt:   review.UpdatedAt,
	}, nil
}

// ReviewRequest and ReviewResponse are placeholders for the actual request and response types.
type ReviewRequest struct {
	ReviewId    int
	ProductId   int
	UserId      int
	Content     string
	Rating      float32
}

type ReviewResponse struct {
	ReviewId    int
	ProductId   int
	UserId      int
	Content     string
	Rating      float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
