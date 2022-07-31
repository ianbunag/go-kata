package valid_braces

type Brace [3]rune

var LeftBraces = Brace{'(', '{', '['}
var RightBraces = Brace{')', '}', ']'}

// Average time complexity: O(log n)
// Worst time complexity: 	O(n)
// Space complexity:				O(log n)
func ValidBraces(str string) bool {
	braceStack := []rune{}

	for _, value := range []rune(str) {
		leftIndex := braceIndexOf(LeftBraces, value)

		if leftIndex > -1 {
			braceStack = append(braceStack, value)
			continue
		}

		rightIndex := braceIndexOf(RightBraces, value)

		if len(braceStack) == 0 {
			return false
		}

		lastBrace := braceStack[len(braceStack)-1]

		if rightIndex > -1 && lastBrace != LeftBraces[rightIndex] {
			return false
		}

		braceStack = braceStack[:len(braceStack)-1]
	}

	return len(braceStack) == 0
}

// Average time complexity:O(log n)
// Worst time complexity: O(n)
// Space complexity: 			O(1)
func braceIndexOf(braces Brace, value rune) int {
	for index, brace := range braces {
		if brace == value {
			return index
		}
	}

	return -1
}
