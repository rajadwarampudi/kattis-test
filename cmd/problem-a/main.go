package main

import (
	"fmt"

	legcounter "github.com/rajadwarampudi/kattis-test/cmd/problem-a/leg-counter"
)

func main() {
	var b, d, c, totalLegs int
	fmt.Println("Enter legs of three animals and totallegs")
	_, err := fmt.Scan(&b, &d, &c, &totalLegs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	legcounter.PrintLegCounts(b, d, c, totalLegs)
}