package surrounded_regions

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(m)
func Solve(board [][]byte) {
	for row := range board {
		capture(board, row, 0)               // Left column.
		capture(board, row, len(board[0])-1) // Right column.
	}

	for column := range board[0] {
		if column == 0 || column == len(board[0])-1 {
			continue
		}

		capture(board, 0, column)            // Top row.
		capture(board, len(board)-1, column) // Bottom row.
	}

	for row := range board {
		for column := range board[row] {
			if board[row][column] == 'O' {
				board[row][column] = 'X'
			}
			if board[row][column] == 'T' {
				board[row][column] = 'O'
			}
		}
	}
}

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(m)
func capture(board [][]byte, row, column int) {
	captureStack := [][]int{{row, column}}

	for len(captureStack) > 0 {
		row, column, captureStack = captureStack[0][0], captureStack[0][1], captureStack[1:]

		if row < 0 ||
			row >= len(board) ||
			column < 0 ||
			column >= len(board[row]) ||
			board[row][column] != 'O' {
			continue
		}

		board[row][column], captureStack = 'T', append(
			captureStack,
			[]int{row - 1, column},
			[]int{row, column - 1},
			[]int{row, column + 1},
			[]int{row + 1, column},
		)
	}
}
