package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"yourdomain/domain" // Replace with your actual package path
)

// GetShipping retrieves shipping information by ID.
func GetShipping(c *gin.Context) {
	shippingId := c.Param("shippingId")
	shipping, err := domain.GetShippingById(shippingId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shipping)
}

// CreateShipping creates shipping information for an order.
func CreateShipping(c *gin.Context) {
	var shipping domain.Shipping
	if err := c.ShouldBindJSON(&shipping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := domain.CreateShipping(&shipping); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, shipping)
}

// UpdateShipping updates shipping information.
func UpdateShipping(c *gin.Context) {
	shippingId := c.Param("shippingId")
	var shipping domain.Shipping
	if err := c.ShouldBindJSON(&shipping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := domain.UpdateShipping(shippingId, &shipping); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shipping)
}

// DeleteShipping deletes shipping information.
func DeleteShipping(c *gin.Context) {
	shippingId := c.Param("shippingId")
	if err := domain.DeleteShipping(shippingId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}