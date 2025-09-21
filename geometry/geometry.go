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

func ST_MakeLine(geometryInput *[]Geometry) error {

	contains := func(slice []string, val string) bool {
		return slices.Contains(slice, val)
	}
	var geomType string
	var allowed bool
	var err error
	allowedTypes := []string{"POINT", "POINT Z", "POINT M", "POINT ZM"}

	zExists := false
	mExists := false
	for _, geom := range *geometryInput {
		geomType = geom.GeometryType()

		allowed = contains(allowedTypes, geomType)

		if !allowed {
			err = errors.New("Type not allowed ")
			break
		}

		if !zExists || mExists {
			switch geomType {
			case "POINT ZM":
				zExists = true
				mExists = true
			case "POINT Z":
				zExists = true
			case "POINT M":
				mExists = true
			}
		}
	}

	var thisGeomType string

	if mExists && zExists {
		thisGeomType = "LINESTRING ZM"
	} else if mExists && !zExists {
		thisGeomType = "LINESTRING M"
	} else if !mExists && zExists {
		thisGeomType = "LINESTRING Z"
	} else if !mExists && !zExists {
		thisGeomType = "LINESTRING"
	}

	switch thisGeomType {
	case "LINESTRING ZM":
	}

	return err
}

func ST_MakePolygon(c [][]Position) Polygon {
	return Polygon{c}
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
