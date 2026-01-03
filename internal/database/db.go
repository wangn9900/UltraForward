package database

import (
	"log"

	"github.com/wangn9900/UltraForward/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
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
}
