package models

import (
	"gorm.io/gorm"
	"weikang/global"
)

type Patient struct {
	gorm.Model
	Name       string `gorm:"type:varchar(255)"`
	Age        int    `gorm:"type:integer"`
	Sex        int    `gorm:"type:integer"`
	IdCard     string `gorm:"type:varchar(255)"`
	Phone      string `gorm:"type:varchar(255)"`
	Department string `gorm:"type:varchar(255)"`
	Doctor     string `gorm:"type:varchar(255)"`
	Status     int    `gorm:"type:int(11)"`
}

func Register(p Patient) (err error) {
	err = global.DB.Create(&p).Error
	return
}
