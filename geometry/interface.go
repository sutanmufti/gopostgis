package GeometryConstructor

type Geometry interface {
	WKT() string
	GeoJSON() string
	GeometryType() string
	String() string
	GetCoordinate() Coordinate
}

type Coordinate interface {
	String() string
	GetCoordinate() []float64
	PromoteM() *PositionM
	PromoteZ() *PositionZ
	PromoteZM() *PositionZM
}

type GeometryCollection []Geometry

type BBox [4]float64
