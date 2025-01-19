package pizzaslicer

import (
	"fmt"
	"math"
)

const FULL_CIRCLE = 2 * math.Pi

// Calculates and prints the maximum area of a slice in the pizza after slicing it repeatedly for given number of people
// at the given angle of degrees, minutes and seconds
// and prints the area of the slice with maximum area
func PrintMaxPossibleSliceArea(radius, noOfPeople, angleDegrees, angleMinutes, angleSeconds int) {
	
	largestSliceArea := CalculateMaxPossibleSliceArea(radius, noOfPeople, angleDegrees, angleMinutes, angleSeconds)

	fmt.Printf("%f", largestSliceArea)
}

// Calculates the maximum area of a slice in the pizza after slicing it repeatedly for given number of people
// at the given angle of degrees, minutes and seconds
// and prints the area of the slice with maximum area
func CalculateMaxPossibleSliceArea(radius, noOfPeople, angleDegrees, angleMinutes, angleSeconds int) float64 {
	fmt.Println("inside pizza slice solution method")
	angle := (math.Pi / 180.0) * (float64(angleDegrees) + float64(angleMinutes)/60.0 + float64(angleSeconds)/3600)

	var maxAngle, currentAngle, prevAngle float64

	for i := 0; i < noOfPeople; i++ {
		currentAngle = math.Mod(currentAngle + angle, FULL_CIRCLE)
		maxAngle = math.Max(maxAngle, currentAngle - prevAngle)
		prevAngle = currentAngle
	}

	maxAngle = math.Max(maxAngle, FULL_CIRCLE - currentAngle)

	largestSliceArea := 0.5 * maxAngle * float64(radius) * float64(radius)

	return math.Round(largestSliceArea * 1e6) / 1e6
}

