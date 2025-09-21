package GeometryConstructor

import "fmt"

type PointM struct {
	coordinates Coordinate
}

func (p *PointM) WKT() string {
	c := p.coordinates
	return fmt.Sprintf("POINT M (%s)", c.String())
}
func (p *PointM) GeoJSON() string           { return "" }
func (p *PointM) GeometryType() string      { return "POINT M" }
func (p *PointM) String() string            { return p.WKT() }
func (p *PointM) GetCoordinate() Coordinate { return p.coordinates }
