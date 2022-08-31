package valid_parenthesis

// Average time complexity: O(log n)
// Worst time complexity:   O(n)
// Space complexity:        O(log n)
func IsValidParentheses(parenthesis string) bool {
	leftParenthesisStack := []rune{}

	for _, currentParenthesis := range parenthesis {
		if currentParenthesis == '(' {
			leftParenthesisStack = append(leftParenthesisStack, currentParenthesis)
			continue
		}

		stackLength := len(leftParenthesisStack)

		if stackLength < 1 {
			return false
		}

		leftParenthesisStack = leftParenthesisStack[:stackLength-1]
	}

	return len(leftParenthesisStack) == 0
}
