package main

import (
	"fmt"
	"math/rand"
	"time"
)

var side int = 9   // side length of the board
var mines int = 10 // number of mines on the board

func createRandomeTable() [9][9]int {
	var auxTable [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; i < 9; i++ {
			auxTable[i][j] = 0
		}
	}
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	fmt.Println(y1.Intn(9))
	return auxTable
}

func printTable(aTable [9][9]int) {

}

func main() {
	table := createRandomeTable()
	printTable(table)

}
