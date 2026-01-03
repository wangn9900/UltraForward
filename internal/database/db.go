package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/wangn9900/UltraForward/internal/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// 使用纯 Go 版 SQLite 驱动 (github.com/glebarez/sqlite)
	// 解决 CGO_ENABLED=0 下的编译运行问题
	DB, err = gorm.Open(sqlite.Open("ultra.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	DB.AutoMigrate(
		&models.User{},
		&models.Plan{},
		&models.UltraNode{},
		&models.Line{},
		&models.Rule{},
		&models.Order{},
	)

	seedData()
}

func seedData() {
	var count int64
	DB.Model(&models.Plan{}).Count(&count)
	if count == 0 {
		DB.Create(&models.Plan{
			Name: "旗舰企业包", Price: 49.9, Traffic: 500, Rules: 20, DurationDays: 30, Description: "全速、全线路、SLA 保证",
		})
	}
}
