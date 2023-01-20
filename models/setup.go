// models/setup.go

package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres dbname=anchors port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("[setup.go] Failed to connect to database!", err)
	}
	database.AutoMigrate(&Tour{}, &Chapter{}, &Properties{})

	DB = database
	Populate(database)
}

func Populate(database *gorm.DB) {
	PopulateTours(database)
	PopulateChapters(database)
}

// Read FeatureCollection from tours.json and create corresponding
// Tour DB Entries
func PopulateTours(database *gorm.DB) {
	jsonFile, err := os.Open("data/tours.json")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("[setup.go] Successfully Opened tours.json")
	}

	defer jsonFile.Close()

	// Get the file's size
	fileInfo, _ := jsonFile.Stat()
	fileSize := fileInfo.Size()

	// Read the file into buffer as byte array
	buffer := make([]byte, fileSize)
	if _, err := io.ReadFull(jsonFile, buffer); err != nil {
		log.Fatal(err)
	}

	type TourCollection struct {
		Features []TourJSON `json:"features"`
	}
	var featureCollection TourCollection

	// Unmarshal byte array, containing our json content into 'features'
	json.Unmarshal(buffer, &featureCollection)

	// Add model representation of each feature to database
	for _, f := range featureCollection.Features {
		database.Create(f.GetTourModel())
	}
}

// Read FeatureCollection from chapters.json and create corresponding
// Chapter DB Entries
func PopulateChapters(database *gorm.DB) {
	// Open json file
	jsonFile, err := os.Open("data/chapters.json")
	// If we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("[setup.go] Successfully Opened chapters.json")
	}

	defer jsonFile.Close()

	// Get the file's size
	fileInfo, _ := jsonFile.Stat()
	fileSize := fileInfo.Size()

	// Read the file into buffer as byte array
	buffer := make([]byte, fileSize)
	if _, err := io.ReadFull(jsonFile, buffer); err != nil {
		log.Fatal(err)
	}

	type ChapterCollection struct {
		Features []ChapterJSON `json:"features"`
	}
	var featureCollection ChapterCollection

	// Unmarshal byte array, containing our json content into 'features'
	json.Unmarshal(buffer, &featureCollection)

	// Iterate over all features
	for _, f := range featureCollection.Features {
		// Check if feature/chapter with id already exists
		// Else add new feature/chapter
		id := f.GetChapterId()

		fmt.Printf("[setup.go] Creating chapter for new ID %s. \n", id)

		// Add model representation of each feature to database
		database.Create(f.GetChapterModel())

	}
}
