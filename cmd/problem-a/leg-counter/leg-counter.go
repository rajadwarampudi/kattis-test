package legcounter

import "fmt"

// PrintLegCounts calculates the number of each type of the three animals
// such that the sum of the total legs of all animals is equal to the given totalLegs
// and prints the result in "x y z" format.
// If the totalLegs is not possible to achive then it prints "impossible"
func PrintLegCounts(firstTypeAnimalLegs, secondTypeAnimalLegs, thirdTypeAnimalLegs, totalLegs int) {
	result := GenerateLegCountsResult(firstTypeAnimalLegs, secondTypeAnimalLegs, thirdTypeAnimalLegs, totalLegs)

	for _, combination := range result {
		fmt.Println(combination)
	}
}

func GenerateLegCountsResult(firstTypeAnimalLegs, secondTypeAnimalLegs, thirdTypeAnimalLegs, totalLegs int) []string {
	maxFirstAnimalCount := calculateMaxPossibleAnimalCount(totalLegs, firstTypeAnimalLegs)
	maxSecondAnimalCount := calculateMaxPossibleAnimalCount(totalLegs, secondTypeAnimalLegs)
	maxThirdAnimalCount := calculateMaxPossibleAnimalCount(totalLegs, thirdTypeAnimalLegs)

	isTotalLegCountPossible := false
	result := []string{}

	for firstCount := 0; firstCount <= maxFirstAnimalCount; firstCount++ {
		for secondCount := 0; secondCount <= maxSecondAnimalCount; secondCount++ {
			for thirdCount := 0; thirdCount <= maxThirdAnimalCount; thirdCount++ {
				if calculateAnimalTypeTotalLegs(firstCount, firstTypeAnimalLegs) +
				 calculateAnimalTypeTotalLegs(secondCount, secondTypeAnimalLegs) +
				 calculateAnimalTypeTotalLegs(thirdCount, thirdTypeAnimalLegs) == totalLegs {
					combination := fmt.Sprintf("%d %d %d", firstCount, secondCount, thirdCount)
					result = append(result, combination)
					isTotalLegCountPossible = true
				}
			}
		}
	}

	if !isTotalLegCountPossible {
		result = append(result, "impossible")
	}

	return result
}


func calculateMaxPossibleAnimalCount(totalLegs, animalLegs int) int {
	return totalLegs / animalLegs;
}

func calculateAnimalTypeTotalLegs(numberOfAnimals, legsForEachAnimal int) int {
	return numberOfAnimals * legsForEachAnimal
}