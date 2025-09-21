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

func contains(slice []string, val string) bool {
	return slices.Contains(slice, val)
}

func ST_MakeLine(geometryInput *[]Geometry) (Geometry, error) {
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

	dim := ""

	switch thisGeomType {
	case "POINT":
		thisGeomType = "LINESTRING"
		dim = ""
	case "POINT Z":
		thisGeomType = "LINESTRING Z"
		dim = "Z"
	case "POINT M":
		thisGeomType = "LINESTRING M"
		dim = "M"
	case "POINT ZM":
		thisGeomType = "LINESTRING ZM"
		dim = "ZM"
	}

	return &LineString{coordinates, thisGeomType, dim}, err
}

func ST_MakePolygon(outter Geometry, inner Geometry) (Geometry, error) {
	if outter == nil {
		return nil, errors.New("Must provide outter LineString.")
	}
	containsInner := inner != nil

	outterGeomType := outter.GeometryType()
	allowedTypes := []string{"LINESTRING", "LINESTRING Z", "LINESTRING M", "LINESTRING ZM"}
	allowedOutter := contains(allowedTypes, outterGeomType)

	if containsInner {
		innerGeomType := inner.GeometryType()
		allowedInner := contains(allowedTypes, innerGeomType)
		if !allowedOutter || !allowedInner {
			return nil, errors.New("Can only make a polygon from LINESTRING")
		}

	}

	if !allowedOutter {
		return nil, errors.New("Can only make a polygon from LINESTRING")
	}

	innerCoords := []Coordinate{}
	outterCoords := []Coordinate{}
	var err error

	if containsInner {
		if outter.GetDim() != inner.GetDim() {
			err = errors.New("Outter and Inner linestrings have different dimensions")
			return nil, err
		}
	}

	for _, c := range outter.GetCoordinates() {
		outterCoords = append(outterCoords, c)
	}

	if containsInner {
		for _, c := range inner.GetCoordinates() {
			innerCoords = append(innerCoords, c)
		}
	}

	if len(outterCoords) < 3 {
		return nil, errors.New("Not enough outter coordinates")
	}

	if len(innerCoords) < 3 {
		return nil, errors.New("Not enough inner coordinates")
	}

	if c1, c2 := outterCoords[0], outterCoords[1]; (c1.GetX() != c2.GetX()) || (c1.GetY() != c2.GetY()) || (c1.GetZ() != c2.GetZ()) || (c1.GetM() != c2.GetM()) {
		return nil, errors.New("Invalid outter coordinates: starting coordinate != closing coordinate")
	}

	if c1, c2 := innerCoords[0], innerCoords[1]; (c1.GetX() != c2.GetX()) || (c1.GetY() != c2.GetY()) || (c1.GetZ() != c2.GetZ()) || (c1.GetM() != c2.GetM()) {
		return nil, errors.New("Invalid inner coordinates: starting coordinate != closing coordinate")
	}

	dim := outter.GetDim()
	var GeomtypeAssign string
	switch dim {
	case "":
		GeomtypeAssign = "POLYGON"
	case "Z":
		GeomtypeAssign = "POLYGON Z"
	case "ZM":
		GeomtypeAssign = "POLYGON ZM"
	case "M":
		GeomtypeAssign = "POLYGON M"
	}

	p := Polygon{
		inner:    innerCoords,
		outter:   outterCoords,
		GeomType: GeomtypeAssign,
		dim:      outter.GetDim(),
	}

	return &p, err
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
