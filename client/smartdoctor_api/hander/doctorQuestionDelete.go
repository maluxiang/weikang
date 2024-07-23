package hander

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"weikang/client/smartdoctor_api/proto/smartDoctor"
	"weikang/client/smartdoctor_api/server"
)

func DoctorQuestionDelete(c *gin.Context) {
	ids := c.PostForm("id")
	id, _ := strconv.Atoi(ids)
	res, err := server.DockerClient.DoctorQuestionDelete(context.Background(), &smartDoctor.DoctorQuestionDeleteRequest{
		Id: int64(id),
	})
	if err != nil {
		c.JSON(500, &gin.H{
			"message": res.Message,
		})
		return
	}
	c.JSON(200, &gin.H{
		"message": res.Message,
	})
}
