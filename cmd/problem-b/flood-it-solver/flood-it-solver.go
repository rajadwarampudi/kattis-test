package flooditsolver

import (
	"fmt"
)

type FloodIt struct {
	Board [][]int
}

// Adds the given row to the Board slice
func (floodIt *FloodIt) AddRow(row []int) {
	floodIt.Board = append(floodIt.Board, row)
}

// Prints the 2-dimentional board at given moment
func (floodIt *FloodIt) PrintBoard() {
	for _, row := range floodIt.Board {
		for j, val := range row {
			if j != 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}

// Checks if the board is already flooded by a single color
func (floodIt *FloodIt) isBoardFlooded() bool {
	color := floodIt.Board[0][0]
	for _, row := range floodIt.Board {
		for _, val := range row {
			if val != color {
				return false
			}
		}
	}

	return true
}

// SolveBoard method solves the given flood-it board using greedy algorithm
// and prints two lines of output.
// The first line represents the number of steps required to solve the board
// The second line represents the count of each color chosen while solving the board.
func (floodIt *FloodIt) SolveBoard() {

	totalStepsToSolve := 0
	var chosenColorCounts [6]int

	for !floodIt.isBoardFlooded() {
		currentColor := floodIt.Board[0][0]
		chosenColor := floodIt.chooseBestColor(currentColor)

		if chosenColor == currentColor {
			break
		}

		floodIt.floodPossibleTilesWithChosenColor(0, 0, currentColor, chosenColor)

		totalStepsToSolve++
		chosenColorCounts[chosenColor - 1]++
	}

	fmt.Println(totalStepsToSolve)

	for _, chosenColorCount := range chosenColorCounts {
		fmt.Printf("%d ", chosenColorCount)
	}
	fmt.Println()

}

func (floodIt *FloodIt) floodPossibleTilesWithChosenColor(x, y, currentColor, chosenColor int) {
	if x < 0 || y < 0 || x >= len(floodIt.Board) || y >= len(floodIt.Board) || floodIt.Board[x][y] != currentColor {
		return
	}

	floodIt.Board[x][y] = chosenColor

	floodIt.floodPossibleTilesWithChosenColor(x + 1, y, currentColor, chosenColor)
	floodIt.floodPossibleTilesWithChosenColor(x - 1, y, currentColor, chosenColor)
	floodIt.floodPossibleTilesWithChosenColor(x, y + 1, currentColor, chosenColor)
	floodIt.floodPossibleTilesWithChosenColor(x, y - 1, currentColor, chosenColor)

}

func (floodIt *FloodIt) chooseBestColor(currentColor int) int {
	size := len(floodIt.Board)
	var colorCount [7]int

	var visited [][]bool

	for i := 0; i < size; i++ {
		row := make([]bool, size)
		visited = append(visited, row)
	}

	floodIt.countColors(0, 0, visited, &colorCount)
	maxCount := 0
	bestColor := currentColor

	for color := 1; color <= 6; color++ {
		if color != currentColor {
			if colorCount[color] > maxCount {
				maxCount = colorCount[color]
				bestColor = color
			}
		}
	}

	return bestColor

}

func (floodIt *FloodIt) countColors(x, y int, visited [][]bool, colorCount *[7]int) {
	if x < 0 || y < 0 || x >= len(floodIt.Board) || y >= len(floodIt.Board) || visited[x][y] {
		return
	}

	visited[x][y] = true

	color := floodIt.Board[x][y]

	colorCount[color]++

	floodIt.countColors(x + 1, y, visited, colorCount)
	floodIt.countColors(x - 1, y, visited, colorCount)
	floodIt.countColors(x, y + 1, visited, colorCount)
	floodIt.countColors(x, y - 1, visited, colorCount)
}
