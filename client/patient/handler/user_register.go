package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weikang/client/patient/form"
	"weikang/client/patient/proto/patient"
	"weikang/client/patient/server"
)

func Register(c *gin.Context) {

	var json form.Patient
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	register, err := server.PatientClient.Register(c, &patient.RegisterReq{
		Name:       json.Name,
		Age:        int64(json.Age),
		Sex:        patient.Sex(json.Sex),
		IdCard:     json.IdCard,
		Phone:      json.Phone,
		Department: json.Department,
		Doctor:     json.Doctor,
		Status:     int64(json.Status),
	})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": register.Msg})
}
