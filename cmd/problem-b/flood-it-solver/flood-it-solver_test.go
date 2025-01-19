package flooditsolver

import (
	"reflect"
	"testing"
)

func TestAddRow(t *testing.T) {
	tests := []struct {
		name         		string
		initialBoard        [][]int
		expectedSteps       int
		expectedColorCounts [6]int
	}{
		{
			name:         "Solving the given flood-it board using greedy algorithm",
			initialBoard: [][]int{
				{1,2,3,4,2,3},
				{3,3,4,5,2,1},
				{4,3,3,1,2,3},
				{5,4,3,6,2,1},
				{3,2,4,3,4,3},
				{2,3,4,1,5,6},},
				expectedSteps:       12,
				expectedColorCounts: [6]int{2, 2, 4, 2, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			floodIt := &FloodIt{Board: tt.initialBoard}
			totalSteps, chosenColorCounts := floodIt.SolveBoard()

			if totalSteps != tt.expectedSteps {
				t.Errorf("Expected %d steps, but got %d", tt.expectedSteps, totalSteps)
			}

			if !reflect.DeepEqual(chosenColorCounts, tt.expectedColorCounts) {
				t.Errorf("Expected color counts %v, but got %v", tt.expectedColorCounts, chosenColorCounts)
			}
		})
	}
}
