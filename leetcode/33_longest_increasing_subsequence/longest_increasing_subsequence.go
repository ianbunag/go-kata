package longest_increasing_subsequence

import "math"

// Average time complexity: O(n ^ 2)
// Worst time complexity:   O(n ^ 2)
// Space complexity:        O(n)
func LengthOfLIS(nums []int) int {
	length := len(nums)
	results := make([]int, length)

	for current := length - 1; current >= 0; current -= 1 {
		results[current] = 1

		for next := current + 1; next < length; next += 1 {
			if nums[current] < nums[next] {
				results[current] = max(results[current], results[next]+1)
			}
		}
	}

	return max(results...)
}

func max(values ...int) int {
	result := math.MinInt

	for _, value := range values {
		if value > result {
			result = value
		}
	}

	return result
}
