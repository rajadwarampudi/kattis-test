package pizzaslicer

import (
	"math"
	"testing"
)

func TestCalculateMaxPossibleSliceArea(t *testing.T) {
	tests := []struct {
		name                  string
		radius                int
		noOfPeople            int
		angleDegrees          int
		angleMinutes          int
		angleSeconds          int
		expectedLargestSliceArea float64
	}{
		{
			name:                  "Given testcase 1-1",
			radius:                20,
			noOfPeople:            6,
			angleDegrees:          60,
			angleMinutes:          0,
			angleSeconds:          0,
			expectedLargestSliceArea: 209.439510,
		},
		{
			name:                  "Given testcase 1-2",
			radius:                20,
			noOfPeople:            6,
			angleDegrees:          59,
			angleMinutes:          59,
			angleSeconds:          59,
			expectedLargestSliceArea: 209.444358,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateMaxPossibleSliceArea(tt.radius, tt.noOfPeople, tt.angleDegrees, tt.angleMinutes, tt.angleSeconds)

			if math.Abs(result-tt.expectedLargestSliceArea) > 1e-6 {
				t.Errorf("Expected %v, but got %v", tt.expectedLargestSliceArea, result)
			}
		})
	}
}
