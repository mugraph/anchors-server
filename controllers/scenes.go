package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/janebuoy/anchors-server/models"
)

type CreateSceneInput struct {
	CommonName string `json:"common_name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

// GET /scenes
// Get all scenes

func FindScenes(c *gin.Context) {
	var scenes []models.Scene
	// Preload associated Models
	// Find Scenes
	models.DB.
	Preload("Properties.FlyToOptions").
	Preload("Layers").
	Preload("Content.Resources").
	Find(&scenes)

	data := make([]*models.SceneJSON, len(scenes))
	for i, s := range scenes {
		data[i] = s.GetJSON()
	}

	resp := struct {
		Type string `json:"type"`
		Scenes []*models.SceneJSON `json:"features"`
	} {
		Type: "FeatureCollection",
		Scenes: data,
	}

	c.JSON(http.StatusOK, resp)
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
	scene := models.Scene{CommonName: input.CommonName}
	models.DB.Create(&scene)

	c.JSON(http.StatusOK, scene)
}