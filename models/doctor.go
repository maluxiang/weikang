package models

import (
	"gorm.io/gorm"
	"weikang/global"
)

type SmartDoctor struct {
	gorm.Model
	Question string
	Answer   string
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
