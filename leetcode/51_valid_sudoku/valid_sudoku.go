package valid_sudoku

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func IsValidSudoku(board [][]byte) bool {
	rows, columns, grids := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}

	for row := 0; row < 9; row += 1 {
		for column := 0; column < 9; column += 1 {
			if board[row][column] == '.' {
				continue
			}

			index, grid := board[row][column]-'1', ((row/3)*3)+column/3

			if rows[row][index] || columns[column][index] || grids[grid][index] {
				return false
			}

			rows[row][index], columns[column][index], grids[grid][index] = true, true, true
		}
	}

	return true
}
