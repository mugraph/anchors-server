package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/anchors-server/models"
)

// GET /tours

func FindTours(c *gin.Context) {
	var tours []models.Tour
	models.DB.
	Preload("Properties").
	Preload("Chapters").
	Find(&tours)

	data := make([]*models.TourJSON, len(tours))
	for i, t := range tours {
		data[i] = t.GetTourJSON()
	}

	resp := struct {
		Type string `json:"type"`
		Tours []*models.TourJSON `json:"features"`
	} {
		Type: "FeatureCollection",
		Tours: data,
	}

	c.JSON(http.StatusOK, resp)
}