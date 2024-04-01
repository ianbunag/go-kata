package two_sum_ii

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func TwoSum(numbers []int, target int) []int {
	leftIndex, rightIndex := 0, len(numbers)-1

	for leftIndex != rightIndex {
		sum := numbers[leftIndex] + numbers[rightIndex]

		if sum == target {
			break
		}

		if sum-target > 0 {
			rightIndex -= 1
			continue
		}

		leftIndex += 1
	}

	return []int{leftIndex + 1, rightIndex + 1}
}
