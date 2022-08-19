package gold_mine

import (
	lib_slice "github.com/yvnbunag/go-kata/lib/slice"
)

// Average time complexity: O(n^2)
// Worst time complexity:   O(n^2)
// Space complexity:        O(n)
func FindMaximumAmount(field [][]int) int {
	if len(field) == 0 {
		return 0
	}

	rowLength := len(field)
	columnLength := len(field[0])
	previousColumn := make([]int, rowLength)
	nextColumn := make([]int, rowLength)

	for columnIndex := 0; columnIndex < columnLength; columnIndex += 1 {
		for rowIndex := 0; rowIndex < rowLength; rowIndex += 1 {
			topLeft, bottomLeft := 0, 0
			adjacentLeft := previousColumn[rowIndex]

			if rowIndex > 0 {
				topLeft = previousColumn[rowIndex-1]
			}

			if rowIndex < rowLength-1 {
				bottomLeft = previousColumn[rowIndex+1]
			}

			highestValue := lib_slice.Max([]int{topLeft, adjacentLeft, bottomLeft})
			nextColumn[rowIndex] = highestValue + field[rowIndex][columnIndex]
		}

		copy(previousColumn, nextColumn)
	}

	return lib_slice.Max(nextColumn)
}
