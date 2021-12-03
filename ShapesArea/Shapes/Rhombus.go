package Shapes

import (
	"example/myTest/Utils"
	"fmt"
)

type Rhombus struct {
	Diagonal1, Diagonal2 float64
}

func (r Rhombus) Area() float64 {
	var result float64 = (r.Diagonal1 * r.Diagonal2) / 2
	Utils.Log("Rhombus", result)
	return result
}

func NewRhombus() *Rhombus {
	var diagonal1, diagonal2 float64
	fmt.Println("Please insert the first and second diagonal of the rhombus")
	fmt.Scan(&diagonal1, &diagonal2)
	return &Rhombus{Diagonal1: diagonal1, Diagonal2: diagonal2}
}
