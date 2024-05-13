package max_area_of_island

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(1)
func MaxAreaOfIsland(grid [][]int) (maxArea int) {
	for row := range grid {
		for column := range grid[row] {
			if area := getArea(grid, row, column); area > maxArea {
				maxArea = area
			}
		}
	}

	return
}

func getArea(grid [][]int, row, column int) (area int) {
	if grid[row][column] == 0 {
		return 0
	}

	grid[row][column] = 0

	if row > 0 {
		area += getArea(grid, row-1, column)
	}
	if row < len(grid)-1 {
		area += getArea(grid, row+1, column)
	}
	if column > 0 {
		area += getArea(grid, row, column-1)
	}
	if column < len(grid[row])-1 {
		area += getArea(grid, row, column+1)
	}

	return area + 1
}
