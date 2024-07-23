package start

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"weikang/global"
	"weikang/models"
)

func Mysql() {
	db, err := gorm.Open(mysql.Open(global.NacosConfig.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		zap.S().Error("MySQL连接失败", err)
		return
	}
	global.DB = db
	err = db.AutoMigrate(&models.User{}, &models.SmartDoctor{}, &models.UploadFile{}, &models.HealthData{})
	if err != nil {
		zap.S().Error("MySQL迁移失败", err)
		return
	}
}
