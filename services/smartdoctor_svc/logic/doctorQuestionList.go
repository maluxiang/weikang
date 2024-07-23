package logic

import (
	"context"
	"encoding/json"
	"weikang/models"
	"weikang/services/smartdoctor_svc/proto/smartDoctor"
)

func (s Server) DoctorQuestionList(ctx context.Context, in *smartDoctor.DoctorQuestionListRequest) (*smartDoctor.DoctorQuestionListResponse, error) {
	list, err := models.DoctorList()
	if err != nil {
		return &smartDoctor.DoctorQuestionListResponse{
			List: nil,
		}, nil
	}
	marshal, _ := json.Marshal(list)
	return &smartDoctor.DoctorQuestionListResponse{
		List: marshal,
	}, nil
}
