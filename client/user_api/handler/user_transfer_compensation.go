package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"weikang/client/user_api/form"
	"weikang/client/user_api/proto/user"
	"weikang/client/user_api/server"
)

func TransferCompensation(c *gin.Context) {
	var req form.Transfer
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := server.UserClient.TransferCompensation(context.Background(), &user.TransferCompensationRequest{
		FromID:        req.FromID,
		ToID:          req.ToID,
		AccountNumber: req.AccountNumber,
		Amount:        req.Amount,
	})
	if err != nil {
		c.JSON(409, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Transfer successful"})
}
