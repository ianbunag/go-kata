package longest_increasing_path_in_matrix

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(n)
func LongestIncreasingPath(matrix [][]int) (longest int) {
	paths := make([][]int, len(matrix))

	for row := range paths {
		paths[row] = make([]int, len(matrix[row]))
	}

	for row := range matrix {
		for column := range matrix[row] {
			longest = max(longest, Traverse(paths, matrix, row, column, -1))
		}
	}

	return longest
}

func Traverse(paths, matrix [][]int, row, column, previousValue int) int {
	if row < 0 ||
		row >= len(matrix) ||
		column < 0 ||
		column >= len(matrix[row]) ||
		matrix[row][column] <= previousValue {
		return 0
	}

	if paths[row][column] != 0 { // Path is visited, return recorded value.
		return paths[row][column]
	}

	currentValue := matrix[row][column]
	paths[row][column] = 1
	paths[row][column] = maxInt(paths[row][column], Traverse(paths, matrix, row, column-1, currentValue)+1)
	paths[row][column] = maxInt(paths[row][column], Traverse(paths, matrix, row, column+1, currentValue)+1)
	paths[row][column] = maxInt(paths[row][column], Traverse(paths, matrix, row-1, column, currentValue)+1)
	paths[row][column] = maxInt(paths[row][column], Traverse(paths, matrix, row+1, column, currentValue)+1)

	return paths[row][column]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}
