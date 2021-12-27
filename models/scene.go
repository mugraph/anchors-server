// models/scene.go

package models

import (
		"github.com/google/uuid"
)

type Scene struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	Type       string `json:"type" gorm:"default:Feature" example:"Feature"`
	CommonName string `json:"common_name" example:"oeverseehabenbecken"`
	Properties Properties `json:"properties" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Properties struct {
	// attributes
	Title string `json:"title" example:"Överseehabenbecken"`
	Zoom uint `json:"zoom" example:"17"`
	FlyTo bool `json:"flyTo"`
	SceneID uuid.UUID 
}


