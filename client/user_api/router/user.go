package router

import (
	"github.com/gin-gonic/gin"
	"weikang/client/user_api/handler"
)

func User(path *gin.RouterGroup) {
	user := path.Group("/user")
	{
		user.POST("/register", handler.Register)
		user.POST("/create_account", handler.CreateAccount)
		user.POST("/transfer", handler.Transfer)
		user.POST("/transfer_compensation", handler.TransferCompensation)
	}
}
