package plus_one

// Average time complexity: O(log n)
// Worst time complexity:   O(log n)
// Space complexity:        O(1)
func PlusOne(digits []int) []int {
	for index := len(digits) - 1; index >= 0; index -= 1 {
		if digits[index] == 9 {
			digits[index] = 0
			continue
		}

		digits[index] += 1
		return digits
	}

	// digits are all zero.
	digits = make([]int, len(digits)+1)
	digits[0] = 1

	return digits
}
