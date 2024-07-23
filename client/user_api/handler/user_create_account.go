package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"weikang/client/user_api/form"
	"weikang/client/user_api/proto/user"
	"weikang/client/user_api/server"
)

func CreateAccount(c *gin.Context) {
	var req form.Account
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := server.UserClient.CreateAccount(context.Background(), &user.CreateAccountRequest{
		UserID:   req.UserID,
		Currency: req.Currency,
		Balance:  req.Balance,
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Account created successfully"})
}
