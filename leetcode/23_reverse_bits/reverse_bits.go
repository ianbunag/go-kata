package reverse_bits

// Average time complexity: O(1)
// Worst time complexity:   O(1)
// Space complexity:        O(1)
func ReverseBits(num uint32) uint32 {
	var result uint32
	remainingBits := 32

	for num != 0 {
		num, result, remainingBits = num>>1, (result<<1)|(num&1), remainingBits-1
	}

	return result << remainingBits
}
