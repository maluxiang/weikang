package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"weikang/models"
	"weikang/services/points_svc/proto/points"
)

func (s Server) UpdatePointsCompensation(ctx context.Context, in *points.UpdatePointsCompensationRequest) (*points.UpdatePointsCompensationResponse, error) {
	// 实现更新积分的逻辑
	p := models.Points{
		UserID: in.UserID,
	}
	info, err := p.GetInfo()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &points.UpdatePointsCompensationResponse{}, status.Error(codes.Aborted, "用户不存在")
	}
	if err != nil {
		return &points.UpdatePointsCompensationResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info.Points -= in.Points
	err = info.Update()
	if err != nil {
		return &points.UpdatePointsCompensationResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &points.UpdatePointsCompensationResponse{Success: true}, nil
}
