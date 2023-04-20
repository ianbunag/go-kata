package palindromic_substrings

type Parity int

const PARITY_ODD Parity = 0
const PARITY_EVEN Parity = 1

// Average time complexity: O(n ^ 2)
// Worst time complexity:   O(n ^ 2)
// Space complexity:        O(1)
func CountSubstrings(value string) (palindromes int) {
	palindromes += countSubstringsWithParity(value, PARITY_ODD)
	palindromes += countSubstringsWithParity(value, PARITY_EVEN)

	return
}

func countSubstringsWithParity(value string, parity Parity) (palindromes int) {
	for index := 0; index < len(value); index += 1 {
		left, right := index, index+int(parity)

		for left >= 0 && right < len(value) && value[left] == value[right] {
			left -= 1
			right += 1
		}

		palindromes += index - left
	}

	return palindromes
}
