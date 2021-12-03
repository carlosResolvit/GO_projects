package Shapes

import (
	"example/myTest/Utils"
	"fmt"
	"math"
)

type Circle struct {
	Radious float64
}

func (c Circle) Area() float64 {
	var result float64 = math.Pi * math.Pow(c.Radious, 2)
	Utils.Log("Circle", result)
	return result
}

func NewCircle() *Circle {
	var radious float64
	fmt.Println("Please insert the radious of the circle")
	fmt.Scan(&radious)
	return &Circle{Radious: radious}
}
