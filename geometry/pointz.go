package GeometryConstructor

import "fmt"

type PointZ struct {
	coordinates Coordinate
}

func (p *PointZ) WKT() string {
	c := p.coordinates
	return fmt.Sprintf("POINT Z (%s)", c.String())
}
func (p *PointZ) GeoJSON() string           { return "" }
func (p *PointZ) GeometryType() string      { return "POINT Z" }
func (p *PointZ) String() string            { return p.WKT() }
func (p *PointZ) GetCoordinate() Coordinate { return p.coordinates }
