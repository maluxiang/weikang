package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"weikang/client/medical_api/proto/medical"
	"weikang/client/medical_api/server"
)

func AddDevice(c *gin.Context) {
	var req medical.AddDeviceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := server.MedicalClient.AddDevice(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
