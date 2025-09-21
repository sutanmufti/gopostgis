package GeometryConstructor

type LineString struct {
	coordinates []Coordinate
	GeomType    string
}

func (p *LineString) WKT() string               { return "" }
func (p *LineString) GeoJSON() string           { return "" }
func (p *LineString) GeometryType() string      { return p.GeomType }
func (p *LineString) String() string            { return p.WKT() }
func (p *LineString) GetCoordinate() Coordinate { return nil }

func MakeLineStringCoordinates(geometryInput *[]Geometry) []Coordinate {
	coords := []Coordinate{}
	for _, g := range *geometryInput {
		coord := g.GetCoordinate()
		coords = append(coords, coord)
	}
	return coords
}
