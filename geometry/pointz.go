package GeometryConstructor

type PointZ struct {
	coordinates PositionZ
}

func (p *PointZ) WKT() string          { return "" }
func (p *PointZ) GeoJSON() string      { return "" }
func (p *PointZ) GeometryType() string { return "POINT Z" }
func (p *PointZ) String() string       { return "POINT Z" }
