package valid_anagram

// Average time complexity: O(n log n)
// Worst time complexity:   O(n log n)
// Space complexity:        O(n log n)
func IsAnagram(source string, target string) bool {
	if len(source) != len(target) {
		return false
	}

	characters := make(map[rune]int)

	for _, character := range source {
		characters[character] += 1
	}

	for _, character := range target {
		characters[character] -= 1
	}

	for _, count := range characters {
		if count != 0 {
			return false
		}
	}

	return true
}
