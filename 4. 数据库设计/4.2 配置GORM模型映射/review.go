package domain

import (
	"gorm.io/gorm"
	"time"
)

type Review struct {
	gorm.Model
	ProductID   int
	Content     string
	Rating      float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ReviewDTO defines the data transfer object for a review.
type ReviewDTO struct {
	ReviewId    int
	ProductId   int
	UserId      int
	Content     string
	Rating      float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ToDTO converts a Review to a ReviewDTO.
func (r *Review) ToDTO() *ReviewDTO {
	return &ReviewDTO{
		ReviewId:    r.ID,
		ProductId:   r.ProductID,
		UserId:      r.UserID,
		Content:     r.Content,
		Rating:      r.Rating,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}