package gopostgis

import (
	"github.com/sutanmufti/gopostgis/geometry"
)

func ST_MakePoint(x *float64, y *float64, z *float64, m *float64) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakePoint(x, y, z, m)
	return val, er
}

func ST_Point(x *float64, y *float64) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakePoint(x, y, nil, nil)
	return val, er
}

func ST_PointM(x *float64, y *float64, m *float64) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakePoint(x, y, nil, m)
	return val, er
}
func ST_PointZ(x *float64, y *float64, z *float64) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakePoint(x, y, z, nil)
	return val, er
}

func ST_PointZM(x *float64, y *float64, z *float64, m *float64) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakePoint(x, y, z, m)
	return val, er
}

func ST_MakeLine(geometryInput *[]GeometryConstructor.Geometry) (GeometryConstructor.Geometry, error) {
	val, er := GeometryConstructor.ST_MakeLine(geometryInput)
	return val, er
}
