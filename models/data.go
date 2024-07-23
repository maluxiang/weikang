package models

import (
	"gorm.io/gorm"
	"weikang/global"
)

type HealthData struct {
	ID                     int     `gorm:"primaryKey"`            // 数据条目ID，设置为主键
	DeviceID               string  `gorm:"size:255;comment:设备ID"` // 设备ID，设置字段大小为255，添加注释
	UserID                 string  `gorm:"size:255;comment:用户ID"` // 用户ID，设置字段大小为255，添加注释
	Timestamp              string  `gorm:"comment:数据记录时间"`        // 数据记录时间，添加注释
	HeartRate              int     `gorm:"comment:心率（每分钟心跳数）"`    // 心率，添加注释
	BloodPressureSystolic  int     `gorm:"comment:血压收缩压"`         // 血压收缩压，添加注释
	BloodPressureDiastolic int     `gorm:"comment:血压舒张压"`         // 血压舒张压，添加注释
	BodyTemperature        float64 `gorm:"comment:体温（摄氏度）"`       // 体温，添加注释
	Steps                  int     `gorm:"comment:步数"`            // 步数，添加注释
	SleepDurationMinutes   int     `gorm:"comment:睡眠时长（分钟）"`      // 睡眠时长，添加注释
	ActivityCaloriesBurned int     `gorm:"comment:活动燃烧的卡路里"`      // 活动燃烧的卡路里，添加注释
	BloodGlucose           float64 `gorm:"comment:血糖浓度"`          // 血糖浓度，添加注释
	Weight                 float64 `gorm:"comment:体重（公斤）"`        // 体重，添加注释
	Height                 float64 `gorm:"comment:身高（厘米）"`        // 身高，添加注释
	DeviceStatus           string  `gorm:"comment:设备状态"`          // 设备状态，添加注释
	gorm.Model
}

func (h *HealthData) DataCreate() error {
	return global.DB.Create(h).Error
}

// 列表
func (h *HealthData) DataList() ([]*HealthData, error) {
	var list []*HealthData
	err := global.DB.Find(&list).Error
	return list, err
}

// 删除
func DataDelete(id int64) error {
	return global.DB.Where("id = ?", id).Delete(&HealthData{}).Error
}
