package router

import (
	"github.com/gin-gonic/gin"
	"weikang/client/smartdoctor_api/hander"
)

func Doctor(path *gin.RouterGroup) {
	doctor := path.Group("/doctor")
	{
		doctor.POST("/question", hander.DoctorQuestion)     //智能医生提问
		doctor.POST("/list", hander.DoctorQuestionList)     //智能医生提问历史列表
		doctor.POST("/delete", hander.DoctorQuestionDelete) //智能医生提问历史删除
	}
}
