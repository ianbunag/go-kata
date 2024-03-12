package search_a_2d_matrix

// Average time complexity: O(log m * n)
// Worst time complexity:   O(log m * n)
// Space complexity:        O(1)
func SearchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1

	for left <= right {
		middle := (left + right) / 2

		if lastIndex := len(matrix[middle]) - 1; matrix[middle][lastIndex] < target {
			left = middle + 1
		} else if matrix[middle][0] > target {
			right = middle - 1
		} else {
			return binarySearch(matrix[middle], target) > -1
		}
	}

	return false
}

// Average time complexity: O(log n)
// Worst time complexity:   O(log n)
// Space complexity:        O(1)
func binarySearch(nums []int, target int) int {
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
