package house_robber_2

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func Rob(houses []int) int {
	if len(houses) == 1 {
		return houses[0]
	}

	robMax := func(houses []int) int {
		highest, previous := 0, 0

		for _, current := range houses {
			highest, previous = max(highest, previous+current), highest
		}

		return highest
	}

	return max(
		robMax(houses[1:]),
		robMax(houses[:len(houses)-1]),
	)
}

func max(first, second int) int {
	if first > second {
		return first
	}

	return second
}
