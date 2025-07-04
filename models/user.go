package models

import (
	"gorm.io/gorm"
	"time"
	"weikang/global"
)

type Users struct {
	ID        uint64         `gorm:"column:id;type:bigint UNSIGNED;primaryKey;not null;" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
	Username  string         `gorm:"column:username;type:longtext;" json:"username"`
	Password  string         `gorm:"column:password;type:longtext;" json:"password"`
	Email     string         `gorm:"column:email;type:longtext;" json:"email"`
	Phone     string         `gorm:"column:phone;type:longtext;" json:"phone"`
}

// Account 账户模型，代表一个用户的账户信息
type Account struct {
	gorm.Model
	UserID        int64   // 用户ID
	Balance       float64 // 账户余额
	AccountNumber string  // 账户号码
	Currency      string  // 账户币种
}

func (u Users) TableName() string {
	return "users"
}
func (u Users) Create() error {
	return global.DB.Create(&u).Error
}
func (u Users) GetInfo() (user Users, err error) {
	return user, global.DB.First(&user, u).Error
}
func (a Account) TableName() string {
	return "accounts"
}
func (a Account) Create() error {
	return global.DB.Create(&a).Error
}
func (a Account) GetInfo() (account Account, err error) {
	return account, global.DB.First(&account, a).Error
}

func (a Account) UpdateBalance() error {
	return global.DB.Model(&a).Update("balance", gorm.Expr("balance + ?", a.Balance)).Error
}
