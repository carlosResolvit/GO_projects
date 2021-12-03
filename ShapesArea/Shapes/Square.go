package Shapes

import (
	"example/myTest/Utils"
	"fmt"
	"math"
)

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	var result float64 = math.Pow(s.Side, 2)
	Utils.Log("Square", result)
	return result
}

func NewSquare() *Square {
	var side float64
	fmt.Println("Please insert the side of the square")
	fmt.Scan(&side)
	return &Square{Side: side}
}
