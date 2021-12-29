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

type Geometry struct {
	Type string `json:"type"`
	Coordinates [2]float64 `json:"coordinates"`
}

func Populate(database *gorm.DB) {

	jsonFile, err := os.Open("data/scenes-test.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened scenes.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var featureCollection FeatureCollection
	
	json.Unmarshal(byteValue, &featureCollection)

	for _, f := range featureCollection.Features {
		data, _ := json.Marshal(f)
	 	fmt.Printf("json: %s\n\n", data)

		database.Create(f.GetModel())
	}
}
