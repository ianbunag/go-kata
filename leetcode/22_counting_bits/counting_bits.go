package counting_bits

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func CountBits(n int) []int {
	bits := make([]int, n+1)
	offset := 1

	for index := 1; index < len(bits); index += 1 {
		if offset*2 == index {
			offset *= 2
		}

		bits[index] = 1 + bits[index-offset]
	}

	return bits
}
