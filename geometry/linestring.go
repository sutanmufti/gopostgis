package GeometryConstructor

type LineString struct {
	coordinates []Position
}

func (p *LineString) WKT() string          { return "" }
func (p *LineString) GeoJSON() string      { return "" }
func (p *LineString) GeometryType() string { return "POINT Z" }
func (p *LineString) String() string       { return "POINT Z" }
