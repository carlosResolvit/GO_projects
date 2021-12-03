package Shapes

import (
	"example/myTest/Utils"
	"fmt"
)

type Parallelogram struct {
	Base, Height float64
}

func (p Parallelogram) Area() float64 {
	var result float64 = p.Base * p.Height
	Utils.Log("Parallelogram", result)
	return result
}

func NewParallelogram() *Parallelogram {
	var base, height float64
	fmt.Println("Please insert the base and height of the parallelogram")
	fmt.Scan(&base, &height)
	return &Parallelogram{Base: base, Height: height}
}
