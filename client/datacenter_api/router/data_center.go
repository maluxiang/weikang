package router

import (
	"github.com/gin-gonic/gin"
	"weikang/client/datacenter_api/hander"
)

func DataCenter(path *gin.RouterGroup) {
	data := path.Group("/data")
	{
		data.POST("reporting", hander.DataReporting)             //数据上报
		data.GET("list", hander.DataCenterList)                  //数据查询
		data.DELETE("delete/:id", hander.ReportHealthDataDelete) //删除数据

	}
}
