package snail

// Average time complexity: O(log n)
// Worst time complexity:   O(log n)
// Space complexity:        O(n)
func Snail(snailMap [][]int) []int {
	trail := []int{}

	if len(snailMap[0]) == 0 {
		return trail
	}

	min := 0
	max := len(snailMap)

	for {
		trail = append(
			trail,
			snailMap[min][min:max]...,
		)
		trail = append(
			trail,
			verticalSlice(snailMap, Point{min + 1, max - 1}, Point{max - 1, max - 1})...,
		)

		max -= 1
		min += 1
		if min > max {
			break
		}

		trail = append(
			trail,
			reverse(snailMap[max][min-1:max])...,
		)
		trail = append(
			trail,
			reverse(verticalSlice(snailMap, Point{min, min - 1}, Point{max - 1, min - 1}))...,
		)
	}

	return trail
}

type Point struct {
	x int
	y int
}

func verticalSlice(matrix [][]int, from, to Point) []int {
	result := []int{}

	for fromX := from.x; fromX <= to.x; fromX += 1 {
		result = append(result, matrix[fromX][from.y])
	}

	return result
}

func reverse[SliceType ~[]Element, Element any](source SliceType) SliceType {
	reversed := make(SliceType, len(source))

	copy(reversed, source)

	for i, j := 0, len(source)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}

	return reversed
}
