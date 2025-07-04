package models

import (
	"gorm.io/gorm"
	"time"
)

type UploadFiles struct {
	ID        int64          `gorm:"column:id;type:bigint;primaryKey;not null;" json:"id"`
	FileName  string         `gorm:"column:file_name;type:longtext;" json:"file_name"`
	FilePath  string         `gorm:"column:file_path;type:longtext;" json:"file_path"`
	FileType  string         `gorm:"column:file_type;type:longtext;" json:"file_type"`
	FileSize  int64          `gorm:"column:file_size;type:bigint;default:NULL;" json:"file_size"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
}

func (u UploadFiles) TableName() string {
	return "upload_files"
}
