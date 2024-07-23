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

func (s Server) UpdatePoints(ctx context.Context, in *points.UpdatePointsRequest) (*points.UpdatePointsResponse, error) {
	// 实现更新积分的逻辑
	p := models.Points{
		UserID: in.UserID,
	}
	
	info, err := p.GetInfo()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &points.UpdatePointsResponse{}, status.Error(codes.Aborted, "User not found")
	}
	if err != nil {
		return &points.UpdatePointsResponse{}, status.Error(codes.Aborted, "Failed to get user info")
	}
	info.Points += in.Points
	err = info.Update()
	if err != nil {
		return &points.UpdatePointsResponse{}, status.Error(codes.Aborted, "Failed to update points")
	}
	if in.Points == 20 {
		return &points.UpdatePointsResponse{}, status.Error(codes.Aborted, "Points cannot be negative")
	}
	return &points.UpdatePointsResponse{Success: true}, nil
}
