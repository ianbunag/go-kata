package missing_number

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(1)
func MissingNumber(numbers []int) int {
	// Initialize with total sum from 0 to n
	missingNumber := (len(numbers) * (len(numbers) + 1)) / 2

	for _, number := range numbers {
		missingNumber -= number
	}

	return missingNumber
}
