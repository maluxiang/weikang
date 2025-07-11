package logic

import (
	"context"
	"testing"
	"time"
	"weikang/models"
	"weikang/pkg/testutils"
	"weikang/services/datacenter_srv/proto/datacenter"
)

func TestServer_ReportHealthData(t *testing.T) {
	// 设置测试数据库
	testDB := testutils.SetupTestDB(t)
	defer testDB.Cleanup()

	server := &Server{}

	// 测试用例1：正常数据上报
	t.Run("正常数据上报", func(t *testing.T) {
		request := &datacenter.ReportHealthDataRequest{
			DeviceId:               "test_device_001",
			UserId:                 "test_user_001",
			Timestamp:              time.Now().Format("2006-01-02 15:04:05"),
			HeartRate:              75,
			BloodPressureSystolic:  120,
			BloodPressureDiastolic: 80,
			BodyTemperature:        36.5,
			Steps:                  8000,
			SleepDurationMinutes:   480,
			ActivityCaloriesBurned: 300,
			BloodGlucose:           5.5,
			Weight:                 65.0,
			Height:                 170.0,
			DeviceStatus:           "normal",
		}

		response, err := server.ReportHealthData(context.Background(), request)
		if err != nil {
			t.Errorf("ReportHealthData() error = %v", err)
		}

		if !response.Success {
			t.Errorf("Expected success to be true, got %v", response.Success)
		}

		if response.Message != "数据上报成功" {
			t.Errorf("Expected message to be '数据上报成功', got %s", response.Message)
		}
	})

	// 测试用例2：无效数据上报
	t.Run("无效数据上报", func(t *testing.T) {
		request := &datacenter.ReportHealthDataRequest{
			DeviceId: "",
			UserId:   "",
			// 缺少必要字段
		}

		response, err := server.ReportHealthData(context.Background(), request)
		if err != nil {
			// 期望有错误
			return
		}

		if response.Success {
			t.Error("Expected success to be false for invalid data")
		}
	})
}

func TestServer_ReportHealthDataList(t *testing.T) {
	// 设置测试数据库
	testDB := testutils.SetupTestDB(t)
	defer testDB.Cleanup()

	server := &Server{}

	// 创建测试数据
	request := &datacenter.ReportHealthDataRequest{
		DeviceId:  "test_device_list",
		UserId:    "test_user_list",
		HeartRate: 75,
	}
	server.ReportHealthData(context.Background(), request)

	// 测试列表查询
	t.Run("获取健康数据列表", func(t *testing.T) {
		listRequest := &datacenter.ReportHealthDataListRequest{}
		response, err := server.ReportHealthDataList(context.Background(), listRequest)
		if err != nil {
			t.Errorf("ReportHealthDataList() error = %v", err)
		}

		if len(response.HealthDataInfo) == 0 {
			t.Error("Expected at least one health data record")
		}

		// 验证返回的数据结构
		for _, data := range response.HealthDataInfo {
			if data.DeviceId == "" {
				t.Error("Expected device_id to be not empty")
			}
			if data.UserId == "" {
				t.Error("Expected user_id to be not empty")
			}
		}
	})
}

func TestServer_ReportHealthDataDelete(t *testing.T) {
	// 设置测试数据库
	testDB := testutils.SetupTestDB(t)
	defer testDB.Cleanup()

	server := &Server{}

	// 创建测试数据
	createRequest := &datacenter.ReportHealthDataRequest{
		DeviceId:  "test_device_delete",
		UserId:    "test_user_delete",
		HeartRate: 70,
	}
	_, _ = server.ReportHealthData(context.Background(), createRequest)

	// 获取创建的记录ID（这里需要从数据库查询）
	var healthData models.HealthData
	testDB.DB.Where("device_id = ?", "test_device_delete").First(&healthData)

	// 测试删除
	t.Run("删除健康数据", func(t *testing.T) {
		deleteRequest := &datacenter.ReportHealthDataDeleteRequest{
			Id: healthData.ID,
		}

		response, err := server.ReportHealthDataDelete(context.Background(), deleteRequest)
		if err != nil {
			t.Errorf("ReportHealthDataDelete() error = %v", err)
		}

		if response.Code != 200 {
			t.Errorf("Expected code to be 200, got %d", response.Code)
		}

		// 验证数据是否被删除
		var result models.HealthData
		err = testDB.DB.Where("id = ?", healthData.ID).First(&result).Error
		if err == nil {
			t.Error("Expected data to be deleted, but it still exists")
		}
	})
}
