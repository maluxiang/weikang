package models

import (
	"context"
	"gorm.io/gorm"
	"time"
	"weikang/global"
)

type MedicalDevice struct {
	ID          uint64         `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	Name        string         `gorm:"type:varchar(255)"`
	Brand       string         `gorm:"type:varchar(30)"`
	Model       string         `gorm:"type:varchar(30)"`
	Description string         `gorm:"type:varchar(30)"`
	Stock       int            `gorm:"type:int(10)"`
	Price       float64        `gorm:"type:decimal(10,2)"`
	Status      int            `gorm:"type:int(11)"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
}

func (m MedicalDevice) TableName() string {
	return "medical_device"
}

func (m MedicalDevice) Create() error {
	err := global.DB.Create(&m).Error
	if err == nil {
		// 同步到 ES
		_, _ = global.EsClient.Index().Index("medical_devices").BodyJson(m).Do(context.Background())
	}
	return err
}

func (m MedicalDevice) MedicalDeviceList() ([]MedicalDevice, error) {
	var devices []MedicalDevice
	err := global.DB.Find(&devices).Error
	return devices, err
}
