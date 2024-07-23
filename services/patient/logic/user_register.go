package logic

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	"weikang/models"
	"weikang/services/patient/proto/patient"
)

func (Server) Register(ctx context.Context, in *patient.RegisterReq) (*patient.RegisterResp, error) {
	// 实现用户注册的逻辑

	err := models.Register(models.Patient{
		Model:      gorm.Model{},
		Name:       in.Name,
		Age:        int(in.Age),
		Sex:        int(in.Sex),
		IdCard:     in.IdCard,
		Phone:      in.Phone,
		Department: in.Department,
		Doctor:     in.Doctor,
		Status:     int(in.Status),
	})
	if err != nil {
		return nil, err
	}

	return &patient.RegisterResp{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}
