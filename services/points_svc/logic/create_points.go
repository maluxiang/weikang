package logic

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"weikang/models"
	"weikang/services/points_svc/proto/points"
)

func (s Server) CreatePoints(ctx context.Context, in *points.CreatePointsRequest) (*points.CreatePointsResponse, error) {
	// 实现创建积分的逻辑
	fmt.Println(in)
	p := models.Points{
		UserID: in.UserID,
		Points: in.Points,
	}
	info, err := p.GetInfo()

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return &points.CreatePointsResponse{}, err
	}
	err = p.FindByID(uint(info.ID))
	if err != nil {
		return &points.CreatePointsResponse{}, errors.New("积分已存在")
	}
	err = p.Create()
	if err != nil {
		return &points.CreatePointsResponse{}, err
	}
	return &points.CreatePointsResponse{Success: true}, nil
}
