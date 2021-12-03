package AreaOptions

import (
	"fmt"
)

func ShowOptions() {
	fmt.Println("Please choose an option")
	fmt.Println("1) Area for Circle ")
	fmt.Println("2) Area for Square")
	fmt.Println("3) Area for Triangle")
	fmt.Println("4) Area for Rhombus")
	fmt.Println("5) Area for Rectangle")
	fmt.Println("6) Area for Parallelogram")
	fmt.Println("/quit to exit")
}

func WelcomeMessage() {
	fmt.Println("Welcome!")
	fmt.Println("Please select a shape to calculate its area.")
	fmt.Println("To quit at any time, just type /quit.")
}
