package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"weikang/client/medical_api/proto/medical"
	"weikang/client/medical_api/server"
)

func ListDevice(c *gin.Context) {
	req := &medical.ListDeviceRequest{}
	resp, err := server.MedicalClient.ListDevice(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
