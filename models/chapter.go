// models/chapter.go

package models

import (
	"time"

	"github.com/google/uuid"
)

// Define `Chapter` database DB model. `Properties` belong to `Chapter`.
type Chapter struct {
	ID           uuid.UUID   `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	CommonName   string      `json:"common_name"`
	PropertiesID *uuid.UUID  `json:"properties_id" gorm:"type:uuid"`
	Properties   *Properties `json:"properties,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Point        Point       `json:"-" gorm:"embedded;embeddedPrefix:point_"`
	TourID       uuid.UUID   `json:"-" gorm:"type:uuid;default:gen_random_uuid();constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Defines `ChapterJSON` representaion.
type ChapterJSON struct {
	ID         uuid.UUID       `json:"id"`
	Type       string          `json:"type"`
	Properties *PropertiesJSON `json:"properties"`
	Geometry   *GeometryJSON   `json:"geometry"`
	TourID     uuid.UUID       `json:"tour_id"`
}

// Returns `ChapterJSON` from its database model (`Chapter`).
func (c *Chapter) GetChapterJSON() *ChapterJSON {
	properties := c.Properties.GetPropertiesJSON()
	return &ChapterJSON{
		ID:         c.ID,
		Type:       "Feature",
		Properties: &properties,
		Geometry:   c.Point.GetPointJSON(),
		TourID:     c.TourID,
	}
}

// Returns chapter id.
func (c *ChapterJSON) GetChapterId() uuid.UUID {
	return c.ID
}

// Returns `Chapter` database model from its JSON representation (`ChapterJSON`).
func (c *ChapterJSON) GetChapterModel() *Chapter {
	return &Chapter{
		ID:         c.ID,
		Properties: GetPropertiesModel(c.Properties, &c.ID),
		Point:      c.Geometry.GetPointModel(),
		TourID:     c.TourID,
	}
}
