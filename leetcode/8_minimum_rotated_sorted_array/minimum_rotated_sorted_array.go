package minimum_rotated_sorted_array

// Average time complexity: O(log n)
// Worst time complexity:   O(log n)
// Space complexity:        O(1)
func FindMin(values []int) int {
	left, right := 0, len(values)-1

	for {
		if values[left] < values[right] || left == right {
			return values[left]
		}

		middle := (left + right) / 2

		if values[middle] >= values[left] {
			left = middle + 1
			continue
		}

		right = middle
	}
}
