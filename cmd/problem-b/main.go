package main

import (
	"fmt"

	flooditsolver "github.com/rajadwarampudi/kattis-test/cmd/problem-b/flood-it-solver"
)

func main() {
	fmt.Println("Inside problem-b")
	var noOfTestCases int
	_, err := fmt.Scan(&noOfTestCases)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	solveBoards(noOfTestCases)

}

func solveBoards(noOfTestCases int) {
	for tcCount := 0; tcCount < noOfTestCases; tcCount++ {
		var size int
		fmt.Scan(&size)
		floodItGame := flooditsolver.FloodIt{}
		for i := 0; i < size; i++ {
			var row []int
			var rowStr string
			fmt.Scan(&rowStr)

			for j := 0; j < len(rowStr); j++ {
				digit := int(rowStr[j] - '0')
				row = append(row, digit)
			}
			floodItGame.AddRow(row)
		}
		floodItGame.PrintSolvedBoardResult()
	}
}
