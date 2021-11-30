package database

import (
	"log"
	"os"

	"github.com/abiiranathan/goclinic/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

var config = &gorm.Config{
	Logger:                                   logger.Default.LogMode(logger.Warn),
	DisableForeignKeyConstraintWhenMigrating: false,
}

func ConnectToDatabase() {
	DSN := os.Getenv("DSN")
	var err error

	DB, err = gorm.Open(postgres.Open(DSN), config)

	if err != nil {
		log.Fatal("Database connection failed!")
	}

	log.Printf("Connected to database...")
	DB.AutoMigrate(&models.User{}, &models.Permission{})

}
