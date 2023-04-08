package HammingWeight

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func HammingWeight(num uint32) (weight int) {
	for num != 0 {
		if num&1 == 1 {
			weight += 1
		}

		num = num >> 1
	}

	return
}
