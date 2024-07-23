package router

import (
	"github.com/gin-gonic/gin"
	"weikang/client/patient/handler"
)

func User(path *gin.RouterGroup) {
	user := path.Group("/patient")
	{
		user.POST("/register", handler.Register)
		user.POST("/search", handler.EsSearch)
	}
}
