package find_odd_int

// Average time complexity: O(n log n)
// Worst time complexity:   O(n log n)
// Space complexity:        O(n log n)
func FindOdd(seq []int) int {
	occurrences := map[int]bool{}

	for _, value := range seq {
		occurrences[value] = !occurrences[value]
	}

	for value, isOdd := range occurrences {
		if isOdd {
			return value
		}
	}

	panic("odd occurring number not found")
}
