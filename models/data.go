package models

import (
	"gorm.io/gorm"
	"time"
	"weikang/global"
)

type HealthData struct {
	ID                     int64          `gorm:"column:id;type:bigint;primaryKey;not null;" json:"id"`
	DeviceId               string         `gorm:"column:device_id;type:varchar(255);comment:设备ID;default:NULL;" json:"device_id"`                             // 设备ID
	UserId                 string         `gorm:"column:user_id;type:varchar(255);comment:用户ID;default:NULL;" json:"user_id"`                                 // 用户ID
	Timestamp              string         `gorm:"column:timestamp;type:longtext;comment:数据记录时间;" json:"timestamp"`                                            // 数据记录时间
	HeartRate              int64          `gorm:"column:heart_rate;type:bigint;comment:心率（每分钟心跳数）;default:NULL;" json:"heart_rate"`                           // 心率（每分钟心跳数）
	BloodPressureSystolic  int64          `gorm:"column:blood_pressure_systolic;type:bigint;comment:血压收缩压;default:NULL;" json:"blood_pressure_systolic"`      // 血压收缩压
	BloodPressureDiastolic int64          `gorm:"column:blood_pressure_diastolic;type:bigint;comment:血压舒张压;default:NULL;" json:"blood_pressure_diastolic"`    // 血压舒张压
	BodyTemperature        float64        `gorm:"column:body_temperature;type:double;comment:体温（摄氏度）;default:NULL;" json:"body_temperature"`                  // 体温（摄氏度）
	Steps                  int64          `gorm:"column:steps;type:bigint;comment:步数;default:NULL;" json:"steps"`                                             // 步数
	SleepDurationMinutes   int64          `gorm:"column:sleep_duration_minutes;type:bigint;comment:睡眠时长（分钟）;default:NULL;" json:"sleep_duration_minutes"`     // 睡眠时长（分钟）
	ActivityCaloriesBurned int64          `gorm:"column:activity_calories_burned;type:bigint;comment:活动燃烧的卡路里;default:NULL;" json:"activity_calories_burned"` // 活动燃烧的卡路里
	BloodGlucose           float64        `gorm:"column:blood_glucose;type:double;comment:血糖浓度;default:NULL;" json:"blood_glucose"`                           // 血糖浓度
	Weight                 float64        `gorm:"column:weight;type:double;comment:体重（公斤）;default:NULL;" json:"weight"`                                       // 体重（公斤）
	Height                 float64        `gorm:"column:height;type:double;comment:身高（厘米）;default:NULL;" json:"height"`                                       // 身高（厘米）
	DeviceStatus           string         `gorm:"column:device_status;type:longtext;comment:设备状态;" json:"device_status"`                                      // 设备状态
	CreatedAt              time.Time      `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"column:updated_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
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
