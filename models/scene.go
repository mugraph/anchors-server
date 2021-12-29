// models/scene.go

package models

import (
		"github.com/google/uuid"
)

type Scene struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	CommonName string `json:"common_name" example:"oeverseehabenbecken"`
	PropertiesID *uuid.UUID `json:"-" gorm:"type:uuid"`
	Properties *Properties `json:"properties,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Point Point `json:"-" gorm:"embedded;embeddedPrefix:point_"`
	Layers []*Layer  `json:"layers" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ContentID *uuid.UUID `json:"-" gorm:"type:uuid"`
	Content *Content `json:"content" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (s *Scene) GetJSON() *SceneJSON {
	return &SceneJSON{
		ID: s.ID,
		Type: "Feature",
		CommonName: s.CommonName,
		Properties: s.Properties,
		Geometry: s.Point.GetJSON(),
		Layers: s.Layers,
		Content: s.Content,
	}
}

type SceneJSON struct {
	ID uuid.UUID `json:"id"`
	Type string `json:"type"`
	CommonName string `json:"common_name" example:"oeverseehabenbecken"`
	Geometry *GeometryJSON `json:"geometry"`
	Properties *Properties `json:"properties"`
	Layers []*Layer `json:"layers"`
	Content *Content `json:"content"`
}

func (s *SceneJSON) GetModel() *Scene {
	return &Scene{
		ID: s.ID,
		CommonName: s.CommonName,
		Properties: s.Properties,
		Point: s.Geometry.GetPoint(),
		Layers: s.Layers,
		Content: s.Content,
	}
}

type Point struct {
	X float64
	Y float64
}

func (p *Point) GetJSON() *GeometryJSON {
	return &GeometryJSON{
		Type: "Point",
		Coordinates: []float64{p.X, p.Y},
	}
}

type GeometryJSON struct {
	Type string `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func (g *GeometryJSON) GetPoint() Point {
	return Point{
		X: g.Coordinates[0],
		Y: g.Coordinates[1],
	}
}


type Properties struct {
	ID uuid.UUID `json:"-" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	Title string `json:"title" example:"Överseehabenbecken"`
	Zoom uint `json:"zoom" example:"17"`
	FlyTo bool `json:"flyTo"`
	FlyToOptionsID *uuid.UUID `json:"-" gorm:"type:uuid"`
	FlyToOptions *FlyToOptions `json:"flyToOptions,omitempty", gorm:"contraint:OnUpdate:CASCASE,OnDelete:CASCADE;"`
}

type FlyToOptions struct {
	ID uuid.UUID `json:"-" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	Duration float64 `json:"duration" gorm:"default:1.5" example:"1.5"`
	EaseLinearity float64 `json:"easeLinearity" gorm:"default:0.2" example:"0.2"`
}

type Layer struct {
	ID uuid.UUID `json:"-" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	Name string `json:"name"`
	SceneID *uuid.UUID `json:"-" gorm:"type:uuid"`
	LayerOptionsID *uuid.UUID `json:"-" gorm:"type:uuid"`
	LayerOptions *LayerOptions `json:"layerOptions,omitempty", gorm:"contraint:OnUpdate:CASCASE,OnDelete:CASCADE;"`
}

type LayerOptions struct {
	ID uuid.UUID `json:"-" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	Source string `json:"source"`
	Type string `json"type"`
	Short string `json:"short"`
	Selector bool `json:"selector"`
	Info bool `json:"info"`
}

type Content struct {
	ID uuid.UUID `json:"-" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	Title string `json:"title"`
	ImageSource string `json:"imageSrc"`
	Description string `json:"description"`
	Resources []*Resource `json:"resources"`
}

type Resource struct {
	ID uuid.UUID `json:"-" gorm:"type:uuid;default:gen_random_uuid()"`
	// attributes
	ContentID *uuid.UUID `json:"-" gorm:"type:uuid"`
	IDx uint `json:"id"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
	AudioSource string `json:"audioSrc,omitempty"`
	Type string `json:"type"`
	Length uint `json:"length,omitempty"`
	TargetUUID *uuid.UUID `json:"targetUUID,omitempty"`
}
