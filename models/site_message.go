package models

import "time"

type SiteMessage struct {
	ID        int64     `gorm:"primaryKey"`
	UserID    int64     `gorm:"not null"`
	Title     string    `gorm:"size:100;not null"`
	Content   string    `gorm:"type:text;not null"`
	IsRead    int8      `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (s SiteMessage) TableName() string {
	return "site_message"
}
