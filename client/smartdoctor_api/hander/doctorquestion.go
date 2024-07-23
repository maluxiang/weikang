package hander

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"weikang/client/smartdoctor_api/proto/smartDoctor"
	"weikang/client/smartdoctor_api/server"
)

func DoctorQuestion(c *gin.Context) {
	question := c.PostForm("question")
	res, err := server.DockerClient.DoctorQuestion(context.Background(), &smartDoctor.DoctorQuestionRequest{
		Question: question,
	})
	if err != nil {
		c.JSON(500, &gin.H{
			"Answer": nil,
		})
		return
	}
	var i interface{}
	json.Unmarshal(res.Answer, &i)
	c.JSON(200, &gin.H{
		"Answer": i,
	})
}
