package models

import (
	"gorm.io/gorm"
	"time"
	"weikang/global"
)

type SmartDoctor struct {
	ID        uint64         `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	Question  string         `gorm:"column:question;type:longtext;" json:"question"`
	Answer    string         `gorm:"column:answer;type:longtext;" json:"answer"`
	UserID    int            `gorm:"column:user_id;type:int(10);" json:"user_id"`
}

func (SmartDoctor) TableName() string {
	return "smart_doctor"
}

func (s SmartDoctor) Create() error {
	return global.DB.Create(&s).Error
}

func DoctorList() ([]SmartDoctor, error) {
	var doctor []SmartDoctor
	global.DB.Find(&doctor)
	return doctor, nil
}

func DoctorDelete(id int64) error {
	return global.DB.Where("id = ?", id).Delete(&SmartDoctor{}).Error
}
