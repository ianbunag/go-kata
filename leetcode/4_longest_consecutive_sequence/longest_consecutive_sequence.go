package longest_consecutive_sequence

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func LongestConsecutive(numbers []int) int {
	if len(numbers) < 1 {
		return 0
	}

	numberMap := make(map[int]bool)

	for _, number := range numbers {
		numberMap[number] = true
	}

	longestSequence := 1

	for number := range numberMap {
		if numberMap[number-1] {
			continue
		}

		nextNumber := number + 1

		for numberMap[nextNumber] {
			nextNumber += 1
		}

		longestSequence = max(longestSequence, nextNumber-number)
	}

	return longestSequence
}

func max(first, second int) int {
	if first > second {
		return first
	}

	return second
}
