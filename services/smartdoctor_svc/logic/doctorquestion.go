package logic

import (
	"context"
	"encoding/json"
	"weikang/models"
	"weikang/pkg"
	"weikang/services/smartdoctor_svc/proto/smartDoctor"
)

func (s Server) DoctorQuestion(ctx context.Context, in *smartDoctor.DoctorQuestionRequest) (*smartDoctor.DoctorQuestionResponse, error) {
	question := in.Question
	que, err := pkg.ChatGpt(question)
	if err != nil {
		return &smartDoctor.DoctorQuestionResponse{
			Answer: nil,
		}, nil
	}
	doctor := models.SmartDoctor{
		Question: question,
		Answer:   que,
	}
	err = doctor.Create()
	if err != nil {
		return &smartDoctor.DoctorQuestionResponse{
			Answer: nil,
		}, nil
	}
	marshal, _ := json.Marshal(que)
	return &smartDoctor.DoctorQuestionResponse{
		Answer: marshal,
	}, nil
}
