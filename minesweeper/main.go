package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const side int = 9   // side length of the board
const mines int = 10 // number of mines on the board

type slot struct {
	value       string
	isDisplayed bool
	flag        bool
}

// Returns True if the cell is valid
func isValid(i int, j int) bool {
	return (i >= 0) && (i < side) && (j >= 0) && (j < side)
}

// Returns True if the cell has a Mine
func isMine(i int, j int, aTable [side][side]slot) bool {
	return aTable[i][j].value == "x"
}

func countAdjacentMines(row int, col int, aBoard [side][side]slot) int {

	count := 0

	//----------- 1st to 8th Neighbours ------------
	// Only process this cells if are a valids ones
	if isValid(row-1, col) && isMine(row-1, col, aBoard) {
		count++
	}

	if isValid(row+1, col) && isMine(row+1, col, aBoard) {
		count++
	}

	if isValid(row, col+1) && isMine(row, col+1, aBoard) {
		count++
	}

	if isValid(row, col-1) && isMine(row, col-1, aBoard) {
		count++
	}

	if isValid(row-1, col+1) && isMine(row-1, col+1, aBoard) {
		count++
	}

	if isValid(row-1, col-1) && isMine(row-1, col-1, aBoard) {
		count++
	}

	if isValid(row+1, col+1) && isMine(row+1, col+1, aBoard) {
		count++
	}

	if isValid(row+1, col-1) && isMine(row+1, col-1, aBoard) {
		count++
	}

	return count
}

func createRandomeTable() [side][side]slot {
	// Populate table with "-"
	var auxTable [side][side]slot
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			auxTable[i][j].isDisplayed = false
			auxTable[i][j].flag = false
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

	// Add numeric values to Table
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			if !isMine(i, j, auxTable) {
				count := countAdjacentMines(i, j, auxTable)
				auxTable[i][j].value = strconv.Itoa(count)
			}
		}
	}

	return auxTable
}

func flipSlot(i int, j int, table [side][side]slot) {
	//if is una bomba?
	//GAME OVER
	//else is a number > 0
	// else is 0
	//display blank neighbours

}

/*func printTable(aTable [side][side]slot) {
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			fmt.Print(aTable[i][j].value + " ")
		}
		fmt.Println(" ")
	}
}*/

func printGameTable(aTable [side][side]slot) {
	fmt.Println("    1   2   3   4   5   6   7   8   9")
	for i := 0; i < side; i++ {
		fmt.Println("  +---+---+---+---+---+---+---+---+---+")
		fmt.Print(i + 1)
		fmt.Print(" ")
		for j := 0; j < side; j++ {
			if aTable[i][j].isDisplayed {
				fmt.Print("| " + aTable[i][j].value + " ")
			} else {
				fmt.Print("| - ")
			}
		}
		fmt.Print("|")
		fmt.Println(" ")
	}
	fmt.Println("  +---+---+---+---+---+---+---+---+---+")
}

func main() {
	gameOver := false
	table := createRandomeTable()
	var row int = -1
	var col int = -1
	for !gameOver {
		fmt.Println("Current Table:")
		printGameTable(table)
		fmt.Println("Choose a row and a column:")
		fmt.Scan(&row, &col)
		if isValid(row, col) {
			flipSlot(row, col, table)
		} else {
			fmt.Println("ERROR")
		}
	}
	//imprimir instrucciones - Ingrese fila y columna:
	//escanear lo que ingresa el usuario

}
