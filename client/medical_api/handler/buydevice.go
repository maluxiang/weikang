package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weikang/client/medical_api/proto/medical"
)

func BuyDevice(c *gin.Context) {
	var req medical.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
