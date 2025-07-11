package models

import "time"

// 订单评论表
type OrderComment struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	OrderID   uint64    `gorm:"not null;index;comment:订单ID"`
	UserID    uint64    `gorm:"not null;index;comment:用户ID"`
	Content   string    `gorm:"type:text;not null;comment:评论内容"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
}

func (o OrderComment) TableName() string {
	return "order_comment"
}
