package GeometryConstructor

import (
	"testing"
)

func F64(n float64) *float64 {
	return &n
}

func TestST_MakePoint(t *testing.T) {

	p, _ := ST_MakePoint(F64(4), F64(5), nil, nil)
	pz, _ := ST_MakePoint(F64(4), F64(5), F64(6), nil)
	pzm, _ := ST_MakePoint(F64(4), F64(5), F64(6), F64(7))
	pm, _ := ST_MakePoint(F64(4), F64(5), nil, F64(7))

	wktP := p.WKT()
	wktPZ := pz.WKT()
	wktPZM := pzm.WKT()
	wktPM := pm.WKT()

	same := "POINT (4.000000 5.000000)" == wktP
	samePZ := "POINT Z (4.000000 5.000000 6.000000)" == wktPZ
	samePZM := "POINT ZM (4.000000 5.000000 6.000000 7.000000)" == wktPZM
	samePM := "POINT M (4.000000 5.000000 7.000000)" == wktPM

	if !same {
		t.Error("same error")
	}
	if !samePZ {
		t.Error("samePZ error")
	}
	if !samePZM {
		t.Error("samePZM error")
	}
	if !samePM {
		t.Error("samePM error")
	}

}

func TestST_MakeLine(t *testing.T) {
	p1, _ := ST_MakePoint(F64(4), F64(5), nil, F64(5))
	p2, _ := ST_MakePoint(F64(5), F64(6), nil, F64(5))

	points := []Geometry{
		p1,
		p2,
	}

	linestring, _ := ST_MakeLine(&points)

	if "LINESTRING M (4.000000 5.000000 5.000000, 5.000000 6.000000 5.000000)" != linestring.WKT() {
		t.Error("Linestring does not match")
	}
}
