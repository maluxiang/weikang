package models

import (
	"gorm.io/gorm"
	"weikang/global"
)

// Points 表示数据库中存储用户积分的模型。
type Points struct {
	gorm.Model       // gorm.Model 包含 ID、CreatedAt、UpdatedAt、DeletedAt 字段，由 GORM 自动管理
	UserID     int64 `gorm:"not null"` // UserID 是与用户关联的外键
	Points     int64 `gorm:"not null;default:0"` // Points 存储用户的实际积分，初始默认为 0
}

func (p Points) TableName() string {
	return "points"
}

func (p Points) GetInfo() (points Points, err error) {
	return points, global.DB.First(&points, p).Error
}

// Create 在数据库中创建新的用户积分记录。
func (p Points) Create() error {
	return global.DB.Create(&p).Error
}

// Update 更新数据库中现有的用户积分记录。
func (p Points) Update() error {
	return global.DB.Save(&p).Error
}

// Delete 从数据库中删除用户积分记录。
func (p Points) Delete() error {
	return global.DB.Delete(p).Error
}

// FindByID 根据ID从数据库检索用户积分记录。
func (p Points) FindByID(id uint) error {
	return global.DB.First(p, id).Error
}

// GetAll 获取所有用户积分记录。
func (p Points) GetAll() ([]Points, error) {
	var points []Points
	
	// 这里使用Preload方法来预加载用户信息，避免出现N+1查询问题。
	// 如果你不使用Preload方法，你需要自己编写查询用户信息的逻辑。
	err := global.DB.Find(&points).Error
	if err != nil {
		return nil, err
	}
	
	return points, nil
}
