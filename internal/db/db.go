package db

import (
	"log"
	"os"
	"time"

	"github.com/NemishGorasiya/Go-Todo/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get db instance: ", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Auto-migrate your models
	err = DB.AutoMigrate(&model.Todo{})
	if err != nil {
		log.Fatal("AutoMigrate failed: ", err)
	}

	log.Println("✅ Connected to database")
}
