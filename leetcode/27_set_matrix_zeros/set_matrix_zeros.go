package set_matrix_zeros

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func SetZeroes(matrix [][]int) {
	firstRow, rows, columns := 1, len(matrix), len(matrix[0])

	for row := 0; row < rows; row += 1 {
		for column := 0; column < columns; column += 1 {
			if matrix[row][column] == 0 {
				if row == 0 {
					firstRow = 0
				} else {
					matrix[row][0] = 0
				}

				matrix[0][column] = 0
			}
		}
	}

	for row := 1; row < rows; row += 1 {
		for column := 1; column < columns; column += 1 {
			if matrix[row][0] == 0 || matrix[0][column] == 0 {
				matrix[row][column] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for row := 0; row < rows; row += 1 {
			matrix[row][0] = 0
		}
	}

	if firstRow == 0 {
		for column := 0; column < columns; column += 1 {
			matrix[0][column] = 0
		}
	}
}
