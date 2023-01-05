package product_of_array_except_self

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func ProductExceptSelf(numbers []int) []int {
	numbersLength := len(numbers)
	output := make([]int, numbersLength)
	output[0] = 1

	for prefixIndex := 1; prefixIndex < numbersLength; prefixIndex += 1 {
		output[prefixIndex] = numbers[prefixIndex-1] * output[prefixIndex-1]
	}

	postfix := numbers[numbersLength-1]
	for postfixIndex := numbersLength - 2; postfixIndex >= 0; postfixIndex -= 1 {
		output[postfixIndex] *= postfix
		postfix *= numbers[postfixIndex]
	}

	return output
}
