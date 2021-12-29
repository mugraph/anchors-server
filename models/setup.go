// models/setup.go

package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=janne dbname=anchors-testing port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	database.AutoMigrate(&Scene{}, &Properties{}, &FlyToOptions{}, &Layer{}, &LayerOptions{}, &Content{}, &Resource{})

	DB = database
	Populate(database)
}

type FeatureCollection struct {
	Features []SceneJSON `json:"features"`
}

func Populate(database *gorm.DB) {

	// Open json file
	jsonFile, err := os.Open("data/scenes-test.json")
	// If we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened scenes.json")
	
	// Defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// Read the opened file as byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialize 'featureCollection' array
	var featureCollection FeatureCollection
	
	// Unmarshal byte array, containing our json content into 'featureCollection'
	json.Unmarshal(byteValue, &featureCollection)

	// Iterate over all features in 'featureCollection'
	for _, f := range featureCollection.Features {
		// Marshal features back to json and print them 
		data, _ := json.Marshal(f)
	 	fmt.Printf("json: %s\n\n", data)

		// Add Model Representation of each feature or 'scene' to Database
		database.Create(f.GetModel())
	}
}
