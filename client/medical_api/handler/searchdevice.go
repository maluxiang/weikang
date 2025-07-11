package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"weikang/client/medical_api/proto/medical"
	"weikang/client/medical_api/server"
)

func SearchDevice(c *gin.Context) {
	keyword := c.Query("keyword")
	req := &medical.SearchDeviceRequest{Keyword: keyword}
	resp, err := server.MedicalClient.SearchDevice(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
