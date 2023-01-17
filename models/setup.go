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
	dsn := "host=localhost user=postgres dbname=anchors-testing port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("[setup.go] Failed to connect to database!", err)
	}
	database.AutoMigrate(&Tour{}, &Chapter{}, &Properties{})

	DB = database
	Populate(database)
}

type TourCollection struct {
	Features []TourJSON `json:"features"`
}

type ChapterCollection struct {
	Features []ChapterJSON `json:"features"`
}

func Populate(database *gorm.DB) {
	PopulateTours(database)
	PopulateChapters(database)
}

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

	// Read the opened file as byte array
	buffer := make([]byte, fileSize)
	if _, err := io.ReadFull(jsonFile, buffer); err != nil {
		log.Fatal(err)
	}

	// Initialize 'featureCollection' array
	var featureCollection TourCollection
	
	// Unmarshal byte array, containing our json content into 'featureCollection'
	json.Unmarshal(buffer, &featureCollection)

	for _, f := range featureCollection.Features {
		// Marshal features back to json and print them 
		// data, _ := json.Marshal(f)
	 	// fmt.Printf("json: %s\n\n", data)
		// Add model representation of each feature to database
		database.Create(f.GetTourModel())
	}

}

func PopulateChapters(database *gorm.DB) {
		// Open json file
	jsonFile, err := os.Open("data/chapters.json")
	// If we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("[setup.go] Successfully Opened chapters.json")
	}
	
	// Defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// Get the file's size
	fileInfo, _ := jsonFile.Stat()
	fileSize := fileInfo.Size()

	// Read the opened file as byte array
	buffer := make([]byte, fileSize)
	if _, err := io.ReadFull(jsonFile, buffer); err != nil {
		log.Fatal(err)
	}

	// Initialize 'featureCollection' array
	var featureCollection ChapterCollection
	
	// Unmarshal byte array, containing our json content into 'featureCollection'
	json.Unmarshal(buffer, &featureCollection)

	// Iterate over all features in 'featureCollection'
	for _, f := range featureCollection.Features {
		
		// Check if feature/chapter with uuid already exists
		// Else add new feature/chapter
		uuid := f.GetChapterUuid()
		exists := false
		err = database.Model(&Chapter{}).
			Select("count(*) > 0").
			Where("id = ?", uuid).
			Find(&exists).
			Error
		if err != nil {
			log.Fatal(err)
		}

		if exists {
			fmt.Printf("[setup.go] Chapter with ID %s already exists and is not created. \n", uuid)
		} else {
			fmt.Printf("[setup.go] Creating chapter for new ID %s. \n", uuid)
			// Add model representation of each feature to database
			database.Create(f.GetChapterModel())
		}
	}
}
