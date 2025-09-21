package gopostgis

import (
	"github.com/sutanmufti/gopostgis/geometry"
)

func GeomArray(n ...GeometryConstructor.Geometry) []GeometryConstructor.Geometry {
	a := []GeometryConstructor.Geometry{}

	for _, g := range n {
		a = append(a, g)
	}

	return a
}

// Creates a 2D XY, 3D XYZ or 4D XYZM Point geometry. This function supports 3d and will not drop the z-index.
func ST_MakePoint(x *float64, y *float64, z *float64, m *float64) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakePoint(x, y, z, m)
	return val, er
}

// Returns a Point with the given X and Y coordinate values. This is the SQL-MM equivalent for ST_MakePoint that takes just X and Y.
func ST_Point(x *float64, y *float64) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakePoint(x, y, nil, nil)
	return val, er
}

// Returns an Point with the given X, Y and M coordinate values.
func ST_PointM(x *float64, y *float64, m *float64) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakePoint(x, y, nil, m)
	return val, er
}

// Returns an Point with the given X, Y and Z coordinate values.
func ST_PointZ(x *float64, y *float64, z *float64) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakePoint(x, y, z, nil)
	return val, er
}

// Returns an Point with the given X, Y, Z and M coordinate values
func ST_PointZM(x *float64, y *float64, z *float64, m *float64) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakePoint(x, y, z, m)
	return val, er
}

// Creates a LineString containing the points of Point, MultiPoint, or LineString geometries. Other geometry types cause an error.
func ST_MakeLine(geometryInput *[]GeometryConstructor.Geometry) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakeLine(geometryInput)
	return val, er
}
