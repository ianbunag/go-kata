package valid_parenthesis_string

// Worst time complexity:   O(n)
// Average time complexity: O(n)
// Space complexity:        O(1)
func CheckValidString(s string) bool {
	leftMin, leftMax := 0, 0

	for _, char := range s {
		if char == '(' {
			leftMin, leftMax = leftMin+1, leftMax+1
		}

		if char == ')' {
			leftMin, leftMax = leftMin-1, leftMax-1
		}

		if char == '*' {
			leftMin, leftMax = leftMin-1, leftMax+1
		}

		if leftMax < 0 {
			return false
		}

		if leftMin < 0 {
			leftMin = 0
		}
	}

	return leftMin == 0
}
