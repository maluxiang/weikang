package models

import (
	"gorm.io/gorm"
	"time"
	"weikang/global"
)

// Points 表示数据库中存储用户积分的模型。
type Points struct {
	ID        uint64         `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	UserID    int64          `gorm:"not null"`           // UserID 是与用户关联的外键
	Points    int64          `gorm:"not null;default:0"` // Points 存储用户的实际积分，初始默认为 0
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
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
	return global.DB.Where("id = ?", id).Find(&p).Limit(1).Error
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
