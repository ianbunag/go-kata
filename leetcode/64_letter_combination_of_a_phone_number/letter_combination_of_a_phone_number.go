package letter_combination_of_a_phone_number

var letters = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

// Worst time complexity:   O(4^n)
// Average time complexity: O(4^n)
// Space complexity:        O(4^n)
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	combinations := letters[digits[len(digits)-1]]

	for index := len(digits) - 2; index >= 0; index -= 1 {
		nextCombinations := []string{}

		for _, letterToCombine := range letters[digits[index]] {
			for _, combination := range combinations {
				nextCombinations = append(nextCombinations, letterToCombine+combination)
			}
		}

		combinations = nextCombinations
	}

	return combinations
}
