package group_anagrams

// Average time complexity: O(n log n)
// Worst time complexity:   O(2n)
// Space complexity:        O(n)
func GroupAnagrams(strings []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, value := range strings {
		anagram := NewAnagram(value)
		groups[anagram.characterCounts] = append(
			groups[anagram.characterCounts],
			anagram.value,
		)
	}

	values := [][]string{}

	for _, group := range groups {
		values = append(values, group)
	}

	return values
}

// Average time complexity: O(n)
// Worst time complexity:   O(n)
// Space complexity:        O(n)
func NewAnagram(value string) Anagram {
	characterCounts := [26]int{}

	for _, character := range value {
		characterCounts[character-'a'] += 1
	}

	return Anagram{
		characterCounts: characterCounts,
		value:           value,
	}
}

type Anagram struct {
	characterCounts [26]int
	value           string
}
