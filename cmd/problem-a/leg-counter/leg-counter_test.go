package legcounter

import (
	"reflect"
	"testing"
)

func TestGenerateLegCountsResult(t *testing.T) {
	tests := []struct {
		name                   string
		firstTypeAnimalLegs    int
		secondTypeAnimalLegs   int
		thirdTypeAnimalLegs    int
		totalLegs              int
		expectedResult         []string
	}{
		{
			name:                 "Basic case with possible combinations",
			firstTypeAnimalLegs:  2,
			secondTypeAnimalLegs: 4,
			thirdTypeAnimalLegs:  4,
			totalLegs:            14,
			expectedResult:       []string{"1 0 3", "1 1 2", "1 2 1", "1 3 0", "3 0 2", "3 1 1", "3 2 0", "5 0 1", "5 1 0", "7 0 0"},
		},
		{
			name:                 "Impossible case",
			firstTypeAnimalLegs:  2,
			secondTypeAnimalLegs: 3,
			thirdTypeAnimalLegs:  5,
			totalLegs:            1,
			expectedResult:       []string{"impossible"},
		},
		{
			name:                 "Zero legs",
			firstTypeAnimalLegs:  2,
			secondTypeAnimalLegs: 4,
			thirdTypeAnimalLegs:  6,
			totalLegs:            0,
			expectedResult:       []string{"0 0 0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateLegCountsResult(tt.firstTypeAnimalLegs, tt.secondTypeAnimalLegs, tt.thirdTypeAnimalLegs, tt.totalLegs)
			if !reflect.DeepEqual(result, tt.expectedResult) {
				t.Errorf("Expected %v, but got %v", tt.expectedResult, result)
			}
		})
	}
}
