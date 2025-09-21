package GeometryConstructor

import (
	"errors"
	"slices"
)

func ST_MakePoint(x *float64, y *float64, z *float64, m *float64) (Geometry, error) {
	var ctype string
	var er error = nil

	if x == nil || y == nil {
		er = errors.New("Must provide X and Y")
	} else if m != nil && z != nil {
		ctype = "PointZM"
	} else if m != nil && z == nil {
		ctype = "PointM"
	} else if m == nil && z != nil {
		ctype = "PointZ"
	} else {
		ctype = "Point"
	}

	switch ctype {
	case "Point":
		return &Point{&Position{x, y}}, er
	case "PointM":
		return &PointM{&PositionM{x, y, m}}, er
	case "PointZ":
		return &PointZ{&PositionZ{x, y, z}}, er
	case "PointZM":
		return &PointZM{&PositionZM{x, y, z, m}}, er
	}

	return &Point{&Position{x, y}}, er
}

func ST_MakeLine(geometryInput *[]Geometry) (Geometry, error) {

	contains := func(slice []string, val string) bool {
		return slices.Contains(slice, val)
	}

	var allowed bool
	var err error
	allowedTypes := []string{"POINT", "POINT Z", "POINT M", "POINT ZM"}

	var thisGeomType string
	for i, geom := range *geometryInput {
		geomType := geom.GeometryType()
		allowed = contains(allowedTypes, geomType)
		if !allowed {
			err = errors.New("Type not allowed")
			return nil, err
		}

		if i == 0 {
			thisGeomType = geomType
		} else if thisGeomType != geomType {
			err = errors.New("Inconsistent positions")
			return nil, err
		}

	}

	coordinates := MakeLineStringCoordinates(geometryInput, thisGeomType)

	switch thisGeomType {
	case "POINT":
		thisGeomType = "LINESTRING"
	case "POINT Z":
		thisGeomType = "LINESTRING Z"
	case "POINT M":
		thisGeomType = "LINESTRING M"
	case "POINT ZM":
		thisGeomType = "LINESTRING ZM"
	}

	return &LineString{coordinates, thisGeomType}, err
}

func ST_MakePolygon(outter []LineString, inner []LineString) Geometry {
	innerCoords := []Coordinate{}
	outterCoords := []Coordinate{}

	for _, c := range outter {
		outterCoords = append(outterCoords, c.coordinates...)
	}

	for _, c := range inner {
		innerCoords = append(innerCoords, c.coordinates...)
	}

	p := Polygon{
		inner:    innerCoords,
		outter:   outterCoords,
		GeomType: "Polygon",
	}

	return &p
}

func ST_MakeEnvelope(xmin float64, ymin float64, xmax float64) {

}

func ST_Collect(geoms ...Geometry) GeometryCollection {
	geomCollection := []Geometry{}

	for _, g := range geoms {
		geomCollection = append(geomCollection, g)
	}
	return geomCollection
}
