package GeometryConstructor

import (
	"fmt"
	"strings"
)

type LineString struct {
	coordinates []Coordinate
	GeomType    string
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

func MakeLineStringCoordinates(geometryInput *[]Geometry, thisGeomType string) []Coordinate {

	coords := []Coordinate{}
	for _, g := range *geometryInput {
		coord := g.GetCoordinate()
		switch thisGeomType {
		case "LINESTRING ZM":
			coords = append(coords, coord.PromoteZM())
		case "LINESTRING M":
			coords = append(coords, coord.PromoteM())
		case "LINESTRING Z":
			coords = append(coords, coord.PromoteM())
		case "LINESTRING":
			coords = append(coords, coord)
		}
	}
	return coords
}
