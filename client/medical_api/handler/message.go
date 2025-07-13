package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	pb "weikang/client/medical_api/proto/medical"
	"weikang/client/medical_api/server" // 假设有 gRPC client 封装
)

// 获取消息列表
func GetMessageList(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	req := &pb.GetMessageListRequest{UserId: userID}
	resp, err := server.MedicalClient.GetMessageList(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "获取失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": resp.Messages})
}

// 标记已读
func MarkMessageRead(c *gin.Context) {
	idStr := c.PostForm("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	req := &pb.MarkMessageReadRequest{Id: id}
	_, err := server.MedicalClient.MarkMessageRead(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": "操作失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "已标记为已读"})
}
