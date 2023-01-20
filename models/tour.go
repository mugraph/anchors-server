// models/tour.go

package models

import (
	"time"

	"github.com/google/uuid"
)

// Defines `Tour` database model. `Properties` belong to `Tour`.
type Tour struct {
	ID           uuid.UUID   `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	PropertiesID *uuid.UUID  `json:"properties_id" gorm:"type:uuid"`
	Properties   *Properties `json:"properties,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Point        Point       `json:"-" gorm:"embedded;embeddedPrefix:point_"`
}

// Define `TourJSON` representaion.
type TourJSON struct {
	ID         uuid.UUID       `json:"id"`
	Type       string          `json:"type"`
	Properties *PropertiesJSON `json:"properties"`
	Geometry   *GeometryJSON   `json:"geometry"`
}

// Returns `TourJSON` from its database model (`Tour`).
func (t *Tour) GetTourJSON() *TourJSON {
	properties := t.Properties.GetPropertiesJSON()
	return &TourJSON{
		ID:         t.ID,
		Type:       "Feature",
		Properties: &properties,
		Geometry:   t.Point.GetPointJSON(),
	}
}

// Returns `Tour` database model from its JSON representation (`TourJSON`).
func (t *TourJSON) GetTourModel() *Tour {
	return &Tour{
		ID:         t.ID,
		Properties: GetPropertiesModel(t.Properties, &t.ID),
		Point:      t.Geometry.GetPointModel(),
	}
}
