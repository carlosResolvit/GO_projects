package main

import (
	"example/myTest/AreaOptions"
	"example/myTest/Shapes"
	"example/myTest/Utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	AreaOptions.WelcomeMessage()
	for {
		AreaOptions.ShowOptions()
		userInput := Utils.UserInput()
		if userInput == "/quit" {
			os.Exit(1)
		}
		shapes, err := Shapes.GenerateAreaForShape(strings.ToLower(userInput))
		if err == nil {
			fmt.Println("\033[2J")
			fmt.Printf("Area for shape is %.4f\n", shapes.Area())
		} else {
			fmt.Println("Invalid input")
		}

	}
}
