package logic

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"time"
	"weikang/global"
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

	result, _ := global.Rds.Get(context.Background(), "health-data").Result()
	if result != "" {
		json.Unmarshal([]byte(result), &list)
		return &datacenter.ReportHealthDataListResponse{
			HealthDataInfo: list,
		}, nil
	}

	for _, healthData := range health {
		list = append(list, &datacenter.HealthDataInfo{
			Id:                     int64(healthData.ID),
			DeviceId:               healthData.DeviceId,
			UserId:                 healthData.UserId,
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

	marshal, _ := json.Marshal(list)
	global.Rds.Set(context.Background(), "health-data", marshal, time.Hour*5)

	return &datacenter.ReportHealthDataListResponse{
		HealthDataInfo: list,
	}, nil

}
