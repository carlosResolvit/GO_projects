package Shapes

import (
	"example/myTest/Utils"
	"fmt"
)

type Rectangle struct {
	Height, Width float64
}

func (r Rectangle) Area() float64 {
	var result float64 = r.Height * r.Width
	Utils.Log("Rectangle", result)
	return result
}

func NewRectangle() *Rectangle {
	var height, width float64
	fmt.Println("Please insert the height and width of the rectangle")
	fmt.Scan(&height, &width)
	return &Rectangle{Height: height, Width: width}
}
