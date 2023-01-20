package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mugraph/anchors-server/models"
	"github.com/paulmach/orb"
)

// Get /tour/:id

type FeatureCollection struct {
	Type     string        `json:"type"`
	BBox		 []float64		 `json:"bbox"`
	Features []interface{} `json:"features"`
}

func calculateBoundingBox(chapters []models.Chapter) []float64 {
	var points []orb.Point

	for i := range chapters {
		lat := chapters[i].GetChapterJSON().Geometry.GetPointModel().Latitude
		lon := chapters[i].GetChapterJSON().Geometry.GetPointModel().Longitude
		points = append(points, orb.Point{lon, lat})
	}
	
	// Create a MultiPoint from the point geometries
	multiPoint := orb.MultiPoint(points)

	// Calculate the bounding box of the MultiPoint
	bBox := multiPoint.Bound()
	bounds := []float64{bBox.Min.Y(), bBox.Min.X(), bBox.Max.Y(), bBox.Max.X()}

	return bounds

}

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

	resp := FeatureCollection{
		Type:     "FeatureCollection",
		BBox: 		calculateBoundingBox(chapters),
		Features: features,
	}

	c.JSON(http.StatusOK, resp)
}
