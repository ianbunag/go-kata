package maximum_subarray_sum

// Average time complexity: O(n log n)
// Worst time complexity:   O(n log n)
// Space complexity:        O(n log n)
func MaximumSubarraySum(numbers []int) int {
	numbersLength := len(numbers)

	if numbersLength == 0 {
		return 0
	}

	if numbersLength == 1 {
		if numbers[0] > 0 {
			return numbers[0]
		}

		return 0
	}

	maximum := 0
	maxGenerations := numbersLength + 1
	grandParentGeneration := make([]int, maxGenerations)
	parentGeneration := numbers

	for generation := 2; generation < maxGenerations; generation += 1 {
		currentGeneration := make([]int, maxGenerations-generation)

		for childIndex := range currentGeneration {
			childValue := parentGeneration[childIndex] + parentGeneration[childIndex+1] - grandParentGeneration[childIndex+1]
			currentGeneration[childIndex] = childValue

			if childValue > maximum {
				maximum = childValue
			}
		}

		grandParentGeneration = parentGeneration
		parentGeneration = currentGeneration
	}

	return maximum
}
