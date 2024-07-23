package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"weikang/client/user_api/form"
	"weikang/client/user_api/proto/user"
	"weikang/client/user_api/server"
)

func Register(c *gin.Context) {
	var req form.User
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err = server.UserClient.Register(context.Background(), &user.RegisterRequest{
		User: &user.UserInfo{
			Username: req.Username,
			Password: req.Password,
			Email:    req.Email,
			Phone:    req.Phone,
		},
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "注册成功"})
}
