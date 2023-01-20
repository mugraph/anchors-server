package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/anchors-server/models"
)

type CreateChapterInput struct {
	Type string `json:"type" binding:"required"`
}

// GET /chapters
// Get all chapters

func FindChapters(c *gin.Context) {
	var chapters []models.Chapter
	// Preload associated Models
	// Find Chapters
	models.DB.
		Preload("Properties").
		Find(&chapters)

	data := make([]*models.ChapterJSON, len(chapters))
	for i, c := range chapters {
		data[i] = c.GetChapterJSON()
	}

	resp := struct {
		Type     string                `json:"type"`
		Chapters []*models.ChapterJSON `json:"features"`
	}{
		Type:     "FeatureCollection",
		Chapters: data,
	}

	c.JSON(http.StatusOK, resp)
}

// POST /chapters
// Create new chapter
func CreateChapter(c *gin.Context) {
	// Validate input
	var input CreateChapterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Create Chapter
	chapter := models.Chapter{}
	models.DB.Create(&chapter)

	c.JSON(http.StatusOK, chapter)
}
