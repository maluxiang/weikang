package handler

import (
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	"weikang/client/points_api/proto/points"
)

func UpdatePoints(c *gin.Context) {
	var req *points.UpdatePointsRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	gid := shortuuid.New()
	// req := &busi.BusiReq{Amount: 30, TransInResult: "FAILURE"}
	saga := dtmgrpc.NewSagaGrpc("localhost:36790", gid). // server:"localhost:36790"
		Add("localhost:8080/points.Points/UpdatePoints", "localhost:8080/points.Points/UpdatePointsCompensation", req)
	// "localhost:8080/points.Points/UpdatePoints"为服务端接口正向操作地址，UpdatePointsCompensation为服务端接口的补偿操作
	err := saga.Submit()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Points updated successfully"})
}
