package models

import (
	"gorm.io/gorm"
)

type UploadFile struct {
	Id       int
	FileName string
	FilePath string
	FileType string
	FileSize int64
	gorm.Model
}
