package GeometryConstructor

type Geometry interface {
	WKT() string
	GeoJSON() string
	GeometryType() string
	String() string
	GetCoordinate() Coordinate
}

type GeometryCollection []Geometry

type BBox [4]float64
