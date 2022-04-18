package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"learn/pkg/setting"
	"time"
)

//Model 嵌入结构体，精简代码
type Model struct {
	ID        uint           `gorm:"primaryKey"`
	UpdatedAt int            // 在创建时该字段值为零或更新时使用当前时间戳秒数填充
	CreatedAt int            // 在创建时该字段值为零时使用当前时间戳秒数填充
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var (
	db        *gorm.DB
	err       error
	ormLogger logger.Interface
)

func init() {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", setting.DBUser, setting.DBPasswd, setting.DBName, setting.DBHost, setting.DBPort)
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: connString,
	}), &gorm.Config{
		Logger: ormLogger,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	err = db.AutoMigrate(&User{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&Log{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&Collection{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&Handling{})
	if err != nil {
		return
	}

}
