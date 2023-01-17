// models/properties.go

package models

import (
	"github.com/google/uuid"
)

// Define Properties DB Model
// Belongs to Scene
type Properties struct {
	ID uuid.UUID `json:"uuid" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	Title string `json:"title" example:"Architecture"`
	Description string `json:"description" example:"Places that showcase architecture in Nantes, such as the Château des ducs de Bretagne, Notre-Dame Cathedral and Passage Pommeraye."`
}

// Return JSON from Properties DB Model
func (p *Properties) GetPropertiesJSON() PropertiesJSON {
	return PropertiesJSON{
		ID: p.ID,
		Title: p.Title,
		Description: p.Description,
	}
}

type PropertiesJSON struct {
		ID uuid.UUID `json:"uuid"`	
    Title       string `json:"title"`
    Description string `json:"description"`
}

// Return Properties DB Model from PropertiesJSON
func (p *PropertiesJSON) GetPropertiesModel() *Properties {
	return &Properties{
		Title: p.Title,
		Description: p.Description,
	}
}