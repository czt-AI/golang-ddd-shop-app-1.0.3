package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"yourdomain/domain" // Replace with your actual package path
)

// GetPayment retrieves a payment by ID.
func GetPayment(c *gin.Context) {
	paymentId := c.Param("paymentId")
	payment, err := domain.GetPaymentById(paymentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

// CreatePayment creates a new payment for an order.
func CreatePayment(c *gin.Context) {
	var payment domain.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := domain.CreatePayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payment)
}

// UpdatePayment updates an existing payment.
func UpdatePayment(c *gin.Context) {
	paymentId := c.Param("paymentId")
	var payment domain.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := domain.UpdatePayment(paymentId, &payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

// DeletePayment deletes a payment.
func DeletePayment(c *gin.Context) {
	paymentId := c.Param("paymentId")
	if err := domain.DeletePayment(paymentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}