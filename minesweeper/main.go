package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const side int = 9   // side length of the board
const mines int = 10 // number of mines on the board

type game struct {
	aTable   [side][side]slot
	gameOver bool
	aWinner  bool
}

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

func createRandomeGame() game {
	var auxGame game
	auxGame.aWinner = false
	auxGame.gameOver = false

	// Populate table with "-"
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			auxGame.aTable[i][j].isDisplayed = false
			auxGame.aTable[i][j].flag = false
			auxGame.aTable[i][j].value = "-"
		}
	}
	// Charge random mines "x" in table
	for i := 0; i < mines; i++ {
		x1 := rand.NewSource(time.Now().UnixNano())
		y1 := rand.New(x1)
		rand1 := y1.Intn(side)
		rand2 := y1.Intn(side)
		if auxGame.aTable[rand1][rand2].value != "-" {
			i--
		} else {
			auxGame.aTable[rand1][rand2].value = "x"
		}
	}

	// Add numeric values to Table
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			if !isMine(i, j, auxGame.aTable) {
				count := countAdjacentMines(i, j, auxGame.aTable)
				auxGame.aTable[i][j].value = strconv.Itoa(count)
			}
		}
	}

	return auxGame
}

// A recursive function to display close values to empty spaces
func (aGame *game) displayCloseValues(row int, col int) {
	aGame.aTable[row][col].isDisplayed = true
	//LEFT
	/*if isValid(row, col-1) && !isMine(row, col-1, aGame.aTable) {
		if aGame.aTable[row][col-1].value == "0" {
			aGame.displayCloseValues(row, col-1)
		} else {
			aGame.aTable[row][col-1].isDisplayed = true
		}
	}*/
	//RIGHT
	if isValid(row, col+1) && !isMine(row, col+1, aGame.aTable) {
		if aGame.aTable[row][col+1].value == "0" {
			aGame.displayCloseValues(row, col+1)
		} else {
			aGame.aTable[row][col+1].isDisplayed = true
		}
	}
	//UP
	/*if isValid(row-1, col) && !isMine(row-1, col, aGame.aTable) {
		if aGame.aTable[row-1][col].value == "0" {
			aGame.displayCloseValues(row-1, col)
		} else {
			aGame.aTable[row-1][col].isDisplayed = true
		}
	}
	//DOWN
	if isValid(row+1, col) && !isMine(row+1, col, aGame.aTable) {
		if aGame.aTable[row+1][col].value == "0" {
			aGame.displayCloseValues(row+1, col)
		} else {
			aGame.aTable[row+1][col].isDisplayed = true
		}
	}*/
}

func (aGame *game) flipSlot(row int, col int) {
	if aGame.aTable[row][col].value == "x" {
		aGame.aTable[row][col].value = "X"
		aGame.aTable[row][col].isDisplayed = true
		aGame.gameOver = true
		for i := 0; i < side; i++ {
			for j := 0; j < side; j++ {
				if isMine(i, j, aGame.aTable) {
					aGame.aTable[i][j].isDisplayed = true
				}
			}
		}
	} else if aGame.aTable[row][col].value == "0" {
		aGame.displayCloseValues(row, col)
	} else {
		aGame.aTable[row][col].isDisplayed = true
	}
}

func printTable(aTable [side][side]slot) {
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			fmt.Print(aTable[i][j].value + " ")
		}
		fmt.Println(" ")
	}
}

func (aGame *game) printGameTable() {
	fmt.Println("    1   2   3   4   5   6   7   8   9")
	for i := 0; i < side; i++ {
		fmt.Println("  +---+---+---+---+---+---+---+---+---+")
		fmt.Print(i + 1)
		fmt.Print(" ")
		for j := 0; j < side; j++ {
			if aGame.aTable[i][j].isDisplayed {
				fmt.Print("| " + aGame.aTable[i][j].value + " ")
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
	auxGame := createRandomeGame()
	printTable(auxGame.aTable) //TO DEBUG
	var row int
	var col int
	for !auxGame.gameOver && !auxGame.aWinner {
		//fmt.Println("\033[2J")
		fmt.Println("Current Table:")
		auxGame.printGameTable()
		fmt.Println("Choose a row and a column:")
		fmt.Scan(&row, &col)
		if isValid(row-1, col-1) {
			auxGame.flipSlot(row-1, col-1)
		} else {
			fmt.Println("The value is not valid")
		}
	}
	fmt.Println("GAME OVER")
	fmt.Println("Current Table:")
	auxGame.printGameTable()
}
