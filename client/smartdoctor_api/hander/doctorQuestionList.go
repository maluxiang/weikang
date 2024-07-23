package hander

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"weikang/client/smartdoctor_api/proto/smartDoctor"
	"weikang/client/smartdoctor_api/server"
)

func DoctorQuestionList(c *gin.Context) {
	res, err := server.DockerClient.DoctorQuestionList(context.Background(), &smartDoctor.DoctorQuestionListRequest{})
	if err != nil {
		c.JSON(500, &gin.H{
			"List": nil,
		})
		return
	}
	var i interface{}
	json.Unmarshal(res.List, &i)
	c.JSON(200, &gin.H{
		"List": i,
	})
}
