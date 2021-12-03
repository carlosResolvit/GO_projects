package main

import (
	"fmt"
	"math/rand"
	"time"
)

type slot struct {
	value      string
	isSelected bool
}

var side int = 9   // side length of the board
var mines int = 10 // number of mines on the board

func createRandomeTable() [9][9]slot {
	var auxTable [9][9]slot
	for i := 0; i < 9; i++ {
		for j := 0; i < 9; i++ {
			auxTable[i][j].isSelected = false
			auxTable[i][j].value = "-"
		}
	}
	for i := 0; i < 10; i++ {
		x1 := rand.NewSource(time.Now().UnixNano())
		y1 := rand.New(x1)
		rand1 := y1.Intn(9)
		rand2 := y1.Intn(9)
		if auxTable[rand1][rand2].value == "-" {
			auxTable[rand1][rand2].value = "x"
			i++
		}
	}
	return auxTable
}

func printTable(aTable [9][9]slot) {
	for i := 0; i < 9; i++ {
		for j := 0; i < 9; i++ {
			fmt.Print(aTable[i][j].value + " ")
		}
		fmt.Println(" ")
	}
}

func main() {
	table := createRandomeTable()
	printTable(table)

}
