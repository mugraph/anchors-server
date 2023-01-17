// models/geometry.go

package models

// Define Geometry in JSON format
type GeometryJSON struct {
	Type string `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Return Point DB Model from GeometryJSON
func (g *GeometryJSON) GetPointModel() Point {
	return Point{
		Latitude: g.Coordinates[0],
		Longitude: g.Coordinates[1],
	}
}

type Point struct {
	Latitude float64
	Longitude float64
}

// Return GeometryJSON from Point DB Model
func (p *Point) GetPointJSON() *GeometryJSON {
	return &GeometryJSON{
		Type: "Point",
		Coordinates: []float64{p.Latitude, p.Longitude},
	}
}