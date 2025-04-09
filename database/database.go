package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"it-courses/models"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=assima05 dbname=itcourses port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = database.AutoMigrate(&models.Course{}, &models.Category{}, &models.Instructor{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	DB = database
}
