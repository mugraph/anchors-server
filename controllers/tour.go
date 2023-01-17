package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/anchors-server/models"
)

// Get /tour/:id

func FindTour(c *gin.Context) {

	var tour models.Tour

	if err := models.DB.
	Preload("Properties").
	Where("id = ?", c.Param("id")).
	First(&tour).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Tour not found!"})
		return
	}

	data := tour.GetTourJSON()

	var chapters []models.Chapter

	if err := models.DB.
	Preload("Properties").
	Where("tour_id = ?", c.Param("id")).
	Find(&chapters).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Chapter not found!"})
		return
	}

	chaptersData := make([]*models.ChapterJSON, len(chapters))
	for i, c := range chapters {
		chaptersData[i] = c.GetChapterJSON()
	}

	data.Chapters = chaptersData

	c.JSON(http.StatusOK, data)
}