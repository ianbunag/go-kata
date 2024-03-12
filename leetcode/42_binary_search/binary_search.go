package binary_search

// Average time complexity: O(log n)
// Worst time complexity:   O(log n)
// Space complexity:        O(1)
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2

		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			return middle
		}
	}

	return -1
}
