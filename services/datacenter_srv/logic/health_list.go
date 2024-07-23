package logic

import (
	"context"
	"go.uber.org/zap"
	"weikang/models"
	"weikang/services/datacenter_srv/proto/datacenter"
)

func (s Server) ReportHealthDataList(ctx context.Context, request *datacenter.ReportHealthDataListRequest) (*datacenter.ReportHealthDataListResponse, error) {
	var data models.HealthData
	health, err := data.DataList()
	if err != nil {
		zap.S().Info("获取健康数据失败")
		return nil, err
	}
	var list []*datacenter.HealthDataInfo
	for _, healthData := range health {
		list = append(list, &datacenter.HealthDataInfo{
			Id:                     int64(healthData.ID),
			DeviceId:               healthData.DeviceID,
			UserId:                 healthData.UserID,
			Timestamp:              healthData.Timestamp,
			HeartRate:              int32(healthData.HeartRate),
			BloodPressureSystolic:  int32(healthData.BloodPressureSystolic),
			BloodPressureDiastolic: int32(healthData.BloodPressureDiastolic),
			BodyTemperature:        healthData.BodyTemperature,
			Steps:                  int32(healthData.Steps),
			SleepDurationMinutes:   int32(healthData.SleepDurationMinutes),
			ActivityCaloriesBurned: int32(healthData.ActivityCaloriesBurned),
			BloodGlucose:           healthData.BloodGlucose,
			Weight:                 healthData.Weight,
			Height:                 healthData.Height,
			DeviceStatus:           healthData.DeviceStatus,
		})

	}

	return &datacenter.ReportHealthDataListResponse{
		HealthDataInfo: list,
	}, nil

}
