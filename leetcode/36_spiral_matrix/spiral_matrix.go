package spiral_matrix

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func SpiralOrder(matrix [][]int) []int {
	order := make([]int, 0, len(matrix)*len(matrix[0]))
	top, left, bottom, right := 0, 0, len(matrix)-1, len(matrix[0])-1
	isWithinBounds := func() bool {
		return top <= bottom && left <= right
	}

	for isWithinBounds() {
		for topIndex := left; topIndex <= right; topIndex += 1 {
			order = append(order, matrix[top][topIndex])
		}
		top += 1

		for rightIndex := top; rightIndex <= bottom; rightIndex += 1 {
			order = append(order, matrix[rightIndex][right])
		}
		right -= 1

		if !isWithinBounds() {
			break
		}

		for bottomIndex := right; bottomIndex >= left; bottomIndex -= 1 {
			order = append(order, matrix[bottom][bottomIndex])
		}
		bottom -= 1

		for leftIndex := bottom; leftIndex >= top; leftIndex -= 1 {
			order = append(order, matrix[leftIndex][left])
		}
		left += 1
	}

	return order
}
