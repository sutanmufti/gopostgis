package GeometryConstructor

import (
	"fmt"
	"strings"
)

type Position [2]*float64
type PositionZ [3]*float64
type PositionZM [4]*float64
type PositionM [3]*float64

func (p *Position) GetX() *float64 {
	return p[0]
}

func (p *Position) GetY() *float64 {
	return p[1]
}

func (p *Position) GetM() *float64 {
	return nil
}

func (p *Position) GetZ() *float64 {
	return nil
}

func (p *PositionM) GetX() *float64 {
	return p[0]
}

func (p *PositionM) GetY() *float64 {
	return p[1]
}

func (p *PositionM) GetM() *float64 {
	return p[2]
}

func (p *PositionM) GetZ() *float64 {
	return nil
}

func (p *PositionZ) GetX() *float64 {
	return p[0]
}

func (p *PositionZ) GetY() *float64 {
	return p[1]
}

func (p *PositionZ) GetM() *float64 {
	return nil
}

func (p *PositionZ) GetZ() *float64 {
	return p[2]
}

func (p *PositionZM) GetX() *float64 {
	return p[0]
}

func (p *PositionZM) GetY() *float64 {
	return p[1]
}

func (p *PositionZM) GetM() *float64 {
	return p[3]
}

func (p *PositionZM) GetZ() *float64 {
	return p[2]
}

func (p *Position) IsM() bool {
	return false
}

func (p *Position) IsZM() bool {
	return false
}

func (p *Position) IsZ() bool {
	return false
}

func (p *PositionM) IsM() bool {
	return true
}

func (p *PositionM) IsZM() bool {
	return false
}

func (p *PositionM) IsZ() bool {
	return false
}

func (p *PositionZ) IsM() bool {
	return false
}

func (p *PositionZ) IsZM() bool {
	return false
}

func (p *PositionZ) IsZ() bool {
	return true
}

func (p *PositionZM) IsM() bool {
	return false
}

func (p *PositionZM) IsZM() bool {
	return true
}

func (p *PositionZM) IsZ() bool {
	return false
}

func (p *Position) String() string {
	components := []string{}
	for _, num := range *p {

		if num == nil {
			c := "NaN"
			components = append(components, c)
			continue
		} else {
			c := fmt.Sprintf("%f", (*num))
			components = append(components, c)
		}
	}
	retval := strings.Join(components, " ")
	return retval
}

func (p *PositionM) String() string {
	components := []string{}
	for _, num := range *p {

		if num == nil {
			c := "NaN"
			components = append(components, c)
			continue
		} else {
			c := fmt.Sprintf("%f", (*num))
			components = append(components, c)
		}
	}
	retval := strings.Join(components, " ")
	return retval
}

func (p *PositionZ) String() string {
	components := []string{}
	for _, num := range *p {

		if num == nil {
			c := "NaN"
			components = append(components, c)
			continue
		} else {
			c := fmt.Sprintf("%f", (*num))
			components = append(components, c)
		}
	}
	retval := strings.Join(components, " ")
	return retval
}

func (p *PositionZM) String() string {
	components := []string{}
	for _, num := range *p {

		if num == nil {
			c := "NaN"
			components = append(components, c)
			continue
		} else {
			c := fmt.Sprintf("%f", (*num))
			components = append(components, c)
		}
	}
	retval := strings.Join(components, " ")
	return retval
}

func (p *Position) GetCoordinate() []float64 {

	return []float64{}
}

func (p *PositionZ) GetCoordinate() []float64 {

	return []float64{}
}

func (p *PositionZM) GetCoordinate() []float64 {

	return []float64{}
}

func (p *PositionM) GetCoordinate() []float64 {

	return []float64{}
}

func (p *Position) PromoteZ() *PositionZ {
	return &PositionZ{p[0], p[1], nil}
}

func (p *Position) PromoteZM() *PositionZM {
	return &PositionZM{p[0], p[1], nil, nil}
}

func (p *Position) PromoteM() *PositionM {
	return &PositionM{p[0], p[1], nil}
}

func (p *PositionZ) PromoteZM() *PositionZM {
	return &PositionZM{p[0], p[1], p[2], nil}
}

func (p *PositionZ) PromoteM() *PositionM {
	return &PositionM{p[0], p[1], nil}
}

func (p *PositionZM) PromoteM() *PositionM {
	return &PositionM{p[0], p[1], p[3]}
}

func (p *PositionZM) PromoteZ() *PositionZ {
	return &PositionZ{p[0], p[1], p[2]}
}

func (p *PositionZM) PromoteZM() *PositionZM {
	return p
}

func (p *PositionZ) PromoteZ() *PositionZ {
	return p
}

func (p *PositionM) PromoteZM() *PositionZM {
	return &PositionZM{p[0], p[1], nil, p[2]}
}

func (p *PositionM) PromoteM() *PositionM {
	return &PositionM{p[0], p[1], p[2]}
}

func (p *PositionM) PromoteZ() *PositionZ {
	return &PositionZ{p[0], p[1], nil}
}

func (p *PositionZ) DropZ() *Position {
	return &Position{p[0], p[1]}
}

func (p *PositionM) DropM() *Position {
	return &Position{p[0], p[1]}
}

func (p *PositionZM) DropZ() *PositionM {
	return &PositionM{p[0], p[1], p[3]}
}

func (p *PositionZM) DropM() *PositionZ {
	return &PositionZ{p[0], p[1], p[2]}
}

func (p *PositionZM) DropZM() *Position {
	return &Position{p[0], p[1]}
}
