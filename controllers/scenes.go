package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/janebuoy/anchors-server/models"
	"github.com/google/uuid"
	"strconv"
)

type CreateSceneInput struct {
	CommonName string `json:"common_name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type FeatureCollection struct {
	Type string `json:"type"`
	Scenes []*models.SceneJSON `json:"features"`
}

// MakeJSON()
func MakeJSON(scenes []models.Scene) FeatureCollection {
	data := make([]*models.SceneJSON, len(scenes))
	for i, s := range scenes {
		data[i] = s.GetJSON()
	}

	return FeatureCollection {
		Type: "FeatureCollection",
		Scenes: data,
	}
}

// GET /scenes
// Get all scenes
func FindScenes(c *gin.Context) {
		var scenes []models.Scene
		// Preload associated Models
		// Find Scenes
		db := models.DB.
		Preload("Properties.FlyToOptions").
		Preload("Layers.LayerOptions").
		Preload("Content.Resources")

		if v, ok := c.GetQuery("is_main"); ok {
			b, err := strconv.ParseBool(v)
			if err != nil {
				c.String(http.StatusBadRequest, "is_main is not a boolean: %s", err)
				return
			}
			db = db.Where("is_main", b)
		}
		if v, ok := c.GetQuery("id"); ok {
			uuid, err := uuid.Parse(v)
			if err != nil {
				c.String(http.StatusBadRequest, "id is not a valid uuid: %s", err)
				return
			}
			db = db.Where("id", uuid)
		}
		
		result := db.Find(&scenes)
		if err := result.Error;	err != nil {
			c.String(http.StatusInternalServerError, "database error: %s", err)
			return
		}
		if result.RowsAffected == 0 {
			c.String(http.StatusNotFound, "no scenes found")
			return
		}

		resp := MakeJSON(scenes)

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