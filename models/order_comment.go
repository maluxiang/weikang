package models

import (
	"gorm.io/gorm"
	"time"
	"weikang/global"
)

// 订单评论表 - 支持无限评论
type OrderComment struct {
	ID          uint64         `gorm:"primaryKey;autoIncrement"`
	OrderID     uint64         `gorm:"not null;index;comment:订单ID"`
	UserID      uint64         `gorm:"not null;index;comment:用户ID"`
	Content     string         `gorm:"type:text;not null;comment:评论内容"`
	Rating      int            `gorm:"type:int;default:5;comment:评分(1-5星)"`
	Images      string         `gorm:"type:text;comment:评论图片，JSON格式存储"`
	ReplyTo     uint64         `gorm:"default:0;comment:回复的评论ID，0表示顶级评论"`
	IsAnonymous bool           `gorm:"default:false;comment:是否匿名评论"`
	Status      int            `gorm:"type:int;default:1;comment:评论状态(1-正常,2-隐藏,3-删除)"`
	Likes       int            `gorm:"type:int;default:0;comment:点赞数"`
	CreatedAt   time.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index"`

	// 关联字段
	User  Users `gorm:"foreignKey:UserID;references:ID"`
	Order Order `gorm:"foreignKey:OrderID;references:ID"`
}

// 评论点赞表
type CommentLike struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement"`
	CommentID uint64         `gorm:"not null;index;comment:评论ID"`
	UserID    uint64         `gorm:"not null;index;comment:用户ID"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index"`
}

func (o OrderComment) TableName() string {
	return "order_comment"
}

func (c CommentLike) TableName() string {
	return "comment_like"
}

// 获取订单的所有评论
func (o OrderComment) GetOrderComments(orderID uint64, page, pageSize int) ([]OrderComment, int64, error) {
	var comments []OrderComment
	var total int64

	query := global.DB.Model(&OrderComment{}).Where("order_id = ? AND status = 1", orderID)

	// 获取总数
	query.Count(&total)

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Preload("User").Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&comments).Error

	return comments, total, err
}

// 获取评论的回复
func (o OrderComment) GetCommentReplies(commentID uint64, page, pageSize int) ([]OrderComment, int64, error) {
	var replies []OrderComment
	var total int64

	query := global.DB.Model(&OrderComment{}).Where("reply_to = ? AND status = 1", commentID)

	// 获取总数
	query.Count(&total)

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Preload("User").Offset(offset).Limit(pageSize).Order("created_at ASC").Find(&replies).Error

	return replies, total, err
}

// 点赞评论
func (c CommentLike) LikeComment(commentID, userID uint64) error {
	// 检查是否已经点赞
	var existingLike CommentLike
	err := global.DB.Where("comment_id = ? AND user_id = ?", commentID, userID).First(&existingLike).Error
	if err == nil {
		// 已经点赞，取消点赞
		global.DB.Delete(&existingLike)
		// 减少评论点赞数
		global.DB.Model(&OrderComment{}).Where("id = ?", commentID).Update("likes", gorm.Expr("likes - 1"))
		return nil
	}

	// 创建点赞记录
	like := CommentLike{
		CommentID: commentID,
		UserID:    userID,
	}
	err = global.DB.Create(&like).Error
	if err != nil {
		return err
	}

	// 增加评论点赞数
	return global.DB.Model(&OrderComment{}).Where("id = ?", commentID).Update("likes", gorm.Expr("likes + 1")).Error
}

// 检查用户是否已点赞
func (c CommentLike) IsLiked(commentID, userID uint64) bool {
	var like CommentLike
	err := global.DB.Where("comment_id = ? AND user_id = ?", commentID, userID).First(&like).Error
	return err == nil
}
