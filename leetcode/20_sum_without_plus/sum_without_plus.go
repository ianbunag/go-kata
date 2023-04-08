package sum_without_plus

// Average time complexity: O(log n)
// Worst time complexity:   O(log n)
// Space complexity:        O(1)
func GetSum(sum int, carry int) int {
	for carry != 0 {
		sum, carry = sum^carry, (sum&carry)<<1
	}

	return sum
}
