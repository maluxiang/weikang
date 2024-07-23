package logic

import (
	"context"
	"weikang/models"
	"weikang/services/smartdoctor_svc/proto/smartDoctor"
)

func (s Server) DoctorQuestionDelete(ctx context.Context, in *smartDoctor.DoctorQuestionDeleteRequest) (*smartDoctor.DoctorQuestionDeleteResponse, error) {
	id := in.Id
	err := models.DoctorDelete(id)
	if err != nil {
		return &smartDoctor.DoctorQuestionDeleteResponse{
			Code:    500,
			Message: "历史记录删除失败",
		}, nil
	}
	return &smartDoctor.DoctorQuestionDeleteResponse{
		Code:    200,
		Message: "历史记录删除成功",
	}, nil
}
