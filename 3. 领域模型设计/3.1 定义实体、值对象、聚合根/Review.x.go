package domain

// Review defines the structure for a review entity.
type Review struct {
	ReviewId    int
	UserId      int
	ProductId   int
	Content     string
	Rating      float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewReview creates a new Review instance.
func NewReview(reviewId, userId, productId int, content string, rating float32, createdAt, updatedAt time.Time) *Review {
	return &Review{
		ReviewId:    reviewId,
		UserId:      userId,
		ProductId:   productId,
		Content:     content,
		Rating:      rating,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

// ReviewDTO defines the data transfer object for a review.
type ReviewDTO struct {
	ReviewId    int
	UserId      int
	ProductId   int
	Content     string
	Rating      float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ToDTO converts a Review to a ReviewDTO.
func (r *Review) ToDTO() *ReviewDTO {
	return &ReviewDTO{
		ReviewId:    r.ReviewId,
		UserId:      r.UserId,
		ProductId:   r.ProductId,
		Content:     r.Content,
		Rating:      r.Rating,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}
