// models/setup.go

package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres dbname=anchors_testing port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&Scene{})

	DB = database
	seed(database)
}

func seed(database *gorm.DB) {
	scenes := []Scene{
		{Type: "Feature", CommonName: "oeverseehabenbecken"},
		{Type: "Feature", CommonName: "europahafenbecken"},
		{Type: "Feature", CommonName: "speicherXI"},
	}
	for _, s := range scenes {
		database.Create(&s)
	}
}