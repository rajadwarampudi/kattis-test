package main

import (
	"fmt"

	pizzaslicer "github.com/rajadwarampudi/kattis-test/cmd/problem-c/pizza-slicer"
)

func main() {
	fmt.Println("Inside problem-c")
	var noOfTestCases int
	_, err := fmt.Scan(&noOfTestCases)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	slicePizzas(noOfTestCases)
}

func slicePizzas(noOfTestCases int) {
	for tcCount := 0; tcCount < noOfTestCases; tcCount++ {
		var radius, noOfPeople, angleDegrees, angleMinutes, angleSeconds int
		fmt.Scan(&radius, &noOfPeople, &angleDegrees, &angleMinutes, &angleSeconds)

		pizzaslicer.PrintMaxPossibleSliceArea(radius, noOfPeople, angleDegrees, angleMinutes, angleSeconds)
	}
}