package GeometryConstructor

type PointZM struct {
	coordinates PositionZM
}

func (p *PointZM) WKT() string          { return "" }
func (p *PointZM) GeoJSON() string      { return "" }
func (p *PointZM) GeometryType() string { return "POINT ZM" }
func (p *PointZM) String() string       { return "POINT ZM" }
