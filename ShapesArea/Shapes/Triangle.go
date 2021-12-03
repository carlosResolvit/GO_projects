package Shapes

import (
	"example/myTest/Utils"
	"fmt"
)

type Triangle struct {
	Height, Width float64
}

func (t Triangle) Area() float64 {
	var result float64 = (t.Height * t.Width) / 2
	Utils.Log("Triangle", result)
	return result
}

func NewTriangle() *Triangle {
	var height, width float64
	fmt.Println("Please insert the height and width of the triangle")
	fmt.Scan(&height, &width)
	return &Triangle{Height: height, Width: width}
}
