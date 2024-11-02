package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"yourdomain/domain" // Replace with your actual package path
)

// GetReview retrieves a review by ID.
func GetReview(c *gin.Context) {
	reviewId := c.Param("reviewId")
	review, err := domain.GetReviewById(reviewId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, review)
}

// CreateReview creates a new review for a product.
func CreateReview(c *gin.Context) {
	var review domain.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := domain.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, review)
}

// UpdateReview updates an existing review.
func UpdateReview(c *gin.Context) {
	reviewId := c.Param("reviewId")
	var review domain.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := domain.UpdateReview(reviewId, &review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, review)
}

// DeleteReview deletes a review.
func DeleteReview(c *gin.Context) {
	reviewId := c.Param("reviewId")
	if err := domain.DeleteReview(reviewId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}