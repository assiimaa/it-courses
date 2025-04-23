package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=assima05 dbname=itcourses port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	//if migrateErr := database.AutoMigrate(
	//	&models.User{},
	//	&models.Course{},
	//	&models.Category{},
	//	&models.Instructor{},
	//); migrateErr != nil {
	//	log.Fatalf("Failed to migrate database: %v", migrateErr)
	//	return
	//}

	DB = database
}
