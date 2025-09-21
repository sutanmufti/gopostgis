package GeometryConstructor

import "fmt"

type Point struct {
	coordinates Coordinate
}

func (p *Point) WKT() string {
	c := p.coordinates
	return fmt.Sprintf("POINT (%s)", c.String())
}
func (p *Point) GeoJSON() string           { return "" }
func (p *Point) GeometryType() string      { return "POINT" }
func (p *Point) String() string            { return p.WKT() }
func (p *Point) GetCoordinate() Coordinate { return p.coordinates }
func (p *Point) GetDim() string            { return "" }
func (p *Point) GetCoordinates() []Coordinate {
	return []Coordinate{p.coordinates}
}
