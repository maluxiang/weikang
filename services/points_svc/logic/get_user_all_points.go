package logic

import (
	"context"
	"weikang/models"
	"weikang/services/points_svc/proto/points"
)

func (s Server) GetUserAllPoints(ctx context.Context, in *points.GetUserAllPointsRequest) (*points.GetUserAllPointsResponse, error) {
	// 实现获取用户所有积分的逻辑
	var p models.Points
	all, err := p.GetAll()
	if err != nil {
		return &points.GetUserAllPointsResponse{}, err
	}
	var pointsList []*points.PointsInfo
	for _, v := range all {
		pointsList = append(pointsList, &points.PointsInfo{
			UserID: v.UserID,
			Points: v.Points,
		})
	}
	return &points.GetUserAllPointsResponse{PointsList: pointsList}, nil
}
