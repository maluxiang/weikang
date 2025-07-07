package models

import (
	"gorm.io/gorm"
	"time"
	"weikang/global"
)

type Patient struct {
	ID         uint64         `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	Name       string         `gorm:"type:varchar(255)"`
	Age        int            `gorm:"type:integer"`
	Sex        int            `gorm:"type:integer"`
	IdCard     string         `gorm:"type:varchar(255)"`
	Phone      string         `gorm:"type:varchar(255)"`
	Department string         `gorm:"type:varchar(255)"`
	Doctor     string         `gorm:"type:varchar(255)"`
	Status     int            `gorm:"type:int(11)"`
	CreatedAt  time.Time      `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
}

func Register(p Patient) (err error) {
	err = global.DB.Create(&p).Error
	return
}

func Find(name string, p Patient) Patient {
	global.DB.Where("name = ?", name).Find(&p)
	return p
}
