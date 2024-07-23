package hander

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"weikang/client/datacenter_api/proto/datacenter"
	"weikang/client/datacenter_api/server"
)

func ReportHealthDataDelete(c *gin.Context) {
	id := c.Param("id") // 从路径参数中获取id
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// 使用获取到的idNum进行处理，例如：
	res, err := server.DataCenterClient.ReportHealthDataDelete(context.Background(), &datacenter.ReportHealthDataDeleteRequest{
		Id: int64(idNum),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": res.Message})
}
