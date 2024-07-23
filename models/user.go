package models

import (
	"gorm.io/gorm"
	"weikang/global"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
	Phone    string
}

// Account 账户模型，代表一个用户的账户信息
type Account struct {
	gorm.Model
	UserID        int64   // 用户ID
	Balance       float64 // 账户余额
	AccountNumber string  // 账户号码
	Currency      string  // 账户币种
}

func (u User) TableName() string {
	return "users"
}
func (u User) Create() error {
	return global.DB.Create(&u).Error
}
func (u User) GetInfo() (user User, err error) {
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
