package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"yourdomain/domain" // Replace with your actual package path
)

// GetProduct retrieves a product by ID.
func GetProduct(c *gin.Context) {
	productId := c.Param("productId")
	product, err := domain.GetProductById(productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product.
func CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := domain.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// UpdateProduct updates an existing product.
func UpdateProduct(c *gin.Context) {
	productId := c.Param("productId")
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := domain.UpdateProduct(productId, &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// DeleteProduct deletes a product.
func DeleteProduct(c *gin.Context) {
	productId := c.Param("productId")
	if err := domain.DeleteProduct(productId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}