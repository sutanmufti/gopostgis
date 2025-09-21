package GeometryConstructor

type PointM struct {
	coordinates Coordinate
}

func (p *PointM) WKT() string          { return "" }
func (p *PointM) GeoJSON() string      { return "" }
func (p *PointM) GeometryType() string { return "POINT M" }
func (p *PointM) String() string       { return "POINT M" }
