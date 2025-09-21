package GeometryConstructor

import "fmt"

type PointZM struct {
	coordinates Coordinate
}

func (p *PointZM) WKT() string {
	c := p.coordinates
	return fmt.Sprintf("POINT ZM (%s)", c.String())
}
func (p *PointZM) GeoJSON() string           { return "" }
func (p *PointZM) GeometryType() string      { return "POINT ZM" }
func (p *PointZM) String() string            { return p.WKT() }
func (p *PointZM) GetCoordinate() Coordinate { return p.coordinates }
