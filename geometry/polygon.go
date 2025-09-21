package GeometryConstructor

import (
	"fmt"
	"strings"
)

type Polygon struct {
	inner    []Coordinate
	outter   []Coordinate
	GeomType string
	dim      string
}

func (p *Polygon) WKT() string {
	innerC := []string{}
	outterC := []string{}

	for _, c := range p.inner {
		innerC = append(innerC, c.String())
	}

	for _, c := range p.outter {
		outterC = append(outterC, c.String())
	}

	innerS := fmt.Sprintf("(%s)", strings.Join(innerC, ", "))
	outterS := fmt.Sprintf("(%s)", strings.Join(outterC, ", "))

	val := fmt.Sprintf("%s (%s, %s)", p.GeomType, outterS, innerS)

	return val
}

func (p *Polygon) GeoJSON() string           { return "" }
func (p *Polygon) GeometryType() string      { return p.GeomType }
func (p *Polygon) String() string            { return p.WKT() }
func (p *Polygon) GetCoordinate() Coordinate { return nil }
func (p *Polygon) GetDim() string            { return "" }
func (p *Polygon) GetCoordinates() []Coordinate {
	return p.outter
}
