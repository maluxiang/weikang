package logic

import (
	"context"
	"weikang/models"
	"weikang/pkg"
	"weikang/services/datacenter_srv/proto/datacenter"
)

func (s Server) ReportHealthData(ctx context.Context, request *datacenter.ReportHealthDataRequest) (*datacenter.ReportHealthDataResponse, error) {
	health := &models.HealthData{
		DeviceID:               request.DeviceId,
		UserID:                 request.UserId,
		Timestamp:              request.Timestamp,
		HeartRate:              int(request.HeartRate),
		BloodPressureSystolic:  int(request.BloodPressureSystolic),
		BloodPressureDiastolic: int(request.BloodPressureDiastolic),
		BodyTemperature:        request.BodyTemperature,
		Steps:                  int(request.Steps),
		SleepDurationMinutes:   int(request.SleepDurationMinutes),
		ActivityCaloriesBurned: int(request.ActivityCaloriesBurned),
		BloodGlucose:           request.BloodGlucose,
		Weight:                 request.Weight,
		Height:                 request.Height,
		DeviceStatus:           request.DeviceStatus,
	}

	// 使用 GORM 插入数据到数据库
	if err := health.DataCreate(); err != nil {
		return &datacenter.ReportHealthDataResponse{
			Success: false,
			Message: "数据上报失败",
		}, err
	}

	// 插入文档到 MongoDB
	if err := pkg.InsertDocument("data", health); err != nil {
		return &datacenter.ReportHealthDataResponse{
			Success: false,
			Message: "mongodb失败",
		}, err
	}

	return &datacenter.ReportHealthDataResponse{
		Success: true,
		Message: "数据上报成功",
	}, nil
}
