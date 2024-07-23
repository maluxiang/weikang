package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"weikang/client/points_api/proto/points"
	"weikang/client/points_api/server"
)

func GetUserAllPoints(c *gin.Context) {
	allPoints, err := server.PointsClient.GetUserAllPoints(context.Background(), &points.GetUserAllPointsRequest{})
	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success", "data": allPoints.PointsList})
}
