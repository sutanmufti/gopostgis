package GeometryConstructor

type Geometry interface {
	WKT() string
	GeoJSON() string
	GeometryType() string
	String() string
	GetCoordinate() Coordinate
	GetDim() string
	GetCoordinates() []Coordinate
}

type Coordinate interface {
	String() string
	GetCoordinate() []float64
	PromoteM() *PositionM
	PromoteZ() *PositionZ
	PromoteZM() *PositionZM
	IsM() bool
	IsZ() bool
	IsZM() bool
	GetX() *float64
	GetY() *float64
	GetZ() *float64
	GetM() *float64
}

type GeometryCollection []Geometry

type BBox [4]float64
