package models

import (
	"fashora-backend/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.AppConfig.PostgresHost,
		config.AppConfig.PostgresUser,
		config.AppConfig.PostgresPassword,
		config.AppConfig.PostgresDB,
		config.AppConfig.PostgresPort,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = database.AutoMigrate(&Users{}, &Stores{}, &Item{})
	if err != nil {
		return
	}
	DB = database
}
