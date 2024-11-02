package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"yourdomain/domain" // Replace with your actual package path
)

// GetOrder retrieves an order by ID.
func GetOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	order, err := domain.GetOrderById(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

// CreateOrder creates a new order.
func CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := domain.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

// UpdateOrder updates an existing order.
func UpdateOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := domain.UpdateOrder(orderId, &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

// DeleteOrder deletes an order.
func DeleteOrder(c *gin.Context) {
	orderId := c.Param("orderId")
	if err := domain.DeleteOrder(orderId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}