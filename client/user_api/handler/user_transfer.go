package handler

import (
	"context"
	"github.com/dtm-labs/client/dtmcli"
	
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	"weikang/client/user_api/form"
	"weikang/client/user_api/proto/user"
	"weikang/client/user_api/server"
)

func Transfer(c *gin.Context) {
	var req form.Transfer
	const qsBusi = "http://localhost:8888"           // 微服务地址
	DtmServer := "http://localhost:36789/api/dtmsvr" // DTM服务地址
	// 创建Saga事务
	saga := dtmcli.NewSaga(DtmServer, shortuuid.New()).
		// 添加子事务
		Add(qsBusi+"/v1/user/transfer", qsBusi+"/v1/user/transfer_compensation", form.Transfer{})
	// 提交事务
	err := saga.Submit()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err = c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err = server.UserClient.Transfer(context.Background(), &user.TransferRequest{
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
