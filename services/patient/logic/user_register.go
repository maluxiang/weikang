package logic

import (
	"context"
	"net/http"
	"weikang/models"
	"weikang/pkg"
	"weikang/services/patient/proto/patient"
)

func (Server) Register(ctx context.Context, in *patient.RegisterReq) (*patient.RegisterResp, error) {
	// 实现用户注册的逻辑

	user := models.Patient{
		Name:       in.Name,
		Age:        int(in.Age),
		Sex:        int(in.Sex),
		IdCard:     in.IdCard,
		Phone:      in.Phone,
		Department: in.Department,
		Doctor:     in.Doctor,
		Status:     int(in.Status),
	}
	err := models.Register(user)
	if err != nil {
		return nil, err
	}

	p := models.Patient{}
	userlogic := models.Find(in.Name, p)
	pkg.ElasticInit("userlogic", userlogic)

	return &patient.RegisterResp{
		Code: http.StatusOK,
		Msg:  "success",
	}, nil
}
