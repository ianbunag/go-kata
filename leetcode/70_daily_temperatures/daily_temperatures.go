package daily_temperatures

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(n)
func DailyTemperatures(temperatures []int) []int {
	waitDays, stack := make([]int, len(temperatures)), []int{}

	for index, temperature := range temperatures {
		for len(stack) > 0 {
			previousIndex := stack[len(stack)-1]

			if temperatures[previousIndex] >= temperature {
				break
			}

			waitDays[previousIndex], stack = index-previousIndex, stack[:len(stack)-1]
		}

		stack = append(stack, index)
	}

	return waitDays
}
