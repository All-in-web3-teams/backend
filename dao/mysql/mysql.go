package mysql

import (
	"backend/models"
	"backend/settings"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.Table(models.User{}.TableName()).AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	err = db.Table(models.ContractInfo{}.TableName()).AutoMigrate(&models.ContractInfo{})
	if err != nil {
		return
	}
	return
}
