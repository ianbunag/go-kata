package word_search

// Average time complexity: O(4^n)
// Worst time complexity:   O(4^n)
// Space complexity:        O(1)
func Exist(board [][]byte, word string) bool {
	included := make(map[[2]int]bool)

	for row := range board {
		for column := range board[row] {
			if wordExistsFromPosition(board, word, [2]int{row, column}, included) {
				return true
			}
		}
	}
	return false
}

// Average time complexity: O(4^n)
// Worst time complexity:   O(4^n)
// Space complexity:        O(1)
func wordExistsFromPosition(
	board [][]byte,
	word string,
	position [2]int,
	included map[[2]int]bool,
) bool {
	if word == "" {
		return true
	}

	row, column := position[0], position[1]

	if included[position] || // Cell is already part of the word
		row < 0 || // Row is out of top bound
		row >= len(board) || // Row is out of bottom bound
		column < 0 || // Column is out of left bound
		column >= len(board[row]) || // Column is out of right bound
		board[row][column] != word[0] { // Cell is not equal to first letter of word
		return false
	}

	nextWord := word[1:]
	included[position] = true

	if wordExistsFromPosition(board, nextWord, [2]int{row, column + 1}, included) || // Check right cell for succeeding word
		wordExistsFromPosition(board, nextWord, [2]int{row + 1, column}, included) || // Check bottom cell for succeeding word
		wordExistsFromPosition(board, nextWord, [2]int{row, column - 1}, included) || // Check left cell for succeeding word
		wordExistsFromPosition(board, nextWord, [2]int{row - 1, column}, included) { // Check top cell for succeeding word
		return true
	}

	// Allow cell to be checked again since it is not part of the word
	included[position] = false

	return false
}
