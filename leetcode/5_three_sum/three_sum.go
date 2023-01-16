package three_sum

import (
	"sort"
)

// Average time complexity: O(n ^ 2)
// Worst time complexity:   O(n ^ 2)
// Space complexity:        O(n)
func GetThreeSum(numbers []int) [][]int {
	threeSums := [][]int{}

	sort.Ints(numbers)

	lastIndex := len(numbers) - 1

	for index, value := range numbers {
		if index > 0 && value == numbers[index-1] {
			continue
		}

		targetValue := -value
		left, right := index+1, lastIndex

		for left < right {
			leftValue, rightValue := numbers[left], numbers[right]
			sum := leftValue + rightValue

			if sum < targetValue {
				left += 1
				continue
			}

			if sum > targetValue {
				right -= 1
				continue
			}

			threeSums = append(threeSums, []int{value, leftValue, rightValue})
			left += 1
			right -= 1

			for left < right && numbers[left] == numbers[left-1] {
				left += 1
			}

			for left < right && numbers[right] == numbers[right+1] {
				right -= 1
			}
		}
	}

	return threeSums
}
