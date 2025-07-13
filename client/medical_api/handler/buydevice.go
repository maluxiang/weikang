package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"weikang/client/medical_api/proto/medical"
	"weikang/client/medical_api/server"
)

// BuyDevice 下单接口
func BuyDevice(c *gin.Context) {
	var req medical.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "参数错误: " + err.Error()})
		return
	}

	// 参数基本校验
	if req.ProductId == 0 || req.UserId == 0 || req.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "商品ID、用户ID、数量不能为空且数量需大于0"})
		return
	}

	resp, err := server.MedicalClient.CreateOrder(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "下单失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": resp.Message,
		"url":     resp.Url,
	})
}
