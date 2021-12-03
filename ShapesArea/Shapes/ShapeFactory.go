package Shapes

import (
	"errors"
	"fmt"
)

func GenerateAreaForShape(shape string) (IArea, error) {
	fmt.Printf("Calculate area for a %v \n", shape)
	switch shape {
	case "1", "circle":
		return NewCircle(), nil
	case "2", "square":
		return NewSquare(), nil
	case "3", "triangle":
		return NewTriangle(), nil
	case "4", "rhombus":
		return NewRhombus(), nil
	case "5", "rectangle":
		return NewRectangle(), nil
	case "6", "parallelogram":
		return NewParallelogram(), nil
	default:
		fmt.Println("We dont recognize that command :(")
		return nil, errors.New("help")
	}
}
