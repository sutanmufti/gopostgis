package GeometryConstructor

type Point struct {
	coordinates Coordinate
}

func (p *Point) WKT() string          { return "" }
func (p *Point) GeoJSON() string      { return "" }
func (p *Point) GeometryType() string { return "POINT" }
func (p *Point) String() string       { return "POINT" }
