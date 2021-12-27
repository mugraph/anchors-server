package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/janebuoy/anchors-server/models"
	//"github.com/twpayne/go-geom/encoding/geojson"
)

type CreateSceneInput struct {
	CommonName string `json:"common_name" binding:"required"`
	Type string `json:"type" binding:"required"`
	//Geometry geojson.Geometry `json:"geometry" binding:"required"`
}

// GET /scenes
// Get all scenes

func FindScenes(c *gin.Context) {
	var scenes []models.Scene
	models.DB.Find(&scenes)

	c.JSON(http.StatusOK, gin.H{"data": scenes})
}

// GET /properties
// Get all properties

func FindProperties(c *gin.Context) {
	var properties []models.Properties
	models.DB.Find(&properties)

	c.JSON(http.StatusOK, gin.H{"data": properties})
}

// POST /scenes
// Create new scene
func CreateScene(c *gin.Context) {
	// Validate input
	var input CreateSceneInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Create Scene
	scene := models.Scene{CommonName: input.CommonName, Type: input.Type}
	models.DB.Create(&scene)

	c.JSON(http.StatusOK, gin.H{"data": scene})
}