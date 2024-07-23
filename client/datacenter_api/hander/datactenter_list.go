package hander

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"weikang/client/datacenter_api/proto/datacenter"
	"weikang/client/datacenter_api/server"
)

func DataCenterList(c *gin.Context) {
	data, err := server.DataCenterClient.ReportHealthDataList(context.Background(), &datacenter.ReportHealthDataListRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("健康设备展示数据失败: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "健康设备展示数据成功", "data": data.HealthDataInfo})
	return
}
