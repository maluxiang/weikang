package logic

import (
	"time"
	"weikang/global"
	"weikang/models"
)

// 发送站内信
func (s Server) SendSiteMessage(userID int64, title, content string) error {
	msg := models.SiteMessage{
		UserID:    userID,
		Title:     title,
		Content:   content,
		IsRead:    0,
		CreatedAt: time.Now(),
	}
	return global.DB.Create(&msg).Error
}

// 获取消息列表
func GetSiteMessages(userID int64) ([]models.SiteMessage, error) {
	var msgs []models.SiteMessage
	err := global.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&msgs).Error
	return msgs, err
}

// 标记已读
func MarkMessageRead(id int64) error {
	return global.DB.Model(&models.SiteMessage{}).Where("id = ?", id).Update("is_read", 1).Error
}
