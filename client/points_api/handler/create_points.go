package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"weikang/client/points_api/form"
	"weikang/client/points_api/proto/points"
	"weikang/client/points_api/server"
)

func CreatePoints(c *gin.Context) {
	var req form.Points
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := server.PointsClient.CreatePoints(context.Background(), &points.CreatePointsRequest{
		UserID: req.UserID,
		Points: req.Points,
	})
	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Points created successfully"})
}
