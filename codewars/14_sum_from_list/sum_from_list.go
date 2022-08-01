package sum_from_list

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(n log n)
func TwoSum(numbers []int, target int) [2]int {
	differenceMap := make(map[int]int, len(numbers))

	for index, number := range numbers {
		differenceIndex, found := differenceMap[number]

		if found {
			if index > differenceIndex {
				return [2]int{differenceIndex, index}
			}

			return [2]int{index, differenceIndex}
		}

		// Store current difference
		differenceMap[target-number] = index
	}

	return [2]int{}
}
