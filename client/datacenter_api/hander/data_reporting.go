package hander

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"weikang/client/datacenter_api/proto/datacenter"
	"weikang/client/datacenter_api/server"
)

func DataReporting(c *gin.Context) {
	// 获取POST请求参数
	DeviceId := c.PostForm("device_id")
	UserID := c.PostForm("user_id")
	Timestamp := c.PostForm("timestamp")
	Rate := c.PostForm("heart_rate")
	Blood := c.PostForm("blood_pressure_systolic")
	Diastolic := c.PostForm("blood_pressure_diastolic")
	Body := c.PostForm("body_temperature")
	St := c.PostForm("steps")
	Sleep := c.PostForm("sleep_duration_minutes")
	Activity := c.PostForm("activity_calories_burned")
	Bloods := c.PostForm("blood_glucose")
	We := c.PostForm("weight")
	He := c.PostForm("height")
	DeviceStatus := c.PostForm("device_status")

	// 转换数据类型并处理错误
	HeartRate, err := strconv.Atoi(Rate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "心率输入错误"})
		return
	}

	BloodPressureSystolic, err := strconv.Atoi(Blood)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "血压收缩压输入错误"})
		return
	}

	BloodPressureDiastolic, err := strconv.Atoi(Diastolic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "血压舒张压输入错误"})
		return
	}

	BodyTemperature, err := strconv.ParseFloat(Body, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "体温输入错误"})
		return
	}

	Steps, err := strconv.Atoi(St)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "步数输入错误"})
		return
	}

	SleepDurationMinutes, err := strconv.Atoi(Sleep)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "睡眠时长输入错误"})
		return
	}

	ActivityCaloriesBurned, err := strconv.Atoi(Activity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "活动热量消耗输入错误"})
		return
	}

	BloodGlucose, err := strconv.ParseFloat(Bloods, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "血糖输入错误"})
		return
	}

	Weight, err := strconv.ParseFloat(We, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "体重输入错误"})
		return
	}

	Height, err := strconv.ParseFloat(He, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "身高输入错误"})
		return
	}

	// 构造 Protocol Buffers 请求体并发送给数据中心服务
	_, err = server.DataCenterClient.ReportHealthData(context.Background(), &datacenter.ReportHealthDataRequest{
		DeviceId:               DeviceId,
		UserId:                 UserID,
		Timestamp:              Timestamp,
		HeartRate:              int32(HeartRate),
		BloodPressureSystolic:  int32(BloodPressureSystolic),
		BloodPressureDiastolic: int32(BloodPressureDiastolic),
		BodyTemperature:        BodyTemperature,
		Steps:                  int32(Steps),
		SleepDurationMinutes:   int32(SleepDurationMinutes),
		ActivityCaloriesBurned: int32(ActivityCaloriesBurned),
		BloodGlucose:           BloodGlucose,
		Weight:                 Weight,
		Height:                 Height,
		DeviceStatus:           DeviceStatus,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("数据上报失败: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "数据上报成功"})
}
