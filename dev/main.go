package main

import (
	"fmt"

	gogeos "github.com/sutanmufti/gopostgis"
)

func main() {

	fmt.Println("hello world")

	x := 2.0

	v, _ := gogeos.ST_PointZ(&x, &x, &x)

	fmt.Println(v)
}
