package move_zeros_to_end

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func MoveZeros(arr []int) []int {
	nonZeros := make([]int, len(arr))
	nonZeroIndex := 0

	for _, value := range arr {
		if value == 0 {
			continue
		}

		nonZeros[nonZeroIndex] = value
		nonZeroIndex += 1
	}

	return nonZeros
}
