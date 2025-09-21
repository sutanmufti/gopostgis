package GeometryConstructor

import (
	"fmt"
	"strings"
)

type LineString struct {
	coordinates []Coordinate
	GeomType    string
	Dim         string
}

func (p *LineString) WKT() string {

	components := []string{}
	for _, coor := range p.coordinates {
		cstr := coor.String()
		components = append(components, cstr)

	}

	retval := strings.Join(components, ", ")
	geomType := p.GeomType

	return fmt.Sprintf("%s (%s)", geomType, retval)
}
func (p *LineString) GeoJSON() string           { return "" }
func (p *LineString) GeometryType() string      { return p.GeomType }
func (p *LineString) String() string            { return p.WKT() }
func (p *LineString) GetCoordinate() Coordinate { return nil }
func (p *LineString) GetDim() string            { return "" }

func MakeLineStringCoordinates(geometryInput *[]Geometry, thisGeomType string) []Coordinate {

	coords := []Coordinate{}
	for _, g := range *geometryInput {
		coords = append(coords, g.GetCoordinate())
	}
	return coords
}
