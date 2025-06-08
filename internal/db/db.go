package db

import (
	"log"

	"devlink/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbURL string) *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Resource{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	return DB
}
