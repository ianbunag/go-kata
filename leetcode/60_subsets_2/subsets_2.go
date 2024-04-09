package subsets_2

import (
	"sort"
)

// Worst time complexity:   O(2^n)
// Average time complexity: O(2^n)
// Space complexity:        O(n)
func SubsetsWithDup(nums []int) (subsets [][]int) {
	sort.Ints(nums)

	var backtrack func(current, previous int, subset []int)
	backtrack = func(current, previous int, subset []int) {
		if current == len(nums) {
			subsets = append(subsets, append([]int{}, subset...))
			return
		}

		// Include current value in subset.
		subset = append(subset, nums[current])
		backtrack(current+1, current, subset)
		subset = subset[:len(subset)-1]

		// Skip subset if previous is equal to current value.
		if previous >= 0 && nums[previous] == nums[current] {
			return
		}

		// Retain current subset.
		backtrack(current+1, -1, subset) // Always keep current subset.
	}

	backtrack(0, -1, []int{})

	return
}
