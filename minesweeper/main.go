package main

import (
	"fmt"
	"math/rand"
	"time"
)

const side int = 9   // side length of the board
const mines int = 10 // number of mines on the board

type slot struct {
	value      string
	isSelected bool
}

func createRandomeTable() [side][side]slot {
	// Populate table with "-"
	var auxTable [side][side]slot
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			auxTable[i][j].isSelected = false
			auxTable[i][j].value = "-"
		}
	}
	// Charge random mines "x" in table
	for i := 0; i < mines; i++ {
		x1 := rand.NewSource(time.Now().UnixNano())
		y1 := rand.New(x1)
		rand1 := y1.Intn(side)
		rand2 := y1.Intn(side)
		if auxTable[rand1][rand2].value != "-" {
			i--
		} else {
			auxTable[rand1][rand2].value = "x"
		}
	}

	// Add numeric values

	return auxTable
}

func printTable(aTable [side][side]slot) {
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			fmt.Print(aTable[i][j].value + " ")
		}
		fmt.Println(" ")
	}
}

func main() {
	table := createRandomeTable()
	printTable(table)

}
