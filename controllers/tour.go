package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/anchors-server/models"
)

// Get /tour/:id

func FindTour(c *gin.Context) {

	var tour models.Tour
	var chapters []models.Chapter

	if err := models.DB.
		Preload("Properties").
		Where("id = ?", c.Param("id")).
		First(&tour).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tour not found!"})
		return
	}

	if err := models.DB.
		Preload("Properties").
		Where("tour_id = ?", c.Param("id")).
		Find(&chapters).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Chapter not found!"})
		return
	}

	features := make([]interface{}, len(chapters)+1)
	tourData := tour.GetTourJSON()

	for i := range features {
		if i == 0 {
			features[i] = tourData
		} else {
			features[i] = chapters[i-1].GetChapterJSON()
		}
	}

	resp := struct {
		Type     string        `json:"type"`
		Features []interface{} `json:"features"`
	}{
		Type:     "FeatureCollection",
		Features: features,
	}

	c.JSON(http.StatusOK, resp)
}
