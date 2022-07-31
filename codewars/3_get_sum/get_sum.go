package get_sum

// Average time complexity: O(1)
// Worst time complexity: 	O(1)
// Space complexity: 				O(1)
func GetSum(first, second int) int {
	if first == second {
		return first
	}

	if first < second {
		first, second = second, first
	}

	/*
		Carl Gauss' formula for adding consecutive numbers
		https://study.com/academy/lesson/finding-the-sum-of-consecutive-numbers.html
	*/
	return (first + second) * (first - second + 1) / 2
}
