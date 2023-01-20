// models/properties.go

package models

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
	"github.com/mugraph/anchors-server/internal"
	"gorm.io/gorm"
)

// Defines `Properties` database model. `Properties` belong to either `Chapter` or `Tour` as defined on the parent.
type Properties struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	CommonName  string    `json:"common_name"`
	Title       string    `json:"title" example:"Architecture"`
	Description string    `json:"description" example:"Places in Nantes that,..."`
	Hash        string    `json:"-" gorm:"unique_index"`
}

// Defines `PropertiesJSON` representaion.
type PropertiesJSON struct {
	ID          uuid.UUID `json:"id"`
	CommonName  string    `json:"common_name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Hash        string    `json:"-"`
}

// Calculates and adds a sum265 hash on `Properties` before creation.
func (p *Properties) BeforeCreate(database *gorm.DB) (err error) {
	h := sha256.New()
	h.Write([]byte(p.Title))
	h.Write([]byte(p.Title))
	p.Hash = hex.EncodeToString(h.Sum(nil))
	return nil
}

// Calculates and adds a sum265 hash on `Properties` before save.
func (p *Properties) BeforeSave(database *gorm.DB) (err error) {
	h := sha256.New()
	h.Write([]byte(p.Title))
	h.Write([]byte(p.Title))
	p.Hash = hex.EncodeToString(h.Sum(nil))
	return nil
}

// Returns `PropertiesJSON` from its database model (`Properties`).
func (p *Properties) GetPropertiesJSON() PropertiesJSON {
	return PropertiesJSON{
		ID:          p.ID,
		CommonName:  p.CommonName,
		Title:       p.Title,
		Description: p.Description,
	}
}

// Returns `Properties` database model from its JSON representaion (`PropertiesJSON`). The `id` is set by parent `id` to facilitate usage on client-side. `common_name` is derived from the title and is an all lower-case string.
func GetPropertiesModel(p *PropertiesJSON, id *uuid.UUID) *Properties {
	return &Properties{
		ID:          *id,
		CommonName:  internal.GetSanitizedCommonName(p.Title),
		Title:       p.Title,
		Description: p.Description,
		Hash:        p.Hash,
	}
}
