package anagrams

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func Anagrams(word string, words []string) []string {
	anagrams := []string{}

	for _, entry := range words {
		if isAnagram(word, entry) {
			anagrams = append(anagrams, entry)
		}
	}

	return anagrams
}

// Average time complexity: O(n log n)
// Worst time complexity:   O(2n)
// Space complexity:        O(log n)
func isAnagram(first, second string) bool {
	characterMap := map[rune]int{}

	for _, character := range first {
		characterMap[character] += 1
	}

	for _, character := range second {
		if characterMap[character] == 0 {
			return false
		}

		characterMap[character] -= 1
	}

	return true
}
