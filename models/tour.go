// models/tour.go

package models

import (
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mugraph/anchors-server/internal"
)

// Define Tour DB Model
// Chapters belong to Tour
type Tour struct {
	ID uuid.UUID `json:"uuid" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CommonName string `json:"common_name"`
	PropertiesID *uuid.UUID `json:"properties_id" gorm:"type:uuid"`
	Properties *Properties `json:"properties,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Chapters []Chapter `json:"chapters"`
	Point Point `json:"-" gorm:"embedded;embeddedPrefix:point_"`
}

// func (t *Tour) convertChapters() []*ChapterJSON {
//     var chapterJSON []*ChapterJSON
//     for _, c := range t.Chapters {
//         chapterJSON = append(chapterJSON, c.GetChapterJSON())
//     }
//     return chapterJSON
// }

// Return JSON form Tour DB Model
func (t *Tour) GetTourJSON() *TourJSON {
	properties := t.Properties.GetPropertiesJSON()
	return &TourJSON{
		ID: t.ID,
		Type: "Feature",
		CommonName: t.CommonName,
		Properties: &properties,
		Geometry: t.Point.GetPointJSON(),
		// Chapters: t.convertChapters(),
	}
}

// Define Tour in JSON format
type TourJSON struct {
	ID uuid.UUID `json:"uuid"`
	Type string `json:"type"`
	CommonName string `json:"common_name"`
	Properties *PropertiesJSON `json:"properties"`
	Geometry *GeometryJSON `json:"geometry"`
	Chapters []*ChapterJSON `json:"chapters,omitempty"`
}

func (t *TourJSON) GetTourModel() *Tour {
	return &Tour{
		ID: t.ID,
		CommonName: t.GetSanitizedCommonName(),
		Properties: t.Properties.GetPropertiesModel(),
		Point: t.Geometry.GetPointModel(),
	}
}

func (t *TourJSON) GetSanitizedCommonName() string {
	if len(t.CommonName) > 0 {
		return t.CommonName
	} else {
		title := t.Properties.GetPropertiesModel().Title
		clean := internal.RemoveAccents(title)
		str := url.PathEscape(strings.ToLower(clean))
		re := regexp.MustCompile("%[0-9A-Fa-f]{2}")
		s := re.ReplaceAllString(str, "-")
		return s
	}
}