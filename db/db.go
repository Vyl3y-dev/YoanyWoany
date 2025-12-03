package db

import (
	"log"

	"YoanyWoany/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	// create or open your playground database
	DB, err = gorm.Open(sqlite.Open("playground.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ failed to connect database: %v", err)
	}

	// auto-create tables based on your structs
	err = DB.AutoMigrate(&models.Account{})
	if err != nil {
		log.Fatalf("❌ migration failed: %v", err)
	}

	log.Println("✨ Database ready kotence!")
}
