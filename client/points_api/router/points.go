package router

import (
	"github.com/gin-gonic/gin"
	"weikang/client/points_api/handler"
)

func Points(path *gin.RouterGroup) {
	points := path.Group("/points")
	{
		points.POST("/create_points", handler.CreatePoints)
		points.GET("/get_user_all_points", handler.GetUserAllPoints)
		points.POST("/update_points", handler.UpdatePoints)
		points.POST("/update_points_compensation", handler.UpdatePointsCompensation)
	}
}
