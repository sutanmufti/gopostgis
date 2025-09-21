package GeometryConstructor

type Polygon struct {
	inner    []Coordinate
	outter   []Coordinate
	GeomType string
	dim      string
}

func (p *Polygon) WKT() string {
	return ""
}
func (p *Polygon) GeoJSON() string           { return "" }
func (p *Polygon) GeometryType() string      { return p.GeomType }
func (p *Polygon) String() string            { return p.WKT() }
func (p *Polygon) GetCoordinate() Coordinate { return nil }
func (p *Polygon) GetDim() string            { return "" }
