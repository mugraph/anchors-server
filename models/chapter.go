// models/chapter.go

package models

import (
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mugraph/anchors-server/internal"
)

// Define Chapter DB model
// Properties belongs to Chapter
type Chapter struct {
	ID uuid.UUID `json:"uuid" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CommonName string `json:"common_name"`
	PropertiesID *uuid.UUID `json:"properties_id" gorm:"type:uuid"`
	Properties *Properties `json:"properties,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Point Point `json:"-" gorm:"embedded;embeddedPrefix:point_"`
	TourID uuid.UUID `json:"-" gorm:"type:uuid;default:gen_random_uuid();constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Return JSON from Chapter DB Model
func (c *Chapter) GetChapterJSON() *ChapterJSON {
	properties := c.Properties.GetPropertiesJSON()
	return &ChapterJSON{
		ID: c.ID,
		Type: "Feature",
		CommonName: c.CommonName,
		Properties: &properties,
		Geometry: c.Point.GetPointJSON(),
		TourID: c.TourID,
	}
}

// Define Chapter in JSON format
type ChapterJSON struct {
	ID uuid.UUID `json:"uuid"`
	Type string `json:"type"`
	CommonName string `json:"common_name"`
	Properties *PropertiesJSON `json:"properties"`
	Geometry *GeometryJSON `json:"geometry"`
	TourID uuid.UUID `json:"tour_id"`
}

// Return Chapter JSON uuid
func (c *ChapterJSON) GetChapterUuid() uuid.UUID {
	return c.ID
}

// Return Chapter DB Model from ChapterJSON
func (c *ChapterJSON) GetChapterModel() *Chapter {
	return &Chapter{
		ID: c.ID,
		CommonName: c.GetSanitizedCommonName(),
		Properties: c.Properties.GetPropertiesModel(),
		Point: c.Geometry.GetPointModel(),
		TourID: c.TourID,
	}
}

func (c *ChapterJSON) GetSanitizedCommonName() string {
	if len(c.CommonName) > 0 {
		return c.CommonName
	} else {
		title := c.Properties.GetPropertiesModel().Title
		clean := internal.RemoveAccents(title)
		str := url.PathEscape(strings.ToLower(clean))
		re := regexp.MustCompile("%[0-9A-Fa-f]{2}")
		s := re.ReplaceAllString(str, "-")
		return s
	}
}