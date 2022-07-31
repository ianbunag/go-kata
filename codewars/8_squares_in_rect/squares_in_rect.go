package squares_in_rect

// Average time complexity: O(n)
// Worst time complexity: 	O(n)
// Space complexity: 				O(n)
func SquaresInRect(length int, width int) (result []int) {
	if length == width {
		return
	}

	for length != width {
		if length < width {
			length, width = width, length
		}

		result = append(result, width)
		length = length - width
	}

	result = append(result, width)
	return
}
