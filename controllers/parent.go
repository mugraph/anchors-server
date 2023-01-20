package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/anchors-server/models"
)

// Get /parent/:id

func FindParent(c *gin.Context) {
	var parent models.Chapter
	if err := models.DB.
		Preload("Properties").
		Where("properties_id = ?", c.Param("id")).
		First(&parent).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parent chapter not found!"})
		return
	}
	data := parent.GetChapterJSON()

	c.JSON(http.StatusOK, data)
}
